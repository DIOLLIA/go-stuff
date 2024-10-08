package main

import "fmt"

type Account struct {
	name string
}

func newAccount(name string) *Account {
	return &Account{name: name}
}

func (a *Account) checkAccount(accName string) error {
	if accName != a.name {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
