package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Server started on port 3000")
	fmt.Println(r)

	CreateBook := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "Book Title: %s\n", title)
	}

	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")

	// r.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	page := vars["Page"]
	// 	fmt.Fprintf(w, "Books Page: %s\n", page)
	// })

	r.HandleFunc("/title", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "Title: %s\n", title)
	})

	http.ListenAndServe(":3000", r)
}
