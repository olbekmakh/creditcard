package creditcard

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func GenerateAll(pattern string) ([]string, error) {
	count := strings.Count(pattern, "*")
	if count > 4 {
		return nil, errors.New("too many asterisks")
	}
	if !strings.HasSuffix(pattern, strings.Repeat("*", count)) {
		return nil, errors.New("asterisks must be at the end")
	}

	base := strings.TrimRight(pattern, "*")
	limit := 1
	for i := 0; i < count; i++ {
		limit *= 10
	}

	var result []string
	for i := 0; i < limit; i++ {
		suffix := fmt.Sprintf("%0*d", count, i)
		full := base + suffix
		if IsValidLuhn(full) {
			result = append(result, full)
		}
	}

	sort.Strings(result)
	return result, nil
}

func GenerateOne(pattern string) (string, error) {
	valids, err := GenerateAll(pattern)
	if err != nil || len(valids) == 0 {
		return "", errors.New("no valid numbers")
	}
	rand.Seed(time.Now().UnixNano())
	return valids[rand.Intn(len(valids))], nil
}
