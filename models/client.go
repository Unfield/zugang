package models

import "time"

type ClientModel struct {
	ID            string     `json:"id"`
	TenantID      string     `json:"tenant_id"`
	Name          string     `json:"name"`
	Secret        string     `json:"secret"`
	RedirectURIs  string     `json:"redirect_uris"`
	AllowedScopes string     `json:"allowed_scopes"`
	GrantTypes    string     `json:"grant_types"`
	Type          string     `json:"type"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeleteAt      *time.Time `json:"deleted_at"`
}
