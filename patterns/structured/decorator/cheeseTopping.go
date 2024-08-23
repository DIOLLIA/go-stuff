package main

type CheeseTopping struct {
	pizza IPizza
}

func (pz *CheeseTopping) getPrice() int {

	return pz.pizza.getPrice() + 10
}
