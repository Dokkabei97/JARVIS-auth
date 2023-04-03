package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"is-deploy-auth/api"
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

	authRouterAdmin := router.Group("/api/v1/auth-router-admin")
	authRouterAdmin.Use(jwtAuthAdminMiddleware)
	{

	}

	authRouter := router.Group("/api/v1/auth-router")
	authRouter.Use(jwtAuthMiddleware)
	{
		authRouter.Group("/load-balance")
		{
			authRouter.PUT("/exclude", api.Exclude)
			authRouter.PUT("/restore", api.Restore)
		}

		authRouter.Group("/deploy")
		{
			authRouter.PUT("/shell", api.Deploy)
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
	}

	return router
}
