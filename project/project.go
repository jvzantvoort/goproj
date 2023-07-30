package project


// Project the project object
//
//	proj := NewProject("/home/foo/project")
type Project struct {
	MetaData  MetaData  `json:"metadata"`
	Locations Locations `json:"-"`
	Functions Functions `json:"-"`
}

func NewProject(path string) *Project {
	retv := &Project{}
	retv.Locations.RootDir = path
	fn := NewFunctions(retv.Locations)
	retv.Functions = *fn
	retv.Functions.Setup()
	return retv
}
