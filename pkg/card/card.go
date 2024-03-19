package card

type CardDetails struct {
	CardNumber      string `json:"card_number"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
}