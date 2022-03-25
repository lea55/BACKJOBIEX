package repository

import "github.com/lea55/BACKJOBIEX/src/domain/entity"

type Message interface {
	Save(message entity.Message) error
	FindByUser(userID string, senderUserID string, from uint32, limit uint8) ([]entity.Message, error)
}
