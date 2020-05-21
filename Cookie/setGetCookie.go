package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "1st cookie",
		Value: "Hi first cookie!",
	})
	fmt.Fprintln(w, "Cookie set!")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("1st cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	}
	fmt.Fprintln(w, "Your cookie:", c)
}

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", getCookie)

	http.ListenAndServe(":8080", nil)
}
