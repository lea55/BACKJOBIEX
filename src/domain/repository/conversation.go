package repository

import "github.com/lea55/BACKJOBIEX/src/domain/entity"

type UserConversation interface {
	Save(doc entity.UserConversation) error
	FindByRelatedUser(userID string) (bool, entity.UserConversation, error)
	UpdateItem(ID string, items []entity.UserConversationItem) error
}
