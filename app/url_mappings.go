package app

import (
	"github.com/beltranbot/bookstore_users-api/controllers/ping"
	"github.com/beltranbot/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	// router.GET("/users/search", controllers.SearchUser)
}
