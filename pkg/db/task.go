package db

import "database/sql"

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Comment string `json:"comment"`
	Title   string `json:"title"`
	Repeat  string `json:"repeat"`
}

//id INTEGER PRIMARY KEY AUTOINCREMENT,
//date CHAR(8) NOT NULL DEFAULT "",
//comment TEXT,
//title VARCHAR(64) NOT NULL,
//repeat VARCHAR

func AddTask(task *Task) (int64, error) {
	var id int64
	// определите запрос
	query := `INSERT INTO scheduler (date,comment,title,repeat) VALUES (:date,:comment,:title,:repeat,)`
	res, err := db.Exec(query, sql.Named("date", task.Date),
		sql.Named("comment", task.Comment),
		sql.Named("title", task.Title),
		sql.Named("repeat", task.Repeat) /*передайте параметры task.Date, task.Title и т.д.*/)

	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}
