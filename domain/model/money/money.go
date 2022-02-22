package money

import (
	"fmt"

	"github.com/Msksgm/itddd-go-02-valueobject/iterrors"
	"github.com/shopspring/decimal"
)

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
	defer iterrors.Wrap(&err, "money.Add(%v)", arg)
	if money.currency != arg.currency {
		err = fmt.Errorf("currency is different between money:%s and arg:%s", money.currency, arg.currency)
		return nil, err
	}
	newMoney, err := NewMoney(money.amount.Add(arg.amount), money.currency)
	return newMoney, err
}
