package card

import (
	"testing"
	"time"
)

func TestIsValidLuhn(t *testing.T) {
	luhnTests := []struct {
		cardNumber string
		expected   bool
	}{
		{"4242424242424242", true},
		{"4000056655665556", true},
		{"5555555555554444", true},
		{"2223003122003222", true},
		{"5200828282828210", true},
		{"378282246310005", true},
		{"6011111111111117", true},
		{"6011000990139424", true},
		{"3530111333300000", true},
		{"3566002020360505", true},
		{"abcd", false},
		{"abcdefghijklmnop", false},
		{"1234567890123456", false},
		{"4242424242424241", false},
		{"1111111111111111", false},
	}

	for _, test := range luhnTests {
		if result := isValidLuhn(test.cardNumber); result != test.expected {
			t.Errorf("isValidLuhn(%s) = %v, want %v", test.cardNumber, result, test.expected)
		}
	}
}

func TestIsValidCardNumberLength(t *testing.T) {
	cardNumberLengthTests := []struct {
		length   int
		expected bool
	}{
		{12, true},
		{13, true},
		{16, true},
		{19, true},
		{11, false},
		{20, false},
		{0, false},
		{5, false},
	}

	for _, test := range cardNumberLengthTests {
		if result := isValidCardNumberLength(test.length); result != test.expected {
			t.Errorf("isValidCardNumberLength(%d) = %v, want %v", test.length, result, test.expected)
		}
	}
}

func TestIsValidExpirationMonth(t *testing.T) {
	expirationMonthTests := []struct {
		month    int
		expected bool
	}{
		{1, true},
		{5, true},
		{12, true},
		{0, false},
		{13, false},
	}

	for _, test := range expirationMonthTests {
		if result := isValidExpirationMonth(test.month); result != test.expected {
			t.Errorf("isValidExpirationMonth(%d) = %v, want %v", test.month, result, test.expected)
		}
	}
}

func TestIsValidExpirationYear(t *testing.T) {
	currentYear := time.Now().Year()
	expirationYearTests := []struct {
		year     int
		expected bool
	}{
		{currentYear, true},
		{currentYear + 1, true},
		{currentYear - 1, false},
	}

	for _, test := range expirationYearTests {
		if result := isValidExpirationYear(test.year); result != test.expected {
			t.Errorf("isValidExpirationYear(%d) = %v, want %v", test.year, result, test.expected)
		}
	}
}

func TestIsValidCardExpiration(t *testing.T) {
	currentYear, currentMonth, _ := time.Now().Date()
	lastMonth := int(currentMonth) - 1
	if lastMonth == 0 {
		lastMonth = 12
	}
	cardExpirationTests := []struct {
		month    int
		year     int
		expected bool
	}{
		{int(currentMonth), currentYear, true},
		{1, currentYear + 1, true},
		{lastMonth, currentYear, false},
		{12, currentYear - 1, false},
	}

	for _, test := range cardExpirationTests {
		if got := isValidCardExpiration(test.month, test.year); got != test.expected {
			t.Errorf("isValidCardExpiration(%d, %d) = %v, want %v", test.month, test.year, got, test.expected)
		}
	}
}

func TestIsValidCard(t *testing.T) {
	cardDetailsTests := []struct {
		card     Details
		expected error
	}{
		{Details{CardNumber: "4242424242424242", ExpirationMonth: 12, ExpirationYear: time.Now().Year() + 1}, nil},
		{Details{CardNumber: "5555555555554444", ExpirationMonth: int(time.Now().Month()) + 1, ExpirationYear: time.Now().Year()}, nil},
		{Details{CardNumber: "1234567890123456", ExpirationMonth: 12, ExpirationYear: time.Now().Year() + 1}, invalidCardNumber},
		{Details{CardNumber: "4242424242", ExpirationMonth: 12, ExpirationYear: time.Now().Year() + 1}, invalidCardNumber},
		{Details{CardNumber: "4242424242424242", ExpirationMonth: 12, ExpirationYear: time.Now().Year() - 1}, invalidCardExpirationYear},
		{Details{CardNumber: "4242424242424242", ExpirationMonth: int(time.Now().Month()) - 1, ExpirationYear: time.Now().Year()}, cardHasExpired},
		{Details{CardNumber: "4242424242424242", ExpirationMonth: 13, ExpirationYear: time.Now().Year()}, invalidCardExpirationMonth},
	}

	for _, test := range cardDetailsTests {
		if _, err := IsValidCard(&test.card); err != test.expected {
			t.Errorf("IsValidCard() = %v, want %v for card %+v", err, test.expected, test.card)
		}
	}
}
