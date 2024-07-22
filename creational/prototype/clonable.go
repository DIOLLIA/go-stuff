package main

var INDENTATION = "  "

type Cloneable interface {
	print()
	clone() Cloneable
}
