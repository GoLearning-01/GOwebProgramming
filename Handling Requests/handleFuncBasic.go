package main

import (
	"fmt"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is first page Dum Dum!")
}

func secondPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is second page Pop Tart!")
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/firstPage", firstPage)
	http.HandleFunc("/secondPage", secondPage)

	server.ListenAndServe()
}
