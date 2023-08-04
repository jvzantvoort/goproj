# config
--
    import "github.com/jvzantvoort/goproj/config"

Example:

    import (
      "fmt"
      "github.com/jvzantvoort/goproj/config"
    )

    mainconfig := NewMainConfig()
    fmt.Printf("home dir: %s", mainconfig.HomeDir)
    fmt.Printf("tmux dir: %s", mainconfig.TmuxDir)
    fmt.Printf("project type config dir: %s", mainconfig.ProjTypeConfigDir)

Package config provides configuration data globally used

## Usage

#### type MainConfig

```go
type MainConfig struct {
	HomeDir            string
	TmuxDir            string
	ProjTypeConfigDir  string
	ProjTypeConfigMode int
}
```


#### func  NewMainConfig

```go
func NewMainConfig() *MainConfig
```

#### func (MainConfig) ExpandHome

```go
func (m MainConfig) ExpandHome(pathstr string) (string, error)
```
ExpandHome expand the tilde in a given path.

#### func (*MainConfig) GetHomeDir

```go
func (m *MainConfig) GetHomeDir() string
```

#### func (MainConfig) GetProjTypeConfigDir

```go
func (m MainConfig) GetProjTypeConfigDir() (string, int)
```

#### func (*MainConfig) GetTmuxDir

```go
func (m *MainConfig) GetTmuxDir() string
```

#### func (MainConfig) MkdirAll

```go
func (m MainConfig) MkdirAll(path string, mode int)
```

#### func (MainConfig) Prefix

```go
func (m MainConfig) Prefix() string
```
