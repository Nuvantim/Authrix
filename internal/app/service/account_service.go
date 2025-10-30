package service

import (
	db "api/database"
	model "api/internal/app/repository"
	req "api/internal/app/request"
	"api/pkgs/guards"

	ctx "context"
	"errors"
	str "strings"
)

func GetProfile(userID int32) (model.GetProfileRow, error) {
	data, err := db.Queries.GetProfile(ctx.Background(), userID)
	if err != nil {
		return model.GetProfileRow{}, errors.New("account not found")
	}
	data.UserAccount.ID = 0
	data.UserProfile.UserID = 0
	data.UserAccount.Password = ""
	return data, nil
}

func UpdateAccount(user req.UpdateAccount, userIDs int32) (model.GetProfileRow, error) {
	// Define update profile
	var updateAccount = model.UpdateAccountParams{
		UserID:   userIDs,
		Name:     user.Name,
		Age:      user.Age,
		Phone:    user.Phone,
		District: user.District,
		City:     user.City,
		Country:  user.Country,
	}

	// Create a buffered channel to receive any error from the goroutine
	errChan := make(chan error, 1)

	// Run user creation and OTP deletion in a separate goroutine
	go func() {
		// Update password is available
		if str.TrimSpace(user.Password) != "" {
			psw := guard.HashBycrypt(user.Password)
			passUpdate := model.UpdatePasswordParams{
				ID:       userIDs,
				Password: string(psw),
			}
			if err := db.Queries.UpdatePassword(ctx.Background(), passUpdate); err != nil {
				errChan <- err
				return
			}
		}
		// Update Profile User
		if err := db.Queries.UpdateAccount(ctx.Background(), updateAccount); err != nil {
			errChan <- err
			return
		}
		// Both operations succeeded
		errChan <- nil
	}()

	// Wait for the result from the goroutine
	if err := <-errChan; err != nil {
		return model.GetProfileRow{}, db.Fatal(err)
	}

	// Returning data
	usr, err := GetProfile(userIDs)
	if err != nil {
		return model.GetProfileRow{}, db.Fatal(err)
	}
	return usr, nil
}

func DeleteAccount(userID int32) (string, error) {
	if err := db.Queries.DeleteAccount(ctx.Background(), userID); err != nil {
		return "", db.Fatal(err)
	}
	return "Your account successfuly delete", nil
}
