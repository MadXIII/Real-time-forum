package server

import (
	"forum/database"
	"log"
	"net/http"
	"text/template"
)

type Server struct {
	store  database.Repository
	router http.ServeMux
	temp   *template.Template
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
	s.router.HandleFunc("/post", s.GetPost)
}

func (s *Server) ListenAndServe(port string) {
	s.Conf()
	http.ListenAndServe(port, &s.router)
}

func (s *Server) Parser() {
	s.temp, err = template.ParseFiles("../client/index.html")
	if err != nil {
		log.Println(err)
	}
}
