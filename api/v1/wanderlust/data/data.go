package data

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mezmerizxd/go-app/pkg/responder"
)

func handleData(ctx context.Context, w responder.Responder, r *http.Request) {
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)

	body := Data{
		Title: request["title"],
		Message: request["message"],
		Other: request["other"],
	}

	// Check for title
	if (body.Title == "") {
		w.Error("json", "Please enter a title")
		return
	}

	// Check for message
	if (body.Message == "") {
		w.Error("json", "Please enter a message")
		return
	}
	// Other can be nil

	log.Printf("\n\nData service used.\n-> %s\n<- %s\n\n", request, body)

	w.Success("json", body)
}