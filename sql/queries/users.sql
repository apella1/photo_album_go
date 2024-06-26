-- name: CreateUser :one
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
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: IsUniqueEmail :one
SELECT COUNT(*) AS is_unique
FROM users
WHERE email = $1;

-- name: IsUniqueUsername :one
SELECT COUNT(*) AS is_unique
FROM users
WHERE username = $1;


-- name: FetchAllUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;