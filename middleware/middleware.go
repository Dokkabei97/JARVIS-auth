package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"is-deploy-auth/application"
)

// JwtAuthMiddleware JWT 인증 미들웨어
// JWT 인증이 필요한 라우터에 사용
func JwtAuthMiddleware(jwtService application.JwtService) gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := getTokenFromHeader(context)
		if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		_, err = jwtService.Validate(token)
		if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		context.Next()
	}
}

// JwtAuthAdminMiddleware JWT 인증 미들웨어
// 관리자 권한이 있는지 확인
// 관리자 권한이 없으면 403 Forbidden
// 관리자 권한이 있으면 다음 미들웨어로 이동
func JwtAuthAdminMiddleware(jwtService application.JwtService) gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := getTokenFromHeader(context)
		if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		valid, isAdmin, err := jwtService.ValidateAdmin(token)
		if err != nil || !valid {
			context.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		if !isAdmin {
			context.AbortWithStatusJSON(403, gin.H{"message": "Forbidden: insufficient privileges"})
			return
		}

		context.Next()
	}
}

// getTokenFromHeader HTTP 헤더에서 토큰을 가져옵니다.
// 토큰이 없으면 에러를 반환합니다.
func getTokenFromHeader(context *gin.Context) (string, error) {
	token := context.GetHeader("Authorization")
	if token == "" {
		return "", errors.New("token not found in header")
	}
	return token, nil
}
