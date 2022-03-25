package dto

type RqRegisterProject struct {
	CategoryID        string   `json:"category_id" validate:"required"`
	SubCategoryID     string   `json:"sub_category_id" validate:"required"`
	Title             string   `json:"title" validate:"required"`
	Description       string   `json:"description" validate:"required"`
	Properties        []string `json:"properties" validate:"required"`
	Budget            string   `json:"budget" validate:"required"`
	PaymentType       string   `json:"payment_type" validate:"required"`
	CustomerID        string   `json:"customer_id" validate:"required"`
	RequiredSkillsIDS []string `json:"required_skills_ids" validate:"required"`
	DeliveryTime      string   `json:"delivery_time"`
}
