package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	_ = Greet(&buffer, "Andrei")

	result := buffer.String()
	expected := "Hello, Andrei!"

	if expected != result {
		t.Errorf("expected %q result %q", expected, result)
	}
}
