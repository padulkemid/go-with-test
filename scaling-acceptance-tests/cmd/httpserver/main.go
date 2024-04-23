package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
