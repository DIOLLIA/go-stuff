package main

import "fmt"

type SportFactory interface {
	makeShoe() AbstractShoe
	makeTShirt() AbstractTShirt
}

func getSportFactory(brand string) (SportFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "Nike" {
		return &Nike{}, nil
	} else {
		return nil, fmt.Errorf("unknown brand: %s", brand)
	}
}
