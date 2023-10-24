package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func TestAuthHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	githubKey := os.Getenv("GITHUB_AUTH")
	fmt.Println(fmt.Sprintf("token %s", githubKey))

	fmt.Println("Pinged -> TestAuth")
	repoName := "nathan-barry/code-hist"
	url := "https://api.github.com/repos/" + repoName + "/commits"
	fmt.Println(url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Authorization header with your access token
	req.Header.Set("Authorization", fmt.Sprintf("token %s", githubKey))
	fmt.Println(fmt.Sprintf("token %s", githubKey))
	fmt.Println("REQUEST", req)

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
}
