package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func route(listenAddr string) {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/pollers", Pollers)
	router.GET("/pollers/:id", Poller)
	router.GET("/pollers/:id/history", PollerHistory)
	router.DELETE("/pollers/:id", PollerDelete)

	log.Fatal(http.ListenAndServe(listenAddr, router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Poller(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if !idValidator.MatchString(id) {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Poller %s\n", ps.ByName("id"))
}

func PollerHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Poller history of %s\n", ps.ByName("id"))
}

func Pollers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	for _, poller := range gomon.Pollers() {
		fmt.Fprintf(w, "%s\n", poller.Json())
	}
}

func PollerDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if !idValidator.MatchString(id) {
		http.NotFound(w, r)
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	err = gomon.Remove(i)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "Deleted poller %s\n", id)
}

var idValidator = regexp.MustCompile("^[1-9][0-9]*$")
