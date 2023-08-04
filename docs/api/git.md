# git
--
    import "github.com/jvzantvoort/goproj/git"


## Usage

```go
const (
	PackageName string = "GitCmd"
)
```

#### func  Buffer2Slice

```go
func Buffer2Slice(stream io.ReadCloser) []string
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

#### func (GitCmd) Execute

```go
func (g GitCmd) Execute(args ...string) ([]string, []string, error)
```

#### func (GitCmd) LogDebugf

```go
func (g GitCmd) LogDebugf(format string, args ...interface{})
```

#### func (GitCmd) LogFatalf

```go
func (g GitCmd) LogFatalf(format string, args ...interface{})
```

#### func (GitCmd) Prefix

```go
func (g GitCmd) Prefix() string
```

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
