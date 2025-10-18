package handlers

import (
	"database/sql"

	"github.com/Unfield/zugang/config"
)

type Handler struct {
	db     *sql.DB
	config *config.ZugangConfig
}

func NewHandler(db *sql.DB, config *config.ZugangConfig) (*Handler, error) {
	return &Handler{
		db:     db,
		config: config,
	}, nil
}
