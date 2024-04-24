package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func assertResponseStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d status, want %d status", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return req
}

func newPostWinRequest(name string) *http.Request {
	route := fmt.Sprintf("/players/%s", name)
	req, _ := http.NewRequest(http.MethodPost, route, nil)

	return req
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"padul": 20,
			"sena":  10,
		},
    []string{},
	}

	server := &PlayerServer{&store}

	t.Run("returns padul's score", func(t *testing.T) {
		// Given
		req := newGetScoreRequest("padul")
		res := httptest.NewRecorder()

		// When
		server.ServeHTTP(res, req)

		// Then
		assertResponseStatus(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body.String(), "20")
	})

	t.Run("returns sena's score", func(t *testing.T) {
		req := newGetScoreRequest("sena")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseStatus(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body.String(), "10")
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("damn")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseStatus(t, res.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
    nil,
	}

	server := &PlayerServer{
		&store,
	}

	t.Run("it returns accepted on POST", func(t *testing.T) {
    req := newPostWinRequest("padul")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseStatus(t, res.Code, http.StatusAccepted)

    if len(store.winCalls) != 1 {
      t.Errorf("got %d calls after winning record want %d", len(store.winCalls), 1)
    }

    if store.winCalls[0] != "padul" {
      t.Errorf("not correct it should be padul not %q", store.winCalls[0])
    }
	})
}
