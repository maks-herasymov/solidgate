package main

import (
	"fmt"
	"github.com/maks-herasymov/solidgate/internal/response"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"
)

func (app *application) reportServerError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)
	app.logger.Error(message, requestAttrs, "trace", trace)
}

func (app *application) errorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, map[string]string{"error": message}, headers)
	if err != nil {
		app.reportServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.reportServerError(r, err)
	app.errorMessage(w, r, http.StatusInternalServerError, "The server encountered a problem and could not process your request", nil)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	app.errorMessage(w, r, http.StatusNotFound, "The requested resource could not be found", nil)
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	app.errorMessage(w, r, http.StatusMethodNotAllowed, fmt.Sprintf("The %s method is not supported for this resource", r.Method), nil)

}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, e error) {
	app.errorMessage(w, r, http.StatusBadRequest, e.Error(), nil)
}
