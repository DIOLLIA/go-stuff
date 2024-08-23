package main

import "fmt"

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (a *SecurityCode) checkSecurityCode(incomeSecCode int) error {
	if incomeSecCode != a.code {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("Security Code Verified")
	return nil
}
