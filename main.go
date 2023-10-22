package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Create URL
	repoName := "nathan-barry/code-hist"
	url := "https://api.github.com/repos/" + repoName + "/commits"

	var data any
	getJSON(url, &data)

	// prettyJSON, _ := json.MarshalIndent(data, "", "    ")
	// fmt.Println(string(prettyJSON))

	var commits []Commit
	getJSON(url, &commits)

	for _, c := range commits {
		fmt.Println(c.String())
	}

	url = "https://api.github.com/repos/" + repoName + "/commits/" + commits[0].SHA
	var commitData any
	getJSON(url, &commitData)

	prettyJSON, _ := json.MarshalIndent(commitData, "", "    ")
	fmt.Println(string(prettyJSON))

}

type Commit struct {
	SHA string `json:"sha"`
}

func (c Commit) String() string {
	return c.SHA
}

func getBody(url string) []byte {
	fmt.Println("Getting from URL:", url)
	// Get response from URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read body of response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func getJSON(url string, data any) {
	body := getBody(url)
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
}
