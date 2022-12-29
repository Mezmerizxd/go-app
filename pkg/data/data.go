package data

import (
	"github.com/mezmerizxd/go-app/pkg/cache"
	"github.com/mezmerizxd/go-app/pkg/data/gta"
	"github.com/mezmerizxd/go-app/pkg/data/users"
)

type Config struct {
	Cache cache.Cache
	Users users.Users
	Gta gta.Gta
}

type Data struct {
	Users users.Users
	Gta gta.Gta
}

func New(cfg *Config) *Data {
	return &Data{
		Users: users.New(cfg.Cache),
		Gta: gta.New(cfg.Cache),
	}
}
