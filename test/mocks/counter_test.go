package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	result := buffer.String()
	expected := `3
2
1
Go!`

	if expected != result {
		t.Errorf("expected %q result %q", expected, result)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, expected 3 got %d", spySleeper.Calls)
	}

}
