package creditcard

import (
	"regexp"
	"strings"
)

func IsValidCardNumber(cardNumber string) bool {
	// clean up spaces
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")

	if !isValidFormat(cardNumber) {
		return false
	}

	checksum := 0
	var err error
	// calculate sum from right to left, double each second digit
	for i, j := len(cardNumber)-1, 0; i >= 0; i, j = i-1, j+1 {
		digit := getDigitFromRune(cardNumber[i])

		if j % 2 == 1 {
			digit, err = sumUntilSingleDigit(digit * 2)
			if err != nil {
				return false
			}
		}
		checksum += digit
	}

	if checksum % 10 == 0 {
		return true
	}

	return false
}

func isValidFormat(number string) bool {
	// limit 8 to 19 based on https://en.wikipedia.org/wiki/Payment_card_number#Structure
	reg := regexp.MustCompile(`^\d{8,19}$`)
	return reg.MatchString(number)
}

func getDigitFromRune(r uint8) int {
	return int(r - '0')
}

type OutsideScopeNumberError struct{}

func (e *OutsideScopeNumberError) Error() string {
	return "Number outside of scope"
}

func sumUntilSingleDigit(number int) (int, error) {
	if number > 18 {
		return -1, &OutsideScopeNumberError{}
	}

	if number < 10 {
		return number, nil
	}
	ones := number % 10
	tens := number / 10
	number = ones + tens
	return number, nil
}
