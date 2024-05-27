-- name: CreateAlbum :one
INSERT INTO
    albums (
        id,
        created_at,
        updated_at,
        title,
        user_id
    )
VALUES($1, $2, $3, $4, $5)
RETURNING *;

-- name: FetchUserAlbums :many
SELECT * FROM albums WHERE user_id = $1;

-- name: DeleteAlbum :exec
DELETE FROM albums
WHERE id = $1
AND user_id = $2;