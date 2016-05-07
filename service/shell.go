// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package service

import (
"github.com/codeskyblue/go-sh"
 "github.com/rmorriso/gomon/monitor"
)

// ShellService represents a program execution to be polled and the
// interval between polls in seconds

// ShellService implements the pollable interface
type HTTPService struct {
	Command          string
	PollInterval int64
}

// NewShellService constructs and initializes a ShellService pollable
func NewShellService(command string) *HTTPService {
	return &HTTPService{Command: command}
}

// Poll executes command and returns health status and a detail string;
func (svc *ShellService) Poll() (monitor.Health, string, error) {
	resp, err := http.Head(svc.URL)
	if err != nil {
		return monitor.Critical, err.Error(), nil
	}
	return monitor.Ok, resp.Status, nil
}

// ID returns the ID string associated with the HTTPService
func (svc *HTTPService) ID() string {
	return svc.URL
}

// Interval returns the polling interval
func (svc *HTTPService) Interval() int64 {
	return svc.PollInterval
}

// String returns a printable description of the HTTPService
func (svc *HTTPService) String() string {
	return "[HTTPService:" + svc.ID() + "]"
}
