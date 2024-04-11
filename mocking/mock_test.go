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
}
