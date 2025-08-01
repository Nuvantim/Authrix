-- name: GetRole :one
SELECT id,name FROM role WHERE id = $1;

-- name: CreateRole :one
INSERT INTO role (name) VALUES ($1) RETURNING id;

-- name: ListRole :many
SELECT
    r.id,
    r.name,
    ARRAY_AGG(p.name ORDER BY p.name) AS permissions
FROM
    public.role AS r
JOIN
    public.role_permission AS rp ON r.id = rp.id_role
JOIN
    public.permission AS p ON rp.id_permission = p.id
GROUP BY
    r.id, r.name
ORDER BY
    r.name;

-- name: UpdateRole :exec
UPDATE role SET name=$2 WHERE id = $1;

-- name: VerifyRole :many
SELECT DISTINCT id FROM role WHERE id = ANY($1:: int[]);

-- name: AddPermissionRole :exec
INSERT INTO role_permission (id_role, id_permission) SELECT $1 AS role_id_params,
unnested_permission_id FROM UNNEST($2::int[]) AS unnested_permission_id;

-- name: UpdatePermissionRole :exec
WITH delete_permission AS (
  DELETE FROM role_permission
  WHERE id_role = $1
)
INSERT INTO role_permission (id_role, id_permission)
SELECT $1,UNNEST($2::int[]);

-- name: GetPermissionRole :one
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

-- name: DeleteRole :exec
DELETE FROM role WHERE id = $1;
