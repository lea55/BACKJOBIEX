package usecase

import (
	"BACKJOBIEX/src/core/constants"

	"BACKJOBIEX/src/application/dto"
	"BACKJOBIEX/src/application/helpers"
	"BACKJOBIEX/src/application/mappers"
	"BACKJOBIEX/src/domain/entity"
	"BACKJOBIEX/src/domain/repository"

	"github.com/pkg/errors"
)

type Project struct {
	repo         repository.Project
	userRepo     repository.User
	cateRepo     repository.ProjectCategoryRepo
	subCatRepo   repository.ProjectSubCategory
	statusRepo   repository.ProjectStatus
	devSkillRepo repository.DevSkill
	mapper       *mappers.Project
	codeHelp     *helpers.Code
	userMapper   *mappers.User
}

func NewProjectUseCase(repo repository.Project, userRepo repository.User, cateRepo repository.ProjectCategoryRepo,
	subCatRepo repository.ProjectSubCategory, statusRepo repository.ProjectStatus,
	devSkill repository.DevSkill) *Project {
	return &Project{
		userMapper:   mappers.NewUser(),
		codeHelp:     helpers.NewCode(),
		mapper:       mappers.NewProject(),
		repo:         repo,
		userRepo:     userRepo,
		cateRepo:     cateRepo,
		subCatRepo:   subCatRepo,
		statusRepo:   statusRepo,
		devSkillRepo: devSkill,
	}
}

func (p Project) FindPaginated(page uint16) ([]entity.Project, error) {
	list := make([]entity.Project, 0)

	if page < 0 {
		return list, errors.New("La página enviada no es válida")
	}

	limit := uint32(page * 10)
	from := limit - 10

	list, err := p.repo.FindPaginated(from, limit)
	if err != nil {
		return nil, errors.Wrap(err, "Error en la busqueda en la base de datos")
	}

	return list, nil
}

func (p Project) FindByID(ID string) (entity.Project, error) {
	result, err := p.repo.FindByID(ID)
	if err != nil {
		return result, errors.Wrap(err, "Error en la consulta a la base de datos")
	}

	return result, nil
}

func (p Project) GetByCustomerID(customerID string) ([]entity.Project, error) {
	resultList := make([]entity.Project, 0)

	_, err := p.userRepo.FindByID(customerID)
	if err != nil {
		return resultList, errors.Wrap(err, "Error en validación de usuario en la base de datos")
	}

	resultList, err = p.repo.FindByCustomerID(customerID)
	if err != nil {
		return resultList, errors.Wrap(err, "Error en consulta de listado a la base de datos")
	}

	return resultList, nil
}

func (p Project) Save(rqModel dto.RqRegisterProject) (string, error) {

	category, err := p.cateRepo.FindByID(rqModel.CategoryID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de categoría en la base de datos")
	}

	subCategory, err := p.subCatRepo.FindByID(rqModel.SubCategoryID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de sub categoría en la base de datos")
	}

	customer, err := p.userRepo.FindByID(rqModel.CustomerID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de usuario en la base de datos")
	}

	status, err := p.statusRepo.GetByCode(constants.NewProjectStatus().Opened)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de usuario en la base de datos")
	}

	devSkillList := make([]entity.DevSkill, 0)
	for _, id := range rqModel.RequiredSkillsIDS {
		devSkill, findDevSkillErr := p.devSkillRepo.FindByID(id)
		if findDevSkillErr != nil {
			return "", errors.Wrap(findDevSkillErr, "Error en validación de campo")
		}
		devSkillList = append(devSkillList, devSkill)
	}

	count, err := p.repo.Count()
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de datos actuales")
	}

	code := p.codeHelp.Generate(count+1, "PJC", 6)

	newProject := p.mapper.GetFromRqModel(rqModel, code, category, subCategory, customer, devSkillList, status)

	idSaved, err := p.repo.Save(newProject)
	if err != nil {
		return "", errors.Wrap(err, "Error al guardar proyecto")
	}

	return idSaved, nil
}

func (p Project) AddProposalItems(projectID string, model entity.UserProposal) error {
	project, err := p.FindByID(projectID)
	if err != nil {
		return errors.Wrap(err, "Error en validación de proyecto ")
	}

	newList := make([]entity.ProjectProposalItem, 0)

	if project.Proposals != nil {
		for _, v := range project.Proposals {
			newList = append(newList, v)
		}
	}

	newProposal := entity.ProjectProposalItem{
		User:          model.User,
		ProposalID:    model.ID,
		ProposalTitle: model.Title,
	}

	newList = append(newList, newProposal)

	err = p.repo.UpdateProposals(project.ID, newList)
	if err != nil {
		return errors.Wrap(err, "Error al guardar propuestas en la base de datos")
	}

	return nil
}

func (p Project) UpdateProposalAverage(projectID string, proposalAverage float32) error {
	project, err := p.FindByID(projectID)
	if err != nil {
		return errors.Wrap(err, "Error en validación de proyecto ")
	}

	err = p.repo.UpdateProposalAverage(project.ID, proposalAverage)
	if err != nil {
		return errors.Wrap(err, "Error en actualización de promedio de propuesta")
	}

	return nil
}
