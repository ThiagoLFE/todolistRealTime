package sse

import (
	"realtime-todolist/internal/components"
	"realtime-todolist/internal/todo"
	"sync"

	"github.com/starfederation/datastar-go/datastar"
)

type Broker struct {
	mu sync.RWMutex

	clients map[*datastar.ServerSentEventGenerator]bool
}

func NewBroker() *Broker {
	return &Broker{
		clients: make(map[*datastar.ServerSentEventGenerator]bool),
	}
}

func (b *Broker) AddClient(
	sse *datastar.ServerSentEventGenerator,
) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.clients[sse] = true
}

func (b *Broker) RemoveClient(
	sse *datastar.ServerSentEventGenerator,
) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.clients, sse)
}

func (b *Broker) BroadcastTable(items []todo.Todo) {

	b.mu.RLock()
	defer b.mu.RUnlock()

	for client := range b.clients {
		client.PatchElementTempl(components.Table(items))
	}
}
