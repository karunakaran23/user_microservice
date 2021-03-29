package server

import (
	"net/http"
	"user_microservice/handler"
	"user_microservice/middleware"
)

type Server struct {
	Router *http.ServeMux
}

func (s *Server) InitRoute(h *handler.Handler) {
	s.Router.HandleFunc("/Microservice/Name", middleware.JSONandCORS(h.MicroserviceName))
}
