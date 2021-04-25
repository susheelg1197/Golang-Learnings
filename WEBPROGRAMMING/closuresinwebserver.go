package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/hello", timed(hello)) // This will return an anonymous function with modified changes
	http.ListenAndServe(":3000", nil)
}

// this function will act as a middleware
func timed(f func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()
		fmt.Println(start, end)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<h1>Hello!!</h1>")
}
