package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOp struct {
	Calls []string
}

func (s *SpyCountdownOp) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOp) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("should pass countdown", func(t *testing.T) {
		b := &bytes.Buffer{}
		s := &SpySleeper{}

		Countdown(b, s)

		got := b.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if s.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 3 got %d", s.Calls)
		}
	})

	t.Run("should sleep before every new print", func(t *testing.T) {
		s := &SpyCountdownOp{}

		Countdown(s, s)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(want, s.Calls) {
			t.Errorf("wanted calls %v got %v", want, s.Calls)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(d time.Duration) {
	s.durationSlept = d
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
