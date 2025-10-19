package models

import "time"

type TokenModel struct {
	ID        string     `json:"id"`
	TenantID  string     `json:"tenant_id"`
	UserID    string     `json:"user_id"`
	ClientID  string     `json:"client_id"`
	Type      string     `json:"type"`
	Scopes    string     `json:"scopes"`
	ExpiresAt int        `json:"expires_at"`
	Revoked   bool       `json:"revoked"`
	Signature string     `json:"signature"`
	IssuedAt  time.Time  `json:"issued_at"`
	Metadata  any        `json:"metadata"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `json:"deleted_at"`
}
