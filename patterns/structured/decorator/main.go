package main

import "fmt"

func main() {
	veggie := &VeggiMania{}

	tomatoVeggie := &TomatoTopping{pizza: veggie}

	cheesyTomatoVeggiePizza := &CheeseTopping{pizza: tomatoVeggie}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", cheesyTomatoVeggiePizza.getPrice())

}
