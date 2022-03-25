package entity

import "time"

type User struct {
	ID              string      `json:"id"`
	Names           string      `json:"names"`
	Surnames        string      `json:"surnames"`
	Email           string      `json:"email"`
	Password        string      `json:"password"`
	LastSessionIp   string      `json:"last_session_ip"`
	LastSessionDate time.Time   `json:"last_session_date"`
	Image           string      `json:"image"`
	Contact         UserContact `json:"contact"`
	Enabled         bool        `json:"enabled"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Rol             Rol         `json:"rol"`
	DeviceID        string      `json:"device_id"`

	//new thigns
	NickName          string          `json:"nick_name"`
	Category          ProjectCategory `json:"category"`
	HourlyCharge      uint16          `json:"hourly_charge"`
	CoverImage        string          `json:"cover_image"`
	ProfilePhoto      string          `json:"profile_photo"`
	Education         []UserEducation `json:"education"`
	Skills            []DevSkill      `json:"skills"`
	CompletedProjects uint16          `json:"completed_projects"`
}

type BasicUser struct {
	Names        string `json:"names"`
	Surnames     string `json:"surnames"`
	ID           string `json:"id"`
	Email        string `json:"email"`
	RolCode      string `json:"rol_code"`
	ProfileImage string `json:"image"`
}

type UserContact struct {
	Country     string `json:"country"`
	City        string `json:"city"`
	CountryCode string `json:"country_code"`
	Currency    string `json:"currency"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
}

type UserEducation struct {
	Institution string `json:"institution"`
	DegreeType  string `json:"degree_type"`
	DegreeTitle string `json:"degree_title"`
	MonthStart  string `json:"month_start"`
	YearStart   string `json:"year_start"`
	MonthEnd    string `json:"month_end"`
	YearEnd     string `json:"year_end"`
}
