package card

// Details represents the expected format of the JSON payload
// @Description Card details including the card number, expiration month, and expiration year.
// @Description Used for validating card information.
type Details struct {
	CardNumber      string `json:"card_number"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
}
