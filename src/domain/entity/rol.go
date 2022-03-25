package entity

type Rol struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RolCode struct {
	Admin string
	User  string
}

func NewUserRol() *RolCode {
	return &RolCode{Admin: "ADMIN", User: "USER"}
}
