package http

import (
	"context"

	"github.com/go-chi/chi/v5"

	"net/http"
)

type Server struct {
	Router *chi.Mux
	server *http.Server
}

func NewServer(addr string) Server {
	r := chi.NewRouter()

	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	return Server{
		Router: r,
		server: &srv,
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
