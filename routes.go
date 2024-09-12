package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.healthCheckHandler)
	mux.HandleFunc("POST /metadata", app.metadataHandler)

	middlewares := alice.New(app.logRequests)

	return middlewares.Then(mux)
}
