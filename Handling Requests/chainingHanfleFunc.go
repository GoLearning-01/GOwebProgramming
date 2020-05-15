package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a Hello handle func!")
}

// Chaining part, take and retuen handlefunc
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Don't know what this part does yet, but understand the logic
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

// This way, you can chain more handlefuncs

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/hello", log(helloHandler))

	server.ListenAndServe()

}
