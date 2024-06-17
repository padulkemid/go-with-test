package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

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

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)

	return req
}

func createPlayerServer(t testing.TB, store PlayerStore) *PlayerServer {
	t.Helper()

	server, err := NewPlayerServer(store)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}

	return server
}

func createWSDial(t testing.TB, url string) *websocket.Conn {
	t.Helper()

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func createWSMessage(t testing.TB, conn *websocket.Conn, msg string) {
	t.Helper()

	err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
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

	server := createPlayerServer(t, &store)

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
	store := &StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}

	server := createPlayerServer(t, store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		req := newPostWinRequest("padul")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseStatus(t, res.Code, http.StatusAccepted)
		AssertPlayerWins(t, store, "padul")
	})
}

func TestLeague(t *testing.T) {
	league := []Player{
		{"Padul", 20},
		{"Sena", 21},
	}
	store := StubPlayerStore{nil, nil, league}
	server := createPlayerServer(t, &store)

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

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		s := &StubPlayerStore{}
		server := createPlayerServer(t, s)

		req := newGameRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseStatus(t, res.Code, http.StatusOK)
	})

	t.Run("when we get a message over a websocket it is a winner of a game", func(t *testing.T) {
		store := &StubPlayerStore{}
		winner := "Padul"
		server := httptest.NewServer(createPlayerServer(t, store))

		defer server.Close()

		wsUrl := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws := createWSDial(t, wsUrl)
		defer ws.Close()

		createWSMessage(t, ws, "Padul")

		time.Sleep(10 * time.Millisecond)
		AssertPlayerWins(t, store, winner)
	})
}
