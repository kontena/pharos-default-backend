package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/TV4/graceful"
	"github.com/gorilla/mux"
)

var default_page []byte
var health_page []byte

func main() {
	// read output pages into mem
	var err error
	default_page, err = ioutil.ReadFile("static/default.html")
	health_page, err = ioutil.ReadFile("static/healthz.html")
	if err != nil {
		log.Fatal(err)
		os.Exit(42)
	}
	// Setup routes
	router := mux.NewRouter()
	router.HandleFunc("/healthz", GetHealth)
	router.PathPrefix("/").HandlerFunc(GetRoot)
	// Handles gracefull shutdown nicely
	graceful.LogListenAndServe(&http.Server{
		Addr:    ":8080",
		Handler: router,
	})
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write(default_page)
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	w.Write(health_page)
}
