package repository

import "BACKJOBIEX/src/domain/entity"

type ProjectCategoryRepo interface {
	Save(doc entity.ProjectCategory) error
	Find() ([]entity.ProjectCategory, error)
	UpdateImage(ID string, image string) error
	UpdateIcon(ID string, image string) error
	FindByID(ID string) (entity.ProjectCategory, error)
}
