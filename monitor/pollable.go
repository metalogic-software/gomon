// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

// an interface for pollable objects
package monitor

type Pollable interface {
	Id() string
	Poll() (Health, string, error)
	Interval() int64
	String() string
}
