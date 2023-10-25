package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nathan-barry/pretty-commit/pkg/api"
)

func FetchCodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pinged -> FetchCode")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	githubKey := os.Getenv("GITHUB_AUTH")

	url := r.FormValue("raw_url")
	fileName := r.FormValue("file_name")
	body := api.GetBody(url, githubKey)

	t := template.Must(template.ParseFiles("./views/home/code.html"))

	data := map[string]any{
		"Code":     string(body),
		"FileName": fileName,
	}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Template error:", err) // Log the error
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
