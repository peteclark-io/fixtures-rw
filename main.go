package main

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/__ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	server := &http.Server{
		Handler: r,
		Addr:    ":80",

		WriteTimeout: 60 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Info("Starting server on localhost:80")
	server.ListenAndServe()
}
