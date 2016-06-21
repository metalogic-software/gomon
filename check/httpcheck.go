// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package check

import (
	"github.com/rmorriso/gomon/monitor"
	"net/http"
)

// HTTPCheck represents an HTTP URL to be polled and the
// interval between polls in seconds
// TODO: implement backoff

// HTTPCheck implements the pollable interface
type HTTPCheck struct {
	URL          string
	PollInterval int64
}

// NewHTTPCheck constructs and initializes an HTTPCheck pollable
func NewHTTPCheck(url string) *HTTPCheck {
	return &HTTPCheck{URL: url}
}

// Poll executes an HTTP HEAD request for the resource url
// and returns health status and a detail string;
// TODO: if HTTP status is 403 or 404 returns an os.Error (NewError)
func (httpCheck *HTTPCheck) Poll() (monitor.Health, string, error) {
	resp, err := http.Head(httpCheck.URL)
	if err != nil {
		return monitor.Critical, err.Error(), nil
	}
	return monitor.Ok, resp.Status, nil
}

// ID returns the ID string associated with the HTTPCheck
func (httpCheck *HTTPCheck) ID() string {
	return httpCheck.URL
}

// Interval returns the polling interval
func (httpCheck *HTTPCheck) Interval() int64 {
	return httpCheck.PollInterval
}

// String returns a printable description of the HTTPCheck
func (httpCheck *HTTPCheck) String() string {
	return "[HTTPCheck:" + httpCheck.ID() + "]"
}
