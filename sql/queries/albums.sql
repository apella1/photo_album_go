-- name: CreateAlbum :one
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
RETURNING *;

-- name: FetchUserAlbums :many
SELECT * FROM albums WHERE user_id = $1;

-- name: DeleteAlbum :exec
DELETE FROM albums
WHERE id = $1
AND user_id = $2;

-- name: GetAlbumById :one
SELECT * FROM albums WHERE id = $1;

-- name: FetchAllAlbums :many
SELECT * FROM albums;