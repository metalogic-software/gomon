// Copyright 2015 The Metalogic Software Authors. All rights reserved.
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/rmorriso/gomon/monitor"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
	"strconv"
	"text/template"
	//	"sync"
	//	"time"
)

const (
	defaultConf = "./monitor.conf"

	second         = 1e9         // one second is 1e9 nanoseconds
	statusInterval = 10 * second // how often to log status to stdout
)

var (
	confFile   string
	services     = new(Services)
	listenAddr = flag.String("port", ":8080", "http port")
	mon        = monitor.NewMonitor()
	templates  = make(map[string]*template.Template)
)

func init() {
	flag.StringVar(&confFile, "c", defaultConf, "the monitor services config file")
	log.SetPrefix("[gomon]: ")
	loadTemplates()
}

// associate templates with their respective pages
func loadTemplates() {
	for _, tmpl := range []string{"error", "list", "view"} {
		t := template.Must(template.ParseFiles("html/"+tmpl+".html", "html/header.html", "html/footer.html"))
		templates[tmpl] = t
	}
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
	services.Init(confFile)

	// submit httpservices for monitoring
	for _, httpservice := range services.HttpServices {
		mon.Add(httpservice)
	}

	// submit tcpservices for monitoring
	for _, tcpservice := range services.TcpServices {
		mon.Add(tcpservice)
	}

	//submit fileservices for monitoring
	for _, file := range services.Files {
		mon.Add(file)
	}

	// handle static content
	http.Handle("/inc/", http.FileServer(http.Dir("")))

	http.HandleFunc("/", root)
	err := http.ListenAndServe(*listenAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// root directory handler
func root(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	if id == "favicon.ico" || id == "" {
		err := templates["list"].Execute(w, mon)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if !idValidator.MatchString(id) {
		http.NotFound(w, r)
		return
	}

	if pollerId, err := strconv.Atoi(id); err == nil {
		switch r.Method {
		case "DELETE":
			err := mon.Remove(pollerId)
			if err == nil {
				fmt.Fprintf(w, "Deleted poller %s\n", id)
			} else {
				fmt.Fprintf(w, err.Error())
			}
		case "GET":
			if poller, err := mon.Poller(pollerId); err == nil {
				err = templates["view"].Execute(w, poller)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}
		return
	} else {
		fmt.Fprintf(w, err.Error())
	}
	return
}
