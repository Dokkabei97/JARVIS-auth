package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const PATH = "/home/danawa/is-deploy-console"

// GetSecretKey PATH경로에 .env파일에  NEXTAUTH_SECRET를 가져온다.
func GetSecretKey() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[ERROR] LoadEnvFile : %s\n", err)
	}

	return []byte(os.Getenv("NEXTAUTH_SECRET"))
}
