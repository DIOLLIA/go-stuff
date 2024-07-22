package main

func getGun(automatic bool) IGun {
	if automatic {
		return newAK()
	} else {
		return newMosina()
	}
}
