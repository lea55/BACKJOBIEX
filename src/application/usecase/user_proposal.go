package usecase

import (
	"github.com/jdpadillaac/jobiex-api/src/application/dto"
	"github.com/jdpadillaac/jobiex-api/src/application/mappers"
	"github.com/jdpadillaac/jobiex-api/src/core/constants"
	"github.com/jdpadillaac/jobiex-api/src/domain/entity"
	"github.com/jdpadillaac/jobiex-api/src/domain/repository"
	"github.com/pkg/errors"
)

type UserProposal struct {
	repository             repository.UserProposal
	userProposalStatusRepo repository.UserProposalStatus
	userRepo               repository.User
	projectRepo            repository.Project
	mapper                 *mappers.UserProposalMapper
	projectUc              *Project
}

func NewUserProposal(
	repository repository.UserProposal,
	userProposalStatusRepo repository.UserProposalStatus,
	user repository.User,
	projectRepo repository.Project,
	projectUc *Project,
) *UserProposal {
	return &UserProposal{
		repository:             repository,
		userProposalStatusRepo: userProposalStatusRepo,
		userRepo:               user,
		projectRepo:            projectRepo,
		mapper:                 mappers.NewUserProposalMapper(),
		projectUc:              projectUc,
	}
}

func (p UserProposal) FindByUserID(userID string) ([]entity.UserProposal, error) {
	user, err := p.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.Wrap(err, "El usuario no existe")
	}

	result, err := p.repository.FindByUserID(user.ID)
	if err != nil {
		return result, errors.Wrap(err, "Error en la base de datos")
	}

	return result, nil
}

func (p UserProposal) FindByProjectID(projectID string) ([]entity.UserProposal, error) {
	project, err := p.projectUc.FindByID(projectID)
	if err != nil {
		return nil, err
	}

	result, err := p.repository.FindByProjectID(project.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p UserProposal) Save(rqModel dto.RqUserProposalRegister) error {
	user, err := p.userRepo.FindByID(rqModel.UserID)
	if err != nil {
		return errors.Wrap(err, "Error en validación de usuario que registra ")
	}

	project, err := p.projectRepo.FindByID(rqModel.ProjectID)
	if err != nil {
		return errors.Wrap(err, "Error en validación de proyecto")
	}

	err = p.validateIfProposalExist(project, user.ID)
	if err != nil {
		return err
	}

	status, err := p.userProposalStatusRepo.FindByCode(coreconstants.NewUserProposalStatusCode().Created)
	if err != nil {
		return errors.Wrap(err, "Error en validación de estado de propuesta")
	}

	proposal := p.mapper.FromRegisterModel(rqModel, status, user)

	err = p.repository.Save(proposal)
	if err != nil {
		return errors.Wrap(err, "Error al guardar propuesta en la base de datos")
	}

	err = p.projectUc.AddProposalItems(rqModel.ProjectID, proposal)
	if err != nil {
		return errors.Wrap(err, "Error al actualizar datos de propuesta")
	}

	err = p.updateProposalAverage(project.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p UserProposal) validateIfProposalExist(project entity.Project, userID string) error {
	if project.Proposals != nil && len(project.Proposals) > 0 {
		for _, v := range project.Proposals {
			if v.User.ID == userID {
				return errors.New("Ya has registrado una propuesta anteriormente.")
			}
		}
	}
	return nil
}

func (p UserProposal) updateProposalAverage(projectID string) error {
	projectProposals, err := p.repository.FindByProjectID(projectID)
	if err != nil {
		return errors.Wrap(err, "Error en consulta de propuesta registradas")
	}

	var proposalSum float32

	for _, v := range projectProposals {
		proposalSum = proposalSum + v.BudgetUSD
	}

	proposalAverage := proposalSum / float32(len(projectProposals))

	err = p.projectUc.UpdateProposalAverage(projectID, proposalAverage)
	if err != nil {
		return err
	}

	return nil
}
