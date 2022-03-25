package mappers

import (
	"time"

	"github.com/lea55/BACKJOBIEX/src/domain/entity"

	"github.com/google/uuid"
)

type Message struct {
	userMapper *User
}

func NewMessage() *Message {
	return &Message{}
}

func (m Message) FromCreateRequest(userFrom entity.User, userTo entity.User, content string) entity.Message {
	return entity.Message{
		ID: uuid.New().String(),
		UserFrom: entity.BasicUser{
			Names:        userFrom.Names,
			Surnames:     userFrom.Surnames,
			ID:           userFrom.ID,
			Email:        userFrom.Email,
			RolCode:      userFrom.Rol.Code,
			ProfileImage: userFrom.ProfilePhoto,
		},
		UserTo: entity.BasicUser{
			Names:        userTo.Names,
			Surnames:     userTo.Surnames,
			ID:           userTo.ID,
			Email:        userTo.Email,
			RolCode:      userTo.Rol.Code,
			ProfileImage: userTo.ProfilePhoto,
		},
		Content:   content,
		CreatedAt: time.Now(),
	}
}
