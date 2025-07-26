package service

import (
	db "api/database"
	repo "api/internal/app/repository"
	req "api/internal/app/request"

	ctx "context"
	"errors"
)

func GetPermission(id int32) (repo.GetPermissionRow, error) {
	permission, err := db.Queries.GetPermission(ctx.Background(), id)
	if err != nil {
		return repo.GetPermissionRow{}, errors.New("Permission Not Found")
	}
	return permission, nil
}
func ListPermission() ([]repo.Permission, error) {
	permission, err := db.Queries.ListPermission(ctx.Background())
	if err != nil {
		return []repo.Permission{}, errors.New("Permission is empty")
	}
	return permission, nil
}
func CreatePermission(data req.Permission) ([]repo.Permission, error) {
	if err := db.Queries.CreatePermission(ctx.Background(), data.Name); err != nil {
		return []repo.Permission{}, err
	}
	var permission, err = ListPermission()
	if err != nil {
		return []repo.Permission{}, err
	}
	return permission, nil
}
func UpdatePermission(data req.Permission, id int32) (repo.GetPermissionRow, error) {
	var permission_data = repo.UpdatePermissionParams{
		ID:   id,
		Name: data.Name,
	}

	if err := db.Queries.UpdatePermission(ctx.Background(), permission_data); err != nil {
		return repo.GetPermissionRow{}, err
	}
	var permission, err = GetPermission(id)
	if err != nil {
		return repo.GetPermissionRow{}, err
	}
	return permission, nil
}

func DeletePermission(id int32) (string, error) {
	if err := db.Queries.DeletePermission(ctx.Background(), id); err != nil {
		return "", err
	}
	return "permission deleted", nil
}
