package main

import "fmt"

type House struct {
	windowType string
	doorType   string
	floor      int
}

func (h House) ToString() string {
	return fmt.Sprintf("doors: %s, windows: %s, floors: %d", h.doorType, h.windowType, h.floor)
}
