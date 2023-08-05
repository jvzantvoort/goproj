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
    fmt.Printf("project type config dir: %s", mainconfig.ConfigDir)

Package config provides configuration data globally used

## Usage

```go
const (
	SettingsFile string = "settings.json"
	RegistryFile string = "registry.json"
)
```

#### type MainConfig

```go
type MainConfig struct {
	ForceInit      bool
	HomeDir        string
	ArchiveDir     string
	ConfigDir      string
	TemplatesDir   string
	ConfigDirPerms int
	AppVersion     string     `json:"version"`
	UserConfig     UserConfig `json:"user"`
}
```

MainConfig configuration for goproj

#### func  NewMainConfig

```go
func NewMainConfig() *MainConfig
```
NewMainConfig initialize a MainConfig and initialize it.

    mc := config.NewMainConfig()
    fmt.Printf("config dir: %s\n", mc.ConfigDir)

#### func (MainConfig) ConfigFile

```go
func (m MainConfig) ConfigFile(name string) string
```

#### func (MainConfig) CreateDirs

```go
func (m MainConfig) CreateDirs()
```
CreateDirs create the main config dir

#### func (MainConfig) Fields

```go
func (m MainConfig) Fields() []string
```

#### func (*MainConfig) Get

```go
func (m *MainConfig) Get(name string) (string, error)
```

#### func (*MainConfig) GetArchiveDir

```go
func (m *MainConfig) GetArchiveDir() string
```

#### func (*MainConfig) GetConfigDir

```go
func (m *MainConfig) GetConfigDir() string
```

#### func (*MainConfig) GetHomeDir

```go
func (m *MainConfig) GetHomeDir() string
```
GetHomeDir get the user's homedir

#### func (*MainConfig) GetTemplatesDir

```go
func (m *MainConfig) GetTemplatesDir() string
```

#### func (*MainConfig) Init

```go
func (m *MainConfig) Init()
```
Init initialize the MainConfig struct

#### func (*MainConfig) Read

```go
func (m *MainConfig) Read(reader io.Reader) error
```
File handling

Read

#### func (*MainConfig) ReadFromFile

```go
func (m *MainConfig) ReadFromFile(name string) error
```
ReadFromFile

#### func (*MainConfig) ResetConfig

```go
func (m *MainConfig) ResetConfig()
```

#### func (MainConfig) Save

```go
func (m MainConfig) Save()
```

#### func (MainConfig) Write

```go
func (m MainConfig) Write(writer io.Writer) error
```
Write

#### func (MainConfig) WriteToFile

```go
func (m MainConfig) WriteToFile(name string) error
```
WriteToFile

#### type UserConfig

```go
type UserConfig struct {
	MailAddress string `json:"mailaddress"`
	Company     string `json:"company"`
	Copyright   string `json:"copyright"`
	License     string `json:"license"`
	User        string `json:"user"`
	Username    string `json:"username"`
}
```
