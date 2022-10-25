package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Simple struct {
	ID       string
	Name     string
	Birthday time.Time
	Phone    string
	Email    string
}

type CreateSimpleHandler struct {
	repo SimpleRepository
}

// TODO: Add validation using the validator.v9 package
func (c CreateSimpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var request Simple
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := c.repo.Create(request)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))

}
