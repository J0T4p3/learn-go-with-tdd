package bank

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) GoString() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

var ErrNegativeOrZeroTransaction = errors.New("cannot transact negative or zero values")
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}
	if amount <= 0 {
		return ErrNegativeOrZeroTransaction
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount <= 0 {
		return ErrNegativeOrZeroTransaction
	}
	w.balance += amount
	return nil
}
