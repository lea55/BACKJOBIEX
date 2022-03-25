package usecase

import (
	"github.com/lea55/BACKJOBIEX/src/application/helpers"
	"github.com/lea55/BACKJOBIEX/src/domain/entity"
	"github.com/lea55/BACKJOBIEX/src/domain/repository"
	"github.com/pkg/errors"
)

type Rol struct {
	repo   repository.UserRol
	helper *helpers.Rol
}

//NewRol Constructor de caso de uso de roles de usuario
func NewRol(repo repository.UserRol) *Rol {
	return &Rol{
		repo:   repo,
		helper: helpers.NewRol(),
	}
}

//FindByCode Busca en la base de datos un registro que coincida con el codigo recibido, de no existir lee de un json
//en los archivos locales el cual contiene la lista de roles segun la logica de negocio, retorna un valor que coincida
//con el codigo y guarda en la base de datos esa lista
func (r Rol) FindByCode(code string) (entity.Rol, error) {
	var rol entity.Rol

	rol, _ = r.repo.FindByCode(code)

	if rol.Code == "" {
		list := r.helper.ReadListFromFile()
		err := r.repo.SaveAll(list)
		if err != nil {
			return rol, errors.Wrap(err, "Error al guardar lista en la base de datos")
		}
		for _, value := range list {
			if value.Code == code {
				rol = value
				break
			}
		}
	}

	return rol, nil
}
