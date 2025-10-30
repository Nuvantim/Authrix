package request

import model "api/internal/app/repository"

type Permission struct {
	Name string `validate:"required" json:"name"`
}

type Role struct {
	Name         string  `validate:"required" json:"name"`
	PermissionID []int32 `validate:"omitempty,dive,gt=0" json:"permission_id"`
}

type GetRole struct {
	ID         int32                       `json:"id"`
	Name       string                      `json:"name"`
	Permission []model.GetPermissionRoleRow `json:"permission"`
}

type GetClient struct {
	ID    int32                   `json:"id"`
	Name  string                  `json:"name"`
	Email string                  `json:"email"`
	Role  []model.GetRoleClientRow `json:"permission"`
}
