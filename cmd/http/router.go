package main

import (
	"fmt"
	_ "github.com/maks-herasymov/solidgate/api"
	httpSwagger "github.com/swaggo/http-swagger"

	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) newRouter() http.Handler {
	r := mux.NewRouter()

	r.NotFoundHandler = app.logAccess(http.HandlerFunc(app.notFound))
	r.MethodNotAllowedHandler = app.logAccess(http.HandlerFunc(app.methodNotAllowed))

	r.Use(app.logAccess)
	r.Use(app.recoverPanic)

	r.PathPrefix("/api").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/api/doc.json", app.config.httpPort)),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	r.HandleFunc("/healthcheck", app.healthcheck).Methods(http.MethodGet)
	r.HandleFunc("/", app.validateCard).Methods(http.MethodPost)

	return r
}
