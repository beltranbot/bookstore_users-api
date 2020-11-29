package users

import (
	"fmt"

	"github.com/beltranbot/bookstore_users-api/datasources/mysql/usersdb"

	"github.com/beltranbot/bookstore_users-api/utils/mysqlutils"

	"github.com/beltranbot/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser       = "DELETE FROM users WHERE id = ?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
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
	scanErr := result.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status,
	)
	if scanErr != nil {
		return mysqlutils.ParseError(scanErr)
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

	insertResult, saveErr := statement.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password,
	)
	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlutils.ParseError(err)
	}

	user.ID = userID

	return nil
}

// Update func
func (user *User) Update() *errors.RestErr {
	statement, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	_, err = statement.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}

// Delete func
func (user *User) Delete() *errors.RestErr {
	statement, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	if _, err = statement.Exec(user.ID); err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}

// FindByStatus func
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	statement, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	rows, err := statement.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysqlutils.ParseError(err)
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
