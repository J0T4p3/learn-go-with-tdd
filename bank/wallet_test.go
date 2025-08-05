package bank

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Retrieve balance", func(t *testing.T) {
		testAmount := Bitcoin(10)
		wallet := Wallet{}
		wallet.Deposit(testAmount)

		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("Got %d want %d", got, want)
		}
	})
}
