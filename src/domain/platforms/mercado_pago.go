package platforms

import "github.com/lea55/BACKJOBIEX/src/domain/entity"

type MercadoPago interface {
	GenerateProductPayment(model entity.GeneratePaymentPreference) (entity.GeneratePaymentPreferenceResponse, error)
}
