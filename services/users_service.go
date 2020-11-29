package services

import (
	"github.com/beltranbot/bookstore_users-api/domain/users"
	"github.com/beltranbot/bookstore_users-api/utils/errors"
)

// Get func
func Get(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// Create func
func Create(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}

// Update func
func Update(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := Get(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

// Delete func
func Delete(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}
