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
SELECT DISTINCT id FROM permission WHERE id = ANY($1:: int[]);




-- SELECT
--     STRING_AGG(p.name, ', ') AS permissions_list
-- FROM
--     "public"."role" AS r
-- JOIN
--     "public".role_permission AS rp ON r.id = rp.id_role
-- JOIN
--     "public".permission AS p ON rp.id_permission = p.id
-- WHERE
--     r.id = $1
-- GROUP BY
--     r.id, r.name
-- ORDER BY
--     r.name;