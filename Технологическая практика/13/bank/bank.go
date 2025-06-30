package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

func (acc *BankAccount) Deposit(amount float64) {
	acc.balance += amount
}

func (acc *BankAccount) Withdraw(amount float64) error {
	if amount > acc.balance {
		return errors.New("недостаточно средств")
	}
	acc.balance -= amount
	return nil
}

func (acc BankAccount) GetBalance() float64 {
	return acc.balance
}

func main() {
	account := BankAccount{owner: "Алексей", balance: 500}

	account.Deposit(500)
	fmt.Println(account.GetBalance())

	account.Withdraw(200)
	fmt.Println(account.GetBalance())

	account.Withdraw(2000)
	fmt.Println(account.GetBalance())
}
