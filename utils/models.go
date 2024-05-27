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
