package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Takes in the URL and an optional github api key
func GetBody(url string, key string) []byte {
	fmt.Println("\nGetting from URL:", url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// If key isn't empty, add it to header
	if key != "" {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", key))
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("STATUS CODE:", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func GetJSON(url string, data any, key string) {
	body := GetBody(url, key)
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

func PrintJSON(name string, data any) {
	fmt.Println("PRINT_JSON: " + name + ":\n" + string(PrettyJSON(data)))
}
