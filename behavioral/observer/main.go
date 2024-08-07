package main

func main() {
	shirtItem := newItem("Adik shirt")

	observer_1 := &Customer{id: "Alib@ba.com"}
	observer_2 := &Customer{id: "y@hoo.com"}

	shirtItem.register(observer_1)
	shirtItem.register(observer_2)

	shirtItem.updateAvailability()

	shirtItem.unregister(observer_1)
	shirtItem.inStock = false

	shirtItem.updateAvailability()
}
