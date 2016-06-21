// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package check

import (
	"fmt"
	"strings"

	"github.com/codeskyblue/go-sh"
	"github.com/rmorriso/gomon/monitor"
)

// ShellCheck represents an shell script or binary to be executed and the
// interval between executions in seconds
// the shell script is expected to follow the Nagio convention for return codes:
// 0 = OK, 1 = WARN, 2 = ERROR
// TODO: implement backoff

// ShellCheck implements the pollable interface
type ShellCheck struct {
	Cmd          string
	Args         string
	PollInterval int64
}

// NewShellCheck constructs and initializes a ShellCheck pollable
func NewShellCheck(cmd, args string) *ShellCheck {
	return &ShellCheck{Cmd: cmd, Args: args}
}

// Poll executes a ShellCheck and returns the output of the execution;
func (shellCheck *ShellCheck) Poll() (monitor.Health, string, error) {
	session := sh.NewSession()
	out, err := session.Command(shellCheck.Cmd, getargs(shellCheck.Args)).Output()
	if err != nil {
		return monitor.Critical, err.Error(), nil
	}
	return monitor.Ok, string(out), nil
}

// ID returns the ID string associated with the ShellCheck
func (shellCheck *ShellCheck) ID() string {
	return shellCheck.Cmd
}

// Interval returns the polling interval
func (shellCheck *ShellCheck) Interval() int64 {
	return shellCheck.PollInterval
}

// String returns a printable description of the ShellCheck
func (shellCheck *ShellCheck) String() string {
	return "[ShellCheck:" + shellCheck.ID() + "]"
}

func getargs(format string, a ...interface{}) []string {
	s := fmt.Sprintf(format, a...)
	return strings.Split(s, " ")
}
