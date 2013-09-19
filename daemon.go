package godaemon

// Copyright (c) 2013 VividCortex, Inc. All rights reserved.
// Please see the LICENSE file for applicable license terms.

// A DaemonAttr describes the options that apply to daemonization
type DaemonAttr struct {
	CaptureOutput bool // whether to capture stdout/stderr
}
