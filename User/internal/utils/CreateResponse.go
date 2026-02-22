package utils

import (
	"User/internal/model"
	"encoding/json"
	"net/http"
)

const realm = "Restricted Area"

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.ErrorResponse{Error: message})
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func CreateUserResponse(user *model.User) model.UserResponse {
	return model.UserResponse{
		ID:       user.ID,
		MobileNo: user.MobileNo,
		UserType: user.UserType,
	}
}

func CreateQueryResponse(user *model.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		MobileNo:  user.MobileNo,
		Email:     user.Email,
		UserType:  user.UserType,
		CurStatus: user.CurStatus,
	}
}

func CreateStatusResponse(user *model.User) model.UserStatusResponse {
	return model.UserStatusResponse{
		ID:        user.ID,
		MobileNo:  user.MobileNo,
		UserType:  user.UserType,
		CurStatus: user.CurStatus,
	}
}

func AuthErrorResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Wrong credentials"})
}
