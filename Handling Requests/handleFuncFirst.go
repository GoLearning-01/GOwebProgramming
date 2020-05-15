package main

import (
	"fmt"
	"net/http"
)

type firstPageHandler struct{}

func (h *firstPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the first page!")
}

type secondPageHandler struct{}

func (h *secondPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the second page!")
}

func main() {

	firstpagehandler := firstPageHandler{}
	secondpagehandler := secondPageHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/first", &firstpagehandler)
	http.Handle("/second", &secondpagehandler)

	server.ListenAndServe()

}
