package project

import "path"

// Locations the locations used in the object
type Locations struct {
	RootDir string
}

func (L Locations) ConfigDir() string {
	return path.Join(L.RootDir, "."+ProjectName)
}

func (L Locations) BinPath() string {
	return path.Join(L.RootDir, "bin")

}

func (L Locations) ToolsPath() string {
	return path.Join(L.ConfigDir(), "libexec")

}
