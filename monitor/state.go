// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

import (
	"fmt"
	"time"
)

// State represents the state of a Pollable reported by a given Poller (id)
// start is the time when the Pollable entered this state, npolls is the number
// of consecutive polls returning the same health status since the previous
// state change
type State struct {
	ID     int
	Health Health
	Detail string
	Npolls int
	Start  time.Time
}

// NewState constructs and initializes a Pollable State
func NewState(pollerID int, health Health, detail string, npolls int) *State {
	return &State{ID: pollerID, Health: health, Detail: detail, Npolls: npolls, Start: time.Now()}
}

func (state State) incrementPollCount() {
	state.Npolls = state.Npolls + 1
}

// String returns a string representation of this State suitable for printing
func (state State) String() string {
	return fmt.Sprintf("State: %s - %s, npolls: %d, start %s", state.Health, state.Detail, state.Npolls, state.Start.String())
}
