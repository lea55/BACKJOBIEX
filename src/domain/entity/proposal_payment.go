package entity

import "time"

type ProposalPayment struct {
	ID                     string    `json:"id"`
	Reference              string    `json:"reference"`
	ProposalConfirmationID string    `json:"proposal_confirmation_id"`
	ProposalID             string    `json:"proposal_id"`
	Date                   time.Time `json:"date"`
	User                   BasicUser `json:"user"`
}
