package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Commit struct {
	SHA string `json:"sha"`
}

func (c Commit) String() string {
	return c.SHA
}

func main() {
	// Create URL
	repoName := "nathan-barry/code-hist"
	url := "https://api.github.com/repos/" + repoName + "/commits"
	fmt.Println("URL:", url)

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
		return
	}

	// Turn body into JSON
	var data interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	// prettyJSON, _ := json.MarshalIndent(data, "", "    ")

	var commits []Commit
	if err = json.Unmarshal(body, &commits); err != nil {
		log.Fatal(err)
	}

	for _, c := range commits {
		fmt.Println(c.String())
	}
}
