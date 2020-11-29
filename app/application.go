package app

import (
	"fmt"

	"github.com/beltranbot/bookstore_users-api/config"
	"github.com/beltranbot/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication func
func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")
	router.Run(fmt.Sprintf(":%s", config.Config.AppPort))
}
