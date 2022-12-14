package form

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mezmerizxd/go-app/pkg/responder"
)

func handleForm(ctx context.Context, w responder.Responder, r *http.Request) {
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)

	body := Form{
		Company: request["company"],
		Contact: request["contact"],
		Country: request["country"],
	}

	if (body.Company == "") {
		w.Error("json", "Please enter a company")
		return
	}
	if (body.Contact == "") {
		w.Error("json", "Please enter a contact")
		return
	}
	if (body.Country == "") {
		w.Error("json", "Please enter a country")
		return
	}
	
	output := append(table, body)
	table = output

	log.Printf("\n\nData service used.\n-> %s\n<- %s\n\n", request, output)

	w.Success("json", table)
}