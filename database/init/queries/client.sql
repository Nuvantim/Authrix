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

-- name: UpdateClient :exec
UPDATE user_account SET 
	name = $2, 
    email = COALESCE(NULLIF(TRIM($3::varchar),''), email), 
	password = COALESCE(NULLIF(TRIM($4::varchar),''), password)
WHERE id = $1;

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
INSERT INTO user_role (id_user, id_role)
SELECT $1, UNNEST($2::int[]);

-- name: AllRoleClient :many
SELECT
    r.id,
    r.name,
    COALESCE(
        (
            SELECT jsonb_agg(
                jsonb_build_object(
                    'id', p.id,
                    'name', p.name
                )
                ORDER BY p.name
            )
            FROM public.permission AS p
            LEFT JOIN public.role_permission AS rp ON p.id = rp.id_permission
            WHERE rp.id_role = r.id
        ),
        '[]'::jsonb
    ) AS permissions
FROM
    public.role AS r
WHERE
    r.id IN (SELECT id_role FROM user_role WHERE id_user = $1)
GROUP BY
    r.id, r.name
ORDER BY
    r.name;


-- name: DeleteRoleClient :exec
DELETE FROM user_role WHERE id_user = $1;
