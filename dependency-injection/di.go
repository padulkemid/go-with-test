package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Oi, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "padoels")
}

func main() {
	h := http.ListenAndServe(":5555", http.HandlerFunc(MyGreeterHandler))

	log.Fatal(h)
}
