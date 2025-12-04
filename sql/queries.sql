-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: AddUser :exec
INSERT INTO users (name, email, password_hash)
VALUES ($1, $2, $3);
