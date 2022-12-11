package feature

import (
	"net/http"

	"github.com/mezmerizxd/go-app/pkg/cache"
	"github.com/mezmerizxd/go-app/pkg/data"
	"github.com/mezmerizxd/go-app/pkg/data/users"
)

type Config struct {
	Cache cache.Cache
	Data  *data.Data
	Users users.Users
}

type Feature struct {
	routes []*Route

	Cache cache.Cache
	Data  *data.Data
	Users users.Users
}

func New(cfg *Config) *Feature {
	return &Feature{
		Cache: cfg.Cache,
		Data:  cfg.Data,
		Users: cfg.Users,
	}
}

func (f *Feature) Get(pattern string, handler handlerFunc) *Route {
	return f.registerRoute(http.MethodGet, pattern, handler)
}

func (f *Feature) Post(pattern string, handler handlerFunc) *Route {
	return f.registerRoute(http.MethodPost, pattern, handler)
}

func (f *Feature) Routes() []*Route {
	return f.routes
}

func (f *Feature) registerRoute(method, pattern string, handler handlerFunc) *Route {
	r := &Route{
		method:  method,
		pattern: pattern,
		handler: handler,
	}
	f.routes = append(f.routes, r)
	return r
}