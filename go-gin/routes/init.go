package routes

import (
	"learn/go-gin/controllers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	ginRouter := gin.Default()
	router = ginRouter
	router.GET("/albums", controllers.GetAlbums)
	router.Run("localhost:8080")
}
