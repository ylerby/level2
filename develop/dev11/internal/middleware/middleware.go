package middleware

import (
	"dev11/schemas"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func LoggingMiddleware(httpMethod string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != httpMethod {
			jsonResponse, err := json.Marshal(schemas.ErrorResponseSchema{
				Error: "Method not allowed",
			})
			if err != nil {
				os.Exit(1)
			}

			w.WriteHeader(http.StatusMethodNotAllowed)
			_, err = w.Write(jsonResponse)
			if err != nil {
				os.Exit(1)
			}
		} else {
			log.Printf("method: %s, request: %v", r.Method, r)
			next.ServeHTTP(w, r)
		}
	}
}
