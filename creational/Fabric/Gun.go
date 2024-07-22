package main

type Gun struct {
	name     string
	power    int
	fullAuto bool
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getPower() int {
	return g.power
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getAutoFire() bool {
	return g.fullAuto
}

func (g *Gun) setAutoFire(autoFire bool) {
	g.fullAuto = autoFire
}
