package drivers

import (
	"database/sql"

	"github.com/Unfield/zugang/config"
)

type DefaultStorageProvider struct {
	cfg *config.ZugangConfig
	db  *sql.DB
	kv  any
}

func NewDefaultStorageProvider(cfg *config.ZugangConfig) (DefaultStorageProvider, error) {
	return DefaultStorageProvider{
		cfg: cfg,
		db:  nil,
		kv:  nil,
	}, nil
}
