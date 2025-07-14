-- name: CreateUser :one
INSERT INTO user_account(name,email,password) VALUES ($1,$2,$3) RETURNING *;

-- name: FindEmail :one
SELECT * FROM user_account WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE user_account SET name=$2, email=$3, password=$4 WHERE id = $1 RETURNING *;