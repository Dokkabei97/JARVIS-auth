package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const PATH = "/home/danawa/is-deploy-console"

func GetSecretKey() []byte {
	err := godotenv.Load(PATH)
	if err != nil {
		log.Fatalf("[ERROR] LoadEnvFile : %s\n", err)
	}

	return []byte(os.Getenv("NEXTAUTH_SECRET"))
}
