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

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pinged -> Test")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	githubKey := os.Getenv("GITHUB_AUTH")

	// Get Commit History
	repoName := "nathan-barry/code-hist"
	url := "https://api.github.com/repos/" + repoName + "/commits"
	var rawCommits []RawCommit
	api.GetJSON(url, &rawCommits, githubKey)
	api.PrintJSON("rawCommits[0]", rawCommits[0])

	// Get Latest commit info
	url = "https://api.github.com/repos/" + repoName + "/commits/" + rawCommits[0].SHA
	var commitFiles Files
	api.GetJSON(url, &commitFiles, githubKey)
	api.PrintJSON("commitFiles from rawCommits[0].SHA", commitFiles)

	// Gets one raw file
	url = commitFiles.Files[0].RawURL
	fileBody := api.GetBody(url, githubKey)
	api.PrintJSON("commitFiles.Files[0].RawURL", string(fileBody))

	t := template.Must(template.ParseFiles("./views/partials/base.html", "./views/test/index.html"))

	data := map[string]any{
		"Title":     "Whazzup Bitches 2",
		"Commits":   rawCommits,
		"FileArray": commitFiles.Files,
		"FileBody":  string(fileBody),
	}

	err = t.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		fmt.Println("Template error:", err) // Log the error
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}

}
