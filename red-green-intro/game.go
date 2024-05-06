package poker

import "time"

type Texas interface {
	Start(numOfPlayers int)
	Finish(winner string)
}

type Game struct {
	alerter BlindAlerter
	store   PlayerStore
}

func (g *Game) Start(numOfPlayers int) {
	increment := time.Duration(5+numOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + increment
	}
}

func (g *Game) Finish(winner string) {
	g.store.RecordWin(winner)
}

func NewGame(alerter BlindAlerter, store PlayerStore) Texas {
	return &Game{
		alerter,
		store,
	}
}
