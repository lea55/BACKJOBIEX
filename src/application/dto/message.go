package dto

type RqSaveMessage struct {
	UserFromID string `json:"user_from" validate:"required"`
	UserToID   string `json:"user_to" validate:"required"`
	Content    string `json:"content" validate:"required"`
}
