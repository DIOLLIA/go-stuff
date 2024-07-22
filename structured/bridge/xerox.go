package main

import "fmt"

type Xerox struct {
}

func (x *Xerox) PrintFile() {
	fmt.Println("Printing by a XEROX Printer")
}
