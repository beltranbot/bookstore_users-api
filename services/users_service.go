package services

import (
	"net/http"

	"github.com/beltranbot/bookstore_users-api/domain/users"
	"github.com/beltranbot/bookstore_users-api/utils/errors"
)

// CreateUser func
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
