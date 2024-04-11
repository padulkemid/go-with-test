package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(b io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(b, i)
		s.Sleep()
	}

	fmt.Fprint(b, finalWord)
}

func main() {
  s := &DefaultSleeper{}
	Countdown(os.Stdout, s)
}
