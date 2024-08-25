package main

import "fmt"

func main() {
	fmt.Print(Hello("Andrei"))
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
