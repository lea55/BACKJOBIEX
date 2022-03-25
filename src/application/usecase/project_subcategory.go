package usecase

import (
	"github.com/lea55/BACKJOBIEX/src/application/dto"
	"github.com/lea55/BACKJOBIEX/src/domain/entity"
	"github.com/lea55/BACKJOBIEX/src/domain/repository"

	"github.com/pkg/errors"
)

type ProjectSubCategory struct {
	repo         repository.ProjectSubCategory
	categoryRepo repository.ProjectCategoryRepo
}

func NewProjectSubCategory(repo repository.ProjectSubCategory,
	categoryRepo repository.ProjectCategoryRepo) *ProjectSubCategory {
	return &ProjectSubCategory{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (c ProjectSubCategory) GetByCategoryID(pjCatID string) ([]entity.ProjectSubCategory, error) {
	list, err := c.repo.FindByCategoryID(pjCatID)
	if err != nil {
		return list, errors.Wrap(err, "Error en consulta a la base de datos")
	}

	return list, nil
}

func (c ProjectSubCategory) Save(model dto.RqRegisterProjectSubCat) error {
	cat, err := c.categoryRepo.FindByID(model.CategoryID)
	if err != nil {
		return errors.Wrap(err, "Error en consulta de categoría")
	}

	items := make([]entity.SubCategoryItem, 0)

	for _, v := range model.Items {
		newItem := entity.SubCategoryItem{
			Title:   v.Title,
			Options: v.Options,
		}
		items = append(items, newItem)
	}

	newSubCat := entity.ProjectSubCategory{
		Name:        model.Name,
		Description: model.Description,
		Category:    cat,
		Items:       items,
	}

	err = c.repo.Save(newSubCat)
	if err != nil {
		return errors.Wrap(err, "Error al guardar en la base de datos")
	}

	return nil
}
