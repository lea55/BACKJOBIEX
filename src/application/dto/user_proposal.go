package dto

type RqUserProposalRegister struct {
	BudgetUSD      float32 `json:"budget_usd" validate:"required"`
	DaysToDelivery uint16  `json:"days_to_delivery" validate:"required"`
	Title          string  `json:"title" validate:"required"`
	Description    string  `json:"description" validate:"required"`
	Requirements   string  `json:"requirements" validate:"required"`
	ProjectID      string  `json:"project_id" validate:"required"`
	UserID         string  `json:"user_id" validate:"required"`
}

type RqConfirmProposal struct {
	ProposalID    string `json:"proposal_id" validate:"required"`
	RequestUserID string `json:"request_user_id" validate:"required"`
}
