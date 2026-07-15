package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/ka4oks1/FinalGolangProject/pkg/api"
	"github.com/ka4oks1/FinalGolangProject/pkg/db"
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

	actualPort := os.Getenv("TODO_LIST_PORT")

	address := fmt.Sprintf("localhost:%s", actualPort)

	//nextDate, err := api.NextDate(time.Now(), "20260113", "d 7")
	//fmt.Println(nextDate)
	//
	//nextDate2, err := rules.NextDate(time.Now(), "20260812", "y")
	//fmt.Println(nextDate2)

	err = http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal(err)
	}
}
