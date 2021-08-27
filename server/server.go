package server

import (
	"forum/database"
	"net/http"
)

type Server struct {
	db     database.Repository
	router http.ServeMux
}

func Init(db database.Repository) *Server {
	return &Server{
		db:     db,
		router: *http.NewServeMux(),
	}
}

func (s *Server) Conf() {
	s.router.HandleFunc("/registration", s.Registration)
}

func (s *Server) ListenAndServe(port string) {
	s.Conf()
	http.ListenAndServe(port, &s.router)
}
