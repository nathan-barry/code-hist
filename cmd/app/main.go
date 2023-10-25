package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nathan-barry/pretty-commit/pkg/handlers"
)

func main() {
	fmt.Println("Starting server...")
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/fetch-repo", handlers.FetchRepoHandler)
	http.HandleFunc("/fetch-files", handlers.FetchFilesHandler)
	http.HandleFunc("/fetch-code", handlers.FetchCodeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
