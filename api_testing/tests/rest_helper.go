package tests

import (
	. "api_testing/entities"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
	"time"
)

const (
	AuthPath    = "/auth"
	PingPath    = "/ping"
	BookingPath = "/booking"
	Slash       = "/"

	Cookie          = "Cookie"
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
	TokenIs         = "token="
)

var (
	BaseURL = ""
	body    = map[string]string{
		"username": "admin",
		"password": "password123",
	}
	tokenInstance *Token
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Panicf("Error getting WorkDirectory: %s", err)
	}

	err = godotenv.Load(pwd + "/../.env")

	if err != nil {
		log.Panicf("Error loading .env file: %s", err)
	}

	BaseURL = os.Getenv("BASE_URL")
}

func newRestyClient() *resty.Client {
	return resty.New().
		SetRetryWaitTime(5 * time.Second).
		EnableTrace()
}

func errHandle(t *testing.T, format string, err error) {
	if err != nil {
		t.Fatalf(format, err)
	}
}

func getToken() *Token {
	if tokenInstance == nil {
		resp, err := newRestyClient().R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post(BaseURL + AuthPath)
		if err != nil {
			log.Fatalf("Can not obtain Token: %v", err)
		}

		err = json.Unmarshal(resp.Body(), &tokenInstance)
		if err != nil {
			log.Fatalf("Can not parse Token: %v", err)
		}
	}

	return tokenInstance
}
