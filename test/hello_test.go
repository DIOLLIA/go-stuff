package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello, Andrei!"
	result := Hello("Andrei")
	if result != expected {
		t.Errorf("Expected %q got %q", expected, result)
	}
}
