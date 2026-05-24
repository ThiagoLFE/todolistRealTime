package main

import (
	log "log/slog"
	"net/http"
	"realtime-todolist/internal/todo"
	"realtime-todolist/internal/web"
)

func main() {

	todolist := todo.InitTodolist()

	log.Info("server started on http://localhost:8080")
	http.ListenAndServe(":8080", web.NewRouter(&todolist))
}
