package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player  {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}

func NewLeague(r io.Reader) ([]Player, error) {
	l := make([]Player, 0)
	err := json.NewDecoder(r).Decode(&l)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return l, err
}
