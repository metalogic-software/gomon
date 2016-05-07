// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

// Health defines allowed values Ok, Warning, Critical and Unknown
type Health int

// ok = 0, warning = 1, critical = 2, unknown = 3
const (
	Ok Health = iota
	Warning
	Critical
	Unknown
)

var healthName = []string{"OK", "Warning", "Critical", "Unknown"}

func (health Health) String() string {
	return healthName[health]
}
