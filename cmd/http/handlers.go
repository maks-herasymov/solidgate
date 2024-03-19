package main

import (
	"encoding/json"
	"fmt"
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
	details := &card.CardDetails{}

	err := json.NewDecoder(r.Body).Decode(details)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	isValid := card.IsValidCard(details)

	type ValidationResponse struct {
		Valid bool `json:"valid"`
	}

	vr := &ValidationResponse{isValid}
	err = response.JSON(w, http.StatusOK, vr)
	if err != nil {
		app.serverError(w, r, err)
	}
}
