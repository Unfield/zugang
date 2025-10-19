package models

import "time"

type CredentialsModel struct {
	ID           string     `json:"id"`
	UserID       string     `json:"user_id"`
	TenantID     string     `json:"tenant_id"`
	Type         string     `json:"type"`
	Provider     string     `json:"provider"`
	SecretHash   string     `json:"secret_hash"`
	IdentifierID string     `json:"identifier_id"`
	LastUsed     *time.Time `json:"last_used"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeleteAt     *time.Time `json:"deleted_at"`
}
