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
		fmt.Println("Get JSON Fucked up", err)
		log.Fatal(err)
	}
}

func PrettyJSON(data any) []byte {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("PrettyJSON Fucked up")
		log.Fatal(err)
	}
	return prettyJSON
}

func PrintJSON(data any) {
	fmt.Println(string(PrettyJSON(data)))
}
