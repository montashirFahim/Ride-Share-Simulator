package model

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	MobileNo  string `json:"mobile_no"`
	Email     string `json:"email"`
	UserType  string `json:"user_type"`  // driver or rider
	CurStatus string `json:"cur_status"` // online or offline
}
