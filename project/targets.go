package project

import "os"

// VCSUrl missing godoc.
// Version Constrol Service Url
type VCSUrl struct {
	Url         string `json:"url"`
	Type        string `json:"type"`
	Branch      string `json:"branch"`
	Destination string `json:"destination"`
}

// Targets missing godoc.
type Targets struct {
	Files []os.FileInfo
	Repos []VCSUrl
}
