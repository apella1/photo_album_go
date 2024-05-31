package handlers

import (
	"album/internal/database"
	"album/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handler) CreatePhoto(w http.ResponseWriter, r *http.Request, user database.User) {
	albumIdStr := chi.URLParam(r, "albumId")
	albumId, err := uuid.Parse(albumIdStr)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't parse album id: %v", err))
		return
	}

	const maxUploadSize = 400 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	err = r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing multipart form: %v", err))
		return
	}

	title := r.FormValue("title")
	if title == "" {
		utils.RespondWithError(w, 400, "Title is required")
		return
	}

	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		if err == http.ErrMissingFile {
			utils.RespondWithError(w, 400, "No file uploaded!")
			return
		}
		utils.RespondWithError(w, 500, fmt.Sprintf("Error getting uploaded file: %v", err))
		return
	}
	defer file.Close()

	if fileHeader.Size > maxUploadSize {
		utils.RespondWithError(w, 400, "File size exceeds 400KB limit!")
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error reading uploaded file: %v", err))
		return
	}

	album, err := h.Cfg.DB.GetAlbumById(r.Context(), albumId)
	if err != nil {
		utils.RespondWithError(w, 400, "Error retrieving album!")
		return
	}

	if album.UserID != user.ID {
		utils.RespondWithError(w, 403, "Unauthorized! You can only upload photos to your own albums!")
		return
	}

	base64URL := fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(fileBytes))

	photo, err := h.Cfg.DB.CreatePhoto(r.Context(), database.CreatePhotoParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     title,
		Body:      fileBytes,
		UserID:    user.ID,
		AlbumID:   albumId,
		ImgUrl:    base64URL,
	})

	if err != nil {
		utils.RespondWithError(w, 500, "Error creating photo!")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, utils.DatabasePhotoToPhoto(photo))
}

func (h *Handler) FetchAlbumPhotos(w http.ResponseWriter, r *http.Request) {
	albumIdStr := chi.URLParam(r, "albumId")
	albumId, err := uuid.Parse(albumIdStr)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't parse album id: %v", err))
		return
	}

	photos, err := h.Cfg.DB.FetchAlbumPhotos(r.Context(), albumId)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching album photos, %v", err))
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabasePhotosToPhotos(photos))
}

func (h *Handler) FetchPhoto(w http.ResponseWriter, r *http.Request) {
	photoIdStr := chi.URLParam(r, "photoId")
	photoId, err := uuid.Parse(photoIdStr)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't parse photo id: %v", err))
		return
	}

	photo, err := h.Cfg.DB.FetchPhoto(r.Context(), photoId)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching photo, %v", err))
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabasePhotoToPhoto(photo))
}

func (h *Handler) UpdatePhotoTitle(w http.ResponseWriter, r *http.Request, user database.User) {
	photoIdStr := chi.URLParam(r, "photoId")
	photoId, err := uuid.Parse(photoIdStr)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't parse photo id: %v", err))
		return
	}

	type parameters struct {
		Title string `json:"title"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	err = h.Cfg.DB.UpdatePhotoTitle(r.Context(), database.UpdatePhotoTitleParams{
		Title:  params.Title,
		ID:     photoId,
		UserID: user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error updating photo: %v", err))
	}
}

func (h *Handler) FetchAllPhotos(w http.ResponseWriter, r *http.Request) {
	dbPhotos, err := h.Cfg.DB.GetAllPhotos(r.Context())

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error fetching photos, %v", err))
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DatabasePhotosToPhotos(dbPhotos))
}

func (h *Handler) DeletePhoto(w http.ResponseWriter, r *http.Request, user database.User) {
	photoIdStr := chi.URLParam(r, "photoId")
	photoId, err := uuid.Parse(photoIdStr)

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't parse photo id: %v", err))
		return
	}

	photo, err := h.Cfg.DB.FetchPhoto(r.Context(), photoId)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't fetch photo: %v", err))
		return
	}

	if photo.UserID != user.ID {
		utils.RespondWithError(w, http.StatusBadRequest, "You don't have permission to delete this photo")
		return
	}

	err = h.Cfg.DB.DeletePhoto(r.Context(), database.DeletePhotoParams{
		ID:     photoId,
		UserID: user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error deleting photo: %v", err))
	}
}
