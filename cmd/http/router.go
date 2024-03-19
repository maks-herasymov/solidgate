package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) newRouter() http.Handler {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(app.notFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowed)

	r.Use(app.logAccess)
	r.Use(app.recoverPanic)

	r.HandleFunc("/healthcheck", app.healthcheck).Methods("GET")
	r.HandleFunc("/", app.validateCard).Methods("POST")

	return r
}
