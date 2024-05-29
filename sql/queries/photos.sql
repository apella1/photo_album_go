-- name: CreatePhoto :one
INSERT INTO
    photos (
        id,
        created_at,
        updated_at,
        title,
        body,
        album_id,
        user_id,
        img_url
    )
VALUES($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: FetchAlbumPhotos :many
SELECT * FROM photos WHERE album_id = $1;

-- name: DeletePhoto :exec
DELETE FROM photos
WHERE id = $1
AND album_id = $2
AND user_id = $3;

-- name: FetchPhoto :one
SELECT * FROM photos WHERE id = $1;

-- name: UpdatePhotoTitle :exec
UPDATE photos
SET title = $1
WHERE id = $2 AND user_id = $3;

-- name: GetAllPhotos :many
SELECT * FROM photos;