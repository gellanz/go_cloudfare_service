package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world, you've requested: %s \n", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
