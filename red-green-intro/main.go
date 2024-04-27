package main

import (
	"log"
	"net/http"
	"os"
)

const PORT = "localhost:6942"
const DB_FILE_NAME = "game.db.json"

func main() {
	db, err := os.OpenFile(DB_FILE_NAME, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s -> %v", DB_FILE_NAME, err)
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system for store -> %v \n", err)
	}


	server := NewPlayerServer(store)

	if err := http.ListenAndServe(PORT, server); err != nil {
		log.Fatalf("could not listen on port %s -> %v", PORT, err)
	}
}
