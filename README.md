# Description

Your goal is to create two functions in Golang that verifies the correctness of a supplied credit card number and determine its card scheme. You can decide on the function's signature and types. It is mandatory that you add automated tests.

When ready, push your code.

## Function #1: Validity of the card number 

The following algorithm can be used to check validity of a card number:

1. Starting from the right, replace each **second** digit of the card number with its doubled value
2. When doubling a digit produces a 2-digit number (e.g 6 produces 12), then add those 2 digits (1+2 = 3)
3. Sum up all the digits

The card number is valid if the sum is divisible by 10

**Example**: Let's check if `5237 2516 2477 8133` is a valid credit card number.

1. Double each second digit: **10** 2 **6** 7 **4** 5 **2** 6 **4** 4 **14** 7 **16** 1 **6** 3
2. Add 2-digit numbers: **1** 2 6 7 4 5 2 6 4 4 **5** 7 **7** 1 6 3
3. Sum up all the digits: 70

70 is divisible by 10, so `5237 2516 2477 8133` is a **valid** credit card number

Please implement a function that given a credit card number returns if it is valid 

## Function #2: Known/supported card schemes

Card Scheme (Visa, MasterCard, JCB, etc) can be detected by the first digits of the card and the length of the card. 

**Example**

| Scheme           | Ranges           | Number of Digits | Example number   |
|---               |---               |---               |---
| American Express | 34,37            | 15               | 378282246310005  |
| JCB              | 3528-3589        | 16-19            | 3530111333300000 |
| Maestro*          | 5018, 5020, 5038, 5893, 6304, 6759, 6761, 6762, 6763, 676770, 676774    | 12-19            | 6759649826438453 |
| Visa             | 4                | 13,16,19         | 4012888888881881 |
| MasterCard       | 2221-2720, 51-55 | 16               | 5105105105105100 |

Based on https://en.wikipedia.org/wiki/Payment_card_number


Please implement a function that given a credit card number returns its card scheme.


## Run

To run code with single example, run in command line: `go run main.go`.
To run whole test suite, run in command line: `go test ./...` 


## Notes

### Possible extensions
Checking bank scheme was implemented only for card providers mentioned in description 
(Amex, JCB, Maestro, Vica, Mastercard).
It's possible to easily extend solution for additional card providers by extending `bankSchemeMap` with 
new entry with provider name and appropriate regular expression.

### * Maestro verification
Defined Maestro range looks outdated at this point, so newer range was implemented based on notes 
from Wikipedia (https://en.wikipedia.org/wiki/Payment_card_number).
There are providers like Dankort and China UnionPay that starts with 6. Also in 2022 new ranges for
Maestro UK were introduced.

### Signatures and returned value
It was decided to create enum for Bank Schemes, so returned value is integer, not string. It's easy
to change that using String() function provided for the enum. 

Values for reference:

| Card Provider | Enum value
|---            |---
| AmericanExpress | 1
| JCB           | 2
| Maestro       | 3
| Visa          | 4
| MasterCard    | 5
| Undefined     | 0

### Note on Golang conventions
As I am more experienced with Python than Go, some parts might be not strictly following Go conventions.
I'm open to discussion on this.
