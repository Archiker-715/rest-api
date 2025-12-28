package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	errResp := ErrorResponse{
		Code:    code,
		Message: message,
	}
	json.NewEncoder(w).Encode(errResp)
}
