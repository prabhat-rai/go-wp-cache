package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice" // New import
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/fetch-wp", http.HandlerFunc(app.fetchWordpressData))

	return standardMiddleware.Then(mux)
}
