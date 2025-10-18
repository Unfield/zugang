package handlers

import (
	"net/http"

	"github.com/Unfield/zugang/utils"
)

type clientCredentialFlowRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

func clientCredentialsFlowHandler(w http.ResponseWriter, r *http.Request) {
	var req clientCredentialFlowRequest
	if err := utils.BindJSON(w, r, &req); err != nil {
		utils.BadRequest(w, err.Error())
		return
	}
}
