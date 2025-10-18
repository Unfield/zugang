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
	router.HandleFunc("GET /v1/authorize", handler.NotImplementedHandler)

	router.HandleFunc("POST /v1/token", handler.TokenHandler)
	router.HandleFunc("POST /v1/introspect", handler.NotImplementedHandler)
	router.HandleFunc("POST /v1/revoke", handler.NotImplementedHandler)
	router.HandleFunc("GET /v1/userinfo", handler.NotImplementedHandler)
	router.HandleFunc("POST /v1/logout", handler.NotImplementedHandler)
	router.HandleFunc("GET /v1/{tenant}/.well‑known/openid‑configuration", handler.NotImplementedHandler)
	router.HandleFunc("GET /v1/{tenant}/.well‑known/jwks.json", handler.NotImplementedHandler)

	return router
}
