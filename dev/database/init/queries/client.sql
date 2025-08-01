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

-- name: GetRoleClient :many
SELECT id_role FROM user_role WHERE id_user = $1;

-- name: DeleteClient :exec
DELETE FROM user_account WHERE id = $1;

-- name: CreateRoleClient :exec
INSERT INTO user_role (id_user, id_role) SELECT $1 AS user_id_params,
unnested_role_id FROM UNNEST($2::int[]) AS unnested_role_id;

-- name: UpdateRoleClient :exec
DELETE FROM user_role
WHERE id_user = $1
AND id_role NOT IN (SELECT unnested_role_id FROM UNNEST($2::int[]) AS unnested_role_id);