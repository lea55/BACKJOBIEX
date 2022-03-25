package entity

type PaymentMethod struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
