package poker_test

import (
	"bytes"
	"fmt"
	poker "hello/red-green-intro"
	"io"
	"strings"
	"testing"
	"time"
)

type SpyAlert struct {
	scheduledAt time.Duration
	amount      int
}

func (s SpyAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.scheduledAt)
}

type SpyBlindAlerter struct {
	alerts []SpyAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(dur time.Duration, amt int, to io.Writer) {
	newAlert := SpyAlert{dur, amt}
	s.alerts = append(s.alerts, newAlert)
}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
}

func (g *GameSpy) Start(p int, to io.Writer) {
	g.StartedWith = p
	g.StartCalled = true
}

func (g *GameSpy) Finish(w string) {
	g.FinishedWith = w
}

func assertScheduleAlert(t testing.TB, got, want SpyAlert) {
	t.Helper()
	if got.amount != want.amount {
		t.Errorf("got amount %d, want %d", got.amount, want.amount)
	}

	if got.scheduledAt != want.scheduledAt {
		t.Errorf("got scheduled time of %v, want %v", got.scheduledAt, want.scheduledAt)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCli(t *testing.T) {
	t.Run("record wins with CLI", func(t *testing.T) {
		cases := []struct {
			desc   string
			in     io.Reader
			winner string
		}{
			{
				desc:   "record padul win",
				in:     strings.NewReader("1\nPadul wins!\n"),
				winner: "Padul",
			},
			{
				desc:   "record sena win",
				in:     strings.NewReader("1\nSena wins!\n"),
				winner: "Sena",
			},
		}
		for _, c := range cases {
			t.Run(c.desc, func(t *testing.T) {
				in := c.in
				store := &poker.StubPlayerStore{}
				alerter := &SpyBlindAlerter{}
				out := &bytes.Buffer{}
				game := poker.NewGame(alerter, store)

				cli := poker.NewCli(in, out, game)
				cli.PlayPoker()

				poker.AssertPlayerWins(t, store, c.winner)
			})
		}
	})

	t.Run("it schedules blind values", func(t *testing.T) {
		in := strings.NewReader("1\nPadul wins!\n")
		store := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		out := &bytes.Buffer{}
		game := poker.NewGame(blindAlerter, store)

		cli := poker.NewCli(in, out, game)
		cli.PlayPoker()

		cases := []SpyAlert{
			{0 * time.Second, 100},
			{6 * time.Minute, 200},
			{12 * time.Minute, 300},
			{18 * time.Minute, 400},
			{24 * time.Minute, 500},
			{30 * time.Minute, 600},
			{36 * time.Minute, 800},
			{42 * time.Minute, 1000},
			{48 * time.Minute, 2000},
			{54 * time.Minute, 4000},
			{60 * time.Minute, 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.amount, c.scheduledAt), func(t *testing.T) {
				if len(blindAlerter.alerts) <= 1 {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduleAlert(t, got, c)
			})
		}
	})

	t.Run("it prompts the user to enter number of players", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &SpyBlindAlerter{}
		store := &poker.StubPlayerStore{}
		game := poker.NewGame(blindAlerter, store)

		cli := poker.NewCli(
			in,
			out,
			game,
		)

		cli.PlayPoker()

		got := out.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []SpyAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, c := range cases {
			t.Run(fmt.Sprint(c), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduleAlert(t, got, c)
			})
		}
	})

	t.Run("it prompts the user to enter number of players and start the game", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCli(in, out, game)
		cli.PlayPoker()

		got := out.String()
		want := poker.PlayerPrompt

		assertString(t, got, want)

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run(
		"it prints an error when a non numeric value is entered and does not start the game",
		func(t *testing.T) {
			out := &bytes.Buffer{}
			in := strings.NewReader("Pies\n")
			game := &GameSpy{}

			cli := poker.NewCli(in, out, game)
			cli.PlayPoker()

			got := out.String()
			want := poker.PlayerPrompt + poker.SillyWord

			assertString(t, got, want)

			if game.StartCalled {
				t.Errorf("game shouldn't start")
			}
		},
	)
}
