// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
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
	checks     = new(Checks)
	listenAddr = flag.String("port", ":8080", "http port")
	gomon      = monitor.NewMonitor()
)

func init() {
	flag.StringVar(&confFile, "c", defaultConf, "the monitor services config file")
	log.SetPrefix("[gomon]: ")
}

var idValidator = regexp.MustCompile("^[1-9][0-9]*$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[1:]
		if !idValidator.MatchString(id) {
			http.NotFound(w, r)
			return
		}
		fn(w, r, id)
	}
}

func main() {
	flag.Parse()
	checks.Init(confFile)

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

	http.HandleFunc("/", root)
	err := http.ListenAndServe(*listenAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// root directory handler
func root(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]

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
			if poller, err := gomon.Poller(pollerID); err == nil {
				fmt.Fprintf(w, "%s\n", poller)
			}
		}
		return
	}
	fmt.Fprintf(w, err.Error())
	return
}
