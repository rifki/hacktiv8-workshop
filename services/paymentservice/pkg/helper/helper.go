package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	value := os.Getenv(key)
	if len(value) == 0 {
		return ""
	}
	return value
}
