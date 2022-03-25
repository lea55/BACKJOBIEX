package repository

import "BACKJOBIEX/src/domain/entity"

type DevSkill interface {
	Save(doc entity.DevSkill) (string, error)
	GetAll() ([]entity.DevSkill, error)
	FindByID(ID string) (entity.DevSkill, error)
}
