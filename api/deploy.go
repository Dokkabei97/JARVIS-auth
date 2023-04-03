package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeployBody struct {
	CommonBody CommonBody `json:"commonBody"`
	Worker     string     `json:"worker"`
	Arguments  string     `json:"arguments"`
}

func Deploy(context *gin.Context) {
	var body DeployBody
	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{}
	api := fmt.Sprintf("http://%s%s/api/v1/health-check/agent?worker=%s&arguments=%s", body.CommonBody.Host, body.CommonBody.Port, body.Worker, body.Arguments)
	req, err := http.NewRequest("PUT", api, nil)
	if err != nil {

	}
	res, err := client.Do(req)
	if err != nil {

	}
	defer res.Body.Close()

	context.JSON(http.StatusOK, gin.H{
		"message": "배포 쉘 스크립트가 실행 되었습니다.",
		"output":  res,
	})
}