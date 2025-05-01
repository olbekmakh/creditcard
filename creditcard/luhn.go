package creditcard

import (
	"unicode"
)

func ValidateCardNumber(number string) bool {
	if len(number) < 13 {
		return false
	}

	var sum int
	double := false

	for i := len(number) - 1; i >= 0; i-- {
		r := rune(number[i])
		if !unicode.IsDigit(r) {
			return false
		}
		digit := int(r - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}

	return sum%10 == 0
}

