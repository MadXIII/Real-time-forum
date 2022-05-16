package server

import (
	"encoding/json"
	"fmt"
	"forum/internal/database"
	newErr "forum/internal/error"
	session "forum/internal/sessions"
	"log"
	"net/http"
	"text/template"
)

// Server - store of DB, routes, cookies
type Server struct {
	store   database.Repository
	router  http.ServeMux
	session session.Repository
	hub     *Hub
}

// Init - Generator of Server struct
func NewServer(store database.Repository, session session.Repository) *Server {
	return &Server{
		store:   store,
		router:  *http.NewServeMux(),
		session: session,
		hub:     NewHub(),
	}
}

// Conf - Hanlders to routes
func (s *Server) Conf() {
	go s.hub.messageListener()
	s.router.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("../client/js"))))
	s.router.HandleFunc("/", s.Index)

	s.router.HandleFunc("/api/", s.MainPage)
	s.router.HandleFunc("/api/signin", s.SignIn)
	s.router.HandleFunc("/api/signup", s.SignUp)
	s.router.HandleFunc("/api/newpost", s.middleWare(s.CreatePost))
	s.router.HandleFunc("/api/post", s.GetPost)
	s.router.HandleFunc("/api/logout", s.middleWare(s.LogOut))
	s.router.HandleFunc("/api/chat", func(w http.ResponseWriter, r *http.Request) {
		s.WSChat(s.hub, w, r)
	})
}

func (h *Hub) messageListener() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

// ListenAndServe - Listener with Configurations to ServMUX
func (s *Server) ListenAndServe(port string) {
	s.Conf()
	log.Println("Server is listening" + port)
	http.ListenAndServe(port, &s.router)
}

// Parser - to parse indexhtml
func Parser() *template.Template {
	temp, err := template.ParseFiles("../client/index.html")
	if err != nil {
		log.Println(fmt.Errorf("ParseFiles: %w", err))
	}
	return temp
}

// logger - send Notification to Client & log error
func logger(w http.ResponseWriter, status int, inputErr error) {
	if newErr.CheckErr(inputErr) {
		bytes, err := json.Marshal(inputErr.Error())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.WriteHeader(status)
		w.Write(bytes)
		log.Println(inputErr)
		return
	}
	w.WriteHeader(status)
	log.Println(inputErr)
}

// logout set cookies max age to -1
func logout(w http.ResponseWriter, ck *http.Cookie) {
	ck.MaxAge = -1
	http.SetCookie(w, ck)
	w.WriteHeader(http.StatusOK)
}

// success - send Notification to Client about success
func success(w http.ResponseWriter, notify string) {
	bytes, err := json.Marshal(notify)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(bytes)
}
