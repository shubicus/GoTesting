package tests

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
	"sync"
	"testing"
	"time"
)

const (
	BaseURL     = "https://restful-booker.herokuapp.com"
	AuthPath    = "/auth"
	PingPath    = "/ping"
	BookingPath = "/booking"

	Cookie          = "Cookie"
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
	TokenIs         = "token="
)

var (
	Body = map[string]string{
		"username": "admin",
		"password": "password123",
	}

	TokenInstance *Token
	once          sync.Once
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

func GetToken() *Token {
	once.Do(func() {
		resp, err := newRestyClient().R().
			SetHeader("Content-Type", "application/json").
			SetBody(Body).
			Post(BaseURL + AuthPath)

		if err != nil {
			log.Fatalf("Can not obtain Token: %v", err)
		}

		err = json.Unmarshal(resp.Body(), &TokenInstance)
		if err != nil {
			log.Fatalf("Can not parse Token: %v", err)
		}
	})
	return TokenInstance
}
