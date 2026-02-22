package utils

import "User/internal/model"

func CreateUser(name, phone, email, userType string) *model.User {
	return &model.User{
		Name:      name,
		MobileNo:  phone,
		Email:     email,
		UserType:  userType,
		CurStatus: "online",
	}
}
