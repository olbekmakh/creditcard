# creditcard

A command-line utility written in **Go** for working with credit card numbers.
This project implements validation, generation, information lookup, and issuing of credit card numbers according to the **Alem School** specification.

The project is implemented using **only Go standard libraries**.

---

## Features

* **validate** — validate credit card numbers using the Luhn algorithm
* **generate** — generate possible valid credit card numbers
* **information** — detect card brand and issuer
* **issue** — issue a new valid credit card number

---

## Project Structure

```
creditcard/
├── cmd/
│   ├── generate.go
│   ├── info.go
│   ├── issue.go
│   └── luhn.go
├── brands.txt        # Card brands dictionary
├── issuers.txt       # Card issuers dictionary
├── go.mod
└── main.go
```

---

## Build

The project **must** be built from the root directory:

```bash
go build -o creditcard .
```

---

## Usage

### Validate

Validate a credit card number using the Luhn algorithm.

```bash
./creditcard validate "4400430180300003"
```

* On success: prints `OK` to stdout and exits with code `0`
* On failure: prints `INCORRECT` to stderr and exits with code `1`

Multiple numbers are supported:

```bash
./creditcard validate "4400430180300003" "4400430180300011"
```

Using stdin:

```bash
echo "4400430180300003" | ./creditcard validate --stdin
```

---

### Generate

Generate possible credit card numbers by replacing `*` (up to 4) at the end of the number.

```bash
./creditcard generate "440043018030****"
```

* All valid numbers are printed in **ascending order**
* Exits with status `1` if an error occurs

Randomly pick one valid number:

```bash
./creditcard generate --pick "440043018030****"
```

---

### Information

Get information about a credit card number using `brands.txt` and `issuers.txt`.

```bash
./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"
```

Example output:

```
4400430180300003
Correct: yes
Card Brand: VISA
Card Issuer: Kaspi Gold
```

If the card number is invalid:

```
Correct: no
Card Brand: -
Card Issuer: -
```

Supports stdin and multiple inputs.

---

### Issue

Generate a random valid credit card number for a specific brand and issuer.

```bash
./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"
```

Exits with status `1` if an error occurs.

---

## Dictionary Format

### brands.txt

```
VISA:4
MASTERCARD:51
MASTERCARD:52
AMEX:34
```

### issuers.txt

```
Kaspi Gold:440043
Forte Black:404243
Halyk Bonus:440563
```

Each line follows the format:

```
<NAME>:<PREFIX>
```

---

## Luhn Algorithm

The Luhn algorithm is used to validate credit card numbers:

1. Every second digit is doubled
2. If the result is greater than 9, subtract 9
3. Sum all digits
4. If the sum is divisible by 10, the number is valid

Implementation is located in `cmd/luhn.go`.

---

## Constraints

* Only Go standard libraries are allowed
* Code is formatted with `gofumpt`
* The program must not panic

---

## Author

**Oljas Bekmahan**
GitHub: [https://github.com/olbekmakh](https://github.com/olbekmakh)

---

## License

Educational project (Alem School)
