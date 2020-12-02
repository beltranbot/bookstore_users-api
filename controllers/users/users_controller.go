package users

import (
	"net/http"
	"strconv"

	"github.com/beltranbot/bookstore_users-api/domain/users"
	"github.com/beltranbot/bookstore_users-api/services"
	"github.com/beltranbot/bookstore_users-api/utils/errors"

	"github.com/beltranbot/bookstore_oauth-go/oauth"

	"github.com/gin-gonic/gin"
)

// Create func
func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	result, createErr := services.UsersService.Create(user)
	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Get func
func Get(c *gin.Context) {
	if getErr := oauth.AuthenticateRequest(c.Request); getErr != nil { // using the custom oauth library
		c.JSON(getErr.Status, getErr)
		return
	}

	if callerID := oauth.GetCallerID(c.Request); callerID == 0 {
		err := errors.RestErr{
			Status:  http.StatusUnauthorized,
			Message: "resource not avaiable",
		}
		c.JSON(err.Status, err)
		return
	}

	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := services.UsersService.Get(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if oauth.GetCallerID(c.Request) == user.ID {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}

	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}

// Update func
func Update(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UsersService.Update(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Delete func
func Delete(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	if deleteErr := services.UsersService.Delete(userID); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search func
func Search(c *gin.Context) {
	status := c.Query("status")
	users, findErr := services.UsersService.Search(status)
	if findErr != nil {
		c.JSON(findErr.Status, findErr)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}

// Login func
func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
