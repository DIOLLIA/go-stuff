package main

type AKGun struct {
	Gun
}

func newAK() IGun {
	return &AKGun{
		Gun: Gun{
			name:     "AK47 gun",
			power:    4,
			fullAuto: true,
		},
	}
}
