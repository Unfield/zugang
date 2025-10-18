package handlers

import (
	"net/http"

	"github.com/Unfield/zugang/utils"
)

type refreshTokenFlowRequest struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
	ClientID     string `json:"client_id"`
	Scope        string `json:"scope"`
}

func refreshTokenFlowHandler(w http.ResponseWriter, r *http.Request) {
	var req refreshTokenFlowRequest
	if err := utils.BindJSON(w, r, &req); err != nil {
		utils.BadRequest(w, err.Error())
		return
	}
}
