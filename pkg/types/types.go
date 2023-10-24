package types

type RawCommit struct {
	SHA    string `json:"sha"` // sha to file
	Commit struct {
		Committer struct {
			Name string `json:"name"`
			Date string `json:"date"`
		}
		Message string `json:"message"`
	}
}

type Files struct {
	Files []File `json:"files"`
}

type File struct {
	FileName    string `json:"filename"`
	Changes     int    `json:"changes"`
	Additions   int    `json:"additions"`
	Deletions   int    `json:"deletions"`
	BlobURL     string `json:"blob_url"` // Link to github
	ContentsURL string `json:"contents_url"`
	RawURL      string `json:"raw_url"` // Contains entire file in resp.Body
	Patch       string `json:"patch"`
	SHA         string `json:"sha"`
	Status      string `json:"status"`
}
