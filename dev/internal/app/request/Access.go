package request

import repo "api/internal/app/repository"

type Permission struct {
	Name string "json:name"
}

type Role struct {
	Name         string  `json:"name"`
	PermissionID []int32 `json:"permission_id"`
}

type GetRole struct {
	Role       repo.GetRoleRow `json:"role"`
	Permission string          `json:"permission"`
}
