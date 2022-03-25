package mappers

import (
	"time"

	"github.com/lea55/BACKJOBIEX/src/application/dto"
	"github.com/lea55/BACKJOBIEX/src/domain/entity"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

func (a Auth) BuildSessionToSave(IP string, token string, user entity.User) entity.UserSession {
	return entity.UserSession{
		ConnectionIP: IP,
		SigInDate:    time.Now().UTC(),
		Token:        token,
		UserID:       user.ID,
		UserRolCode:  user.Rol.Code,
	}
}
func (a Auth) BuildUser(user entity.User, token string, sessionID string) dto.AuthenticatedUser {
	return dto.AuthenticatedUser{
		ID:      user.ID,
		Name:    user.Names,
		Surname: user.Surnames,
		Token:   token,
		Email:   user.Email,
		TokenID: sessionID,
	}
}
