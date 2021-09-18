package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	
)

func StartApplication() {
	router.Use(CORSMiddleware())
	mapUrls()
	router.Run()
}
