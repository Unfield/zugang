package models

import (
	"encoding/json"
	"time"
)

type TenantModel struct {
	ID        string          `json:"id"`
	TenantID  string          `json:"tenant_id"`
	Name      string          `json:"name"`
	Issuer    string          `json:"issuer"`
	Branding  any             `json:"branding"`
	Settings  json.RawMessage `json:"settings"`
	Active    bool            `json:"active"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeleteAt  *time.Time      `json:"deleted_at"`
}
