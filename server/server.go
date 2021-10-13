package server

import (
	"forum/database"
	session "forum/sessions"
	"log"
	"net/http"
	"text/template"
)

//Server - store of DB, routes, cookies
type Server struct {
	store        database.Repository
	router       http.ServeMux
	cookiesStore session.Repository
}

//Init - Generator of Server struct
func Init(store database.Repository, cookiesStore session.Repository) *Server {
	return &Server{
		store:        store,
		router:       *http.NewServeMux(),
		cookiesStore: cookiesStore,
	}
}

//Conf - Hanlders to routes
func (s *Server) Conf() {
	s.router.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("../client/js"))))
	s.router.HandleFunc("/", s.MainPage)
	s.router.HandleFunc("/signin", s.middleWare(false, s.SignIn))
	s.router.HandleFunc("/signup", s.middleWare(false, s.SignUp))
	s.router.HandleFunc("/newpost", s.middleWare(true, s.CreatePost))
	s.router.HandleFunc("/logout", s.middleWare(true, s.LogOut))
}

//ListenAndServe - Listener with Configurations to ServMUX
func (s *Server) ListenAndServe(port string) {
	s.Conf()
	http.ListenAndServe(port, &s.router)
}

//Parser - to parse indexhtml
func Parser() *template.Template {
	temp, err := template.ParseFiles("../client/index.html")
	if err != nil {
		log.Println(err)
	}
	return temp
}
