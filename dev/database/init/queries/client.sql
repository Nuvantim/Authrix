-- name: ListClient :many
SELECT
    u.id,
    u.name,
    u.email,
    ARRAY_AGG(r.name ORDER BY r.name) FILTER (WHERE r.name IS NOT NULL) AS role
FROM
    public.user_account AS u
LEFT JOIN
    public.user_role AS ur ON u.id = ur.id_user
LEFT JOIN
    public.role AS r ON ur.id_role = r.id
GROUP BY
    u.id, u.name, u.email
ORDER BY
    u.name;

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
SELECT id,name FROM role WHERE id IN (SELECT id_role FROM user_role WHERE id_user = $1);

-- name: DeleteClient :exec
DELETE FROM user_account WHERE id = $1;

-- name: CreateRoleClient :exec
INSERT INTO user_role (id_user, id_role) SELECT $1 AS user_id_params,
unnested_role_id FROM UNNEST($2::int[]) AS unnested_role_id;

-- name: UpdateRoleClient :exec
WITH delete_role AS (
  DELETE FROM user_role
  WHERE id_user = $1 
)
INSERT INTO user_role (id_user, id_rol)
SELECT $1, UNNEST($2::int[]);

--name: AnyRoleClient :many
SELECT
    r.id,
    r.name,
    COALESCE(
        jsonb_agg(
            DISTINCT jsonb_build_object(
                'id', p.id,
                'name', p.name
            )
            ORDER BY p.name
        ) FILTER (WHERE p.id IS NOT NULL),
        '[]'
    ) AS permissions
FROM
    public.role AS r
LEFT JOIN
    public.role_permission AS rp ON r.id = rp.id_role
LEFT JOIN
    public.permission AS p ON rp.id_permission = p.id
GROUP BY
    r.id, r.name
WHERE
    id IN (SELECT id_role FROM user_role WHERE id_user = $1)
ORDER BY
    r.name;

-- name: DeleteRoleClient :exec
DELETE FROM user_role WHERE id_user = $1;
