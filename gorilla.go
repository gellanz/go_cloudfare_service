package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
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

	// template
	type Todo struct {
		Title string
		Done  bool
	}

	type TodoPageData struct {
		PageTitle string
		Todos     []Todo
		Style     template.CSS
	}

	style, err := os.ReadFile("template/style.css")

	if err != nil {
		return
	}

	page := TodoPageData{
		PageTitle: "TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
		Style: template.CSS(style),
	}

	r.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("template/layout.html")

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, page)
	})

	http.ListenAndServe(":80", r)
}
