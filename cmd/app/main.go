package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nathan-barry/code-hist/pkg/handlers"
)

func main() {
	fmt.Println("Starting server...")
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/test", handlers.TestHandler)
	http.HandleFunc("/fetch-repo", handlers.FetchRepoHandler)
	http.HandleFunc("/fetch-files", handlers.FetchFilesHandler)
	http.HandleFunc("/fetch-code", handlers.FetchCodeHandler)
	http.HandleFunc("/testauth", handlers.TestAuthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
