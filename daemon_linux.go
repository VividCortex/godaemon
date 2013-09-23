// Package godaemon runs a program as a Unix daemon.
package godaemon

// Copyright (c) 2013 VividCortex, Inc. All rights reserved.
// Please see the LICENSE file for applicable license terms.

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
 * This returns the absolute path to the currently running executable.
 *
 * It is used internally by the godaemon package.
 * It may also be used elsewhere in the VividCortex codebase.
 */
func GetExecutablePath() (string, error) {
	exePath, err := os.Readlink("/proc/self/exe")

	if err != nil {
		err = fmt.Errorf("can't read /proc/self/exe: %v", err)
	}

	return filepath.Clean(exePath), err
}
