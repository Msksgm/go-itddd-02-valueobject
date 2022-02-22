package money

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
)

func TestNewMoney(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount, err := decimal.NewFromString("1.234")
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
