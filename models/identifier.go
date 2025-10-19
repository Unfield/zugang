package models

import "time"

type IdentifierModel struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	TenantID  string     `json:"tenant_id"`
	Type      string     `json:"type"`
	Value     string     `json:"value"`
	Verified  bool       `json:"verified"`
	Primary   bool       `json:"primary"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `json:"deleted_at"`
}
