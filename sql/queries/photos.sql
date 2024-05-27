-- name: CreatePhoto :one
INSERT INTO
    photos (
        id,
        created_at,
        updated_at,
        title,
        body,
        album_id
    )
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: FetchAlbumPhotos :many
SELECT * FROM photos WHERE album_id = $1;

-- name: DeletePhoto :exec
DELETE FROM photos
WHERE id = $1
AND album_id = $2
AND user_id = $3;