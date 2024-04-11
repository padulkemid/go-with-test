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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep()  {
  c.sleep(c.duration)
}

func Countdown(b io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(b, i)
		s.Sleep()
	}

	fmt.Fprint(b, finalWord)
}

func main() {
  longSleep := 1 * time.Second
	s := &ConfigurableSleeper{longSleep, time.Sleep}

	Countdown(os.Stdout, s)
}
