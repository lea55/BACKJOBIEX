package repository

import "BACKJOBIEX/src/domain/entity"

type UserSession interface {
	Save(doc entity.UserSession) (string, error)
	FindByID(Id string) (entity.UserSession, error)
}
