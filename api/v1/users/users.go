package users

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mezmerizxd/go-app/pkg/data/users"
	"github.com/mezmerizxd/go-app/pkg/responder"
)

func (f *usersFeature) getUser(ctx context.Context, w responder.Responder, r *http.Request) {
	urlparam := chi.URLParam(r, "user")

	if urlparam == "" {
		w.Error("json", "Invalid user")
		return
	}

	user, err := f.Data.Users.GetUserByUsername(ctx, urlparam)
	if err != nil {
		w.Error("json", "No user found")
		return
	}

	w.Success("json", makeUser(user))
}

func (f *usersFeature) createUser(ctx context.Context, w responder.Responder, r *http.Request) {
	request := map[string]string{}
	json.NewDecoder(r.Body).Decode(&request)

	body := users.User{
		Firstname: request["firstname"],
		Lastname: request["lastname"],
		Username: request["username"],
		Email: request["email"],
	}

	log.Printf("%s", body)

	success, err := f.Data.Users.CreateUser(ctx, body)
	if err != nil && success == false {
		w.Error("json", "Failed to create user")
		return
	}

	w.Success("json", nil)
}