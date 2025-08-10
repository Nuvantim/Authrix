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
		return repo.GetPermissionRow{}, errors.New("permission not found")
	}
	return permission, nil
}
func ListPermission() ([]repo.Permission, error) {
	permission, err := db.Queries.ListPermission(ctx.Background())
	if err != nil {
		return []repo.Permission{}, errors.New("permission is empty")
	}
	return permission, nil
}
func CreatePermission(data req.Permission) ([]repo.Permission, error) {
	if err := db.Queries.CreatePermission(ctx.Background(), data.Name); err != nil {
		return []repo.Permission{}, db.Fatal(err)
	}
	var permission, err = ListPermission()
	if err != nil {
		return []repo.Permission{}, db.Fatal(err)
	}
	return permission, nil
}
func UpdatePermission(data req.Permission, id int32) (repo.GetPermissionRow, error) {
	var permission_data = repo.UpdatePermissionParams{
		ID:   id,
		Name: data.Name,
	}

	if err := db.Queries.UpdatePermission(ctx.Background(), permission_data); err != nil {
		return repo.GetPermissionRow{}, db.Fatal(err)
	}
	var permission, err = GetPermission(id)
	if err != nil {
		return repo.GetPermissionRow{}, db.Fatal(err)
	}
	return permission, nil
}

func DeletePermission(id int32) (string, error) {
	if err := db.Queries.DeletePermission(ctx.Background(), id); err != nil {
		return "", db.Fatal(err)
	}
	return "permission deleted", nil
}
