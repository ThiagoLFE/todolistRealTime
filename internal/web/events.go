package web

import (
	"net/http"

	"realtime-todolist/internal/sse"

	"github.com/starfederation/datastar-go/datastar"
)

func EventsHandler(broker *sse.Broker) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		sse := datastar.NewSSE(w, r)

		broker.AddClient(sse)

		defer broker.RemoveClient(sse)

		<-r.Context().Done()
	}
}
