package resources

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Ping returns a pong if the service is reachable.
func Ping() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pong, _ := json.Marshal(struct {
			Pong bool `json:"pong"`
		}{true})
		w.Write(pong)
	}
}

// Version returns info.
func Version() func(w http.ResponseWriter, r *http.Request) {
	version, _ := ioutil.ReadFile("/version.json")

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(version)
	}
}
