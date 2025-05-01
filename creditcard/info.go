package creditcard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadMapping(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mapping := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		name := strings.TrimSpace(parts[0])
		prefix := strings.TrimSpace(parts[1])
		mapping[prefix] = name
	}
	return mapping, nil
}

func findMatch(number string, mapping map[string]string) string {
	for prefix, name := range mapping {
		if strings.HasPrefix(number, prefix) {
			return name
		}
	}
	return "-"
}

func Information(numbers []string, brandsFile, issuersFile string) {
	brands, err := loadMapping(brandsFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read brands:", err)
		return
	}
	issuers, err := loadMapping(issuersFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read issuers:", err)
		return
	}

	for _, number := range numbers {
		fmt.Println(number)
		valid := ValidateCardNumber(number)
		if valid {
			fmt.Println("Correct: yes")
		} else {
			fmt.Println("Correct: no")
		}
		fmt.Println("Card Brand:", findMatch(number, brands))
		fmt.Println("Card Issuer:", findMatch(number, issuers))
	}
}

