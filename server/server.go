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
	s.router.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("../client/js"))))
	s.router.HandleFunc("/", s.MainPage)
	s.router.HandleFunc("/signup", s.SignUp)
	// s.router.HandleFunc("/post", s.GetPost)

	// s.router.HandleFunc("/registration", s.Registration)
}

func (s *Server) ListenAndServe(port string) {
	s.Conf()
	http.ListenAndServe(port, &s.router)
}
