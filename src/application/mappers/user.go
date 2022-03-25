package mappers

import (
	"time"

	"BACKJOBIEX/src/application/dto"
	"BACKJOBIEX/src/domain/entity"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u User) GetToSimpleReg(doc dto.RqRegisterUser, pass string, rol entity.Rol) entity.User {
	return entity.User{
		Names:             doc.Names,
		Surnames:          doc.Surnames,
		Email:             doc.Email,
		Password:          pass,
		LastSessionIp:     "__",
		LastSessionDate:   time.Now().UTC(),
		Enabled:           true,
		CreatedAt:         time.Now().UTC(),
		Rol:               rol,
		NickName:          doc.UserName,
		CompletedProjects: 0,
	}
}

func (u User) GetBasicFromEntity(user entity.User) entity.BasicUser {
	return entity.BasicUser{
		Names:        user.Names,
		Surnames:     user.Surnames,
		ID:           user.ID,
		Email:        user.Email,
		RolCode:      user.Rol.Code,
		ProfileImage: user.Image,
	}
}
