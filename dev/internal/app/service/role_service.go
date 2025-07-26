package service

import (
	db "api/database"
	repo "api/internal/app/repository"
	req "api/internal/app/request"

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
	// Create Role
	role_id, err := db.Queries.CreateRole(ctx.Background(), data.Name)
	if err != nil {
		return []repo.Role{}, err
	}

	// check length data permission_id
	var check int = len(data.PermissionID)

	if check != 0 {
		// VerfyPermission
		permission_id, err := db.Queries.VerifyPermission(ctx.Background(), data.PermissionID)
		if err != nil {
			return []repo.Role{}, errors.New("Permission not found")
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
	}

	role, err := ListRole()
	if err != nil {
		return []repo.Role{}, err
	}
	return role, nil
}

func UpdateRole(data req.Role, id int32) (repo.GetRoleRow, error) {
	// Update Role
	var role_data = repo.UpdateRoleParams{
		ID:   id,
		Name: data.Name,
	}

	if err := db.Queries.UpdateRole(ctx.Background(), role_data); err != nil {
		return repo.GetRoleRow{}, err
	}
	// Check lenght PermissionID
	var check int = len(data.PermissionID)

	if check > 0 {
		// UpdatePermissionRole
		var role_permission = repo.UpdatePermissionRoleParams{
			IDRole:  id,
			Column2: data.PermissionID,
		}

		if err := db.Queries.UpdatePermissionRole(ctx.Background(), role_permission); err != nil {
			return repo.GetRoleRow{}, err
		}
	} else {
		_ = db.Queries.DeletePermissionRole(ctx.Background(), id)
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
