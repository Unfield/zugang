package handlers

import (
	"github.com/Unfield/zugang/config"
	"github.com/Unfield/zugang/db"
)

type Handler struct {
	storageProvider db.StorageProvider
	config          *config.ZugangConfig
}

func NewHandler(storageProvider db.StorageProvider, config *config.ZugangConfig) (*Handler, error) {
	return &Handler{
		storageProvider: storageProvider,
		config:          config,
	}, nil
}
