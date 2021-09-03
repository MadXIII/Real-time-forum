package server

import (
	"encoding/json"
	"forum/models"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type PageData struct {
	PageTitle string
	Data      interface{}
}

func Parse() (*template.Template, error) {
	temp, err := template.ParseFiles("../client/index.html")
	if err != nil {
		log.Fatal(err)
	}
	return temp, err
}

// func Execute(temp *template.Template) {
// 	temp.Execute()
// }

// func (s *Server) Registration(w http.ResponseWriter, r *http.Request) {
// 	var newUser models.User
// 	if r.Method == http.MethodGet {
// 		w.Write([]byte("TEST"))
// 	}
// 	if r.Method == "POST" {
// 		bytes, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		err = json.Unmarshal(bytes, &newUser) //
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		err = s.store.InsertUser(newUser)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// }
func (s *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	log.Println("post huest")

}

func (s *Server) MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// if r.URL.Path != "/" {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	w.Write([]byte("404 not found"))
		// 	return
		// }
		temp, err := Parse()
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		err = temp.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 Method not allowed"))
	}
}

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if r.Method == http.MethodPost {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(bytes, &newUser)
		if err != nil {
			log.Fatal(err)
		}

		err = s.store.InsertUser(newUser)
		if err != nil {
			log.Fatal(err)
		}
	}
}
