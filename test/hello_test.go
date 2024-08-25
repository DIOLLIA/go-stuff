package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("If name passed, hello <name> returned", func(t *testing.T) {
		expected := "Hello, Andrei!"
		actual := Hello("Andrei")
		assertMessage(expected, actual, t)
	})
	t.Run("If empty string passed, hello world returned", func(t *testing.T) {
		expected := "Hello, World!"
		actual := Hello("")
		assertMessage(expected, actual, t)
	})
}
func assertMessage(expected, actual string, t *testing.T) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %q got %q", expected, actual)
	}
}
