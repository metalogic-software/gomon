// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

import (
	"log"
	"time"
)

const (
	logInterval = 10 * second
)

var (
	logEnabled = true
)

// Logger maintains a map that stores the most recent states reported by Pollers,
// and logs the states every logInterval seconds;
// It returns a chan State to which state updates should be sent.
func Logger(logInterval time.Duration) (updates chan State, control chan bool) {
	updates = make(chan State)
	control = make(chan bool)
	stateMap := make(map[int]State)
	ticker := time.NewTicker(logInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				if logEnabled {
					logState(stateMap)
				}
			case state := <-updates:
				stateMap[state.id] = state
			case logEnabled = <-control:
				log.Printf("set logging: %s", logging)
			}
		}
	}()
	return updates, control
}

// logState prints a state map.
func logState(states map[int]State) {
	log.Println("Current state:")
	for id, state := range states {
		log.Printf("Poller[%d]: %s", id, state)
	}
}
