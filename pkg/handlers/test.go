package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nathan-barry/code-hist/pkg/api"
	. "github.com/nathan-barry/code-hist/pkg/types"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	githubKey := os.Getenv("GITHUB_AUTH")
	fmt.Println(fmt.Sprintf("Bearer %s", githubKey))

	fmt.Println("Pinged -> Test")
	repoName := "nathan-barry/code-hist"
	url := "https://api.github.com/repos/" + repoName + "/commits"
	fmt.Println(url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Authorization header with your access token
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", githubKey))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("STATUS CODE:", resp.StatusCode)

	// Handle the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	var commits []Commit
	if err := json.Unmarshal(body, &commits); err != nil {
		fmt.Println("Get JSON Fucked up")
		log.Fatal(err)
	}

	for i, c := range commits {
		// Prints SHA
		api.PrintJSON(c)
		if i > 2 {
			break
		}
	}

	// Latest commit
	url = "https://api.github.com/repos/" + repoName + "/commits/" + commits[0].SHA
	fmt.Println(url)

	var commitFiles Files
	api.GetJSON(url, &commitFiles)
	// api.PrintJSON(commitFiles.Files[0])
	url = commitFiles.Files[0].RawURL
	fileBody := api.GetBody(url)

	// fmt.Print(string(api.GetBody(commitFiles.Files[0].RawURL)))

	t := template.Must(template.ParseFiles("./templates/partials/base.html", "./templates/test/index.html"))

	data := map[string]any{
		"Title":     "Whazzup Bitches 2",
		"Commits":   commits,
		"FileArray": commitFiles.Files,
		"FileBody":  string(fileBody),
	}

	err = t.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		fmt.Println("Template error:", err) // Log the error
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}

}
