package main

import (
	"CloudServer/controller"
	"fmt"
	"net/http"
)

var Route map[string]

type Mux struct {
}

func (this *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		Demo(w, r)
	case "/user":

	default:
		NotFound(w, r)
	}
}

func Demo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func NotFound(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := &Mux{}
	http.ListenAndServe(":9000", mux)
}
