package handler

import (
	"User/internal/repository"
	"User/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type StatusRequest struct {
	Status string `json:"status"`
}

type StatusHandler struct {
	repo repository.UserRepository
}

func NewStatusHandler(repo repository.UserRepository) *StatusHandler {
	return &StatusHandler{repo: repo}
}

func (h *StatusHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	var req StatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Status != "online" && req.Status != "offline" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid Status")
		return
	}

	user, err := h.repo.GetByID(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	if user.UserType != "driver" {
		utils.ErrorResponse(w, http.StatusBadRequest, "User is not a driver")
		return
	}

	user.CurStatus = req.Status
	if err := h.repo.UpdateStatus(user.ID, user.CurStatus); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed to update status")
		return
	}

	response := utils.CreateStatusResponse(user)
	utils.SuccessResponse(w, response)
}
