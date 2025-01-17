package tests

import (
	. "MyRestyTesty/entities"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
	"testing"
	"time"
)

const (
	BaseURL     = "https://restful-booker.herokuapp.com"
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
	body = map[string]string{
		"username": "admin",
		"password": "password123",
	}

	tokenInstance *Token
)

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
