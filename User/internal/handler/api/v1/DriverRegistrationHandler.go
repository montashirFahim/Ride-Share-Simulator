package handler

import (
	"User/internal/model"
	"User/internal/repository"
	"User/internal/utils"
	"encoding/json"
	"net/http"
)

type DriverRegisterHandler struct {
	repo repository.UserRepository
}

func NewDriverRegisterHandler(repo repository.UserRepository) *DriverRegisterHandler {
	return &DriverRegisterHandler{repo: repo}
}

func (h *DriverRegisterHandler) RegisterDriver(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest

	// Validate request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate Bangladeshi phone
	ContactNumber, flag := utils.ContactNumberValidation(req.Phone, "BD")
	if !flag {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid Phone Number")
		return
	}
	req.Phone = ContactNumber

	// E-mail validation
	if !utils.MailValidation(req.Email) {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid E-mail Address")
		return
	}

	//Check e-mail registered or not
	if exists := h.repo.EmailExists(req.Email); exists {
		utils.ErrorResponse(w, http.StatusConflict, "The E-mail is already registered")
		return
	}

	// Check if user already exists
	if exists := h.repo.UserExists(req.Phone); exists {
		utils.ErrorResponse(w, http.StatusConflict, "User already exists")
		return
	}

	//Create User
	user := utils.CreateUser(req.Name, req.Phone, req.Email, "driver")
	if err := h.repo.Create(user); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	//Send Response
	response := utils.CreateUserResponse(user)
	utils.SuccessResponse(w, response)

}
