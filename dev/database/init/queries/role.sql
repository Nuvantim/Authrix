-- name: GetRole :one
SELECT id,name FROM role WHERE id = $1;

-- name: CreateRole :one
INSERT INTO role (name) VALUES ($1) RETURNING id;

-- name: ListRole :many
SELECT * FROM role;

-- name: UpdateRole :exec
UPDATE role SET name=$2 WHERE id = $1;

-- name: VerifyRole :many
SELECT DISTINCT id FROM role WHERE id = ANY($1:: int[]);

-- name: CreateRoleClient :exec
INSERT INTO user_role (id_user, id_role) SELECT $1 AS user_id_params,
unnested_role_id FROM UNNEST($2::int[]) AS unnested_role_id;

-- name: UpdateRoleClient :exec
DELETE FROM user_role
WHERE id_user = $1
AND id_role NOT IN (SELECT unnested_role_id FROM UNNEST($2::int[]) AS unnested_role_id);

-- name: GetRoleClient


-- name: DeleteRole :exec
DELETE FROM role WHERE id = $1;
