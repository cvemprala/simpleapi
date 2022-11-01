package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := mux.NewRouter()
	r = CreateSimpleapiRouter(r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func CreateSimpleapiRouter(r *mux.Router) *mux.Router {
	repo, err := NewSimpleRepository()

	if err != nil {
		return nil
	}
	createSimpleHandler := CreateSimpleHandler{
		repo: repo,
	}

	getSimpleHandler := GetSimpleHandler{
		repository: repo,
	}

	apiRouter := r.PathPrefix("/api/simpleapi").Subrouter()

	apiRouter.Handle("", NewLogger(createSimpleHandler)).Methods("POST")
	apiRouter.Handle("/{id}", NewLogger(getSimpleHandler)).Methods("GET")
	return r

}
