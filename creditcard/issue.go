package creditcard

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Issue(brandsFile, issuersFile, brand, issuer string) error {
	brands, err := loadMapping(brandsFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load brands:", err)
		return err
	}
	issuers, err := loadMapping(issuersFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load issuers:", err)
		return err
	}

	var brandPrefix string
	for prefix, b := range brands {
		if b == brand {
			brandPrefix = prefix
			break
		}
	}
	if brandPrefix == "" {
		fmt.Fprintln(os.Stderr, "Brand not found")
		return fmt.Errorf("brand not found")
	}

	var issuerPrefix string
	for prefix, i := range issuers {
		if i == issuer {
			issuerPrefix = prefix
			break
		}
	}
	if issuerPrefix == "" {
		fmt.Fprintln(os.Stderr, "Issuer not found")
		return fmt.Errorf("issuer not found")
	}

	if !strings.HasPrefix(issuerPrefix, brandPrefix) {
		fmt.Fprintln(os.Stderr, "Issuer does not match brand prefix")
		return fmt.Errorf("issuer mismatch")
	}

	rand.Seed(time.Now().UnixNano())
	for {
		bodyLen := 15 - len(issuerPrefix)
		body := ""
		for i := 0; i < bodyLen; i++ {
			body += fmt.Sprint(rand.Intn(10))
		}
		for i := 0; i < 10; i++ {
			lastDigit := fmt.Sprint(i)
			full := issuerPrefix + body + lastDigit
			if ValidateCardNumber(full) {
				fmt.Println(full)
				return nil
			}
		}
	}
}
