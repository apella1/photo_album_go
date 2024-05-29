package utils

import (
	"album/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
	}
}

func DatabaseUsersToUsers(dbUsers []database.User) []User {
	users := []User{}

	for _, dbUser := range dbUsers {
		users = append(users, DatabaseUserToUser(dbUser))
	}
	return users
}

type Album struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Photos    [][]byte  `json:"photos"`
	UserID    uuid.UUID `json:"user_id"`
}

func DatabaseAlbumToAlbum(dbAlbum database.Album) Album {
	return Album{
		ID:        dbAlbum.ID,
		CreatedAt: dbAlbum.CreatedAt,
		UpdatedAt: dbAlbum.UpdatedAt,
		Title:     dbAlbum.Title,
		UserID:    dbAlbum.UserID,
	}
}

func DatabaseAlbumsToAlbums(dbAlbums []database.Album) []Album {
	albums := []Album{}

	for _, dbAlbum := range dbAlbums {
		albums = append(albums, DatabaseAlbumToAlbum(dbAlbum))
	}
	return albums
}

type Photo struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AlbumID   uuid.UUID `json:"album_id"`
	Title     string    `json:"title"`
	UserId    uuid.UUID `json:"user_id"`
	Body      []byte    `json:"body"`
	ImageURL  string    `json:"image_url"`
}

func DatabasePhotoToPhoto(dbPhoto database.Photo) Photo {
	return Photo{
		ID:        dbPhoto.ID,
		CreatedAt: dbPhoto.CreatedAt,
		UpdatedAt: dbPhoto.UpdatedAt,
		AlbumID:   dbPhoto.AlbumID,
		Title:     dbPhoto.Title,
		UserId:    dbPhoto.UserID,
		Body:      dbPhoto.Body,
		ImageURL:  dbPhoto.ImgUrl.String,
	}
}

func DatabasePhotosToPhotos(dbPhotos []database.Photo) []Photo {
	photos := []Photo{}

	for _, dbPhoto := range dbPhotos {
		photos = append(photos, DatabasePhotoToPhoto(dbPhoto))
	}
	return photos
}
