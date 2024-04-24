package main

import (
	"log"
	"net/http"
)

const PORT = "localhost:6942"

func main() {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store}
	log.Fatal(http.ListenAndServe(PORT, server))
}
