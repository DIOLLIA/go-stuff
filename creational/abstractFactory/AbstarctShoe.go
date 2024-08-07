package main

type AbstractShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (t *Shoe) setLogo(logo string) {
	t.logo = logo
}
func (t *Shoe) setSize(size int) {
	t.size = size
}
func (t *Shoe) getLogo() string {
	return t.logo
}
func (t *Shoe) getSize() int {
	return t.size
}
