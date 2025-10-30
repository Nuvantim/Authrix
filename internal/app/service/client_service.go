package service

import (
	db "api/database"
	model "api/internal/app/repository"
	req "api/internal/app/request"
	"api/pkgs/guards"

	ctx "context"
	str "strings"
)

func ListClient() ([]model.ListClientRow, error) {
	data, err := db.Queries.ListClient(ctx.Background())
	if err != nil {
		return []model.ListClientRow{}, db.Fatal(err)
	}
	return data, nil
}

func GetClient(id int32) (req.GetClient, error) {
	client, err := db.Queries.GetClient(ctx.Background(), id)
	if err != nil {
		return req.GetClient{}, db.Fatal(err)
	}
	role, err := db.Queries.GetRoleClient(ctx.Background(), id)
	if err != nil {
		return req.GetClient{}, db.Fatal(err)
	}
	var data = req.GetClient{
		ID:    client.ID,
		Name:  client.Name,
		Email: client.Email,
		Role:  role,
	}
	return data, nil
}

func UpdateClient(Id int32, client req.UpdateClient) (req.GetClient, error) {
	var update_data = model.UpdateClientParams{
		ID:      Id,
		Name:    client.Name,
		Column3: client.Email,
	}

	if str.TrimSpace(client.Password) != "" {
		psw := guard.HashBycrypt(client.Password)
		update_data.Column4 = string(psw)
	}

	// Update client data
	if err := db.Queries.UpdateClient(ctx.Background(), update_data); err != nil {
		return req.GetClient{}, db.Fatal(err)
	}

	// verify role
	role, err := db.Queries.VerifyRole(ctx.Background(), client.Role)
	if err != nil {
		return req.GetClient{}, db.Fatal(err)
	}
	var check int = len(role)
	if check != 0 {
		// update client role
		var client_role = model.UpdateRoleClientParams{
			IDUser:  Id,
			Column2: role,
		}

		if err := db.Queries.UpdateRoleClient(ctx.Background(), client_role); err != nil {
			return req.GetClient{}, db.Fatal(err)
		}
	} else {
		_ = db.Queries.DeleteRoleClient(ctx.Background(), Id)
	}

	// Get Client data
	client_data, err := GetClient(Id)
	if err != nil {
		return req.GetClient{}, db.Fatal(err)
	}

	return client_data, nil

}

func DeleteClient(id int32) (string, error) {
	if err := db.Queries.DeleteClient(ctx.Background(), id); err != nil {
		return "", db.Fatal(err)
	}
	return "Client Deleted", nil
}
