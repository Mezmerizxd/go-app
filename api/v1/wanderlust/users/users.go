package users

import (
	"context"
	"encoding/json"
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

func (f *usersFeature) getAllUsers(ctx context.Context, w responder.Responder, r *http.Request) {
	users, err := f.Data.Users.GetAllUsers(ctx)
	if err != nil {
		w.Error("json", "No users found")
		return
	}

	w.Success("json", users)
}

func (f *usersFeature) createUser(ctx context.Context, w responder.Responder, r *http.Request) {
	body := map[string]string{}
	json.NewDecoder(r.Body).Decode(&body)

	data := users.User{
		Firstname: body["firstname"],
		Lastname: body["lastname"],
		Username: body["username"],
		Email: body["email"],
		Description: body["description"],
	}

	success, err := f.Data.Users.CreateUser(ctx, data)
	if err != nil && success == false {
		w.Error("json", "Failed to create user")
		return
	}

	w.Success("json", nil)
}