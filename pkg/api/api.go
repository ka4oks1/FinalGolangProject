package api

import "net/http"

func init() {
	webDir := "./web"

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	http.HandleFunc("/api/nextdate", HandleNextDate)
	http.HandleFunc("/api/task", mainTaskHandler)
	http.HandleFunc("/api/tasks", getTasksHandler)

}
