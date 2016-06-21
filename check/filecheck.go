// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package check

import (
	"fmt"

	"github.com/rmorriso/gomon/monitor"
)

// FileCheck represents a  file to be checked for a change in its hash
// since the last check
// FileCheck implements the pollable interface
type FileCheck struct {
	Path         string
	PollInterval int64
	hash         string
}

// NewFileCheck constructs and initializes a FileCheck Pollable
func NewFileCheck(path string) *FileCheck {
	return &FileCheck{Path: path}
}

// Poll computes a hash on the contents of the file, returning health status,
// the hash as the detail string of the Poll and possibly error;
// if the hash cannot be computed the empty string is returned with false representing failure
func (fileCheck *FileCheck) Poll() (monitor.Health, string, error) {
	hash, err := getHash(fileCheck.Path)
	if err != nil {
		return monitor.Critical, err.Error(), err
	}
	if fileCheck.hash != hash {
		if fileCheck.hash == "" {
			fileCheck.hash = hash
			return monitor.Ok, fmt.Sprintf("file hash: %s", fileCheck.hash), nil
		}
		fileCheck.hash = hash
		return monitor.Warning, fmt.Sprintf("file hash changed: %s", fileCheck.hash), nil
	}
	return monitor.Ok, fmt.Sprintf("file hash: %s", fileCheck.hash), nil
}

// ID returns the ID of this FileCheck pollable
func (fileCheck *FileCheck) ID() string {
	return fileCheck.Path
}

// Interval returns the polling Interval (in sec) of this FileCheck pollable
func (fileCheck *FileCheck) Interval() int64 {
	return fileCheck.PollInterval
}

// String returns a string representation of this FileCheck pollable suitable for printing
func (fileCheck *FileCheck) String() string {
	return "[File:" + fileCheck.ID() + "]"
}
