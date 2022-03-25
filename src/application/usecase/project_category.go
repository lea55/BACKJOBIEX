package usecase

import (
	"math/rand"

	"BACKJOBIEX/src/application/dto"
	"BACKJOBIEX/src/domain/entity"
	"BACKJOBIEX/src/domain/repository"

	"github.com/pkg/errors"
)

type ProjectCategory struct {
	repo repository.ProjectCategoryRepo
}

func NewProjectCategory(repo repository.ProjectCategoryRepo) *ProjectCategory {
	return &ProjectCategory{repo: repo}
}

func (c ProjectCategory) Save(model dto.RqCommonRegister) error {
	newCat := entity.ProjectCategory{
		Name:        model.Name,
		Description: model.Description,
	}

	err := c.repo.Save(newCat)
	if err != nil {
		return errors.Wrap(err, "Error al guardar categoría")
	}
	return nil
}

func (c ProjectCategory) GetList() ([]entity.ProjectCategory, error) {
	result, err := c.repo.Find()
	if err != nil {
		return nil, errors.Wrap(err, "Error en consulta a la base de datos")
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result, nil
}

func (c ProjectCategory) UpdateImage(ID string, image string) error {
	_, err := c.repo.FindByID(ID)
	if err != nil {
		return errors.New("la cateogoria no existe en la base de datos")
	}

	err = c.repo.UpdateImage(ID, image)
	if err != nil {
		return errors.Wrap(err, "Error al giardar imagen en la base de datso")
	}

	return nil
}

func (c ProjectCategory) UpdateIcon(ID string, image string) error {
	_, err := c.repo.FindByID(ID)
	if err != nil {
		return errors.New("la cateogoria no existe en la base de datos")
	}

	err = c.repo.UpdateIcon(ID, image)
	if err != nil {
		return errors.Wrap(err, "Error al guardar imagen en la base de datos")
	}

	return nil
}
