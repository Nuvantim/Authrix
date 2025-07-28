-- name: ListClient :many
SELECT id,name,email FROM user_account;

-- name: GetClient :one
SELECT id,name,email FROM user_account WHERE id = $1;

-- name: UpdateClient :one
UPDATE user_account SET 
	name = $2, 
	email = $3, 
	password = CASE
		WHEN $4 IS NULL OR TRIM($4) = '' THEN password
		ELSE $4
	END
WHERE id = $1 RETURNING *;

-- name: DeleteClient :exec
DELETE FROM user_account WHERE id = $1;