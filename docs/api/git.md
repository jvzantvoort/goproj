# commands
--
    import "github.com/jvzantvoort/goproj/git"

PATH type handling.

## Usage

#### type Path

```go
type Path struct {
	Type        string
	Home        string
	Directories []string
}
```


#### func  NewPath

```go
func NewPath(pathname string) *Path
```

#### func (*Path) AppendPath

```go
func (p *Path) AppendPath(inputdir string) error
```
AppendPath append a path to the list of Directories

#### func (Path) HavePath

```go
func (p Path) HavePath(inputdir string) bool
```

#### func (*Path) Import

```go
func (p *Path) Import(path string)
```

#### func (Path) IsEmpty

```go
func (p Path) IsEmpty() bool
```

#### func (Path) Lookup

```go
func (p Path) Lookup(target string) (string, error)
```

#### func (Path) LookupMulti

```go
func (p Path) LookupMulti(targets ...string) (string, error)
```

#### func (Path) LookupPlatform

```go
func (p Path) LookupPlatform(pathmap map[string]string) (string, error)
```
LookupPlatform lookup paths based on platform

#### func (Path) MapGetPlatform

```go
func (p Path) MapGetPlatform(pathmap map[string]string) (string, error)
```

#### func (Path) Prefix

```go
func (p Path) Prefix() string
```

#### func (*Path) PrependPath

```go
func (p *Path) PrependPath(inputdir string) error
```

#### func (Path) ReturnExport

```go
func (p Path) ReturnExport() string
```
