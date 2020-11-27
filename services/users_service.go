package services

import (
	"github.com/beltranbot/bookstore_users-api/domain/users"
	"github.com/beltranbot/bookstore_users-api/utils/errors"
)

// GetUser func
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser func
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}
