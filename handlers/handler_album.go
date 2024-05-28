package handlers

import (
	"album/internal/database"
	"album/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handler) CreateAlbum(w http.ResponseWriter, r *http.Request, user database.User) {
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

	album, err := h.Cfg.DB.CreateAlbum(r.Context(), database.CreateAlbumParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     params.Title,
		UserID:    user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 500, "Error creating album!")
		return
	}

	utils.RespondWithJSON(w, http.StatusAccepted, utils.DatabaseAlbumToAlbum(album))

}

func (h *Handler) FetchUserAlbums(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")
	userId, err := uuid.Parse(userIdStr)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't parse user id: %v", err))
		return
	}

	albums, err := h.Cfg.DB.FetchUserAlbums(r.Context(), userId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't fetch user albums: %v", err))
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseAlbumsToAlbums(albums))

}

func (h *Handler) FetchAllAlbums(w http.ResponseWriter, r *http.Request) {
	dbAlbums, err := h.Cfg.DB.FetchAllAlbums(r.Context())

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching albums, %v", err))
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabaseAlbumsToAlbums(dbAlbums))
}
