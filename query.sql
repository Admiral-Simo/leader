-- Create a new user
-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, password;

-- Get a user by ID
-- name: GetUserByID :one
SELECT id, name, email, password
FROM users
WHERE id = $1;

-- Get a user by email
-- name: GetUserByEmail :one
SELECT id, name, email, password
FROM users
WHERE email = $1;

-- Update a user's name, email, and password by ID
-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, password = $4
WHERE id = $1
RETURNING id, name, email, password;

-- Delete a user by ID
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- List all users
-- name: ListUsers :many
SELECT id, name, email, password
FROM users
ORDER BY id;
