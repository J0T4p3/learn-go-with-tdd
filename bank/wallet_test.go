package bank

import (
	"testing"
)

func TestWallet_deposit(t *testing.T) {
	t.Run("Deposit positive amount", func(t *testing.T) {
		const amount = Bitcoin(20)
		w := Wallet{}
		err := w.Deposit(amount)

		assertNoError(t, err)
		assertBalance(t, w, amount)
	})

	t.Run("Deposit zero-or-negative amount", func(t *testing.T) {
		const negativeAmount = Bitcoin(-20)
		const zeroAmount = Bitcoin(0)
		const balance = Bitcoin(10)
		w := Wallet{balance: balance}

		err := w.Deposit(negativeAmount)
		assertError(t, err, ErrNegativeOrZeroTransaction)
		assertBalance(t, w, balance)

		err = w.Deposit(zeroAmount)
		assertError(t, err, ErrNegativeOrZeroTransaction)
		assertBalance(t, w, balance)
	})
}

func TestWallet_withdraw(t *testing.T) {
	t.Run("Positive withdraw", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(20)}
		err := w.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("Withdraw more than balance", func(t *testing.T) {
		const balance = Bitcoin(10)
		const withdrawAmount = Bitcoin(30)

		w := Wallet{balance: balance}
		err := w.Withdraw(withdrawAmount)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, w, balance)
	})

	t.Run("Withdraw negative or zero", func(t *testing.T) {
		const zeroAmount = Bitcoin(0)
		const negativeAmount = Bitcoin(-1)
		const balance = Bitcoin(10)
		w := Wallet{balance: balance}

		err := w.Withdraw(zeroAmount)
		assertError(t, err, ErrNegativeOrZeroTransaction)
		assertBalance(t, w, balance)

		err = w.Withdraw(negativeAmount)
		assertError(t, err, ErrNegativeOrZeroTransaction)
		assertBalance(t, w, balance)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got %#v want nil", err)
	}
}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("any errors to compare")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
