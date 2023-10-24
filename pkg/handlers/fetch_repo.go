package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func FetchRepoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pinged -> FetchRepo")
	fmt.Println(r.Method)

	t := template.Must(template.ParseFiles("./views/home/commits.html"))

	data := map[string]any{
		"Method": r.Method,
	}

	err := t.Execute(w, data)
	if err != nil {
		fmt.Println("Template error:", err) // Log the error
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
