// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

import (
	"fmt"
	"log"
	"time"
)

const (
	second          = 1e9         // one second is 1e9 nanoseconds
	errTimeout      = 10 * second // back-off timeout on error
	defaultInterval = 60 * second // how often to poll each Resource
)

// Poller controls execution of a Pollable and records its history
type Poller struct {
	id       int
	pollable Pollable
	nerrs    int
	history  []*State
	running  bool
	logging  bool
	ticker   *time.Ticker
	control  chan Operation
}

// NewPoller constructs and initializes a Poller
func NewPoller(id int, pollable Pollable) *Poller {
	return &Poller{
		id:       id,
		pollable: pollable,
		nerrs:    0,
		history:  []*State{NewState(id, Unknown, "initializing poller", 0)},
		running:  true,
		ticker:   time.NewTicker(time.Duration(pollable.Interval() * second)),
		control:  make(chan Operation),
	}
}

// Exec executes the Poller run loop in a goroutine
func (poller *Poller) Exec() {
	go func() {
		for {
			select {
			case <-poller.ticker.C:
				if poller.running { // TODO: handle error
					health, detail, _ := poller.pollable.Poll()
					poller.updateHistory(health, detail)
				}
			case operation := <-poller.control:
				log.Printf("control operation: %s", operation)
				switch operation {
				case run:
					poller.running = true
				case pause:
					poller.running = false
				case terminate:
					return
				case logging:
					poller.logging = !poller.logging
					fmt.Println("set logging:", poller.logging)
				}
			}
		}
	}()
}

// Run starts a paused Poller
func (poller *Poller) Run() {
	poller.control <- run
}

// Pause pauses a running Poller
func (poller *Poller) Pause() {
	poller.control <- pause
}

// Terminate terminates a Poller
func (poller *Poller) Terminate() {
	poller.control <- terminate
}

// Log toggles Poller logging off and on
func (poller *Poller) Log(enable bool) {
	if poller.logging == enable {
		return
	}
	poller.control <- logging
}

// ID returns the ID of this Poller
func (poller *Poller) ID() int {
	return poller.id
}

// History returns the state history
func (poller *Poller) History() []*State {
	return poller.history
}

func (poller *Poller) updateHistory(health Health, detail string) {
	current := poller.history[len(poller.history)-1]

	if current.health == health { //health status unchanged
		current.incrementPollCount()
	} else { // health status changed
		newState := NewState(poller.id, health, detail, 1)
		poller.history = append(poller.history, newState)
	}
}

// Pollable returns the Pollable controlled by this Poller
func (poller *Poller) Pollable() Pollable {
	return poller.pollable
}

// String returns a string representation of the Poller suitable for printing
func (poller *Poller) String() string {
	return fmt.Sprintf("Poller[%d]: %s", poller.id, poller.pollable.String())
}
