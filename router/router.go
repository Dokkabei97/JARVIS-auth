package router

import "github.com/gin-gonic/gin"

func SetRouter() *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {

		c.Next()
	})

	auth := router.Group("/api/v1/auth")
	{
		auth.GET("")
		auth.POST("")
	}

	authRouter := router.Group("/api/v1/auth-router")
	{
		authRouter.POST("")
	}

	return router
}
