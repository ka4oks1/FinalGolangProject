package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ka4oks1/FinalGolangProject/pkg/db"
)

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postTaskHandler(w, r)
		break
	case http.MethodDelete:
		deleteTaskHandler(w, r)
		break
	case http.MethodGet:
		getTaskHandler(w, r)
		break
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type IdResponse struct {
	Id string `json:"id"`
}

func postTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task
	body, err := io.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	err = json.Unmarshal(body, &task)

	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	if task.Title == "" {
		//http.Error(w, "task title is empty", http.StatusInternalServerError)
		writeJson(w, ErrorResponse{"task title is empty"})
		return
	}

	//if task.Repeat != "" {

	//		parsedTime, err := time.Parse(dateFormat, task.Date)

	//if err != nil {
	//	writeJson(w, ErrorResponse{err.Error()})
	//	return
	//}

	//	if !afterNow(parsedTime, time.Now()) {
	_, err = NextDate(time.Now(), task.Date, task.Repeat)
	//}

	if err != nil {
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	//}

	var id int64
	err = checkDate(&task)

	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		writeJson(w, ErrorResponse{err.Error()})
		return
	}

	id, err = db.AddTask(&task)

	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		writeJson(w, ErrorResponse{err.Error()})
		return

	}

	idStr := strconv.Itoa(int(id))
	writeJson(w, IdResponse{idStr})

}
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func checkDate(task *db.Task) error {
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format("20060102")
	}

	t, err := time.Parse("20060102", task.Date)

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
			task.Date = now.Format("20060102")
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
