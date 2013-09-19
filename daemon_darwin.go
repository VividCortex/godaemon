package godaemon

import (
	"io"
)

// Copyright (c) 2013 VividCortex, Inc. All rights reserved.
// Please see the LICENSE file for applicable license terms.

// A DaemonAttr describes the options that apply to daemonization
type DaemonAttr struct {
	CaptureOutput bool // whether to capture stdout/stderr
}

// Daemonize is a no-op on MacOSX Darwin.
func Daemonize(child bool) {
}

// DaemonizeWithCapture is a no-op on MacOSX Darwin.
func DaemonizeWithCapture(child bool) (io.Reader, io.Reader) {
	return nil, nil
}

// MakeDaemon is a no-op on MacOSX Darwin.
func MakeDaemon(attrs *DaemonAttr) (io.Reader, io.Reader) {
	return nil, nil
}
