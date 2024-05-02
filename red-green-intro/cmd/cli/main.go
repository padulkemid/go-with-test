package main

import (
	"fmt"
	poker "hello/red-green-intro"
	"log"
	"os"
)

const DB_FILE_NAME = "game_cli.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(DB_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer close()

	fmt.Println("let's play some poker will ya!")
	fmt.Println("type '<name> wins!' to record a win!")
	poker.NewCli(store, os.Stdin).PlayPoker()
}
