# git
--
    import "github.com/jvzantvoort/goproj/git"

PATH type handling.

## Usage

```go
const (
	PackageName string = "GitCmd"
)
```
PackageName missing godoc.

#### func  Buffer2Slice

```go
func Buffer2Slice(stream io.ReadCloser) []string
```
Buffer2Slice translate a io stream into a slice of strings.

    stdout_list := Buffer2Slice(stdout)

#### func  PanicOnError

```go
func PanicOnError(fmtstr string, err error)
```
PanicOnError missing godoc.

#### func  PrintError

```go
func PrintError(fmtstr string, err error) error
```
PrintError if err is not nil print fmtstr as error.

#### func  PrintFatal

```go
func PrintFatal(fmtstr string, err error) error
```
PrintFatal missing godoc.

#### type GitCmd

```go
type GitCmd struct {
	Path       *Path
	Cwd        string
	Command    string
	CommandMap map[string]string
}
```

GitCmd object for git

#### func  NewGitCmd

```go
func NewGitCmd() *GitCmd
```
NewGitCmd create a new git object

#### func (GitCmd) Branch

```go
func (g GitCmd) Branch() string
```
Branch function returning the current git branch

#### func (GitCmd) Commit

```go
func (g GitCmd) Commit(message string, args ...string) ([]string, []string, error)
```
Commit missing godoc.

#### func (GitCmd) Execute

```go
func (g GitCmd) Execute(args ...string) ([]string, []string, error)
```
Execute missing godoc.

#### func (GitCmd) LogDebugf

```go
func (g GitCmd) LogDebugf(format string, args ...interface{})
```
LogDebugf missing godoc.

#### func (GitCmd) LogFatalf

```go
func (g GitCmd) LogFatalf(format string, args ...interface{})
```
LogFatalf missing godoc.

#### func (GitCmd) Prefix

```go
func (g GitCmd) Prefix() string
```
Prefix missing godoc.

#### func (GitCmd) Root

```go
func (g GitCmd) Root() string
```
Root function returning the git root

#### func (GitCmd) URL

```go
func (g GitCmd) URL() string
```
Aliasses URL function returning the git url

#### type Path

```go
type Path struct {
	Type        string
	Home        string
	Directories []string
}
```

Path missing godoc.

#### func  NewPath

```go
func NewPath(pathname string) *Path
```
NewPath missing godoc.

#### func (*Path) AppendPath

```go
func (p *Path) AppendPath(inputdir string) error
```
AppendPath append a path to the list of Directories

#### func (Path) HavePath

```go
func (p Path) HavePath(inputdir string) bool
```
HavePath missing godoc.

#### func (*Path) Import

```go
func (p *Path) Import(path string)
```
Import missing godoc.

#### func (Path) IsEmpty

```go
func (p Path) IsEmpty() bool
```
IsEmpty missing godoc.

#### func (Path) Lookup

```go
func (p Path) Lookup(target string) (string, error)
```
Lookup missing godoc.

#### func (Path) LookupMulti

```go
func (p Path) LookupMulti(targets ...string) (string, error)
```
LookupMulti missing godoc.

#### func (Path) LookupPlatform

```go
func (p Path) LookupPlatform(pathmap map[string]string) (string, error)
```
LookupPlatform lookup paths based on platform

#### func (Path) MapGetPlatform

```go
func (p Path) MapGetPlatform(pathmap map[string]string) (string, error)
```
MapGetPlatform missing godoc.

#### func (Path) Prefix

```go
func (p Path) Prefix() string
```
Prefix missing godoc.

#### func (*Path) PrependPath

```go
func (p *Path) PrependPath(inputdir string) error
```
PrependPath missing godoc.

#### func (Path) ReturnExport

```go
func (p Path) ReturnExport() string
```
ReturnExport missing godoc.
