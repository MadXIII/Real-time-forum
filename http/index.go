package http

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	newErr "github.com/madxiii/real-time-forum/error"
)

func (a *API) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := checkURLPath(r.URL.Path); err != nil {
			logger(w, http.StatusNotFound, err)
			return
		}
		temp := parser()
		if err := temp.Execute(w, nil); err != nil {
			logger(w, http.StatusInternalServerError, fmt.Errorf("Index, Execute: %w", err))
			return
		}
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func checkURLPath(path string) error {
	if path != "/" && path != "/signin" && path != "/signup" && path != "/newpost" && path != "/post" && path != "/logout" {
		return newErr.ErrNotFound
	}
	return nil
}

func parser() *template.Template {
	temp, err := template.ParseFiles("./client/index.html")
	if err != nil {
		log.Println(fmt.Errorf("ParseFiles: %w", err))
	}
	return temp
}
