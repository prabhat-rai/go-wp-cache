package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from WP"))
}


func (app *application) fetchWordpressData(w http.ResponseWriter, r *http.Request) {
	site := r.URL.Query().Get("site")
	if site == "" {
		site = "pkrai.wordpress.com"
	}
	response := fetchWordpressData(site, "71978482")
	finalResponse := "{\"posts\" : " + response[0] + ", \"categories\" : " + response[1] + ", \"tags\" : " + response[2] +"}"

	//putSiteTerms("test", response)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(finalResponse))
}
