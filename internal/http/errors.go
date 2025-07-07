package http

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type httpError interface {
	GetCode() string
	GetMessage() string
	GetField() string
}

func respondError[T httpError](w http.ResponseWriter, err []T) {
	w.Header().Set("Content-Type", "application/json")

	var response []ErrorResponse

	for _, httpError := range err {
		response = append(response, ErrorResponse{
			Code:    httpError.GetCode(),
			Message: httpError.GetMessage(),
			Field:   httpError.GetField(),
		})
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
