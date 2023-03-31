package router

import (
	"github.com/gin-gonic/gin"
	"is-deploy-auth/application"
	"is-deploy-auth/infrastructure"
	"is-deploy-auth/middleware"
)

func SetRouter() *gin.Engine {
	jwtRepository := infrastructure.NewJwtRepository()
	jwtService := application.NewJwtTokenService(jwtRepository)
	jwtAuthMiddleware := middleware.JwtAuthMiddleware(jwtService)

	router := gin.Default()

	auth := router.Group("/api/v1/auth")
	{
		auth.GET("")
		auth.POST("")
	}

	authRouter := router.Group("/api/v1/auth-router")
	authRouter.Use(jwtAuthMiddleware)
	{
		authRouter.POST("")
	}

	return router
}
