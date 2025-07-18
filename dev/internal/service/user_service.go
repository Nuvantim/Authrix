package service

import (
	db "api/database"
	repo "api/internal/repository"
	req "api/internal/request"
	"api/pkgs/utils"

	ctx "context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GetProfile(userID int32) (repo.GetUserRow,error){
	data, err := db.Queries.GetUser(ctx.Background(), userID); 
	if err != nil{
		return repo.GetUserRow{},errors.New("Account not found !")
	}
	return data, nil	
}

func UpdateAccount()(){}