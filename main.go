package main

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/peteclark-io/match-rw/resources"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/__ping", resources.Ping()).Methods("GET")
	r.HandleFunc("/__version", resources.Version()).Methods("GET")

	server := &http.Server{
		Handler: r,
		Addr:    ":80",

		WriteTimeout: 60 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Info("Starting server on localhost:80")
	server.ListenAndServe()
}
