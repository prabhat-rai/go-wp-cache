package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from WP"))
}


func (app *application) fetchWordpressData(w http.ResponseWriter, r *http.Request) {
	response := fetchWordpressData("pkrai.wordpress.com", "71978482")
	finalResponse := "{\"posts\" : " + response[0] + ", \"categories\" : " + response[1] + ", \"tags\" : " + response[2] +"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(finalResponse))
}
