package server

import (
	"encoding/json"
	"forum/utils/database"
	newErr "forum/utils/internal/error"
	session "forum/utils/sessions"
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
	s.router.HandleFunc("/", s.Index)

	s.router.HandleFunc("/newpost", s.middleWare(true, s.Index))
	s.router.HandleFunc("/logout", s.middleWare(true, s.LogOut))

	s.router.HandleFunc("/api/", s.MainPage)
	s.router.HandleFunc("/api/signin", s.middleWare(false, s.SignIn))
	s.router.HandleFunc("/api/signup", s.middleWare(false, s.SignUp))
	s.router.HandleFunc("/api/newpost", s.middleWare(true, s.CreatePost))
	s.router.HandleFunc("/api/post", s.middleWare(false, s.GetPost))

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

//logger - send Notification to Client & log error
func logger(w http.ResponseWriter, status int, inputErr error) {
	if newErr.CheckErr(inputErr) {
		notify, err := json.Marshal(inputErr.Error())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.WriteHeader(status)
		w.Write(notify)
		log.Println(inputErr)
		return
	}
	w.WriteHeader(status)
	log.Println(inputErr)
}

//logout set cookies max age to -1
func logout(w http.ResponseWriter, ck *http.Cookie) {
	ck.MaxAge = -1
	http.SetCookie(w, ck)
	w.WriteHeader(http.StatusOK)
}
