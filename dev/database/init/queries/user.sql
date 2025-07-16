-- name: CreateUser :exec
INSERT INTO user_account(name,email,password) VALUES ($1,$2,$3);

-- name: FindEmail :one
SELECT id,email FROM user_account WHERE email = $1 LIMIT 1;

-- name: LoginUser :one
SELECT id,email FROM user_account WHERE email = $1 AND password = $2 LIMIT 1;

-- name: UpdateUser :one
UPDATE user_account SET name=$2, email=$3, password=$4 WHERE id = $1 RETURNING *;

-- name: UpdatePassword :exec
UPDATE user_account SET password=$2 WHERE email=$1;