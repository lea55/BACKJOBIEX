package dto

import "github.com/golang-jwt/jwt"

type RqUpdateUserSkill struct {
	UserID         string   `json:"user_id" validate:"required"`
	DevSkillIDList []string `json:"dev_skill_id_list" validate:"required,min=1"`
}

type RqUpdateUserEducation struct {
	Institution string `json:"institution" validate:"required"`
	DegreeType  string `json:"degree_type" validate:"required"`
	DegreeTitle string `json:"degree_title" validate:"required"`
	MonthStart  string `json:"month_start" validate:"required"`
	YearStart   string `json:"year_start" validate:"required"`
	MonthEnd    string `json:"month_end" validate:"required"`
	YearEnd     string `json:"year_end" validate:"required"`
}

type RqUpdateUserProfile struct {
	Names       string `json:"names" validate:"required"`
	Surnames    string `json:"surnames" validate:"required"`
	CategoryID  string `json:"category_id" validate:"required"`
	Country     string `json:"country" validate:"required"`
	City        string `json:"city" validate:"required"`
	CostPerHour uint32 `json:"cost_per_hour" validate:"required"`
	UserID      string `json:"user_id" validate:"required"`
	NickName    string `json:"nick_name" validate:"required"`
}

type RqRegisterUser struct {
	Names    string `json:"names" validate:"required,min=5"`
	Surnames string `json:"surnames"  validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}

type RqAuthentication struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=5"`
	DeviceID string `json:"device_id"`
	CnnIP    string `json:"cnn_ip"`
}

type RqRefreshToken struct {
	Email    string `json:"email" validate:"email"`
	DeviceID string `json:"device_id" validate:"required"`
	CnnIP    string `json:"cnn_ip"`
	TokenID  string `json:"token_id"`
}

type AuthenticatedUser struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Token   string `json:"token"`
	Email   string `json:"email"`
	TokenID string `json:"token_id"`
}

type RqSocialMediaAuth struct {
	Email          string `json:"email" validate:"email"`
	Names          string `json:"names" validate:"required"`
	PlatformToken  string `json:"platform_token"`
	Image          string `json:"image"`
	Type           string `json:"type" validate:"required"`
	PlatformUserID string `json:"platform_user_id"`
	CnnIP          string `json:"cnn_ip"`
	DeviceID       string `json:"device_id"`
}

type TokenUser struct {
	Email   string `json:"email"`
	ID      string `json:"id"`
	RolCode string `json:"rol_code"`
	jwt.StandardClaims
}
