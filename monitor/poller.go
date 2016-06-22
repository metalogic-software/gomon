// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

import (
	"encoding/json"
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
	ID       int
	Pollable Pollable
	Nerrs    int
	History  []*State
	Running  bool
	Logging  bool
	ticker   *time.Ticker
	control  chan Operation
}

// NewPoller constructs and initializes a Poller
func NewPoller(id int, pollable Pollable) *Poller {
	return &Poller{
		ID:       id,
		Pollable: pollable,
		Nerrs:    0,
		History:  []*State{NewState(id, Unknown, "initializing poller", 0)},
		Running:  true,
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
				if poller.Running { // TODO: handle error
					health, detail, _ := poller.Pollable.Poll()
					poller.updateHistory(health, detail)
				}
			case operation := <-poller.control:
				log.Printf("control operation: %s", operation)
				switch operation {
				case run:
					poller.Running = true
				case pause:
					poller.Running = false
				case terminate:
					return
				case logging:
					poller.Logging = !poller.Logging
					fmt.Println("set logging:", poller.Logging)
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
	if poller.Logging == enable {
		return
	}
	poller.control <- logging
}

func (poller *Poller) updateHistory(health Health, detail string) {
	current := poller.History[len(poller.History)-1]

	if current.Health == health { //health status unchanged
		current.incrementPollCount()
	} else { // health status changed
		newState := NewState(poller.ID, health, detail, 1)
		poller.History = append(poller.History, newState)
	}
}

// String returns a string representation of the Poller suitable for printing
func (poller *Poller) String() string {
	return fmt.Sprintf("Poller[%d]: %s", poller.ID, poller.Pollable.String())
}

func (poller *Poller) Json() []byte {
	jout, err := json.Marshal(poller)
	if err != nil {
		return []byte(fmt.Sprintf(`{ "error" : "failed to marshall poller %d" }`, poller.ID))
	}
	return jout
}
