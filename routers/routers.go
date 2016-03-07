package routers

import (
	"github.com/gin-gonic/gin"
	"helloGin/controllers"
)

func Init() *gin.Engine{
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", controllers.Index)
	router.GET("/hello", controllers.Hello)

	return router
}
