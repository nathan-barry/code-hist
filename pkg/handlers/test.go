package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nathan-barry/code-hist/pkg/api"
	. "github.com/nathan-barry/code-hist/pkg/types"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Pinged -> Test")
	repoName := "nathan-barry/nathan.rs"
	url := "https://api.github.com/repos/" + repoName + "/commits"

	var commits []Commit
	api.GetJSON(url, &commits)
	for i, c := range commits {
		// Prints SHA
		api.PrintJSON(c)
		if i > 2 {
			break
		}
	}

	// Latest commit
	url = "https://api.github.com/repos/" + repoName + "/commits/" + commits[0].SHA

	var commitFiles Files
	api.GetJSON(url, &commitFiles)
	// api.PrintJSON(commitFiles.Files[0])

	// fmt.Print(string(api.GetBody(commitFiles.Files[0].RawURL)))

	t := template.Must(template.ParseFiles("./templates/partials/base.html", "./templates/test/index.html"))

	err := t.ExecuteTemplate(w, "base.html", url)
	if err != nil {
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}
