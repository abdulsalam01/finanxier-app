package http

import (
	"encoding/json"
	"net/http"

	"github.com/finanxier-app/internal/entity/base"
)

func GenericMiddleware(next func(http.ResponseWriter, *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err          error
			jsonResponse []byte
		)

		// Set the Content-Type header to indicate JSON response.
		w.Header().Set("Content-Type", "application/json")

		// Call the handler and capture its response and error.
		response, err := next(w, r)
		mapData := base.Response[interface{}]{
			Data:    response,
			Message: "Successfully executed",
			Success: err == nil,
		}

		if err != nil {
			// Catch the errors.
			mapData.Data = nil
			mapData.Message = err.Error()
			// Safe operation, without breaking changes or flow.
			jsonResponse, err = json.Marshal(mapData) //nolint:all

			// Write the JSON response.
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonResponse)
			return
		}

		// Marshal the response data to JSON.
		jsonResponse, err = json.Marshal(mapData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the JSON response.
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse) //nolint:all
	}
}
