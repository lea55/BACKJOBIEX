package entity

import "time"

type Message struct {
	ID        string    `json:"id"`
	UserFrom  BasicUser `json:"user_from"`
	UserTo    BasicUser `json:"user_to"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
