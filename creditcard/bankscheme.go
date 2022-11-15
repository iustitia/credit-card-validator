package creditcard

import (
	"regexp"
	"strings"
)


type InvalidCardError struct{}

func (e *InvalidCardError) Error() string {
	return "Invalid card number"
}

type CardProvider int64

const (
	UndefinedCardProvider CardProvider = iota
	AmericanExpress
	JCB
	Maestro
	Visa
	MasterCard
)

func (s CardProvider) String() string {
	switch s {
	case UndefinedCardProvider:
		return "UndefinedCardProvider"
	case AmericanExpress:
		return "AmericanExpress"
	case JCB:
		return "JCB"
	case Maestro:
		return "Maestro"
	case Visa:
		return "Visa"
	case MasterCard:
		return "MasterCard"
	}
	return "unknown"
}

// Card schemes verified with https://en.wikipedia.org/wiki/Payment_card_number
var bankSchemeMap = map[CardProvider]string{
	AmericanExpress: `^3[4|7]\d{13}$`, // 34, 37
	JCB: `^352[8|9]|35[3-8][0-9]\d{12,15}$`, // 3528-3589
	Maestro: `^((5018|5020|5038|5893|6304|6759|6761|6762|6763)[0-9][0-9]|67677[0|4])\d{6,13}$`,
	Visa: `^4\d{12}(?:\d{3})?(?:\d{3})?$`, // 4
	MasterCard: `^(222[1-9]|22[3-9][0-9]|2[3-6][0-9][0-9]|27[0-1][0-9]|2720|5[1-5][0-9][0-9])\d{12}$`, // 2221-2720, 5100-5599
}


func getCardProvider(cardNumber string) CardProvider {
	for bankScheme, _ := range bankSchemeMap {
		re := regexp.MustCompile(bankSchemeMap[bankScheme])
		if re.MatchString(cardNumber) {
			return bankScheme
		}
	}
	return UndefinedCardProvider
}


func GetBankScheme(cardNumber string) (CardProvider, error) {
	// clean up spaces
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")

	if !IsValidCardNumber(cardNumber) {
		return UndefinedCardProvider, &InvalidCardError{}
	}

	return getCardProvider(cardNumber), nil
}
