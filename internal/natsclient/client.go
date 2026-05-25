package natsclient

import (
	"log"

	"github.com/nats-io/nats.go"
)

func New() *nats.Conn {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to nats")

	return nc
}
