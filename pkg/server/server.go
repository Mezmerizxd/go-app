package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	v1 "github.com/mezmerizxd/go-app/api/v1"
	"github.com/mezmerizxd/go-app/pkg/version"
)

type Server struct {
	srv *http.Server
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, _ := upgrader.Upgrade(w, r, nil)
	defer c.Close()

	for {
		mt, message, _ := c.ReadMessage()

		log.Printf("received: %s", message)
		c.WriteMessage(mt, message)
	}

}

func New(addr string, cfg *version.Config) *Server {
	r := chi.NewRouter()

	// Middlewares
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*", "*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		ExposedHeaders:   []string{"Link", "Access-Control-Allow-Origin", "Accept"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Second * 15))

	// API
	r.Mount("/api/v1", v1.New(cfg))

	r.HandleFunc("/echo", echo)

	return &Server{
		srv: &http.Server{
			Addr:    addr,
			Handler: r,
		},
	}
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}