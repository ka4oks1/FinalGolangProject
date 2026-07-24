package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ka4oks1/FinalGolangProject/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

type IdResponse struct {
	Id string `json:"id"`
}

type EmptyStruct struct {
}

func mainTaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postTaskHandler(w, r)
		return
	case http.MethodGet:
		getOneTaskHandler(w, r)
		return
	case http.MethodPut:
		putTaskHandler(w, r)
		return
	case http.MethodDelete:
		deleteTaskHandler(w, r)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

const tasksLimit int = 50

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(tasksLimit)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}

func getOneTaskHandler(w http.ResponseWriter, r *http.Request) {
	idValue := r.FormValue("id")

	currTask, err := db.GetTask(idValue)

	if err != nil {
		writeJson(w, ErrorResponse{Error: err.Error()})
		return
	}

	writeJson(w, &currTask)

}
func postTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task
	body, err := io.ReadAll(r.Body)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	err = json.Unmarshal(body, &task)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	if task.Title == "" {
		writeJson(w, ErrorResponse{"task title is empty"})
		return
	}

	_, err = NextDate(time.Now(), task.Date, task.Repeat)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	var id int64
	err = checkDate(&task)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	id, err = db.AddTask(&task)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	idStr := strconv.Itoa(int(id))
	writeJson(w, IdResponse{idStr})

}

func putTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task
	body, err := io.ReadAll(r.Body)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	err = json.Unmarshal(body, &task)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	if task.Title == "" {
		writeJson(w, ErrorResponse{"task title is empty"})
		return
	}

	_, err = strconv.Atoi(task.ID)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	_, err = NextDate(time.Now(), task.Date, task.Repeat)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	err = checkDate(&task)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	err = db.UpdateTask(&task)

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	writeJson(w, EmptyStruct{})

}

func checkDate(task *db.Task) error {

	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(dateFormat)
	}

	t, err := time.Parse(dateFormat, task.Date)

	if err != nil {
		return err
	}
	var next string

	next, err = NextDate(now, task.Date, task.Repeat)
	if err != nil {
		return err
	}

	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(dateFormat)
		} else {
			if !(now.Format(dateFormat) == t.Format(dateFormat)) {
				task.Date = next
			}
		}
	}

	return err
}

func writeJson(w http.ResponseWriter, data any) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	jsData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(jsData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	idValue := r.FormValue("id")

	task, err := db.GetTask(idValue)

	if err != nil {
		writeJson(w, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.DeleteTask(task.ID)
	if err != nil {
		writeJson(w, ErrorResponse{Error: err.Error()})
		return
	}
	writeJson(w, EmptyStruct{})

}

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {

	idValue := r.FormValue("id")

	task, err := db.GetTask(idValue)

	if err != nil {
		writeJson(w, ErrorResponse{Error: err.Error()})
		return
	}

	if task.Repeat == "" {

		err = db.DeleteTask(task.ID)

		if err != nil {
			writeJson(w, ErrorResponse{Error: err.Error()})
			return
		}

		writeJson(w, EmptyStruct{})
		return
	}

	nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)

	if err != nil {
		writeJson(w, ErrorResponse{Error: err.Error()})
		return
	}

	err = db.UpdateDate(nextDate, task.ID)

	if err != nil {
		writeJson(w, ErrorResponse{Error: err.Error()})
		return
	}

	writeJson(w, EmptyStruct{})

}

func wrongMethod() {}
