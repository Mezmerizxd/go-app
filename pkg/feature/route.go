package feature

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mezmerizxd/go-app/pkg/responder"
)

type routeHandlerFunc func(ctx context.Context, w responder.Responder, req *http.Request)
type socketHandlerFunc func(ctx context.Context, w http.ResponseWriter, req *http.Request)

type Route struct {
	method  string
	pattern string
	handler routeHandlerFunc
}

type Socket struct {
	pattern string
	handler socketHandlerFunc
}

func (r *Route) InjectRoute(mux *chi.Mux) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		resp := responder.New(w, req)
		r.handler(req.Context(), resp, req)
	})

	mux.Group(func(group chi.Router) {
		switch r.method {
		case http.MethodGet:
			group.Get(r.pattern, handler)

		case http.MethodPost:
			group.Post(r.pattern, handler)
		}
	})
}

func (s *Socket) InjectSocket(mux *chi.Mux) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		s.handler(req.Context(), w, req)
	})

	mux.Group(func(group chi.Router) {
		group.HandleFunc(s.pattern, handler)
	})
}