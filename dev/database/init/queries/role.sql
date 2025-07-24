-- name: GetRole :one
SELECT id,name FROM role WHERE id = $1;

-- name: CreateRole :exec
INSERT INTO role (name) VALUES ($1);

-- name: ListRole :many
SELECT * FROM role;

-- name: UpdateRole :exec
UPDATE role SET name=$2 WHERE id = $1;

-- name: DeleteRole :exec
DELETE FROM role WHERE id = $1;