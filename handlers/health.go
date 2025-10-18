package handlers

import (
	"net/http"
	"time"

	"github.com/Unfield/zugang/utils"
)

type HealthResponse struct {
	Status    string `json:"status" xml:"status"`
	Timestamp string `json:"timestamp" xml:"timestamp"`
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	utils.OK(w, HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}
