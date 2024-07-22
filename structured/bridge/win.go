package main

import "fmt"

type Win struct {
	printer Printer
}

func (w *Win) Print() {
	w.printer.PrintFile()
	fmt.Println("Print request for windows")
}

func (w *Win) SetPrinter(p Printer) {
	w.printer = p
}
