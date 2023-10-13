package helpers

import (
	"encoding/json"
	"net/http"
)

func Header(w http.ResponseWriter, methods string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", methods)
}

func ResponseSuccess(w http.ResponseWriter, message string, status string, code int, args map[string]interface{}) {
	w.WriteHeader(code)
	response := map[string]interface{}{
		"message":    message,
		"status":     status,
		"statusCode": code,
	}
	for key, value := range args {
		response[key] = value
	}
	_ = json.NewEncoder(w).Encode(response)
}

func ResponseErrorSender(w http.ResponseWriter, message string, status string, code int) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    message,
		"status":     status,
		"statusCode": code,
	})
}

const (
	FAILED  = "FAILED"
	SUCCESS = "SUCCESS"
)
