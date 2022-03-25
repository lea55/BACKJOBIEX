package repository

import "github.com/lea55/BACKJOBIEX/src/domain/entity"

type UserSession interface {
	Save(doc entity.UserSession) (string, error)
	FindByID(Id string) (entity.UserSession, error)
}
