package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nathan-barry/pretty-commit/api"
)

func FetchCodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pinged -> FetchCode")

	url := r.FormValue("raw_url")
	fileName := r.FormValue("file_name")
	body := api.GetBody(url, githubKey)

	t := template.Must(template.ParseFiles("./views/home/code.html"))

	data := map[string]any{
		"Code":     string(body),
		"FileName": fileName,
	}

	err := t.Execute(w, data)
	if err != nil {
		fmt.Println("Template error:", err) // Log the error
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
