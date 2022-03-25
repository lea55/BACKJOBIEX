package mappers

import (
	"time"

	"github.com/lea55/BACKJOBIEX/src/application/dto"
	"github.com/lea55/BACKJOBIEX/src/domain/entity"

	"github.com/google/uuid"
)

type UserProposalMapper struct {
	userMapper *User
}

func NewUserProposalMapper() *UserProposalMapper {
	return &UserProposalMapper{
		userMapper: NewUser(),
	}
}

func (m UserProposalMapper) FromRegisterModel(
	register dto.RqUserProposalRegister,
	status entity.UserProposalStatus,
	user entity.User,

) entity.UserProposal {
	ID := uuid.New().String()
	return entity.UserProposal{
		ID:             ID,
		BudgetUSD:      register.BudgetUSD,
		DaysToDelivery: register.DaysToDelivery,
		Title:          register.Title,
		Description:    register.Description,
		Status:         status,
		Requirements:   register.Requirements,
		ProjectID:      register.ProjectID,
		User:           m.userMapper.GetBasicFromEntity(user),
		CreatedAt:      time.Now().UTC(),
		UpdateAt:       time.Now().UTC(),
	}
}
