package model

type UserStatusResponse struct {
	ID        int    `json:"id"`
	MobileNo  string `json:"mobile_no"`
	UserType  string `json:"user_type"`  // driver or rider
	CurStatus string `json:"cur_status"` // online or offline
}
