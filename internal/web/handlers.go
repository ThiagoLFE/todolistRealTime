package web

import (
	"net/http"
	"realtime-todolist/internal/components"
	"realtime-todolist/internal/todo"
	"strconv"

	"github.com/nats-io/nats.go"
	"github.com/starfederation/datastar-go/datastar"
)

func basePage(tl *todo.TodoList, nc *nats.Conn) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		components.BasePage(components.Table(tl.Items)).Render(req.Context(), w)
	}
}

func createTask(tl *todo.TodoList, nc *nats.Conn) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		signals := &components.NewTaskForm{}

		if err := datastar.ReadSignals(req, signals); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sse := datastar.NewSSE(w, req)

		if err := tl.Add(signals.Form.Name); err != nil {
			sse.PatchElementTempl(components.ErrorMsg(err.Error()))
		}
		nc.Publish("todo.events", []byte("New event"))
		sse.PatchElementTempl(components.Table(tl.Items))
	}
}

func toggleTask(tl *todo.TodoList, nc *nats.Conn) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		tl.ChangeState(id)
		sse := datastar.NewSSE(w, req)

		nc.Publish("todo.events", []byte("New event"))
		sse.PatchElementTempl(components.Table(tl.Items))
	}
}

func deleteTask(tl *todo.TodoList, nc *nats.Conn) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		tl.Remove(id)
		sse := datastar.NewSSE(w, req)

		nc.Publish("todo.events", []byte("New event"))
		sse.PatchElementTempl(components.Table(tl.Items))
	}
}
