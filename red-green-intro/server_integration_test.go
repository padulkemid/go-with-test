package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	db, clean := createTempFile(t, "[]")
	defer clean()

	store, err := NewFileSystemPlayerStore(db)

	assertNoError(t, err)

	server := NewPlayerServer(store)
	player := "padul"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newGetScoreRequest(player))

		assertResponseStatus(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newLeagueRequest())

		assertResponseStatus(t, res.Code, http.StatusOK)

		got := getLeagueResponse(t, res.Body)
		want := []Player{
			{"padul", 3},
		}

		assertLeague(t, got, want)
	})
}
