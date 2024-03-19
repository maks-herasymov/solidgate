package main

import (
	"strconv"
	"time"
)

type CardDetails struct {
	CardNumber      string
	ExpirationMonth int
	ExpirationYear  int
}

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

func IsValidCard(card *CardDetails) bool {
	return isValidCardNumberLength(len(card.CardNumber)) &&
		isValidLuhn(card.CardNumber) &&
		isValidExpirationMonth(card.ExpirationMonth) &&
		isValidExpirationYear(card.ExpirationYear) &&
		isValidCardExpiration(card.ExpirationMonth, card.ExpirationYear)
}
