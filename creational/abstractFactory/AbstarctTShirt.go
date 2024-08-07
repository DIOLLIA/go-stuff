package main

type AbstractTShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type TShirt struct {
	logo string
	size int
}

func (t *TShirt) setLogo(logo string) {
	t.logo = logo
}
func (t *TShirt) setSize(size int) {
	t.size = size
}
func (t *TShirt) getLogo() string {
	return t.logo
}
func (t *TShirt) getSize() int {
	return t.size
}
