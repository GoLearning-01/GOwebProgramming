package main

import (
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

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/cookies", setCookie)

	server.ListenAndServe()
}
