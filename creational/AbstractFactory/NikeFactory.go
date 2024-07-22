package main

// Nike Fabric

type Nike struct {
}

func (a *Nike) makeShoe() AbstractShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "Nike",
			size: 34,
		},
	}
}
func (a *Nike) makeTShirt() AbstractTShirt {
	return &NikeTShirt{
		TShirt: TShirt{
			logo: "Nike",
			size: 46,
		},
	}
}
