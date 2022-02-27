package main

import (
	"fmt"
	"log"

	"github.com/Msksgm/itddd-go-02-valueobject/domain/model/money"
	"github.com/shopspring/decimal"
)

func main() {
	amount1, err := decimal.NewFromString("1000")
	if err != nil {
		log.Fatal(err)
	}
	currency := "JPY"
	myMoney, err := money.NewMoney(amount1, currency)
	if err != nil {
		log.Fatal(err)
	}

	amount2, err := decimal.NewFromString("3000")
	if err != nil {
		log.Fatal(err)
	}
	allowance, err := money.NewMoney(amount2, currency)
	if err != nil {
		log.Fatal(err)
	}

	result, err := myMoney.Add(*allowance)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
