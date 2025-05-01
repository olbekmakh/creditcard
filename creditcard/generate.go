package creditcard

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func Generate(template string, pick bool) error {
	stars := strings.Count(template, "*")
	if stars == 0 {
		return fmt.Errorf("no asterisks to replace")
	}
	if stars > 4 || !strings.HasSuffix(template, strings.Repeat("*", stars)) {
		return fmt.Errorf("invalid template format")
	}

	prefix := template[:len(template)-stars]
	max := intPow(10, stars)

	var results []string
	for i := 0; i < max; i++ {
		suffix := fmt.Sprintf("%0*d", stars, i)
		candidate := prefix + suffix
		if ValidateCardNumber(candidate) {
			results = append(results, candidate)
		}
	}

	if len(results) == 0 {
		return fmt.Errorf("no valid cards generated")
	}

	if pick {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(results[rand.Intn(len(results))])
		return nil
	}

	sort.Strings(results)
	for _, r := range results {
		fmt.Println(r)
	}
	return nil
}

func intPow(a, b int) int {
	result := 1
	for b > 0 {
		result *= a
		b--
	}
	return result
}
