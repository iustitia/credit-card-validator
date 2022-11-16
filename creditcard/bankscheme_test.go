package creditcard

import "testing"

type cardSchemeTest struct {
	creditCardNumber string
	expected CardProvider
	expectedNoError bool
}

var schemeCases = []cardSchemeTest{
	// wrong
	{"abcdabcdabcd", UndefinedCardProvider, false},
	{"", UndefinedCardProvider, false},
	{"123", UndefinedCardProvider, false},
	{"5237 2516\n 2477 8133", UndefinedCardProvider, false},
	{"5237-2516-2477-8133", UndefinedCardProvider, false},
	{"111111111111111111", UndefinedCardProvider, false},
	{"2424240", UndefinedCardProvider, false},

	// ok card, different providers
	{"8171 9999 2766 0000", UndefinedCardProvider, true}, // China UnionPay
	{"6243 0300 0000 0001", UndefinedCardProvider, true}, // China UnionPay  - conflicting data
	{"5019 5555 4444 5555", UndefinedCardProvider, true}, // Dankort
	{"6033 4225 5384 5003", UndefinedCardProvider, true}, // Edenred  - conflicting data
	{"3600 6666 3333 44", UndefinedCardProvider, true}, // Diners
	{"5066 9911 1111 1118", UndefinedCardProvider, true}, // Elo
	// valid, no provider details
	{"4242 4242", UndefinedCardProvider, true}, // len 8
	{"42424242424242", UndefinedCardProvider, true},
	{"0042424242424242", UndefinedCardProvider, true},
	{"00000000042", UndefinedCardProvider, true},


	// ok, source: source: https://docs.adyen.com/development-resources/testing/test-card-numbers
	{"378282246310005", AmericanExpress,true},
	{"3714 4963 5398 431", AmericanExpress,true},
	{"374101012180018", AmericanExpress,true},
	{"3530111333300000", JCB,true},
	{"3569 9900 1009 5841", JCB,true},
	{"6759649826438453", Maestro,true},
	{"4012888888881881", Visa, true},
	{"4131 8400 0000 0003", Visa, true},
	{"4017 3400 0000 0003", Visa, true},
	{"4571 0000 0000 0001", Visa, true},
	{"5237 2516 2477 8133", MasterCard, true},
	{"5105105105105100", MasterCard, true},
	{"2222 4000 7000 0005", MasterCard, true},
	{"5577 0000 5577 0004", MasterCard, true},
	{"2223 5204 4356 0010", MasterCard, true},
	{"5105 1051 0510 5100", MasterCard, true},
	{"5105   1051   0510   5100   ", MasterCard, true},
}

func TestGetBankScheme(t *testing.T) {

	for _, test := range schemeCases {
		value, err := GetBankScheme(test.creditCardNumber)

		if err != nil {
			if test.expectedNoError {
				t.Errorf("Unexpected error for '%s'", test.creditCardNumber)
			}
		}

		if value != test.expected {
			t.Errorf("For credit card '%s': Expected provider %d (%s) got %d (%s). Expected no error: %t, got %v ",
				test.creditCardNumber, test.expected, test.expected.String(), value, value.String(), test.expectedNoError, err)
		}
	}
}


type cardProviderTest struct {
	creditCardNumber string
	expected CardProvider
}

var sCases = []cardProviderTest{
	// Amex
	{"378282246310005", AmericanExpress},
	{"348282246310005", AmericanExpress},
	{"3482822463100050", UndefinedCardProvider},
	{"34828224631000", UndefinedCardProvider},

	// JCB
	{"3528000012340000", JCB},
	{"3529000012340000", JCB},
	{"3530000012340000", JCB},
	{"3545000012340000", JCB},
	{"3589123456781234", JCB},
	{"3580123456781234", JCB},
	{"35731234567812345", JCB},
	{"3590123456781234", UndefinedCardProvider},
	{"3527123456781234", UndefinedCardProvider},
	{"3520123456781234", UndefinedCardProvider},

	// Maestro
	{"5018000012340000", Maestro},
	{"5020000012340000", Maestro},
	{"5038000012340000", Maestro},
	{"5893000012340000", Maestro},
	{"6304000012340000", Maestro},
	{"6759000012340000", Maestro},
	{"6761000012340000", Maestro},
	{"6762000012340000", Maestro},
	{"6763000012340000", Maestro},
	{"6763000012340000", Maestro},
	{"6767701234000012340", Maestro},
	{"676774000012340000", Maestro},
	{"676774000000000000123", UndefinedCardProvider},

	// Mastercard
	{"5100001234000012", MasterCard},
	{"5500001234000012", MasterCard},
	{"2221001234000012", MasterCard},
	{"2220001234000012", UndefinedCardProvider},
	{"2229001234000012", MasterCard},
	{"2300001234000012", MasterCard},
	{"2459001234000012", MasterCard},
	{"2638001234000012", MasterCard},
	{"2700001234000012", MasterCard},
	{"2711001234000012", MasterCard},
	{"2720001234000012", MasterCard},
	{"27200012340000120", UndefinedCardProvider},
	{"2721001234000012", UndefinedCardProvider},
	{"2730001234000012", UndefinedCardProvider},

	// Visa
	{"4000123400001", Visa},
	{"4100123400001", Visa},
	{"41001234000012", UndefinedCardProvider},
	{"4000123400001234", Visa},
	{"4917123400001234", Visa},
	{"49171234000012340", UndefinedCardProvider},

}

func TestGetCardProvider(t *testing.T) {

	for _, test := range sCases {
		value := getCardProvider(test.creditCardNumber)
		if value != test.expected {
			t.Errorf("For credit card '%s': Expected %d (%s) but got %d (%s)",
				test.creditCardNumber, test.expected, test.expected.String(), value, value.String())
		}
	}
}