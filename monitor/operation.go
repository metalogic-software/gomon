// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

// Operation defines allowable operations that may be executed on
// a Poller: enable logging, pause, run or terminate the poller
type Operation int

const (
	logging Operation = iota
	pause
	run
	terminate
)

var operationName = []string{"logging", "pause", "run", "terminate"}

// String() returns the string representation of the Operation
func (operation Operation) String() string {
	return operationName[operation]
}
