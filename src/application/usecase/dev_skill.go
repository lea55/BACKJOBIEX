package usecase

import (
	"BACKJOBIEX/src/domain/entity"
	"BACKJOBIEX/src/domain/repository"

	"github.com/pkg/errors"
)

type DevSkill struct {
	repo repository.DevSkill
}

func NewDevSkill(repo repository.DevSkill) *DevSkill {
	return &DevSkill{repo: repo}
}

func (ds DevSkill) GetAll() ([]entity.DevSkill, error) {
	return ds.repo.GetAll()
}

func (ds DevSkill) Save(name string, description string) (string, error) {

	newDevSkill := entity.DevSkill{
		Name:        name,
		Description: description,
	}

	idSaved, err := ds.repo.Save(newDevSkill)
	if err != nil {
		return "", errors.Wrap(err, "Error en la base de datos")
	}

	return idSaved, nil
}
