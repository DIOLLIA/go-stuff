package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalGo   = "Go!"
	countdown = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(output io.Writer, sleeper Sleeper) {
	for i := countdown; i > 0; i-- {
		fmt.Fprintln(output, i)
		sleeper.Sleep()
	}
	fmt.Fprint(output, finalGo)

}
