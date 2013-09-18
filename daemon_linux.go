// Package godaemon runs a program as a Unix daemon.
package godaemon

// Copyright (c) 2013 VividCortex, Inc. All rights reserved.
// Please see the LICENSE file for applicable license terms.

import (
	"fmt"
	"os"
)

func getExecutablePath() (string, error) {
	exePath, err := os.Readlink("/proc/self/exe")

	if err != nil {
		err = fmt.Errorf("can't read /proc/self/exe: %v", err)
	}

	return exePath, err
}
