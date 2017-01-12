package main

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/peteclark-io/fixtures-rw/health"
	"github.com/peteclark-io/fixtures-rw/resources"
)

func main() {
	ping := health.Ping()
	runner := health.Aggregator(ping)

	r := mux.NewRouter()
	r.HandleFunc("/__ping", resources.Ping()).Methods("GET")
	r.HandleFunc("/__version", resources.Version()).Methods("GET")
	r.HandleFunc("/__health", resources.Health(runner)).Methods("GET")

	server := &http.Server{
		Handler: r,
		Addr:    ":80",

		WriteTimeout: 60 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Info("Starting server on localhost:80")
	server.ListenAndServe()
}
