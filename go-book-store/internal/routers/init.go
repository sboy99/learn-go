package routers

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()
	RegisterBookRoutes(router)
	router.Run(`localhost:8080`)
}
