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
	"text/template"
)

const (
	second         = 1e9         // one second is 1e9 nanoseconds
	statusInterval = 10 * second // how often to log status to stdout
)

var (
	confFile   string
	listenAddr string
	api        = flag.String("api", "http://localhost:9080", "api base url")
	port       = flag.Int("port", 8080, "http port")
	templates  = make(map[string]*template.Template)
)

func init() {
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

func main() {
	flag.Parse()

	listenAddr = fmt.Sprintf(":%d", *port)

	// handle static content
	http.Handle("/inc/", http.FileServer(http.Dir("")))

	http.HandleFunc("/", root)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// root directory handler
func root(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "working")
	return
}
