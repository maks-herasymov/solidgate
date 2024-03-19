package card

import (
	"errors"
	"strconv"
	"time"
)

func isValidLuhn(cardNumber string) bool {
	var sum int
	var alternate bool

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}
		if alternate {
			digit *= 2
			if digit > 9 {
				digit = (digit % 10) + 1
			}
		}
		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}

func isValidCardNumberLength(length int) bool {
	return length >= 12 && length <= 19
}

func isValidExpirationMonth(month int) bool {
	return month >= 1 && month <= 12
}

func isValidExpirationYear(year int) bool {
	currentYear := time.Now().Year()
	return year >= currentYear
}

func isValidCardExpiration(month int, year int) bool {
	currentYear, currentMonth, _ := time.Now().Date()
	if year < currentYear {
		return false
	}
	if year == currentYear && month < int(currentMonth) {
		return false
	}
	return true
}

var (
	invalidCardNumber          = errors.New("Invalid card number")
	invalidCardExpirationMonth = errors.New("Invalid card expiration month")
	invalidCardExpirationYear  = errors.New("Invalid card expiration year")
	cardHasExpired             = errors.New("This card has expired")
)

func IsValidCard(card *Details) (int, error) {
	if !isValidCardNumberLength(len(card.CardNumber)) || !isValidLuhn(card.CardNumber) {
		return 1, invalidCardNumber
	}
	if !isValidExpirationMonth(card.ExpirationMonth) {
		return 2, invalidCardExpirationMonth
	}
	if !isValidExpirationYear(card.ExpirationYear) {
		return 3, invalidCardExpirationYear
	}
	if !isValidCardExpiration(card.ExpirationMonth, card.ExpirationYear) {
		return 4, cardHasExpired
	}
	return 0, nil
}
