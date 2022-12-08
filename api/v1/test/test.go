package test

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/mezmerizxd/go-app/pkg/responder"
)

var randomReturns = []string{
	"Hello world",
	"Testing testing 123",
	"I am a robot",
	"I was written in Go",
	"How are you",
	"Hire me to make big $$",
}

func test(ctx context.Context, w responder.Responder, r *http.Request) {
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)

	rdmIndex := rand.Intn(len(randomReturns))

	response := TestData{
		Message: randomReturns[rdmIndex],
	}

	log.Printf("\n\nTest service used.\n-> %s\n<- %s\n\n", request, response)

	w.Error("json", "Error")
}
