package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/beltranbot/bookstore_users-api/config"
	// load the mysql driver that we are going to use to connect to the database
	_ "github.com/go-sql-driver/mysql"
)

var (
	// Client is the users_db connection
	Client *sql.DB

	username = config.Config.MySQLUsersDBUsername
	password = config.Config.MySQLUsersDBPassword
	host     = config.Config.MySQLUsersDBHost
	port     = config.Config.MySQLUsersDBPort
	schema   = config.Config.MySQLUsersDBSchema
)

func init() {
	datasourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, password, host, port, schema,
	)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database sucessfully configured")
}
