package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, r *http.Request, statusCode int, payload interface{}) {

	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)

}

func ResponseWithError(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	ResponseWithJson(w, r, statusCode, map[string]string{"error": message})

}
