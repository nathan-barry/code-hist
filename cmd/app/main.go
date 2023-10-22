package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseGlob("../../templates/*/*.html"))

	data := map[string]interface{}{
		"Title": "Home Page",
	}

	err := template.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
