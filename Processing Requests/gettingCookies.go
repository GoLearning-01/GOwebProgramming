package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "First Cookie",
		Value:    "Go Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "Second Cookie",
		Value:    "Mason's cookie",
		HttpOnly: true,
	}

	// This can also be used
	/*
		w.Header().Set("Set-Cookie", c1.String())
		w.Header().Set("Set-Cookie", c2.String())
	*/

	// But this is easier
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/setcookie", setCookie)
	http.HandleFunc("/getcookie", getCookie)

	server.ListenAndServe()
}
