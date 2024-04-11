package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	b := bytes.Buffer{}

	Greet(&b, "Padul")

	got := b.String()
	want := "Oi, Padul"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
