package model

type RegisterRequest struct {
	Phone    string `json:"phone"`
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	UserType string `json:"user_type,omitempty"`
}
