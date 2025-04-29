package creditcard

import (
	"strings"
	"unicode"
)

func IsValidLuhn(number string) bool {
	number = strings.TrimSpace(number)
	if len(number) < 13 {
		return false
	}

	var sum int
	var alt bool

	for i := len(number) - 1; i >= 0; i-- {
		r := rune(number[i])
		if !unicode.IsDigit(r) {
			return false
		}
		d := int(r - '0')
		if alt {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		alt = !alt
	}

	return sum%10 == 0
}
