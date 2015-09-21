// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

// Operation defines allowable values on a Poller logging, pause, run, terminate
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
