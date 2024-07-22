package main

func main() {

	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	normalBuilder.setDoorType()
	normalBuilder.setWindowType()
	normalBuilder.setNumFloor()

	iglooBuilder.setDoorType()
	iglooBuilder.setWindowType()
	iglooBuilder.setNumFloor()

	print(normalBuilder.getHouse().ToString())
	//println(iglooBuilder.getHouse())
}
