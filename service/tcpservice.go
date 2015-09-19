// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package service

import (
	"fmt"
	"github.com/rmorriso/gomon/monitor"
	"net"
)

// TcpService represents a hostname or IP address to be checked
// for open/close status of a given port
//
// TcpService implements the pollable interface
type TcpService struct {
	Addr         string
	Open         bool
	Port         int16
	PollInterval int64
}

// NewTcpService returns a TcpService type with the address (IP or hostname)
// and the given port number
func NewTcpService(addr string, port int16) *TcpService {
	return &TcpService{Addr: addr, Port: port}
}

// Poll attempts to open a TCP connection to the TcpService on its addr and Port
// and returns OK or an error string.
func (svc *TcpService) Poll() (monitor.Health, string, error) {
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", svc.Addr, svc.Port))
	if err != nil && svc.IsOpen() {
		return monitor.Critical, err.Error(), nil
	}
	return monitor.Ok, "listening", nil
}

func (svc *TcpService) Id() string {
	return svc.Addr
}

func (svc *TcpService) Interval() int64 {
	return svc.PollInterval
}

func (svc *TcpService) IsOpen() bool {
	return svc.Open
}
func (svc *TcpService) String() string {
	return fmt.Sprintf("[TcpService: %v (Port = %d, Open = %v)]", svc.Id(), svc.Port, svc.IsOpen())
}
