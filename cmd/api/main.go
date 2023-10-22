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
	repoName := "nathan-barry/nathan.rs"
	url := "https://api.github.com/repos/" + repoName + "/commits"

	// var data any
	// getJSON(url, &data)

	// prettyJSON, _ := json.MarshalIndent(data, "", "    ")
	// fmt.Println(string(prettyJSON))

	var commits []Commit
	getJSON(url, &commits)

	for _, c := range commits {
		fmt.Println(c.String())
	}

	url = "https://api.github.com/repos/" + repoName + "/commits/" + commits[0].SHA
	var commitFiles Files
	getJSON(url, &commitFiles)

	prettyJSON, _ := json.MarshalIndent(commitFiles, "", "    ")
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

type Files struct {
	Files []File `json:"files"`
}

type File struct {
	FileName    string `json:"filename"`
	Changes     int    `json:"changes"`
	Additions   int    `json:"additions"`
	Deletions   int    `json:"deletions"`
	BlobURL     string `json:"blob_url"` // Link to github
	ContentsURL string `json:"contents_url"`
	RawURL      string `json:"raw_url"` // Contains file in resp.Body
	Patch       string `json:"patch"`
	SHA         string `json:"sha"`
	Status      string `json:"status"`
}
