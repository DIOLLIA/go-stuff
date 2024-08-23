package main

type IGun interface {
	getAutoFire() bool
	getName() string
	getPower() int
	setPower(int)
	setAutoFire(bool)
	setName(string)
}
