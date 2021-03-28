package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from WP"))
}


func (app *application) fetchWordpressData(w http.ResponseWriter, r *http.Request) {
	fetchWordpressData("pkrai.wordpress.com", "71978482")
}
