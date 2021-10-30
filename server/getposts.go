package server

import (
	"net/http"
)

func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.handlerGetPostPage(w)
		return
	}
	if r.Method == http.MethodPost {
		s.handlerGetPost(w, r)
		return
	}

}

func (s *Server) handlerGetPostPage(w http.ResponseWriter) {
	temp := Parser()
	if err := temp.Execute(w, nil); err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}
}

func (s *Server) handlerGetPost(w http.ResponseWriter, r *http.Request) {

}
