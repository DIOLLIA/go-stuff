package main

import "fmt"

type Notification struct {
}

func (n *Notification) sendDebitWalletNotification() {
	fmt.Println("Sending wallet debit notification")
}
func (n *Notification) sendCreditWalletNotification() {
	fmt.Println("Sending wallet credit notification")
}
