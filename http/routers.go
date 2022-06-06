package http

import (
	"net/http"

	"github.com/madxiii/real-time-forum/service"
	"github.com/madxiii/real-time-forum/session"
)

type API struct {
	service *service.Service
	session *session.Sesssion
}

func NewAPI(service *service.Service) *API {
	return &API{service: service}
}

func (a *API) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("../client/js"))))
	mux.HandleFunc("/", a.Index)

	return mux
}
