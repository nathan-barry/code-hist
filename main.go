package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello world!"))
	fmt.Println("Home served")

}

func main() {
	fmt.Println("Starting Server...")
	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
