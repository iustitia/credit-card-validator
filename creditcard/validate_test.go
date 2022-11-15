package creditcard

import "testing"

type cardTest struct {
	creditCardNumber string
	expected bool
}

var tCases = []cardTest{
	// wrong
	cardTest{"abcdabcdabcd", false},
	cardTest{"", false},
	cardTest{"123", false},
	cardTest{"5237 2516\n 2477 8133", false},
	cardTest{"5237-2516-2477-8133", false},
	cardTest{"111111111111111111", false},
	cardTest{"111111111111111111", false},
	cardTest{"2424240", false},
	cardTest{"2424 2424 2424 2424 2424", false},

	// ok
	cardTest{"5237 2516 2477 8133", true},
	cardTest{"378282246310005", true},
	cardTest{"3530111333300000", true},
	cardTest{"6759649826438453", true},
	cardTest{"4012888888881881", true},
	cardTest{"5105105105105100", true},
	cardTest{"5105 1051 0510 5100", true},
	cardTest{"5105   1051   0510   5100   ", true},
	cardTest{"4242 4242", true}, // len 8
	cardTest{"42424242424242", true},
	cardTest{"0042424242424242", true},
	cardTest{"00000000042", true},

	// ok, source: https://docs.adyen.com/development-resources/testing/test-card-numbers
	cardTest{"3700 0000 0000 002", true},
	cardTest{"6703 0000 0000 0000 003", true},
}

func TestIsValidCardNumber(t *testing.T) {

	for _, test := range tCases {
		value := IsValidCardNumber(test.creditCardNumber)
		if value != test.expected {
			t.Errorf("For credit carrd '%s': Expected %t but got %t", test.creditCardNumber, test.expected, value)
		}
	}
}

type runeTest struct {
	rune uint8
	expectedInt int
}

var runeTestCases = []runeTest{
	{'0', 0},
	{'1', 1},
	{'2', 2},
	{'3', 3},
	{'4', 4},
	{'7', 7},
	{'9', 9},
}


func TestGetDigitFromRune(t *testing.T) {

	for _, test := range runeTestCases {
		value := getDigitFromRune(test.rune)
		if value != test.expectedInt {
			t.Errorf("For rune '%v': Expected %d but got %d", test.rune, test.expectedInt, value)
		}
	}
}

type sumUntilSingleDigitTest struct {
	input int
	expectedInt int
	expectedNoErr bool
}

var sumUntilSingleDigitTestCases = []sumUntilSingleDigitTest{
	{0, 0, true},
	{4, 4, true},
	{9, 9, true},
	{10, 1, true},
	{13, 4, true},
	{18, 9, true},
	{99, -1, false},
}


func TestSumUntilSingleDigit(t *testing.T) {

	for _, test := range sumUntilSingleDigitTestCases {
		value, err := sumUntilSingleDigit(test.input)
		if err != nil && test.expectedNoErr == true {
			t.Errorf("Unexpected error for value %d, '%s'", test.input, err)
		}
		if value != test.expectedInt {
			t.Errorf("For in '%d': Expected %d but got %d", test.input, test.expectedInt, value)
		}
	}
}