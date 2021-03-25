package database

import (
	"errors"

	"github.com/bnsngltn/go-fiber-boilerplate/model"
)

// Creates a User on the database
func CreateUser(user *model.User) error {
	// Check if contents of the struct are valid
	if err := user.ValidateContents(); err != nil {
		return err
	}

	// Add the user to the database
	if err := DB.Create(&user).Error; err != nil {
		return errors.New("Username might not be unique")
	}
	return nil
}

// GetUserByUsername
// Retrieve a user from the database through the given username
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	var err error

	if err = DB.Where(&model.User{Username: username}).Preload("Roles").Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
