package handler

import (
	"encoding/json"
	"net/http"

)

type StringProcessor func([]string) []string

func GenericHandler(processor StringProcessor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input []string
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		output := processor(input)

		response, err := json.Marshal(output)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
