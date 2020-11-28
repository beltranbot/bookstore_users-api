package config

import "github.com/spf13/viper"

type config struct {
	MySQLUsersDBUsername string
	MySQLUsersDBPassword string
	MySQLUsersDBHost     string
	MySQLUsersDBSchema   string
	MySQLUsersDBPort     string
	AppPort              string
}

// Config app configuration
var Config *config

func init() {
	loadViper()
}

func loadViper() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Config = &config{
		MySQLUsersDBUsername: viper.GetString("MYSQL_USERSDB_USERNAME"),
		MySQLUsersDBPassword: viper.GetString("MYSQL_USERSDB_PASSWORD"),
		MySQLUsersDBHost:     viper.GetString("MYSQL_USERSDB_HOST"),
		MySQLUsersDBPort:     viper.GetString("MYSQL_USERSDB_PORT"),
		MySQLUsersDBSchema:   viper.GetString("MYSQL_USERSDB_SCHEMA"),
		AppPort:              viper.GetString("APP_PORT"),
	}
}
