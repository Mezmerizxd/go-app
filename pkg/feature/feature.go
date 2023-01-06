package feature

import (
	"net/http"

	"github.com/mezmerizxd/go-app/pkg/cache"
	"github.com/mezmerizxd/go-app/pkg/data"
	"github.com/mezmerizxd/go-app/pkg/data/gta"
	"github.com/mezmerizxd/go-app/pkg/data/users"
	"github.com/mezmerizxd/go-app/pkg/firebase"
)

type Config struct {
	Cache cache.Cache
	Data  *data.Data
	Firebase firebase.Firebase
	Users users.Users
	Gta gta.Gta
}

type Feature struct {
	routes []*Route
	sockets []*Socket

	Cache cache.Cache
	Data  *data.Data
	Firebase firebase.Firebase
	Users users.Users
	Gta gta.Gta
}

func New(cfg *Config) *Feature {
	return &Feature{
		Cache: cfg.Cache,
		Data:  cfg.Data,
		Firebase: cfg.Firebase,
		Users: cfg.Users,
		Gta: cfg.Gta,
	}
}

func (f *Feature) Get(pattern string, handler routeHandlerFunc) *Route {
	return f.registerRoute(http.MethodGet, pattern, handler)
}

func (f *Feature) Post(pattern string, handler routeHandlerFunc) *Route {
	return f.registerRoute(http.MethodPost, pattern, handler)
}

func (f *Feature) Routes() []*Route {
	return f.routes
}

func (f *Feature) registerRoute(method, pattern string, handler routeHandlerFunc) *Route {
	r := &Route{
		method:  method,
		pattern: pattern,
		handler: handler,
	}
	f.routes = append(f.routes, r)
	return r
}

func (f *Feature) Socket(pattern string, handler socketHandlerFunc) *Socket {
	return f.registerSocket(pattern, handler)
}

func (f *Feature) Sockets() []*Socket {
	return f.sockets
}

func (f *Feature) registerSocket(pattern string, handler socketHandlerFunc) *Socket {
	s := &Socket{
		pattern: pattern,
		handler: handler,
	}
	f.sockets = append(f.sockets, s)
	return s
}