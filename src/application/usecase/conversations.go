package usecase

import (
	"BACKJOBIEX/src/domain/entity"
	"BACKJOBIEX/src/domain/repository"

	"github.com/pkg/errors"
)

type UserConversations struct {
	repo repository.UserConversation
}

func NewUserConversations(repo repository.UserConversation) *UserConversations {
	return &UserConversations{repo: repo}
}

func (c UserConversations) FindByUserID(userID string) (entity.UserConversation, error) {
	founded, result, err := c.repo.FindByRelatedUser(userID)
	if err != nil {
		return result, errors.Wrap(err, "Error en consulta a la base de datos")
	}

	if err == nil && !founded {
		return result, nil
	}

	return result, nil
}
