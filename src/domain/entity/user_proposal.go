package entity

import "time"

type UserProposal struct {
	ID             string             `json:"id"`
	BudgetUSD      float32            `json:"budget_usd"`
	DaysToDelivery uint16             `json:"days_to_delivery"`
	Title          string             `json:"title"`
	Description    string             `json:"description"`
	Status         UserProposalStatus `json:"status"`
	Requirements   string             `json:"requirements"`
	ProjectID      string             `json:"project_id"`
	User           BasicUser          `json:"user"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdateAt       time.Time          `json:"update_at"`
}

type UserProposalStatus struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
