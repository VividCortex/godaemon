package godaemon

// Copyright (c) 2013 VividCortex, Inc. All rights reserved.
// Please see the LICENSE file for applicable license terms.

import (
	"io"
	"os"
)

// A DaemonAttr describes the options that apply to daemonization
type DaemonAttr struct {
	CaptureOutput bool // whether to capture stdout/stderr
}

// MakeDaemon is a no-op on MacOSX Darwin.
func MakeDaemon(attrs *DaemonAttr) (io.Reader, io.Reader) {
	var stdout, stderr *os.File

	if attrs.CaptureOutput {
		stdout = os.NewFile(uintptr(1), "stdout")
		stderr = os.NewFile(uintptr(2), "stderr")
	}
	return stdout, stderr
}

// Daemonize is equivalent to MakeDaemon(&DaemonAttr{}). It is kept only for
// backwards API compatibility, but its usage is otherwise discouraged. Use
//
// Daemonize is a no-op on MacOSX Darwin.
func Daemonize(child ...bool) {
}
