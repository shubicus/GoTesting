package entities

type Token struct {
	Token string `json:"token"`
}

type Booking struct {
	Firstname       string       `json:"firstname"`
	Lastname        string       `json:"lastname"`
	TotalPrice      int          `json:"totalprice"`
	DepositPaid     bool         `json:"depositpaid"`
	BookingDates    BookingDates `json:"bookingdates"`
	AdditionalNeeds string       `json:"additionalneeds"`
}

type BookingDates struct {
	Checkin  string `json:"checkin"`
	Checkout string `json:"checkout"`
}
