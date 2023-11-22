-- name: CreateUser :one
INSERT INTO users(name, email, password)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetUserWithEmail :one
SELECT *
FROM users
WHERE email = $1;