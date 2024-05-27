package handlers

import (
	"album/config"
	"album/internal/database"
	"album/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Cfg *config.ApiConfig
}

func (h *Handler) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Couldn't hash password: %v", err))
		return
	}

	exists, err := h.Cfg.DB.IsUniqueEmail(r.Context(), params.Email)

	if err != nil {
		utils.RespondWithError(w, 500, "Error checking unique email!")
		return
	}

	if exists != 0 {
		utils.RespondWithError(w, 400, "User with email already exists!")
		return

	}

	exists, err = h.Cfg.DB.IsUniqueUsername(r.Context(), params.Username)

	if err != nil {
		utils.RespondWithError(w, 400, "Error checking unique username!")
		return
	}

	if exists != 0 {
		utils.RespondWithError(w, 400, "User with username already exists!")
		return

	}

	user, err := h.Cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Username:  params.Username,
		Email:     params.Email,
		Password:  string(hashedPassword),
	})

	if err != nil {
		utils.RespondWithError(w, 500, "Error creating user!")
		return
	}
	utils.RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))
}
