-- name: GetPermission :one
SELECT id, name FROM permission WHERE id = $1;

-- name: ListPermission :many
SELECT * FROM permission;

-- name: CreatePermission :exec
INSERT INTO permission (name) VALUES ($1);

-- name: UpdatePermission :exec
UPDATE permission SET name=$2 WHERE id=$1;

-- name: DeletePermission :exec
DELETE FROM permission WHERE id=$1;

-- name: VerifyPermission :many
SELECT DISTINCT id FROM permission WHERE id = ANY(@ids:: int[]);

-- name: AddPermissionRole :exec
INSERT INTO role_permission (id_role, id_permission) SELECT $1 AS role_id_params,
unnested_permission_id FROM UNNEST($2::int[]) AS unnested_permission_id;
