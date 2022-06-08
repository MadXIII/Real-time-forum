package http

import (
	"net/http"

	"github.com/madxiii/real-time-forum/service"
)

type API struct {
	service *service.Service
}

func NewAPI(service *service.Service) *API {
	return &API{
		service: service,
	}
}

func (a *API) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./client/js"))))
	mux.HandleFunc("/", a.Index)
	mux.HandleFunc("/api/home", a.Home)
	mux.HandleFunc("/api/signup", a.SignUp)
	mux.HandleFunc("/api/signup", a.SignIn)
	mux.HandleFunc("/api/logout", a.Logout)
	mux.HandleFunc("/api/newpost", a.CreatePost)

	return mux
}
