// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package monitor

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"
	//	"time"
)

type Monitor struct {
	index   int
	mu      sync.Mutex
	pollers map[int]*Poller
}

// Monitor constructor
func NewMonitor() *Monitor {
	return &Monitor{index: 0, pollers: make(map[int]*Poller)}
}

func (mon *Monitor) Add(pollable Pollable) {
	mon.mu.Lock()
	defer mon.mu.Unlock()

	mon.index++
	log.Printf("adding Poller[%d]: %v\n", mon.index, pollable)
	poller := NewPoller(mon.index, pollable)
	mon.pollers[mon.index] = poller
	poller.Exec()
}

func (mon *Monitor) Remove(id int) error {
	mon.mu.Lock()
	defer mon.mu.Unlock()

	if poller, present := mon.pollers[id]; present == true {
		fmt.Println("removing", id)
		poller.Terminate()
		delete(mon.pollers, id)
		return nil
	}
	return errors.New("attempt to remove non-existent poller id")
}

func (mon *Monitor) Run(id int) {
	mon.mu.Lock()
	defer mon.mu.Unlock()

	if poller, present := mon.pollers[id]; present == true {
		fmt.Println("running", id)
		poller.Run()
	}
}

func (mon *Monitor) Pause(id int) {
	mon.mu.Lock()
	defer mon.mu.Unlock()

	if poller, present := mon.pollers[id]; present == true {
		fmt.Println("pausing", id)
		poller.Pause()
	}
}

func (mon *Monitor) SetLogging(logging bool) {
	for _, poller := range mon.pollers {
		poller.Log(logging)
	}
}

func (mon *Monitor) Pollers() map[int]*Poller {
	return mon.pollers
}

func (mon *Monitor) ListAll(w io.Writer) {
	for _, poller := range mon.pollers {
		fmt.Fprintf(w, "http://localhost:8080/%d<br/>\n", poller.Id())
	}
}

func (mon *Monitor) PrintDetail(w io.Writer, id int) {
	poller, present := mon.pollers[id]
	if present {
		fmt.Fprintf(w, "%s<br/>", poller.String())
		for _, state := range poller.History() {
			fmt.Fprintf(w, "%s<br/>", state.String())
		}
	} else {
		fmt.Fprint(w, "PrintDetail not a valid poller id:"+strconv.Itoa(id))
	}
}

func (mon *Monitor) PollerDetails(id int) (string, error) {
	poller, present := mon.pollers[id]
	if present {
		details := fmt.Sprintf("%s<br/>", poller.String())
		for _, state := range poller.History() {
			details = details + fmt.Sprintf("%s<br/>", state.String())
		}
		return details, nil
	}
	return "", errors.New(fmt.Sprintf("Invalid poller id %d.", id))
}

func (mon *Monitor) Poller(id int) (*Poller, error) {
	poller, present := mon.pollers[id]
	if present {
		return poller, nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid poller id %d.", id))
}
