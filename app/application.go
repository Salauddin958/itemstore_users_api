package app

import (
	"github.com/federicoleon/bookstore_utils-go/logger"
	gin "github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")
	router.Run(":8081")
}
