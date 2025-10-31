package service

import (
	db "api/database"
	repo "api/internal/app/repository"
	req "api/internal/app/request"

	ctx "context"
	"errors"
)

func GetRole(id int32) (req.GetRole, error) {
	role, err := db.Queries.GetRole(ctx.Background(), id)
	if err != nil {
		return req.GetRole{}, db.Fatal(err)
	}
	permission, err := db.Queries.GetPermissionRole(ctx.Background(), id)
	if err != nil {
		return req.GetRole{}, db.Fatal(err)
	}
	var data = req.GetRole{
		ID:         role.ID,
		Name:       role.Name,
		Permission: permission,
	}
	return data, nil
}

func ListRole() ([]repo.ListRoleRow, error) {
	role, err := db.Queries.ListRole(ctx.Background())
	if err != nil {
		return []repo.ListRoleRow{}, errors.New("role is empty")
	}
	return role, nil
}

func CreateRole(data req.Role) ([]repo.ListRoleRow, error) {
	// Create Role
	role_id, err := db.Queries.CreateRole(ctx.Background(), data.Name)
	if err != nil {
		return []repo.ListRoleRow{}, db.Fatal(err)
	}

	// check length data permission_id
	// VerfyPermission
	permission_id, err := db.Queries.VerifyPermission(ctx.Background(), data.PermissionID)
	if err != nil {
		return []repo.ListRoleRow{}, errors.New("permission not found")
	}
	// check length data permission
	var check int = len(permission_id)

	if check != 0 {
		// Params data Role_Permission
		role_permission := repo.AddPermissionRoleParams{
			IDRole:  role_id,
			Column2: permission_id,
		}

		// Create Role_Permission
		if err := db.Queries.AddPermissionRole(ctx.Background(), role_permission); err != nil {
			return []repo.ListRoleRow{}, err
		}
	}

	role, err := ListRole()
	if err != nil {
		return []repo.ListRoleRow{}, db.Fatal(err)
	}
	return role, nil
}

func UpdateRole(data req.Role, id int32) (req.GetRole, error) {
	// Update Role
	var role_data = repo.UpdateRoleParams{
		ID:   id,
		Name: data.Name,
	}

	if err := db.Queries.UpdateRole(ctx.Background(), role_data); err != nil {
		return req.GetRole{}, db.Fatal(err)
	}

	// Verify Role
	role_id, err := db.Queries.VerifyRole(ctx.Background(), data.PermissionID)
	if err != nil {
		return req.GetRole{}, db.Fatal(err)
	}
	// Check lenght PermissionID
	var check int = len(role_id)

	if check > 0 {
		// UpdatePermissionRole
		var role_permission = repo.UpdatePermissionRoleParams{
			IDRole:  id,
			Column2: data.PermissionID,
		}

		if err := db.Queries.UpdatePermissionRole(ctx.Background(), role_permission); err != nil {
			return req.GetRole{}, db.Fatal(err)
		}
	} else {
		_ = db.Queries.DeletePermissionRole(ctx.Background(), id)
	}

	roles, err := GetRole(id)
	if err != nil {
		return req.GetRole{}, db.Fatal(err)
	}
	return roles, nil
}

func DeleteRole(id int32) (string, error) {
	if err := db.Queries.DeleteRole(ctx.Background(), id); err != nil {
		return "", db.Fatal(err)
	}
	return "permission deleted", nil
}
