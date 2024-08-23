package main

type TomatoTopping struct {
	pizza IPizza
}

func (pz *TomatoTopping) getPrice() int {

	return pz.pizza.getPrice() + 7
}
