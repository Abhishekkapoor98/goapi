package api

import (
	"encoding/json"
	"net/http"
)

// coin Balance Parameters

type CoinBalanceParams struct {
	Username string
}

// coin balance Response
type CoinBalanceResponse struct {
	// Success Code, Usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error Response
type ErrorResponse struct {
	// Error Code
	Code int

	// Error Message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := ErrorResponse{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occuoed", http.StatusInternalServerError)
	}
)
