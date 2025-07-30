package service

import (
	db "api/database"
	repo "api/internal/app/repository"
	req "api/internal/app/request"
	"api/pkgs/utils"

	ctx "context"
	str "strings"
)

func ListClient() ([]repo.ListClientRow, error) {
	data, err := db.Queries.ListClient(ctx.Background())
	if err != nil {
		return []repo.ListClientRow{}, err
	}
	return data, nil
}

func GetClient(id int32) (repo.GetClientRow, error) {
	data, err := db.Queries.GetClient(ctx.Background(), id)
	if err != nil {
		return repo.GetClientRow{}, err
	}
	return data, nil
}

func UpdateClient(Id int32, client req.UpdateClient) (repo.UserAccount, error) {
	var update_data = repo.UpdateClientParams{
		ID:    Id,
		Name:  client.Name,
		Email: client.Email,
	}

	if str.TrimSpace(client.Password) != "" {
		psw := utils.HashBycrypt(client.Password)
		update_data.Password = string(psw)
	}

	// Update client data
	data, err := db.Queries.UpdateClient(ctx.Background(), update_data)
	if err != nil {
		return repo.UserAccount{}, err
	}

	// verify role
	role, err := db.Queries.VerifyRole(ctx.Background(), client.Role)
	if err != nil {
		return repo.UserAccount{}, err
	}
	// update client role
	var client_role = repo.UpdateRoleClientParams{
		IDUser:  Id,
		Column2: role,
	}

	if err := db.Queries.UpdateRoleClient(ctx.Background(), client_role); err != nil {
		return repo.UserAccount{}, err
	}

	return data, nil

}

func DeleteClient(id int32) (string, error) {
	if err := db.Queries.DeleteClient(ctx.Background(), id); err != nil {
		return "", err
	}
	return "Client Deleted", nil
}
