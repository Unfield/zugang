package handlers

import (
	"net/http"

	"github.com/Unfield/zugang/utils"
)

type authorizationCodeFlowRequest struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RedirectUri  string `json:"redirect_uri"`
	ClientID     string `json:"client_id"`
	CodeVerifier string `json:"code_verifier"`
}

func authorizationCodeFlowHandler(w http.ResponseWriter, r *http.Request) {
	var req authorizationCodeFlowRequest
	if err := utils.BindJSON(w, r, &req); err != nil {
		utils.BadRequest(w, err.Error())
		return
	}
}
