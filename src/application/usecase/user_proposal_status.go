package usecase

import (
	"BACKJOBIEX/src/domain/entity"
	"BACKJOBIEX/src/domain/repository"

	"github.com/pkg/errors"
)

type UserProposalStatus struct {
	repo repository.UserProposalStatus
}

func NewUserProposalStatus(repo repository.UserProposalStatus) *UserProposalStatus {
	return &UserProposalStatus{repo: repo}
}

func (s UserProposalStatus) GetAll() ([]entity.UserProposalStatus, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return list, errors.Wrap(err, "Error en consulta a la base de datos")
	}

	return list, err
}
