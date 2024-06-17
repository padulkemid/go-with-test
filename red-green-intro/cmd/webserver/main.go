package main

import (
	poker "hello/red-green-intro"
	"log"
	"net/http"
)

const (
	PORT         = "localhost:6942"
	DB_FILE_NAME = "game.db.json"
)

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(DB_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer close()

	server, err := poker.NewPlayerServer(store)
	if err != nil {
		log.Fatalf("error starting server, reason: %v", err)
	}

	if err := http.ListenAndServe(PORT, server); err != nil {
		log.Fatalf("could not listen on port %s -> %v", PORT, err)
	}
}
