package users

import (
	"context"
	"errors"

	"github.com/mezmerizxd/go-app/pkg/cache"
)

type User struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type Users interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, user User) (bool, error)
}

type users struct {
	cache     cache.Cache
}

var users_list = []User{}

var ErrUserNotFound = errors.New("users: user not found")

func New(cacheImpl cache.Cache) Users {
	return &users{
		cache: cacheImpl,
	}
}

func (u *users) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	if len(users_list) == 0 {
		return nil, ErrUserNotFound
	}

	index := -1

	for i := 0; i < len(users_list); i++ {
		if users_list[i].Username == username {
			index = i
		}
	}

	if index == -1 {
		return nil, ErrUserNotFound
	}

	return &users_list[index], nil
}

func (u *users) CreateUser(ctx context.Context, user User) (bool, error) {
	output := append(users_list, user)
	users_list = output

	return true, nil
}