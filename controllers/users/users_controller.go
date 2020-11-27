package users

import (
	"net/http"
	"strconv"

	"github.com/beltranbot/bookstore_users-api/services"
	"github.com/beltranbot/bookstore_users-api/utils/errors"

	"github.com/beltranbot/bookstore_users-api/domain/users"

	"github.com/gin-gonic/gin"
)

// CreateUser func
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	result, createErr := services.CreateUser(user)
	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser func
func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

// SearchUser func
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
