package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountMiddleware(api *ApiConfig) {
	s.Router.Use(middleware.Logger)
	s.Router.Use(api.MiddlewareCors)
}

func (s *Server) MakeApiHandlers(api *ApiConfig) {
	// Mounting API Router
	s.Router.Get("/health", api.HealthCheck)
	s.Router.Post("/overthewire", api.OverTheWire)
}
