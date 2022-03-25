package entity

import "time"

type Project struct {
	ID              string                `json:"id"`
	Code            string                `json:"code"`
	Category        ProjectCategory       `json:"category"`
	SubCategory     ProjectSubCategory    `json:"sub_category"`
	Title           string                `json:"title"`
	Description     string                `json:"description"`
	Properties      []string              `json:"properties"`
	Budget          string                `json:"budget"`
	PaymentType     string                `json:"payment_type"`
	Customer        BasicUser             `json:"customer"`
	Status          ProjectStatus         `json:"status"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
	RequiredSkills  []DevSkill            `json:"required_skills"`
	Enabled         bool                  `json:"enabled"`
	Proposals       []ProjectProposalItem `json:"proposals"`
	City            string                `json:"city"`
	Country         string                `json:"country"`
	DeliveryTime    string                `json:"delivery_time"`
	ProposalAverage float32               `json:"proposal_average"`
}

type ProjectStatus struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ProjectProposalItem struct {
	User          BasicUser `json:"user"`
	ProposalID    string    `json:"proposal_id"`
	ProposalTitle string    `json:"proposal_title"`
}
