package main

type Mosin struct {
	Gun
}

func newMosina() IGun {
	return &Mosin{
		Gun: Gun{
			name:     "Mosina rifle",
			power:    7,
			fullAuto: false,
		},
	}
}
