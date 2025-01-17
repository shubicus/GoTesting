package tests

import (
	. "MyRestyTesty/entities"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"log"
	"testing"
	"time"
)

var bookingId string

const Format = "Request has failed: %v"

func TestSuit(t *testing.T) {
	healthCheck()

	t.Run("Test01CreateToken", func(t *testing.T) {
		resp, err := newRestyClient().R().
			SetHeader(ContentType, ApplicationJson).
			SetBody(body).
			Post(BaseURL + AuthPath)
		errHandle(t, Format, err)
		token := gjson.Get(resp.String(), "token").String()

		assert.NotEmpty(t, token, "Token should not be empty")
	})

	t.Run("Test02CreateBooking", func(t *testing.T) {
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
		errHandle(t, Format, err)
		bookingId = gjson.Get(resp.String(), "bookingid").String()

		assertSoft.Equal(200, resp.StatusCode(), "Expected status code 200")
		assertSoft.NotEmpty(bookingId, "Booking ID is present")
	})

	t.Run("Test03UpdateBooking", func(t *testing.T) {
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

			headers = map[string]string{
				ContentType: ApplicationJson,
				Cookie:      TokenIs + getToken().Token,
			}

			bookingActual Booking
		)

		resp, err := newRestyClient().R().
			SetHeaders(headers).
			SetBody(booking).
			Put(BaseURL + BookingPath + Slash + bookingId)
		errHandle(t, Format, err)

		err = json.Unmarshal(resp.Body(), &bookingActual)
		errHandle(t, "Unmarshal has failed: %v", err)

		assert.Equal(t, booking, bookingActual, "The entity is not updated")
	})

	t.Run("Test04DeleteBooking", func(t *testing.T) {
		resp, err := newRestyClient().R().
			SetHeaders(map[string]string{
				ContentType: ApplicationJson,
				Cookie:      TokenIs + getToken().Token,
			}).
			Delete(BaseURL + BookingPath + Slash + bookingId)
		errHandle(t, Format, err)

		assert.Equal(t, 201, resp.StatusCode(), "Expected status code 201")
	})
}

func healthCheck() {
	resp, err := newRestyClient().R().Get(BaseURL + PingPath)
	if err != nil {
		log.Fatalf(Format, err)
	}

	if resp.StatusCode() != 201 {
		log.Fatalf("Actual Response StatusCode is: %v", resp.StatusCode())
	}
}
