package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"creditcard"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "no command specified")
		os.Exit(1)
	}

	switch args[1] {
	case "validate":
		validate(args[2:])
	case "generate":
		generate(args[2:])
	default:
		fmt.Fprintln(os.Stderr, "unknown command:", args[1])
		os.Exit(1)
	}
}

func validate(args []string) {
	readFromStdin := false
	var numbers []string

	for _, arg := range args {
		if arg == "--stdin" {
			readFromStdin = true
		} else {
			numbers = append(numbers, arg)
		}
	}

	if readFromStdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			numbers = append(numbers, strings.Fields(line)...)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "stdin error:", err)
			os.Exit(1)
		}
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
	pick := false
	var template string

	for _, arg := range args {
		if arg == "--pick" {
			pick = true
		} else if template == "" {
			template = arg
		}
	}

	if template == "" {
		fmt.Fprintln(os.Stderr, "missing card template")
		os.Exit(1)
	}

	if pick {
		card, err := creditcard.PickRandomCard(template)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(card)
		return
	}

	cards, err := creditcard.GenerateCardNumbers(template)
	if err != nil {
		os.Exit(1)
	}
	for _, c := range cards {
		fmt.Println(c)
	}
}
