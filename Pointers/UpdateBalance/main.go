package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?
func updateBalance(c *customer, t transaction) (err error) {
	if t.transactionType == "deposit" {
		c.balance += t.amount
		return nil
	} else if t.transactionType == "withdrawal" {
		if t.amount > c.balance {
			var e error = errors.New("insufficient funds")
			return e
		} else {
			c.balance -= t.amount
			return nil
		}
	} else {
		var e error = errors.New("unknown transaction type")
		return e
	}
}
