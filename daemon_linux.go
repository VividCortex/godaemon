// Package godaemon runs a program as a Unix daemon.
package godaemon

// Copyright (c) 2013 VividCortex, Inc. All rights reserved.
// Please see the LICENSE file for applicable license terms.

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// The name of the env var that indicates what stage of daemonization we're at.
const stageVar = "__DAEMON_STAGE"

/*
Daemonize turns the process into a daemon. But given the lack of Go's
support for fork(), Daemonize() is forced to run the process all over again,
from the start. Hence, this should probably be your first call after main
begins, unless you understand the effects of calling from somewhere else.
Keep in mind that the PID changes after this function is called, given
that it only returns in the child; the parent will exit without returning.

The child parameter, which defaults to false, indicates whether children
should also Daemonize(). If so, then there will be extra "forks". If not,
then the environment variable stageVar is reserved for the use of this
package. Otherwise it will be restored to its original value.

Daemonizing is a 3-stage process. In stage 0, the program increments the
magical environment variable and starts a copy of itself that's a session
leader, with its STDIN, STDOUT, and STDERR disconnected from any tty. It
then exits.

In stage 1, the (new copy of) the program starts another copy that's not
a session leader, and then exits.

In stage 2, the (new copy of) the program chdir's to /, then sets the umask.
If child is true, it returns the magical variable to its original value.
*/
func Daemonize(child ...bool) {
	stage, advanceStage, resetEnv := getStage()

	if stage == 0 || stage == 1 {
		procName, err := os.Readlink("/proc/self/exe")

		if err != nil {
			fmt.Fprintf(os.Stderr, "can't read /proc/self/exe:", err, "\n")
			os.Exit(1)
		}

		advanceStage()
		dir, _ := os.Getwd()
		attrs := os.ProcAttr{Dir: dir, Env: os.Environ(), Files: []*os.File{nil, nil, nil}}

		if stage == 0 {
			sysattrs := syscall.SysProcAttr{Setsid: true}
			attrs.Sys = &sysattrs
		}

		proc, err := os.StartProcess(procName, os.Args, &attrs)

		if err != nil {
			fmt.Fprintf(os.Stderr, "can't create process %s\n", procName)
			os.Exit(1)
		}

		proc.Release()
		os.Exit(0)
	}

	os.Chdir("/")
	syscall.Umask(0)

	if len(child) == 0 || !child[0] {
		resetEnv()
	}
}

// Returns the current stage in the "daemonization process", that's kept in
// an environment variable. The variable is instrumented with a digital
// signature, to avoid misbehavior if it was present in the user's
// environment. The original value is restored after the last stage, so that
// there's no final effect on the environment the application receives.
func getStage() (stage int, advanceStage func(), resetEnv func()) {
	var origValue string
	stage = 0

	daemonStage := os.Getenv(stageVar)
	stageTag := strings.SplitN(daemonStage, ":", 2)
	stageInfo := strings.SplitN(stageTag[0], "/", 3)

	if len(stageInfo) == 3 {
		stageStr, tm, check := stageInfo[0], stageInfo[1], stageInfo[2]

		hash := sha1.New()
		hash.Write([]byte(stageStr + "/" + tm + "/"))

		if check != hex.EncodeToString(hash.Sum([]byte{})) {
			// This whole chunk is original data
			origValue = daemonStage
		} else {
			stage, _ = strconv.Atoi(stageStr)

			if len(stageTag) == 2 {
				origValue = stageTag[1]
			}
		}
	} else {
		origValue = daemonStage
	}

	advanceStage = func() {
		base := fmt.Sprintf("%d/%09d/", stage+1, time.Now().Nanosecond())
		hash := sha1.New()
		hash.Write([]byte(base))

		tag := base + hex.EncodeToString(hash.Sum([]byte{}))

		if err := os.Setenv(stageVar, tag+":"+origValue); err != nil {
			fmt.Fprintf(os.Stderr, "can't set %s (stage %d)\n", stageVar, stage)
			os.Exit(1)
		}
	}

	resetEnv = func() {
		if err := os.Setenv(stageVar, origValue); err != nil {
			fmt.Fprintf(os.Stderr, "can't reset %s\n", stageVar)
			os.Exit(1)
		}
	}

	return stage, advanceStage, resetEnv
}
