package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2002-01-01"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/time", process)

	server.ListenAndServe()
}
