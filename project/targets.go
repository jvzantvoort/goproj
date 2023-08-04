package project

import "os"

// Version Constrol Service Url
type VCSUrl struct {
	Url         string `json:"url"`
	Type        string `json:"type"`
	Branch      string `json:"branch"`
	Destination string `json:"destination"`
}

type Targets struct {
	Files []os.FileInfo
	Repos []VCSUrl
}
