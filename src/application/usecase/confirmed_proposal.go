package usecase

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lea55/BACKJOBIEX/src/application/dto"
	"github.com/lea55/BACKJOBIEX/src/application/helpers"
	"github.com/lea55/BACKJOBIEX/src/application/mappers"
	"github.com/lea55/BACKJOBIEX/src/domain/entity"
	"github.com/lea55/BACKJOBIEX/src/domain/platforms"
	"github.com/lea55/BACKJOBIEX/src/domain/repository"
	"github.com/pkg/errors"
)

type ProposalConfirmation struct {
	repository              repository.ConfirmedProposal
	proposalRepo            repository.UserProposal
	proposalStsRepo         repository.UserProposalStatus
	userRepo                repository.User
	paymentMethodRepo       repository.PaymentMethod
	paymentStatusRepo       repository.PaymentStatus
	codeHelper              *helpers.Code
	userMapper              *mappers.User
	projectRepo             repository.Project
	projectStatusRepo       repository.ProjectStatus
	mercadoPago             platforms.MercadoPago
	proposalConfPaymentRepo repository.ProposalConfirmationPayment
}

func NewProposalConfirmation(
	repository repository.ConfirmedProposal,
	proposalRepo repository.UserProposal,
	userRepo repository.User,
	paymentMethodRepo repository.PaymentMethod,
	paymentStatusRepo repository.PaymentStatus,
	proposalStatusRepo repository.UserProposalStatus,
	projectRepo repository.Project,
	projectStatusRepo repository.ProjectStatus,
	mercadoPago platforms.MercadoPago,
	proposalConfPaymentRepo repository.ProposalConfirmationPayment,
) *ProposalConfirmation {
	return &ProposalConfirmation{
		mercadoPago:             mercadoPago,
		projectRepo:             projectRepo,
		repository:              repository,
		proposalRepo:            proposalRepo,
		userRepo:                userRepo,
		paymentMethodRepo:       paymentMethodRepo,
		paymentStatusRepo:       paymentStatusRepo,
		codeHelper:              helpers.NewCode(),
		userMapper:              mappers.NewUser(),
		proposalStsRepo:         proposalStatusRepo,
		projectStatusRepo:       projectStatusRepo,
		proposalConfPaymentRepo: proposalConfPaymentRepo,
	}
}

func (p ProposalConfirmation) ConfirmProposal(model dto.RqConfirmProposal) (string, error) {
	proposal, err := p.proposalRepo.FindByID(model.ProposalID)
	if err != nil {
		return "", errors.Wrap(err, "Error En validación de propuesta")
	}

	if proposal.Status.Code == coreconstants.NewUserProposalStatusCode().Approved {
		return "", errors.New("La propuesta ya ha sido aceptada")
	}

	_, err, proposalExist := p.repository.FindByProposalID(model.ProposalID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de existencia de propuesta")
	}

	if proposalExist {
		return "", errors.New("La propuesta ya ha sido aceptada")
	}

	freelancer, err := p.userRepo.FindByID(proposal.User.ID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de freelancer")
	}

	customer, err := p.userRepo.FindByID(model.RequestUserID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de cliente")
	}

	paymentStatus, err := p.paymentStatusRepo.FindByCode(coreconstants.PaymentStsPending)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de estado de pago")
	}

	paymentMethod, err := p.paymentMethodRepo.FindByCode(coreconstants.PaymentMthUndefined)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de metodo de pago")
	}

	totalDocs, err := p.repository.Count()
	if err != nil {
		return "", errors.Wrap(err, "Error al consultar número de documentos")
	}

	codePrefix := "P1" + customer.Names[0:1] + freelancer.Names[0:1]

	code := strings.ToUpper(p.codeHelper.Generate(totalDocs+1, codePrefix, 5))

	timeNow := time.Now().UTC()
	dueDate := timeNow.AddDate(0, 0, int(proposal.DaysToDelivery*1))

	confirmedProposal := entity.ConfirmedProposal{
		ID:             uuid.New().String(),
		Code:           code,
		Customer:       p.userMapper.GetBasicFromEntity(customer),
		Freelancer:     p.userMapper.GetBasicFromEntity(freelancer),
		UsPrice:        proposal.BudgetUSD,
		DaysToDelivery: proposal.DaysToDelivery,
		DueDate:        dueDate,
		PaymentStatus:  paymentStatus,
		PaymentMethod:  paymentMethod,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		ProposalID:     model.ProposalID,
	}

	err = p.repository.Save(confirmedProposal)
	if err != nil {
		return "", errors.Wrap(err, "Error al guardar documento en la base de datos")
	}

	err = p.updateProposalStatus(coreconstants.NewUserProposalStatusCode().Approved, proposal.ID)
	if err != nil {
		return "", err
	}

	err = p.updateProjectStatus(coreconstants.NewProjectStatus().InProgress, proposal.ProjectID)
	if err != nil {
		return "", err
	}

	return confirmedProposal.ID, nil
}

func (p ProposalConfirmation) updateProposalStatus(statusCode string, proposalID string) error {
	proposalStatus, err := p.proposalStsRepo.FindByCode(statusCode)
	if err != nil {
		return errors.Wrap(err, "Error en validación de estado de propuesta")
	}

	err = p.proposalRepo.UpdateStatus(proposalID, proposalStatus)
	if err != nil {
		return errors.Wrap(err, "Error al actualizar estado")
	}

	return nil
}

func (p ProposalConfirmation) updateProjectStatus(statusCode string, projectID string) error {
	pjStatus, err := p.projectStatusRepo.GetByCode(statusCode)
	if err != nil {
		return errors.Wrap(err, "Error al consultar estado de proyecto")
	}

	err = p.projectRepo.UpdateStatus(projectID, pjStatus)
	if err != nil {
		return errors.Wrap(err, "Error en actualización de estado de proyecto")
	}

	return nil
}

func (p ProposalConfirmation) GeneratePaymentLink(proposalID string) (string, error) {

	proposal, err := p.proposalRepo.FindByID(proposalID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de propuesta")
	}

	if proposal.Status.Code != coreconstants.NewUserProposalStatusCode().Approved {
		return "", errors.New("La propuesta aun no se ha aprobado")
	}

	proposalConfirmation, err, founded := p.repository.FindByProposalID(proposalID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de de confirmación")
	}

	if !founded {
		return "", errors.New("No existe una confirmación de propuesta")
	}

	project, err := p.projectRepo.FindByID(proposal.ProjectID)
	if err != nil {
		return "", errors.Wrap(err, "Error en validación de proyecto")
	}

	model := entity.GeneratePaymentPreference{
		Title: "Deposito para proyecto " + project.Title + " " + project.Code,
		Description: "Con esta acción estas haciendo el deposito para que el freelancer " +
			proposalConfirmation.Freelancer.Names + " Empiece a trabajar",
		Quantity:            1,
		UnitPrice:           proposalConfirmation.UsPrice,
		ConfirmedProposalID: proposalConfirmation.ID,
	}

	result, err := p.mercadoPago.GenerateProductPayment(model)
	if err != nil {
		return "", errors.Wrap(err, "Error en generación de link")
	}

	return result.GeneratedLink, nil
}

func (p ProposalConfirmation) PaymentConfirmation(proposalConfirmationID string, reference string) error {
	proposalConfirmation, err := p.repository.FindByID(proposalConfirmationID)
	if err != nil {
		return errors.Wrap(err, "Error en validación de confirmación de propuesta")
	}

	payment := entity.ProposalPayment{
		ID:                     uuid.New().String(),
		Reference:              reference,
		ProposalConfirmationID: proposalConfirmation.ID,
		ProposalID:             proposalConfirmation.ProposalID,
		Date:                   time.Now().UTC(),
		User:                   proposalConfirmation.Customer,
	}

	proposal, err := p.proposalRepo.FindByID(proposalConfirmation.ProposalID)
	if err != nil {
		return errors.Wrap(err, "Error en validación de propuesta")
	}

	err = p.proposalConfPaymentRepo.Save(payment)
	if err != nil {
		return errors.Wrap(err, "Error al guarda pago en la app")
	}

	err = p.updateProposalStatus(coreconstants.NewUserProposalStatusCode().DevProgress, proposalConfirmation.ProposalID)
	if err != nil {
		return err
	}

	err = p.updateProjectStatus(coreconstants.NewProjectStatus().InProgress, proposal.ProjectID)
	if err != nil {
		return err
	}

	status, err := p.paymentStatusRepo.FindByCode(coreconstants.PaymentStsConfirmed)
	if err != nil {
		return errors.Wrap(err, "Error en validación de estado de pago")
	}

	paymentMethod, err := p.paymentMethodRepo.FindByCode(coreconstants.PaymentMthMercadoPago)
	if err != nil {
		return errors.Wrap(err, "Error en validación de metodo de pago")
	}

	err = p.repository.UpdatePaymentMethod(proposalConfirmationID, paymentMethod)
	if err != nil {
		return errors.Wrap(err, "Error en actualización de metodo de paggo")
	}

	err = p.repository.UpdateStatus(proposalConfirmationID, status)
	if err != nil {
		return errors.Wrap(err, "Error en actualización de estado de confirmación")
	}

	return nil
}
