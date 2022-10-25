package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	r := Router{}

	// TODO: Add logging, time measurement middleware
	// link: https://drstearns.github.io/tutorials/gomiddleware/
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Router struct{}

// TODO: Replace routing with gorilla/mux
// Bonus: Write your own router (https://benhoyt.com/writings/go-routing/, https://www.alexedwards.net/blog/which-go-router-should-i-use)
func (ro Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1:]
	fmt.Println(path, len(path))
	if len(path) < 1 || path[0] != "api" || path[1] != "simpleapi" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		repo, err := NewSimpleRepository()
		if len(path) == 2 {
			if r.Method == http.MethodPost {
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				c := CreateSimpleHandler{
					repo: repo,
				}
				c.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else if len(path) == 3 {
			if r.Method == http.MethodGet {
				g := GetSimpleHandler{
					repository: repo,
				}
				g.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}

}

type MyHandler struct {
}

func (m MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
	fmt.Println(strings.Split(r.URL.Path, "/"))
	w.Write([]byte("\nhello"))
}
