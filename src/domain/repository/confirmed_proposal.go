package repository

import "BACKJOBIEX/src/domain/entity"

type ConfirmedProposal interface {
	Save(doc entity.ConfirmedProposal) error
	FindByProposalID(proposalID string) (entity.ConfirmedProposal, error, bool)
	Count() (uint32, error)
	FindByID(ID string) (entity.ConfirmedProposal, error)
	UpdatePaymentMethod(ID string, method entity.PaymentMethod) error
	UpdateStatus(ID string, status entity.PaymentStatus) error
}
