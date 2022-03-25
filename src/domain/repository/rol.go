package repository

import "BACKJOBIEX/src/domain/entity"

type UserRol interface {
	SaveAll(list []entity.Rol) error
	FindByCode(code string) (entity.Rol, error)
}
