package main

import "fmt"

const helloPrefix = "Hello, "

func main() {
	fmt.Print(Hello("Andrei!"))
}

func Hello(name string) string {
	if name == "" {
		return helloPrefix + "World!"
	}
	return helloPrefix + name + "!"
}
