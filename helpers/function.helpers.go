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

func ResponseSuccess(w http.ResponseWriter, message string, status string, code int, data interface{}, count int) {
	w.WriteHeader(code)
	response := map[string]interface{}{
		"message":    message,
		"status":     status,
		"statusCode": code,
	}
	if data != nil {
		response["data"] = data
	}
	if count != 0 {
		response["count"] = count
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
