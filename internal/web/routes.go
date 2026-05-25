package web

import (
	"net/http"
	"realtime-todolist/internal/sse"
	"realtime-todolist/internal/todo"

	"github.com/nats-io/nats.go"
)

func NewRouter(tl *todo.TodoList, nc *nats.Conn, broker *sse.Broker) http.Handler {

	mux := http.NewServeMux()

	// pages
	mux.HandleFunc("/", basePage(tl, nc))

	// realtime
	mux.HandleFunc("/events", EventsHandler(broker))

	// api
	mux.HandleFunc("POST /task/new", createTask(tl, nc))
	mux.HandleFunc("POST /task/{id}/toggle", toggleTask(tl, nc))
	mux.HandleFunc("POST /task/{id}/delete", deleteTask(tl, nc))

	// static
	fs := http.FileServer(http.Dir("../static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
