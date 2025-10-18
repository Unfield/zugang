package api

import (
	"net/http"

	"github.com/Unfield/zugang/config"
	"github.com/Unfield/zugang/handlers"
)

func NewRouter(config *config.ZugangConfig) *http.ServeMux {
	router := http.NewServeMux()

	handler, err := handlers.NewHandler(nil, config)
	if err != nil {
		panic(err)
	}

	router.HandleFunc("GET /v1/health", handler.HealthHandler)

	// CORE OAuth & OIDC
	router.HandleFunc("GET /v1/authorize", nil)

	router.HandleFunc("POST /v1/token", handler.TokenHandler)
	router.HandleFunc("POST /v1/introspect", nil)
	router.HandleFunc("POST /v1/revoke", nil)
	router.HandleFunc("GET /v1/userinfo", nil)
	router.HandleFunc("POST /v1/logout", nil)
	router.HandleFunc("GET /v1/{tenant}/.well‑known/openid‑configuration", nil)
	router.HandleFunc("GET /v1/{tenant}/.well‑known/jwks.json", nil)

	return router
}
