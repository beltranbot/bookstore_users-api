package users

import (
	"fmt"
	"strings"

	"github.com/beltranbot/bookstore_users-api/datasources/mysql/usersdb"

	"github.com/beltranbot/bookstore_users-api/utils/dateutils"

	"github.com/beltranbot/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_unique"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
)

var (
	usersDB = make(map[int64]*User)
)

// Get func
func (user *User) Get() *errors.RestErr {
	statement, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()
	result := statement.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d not found", user.ID),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to get user %d: %s", user.ID, err.Error()),
		)
	}
	return nil
}

// Save func
func (user *User) Save() *errors.RestErr {
	statement, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	user.DateCreated = dateutils.GetNowString()

	insertResult, err := statement.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
	)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}

	user.ID = userID

	return nil
}
