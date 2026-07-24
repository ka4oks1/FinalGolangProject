package db

import (
	"database/sql"
	"errors"
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

		if err = rows.Err(); err != nil {
			return []*Task{}, err
		}

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

func GetTask(id string) (*Task, error) {

	var task Task
	row := db.QueryRow("SELECT * FROM scheduler WHERE id = :id;", sql.Named("id", id))

	err := row.Scan(&task.ID, &task.Date, &task.Comment, &task.Title, &task.Repeat)

	if err != nil {
		return &Task{}, err
	}

	return &task, err
}

func UpdateTask(task *Task) error {

	query := `UPDATE scheduler SET date = :date,comment = :comment,title = :title,repeat = :repeat WHERE id = :id;`

	res, err := db.Exec(query, sql.Named("date", task.Date),
		sql.Named("comment", task.Comment),
		sql.Named("title", task.Title),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID))

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		errors.New(`incorrect id for updating task`)
	}
	return nil
}

func DeleteTask(id string) error {

	query := `DELETE FROM scheduler WHERE id = :id;`

	_, err := db.Exec(query, sql.Named("id", id))

	if err != nil {
		return err
	}

	return nil
}

func UpdateDate(next string, id string) error {

	query := `UPDATE scheduler SET date = :date WHERE id = :id;`

	_, err := db.Exec(query, sql.Named("date", next), sql.Named("id", id))

	if err != nil {
		return err
	}

	return nil
}
