// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package service

import (
	"fmt"

	"github.com/rmorriso/gomon/monitor"
)

// File represents a  file to be checked for a change in its hash
// since the last check
// File implements the pollable interface
type File struct {
	Path         string
	PollInterval int64
	hash         string
}

func NewFile(path string) *File {
	return &File{Path: path}
}

// Poll computes a hash on the contents of the file, returning health status,
// the hash as the detail string of the Poll and possibly error;
// if the hash cannot be computed the empty string is returned with false representing failure
func (file *File) Poll() (monitor.Health, string, error) {
	hash, err := getHash(file.Path)
	if err != nil {
		return monitor.Critical, err.Error(), err
	}
	if file.hash != hash {
		if file.hash == "" {
			file.hash = hash
			return monitor.Ok, fmt.Sprintf("file hash: %s", file.hash), nil
		}
		file.hash = hash
		return monitor.Warning, fmt.Sprintf("file hash changed: %s", file.hash), nil
	}
	return monitor.Ok, fmt.Sprintf("file hash: %s", file.hash), nil
}

func (file *File) Id() string {
	return file.Path
}

func (file *File) Interval() int64 {
	return file.PollInterval
}

func (file *File) String() string {
	return "[File:" + file.Id() + "]"
}
