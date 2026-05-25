package main

import (
	log "log/slog"
	"net/http"
	"realtime-todolist/internal/natsclient"
	"realtime-todolist/internal/sse"
	"realtime-todolist/internal/todo"
	"realtime-todolist/internal/web"
)

func main() {

	// connection with nats
	nc := natsclient.New()

	broker := sse.NewBroker()
	todolist := todo.InitTodolist()

	natsclient.SubscribeTodoEvents(nc, broker, &todolist)

	log.Info("server started on http://localhost:8080")
	http.ListenAndServe(":8080", web.NewRouter(&todolist, nc, broker))
}
