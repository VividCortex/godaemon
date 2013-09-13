package godaemon

// Copyright (c) 2013 VividCortex, Inc. All rights reserved.
// Please see the LICENSE file for applicable license terms.

// Daemonize is a no-op on MacOSX Darwin.
func Daemonize(child bool) {
}

// DaemonizeWithCapture is a no-op on MacOSX Darwin.
func DaemonizeWithCapture(child bool) (io.Reader, io.Reader) {
}
