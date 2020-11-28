package users

import (
	"fmt"

	"github.com/beltranbot/bookstore_users-api/datasources/mysql/usersdb"

	"github.com/beltranbot/bookstore_users-api/utils/dateutils"

	"github.com/beltranbot/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Get func
func (user *User) Get() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// Save func
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", current.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", current.ID))
	}

	user.DateCreated = dateutils.GetNowString()
	usersDB[user.ID] = user

	return nil
}
