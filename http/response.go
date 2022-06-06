package http

import (
	"encoding/json"
	"log"
	"net/http"

	newErr "github.com/madxiii/real-time-forum/error"
)

func logger(w http.ResponseWriter, status int, inputErr error) {
	w.WriteHeader(status)
	log.Println(inputErr)
	if newErr.CheckErr(inputErr) {
		bytes, err := json.Marshal(inputErr.Error())
		if err != nil {
			log.Printf(`logger, json.Marshal with: %s, err: %s`, inputErr.Error(), err.Error())
			return
		}
		w.Write(bytes)
	}
}

func success(w http.ResponseWriter, notify string) {
	bytes, err := json.Marshal(notify)
	if err != nil {
		logger(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(bytes)
}
