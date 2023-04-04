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

	lb := router.Group("/api/v1/load-balance")
	lb.Use(jwtAuthMiddleware)
	{
		lb.PUT("/exclude", api.Exclude)
		lb.PUT("/restore", api.Restore)
	}

	dp := router.Group("/api/v1/deploy")
	dp.Use(jwtAuthMiddleware)
	{
		dp.PUT("/shell", api.Deploy)
	}

	set := router.Group("/api/v1/setting")
	set.Use(jwtAuthMiddleware)
	{
		set.PUT("", api.SyncSettingJson)
	}

	return router
}
