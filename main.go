package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	db "github.com/ka4oks1/FinalGolangProject/pkg"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while reading .env")
	}

	dbFilePath := os.Getenv("TODO_DBFILE")

	if dbFilePath != "" {
		os.Mkdir(dbFilePath, 0755)
	}

	dbFile := dbFilePath + "scheduler.db"

	err = db.Init(dbFile)

	if err != nil {

		log.Fatal(err)
	}

	webDir := "./web"

	actualPort := os.Getenv("TODO_LIST_PORT")

	address := fmt.Sprintf("localhost:%s", actualPort)

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err = http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal(err)
	}
}
