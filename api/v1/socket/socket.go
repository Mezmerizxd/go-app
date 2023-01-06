package socket

import (
	"context"
	"log"
	"net/http"

	"github.com/mezmerizxd/go-app/pkg/feature"
	"github.com/mezmerizxd/go-app/pkg/socket"
)

type socketFeature struct {
	*feature.Feature
}

func New(parent *feature.Feature) {
	f := &socketFeature{Feature: parent}
	
	f.Socket("/echo", echo)
	f.Socket("/echo2", echo2)
}

func echo(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	c := socket.New(w, r)

	log.Printf("Connection established with /echo: %v", c.RemoteAddr())

	for {
		mt, message, _ := c.ReadMessage()

		log.Printf("received: %s", message)
		c.WriteMessage(mt, message)
	}
}

func echo2(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	c := socket.New(w, r)

	log.Printf("Connection established with /echo2: %v", c.RemoteAddr())

	for {
		mt, message, _ := c.ReadMessage()

		log.Printf("received: %s", message)
		c.WriteMessage(mt, message)
	}
}