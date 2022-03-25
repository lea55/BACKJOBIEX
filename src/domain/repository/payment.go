package repository

import "github.com/lea55/BACKJOBIEX/src/domain/entity"

type PaymentStatus interface {
	FindByCode(code string) (entity.PaymentStatus, error)
}

type PaymentMethod interface {
	FindByCode(code string) (entity.PaymentMethod, error)
}

type ProposalConfirmationPayment interface {
	Save(doc entity.ProposalPayment) error
}
