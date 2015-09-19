// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

type Health int

const (
	Ok Health = iota
	Warning
	Critical
	Unknown
)

var healthName []string = []string{"OK", "Warning", "Critical", "Unknown"}

func (health Health) String() string {
	return healthName[health]
}
