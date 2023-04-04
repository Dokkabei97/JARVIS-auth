package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type JsonConfig struct {
	CommonBody CommonBody `json:"commonBody"`
	Config     string     `json:"config"`
}

func SyncSettingJson(context *gin.Context) {
	var body JsonConfig
	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config := strings.NewReader(body.Config)

	client := &http.Client{}
	api := fmt.Sprintf("http://%s%s/api/v1/setting/", body.CommonBody.Host, body.CommonBody.Port)
	req, err := http.NewRequest("PUT", api, config)

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}

	res, err := client.Do(req)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	defer res.Body.Close()

	context.JSON(http.StatusOK, gin.H{
		"message": res.Body,
	})
}
