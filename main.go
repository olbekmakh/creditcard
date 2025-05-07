package main

import (
	"bufio"
	"creditcard/cmd"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No command provided")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "validate":
		handleValidate(args)
	case "generate":
		handleGenerate(args)
	case "information":
		handleInformation(args)
	case "issue":
		handleIssue(args)
	default:
		fmt.Fprintln(os.Stderr, "Invalid command. Use --help for usage details.")
		os.Exit(1)
	}
}

func handleValidate(args []string) {
	useStdin := false
	var numbers []string

	for _, arg := range args {
		if arg == "--stdin" {
			useStdin = true
		} else {
			numbers = append(numbers, arg)
		}
	}

	if useStdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fields := strings.Fields(scanner.Text())
			numbers = append(numbers, fields...)
		}
	}

	for _, number := range numbers {
		if creditcard.ValidateCardNumber(number) {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			os.Exit(1)
		}
	}
}

func handleGenerate(args []string) {
	pick := false
	var template string

	for _, arg := range args {
		if arg == "--pick" {
			pick = true
		} else {
			template = arg
		}
	}

	if err := creditcard.Generate(template, pick); err != nil {
		os.Exit(1)
	}
}

func handleInformation(args []string) {
	var brandsFile, issuersFile string
	var numbers []string
	useStdin := false

	for _, arg := range args {
		if strings.HasPrefix(arg, "--brands=") {
			brandsFile = strings.TrimPrefix(arg, "--brands=")
		} else if strings.HasPrefix(arg, "--issuers=") {
			issuersFile = strings.TrimPrefix(arg, "--issuers=")
		} else if arg == "--stdin" {
			useStdin = true
		} else {
			numbers = append(numbers, arg)
		}
	}

	if useStdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fields := strings.Fields(scanner.Text())
			numbers = append(numbers, fields...)
		}
	}

	creditcard.Information(numbers, brandsFile, issuersFile)
}

func handleIssue(args []string) {
	var brandsFile, issuersFile, brand, issuer string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--brands=") {
			brandsFile = strings.TrimPrefix(arg, "--brands=")
		} else if strings.HasPrefix(arg, "--issuers=") {
			issuersFile = strings.TrimPrefix(arg, "--issuers=")
		} else if strings.HasPrefix(arg, "--brand=") {
			brand = strings.TrimPrefix(arg, "--brand=")
		} else if strings.HasPrefix(arg, "--issuer=") {
			issuer = strings.TrimPrefix(arg, "--issuer=")
		}
	}

	if err := creditcard.Issue(brandsFile, issuersFile, brand, issuer); err != nil {
		os.Exit(1)
	}
}
