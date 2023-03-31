package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"is-deploy-auth/application"
	infrastructure "is-deploy-auth/infrastructure"
	"is-deploy-auth/middleware"
)

func SetRouter(db *gorm.DB) *gin.Engine {
	jwtRepository := infrastructure.NewJwtRepository(db)
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
		authRouter.GET("")
		authRouter.POST("")
	}

	return router
}
