package creditcard

import (
	"bufio"
	"os"
	"strings"
)

type CardInfo struct {
	Number string
	Valid  bool
	Brand  string
	Issuer string
}

func LoadInfoFiles(brandsPath, issuersPath string) (map[string]string, map[string]string, error) {
	brandMap := make(map[string]string)
	issuerMap := make(map[string]string)

	read := func(path string, target map[string]string) error {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if parts := strings.SplitN(line, ":", 2); len(parts) == 2 {
				target[parts[1]] = parts[0]
			}
		}
		return scanner.Err()
	}

	if err := read(brandsPath, brandMap); err != nil {
		return nil, nil, err
	}
	if err := read(issuersPath, issuerMap); err != nil {
		return nil, nil, err
	}
	return brandMap, issuerMap, nil
}

func GetCardInfo(number string, brands, issuers map[string]string) CardInfo {
	info := CardInfo{Number: number, Valid: IsValidLuhn(number)}
	info.Brand = "-"
	info.Issuer = "-"
	for prefix, brand := range brands {
		if strings.HasPrefix(number, prefix) {
			info.Brand = brand
			break
		}
	}
	for prefix, issuer := range issuers {
		if strings.HasPrefix(number, prefix) {
			info.Issuer = issuer
			break
		}
	}
	return info
}
