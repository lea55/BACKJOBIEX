package usecase

import (
	"github.com/jdpadillaac/jobiex-api/src/domain/entity"
	"github.com/jdpadillaac/jobiex-api/src/domain/repository"
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
