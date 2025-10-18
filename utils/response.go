package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, `{"error": "failed to encode response"}`, http.StatusInternalServerError)
	}
}

func OK(w http.ResponseWriter, data any) {
	WriteJSON(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, data any) {
	WriteJSON(w, http.StatusCreated, data)
}

func Error(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}

func BadRequest(w http.ResponseWriter, message string) {
	Error(w, http.StatusBadRequest, message)
}

func NotFound(w http.ResponseWriter, message string) {
	Error(w, http.StatusNotFound, message)
}

func InternalServerError(w http.ResponseWriter, message string) {
	Error(w, http.StatusInternalServerError, message)
}

type ErrorResponse struct {
	Err            string `json:"error"`
	ErrDescription string `json:"error_description"`
}

func SendError(w http.ResponseWriter, status int, err, desc string) {
	WriteJSON(w, http.StatusCreated, ErrorResponse{
		Err:            err,
		ErrDescription: desc,
	})
}
