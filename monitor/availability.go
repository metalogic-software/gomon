// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

// Availability defines allowed values up and down
type Availability int

const (
	up Availability = iota
	down
)

var availabilityName = []string{"Up", "Down"}

func (availability Availability) String() string {
	return availabilityName[availability]
}
