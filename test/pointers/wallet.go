package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Euro int

type Wallet struct {
	balance Euro
}

// Stringer is implemented by any value that has a String method, which defines the “native” format for that value.
// The String method is used to print values passed as an operand to any format that accepts a string or to
// an unformatted printer such as Print.
// type Stringer interface { //this interface is defined in fmt package
//
//		String() string
//	}
func (e Euro) String() string {
	return fmt.Sprintf("%d EUR", e)
}

func (w *Wallet) Deposit(amount Euro) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Euro) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Euro {
	return w.balance
}
