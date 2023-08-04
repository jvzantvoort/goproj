# project
--
    import "github.com/jvzantvoort/goproj/project"

### Project

    MetaData
      ProjectInfo
        Name
        Description
      Base Template
        Name
        Version
    Locations
      RootDir
        ConfigDir
    Resources
      Environment
        PATH
        LD_LIBRARY_PATH
    Functions
      Setup
      Status
      Backup
      Resume
      Package
      Test
    Session Management

## Usage

```go
const (
	ProjectName string = "goproj"
)
```

```go
var (
	// ErrFileNotFound project not found
	ErrFileNotFound = errors.New("File not found")
	// ErrNotFound project not found
	ErrNotFound = errors.New("Project not found")
	// ErrDuplicate project is already in list when it should not
	ErrDuplicate   = errors.New("Project already found")
	ErrListEmpty   = errors.New("List is empty")
	ErrListTooLong = errors.New("List is too long")
)
```

#### func  Buffer2Slice

```go
func Buffer2Slice(stream io.ReadCloser) []string
```

#### func  ExecNonFatal

```go
func ExecNonFatal(f fn, args ...string)
```

#### func  FileExists

```go
func FileExists(fpath string) (bool, os.FileInfo)
```

#### func  FileIsExecutable

```go
func FileIsExecutable(fpath string) bool
```

#### func  MkdirAll

```go
func MkdirAll(path string, mode int)
```

#### func  OneOrLess

```go
func OneOrLess(args ...string) (string, error)
```

#### func  PanicOnError

```go
func PanicOnError(fmtstr string, err error)
```

#### func  PrintError

```go
func PrintError(fmtstr string, err error) error
```

#### func  PrintFatal

```go
func PrintFatal(fmtstr string, err error) error
```

#### func  Reverse

```go
func Reverse(input []string) []string
```

#### type BaseTemplate

```go
type BaseTemplate struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
```

BaseTemplate template the project was originally based on

#### type EnvVar

```go
type EnvVar struct {
	RootPath string   `json:"-"`
	Name     string   `json:"name"`
	Paths    []string `json:"paths"`
}
```


#### func (*EnvVar) Append

```go
func (evar *EnvVar) Append(inputdir string)
```

#### func (EnvVar) Has

```go
func (evar EnvVar) Has(inputdir string) bool
```

#### func (*EnvVar) Prepend

```go
func (evar *EnvVar) Prepend(inputdir string)
```

#### func (EnvVar) Write

```go
func (evar EnvVar) Write(w io.Writer)
```
ExportString write the name and value to the buffer as a bash export string

#### type Environment

```go
type Environment struct {
	RootDir string            `json:"-"`
	Vars    map[string]EnvVar `json:"vars"`
}
```


#### type Executable

```go
type Executable struct {
	PathToScript string
	OS           string
	Extensions   []string
}
```


#### func  NewExecutable

```go
func NewExecutable(command string, extensions ...string) *Executable
```

#### func (*Executable) AppendExt

```go
func (e *Executable) AppendExt(extstr string)
```

#### func (Executable) LookupExt

```go
func (e Executable) LookupExt() (string, error)
```

#### func (*Executable) PrependExt

```go
func (e *Executable) PrependExt(extstr string)
```

#### type Functions

```go
type Functions struct {
	Cwd       string
	ToolsPath string
	Locations Locations
}
```


#### func  NewFunctions

```go
func NewFunctions(locations Locations) *Functions
```

#### func (Functions) Backup

```go
func (f Functions) Backup(args ...string) ([]string, []string, error)
```

#### func (Functions) Build

```go
func (f Functions) Build(args ...string) ([]string, []string, error)
```

#### func (Functions) BuildProject

```go
func (f Functions) BuildProject(args ...string)
```

#### func (Functions) Execute

```go
func (f Functions) Execute(command string, args ...string) ([]string, []string, error)
```

#### func (Functions) LookupExt

```go
func (f Functions) LookupExt(command string) (string, error)
```

#### func (Functions) Package

```go
func (f Functions) Package(args ...string) ([]string, []string, error)
```

#### func (Functions) Publish

```go
func (f Functions) Publish(args ...string) ([]string, []string, error)
```

#### func (Functions) Setup

```go
func (f Functions) Setup(args ...string) ([]string, []string, error)
```

#### func (Functions) SetupProject

```go
func (f Functions) SetupProject(args ...string)
```

#### func (Functions) Status

```go
func (f Functions) Status(args ...string) ([]string, []string, error)
```

#### func (Functions) Teardown

```go
func (f Functions) Teardown(args ...string) ([]string, []string, error)
```

#### func (Functions) Test

```go
func (f Functions) Test(args ...string) ([]string, []string, error)
```

#### func (Functions) Which

```go
func (f Functions) Which(command string) string
```

#### type Locations

```go
type Locations struct {
	RootDir string `json:"root"`
}
```

Locations the locations used in the object

#### func (Locations) BackupDir

```go
func (L Locations) BackupDir(args ...string) string
```

#### func (Locations) BackupDirRotating

```go
func (L Locations) BackupDirRotating(name string, max int) string
```

#### func (Locations) BinPath

```go
func (L Locations) BinPath() string
```

#### func (Locations) ConfigDir

```go
func (L Locations) ConfigDir() string
```

#### func (Locations) ConfigFile

```go
func (L Locations) ConfigFile() string
```

#### func (Locations) ToolsPath

```go
func (L Locations) ToolsPath() string
```

#### type MetaData

```go
type MetaData struct {
	Project      ProjectInfo  `json:"project"`
	BaseTemplate BaseTemplate `json:"basetemplate"`
}
```

MetaData references to other information

#### type Project

```go
type Project struct {
	MetaData  MetaData  `json:"metadata"`
	Locations Locations `json:"locations"`
	Functions Functions `json:"-"`
}
```

Project the project object

    proj := NewProject("/home/foo/project")

#### func  NewProject

```go
func NewProject(path string) *Project
```

#### func (*Project) Description

```go
func (p *Project) Description(args ...string) string
```

#### func (*Project) Name

```go
func (p *Project) Name(args ...string) string
```

#### func (*Project) Read

```go
func (p *Project) Read(reader io.Reader) error
```
File handling

Read

#### func (*Project) ReadFromFile

```go
func (p *Project) ReadFromFile() error
```
ReadFromFile

#### func (Project) Write

```go
func (p Project) Write(writer io.Writer) error
```
Write

#### func (Project) WriteToFile

```go
func (p Project) WriteToFile() error
```
WriteToFile

#### type ProjectInfo

```go
type ProjectInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
```

ProjectInfo basic project info

#### type Targets

```go
type Targets struct {
	Files []os.FileInfo
	Repos []VCSUrl
}
```


#### type VCSUrl

```go
type VCSUrl struct {
	Url         string `json:"url"`
	Type        string `json:"type"`
	Branch      string `json:"branch"`
	Destination string `json:"destination"`
}
```

Version Constrol Service Url
