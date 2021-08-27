package server

import (
	"forum/database"
	"net/http"
)

type Server struct {
	store  database.Repository
	router http.ServeMux
}

func Init(store database.Repository) *Server {
	return &Server{
		store:  store,
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
