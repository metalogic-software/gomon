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

func (this *Monitor) Add(pollable Pollable) {
	this.mu.Lock()
	defer this.mu.Unlock()

	this.index++
	log.Printf("adding Poller[%d]: %v\n", this.index, pollable)
	poller := NewPoller(this.index, pollable)
	this.pollers[this.index] = poller
	poller.Exec()
}

func (this *Monitor) Remove(id int) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	if poller, present := this.pollers[id]; present == true {
		fmt.Println("removing", id)
		poller.Terminate()
		delete(this.pollers, id)
		return nil
	}
	return errors.New("attempt to remove non-existent poller id")
}

func (this *Monitor) Run(id int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if poller, present := this.pollers[id]; present == true {
		fmt.Println("running", id)
		poller.Run()
	}
}

func (this *Monitor) Pause(id int) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if poller, present := this.pollers[id]; present == true {
		fmt.Println("pausing", id)
		poller.Pause()
	}
}

func (this *Monitor) SetLogging(logging bool) {
	for _, poller := range this.pollers {
		poller.Log(logging)
	}
}

func (this *Monitor) Pollers() map[int]*Poller {
	return this.pollers
}

func (this *Monitor) ListAll(w io.Writer) {
	for _, poller := range this.pollers {
		fmt.Fprintf(w, "http://localhost:8080/%d<br/>\n", poller.Id())
	}
}

func (this *Monitor) PrintDetail(w io.Writer, id int) {
	poller, present := this.pollers[id]
	if present {
		fmt.Fprintf(w, "%s<br/>", poller.String())
		for _, state := range poller.History() {
			fmt.Fprintf(w, "%s<br/>", state.String())
		}
	} else {
		fmt.Fprint(w, "PrintDetail not a valid poller id:"+strconv.Itoa(id))
	}
}

func (this *Monitor) PollerDetails(id int) (string, error) {
	poller, present := this.pollers[id]
	if present {
		details := fmt.Sprintf("%s<br/>", poller.String())
		for _, state := range poller.History() {
			details = details + fmt.Sprintf("%s<br/>", state.String())
		}
		return details, nil
	}
	return "", errors.New(fmt.Sprintf("Invalid poller id %d.", id))
}

func (this *Monitor) Poller(id int) (*Poller, error) {
	poller, present := this.pollers[id]
	if present {
		return poller, nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid poller id %d.", id))
}
