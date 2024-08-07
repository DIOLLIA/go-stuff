package main

// Adidas Fabric

type Adidas struct {
}

func (a *Adidas) makeShoe() AbstractShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 43,
		},
	}
}
func (a *Adidas) makeTShirt() AbstractTShirt {
	return &AdidasTShirt{
		TShirt: TShirt{
			logo: "adidas",
			size: 50,
		},
	}
}
