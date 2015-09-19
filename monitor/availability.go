// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file
package monitor

type Availability int

const (
	Up Availability = iota
	Down
)

var availabilityName []string = []string{"Up", "Down"}

func (availability Availability) String() string {
	return availabilityName[availability]
}
