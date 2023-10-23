package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pinged -> Home")
	t := template.Must(template.ParseFiles("./templates/partials/base.html", "./templates/home/index.html"))

	err := t.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
