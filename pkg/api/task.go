package api

import (
	"encoding/json"
	"net/http"

	"github.com/ka4oks1/FinalGolangProject/pkg/db"
)

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// обработка других методов будет добавлена на следующих шагах
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

func postTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task
	json.Unmarshal(r.Body.)

}
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {

}
