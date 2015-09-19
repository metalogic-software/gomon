// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package service

import "fmt"
import "github.com/rmorriso/gomon/monitor"

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

// Poll computes a hash on the contents of the file, returning health status, the
// hash as the detail string and posibly an os.Error if an error occurred when polling
// if the hash cannot be computed the empty string is returned with false representing failure
func (this *File) Poll() (monitor.Health, string, error) {
	hash, err := getHash(this.Path)
	if err != nil {
		return monitor.Critical, err.Error(), err
	}
	if this.hash != hash {
		if this.hash == "" {
			this.hash = hash
			return monitor.Ok, fmt.Sprintf("file hash: %s", this.hash), nil
		}
		this.hash = hash
		return monitor.Warning, fmt.Sprintf("file hash changed: %s", this.hash), nil
	}
	return monitor.Ok, fmt.Sprintf("file hash: %s", this.hash), nil
}

func (this *File) Id() string {
	return this.Path
}

func (this *File) Interval() int64 {
	return this.PollInterval
}

func (this *File) String() string {
	return "[File:" + this.Id() + "]"
}
