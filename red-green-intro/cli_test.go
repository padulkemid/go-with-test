package poker_test

import (
	poker "hello/red-green-intro"
	"io"
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	cases := []struct {
		desc   string
		in     io.Reader
		winner string
	}{
		{
			desc:   "record padul win",
			in:     strings.NewReader("Padul wins!\n"),
			winner: "Padul",
		},
		{
			desc:   "record sena win",
			in:     strings.NewReader("Sena wins!\n"),
			winner: "Sena",
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			in := c.in
			store := &poker.StubPlayerStore{}

			cli := poker.NewCli(store, in)
			cli.PlayPoker()

			poker.AssertPlayerWins(t, store, c.winner)
		})
	}
}
