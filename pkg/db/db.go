package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

// поправить создание таблицы
var shema string = `
CREATE TABLE scheduler (
id INTEGER PRIMARY KEY AUTOINCREMENT,
date CHAR(8) NOT NULL DEFAULT "",
comment TEXT,
title VARCHAR(64) NOT NULL,
repeat VARCHAR
)
`
var db *sql.DB

func Init(dbFile string) error {

	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	db, err = sql.Open("sqlite", dbFile)

	if err != nil {
		//log.Fatal(err)
	}

	if install {
		_, err = db.Exec(shema)
	}

	if err != nil {
		log.Fatal(err)
	}

	return err
	// если install равен true, после открытия БД требуется выполнить
	// sql-запрос с CREATE TABLE и CREATE INDEX

}
