package server

import (
	"context"
	"net/http"

	"github.com/SRsawaguchi/jwt-auth-in-go/internal/endpoint"
	"github.com/SRsawaguchi/jwt-auth-in-go/internal/model"
	"github.com/go-chi/chi"
)

// Server represents http server
type Server struct {
	handler  http.Handler
	endpoint endpoint.Endpoint
}

func (s *Server) setupRouter() {
	r := chi.NewRouter()

	r.Post("/signin", s.endpoint.Signin)
	r.Post("/login", s.endpoint.Login)
	r.Get("/hello", s.endpoint.Hello)

	s.handler = r
}

// Handler returns http.Handler of this server
func (s *Server) Handler() http.Handler {
	return s.handler
}

// Init initializes server
func (s *Server) Init(ctx context.Context) error {
	s.setupRouter()
	s.endpoint.Users = map[string]*model.User{}

	return nil
}
