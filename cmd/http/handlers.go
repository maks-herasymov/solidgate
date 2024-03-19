package main

import (
	"fmt"
	"github.com/maks-herasymov/solidgate/internal/request"
	"github.com/maks-herasymov/solidgate/internal/response"
	"github.com/maks-herasymov/solidgate/pkg/card"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "api at your service! What can I do for you today? (Up)\n")
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) validateCard(w http.ResponseWriter, r *http.Request) {
	var details card.Details

	err := request.DecodeJSON(w, r, &details)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	isValid := card.IsValidCard(&details)

	vr := &struct {
		Valid bool `json:"valid"`
	}{isValid}

	err = response.JSON(w, http.StatusOK, vr)
	if err != nil {
		app.serverError(w, r, err)
	}
}
