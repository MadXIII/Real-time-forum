package server

import (
	"forum/database"
	sessions "forum/sessions"
	"log"
	"net/http"
	"text/template"
)

var err error
var cookies map[int]*http.Cookie = map[int]*http.Cookie{}

type Server struct {
	store        database.Repository
	router       http.ServeMux
	cookiesStore sessions.Repository
}

func Init(store database.Repository, cookiesStore sessions.Repository) *Server {
	return &Server{
		store:        store,
		router:       *http.NewServeMux(),
		cookiesStore: cookiesStore,
	}
}

func (s *Server) Conf() {
	s.router.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("../client/js"))))
	s.router.HandleFunc("/", s.MainPage)
	s.router.HandleFunc("/signup", s.SignUp)
	s.router.HandleFunc("/signin", s.SignIn)
	s.router.HandleFunc("/post", s.GetPost)
}

func (s *Server) ListenAndServe(port string) {
	s.Conf()
	http.ListenAndServe(port, &s.router)
}

func Parser() *template.Template {
	temp, err := template.ParseFiles("../client/index.html")
	if err != nil {
		log.Println(err)
	}
	return temp
}
