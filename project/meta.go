package project

// MetaData references to other information
type MetaData struct {
	Project      ProjectInfo  `json:"project"`
	BaseTemplate BaseTemplate `json:"basetemplate"`
}
