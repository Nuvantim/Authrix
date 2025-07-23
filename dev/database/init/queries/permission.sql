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