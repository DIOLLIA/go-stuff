package main

type CounterTerroristDress struct {
	color string
}

func (td *CounterTerroristDress) getColor() string {
	return td.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "blue"}
}
