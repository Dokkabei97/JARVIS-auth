package api

import (
	"github.com/gin-gonic/gin"
	"is-deploy-auth/domain"
	"net/http"
)

func MakeToken(context *gin.Context) {
	var body domain.UserInfo
	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
