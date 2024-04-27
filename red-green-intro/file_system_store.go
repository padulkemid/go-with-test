package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	db     *json.Encoder
	league League
}

func (f *FileSystemPlayerStore) GetLeague() League {
  sort.Slice(f.league, func (i, j int) bool  {
    return f.league[i].Wins > f.league[j].Wins
  })

	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	p := f.league.Find(name)

	if p != nil {
		p.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.db.Encode(f.league)
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("probelm getting file info %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}

	return nil
}

func NewFileSystemPlayerStore(db *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(db)
	if err != nil {
		return nil, fmt.Errorf("problem initialising file %v", err)
	}

	l, err := NewLeague(db)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", db.Name(), err)
	}

	return &FileSystemPlayerStore{
		json.NewEncoder(&tape{db}),
		l,
	}, nil
}
