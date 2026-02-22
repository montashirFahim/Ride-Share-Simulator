package utils

import "User/internal/model"

func AppendResponse(dest *[]model.UserResponse, src []model.User) {
	for _, d := range src {
		*dest = append(*dest, model.UserResponse{
			ID:        d.ID,
			Name:      d.Name,
			MobileNo:  d.MobileNo,
			Email:     d.Email,
			UserType:  d.UserType,
			CurStatus: d.CurStatus,
		})
	}
}
