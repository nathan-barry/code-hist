package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nathan-barry/pretty-commit/api"
	. "github.com/nathan-barry/pretty-commit/types"
)

func FetchFilesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pinged -> FetchFiles")

	url := r.FormValue("url") + "/" + r.FormValue("sha")

	var commitFiles Files
	api.GetJSON(url, &commitFiles, githubKey)

	t := template.Must(template.ParseFiles("./views/home/files.html"))

	data := map[string]any{
		"FileArray": commitFiles.Files,
	}

	err := t.Execute(w, data)
	if err != nil {
		fmt.Println("Template error:", err) // Log the error
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
