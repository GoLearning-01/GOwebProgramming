package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", counter)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func counter(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}

	count++
	c.Value = strconv.Itoa(count)

	http.SetCookie(w, c)

	io.WriteString(w, c.Value)

}
