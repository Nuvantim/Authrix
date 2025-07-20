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

func GetProfile(userID int32) (repo.GetUserRow, error) {
	data, err := db.Queries.GetUser(ctx.Background(), userID)
	if err != nil {
		return repo.GetUserRow{}, errors.New("Account not found !")
	}
	return data, nil
}

func UpdateProfile(user req.UpdateProfile, userIDs int32) (repo.GetUserRow, error) {
	// Define update profile
	var UpdateProfiles = repo.UpdateProfileParams{
		UserID:   userIDs,
		Name:     user.Name,
		Age:      user.Age,
		Phone:    user.Phone,
		District: user.District,
		City:     user.City,
		Country:  user.Country,
	}

	// Update password is available
	if user.Password != "" {
		psw := utils.HashBycrypt(user.Password)
		var passUpdate = repo.UpdatePasswordParams{
			ID:       userIDs,
			Password: string(psw),
		}
	}

	// Create a buffered channel to receive any error from the goroutine
	errChan := make(chan error, 1)

	// Run user creation and OTP deletion in a separate goroutine
	go func() {
		if user.Password != "" {
			if err := db.Queries.UpdatePassword(ctx.Background(), passUpdate); err != nil {
				errChan <- err
				return
			}
		}
		if err := db.Queries.UpdateProfile(ctx.Background(), UpdateProfiles); err != nil {
			errChan <- err
			return
		}
		// Both operations succeeded
		errChan <- nil
	}()

	// Wait for the result from the goroutine
	if err := <-errChan; err != nil {
		return repo.GetUserRow{}, err
	}

	// Returning data
	usr, err := GetProfile(userIDs)
	if err != nil {
		return repo.GetUserRow{}, err
	}
	return usr, nil
}

func DeleteAccount(userID int32) (string, err) {
	if err := db.Queries.DeleteAccount(ctx.Background(), userID); err != nil {
		return "", err
	}
	return "Your account successfuly delete", nil
}
