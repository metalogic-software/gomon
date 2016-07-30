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
	"os/signal"
	"regexp"
	"syscall"
	"text/template"

	"github.com/rmorriso/gomon/monitor"
)

const (
	defaultConf = "./monitor.conf"

	second         = 1e9         // one second is 1e9 nanoseconds
	statusInterval = 10 * second // how often to log status to stdout
)

var (
	confFile   string
	checks   = new(Checks)
	listenAddr = flag.String("port", ":8080", "http port")
	gomon      = monitor.NewMonitor()
	templates  = make(map[string]*template.Template)
)

func init() {
	flag.StringVar(&confFile, "c", defaultConf, "the monitor checks config file")
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

var idValidator = regexp.MustCompile("^[1-9,a-f][0-9,a-f]*$")

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	load(checks)

	go func() {
		for _ = range c {
			log.Printf("reloading config %s\n", confFile)
			checks := new(Checks)
			checks.Init(confFile)
			load(checks)
		}
	}()

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
		err := templates["list"].Execute(w, gomon)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if !idValidator.MatchString(id) {
		http.NotFound(w, r)
		return
	}

	var err error
	switch r.Method {
	case "DELETE":
		err := gomon.Remove(id)
		if err == nil {
			fmt.Fprintf(w, "Deleted poller %s\n", id)
			return
		}
		fmt.Fprintf(w, err.Error())
	case "GET":
		poller, err := gomon.Poller(id)
		if err == nil {
			err = templates["view"].Execute(w, poller)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	fmt.Fprintf(w, err.Error())
	return
}

func envDefault(variable string, defaultValue string) string {
	result := os.Getenv(variable)
	if result == "" {
		return defaultValue
	}
	return result
}
