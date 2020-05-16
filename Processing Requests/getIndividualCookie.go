package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "First_cookie",
		Value:    "Do you like cookie?",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "Second_cookie",
		Value:    "Yes I do!",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("First_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get first cookie!")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)

	server.ListenAndServe()
}
