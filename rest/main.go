// Copyright 2012-2016 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rmorriso/gomon/monitor"
)

const (
	defaultConf = "./monitor.conf"

	second         = 1e9         // one second is 1e9 nanoseconds
	statusInterval = 10 * second // how often to log status to stdout
)

var (
	confFile   string
	port       = flag.Int("port", 8080, "listen port")
	listenAddr string
	rundir     string
	checks     = new(Checks)
	gomon      = monitor.NewMonitor()
)

func init() {
	log.SetPrefix("[gomon]: ")
	rundir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	flag.StringVar(&confFile, "c", fmt.Sprintf("%s/%s", rundir, defaultConf), "the monitor checks config file")
}

func main() {
	flag.Parse()

	listenAddr = fmt.Sprintf(":%d", *port)

	checks.Init(confFile)

	addChecks()

	route(listenAddr)

}

// root directory handler
func root(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]

	if id == "" {
		for _, poller := range gomon.Pollers() {
			fmt.Fprintf(w, "%s\n", poller.Json())
		}
		return
	}

	if !idValidator.MatchString(id) {
		http.NotFound(w, r)
		return
	}

	var err error
	if pollerID, err := strconv.Atoi(id); err == nil {
		switch r.Method {
		case "DELETE":
			err := gomon.Remove(pollerID)
			if err == nil {
				fmt.Fprintf(w, "Deleted poller %s\n", id)
			} else {
				fmt.Fprintf(w, err.Error())
			}
		case "GET":
			if details, err := gomon.PollerDetails(pollerID); err == nil {
				fmt.Fprintf(w, "%s\n", details)
			}
		}
		return
	}
	fmt.Fprintf(w, err.Error())
	return
}

func addChecks() {
	// submit httpchecks for monitoring
	for _, httpcheck := range checks.HTTPChecks {
		gomon.Add(httpcheck)
	}

	// submit tcpchecks for monitoring
	for _, tcpcheck := range checks.TcpChecks {
		gomon.Add(tcpcheck)
	}

	//submit filechecks for monitoring
	for _, filecheck := range checks.FileChecks {
		gomon.Add(filecheck)
	}

	//submit shellchecks for monitoring
	for _, shellcheck := range checks.ShellChecks {
		gomon.Add(shellcheck)
	}
}
