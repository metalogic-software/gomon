// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

// the operations defined on a poller
type Operation int

const (
	Logging Operation = iota
	Pause
	Run
	Terminate
)

var operationName []string = []string{"logging", "pause", "run", "terminate"}

// String() returns the string representation of the Operation
func (operation Operation) String() string {
	return operationName[operation]
}
