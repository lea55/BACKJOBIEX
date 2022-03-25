package entity

import "time"

type ConfirmedProposal struct {
	ID             string        `json:"id"`
	Code           string        `json:"code"`
	ProposalID     string        `json:"proposal_id"`
	Customer       BasicUser     `json:"customer"`
	Freelancer     BasicUser     `json:"freelancer"`
	UsPrice        float32       `json:"us_price"`
	DaysToDelivery uint16        `json:"days_to_delivery"`
	DueDate        time.Time     `json:"due_date"`
	PaymentStatus  PaymentStatus `json:"payment_status"`
	PaymentMethod  PaymentMethod `json:"payment_method"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type ConfirmedProposalUpdate struct {
	ID         string `json:"id"`
	ProposalID string `json:"proposal_id"`
	UpdateNote string `json:"update_note"`
	CreatedAt  string `json:"created_at"`
}

type GeneratePaymentPreference struct {
	Title               string
	Description         string
	Quantity            uint8
	UnitPrice           float32
	ConfirmedProposalID string
}

type GeneratePaymentPreferenceResponse struct {
	ReferenceID   string
	GeneratedLink string
}
