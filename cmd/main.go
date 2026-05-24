package main

import (
	log "log/slog"
	"net/http"

	"realtime-todolist/internal/todo"
)

func main() {
	mux := http.NewServeMux()

	todolist := todo.InitTodolist()

	log.Info("server started on localhost:8080")
	http.ListenAndServe(":8080", mux)
}
