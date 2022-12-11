package users

import (
	"github.com/mezmerizxd/go-app/pkg/data/users"
	"github.com/mezmerizxd/go-app/pkg/feature"
)

type usersUser struct {
	Username string `json:"username"`
	Description string `json:"description"`
}

type usersFeature struct {
	*feature.Feature
}

func New(parent *feature.Feature) {
	f := &usersFeature{Feature: parent}

	f.Get("/users/u/{user}", f.getUser)
	f.Get("/users/all", f.getAllUsers)
	f.Post("/users/create", f.createUser)
}

func makeUser(user *users.User) *usersUser {
	return &usersUser{
		Username: user.Username,
		Description: user.Description,
	}
}