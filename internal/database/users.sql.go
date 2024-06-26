// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (
        id,
        created_at,
        updated_at,
        first_name,
        last_name,
        username,
        email,
        password
    )
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING id, created_at, updated_at, first_name, last_name, username, email, password
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const fetchAllUsers = `-- name: FetchAllUsers :many
SELECT id, created_at, updated_at, first_name, last_name, username, email, password FROM users
`

func (q *Queries) FetchAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, fetchAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.Email,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, first_name, last_name, username, email, password FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, created_at, updated_at, first_name, last_name, username, email, password FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const isUniqueEmail = `-- name: IsUniqueEmail :one
SELECT COUNT(*) AS is_unique
FROM users
WHERE email = $1
`

func (q *Queries) IsUniqueEmail(ctx context.Context, email string) (int64, error) {
	row := q.db.QueryRowContext(ctx, isUniqueEmail, email)
	var is_unique int64
	err := row.Scan(&is_unique)
	return is_unique, err
}

const isUniqueUsername = `-- name: IsUniqueUsername :one
SELECT COUNT(*) AS is_unique
FROM users
WHERE username = $1
`

func (q *Queries) IsUniqueUsername(ctx context.Context, username string) (int64, error) {
	row := q.db.QueryRowContext(ctx, isUniqueUsername, username)
	var is_unique int64
	err := row.Scan(&is_unique)
	return is_unique, err
}
