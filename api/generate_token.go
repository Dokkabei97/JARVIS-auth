package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-auth/application"
	"is-deploy-auth/domain"
	"net/http"
)

type IssueTokenRequest struct {
	UserInfo     domain.UserInfo `json:"userInfo"`
	AccessToken  string          `json:"accessToken"`
	RefreshToken string          `json:"refreshToken"`
}

func MakeToken(context *gin.Context, jwtService application.JwtService, isReissue bool) {
	if isReissue {
		var body IssueTokenRequest

		if err := context.BindJSON(&body); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		token, err := jwtService.ReissueToken(body.AccessToken, body.RefreshToken, body.UserInfo)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		}

		context.JSON(http.StatusOK, gin.H{
			"status":       true,
			"accessToken":  token.AccessToken,
			"refreshToken": token.RefreshToken,
		})
	} else {
		var body IssueTokenRequest

		if err := context.BindJSON(&body); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		token, err := jwtService.IssueToken(body.UserInfo)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		}
		context.JSON(http.StatusOK, gin.H{
			"status":       true,
			"accessToken":  token.AccessToken,
			"refreshToken": token.RefreshToken,
		})
	}
}
