package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	uname string
	fname string
	lname string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // UserID, User
var dbSessions = map[string]string{} // Session ID, UserID

func init() {
	tpl = template.Must(template.ParseGlob("*gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {

	// Get Cookie
	cookie, err := r.Cookie("session-cookie")
	if err != nil {
		sID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-cookie",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
	}

	// If the user already exist, get user
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	// Process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		fn := r.FormValue("firstname")
		ln := r.FormValue("lastname")
		u = user{un, fn, ln}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func bar(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-cookie")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
