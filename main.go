package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	webDir := "./web"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while reading .env")
	}

	actualPort := os.Getenv("TODO_LIST_PORT")

	address := fmt.Sprintf("localhost:%s", actualPort)

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err = http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal(err)
	}
}
