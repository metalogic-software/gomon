// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package service

import "net/http"
import "github.com/metalogic-software/gomon/monitor"

// HTTPService represents an HTTP URL to be polled and the
// interval between polls in seconds
// TODO: implement backoff

// HTTPService implements the pollable interface
type HTTPService struct {
	URL          string
	PollInterval int64
}

// NewHTTPService constructs and initializes an HTTPService pollable
func NewHTTPService(url string) *HTTPService {
	return &HTTPService{URL: url}
}

// Poll executes an HTTP HEAD request for the resource url
// and returns health status and a detail string;
// TODO: if HTTP status is 403 or 404 returns an os.Error (NewError)
func (svc *HTTPService) Poll() (monitor.Health, string, error) {
	resp, err := http.Head(svc.URL)
	if err != nil {
		return monitor.Critical, err.Error(), nil
	}
	return monitor.Ok, resp.Status, nil
}

// ID returns the ID string associated with the HTTPService
func (svc *HTTPService) ID() string {
	return svc.URL
}

// Interval returns the polling interval
func (svc *HTTPService) Interval() int64 {
	return svc.PollInterval
}

// String returns a printable description of the HTTPService
func (svc *HTTPService) String() string {
	return "[HTTPService:" + svc.ID() + "]"
}
