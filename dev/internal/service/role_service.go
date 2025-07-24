package service

import (
	db "api/database"
	repo "api/internal/repository"
	req "api/internal/request"

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
	if err := db.Queries.CreateRole(ctx.Background(), data.Name); err != nil {
		return []repo.Role{}, err
	}
	var role, err = ListRole()
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
