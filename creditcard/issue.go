package creditcard

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func IssueCardNumber(brand, issuer string, brandMap, issuerMap map[string]string) (string, error) {
	var brandPrefix, issuerPrefix string
	for prefix, name := range brandMap {
		if name == brand {
			brandPrefix = prefix
			break
		}
	}
	for prefix, name := range issuerMap {
		if name == issuer {
			issuerPrefix = prefix
			break
		}
	}
	if brandPrefix == "" || issuerPrefix == "" || !strings.HasPrefix(issuerPrefix, brandPrefix) {
		return "", errors.New("invalid brand/issuer combination")
	}

	length := 16
	if brand == "AMEX" {
		length = 15
	}

	rand.Seed(time.Now().UnixNano())
	for attempts := 0; attempts < 1000; attempts++ {
		restLen := length - len(issuerPrefix)
		var sb strings.Builder
		sb.WriteString(issuerPrefix)
		for i := 0; i < restLen-1; i++ {
			sb.WriteByte('0' + byte(rand.Intn(10)))
		}
		for i := 0; i <= 9; i++ {
			candidate := sb.String() + fmt.Sprint(i)
			if IsValidLuhn(candidate) {
				return candidate, nil
			}
		}
	}
	return "", errors.New("could not generate valid number")
}
