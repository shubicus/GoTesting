package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"testing"
	"time"
)

var bookingId string

func Test01HealthCheck(t *testing.T) {
	resp, err := newRestyClient().R().Get(BaseURL + PingPath)
	errHandle(t, "Request has failed: %v", err)

	assert.Equal(t, 201, resp.StatusCode(), "Expected status code 201")
}

func Test02CreateToken(t *testing.T) {
	token := GetToken().Token
	assert.NotEmpty(t, token, "Token should not be empty")
}

func Test03CreateBooking(t *testing.T) {
	var (
		booking = Booking{
			Firstname:   "Vad",
			Lastname:    "Shu",
			TotalPrice:  123,
			DepositPaid: false,
			BookingDates: BookingDates{
				Checkin: time.Now().Format("2006-01-02"),
				Checkout: time.Now().
					AddDate(0, 0, 15).
					Format("2006-01-02"),
			},
			AdditionalNeeds: "Toilet paper",
		}

		assertSoft = assert.New(t)
	)

	resp, err := newRestyClient().R().
		SetHeader(ContentType, ApplicationJson).
		SetBody(booking).
		Post(BaseURL + BookingPath)
	errHandle(t, "Request has failed: %v", err)
	bookingId = gjson.Get(resp.String(), "bookingid").String()

	assertSoft.Equal(200, resp.StatusCode(), "Expected status code 200")
	assertSoft.NotEmpty(bookingId, "Booking ID is present")
}

func Test04UpdateBooking(t *testing.T) {
	var (
		booking = Booking{
			Firstname:   "Dav",
			Lastname:    "Uhs",
			TotalPrice:  456,
			DepositPaid: false,
			BookingDates: BookingDates{
				Checkin:  time.Now().Format("2006-01-02"),
				Checkout: time.Now().AddDate(0, 0, 30).Format("2006-01-02"),
			},
			AdditionalNeeds: "Newspaper",
		}

		bookingActual Booking
	)

	resp, err := newRestyClient().R().
		SetHeaders(map[string]string{
			ContentType: ApplicationJson,
			Cookie:      TokenIs + GetToken().Token,
		}).
		SetBody(booking).
		Put(BaseURL + BookingPath + "/" + bookingId)
	errHandle(t, "Request has failed: %v", err)

	err = json.Unmarshal(resp.Body(), &bookingActual)
	errHandle(t, "Unmarshal has failed: %v", err)

	assert.Equal(t, booking, bookingActual, "The entity is not updated")
}

func Test05DeleteBooking(t *testing.T) {
	resp, err := newRestyClient().R().
		SetHeaders(map[string]string{
			ContentType: ApplicationJson,
			Cookie:      TokenIs + GetToken().Token,
		}).
		Delete(BaseURL + BookingPath + "/" + bookingId)
	errHandle(t, "Request has failed: %v", err)

	assert.Equal(t, 201, resp.StatusCode(), "Expected status code 201")
}
