package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/aihmed/user-service/internal/models"
	"github.com/aihmed/user-service/internal/repository"
)

type UserHandler struct {
	repo      repository.UserRepository
	validator *validator.Validate
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo:      repo,
		validator: validator.New(),
	}
}

func generateUUID() string {
	return uuid.New().String()
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Hash password before storing
	user.ID = generateUUID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := h.repo.CreateUser(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Add similar methods for Login, GetUser, etc.
