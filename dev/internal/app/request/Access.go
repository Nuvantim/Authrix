package request

import repo "api/internal/app/repository"

type Permission struct {
	Name string `validate:"required" json:"name"`
}

type Role struct {
	Name         string  `json:"name"`
	PermissionID []int32 `json:"permission_id"`
}

type GetRole struct {
	ID         int32                       `json:"id"`
	Name       string                      `json:"name"`
	Permission []repo.GetPermissionRoleRow `json:"permission"`
}

type GetClient struct {
	ID    int32                   `json:"id"`
	Name  string                  `json:"name"`
	Email string                  `json:"email"`
	Role  []repo.GetRoleClientRow `json:"permission"`
}
