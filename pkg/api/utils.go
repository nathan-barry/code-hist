package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetBody(url string) []byte {
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

func GetJSON(url string, data any) {
	body := GetBody(url)
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
}

func PrintJSON(data any) {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(prettyJSON))
}
