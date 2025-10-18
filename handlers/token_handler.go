package handlers

import (
	"net/http"

	"github.com/Unfield/zugang/utils"
)

type TokenRequest struct {
	GrantType string `json:"grant_type"`
}

func (h *Handler) TokenHandler(w http.ResponseWriter, r *http.Request) {
	var req TokenRequest
	if err := utils.BindJSON(w, r, &req); err != nil {
		utils.BadRequest(w, err.Error())
		return
	}

	if !h.config.OAuth.GrantTypesAll && !utils.ContainsString(h.config.OAuth.GrantTypes, req.GrantType) {
		utils.SendError(w, 400, "unsupported_grant_type", "grant type not supported")
		return
	}

	switch req.GrantType {
	case "authorization_code":
		authorizationCodeFlowHandler(w, r)
	case "password":
		passwordFlowHandler(w, r)
	case "client_credentials":
		clientCredentialsFlowHandler(w, r)
	case "refresh_token":
		refreshTokenFlowHandler(w, r)
	default:
		utils.SendError(w, 400, "unsupported_grant_type", "grant type not supported")
	}
}

type SuccessResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	IDToken      string `json:"id_token,omitempty"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}
