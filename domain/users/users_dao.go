package users

import (
	"fmt"
	"strings"

	"github.com/beltranbot/bookstore_users-api/datasources/mysql/usersdb"
	"github.com/beltranbot/bookstore_users-api/logger"

	"github.com/beltranbot/bookstore_users-api/utils/mysqlutils"
	"github.com/beltranbot/bookstore_utils-go/resterrors"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser             = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser             = "DELETE FROM users WHERE id = ?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email = ? AND password = ? and status = ?;"
)

var (
	usersDB = make(map[int64]*User)
)

// Get func
func (user *User) Get() *resterrors.RestErr {
	statement, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
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
		logger.Error("error when trying to get user by id", scanErr)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}
	return nil
}

// Save func
func (user *User) Save() *resterrors.RestErr {
	statement, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare insert user statement", err)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
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
		logger.Error("error when trying to prepare save user", err)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last inserted id after creating new user", err)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}

	user.ID = userID

	return nil
}

// Update func
func (user *User) Update() *resterrors.RestErr {
	statement, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to get to prepare update user statement", err)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}
	defer statement.Close()

	_, err = statement.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("error when trying to execute user update statement", err)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}
	return nil
}

// Delete func
func (user *User) Delete() *resterrors.RestErr {
	statement, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete statement", err)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}
	defer statement.Close()

	if _, err = statement.Exec(user.ID); err != nil {
		logger.Error("error when executing delete user statement", err)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}
	return nil
}

// FindByStatus func
func (user *User) FindByStatus(status string) ([]User, *resterrors.RestErr) {
	statement, err := usersdb.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}
	defer statement.Close()

	rows, err := statement.Query(status)
	if err != nil {
		logger.Error("error when executing find user by status statement", err)
		return nil, resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when parsing find user by status results", err)
			return nil, resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, resterrors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

// FindByEmailAndPassword func
func (user *User) FindByEmailAndPassword() *resterrors.RestErr {
	statement, err := usersdb.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}
	defer statement.Close()

	result := statement.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status,
	); getErr != nil {
		if strings.Contains(getErr.Error(), mysqlutils.ErrorNoRows) {
			return resterrors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", getErr)
		return resterrors.NewInternalServerError("database error", resterrors.NewError("database error"))
	}
	return nil
}
