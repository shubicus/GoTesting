package bdd

import (
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"time"
)

var BaseURL string

const Format = "Request has failed: %v"

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Panicf("Error getting WorkDirectory: %s", err)
	}

	dir := filepath.Dir(pwd)
	err = godotenv.Load(dir + "/.env")

	if err != nil {
		log.Panicf("Error loading .env file: %s", err)
	}

	BaseURL = os.Getenv("BASE_URL")
}

func NewRestyClient() *resty.Client {
	return resty.New().
		SetRetryWaitTime(5 * time.Second).
		EnableTrace()
}

func ErrHandleFatalf(format string, err error) {
	if err != nil {
		log.Fatalf(format, err)
	}
}
