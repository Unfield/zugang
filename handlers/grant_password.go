package handlers

import (
	"net/http"

	"github.com/Unfield/zugang/utils"
)

type passwordFlowRequest struct {
	GrantType    string `json:"grant_type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

func passwordFlowHandler(w http.ResponseWriter, r *http.Request) {
	var req passwordFlowRequest
	if err := utils.BindJSON(w, r, &req); err != nil {
		utils.BadRequest(w, err.Error())
		return
	}
}
