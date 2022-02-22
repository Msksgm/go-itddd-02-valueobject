package money

import "github.com/shopspring/decimal"

type Money struct {
	amount   decimal.Decimal
	currency string
}

func NewMoney(amount decimal.Decimal, currency string) (_ *Money, err error) {
	money := new(Money)

	money.amount = amount
	money.currency = currency

	return money, nil
}

func (money *Money) Add(arg Money) (_ *Money, err error) {
	newMoney, err := NewMoney(money.amount.Add(arg.amount), money.currency)
	return newMoney, err
}
