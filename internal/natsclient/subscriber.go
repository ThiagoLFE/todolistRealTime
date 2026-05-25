package natsclient

import (
	"realtime-todolist/internal/sse"
	"realtime-todolist/internal/todo"

	"github.com/nats-io/nats.go"
)

func SubscribeTodoEvents(nc *nats.Conn, broker *sse.Broker, tl *todo.TodoList) {

	nc.Subscribe("todo.events", func(msg *nats.Msg) {

		broker.BroadcastTable(tl.Items)
	})
}
