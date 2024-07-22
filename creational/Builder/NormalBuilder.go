package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (nb *NormalBuilder) setWindowType() {
	nb.windowType = "Wooden Windows"
}

func (nb *NormalBuilder) setDoorType() {
	nb.doorType = "Wooden Doors"
}

func (nb *NormalBuilder) setNumFloor() {
	nb.floor = 2
}

func (nb *NormalBuilder) getHouse() House {
	return House{
		doorType:   nb.doorType,
		windowType: nb.windowType,
		floor:      nb.floor,
	}
}
