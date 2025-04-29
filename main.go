package main

import (
	"creditcard/creditcard"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: creditcard <command>")
		os.Exit(1)
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "validate":
		validate(args)
	case "generate":
		generate(args)
	case "information":
		information(args)
	case "issue":
		issue(args)
	default:
		fmt.Fprintln(os.Stderr, "unknown command:", cmd)
		os.Exit(1)
	}
}

func validate(args []string) {
	useStdin := len(args) > 0 && args[0] == "--stdin"
	var numbers []string

	if useStdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			parts := strings.Fields(scanner.Text())
			numbers = append(numbers, parts...)
		}
	} else {
		numbers = args
	}

	status := 0
	for _, number := range numbers {
		if creditcard.IsValidLuhn(number) {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			status = 1
		}
	}
	os.Exit(status)
}

func generate(args []string) {
	usePick := false
	var input string
	for _, arg := range args {
		if arg == "--pick" {
			usePick = true
		} else {
			input = arg
		}
	}
	if usePick {
		n, err := creditcard.GenerateOne(input)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(n)
	} else {
		nums, err := creditcard.GenerateAll(input)
		if err != nil {
			os.Exit(1)
		}
		for _, n := range nums {
			fmt.Println(n)
		}
	}
}

func information(args []string) {
	var brandFile, issuerFile string
	useStdin := false
	var numbers []string

	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "--brands="):
			brandFile = strings.TrimPrefix(arg, "--brands=")
		case strings.HasPrefix(arg, "--issuers="):
			issuerFile = strings.TrimPrefix(arg, "--issuers=")
		case arg == "--stdin":
			useStdin = true
		default:
			numbers = append(numbers, arg)
		}
	}

	if brandFile == "" || issuerFile == "" {
		fmt.Fprintln(os.Stderr, "missing --brands or --issuers")
		os.Exit(1)
	}

	brandMap, issuerMap, err := creditcard.LoadInfoFiles(brandFile, issuerFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if useStdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			parts := strings.Fields(scanner.Text())
			numbers = append(numbers, parts...)
		}
	}

	for _, n := range numbers {
		info := creditcard.GetCardInfo(n, brandMap, issuerMap)
		fmt.Println(info.Number)
		if info.Valid {
			fmt.Println("Correct: yes")
		} else {
			fmt.Println("Correct: no")
		}
		fmt.Println("Card Brand:", info.Brand)
		fmt.Println("Card Issuer:", info.Issuer)
	}
}

func issue(args []string) {
	var brandPath, issuerPath, brand, issuer string

	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "--brands="):
			brandPath = strings.TrimPrefix(arg, "--brands=")
		case strings.HasPrefix(arg, "--issuers="):
			issuerPath = strings.TrimPrefix(arg, "--issuers=")
		case strings.HasPrefix(arg, "--brand="):
			brand = strings.TrimPrefix(arg, "--brand=")
		case strings.HasPrefix(arg, "--issuer="):
			issuer = strings.TrimPrefix(arg, "--issuer=")
		}
	}

	if brandPath == "" || issuerPath == "" || brand == "" || issuer == "" {
		fmt.Fprintln(os.Stderr, "missing --brands, --issuers, --brand or --issuer")
		os.Exit(1)
	}

	brandMap, issuerMap, err := creditcard.LoadInfoFiles(brandPath, issuerPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load brand/issuer files:", err)
		os.Exit(1)
	}

	num, err := creditcard.IssueCardNumber(brand, issuer, brandMap, issuerMap)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println(num)
}
