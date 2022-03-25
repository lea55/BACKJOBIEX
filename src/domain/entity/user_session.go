package entity

import "time"

type UserSession struct {
	TokenID      string    `json:"token_id"`
	ConnectionIP string    `json:"connection_ip"`
	SigInDate    time.Time `json:"sig_in_date"`
	Token        string    `json:"token"`
	UserID       string    `json:"user_id"`
	UserRolCode  string    `json:"user_rol_code"`
}
