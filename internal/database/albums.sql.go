// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: albums.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createAlbum = `-- name: CreateAlbum :one
INSERT INTO
    albums (
        id,
        created_at,
        updated_at,
        title,
        photos,
        user_id
    )
VALUES($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, title, photos, user_id
`

type CreateAlbumParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Photos    [][]byte
	UserID    uuid.UUID
}

func (q *Queries) CreateAlbum(ctx context.Context, arg CreateAlbumParams) (Album, error) {
	row := q.db.QueryRowContext(ctx, createAlbum,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		pq.Array(arg.Photos),
		arg.UserID,
	)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		pq.Array(&i.Photos),
		&i.UserID,
	)
	return i, err
}

const deleteAlbum = `-- name: DeleteAlbum :exec
DELETE FROM albums
WHERE id = $1
AND user_id = $2
`

type DeleteAlbumParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteAlbum(ctx context.Context, arg DeleteAlbumParams) error {
	_, err := q.db.ExecContext(ctx, deleteAlbum, arg.ID, arg.UserID)
	return err
}

const fetchUserAlbums = `-- name: FetchUserAlbums :many
SELECT id, created_at, updated_at, title, photos, user_id FROM albums WHERE user_id = $1
`

func (q *Queries) FetchUserAlbums(ctx context.Context, userID uuid.UUID) ([]Album, error) {
	rows, err := q.db.QueryContext(ctx, fetchUserAlbums, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Album
	for rows.Next() {
		var i Album
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			pq.Array(&i.Photos),
			&i.UserID,
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

const getAlbumById = `-- name: GetAlbumById :one
SELECT id, created_at, updated_at, title, photos, user_id FROM albums WHERE id = $1
`

func (q *Queries) GetAlbumById(ctx context.Context, id uuid.UUID) (Album, error) {
	row := q.db.QueryRowContext(ctx, getAlbumById, id)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		pq.Array(&i.Photos),
		&i.UserID,
	)
	return i, err
}
