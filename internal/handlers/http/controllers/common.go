package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	OperationID string `json:"operation_id"`
	Data        any    `json:"data"`
	Metadata    any    `json:"metadata,omitempty"`
}

type ErrorResponse struct {
	OperationID string `json:"operation_id"`
	Message     any    `json:"message"`
	Errors      any    `json:"errors,omitempty"`
}

func writeJSONResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println(err)
	}
}
