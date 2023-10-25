package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nathan-barry/pretty-commit/pkg/api"
	. "github.com/nathan-barry/pretty-commit/pkg/types"
)

func FetchFilesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pinged -> FetchFiles")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	githubKey := os.Getenv("GITHUB_AUTH")

	url := r.FormValue("url") + "/" + r.FormValue("sha")

	var commitFiles Files
	api.GetJSON(url, &commitFiles, githubKey)

	t := template.Must(template.ParseFiles("./views/home/files.html"))

	data := map[string]any{
		"FileArray": commitFiles.Files,
	}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Template error:", err) // Log the error
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
