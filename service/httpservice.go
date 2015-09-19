// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package service

import "net/http"
import "github.com/rmorriso/gomon/monitor"

// HttpService represents an HTTP URL to be polled and the
// interval between polls in seconds
// TODO: implement backoff
type HttpService struct {
	Url          string
	PollInterval int64
}

func NewHttpService(url string) *HttpService {
	return &HttpService{Url: url}
}

// Poll executes an HTTP HEAD request for the resource url
// and returns health status and a detail string;
// TODO: if HTTP status is 403 or 404 returns an os.Error (NewError)
func (this *HttpService) Poll() (monitor.Health, string, error) {
	resp, err := http.Head(this.Url)
	if err != nil {
		return monitor.Critical, err.Error(), nil
	}
	return monitor.Ok, resp.Status, nil
}

// Id returns the ID string associated with the HttpService
func (this *HttpService) Id() string {
	return this.Url
}

// Interval returns the polling interval
func (this *HttpService) Interval() int64 {
	return this.PollInterval
}

// String returns a printable description of the HttpService
func (this *HttpService) String() string {
	return "[HttpService:" + this.Id() + "]"
}