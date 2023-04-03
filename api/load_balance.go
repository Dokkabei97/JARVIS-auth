package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoadBalance struct {
	CommonBody CommonBody `json:"commonBody"`
	Worker     string     `json:"worker"`
}

func Exclude(context *gin.Context) {
	var body LoadBalance
	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{}
	api := fmt.Sprintf("http://%s%s/api/v1/load-balance/exclude?worker=%s", body.CommonBody.Host, body.CommonBody.Port, body.Worker)
	req, err := http.NewRequest("PUT", api, nil)
	if err != nil {

	}
	res, err := client.Do(req)
	if err != nil {

	}
	defer res.Body.Close()

	context.JSON(http.StatusOK, gin.H{
		"message": body.Worker + " 제외되었습니다",
	})
}

func Restore(context *gin.Context) {
	var body LoadBalance
	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{}
	api := fmt.Sprintf("http://%s%s/api/v1/load-balance/restore", body.CommonBody.Host, body.CommonBody.Port)
	req, err := http.NewRequest("PUT", api, nil)
	if err != nil {

	}
	res, err := client.Do(req)
	if err != nil {

	}
	defer res.Body.Close()

	context.JSON(http.StatusOK, gin.H{
		"message": "연결이 복원되었습니다",
	})
}
