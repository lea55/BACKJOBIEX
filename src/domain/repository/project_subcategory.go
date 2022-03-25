package repository

import "BACKJOBIEX/src/domain/entity"

type ProjectSubCategory interface {
	Save(doc entity.ProjectSubCategory) error
	Find() ([]entity.ProjectSubCategory, error)
	UpdateImage(ID string, image string) error
	FindByID(ID string) (entity.ProjectSubCategory, error)
	FindByCategoryID(catID string) ([]entity.ProjectSubCategory, error)
}
