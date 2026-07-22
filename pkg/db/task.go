package db

import (
	"database/sql"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Comment string `json:"comment"`
	Title   string `json:"title"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	// определите запрос
	query := `INSERT INTO scheduler (date,comment,title,repeat) VALUES (:date,:comment,:title,:repeat)`
	res, err := db.Exec(query, sql.Named("date", task.Date),
		sql.Named("comment", task.Comment),
		sql.Named("title", task.Title),
		sql.Named("repeat", task.Repeat))

	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, err
}

func Tasks(limit int) ([]*Task, error) {

	var tasks []*Task

	rows, err := db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit;", sql.Named("limit", limit))

	if err != nil {
		return []*Task{}, err
	}
	var foundOne bool

	for rows.Next() {
		foundOne = true

		var currTask Task

		err := rows.Scan(&currTask.ID, &currTask.Date, &currTask.Comment, &currTask.Title, &currTask.Repeat)

		if err != nil {
			return []*Task{}, err
		}

		tasks = append(tasks, &currTask)
	}
	if !foundOne {
		return []*Task{}, err
	}

	return tasks, err
}
