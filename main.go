package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ka4oks1/FinalGolangProject/pkg/db"
	"github.com/ka4oks1/FinalGolangProject/pkg/rules"
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

	nextDate, err := rules.NextDate(time.Now(), "20240229", "y")
	fmt.Println(nextDate)
	//
	//nextDate2, err := rules.NextDate(time.Now(), "20260812", "y")
	//fmt.Println(nextDate2)

	err = http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal(err)
	}
}
