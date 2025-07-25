package service

import (
	db "api/database"
	req "api/internal/app_request"
	repo "api/internal/repository"

	ctx "context"
	"errors"
)

func GetRole(id int32) (repo.GetRoleRow, error) {
	role, err := db.Queries.GetRole(ctx.Background(), id)
	if err != nil {
		return repo.GetRoleRow{}, errors.New("Role Not Found")
	}
	return role, nil
}
func ListRole() ([]repo.Role, error) {
	role, err := db.Queries.ListRole(ctx.Background())
	if err != nil {
		return []repo.Role{}, errors.New("Role is empty")
	}
	return role, nil
}
func CreateRole(data req.Role) ([]repo.Role, error) {
	// VerfyPermission
	permission_id, err := db.Queries.VerifyPermission(ctx.Background(), data.PermissionID)
	if err != nil {
		return []repo.Role{}, errors.New("Permission not found")
	}

	// Create Role
	role_id, err := db.Queries.CreateRole(ctx.Background(), data.Name)
	if err != nil {
		return []repo.Role{}, err
	}

	// Params data Role_Permission
	role_permission := repo.AddPermissionRoleParams{
		IDRole:  role_id,
		Column2: permission_id,
	}

	// Create Role_Permission
	if err := db.Queries.AddPermissionRole(ctx.Background(), role_permission); err != nil {
		return []repo.Role{}, err
	}
	role, err := ListRole()
	if err != nil {
		return []repo.Role{}, err
	}
	return role, nil
}
func UpdateRole(data req.Role, id int32) (repo.GetRoleRow, error) {
	var role_data = repo.UpdateRoleParams{
		ID:   id,
		Name: data.Name,
	}

	if err := db.Queries.UpdateRole(ctx.Background(), role_data); err != nil {
		return repo.GetRoleRow{}, err
	}
	var role, err = GetRole(id)
	if err != nil {
		return repo.GetRoleRow{}, err
	}
	return role, nil
}
func DeleteRole(id int32) (string, error) {
	if err := db.Queries.DeleteRole(ctx.Background(), id); err != nil {
		return "", err
	}
	return "permission deleted", nil
}
