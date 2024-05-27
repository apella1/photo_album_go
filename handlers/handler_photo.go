package handlers

import (
	"album/internal/database"
	"album/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) CreatePhoto(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title string `json:"title"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	photo, err := h.Cfg.DB.CreatePhoto(r.Context(), database.CreatePhotoParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     params.Title,
		UserID:    user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 500, "Error creating photo!")
		return
	}

	utils.RespondWithJSON(w, http.StatusAccepted, utils.DatabasePhotoToPhoto(photo))
}
