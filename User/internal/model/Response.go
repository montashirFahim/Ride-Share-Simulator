package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	MobileNo  string `json:"mobile_no"`
	Email     string `json:"email,omitempty"`
	UserType  string `json:"user_type"`
	CurStatus string `json:"cur_status,omitempty"`
}
