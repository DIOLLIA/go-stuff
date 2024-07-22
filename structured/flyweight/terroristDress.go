package main

type TerroristDress struct {
	color string
}

func (td *TerroristDress) getColor() string {
	return td.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}
