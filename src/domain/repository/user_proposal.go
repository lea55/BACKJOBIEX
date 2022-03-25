package repository

import "BACKJOBIEX/src/domain/entity"

type UserProposal interface {
	Save(up entity.UserProposal) error
	FindByProjectID(projectID string) ([]entity.UserProposal, error)
	FindByUserID(userID string) ([]entity.UserProposal, error)
	FindByID(ID string) (entity.UserProposal, error)
	UpdateStatus(ID string, status entity.UserProposalStatus) error
}

type UserProposalStatus interface {
	FindByCode(code string) (entity.UserProposalStatus, error)
	FindAll() ([]entity.UserProposalStatus, error)
}
