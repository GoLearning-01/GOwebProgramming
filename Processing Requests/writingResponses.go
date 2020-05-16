package main

import (
	"fmt"
	"net/http"
)

// Writing regular responses
func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
			<head>
			<title>Go Web Programming</title>
			</head> 
			<body><h1>Hello World</h1></body>
			</html>`
	w.Write([]byte(str))
}

// Usually for errors
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "Content not here bro! Visit other pages...")
}

// For redirecting
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://summitkhatiwada.com")
	w.WriteHeader(302)

}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)

	server.ListenAndServe()
}
