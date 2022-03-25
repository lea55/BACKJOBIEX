package entity

type ProjectSubCategory struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Items       []SubCategoryItem `json:"items"`
	Image       string            `json:"image"`
	Category    ProjectCategory   `json:"category"`
}

type SubCategoryItem struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
}
