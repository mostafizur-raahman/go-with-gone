package handlers

import (
	"net/http"

	"github.com/mostafizur-raahman/go-with-gone/internal/models"
	"github.com/mostafizur-raahman/go-with-gone/internal/services"
	"github.com/mostafizur-raahman/go-with-gone/pkg"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := pkg.ParseJSONRequest(w, r, &user); err != nil {
		return
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	user.ID = id

	pkg.JSONResponse(w, http.StatusCreated, user)
}

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	// Call service to get all users
	users, err := h.service.AllUsers()
	if err != nil {
		pkg.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// If no users found, return empty array instead of null
	if len(users) == 0 {
		users = []models.User{}
	}

	// Return users in JSON format
	pkg.JSONResponse(w, http.StatusOK, users)
}
