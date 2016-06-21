// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file

package check

import (
	"fmt"
	"github.com/rmorriso/gomon/monitor"
	"net"
)

// TcpCheck represents a hostname or IP address to be checked
// for open/close status of a given port
//
// TcpCheck implements the pollable interface
type TcpCheck struct {
	Addr         string
	Open         bool
	Port         int16
	PollInterval int64
}

// NewTcpCheck returns a TcpCheck type with the address (IP or hostname)
// and the given port number
func NewTcpCheck(addr string, port int16) *TcpCheck {
	return &TcpCheck{Addr: addr, Port: port}
}

// Poll attempts to open a TCP connection to the TcpCheck on its addr and Port
// and returns OK or an error string.
func (tcpCheck *TcpCheck) Poll() (monitor.Health, string, error) {
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", tcpCheck.Addr, tcpCheck.Port))
	if err != nil && tcpCheck.IsOpen() {
		return monitor.Critical, err.Error(), nil
	}
	return monitor.Ok, "listening", nil
}

// ID returns the ID of this TcpCheck pollable
func (tcpCheck *TcpCheck) ID() string {
	return tcpCheck.Addr
}

// Interval returns the polling interval (in sec) of this TcpCheck pollable
func (tcpCheck *TcpCheck) Interval() int64 {
	return tcpCheck.PollInterval
}

// IsOpen returns true if the desired state of this TcpCheck is open
func (tcpCheck *TcpCheck) IsOpen() bool {
	return tcpCheck.Open
}

// String returns a string representation of this TcpCheck pollable suitable for printing
func (tcpCheck *TcpCheck) String() string {
	return fmt.Sprintf("[TcpCheck: %v (Port = %d, Open = %v)]", tcpCheck.ID(), tcpCheck.Port, tcpCheck.IsOpen())
}
