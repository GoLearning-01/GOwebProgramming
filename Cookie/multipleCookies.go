package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/multiple", multiple)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Cookie-One",
		Value: "First Cookie",
	})
	fmt.Fprintln(w, "First cookie written successfully!")
}

func read(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("Cookie-One")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "First cookie is:", c1)
	}

	c2, err := r.Cookie("Cookie-Two")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Second cookie is:", c2)
	}

	c3, err := r.Cookie("Cookie-Three")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Third cookie is:", c3)
	}
}

func multiple(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Cookie-Two",
		Value: "Second Cookie",
	})
	fmt.Fprintln(w, "Second cookie written sucessfully!")

	http.SetCookie(w, &http.Cookie{
		Name:  "Cookie-Three",
		Value: "Third Cookie",
	})
	fmt.Fprintln(w, "Third cookie written sucessfully!")
}
