// Viewer
//
// Routines to provide output meant for presentation
//

package config

import (
	"fmt"
	"io"
)

func (m *MainConfig) PrintConfig(writer io.Writer) {
	fmt.Printf("\n")
	fmt.Printf("Homedir:   %s\n", m.HomeDir)
	fmt.Printf("Configdir: %s\n", m.ProjConfigDir)
	fmt.Printf("\n")
	fmt.Printf("Directories:\n")
	fmt.Printf("Cache:     %s\n", m.Main.CacheDir)
	fmt.Printf("Types:     %s\n", m.Main.TypesDir)
	fmt.Printf("\n")
	fmt.Printf("Commands:\n")
	fmt.Printf("Edit:      %s\n", m.CmdPaths.Editor)
	fmt.Printf("Vcs:       %s\n", m.CmdPaths.Vcs)
}
