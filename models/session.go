package models

import "time"

type SessionModel struct {
	ID             string     `json:"id"`
	TenantID       string     `json:"tenant_id"`
	UserID         string     `json:"user_id"`
	ClientID       string     `json:"client_id"`
	UserAgent      string     `json:"user_agent"`
	DeviceInfo     string     `json:"device_info"`
	DeviceLocation string     `json:"device_location"`
	IP             string     `json:"ip"`
	Revoked        bool       `json:"revoked"`
	Metadata       any        `json:"metadata"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeleteAt       *time.Time `json:"deleted_at"`
}
