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

-- name: AddPermissionRole :exec
INSERT INTO role_permission (id_role, id_permission) SELECT $1 AS role_id_params,
unnested_permission_id FROM UNNEST($2::int[]) AS unnested_permission_id;

-- name: UpdatePermissionRole :exec
WITH delete_permission AS (DELETE FROM role_permission
WHERE id_role = $1 
)
INSERT INTO role_permission (id_role, id_permission) SELECT $1 AS role_id_params, 
unnested_permission_id FROM UNNEST($2::int[]) AS unnested_permission_id 
ON CONFLICT (id_role, id_permission) DO NOTHING;

-- name: GetPermissionRole :many
SELECT
    CASE
        WHEN ROW_NUMBER() OVER (PARTITION BY r.id ORDER BY p.name) = 1 THEN r.name
        ELSE NULL
    END AS role_name,
    p.name AS permission_name,
    p.id AS permission_id
FROM
    "public"."role" AS r
JOIN
    "public".role_permission AS rp ON r.id = rp.id_role
JOIN
    "public".permission AS p ON rp.id_permission = p.id
WHERE
    r.id = $1
ORDER BY
    r.name, p.name;

-- name: ListPermissionRole :many
SELECT
    CASE
        WHEN ROW_NUMBER() OVER (PARTITION BY r.id ORDER BY p.name) = 1 THEN r.name
        ELSE NULL
    END AS role_name,
    p.name AS permission_name,
    p.id AS permission_id
FROM
    "public"."role" AS r
JOIN
    "public".role_permission AS rp ON r.id = rp.id_role
JOIN
    "public".permission AS p ON rp.id_permission = p.id
WHERE
    r.id IN ($1::int[])
ORDER BY
    r.name, p.name;

-- name: DeletePermissionRole :exec
DELETE FROM role_permission WHERE id_role = $1;


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