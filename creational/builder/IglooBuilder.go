package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (nb *IglooBuilder) setWindowType() {
	nb.windowType = "snow Windows"
}

func (nb *IglooBuilder) setDoorType() {
	nb.doorType = "snow Doors"
}

func (nb *IglooBuilder) setNumFloor() {
	nb.floor = 1
}

func (nb *IglooBuilder) getHouse() House {
	return House{
		doorType:   nb.doorType,
		windowType: nb.windowType,
		floor:      nb.floor,
	}
}
