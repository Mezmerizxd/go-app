package data

import (
	"github.com/mezmerizxd/go-app/pkg/cache"
)

type Config struct {
	Cache cache.Cache
}

type Data struct {
}

func New(cfg *Config) *Data {
	return &Data{}
}
