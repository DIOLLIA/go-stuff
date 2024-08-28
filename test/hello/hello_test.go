package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("If name passed, hello <name> returned", func(t *testing.T) {
		expected := "Hello, Andrei!"
		actual := Hello("Andrei", "EN")
		assertMessage(expected, actual, t)
	})
	t.Run("If empty name passed, hello world returned", func(t *testing.T) {
		expected := "Hello, World!"
		actual := Hello("", "")
		assertMessage(expected, actual, t)
	})
	t.Run("If language = DE and no name then should see hallo welt", func(t *testing.T) {
		expected := "Hallo, Welt!"
		actual := Hello("", "DE")
		assertMessage(expected, actual, t)
	})
	t.Run("If language = DE and name passed then should see hallo <name>", func(t *testing.T) {
		expected := "Hallo, Andre!"
		actual := Hello("Andre", "DE")
		assertMessage(expected, actual, t)
	})
}
func assertMessage(expected, actual string, t *testing.T) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %q got %q", expected, actual)
	}
}
