package godaemon

// Copyright (c) 2013 VividCortex, Inc. All rights reserved.
// Please see the LICENSE file for applicable license terms.

import (
	"os"
)

// GetExecutablePath returns the absolute path to the currently running
// executable.  It is used internally by the godaemon package, and exported
// publicly because it's useful outside of the package too.
func GetExecutablePath() (string, error) {
	return os.Executable()
}
