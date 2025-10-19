package models

import "time"

type UserModel struct {
	ID                  string     `json:"id"`
	TenantID            string     `json:"tenant_id"`
	PrimaryIdentifierID string     `json:"primary_identifier_id"`
	ProfileName         string     `json:"profile_name"`
	Status              string     `json:"status"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeleteAt            *time.Time `json:"deleted_at"`
}
