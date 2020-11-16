package main

import (
	"fmt"
	"html"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(res, "Hello, my name is Anton!")
	fmt.Fprintf(res, "Hello, %q", html.EscapeString(req.URL.Path))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}
