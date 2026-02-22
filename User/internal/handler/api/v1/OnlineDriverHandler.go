package handler

import (
	"User/internal/model"
	"User/internal/repository"
	"User/internal/utils"
	"net/http"
)

type DriverOnlineHandler struct {
	repo repository.UserRepository
}

func NewDriverOnlineHandler(repo repository.UserRepository) *DriverOnlineHandler {
	return &DriverOnlineHandler{repo: repo}
}

func (h *DriverOnlineHandler) ListOnlineDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := h.repo.ListDrivers("online")
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var resp []model.UserResponse
	utils.AppendResponse(&resp, drivers)

	utils.SuccessResponse(w, resp)
}
