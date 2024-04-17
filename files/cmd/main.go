package main

import (
	blogpost "hello/files"
	"log"
	"os"
)

func main() {
	posts, err := blogpost.NewPostFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(posts)
}
