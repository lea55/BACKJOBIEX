package entity

import "time"

type UserConversation struct {
	ID            string                 `json:"id"`
	RelatedUser   BasicUser              `json:"related_user"`
	UpdatedAt     time.Time              `json:"updated_at"`
	CreatedAt     time.Time              `json:"created_at"`
	Conversations []UserConversationItem `json:"conversations"`
}

type UserConversationItem struct {
	DestUser    BasicUser `json:"dest_user"`
	LastMessage string    `json:"last_message"`
	MessageDate time.Time `json:"message_date"`
	Read        bool      `json:"read"`
}
