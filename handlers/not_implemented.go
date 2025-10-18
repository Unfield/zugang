package handlers

import (
	"net/http"

	"github.com/Unfield/zugang/utils"
)

func (h *Handler) NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	utils.BadRequest(w, "Not yet implemented")
}
