package poker_test

import (
	poker "hello/red-green-intro"
	"os"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedule the alerts for 5 players", func(t *testing.T) {
		alerter := &SpyBlindAlerter{}
		store := &poker.StubPlayerStore{}
		game := poker.NewGame(alerter, store)

		game.Start(5, os.Stdout)

		cases := []SpyAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, c := range cases {
			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
			}

			got := alerter.alerts[i]
			assertScheduleAlert(t, got, c)
		}
	})

	t.Run("schedules alerts on game for 7 players", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		alerter := &SpyBlindAlerter{}
		game := poker.NewGame(alerter, store)

		game.Start(7, os.Stdout)

		cases := []SpyAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, c := range cases {
			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
			}

			got := alerter.alerts[i]
			assertScheduleAlert(t, got, c)
		}
	})
}

func TestGame_Finish(t *testing.T) {
	t.Run("should finish the game", func(t *testing.T) {
		alerter := &SpyBlindAlerter{}
		store := &poker.StubPlayerStore{}
		game := poker.NewGame(alerter, store)

		winner := "Padul"

		game.Finish("Padul")
		poker.AssertPlayerWins(t, store, winner)
	})
}
