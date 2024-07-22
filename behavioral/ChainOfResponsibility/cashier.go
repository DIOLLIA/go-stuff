package main

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("5k $ for the first time would be enough.")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
