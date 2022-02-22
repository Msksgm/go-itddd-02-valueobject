package money

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
)

func TestNewMoney(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount, err := decimal.NewFromString("100")
		if err != nil {
			t.Fatal(err)
		}
		currency := "JPY"
		got, err := NewMoney(amount, currency)
		if err != nil {
			t.Fatal(err)
		}
		want := &Money{amount: amount, currency: currency}
		if diff := cmp.Diff(got, want, cmp.AllowUnexported(Money{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
}

func TestAdd(t *testing.T) {
	amount1, err := decimal.NewFromString("100")
	if err != nil {
		t.Fatal(err)
	}
	currency := "JPY"
	money1 := &Money{amount: amount1, currency: currency}
	amount2, err := decimal.NewFromString("200")
	if err != nil {
		t.Fatal(err)
	}
	money2 := &Money{amount: amount2, currency: currency}
	got, err := money1.Add(*money2)
	if err != nil {
		t.Error(err)
	}

	wantAmount, err := decimal.NewFromString("300")
	if err != nil {
		t.Fatal(err)
	}
	want := &Money{amount: wantAmount, currency: currency}

	if diff := cmp.Diff(got, want, cmp.AllowUnexported(Money{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
