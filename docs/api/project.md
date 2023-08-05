# project
--
    import "github.com/jvzantvoort/goproj/project"


## Usage

```go
const (
	ProjectName string = "goproj"
)
```
ProjectName missing godoc.

```go
var (
	// ErrFileNotFound project not found
	ErrFileNotFound = errors.New("File not found")
	// ErrNotFound project not found
	ErrNotFound = errors.New("Project not found")
	// ErrDuplicate project is already in list when it should not
	ErrDuplicate = errors.New("Project already found")
	// ErrListEmpty missing godoc.
	ErrListEmpty = errors.New("List is empty")
	// ErrListTooLong missing godoc.
	ErrListTooLong = errors.New("List is too long")
)
```

#### func  Buffer2Slice

```go
func Buffer2Slice(stream io.ReadCloser) []string
```
Buffer2Slice missing godoc.

#### func  ExecNonFatal

```go
func ExecNonFatal(f fn, args ...string)
```
ExecNonFatal missing godoc.

#### func  FileExists

```go
func FileExists(fpath string) (bool, os.FileInfo)
```
FileExists missing godoc.

#### func  FileIsExecutable

```go
func FileIsExecutable(fpath string) bool
```
FileIsExecutable missing godoc.

#### func  MkdirAll

```go
func MkdirAll(path string, mode int)
```
MkdirAll missing godoc.

#### func  OneOrLess

```go
func OneOrLess(args ...string) (string, error)
```
OneOrLess missing godoc.

#### func  PanicOnError

```go
func PanicOnError(fmtstr string, err error)
```
PanicOnError missing godoc.

#### func  PrintError

```go
func PrintError(fmtstr string, err error) error
```
PrintError missing godoc.

#### func  PrintFatal

```go
func PrintFatal(fmtstr string, err error) error
```
PrintFatal missing godoc.

#### func  Reverse

```go
func Reverse(input []string) []string
```
Reverse missing godoc.

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

EnvVar missing godoc.

#### func (*EnvVar) Append

```go
func (evar *EnvVar) Append(inputdir string)
```
Append missing godoc.

#### func (EnvVar) Has

```go
func (evar EnvVar) Has(inputdir string) bool
```
Has missing godoc.

#### func (*EnvVar) Prepend

```go
func (evar *EnvVar) Prepend(inputdir string)
```
Prepend missing godoc.

#### func (EnvVar) Write

```go
func (evar EnvVar) Write(w io.Writer)
```
Write missing godoc. ExportString write the name and value to the buffer as a
bash export string

#### type Environment

```go
type Environment struct {
	RootDir string            `json:"-"`
	Vars    map[string]EnvVar `json:"vars"`
}
```

Environment missing godoc.

#### type Executable

```go
type Executable struct {
	PathToScript string
	OS           string
	Extensions   []string
}
```

Executable missing godoc.

#### func  NewExecutable

```go
func NewExecutable(command string, extensions ...string) *Executable
```
NewExecutable missing godoc.

#### func (*Executable) AppendExt

```go
func (e *Executable) AppendExt(extstr string)
```
AppendExt missing godoc.

#### func (Executable) LookupExt

```go
func (e Executable) LookupExt() (string, error)
```
LookupExt missing godoc.

#### func (*Executable) PrependExt

```go
func (e *Executable) PrependExt(extstr string)
```
PrependExt missing godoc.

#### type Functions

```go
type Functions struct {
	Cwd       string
	ToolsPath string
	Locations Locations
}
```

Functions missing godoc.

#### func  NewFunctions

```go
func NewFunctions(locations Locations) *Functions
```
NewFunctions missing godoc.

#### func (Functions) Backup

```go
func (f Functions) Backup(args ...string) ([]string, []string, error)
```
Backup missing godoc.

#### func (Functions) Build

```go
func (f Functions) Build(args ...string) ([]string, []string, error)
```
Build missing godoc.

#### func (Functions) BuildProject

```go
func (f Functions) BuildProject(args ...string)
```
BuildProject missing godoc.

#### func (Functions) Execute

```go
func (f Functions) Execute(command string, args ...string) ([]string, []string, error)
```
Execute missing godoc.

#### func (Functions) LookupExt

```go
func (f Functions) LookupExt(command string) (string, error)
```
LookupExt missing godoc.

#### func (Functions) Package

```go
func (f Functions) Package(args ...string) ([]string, []string, error)
```
Package missing godoc.

#### func (Functions) Publish

```go
func (f Functions) Publish(args ...string) ([]string, []string, error)
```
Publish missing godoc.

#### func (Functions) Setup

```go
func (f Functions) Setup(args ...string) ([]string, []string, error)
```
Setup missing godoc.

#### func (Functions) SetupProject

```go
func (f Functions) SetupProject(args ...string)
```
SetupProject missing godoc.

#### func (Functions) Status

```go
func (f Functions) Status(args ...string) ([]string, []string, error)
```
Status missing godoc.

#### func (Functions) Teardown

```go
func (f Functions) Teardown(args ...string) ([]string, []string, error)
```
Teardown missing godoc.

#### func (Functions) Test

```go
func (f Functions) Test(args ...string) ([]string, []string, error)
```
Test missing godoc.

#### func (Functions) Which

```go
func (f Functions) Which(command string) string
```
Which missing godoc.

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
BackupDir missing godoc.

#### func (Locations) BackupDirRotating

```go
func (L Locations) BackupDirRotating(name string, max int) string
```
BackupDirRotating missing godoc.

#### func (Locations) BinPath

```go
func (L Locations) BinPath() string
```
BinPath missing godoc.

#### func (Locations) ConfigDir

```go
func (L Locations) ConfigDir() string
```
ConfigDir missing godoc.

#### func (Locations) ConfigFile

```go
func (L Locations) ConfigFile() string
```
ConfigFile missing godoc.

#### func (Locations) ToolsPath

```go
func (L Locations) ToolsPath() string
```
ToolsPath missing godoc.

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
	Targets   Targets   `json:"targets"`
	Functions Functions `json:"-"`
}
```

Project the project object

    proj := NewProject("/home/foo/project")

#### func  NewProject

```go
func NewProject(path string) *Project
```
NewProject missing godoc.

#### func (*Project) Description

```go
func (p *Project) Description(args ...string) string
```
Description missing godoc.

#### func (*Project) Name

```go
func (p *Project) Name(args ...string) string
```
Name missing godoc.

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

#### func (Project) WriteTable

```go
func (p Project) WriteTable(writer io.Writer)
```

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
	Files []os.FileInfo `json:"files"`
	Repos []VCSUrl      `json:"vcs"`
}
```

Targets missing godoc.

#### type VCSUrl

```go
type VCSUrl struct {
	Url         string `json:"url"`
	Type        string `json:"type"`
	Branch      string `json:"branch"`
	Destination string `json:"destination"`
}
```

VCSUrl missing godoc. Version Constrol Service Url
