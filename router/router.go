package router

import "github.com/gin-gonic/gin"

func SetRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/api/v1/auth")
	{
		auth.GET("")
		auth.POST("")
	}

	return router
}
