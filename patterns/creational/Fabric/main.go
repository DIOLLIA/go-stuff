package main

import "fmt"

func main() {

	akGun := getGun(true)
	mosin := getGun(false)
	printDetails(akGun)
	printDetails(mosin)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
