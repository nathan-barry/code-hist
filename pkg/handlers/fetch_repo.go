package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nathan-barry/code-hist/pkg/api"
	. "github.com/nathan-barry/code-hist/pkg/types"
)

func FetchRepoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pinged -> FetchRepo")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	githubKey := os.Getenv("GITHUB_AUTH")

	url := "https://api.github.com/repos/" + r.FormValue("repoURL") + "/commits"
	var rawCommits []RawCommit
	fmt.Println("HERE 1")
	api.GetJSON(url, &rawCommits, githubKey)
	fmt.Println("HERE 2")
	api.PrintJSON("rawCommits[0]", rawCommits[0])

	t := template.Must(template.ParseFiles("./views/home/commits.html"))

	data := map[string]any{
		"RawCommits": rawCommits,
	}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Template error:", err) // Log the error
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
