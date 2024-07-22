package main

import "fmt"

type Mac struct {
	printer Printer
}

func (mac *Mac) Print() {
	mac.printer.PrintFile()
	fmt.Print("Print request for Mac\n")
}

func (mac *Mac) SetPrinter(printer Printer) {
	mac.printer = printer
}
