package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func InputValidationMiddleware(requiredFields []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Read the request body
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, `{"error": "Unable to read request body"}`, http.StatusBadRequest)
				return
			}

			// Parse the JSON payload
			var payload map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &payload); err != nil {
				http.Error(w, `{"error": "Invalid JSON payload"}`, http.StatusBadRequest)
				return
			}

			// Validate required fields
			for _, field := range requiredFields {
				if _, ok := payload[field]; !ok {
					http.Error(w, `{"error": "Missing required field: `+field+`"}`, http.StatusBadRequest)
					return
				}
			}

			// Reset the request body for the next handler
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
