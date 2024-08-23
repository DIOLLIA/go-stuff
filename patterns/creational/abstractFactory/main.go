package main

import "fmt"

func main() {
	adidasFactory, _ := getSportFactory("adidas")
	nikeFactory, _ := getSportFactory("Nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeTShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeTShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s AbstractShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(t AbstractTShirt) {
	fmt.Printf("Logo: %s", t.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", t.getSize())
	fmt.Println()
}
