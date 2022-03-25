package dto

type RqRegisterProjectSubCat struct {
	Name        string                   `json:"name" validate:"required"`
	Description string                   `json:"description" validate:"required"`
	CategoryID  string                   `json:"category_id" validate:"required"`
	Items       []RqRegProjectSubCatItem `json:"items"`
}

type RqRegProjectSubCatItem struct {
	Title   string   `json:"title" validate:"required"`
	Options []string `json:"options" validate:"required"`
}
