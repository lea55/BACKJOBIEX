package mappers

import (
	"time"

	"github.com/lea55/BACKJOBIEX/src/application/dto"
	"github.com/lea55/BACKJOBIEX/src/domain/entity"
)

type Project struct{}

func NewProject() *Project {
	return &Project{}
}

func (p Project) GetFromRqModel(model dto.RqRegisterProject, code string, category entity.ProjectCategory,
	subCategory entity.ProjectSubCategory, customer entity.User, devSkillList []entity.DevSkill,
	status entity.ProjectStatus) entity.Project {

	return entity.Project{
		Code:        code,
		Category:    category,
		SubCategory: subCategory,
		Title:       model.Title,
		Description: model.Description,
		Properties:  model.Properties,
		Budget:      model.Budget,
		PaymentType: model.PaymentType,
		Customer: entity.BasicUser{
			Names:    customer.Names,
			Surnames: customer.Surnames,
			ID:       customer.ID,
			Email:    customer.Email,
			RolCode:  customer.Rol.Code,
		},
		Status:         status,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		RequiredSkills: devSkillList,
		Enabled:        true,
		Proposals:      nil,
		City:           customer.Contact.City,
		Country:        customer.Contact.Country,
		DeliveryTime:   model.DeliveryTime,
	}
}
