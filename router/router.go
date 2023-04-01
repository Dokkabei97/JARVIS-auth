package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"is-deploy-auth/application"
	"is-deploy-auth/infrastructure"
	"is-deploy-auth/middleware"
)

func SetRouter(db *gorm.DB) *gin.Engine {
	jwtRepository := infrastructure.NewJwtRepository(db)
	jwtService := application.NewJwtTokenService(jwtRepository)
	jwtAuthMiddleware := middleware.JwtAuthMiddleware(jwtService)
	jwtAuthAdminMiddleware := middleware.JwtAuthAdminMiddleware(jwtService)

	router := gin.Default()

	auth := router.Group("/api/v1/auth")
	{
		auth.GET("")
		auth.POST("")
	}

	authRouter := router.Group("/api/v1/auth-router")
	authRouter.Use(jwtAuthMiddleware)
	{
		authRouter.Group("/load-balance")
		{
			authRouter.GET("")
			authRouter.PUT("/exclude")
			authRouter.PUT("/restore")
		}

		authRouter.Group("/deploy")
		{
			authRouter.PUT("/shell")
		}

		authRouter.Group("/setting")
		{
			authRouter.PUT("")
		}

		authRouter.Group("/update")
		{
			authRouter.PUT("/:version")
			authRouter.GET("/version")
		}

		authRouter.Group("/logs")
		{
			authRouter.GET("")
		}

		authRouter.Group("/health-check")
		{
			authRouter.GET("/agent")
			authRouter.GET("/tomcat")
		}
	}

	return router
}
