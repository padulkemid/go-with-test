package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct{
	balance Bitcoin
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amt Bitcoin) error  {
	if amt > w.balance {
		return errors.New("can't withdraw insufficient balance")
	}

	w.balance -= amt

	return nil
}

func (b Bitcoin) String() string  {
	return fmt.Sprintf("%d BTC", b)
}
