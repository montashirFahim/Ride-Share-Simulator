package handler

import (
	"User/internal/repository"
	"User/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type InfoHandler struct {
	UserQuery *repository.UserQuery
}

func NewInfoHandler(userQuery *repository.UserQuery) *InfoHandler {
	return &InfoHandler{UserQuery: userQuery}
}

func (h *InfoHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid User ID format")
		return
	}
	user, err := h.UserQuery.GetUser(id)
	if err != nil {
		log.Printf("Error fetching user %d: %v", id, err)
		utils.ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}
	utils.SuccessResponse(w, user)
}
