package main

import (
	"encoding/json"
	"fmt"
	card "github.com/maks-herasymov/solidgate/pkg/card"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "api at your service! What can I do for you today? (Up)\n")
	})

	type ValidationResponse struct {
		Valid bool `json:"valid"`
	}

	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		details := &card.CardDetails{}

		err := json.NewDecoder(r.Body).Decode(details)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		isValid := card.IsValidCard(details)

		response := &ValidationResponse{isValid}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", mux)
}
