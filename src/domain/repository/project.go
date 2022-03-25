package repository

import "BACKJOBIEX/src/domain/entity"

type Project interface {
	Save(project entity.Project) (string, error)
	Count() (uint32, error)
	FindByCustomerID(customerID string) ([]entity.Project, error)
	FindByID(ID string) (entity.Project, error)
	FindPaginated(from uint32, to uint32) ([]entity.Project, error)
	UpdateProposals(ID string, items []entity.ProjectProposalItem) error
	UpdateProposalAverage(ID string, average float32) error
	UpdateStatus(projectID string, status entity.ProjectStatus) error
}

type ProjectStatus interface {
	GetAll() ([]entity.ProjectStatus, error)
	GetByCode(code string) (entity.ProjectStatus, error)
}
