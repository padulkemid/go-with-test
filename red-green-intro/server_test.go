package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
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

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertContentType(t testing.TB, res *httptest.ResponseRecorder, want string) {
	t.Helper()

	if res.Result().Header.Get(contentType) != want {
		t.Errorf("res didn't have content type of %s, got %v", want, res.Result().Header)
	}
}

func getLeagueResponse(t testing.TB, b io.Reader) (p []Player) {
	t.Helper()

	err := json.NewDecoder(b).Decode(&p)
	if err != nil {
		t.Fatalf("unable to parse from server %q into Player, '%v'", b, err)
	}

	return
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

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)

	return req
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"padul": 20,
			"sena":  10,
		},
		[]string{},
		nil,
	}

	server := NewPlayerServer(&store)

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
		nil,
	}

	server := NewPlayerServer(&store)

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

func TestLeague(t *testing.T) {
	league := []Player{
		{"Padul", 20},
		{"Sena", 21},
	}
	store := StubPlayerStore{nil, nil, league}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		req := newLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := getLeagueResponse(t, res.Body)

		assertResponseStatus(t, res.Code, http.StatusOK)
		assertLeague(t, got, league)
    assertContentType(t, res, jsonType)
	})
}
