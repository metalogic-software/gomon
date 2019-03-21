// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"testing"
	"time"

	"github.com/metalogic-software/gomon/monitor"
)

var pollers map[int]*monitor.Poller
var paused *monitor.Poller

func TestMain(m *testing.M) {
	services.Init("./monitor.conf")
	for _, httpservice := range services.HTTPServices {
		gomon.Add(httpservice)
	}
	pollers = gomon.Pollers()
	paused = pollers[3]
	paused.Pause()
	os.Exit(m.Run())
}

func TestAdd(t *testing.T) {
	id := 1
	poller := pollers[id]

	if id != poller.ID() {
		t.Fatalf("expected %d but got %d\n", id, poller.ID())
	}

	expected := "Poller[1]: [HTTPService:http://blog.golang.org]"
	desc := poller.String()
	if desc != expected {
		t.Fatalf("expected %s but got %s\n", expected, desc)
	}

	h := len(poller.History())
	if 1 != h {
		t.Fatalf("expected %d but got %d\n", 1, h)
	}
}

func TestRemove(t *testing.T) {
	id := 1
	nosuchid := -1

	err := gomon.Remove(id)
	if err != nil {
		t.Fatalf("unexpected error removing poller id %d\n", id)
	}

	if pollers[id] != nil {
		t.Fatalf("pollers[%d] is not empty after Remove()\n", id)
	}

	err = gomon.Remove(nosuchid)
	if err == nil {
		t.Fatalf("expected error removing non-existent poller id %d\n", nosuchid)
	}
}

func TestPollingInterval(t *testing.T) {
	poller := pollers[2]
	pollable := poller.Pollable()

	ticker := time.NewTicker(time.Duration((pollable.Interval() + 5) * second))
	for {
		select {
		case <-ticker.C:
			h := poller.History()
			if 2 != len(h) {
				t.Fatalf("expected poll history of length 2 after pollable interval has elapsed\n")
			}
			return
		}
	}
}

func TestPausePolling(t *testing.T) {
	pollable := paused.Pollable()

	ticker := time.NewTicker(time.Duration((pollable.Interval() + 5) * second))
	for {
		select {
		case <-ticker.C:
			h := paused.History()
			if 1 != len(h) {
				t.Fatalf("expected poll history of length 1 after pollable interval has elapsed on a paused poller but got %d\n", len(h))
			}
			return
		}
	}
}
