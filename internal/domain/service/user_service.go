package service

import (
	"api/internal/database"
	"api/internal/domain/models"
	"api/pkg/utils"
)

type ( // declare type models User & UserTemps
	User = models.User
)

func CheckEmail(email string) bool {
	var user User

	// Check User
	result := database.DB.Where("email = ?", email).Find(&user)
	if result.RowsAffected > 0 {
		return true
	}

	return false
}

func RegisterAccount(Name, Email, Password string) string {
	// hashing password
	hashPassword := utils.HashBycrypt(Password)

	// Create UserTemp
	user := User{
		Name:     Name,
		Email:    Email,
		Password: string(hashPassword),
	}
	// Simpan user ke database
	database.DB.Create(&user)

	return "Success Register, Please Check Your Email"
}

func FindAccount(id uint) User {
	var user User

	// Get data by ID
	database.DB.Take(&user, id)

	// Kembalikan data
	return user
}

func UpdateAccount(users User, user_id uint) User {
	// Declare variable
	var user User

	// Get user data by id
	database.DB.Take(&user, user_id)

	// update user
	user.Name = users.Name
	user.Email = users.Email
	if users.Password != "" {
		hash := utils.HashBycrypt(users.Password)
		user.Password = string(hash)
	}

	// Simpan perubahan user
	database.DB.Save(&user)

	return users
}

func DeleteAccount(user_id uint) error {
	var user User
	if err := database.DB.Take(&user, user_id).Error; err != nil {
		return err
	}
	database.DB.Delete(&user)
	return nil
}
