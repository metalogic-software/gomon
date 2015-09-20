// Copyright 2015 The Metalogic Software Authors. All rights reserved.
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
				case Run:
					poller.running = true
				case Pause:
					poller.running = false
				case Terminate:
					return
				case Logging:
					poller.logging = !poller.logging
					fmt.Println("set logging:", poller.logging)
				}
			}
		}
	}()
}

func (poller *Poller) Run() {
	poller.control <- Run
}

func (poller *Poller) Pause() {
	poller.control <- Pause
}
func (poller *Poller) Terminate() {
	poller.control <- Terminate
}

func (poller *Poller) Log(logging bool) {
	if poller.logging == logging {
		return
	}
	poller.control <- Logging
}

func (poller *Poller) Id() int {
	return poller.id
}

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

func (poller *Poller) Pollable() Pollable {
	return poller.pollable
}

func (poller *Poller) String() string {
	return fmt.Sprintf("Poller[%d]: %s", poller.id, poller.pollable.String())
}
