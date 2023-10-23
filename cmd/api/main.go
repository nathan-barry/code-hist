package main

import (
	"github.com/nathan-barry/code-hist/pkg/api"
	. "github.com/nathan-barry/code-hist/pkg/types"
)

func main() {
	repoName := "nathan-barry/nathan.rs"
	url := "https://api.github.com/repos/" + repoName + "/commits"

	var commits []Commit
	api.GetJSON(url, &commits)
	for _, c := range commits {
		// Prints SHA
		api.PrintJSON(c)
	}

	// Latest commit
	url = "https://api.github.com/repos/" + repoName + "/commits/" + commits[0].SHA

	var commitFiles Files
	api.PrintJSON(commitFiles)
}
