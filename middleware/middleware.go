package middleware

import (
	"github.com/gin-gonic/gin"
	"is-deploy-auth/application"
)

// JwtAuthMiddleware JWT 인증 미들웨어
func JwtAuthMiddleware(jwtService application.JwtService) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")

		if token == "" {
			context.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		_, err := jwtService.Validate(token)
		if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		context.Next()
	}
}
