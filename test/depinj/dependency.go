package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s!", name)
	if err != nil {
		return errors.New("error writing to writer")
	}
	return nil
}

func GreeterHandler(responseWriter http.ResponseWriter, request *http.Request) {
	err := Greet(responseWriter, "world")
	if err != nil {
		return
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreeterHandler)))
}
