package platforms

import "BACKJOBIEX/src/domain/entity"

type MercadoPago interface {
	GenerateProductPayment(model entity.GeneratePaymentPreference) (entity.GeneratePaymentPreferenceResponse, error)
}
