package main

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gorilla/mux"
)

func main() {
	// new router
	r := mux.NewRouter()

	r.HandleFunc("/videogames/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "%s \n", title)
		videogames := []string{"dark_souls", "pokemon"}
		if slices.Contains(videogames, title) {
			fmt.Fprintf(w, "Nice game!")
		} else {
			fmt.Fprintf(w, "I don't know about it")
		}
	})

	http.ListenAndServe(":80", r)
}
