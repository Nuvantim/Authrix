-- name: CreateUser :one
INSERT INTO user_account(name,email,password) 
VALUES ($1,$2,$3) 
RETURNING id;

-- name: CreateProfile :exec
INSERT INTO user_profile (user_id) 
VALUES ($1);

-- name: GetProfile :one
SELECT sqlc.embed(user_account), sqlc.embed(user_profile)
FROM "public".user_account
INNER JOIN "public".user_profile ON (user_account.id = user_profile.user_id) WHERE user_account.id = $1 LIMIT 1;

-- name: FindEmail :one
SELECT id, email,password FROM user_account WHERE email = $1;

-- name: UpdateAccount :exec
WITH updated_account AS (
    UPDATE user_account SET name = $2 WHERE id= $1
)
UPDATE user_profile SET age = $3, phone = $4, district = $5, city = $6, country = $7
WHERE user_id = $1;

-- name: UpdatePassword :exec
UPDATE user_account SET password=$2 WHERE id=$1;

-- name: ResetPassword :exec
UPDATE user_account SET password=$2 WHERE email=$1;

-- name: DeleteAccount :exec
DELETE FROM user_account WHERE id = $1;

