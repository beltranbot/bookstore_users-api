package users

import (
	"fmt"
	"net/http"

	"github.com/beltranbot/bookstore_users-api/services"
	"github.com/beltranbot/bookstore_users-api/utils/errors"

	"github.com/beltranbot/bookstore_users-api/domain/users"

	"github.com/gin-gonic/gin"
)

// CreateUser func
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser func
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// SearchUser func
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
