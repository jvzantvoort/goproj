# template
--
    import "github.com/jvzantvoort/goproj/template"


## Usage

#### type CloneUrl

```go
type CloneUrl struct {
	Name        string `yaml:"name"`
	Url         string `yaml:"url"`
	Destination string `yaml:"destination"`
	Branch      string `yaml:"branch"`
}
```

CloneUrl missing godoc.

#### func (CloneUrl) Clone

```go
func (cu CloneUrl) Clone()
```
Clone missing godoc.

#### func (CloneUrl) Exists

```go
func (cu CloneUrl) Exists()
```
Exists missing godoc.

#### func (CloneUrl) Pull

```go
func (cu CloneUrl) Pull()
```
Pull missing godoc.

#### type File

```go
type File struct {
	Name        string `yaml:"name"`
	Destination string `yaml:"destination"`
	Mode        string `yaml:"mode"`
}
```

File missing godoc.

#### type Setup

```go
type Setup struct {
	Clone    []CloneUrl
	Commands []string
}
```

Setup missing godoc.

#### type Template

```go
type Template struct {
	Name    string `yaml:"name"`
	Pattern string `yaml:"pattern"`
	Version int    `yaml:"version"`
	Setup   Setup  `yaml:"setup"`
	Files   []File `yaml:"files"`
}
```

Template missing godoc.
