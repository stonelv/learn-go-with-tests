package pointersanderrors

import (
	"errors"
)

//Bitcoin from base type 'int'
type Bitcoin int

//Wallet a struct contain balance
type Wallet struct {
	balance Bitcoin
}

//ErrInsufficientFunds err type to describe insufficient funds.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Deposite deposite some amount of money from wallet
func (w *Wallet) Deposite(amount Bitcoin) {
	//fmt.Printf("address of balance in Deposite is %v \n", &w.balance)
	w.balance += amount
}

//Withdraw withdraw some amount of money from wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

//Balance return current balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
