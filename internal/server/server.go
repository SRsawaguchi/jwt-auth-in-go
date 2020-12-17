package server

import (
	"context"
	"net/http"

	"github.com/SRsawaguchi/simple-jwt-auth-in-go/internal/endpoint"
	"github.com/go-chi/chi"
)

// Server represents http server
type Server struct {
	handler http.Handler
}

func (s *Server) setupRouter() {
	r := chi.NewRouter()

	r.Post("/signin", endpoint.Signin)
	r.Post("/login", endpoint.Login)
	r.Get("/refresh-token", endpoint.RefreshToken)
	r.Post("/memo", endpoint.NewMemo)
	r.Get("/memo", endpoint.GetMemo)

	s.handler = r
}

// Handler returns http.Handler of this server
func (s *Server) Handler() http.Handler {
	return s.handler
}

// Init initializes server
func (s *Server) Init(ctx context.Context) error {
	s.setupRouter()

	return nil
}

// Close cleanups used server resources
func (s *Server) Close(ctx context.Context) error {
	return nil
}
