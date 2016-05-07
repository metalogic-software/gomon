// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

// Pollable defines the pollable interface;
type Pollable interface {
	ID() string
	Poll() (Health, string, error)
	Interval() int64
	String() string
}
