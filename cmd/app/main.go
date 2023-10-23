package main

import (
	"fmt"
	"net/http"

	"github.com/nathan-barry/code-hist/pkg/handlers"
)

func main() {
	fmt.Println("Starting server...")

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/test", handlers.TestHandler)
	http.ListenAndServe(":8080", nil)
}
