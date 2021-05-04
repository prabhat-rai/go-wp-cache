package main

import (
	"net/http"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from WP"))
}

func (app *application) fetchWordpressData(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	site := r.URL.Query().Get("site")

	if site == "" {
		site = "pkrai.wordpress.com"
	}

	response := fetchWordPressData(site, "71978482")
	timeTaken := time.Since(startTime).String()

	app.infoLog.Printf("Time Taken to Process the Complete API %v", timeTaken)
	finalResponse := "{" +
		"\"time_taken\" : \"" + timeTaken + "\"," +
		"\"posts\" : " + response["posts"] + ", " +
		"\"categories\" : " + response["categories"] + ", " +
		"\"tags\" : " + response["tags"] +
		"}"

	//putSiteTerms("test", response)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(finalResponse))
}
