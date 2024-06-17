package poker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

const (
	contentType      = "content-type"
	jsonType         = "application/json"
	bufferSize       = 1024
	templateFileName = "game.html"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  bufferSize,
	WriteBufferSize: bufferSize,
}

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, jsonType)

	lt := p.store.GetLeague()
	json.NewEncoder(w).Encode(lt)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	p.template.Execute(w, nil)
}

func (p *PlayerServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	_, msg, _ := conn.ReadMessage()

	p.store.RecordWin(string(msg))
}

func NewPlayerServer(store PlayerStore) (*PlayerServer, error) {
	p := new(PlayerServer)

	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", templateFileName, err)
	}

	p.template = tmpl
	p.store = store

	r := http.NewServeMux()

	r.Handle("/league", http.HandlerFunc(p.leagueHandler))
	r.Handle("/players/", http.HandlerFunc(p.playersHandler))
	r.Handle("/game", http.HandlerFunc(p.gameHandler))
	r.Handle("/ws", http.HandlerFunc(p.wsHandler))
	p.Handler = r

	return p, nil
}
