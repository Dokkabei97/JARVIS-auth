package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-auth/application"
	"is-deploy-auth/domain"
	"net/http"
)

func MakeToken(context *gin.Context, jwtService application.JwtService, isReissue bool) {
	var body domain.UserInfo

	if err := context.BindJSON(&body); err != nil {
		handleError(context, 400, err.Error())
		return
	}

	var token *domain.Token
	var err error

	if isReissue {
		accessToken := context.GetHeader("Authorization")
		if accessToken == "" {
			handleError(context, http.StatusOK, "accessToken not found in header")
			return
		}

		refreshToken, err := context.Cookie("refreshToken")
		if err != nil {
			handleError(context, http.StatusOK, "refreshToken not found in cookie")
			return
		}

		token, err = jwtService.ReissueToken(accessToken, refreshToken, body)
	} else {
		token, err = jwtService.IssueToken(body)
	}

	if err != nil {
		handleError(context, http.StatusOK, err.Error())
		return
	}

	setRefreshTokenCookie(context, token.RefreshToken)

	context.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"accessToken": token.AccessToken,
	})
}

func setRefreshTokenCookie(context *gin.Context, refreshToken string) {
	http.SetCookie(context.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		MaxAge:   30 * 24 * 60 * 60,
		HttpOnly: true,
		Secure:   gin.Mode() != gin.DebugMode,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
}

func handleError(context *gin.Context, status int, errMsg string) {
	context.JSON(status, gin.H{"error": errMsg})
}
