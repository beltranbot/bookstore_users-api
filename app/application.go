package app

import (
	"fmt"

	"github.com/beltranbot/bookstore_users-api/config"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication func
func StartApplication() {
	mapUrls()
	router.Run(fmt.Sprintf(":%s", config.Config.AppPort))
}
