package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type GetSimpleHandler struct {
	repository SimpleRepository
}

func (gsh GetSimpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := strings.Split(r.URL.Path, "/")[1:]
	id := path[2]
	simple, err := gsh.repository.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(simple)
}
