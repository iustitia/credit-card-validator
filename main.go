package main

import (
	"fmt"
	cc "github.com/sumup-challenges/coding-challenge-op-go-iustitia/creditcard"
	"log"
)

func main() {
	cardNumber := "5237 2516 2477 8133"
	fmt.Println(cc.IsValidCardNumber(cardNumber))

	bank, err := cc.GetBankScheme(cardNumber)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(bank.String())

}