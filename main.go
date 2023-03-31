package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"is-deploy-auth/config"
	"is-deploy-auth/router"
	"log"
)

const port = ":8080"

var err error

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.SetDb(config.SetDbInfo())), &gorm.Config{})
	if err != nil {
		log.Fatalf("에러 %s", err)
	}

	server := router.SetRouter(config.DB)
	err = server.Run(port)
	if err != nil {
		log.Fatalf("에러 %s", err)
	}
}
