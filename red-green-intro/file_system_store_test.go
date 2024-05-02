package poker

import (
	"os"
	"testing"
)

func createTempFile(t testing.TB, data string) (*os.File, func()) {
	t.Helper()

	temp, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file -> %v", err)
	}

	temp.Write([]byte(data))

	rmFile := func() {
		fileName := temp.Name()

		temp.Close()
		os.Remove(fileName)
	}

	return temp, rmFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("we got an error, it shouldn't be -> %v \n", err)
	}
}

func TestFileSystemStore(t *testing.T) {
	db, clean := createTempFile(t, `[
		{"Name": "Padul", "Wins": 10},
		{"Name": "Sena", "Wins": 33}
	]`)

	defer clean()

	store, err := NewFileSystemPlayerStore(db)

	assertNoError(t, err)

	t.Run("get league from a reader", func(t *testing.T) {
		got := store.GetLeague()
		want := League{
			{"Sena", 33},
			{"Padul", 10},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		got := store.GetPlayerScore("Padul")
		want := 10

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for players", func(t *testing.T) {
		store.RecordWin("Padul")

		got := store.GetPlayerScore("Padul")
		want := 11

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new player", func(t *testing.T) {
		store.RecordWin("Kiting")

		got := store.GetPlayerScore("Kiting")
		want := 1

		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		file, clean := createTempFile(t, "")
		defer clean()

		_, err := NewFileSystemPlayerStore(file)

		assertNoError(t, err)
	})

	t.Run("should sort league", func(t *testing.T) {
		got := store.GetLeague()
		want := League{
			{"Sena", 33},
			{"Padul", 11},
			{"Kiting", 1},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}
