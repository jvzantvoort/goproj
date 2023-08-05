# projecttype
--
    import "github.com/jvzantvoort/goproj/projecttype"


## Usage

```go
var Content embed.FS
```

#### func  CreateProjectType

```go
func CreateProjectType(projecttype string) error
```

#### type ProjectTypeConfig

```go
type ProjectTypeConfig struct {
	ProjectType    string `yaml:"projecttype"`
	ProjectTypeDir string
	Workdir        string            `yaml:"workdir"`
	Pattern        string            `yaml:"pattern"`
	SetupActions   []string          `yaml:"setupactions"`
	Files          []ProjectTypeFile `yaml:"files"`
}
```

ProjectTypeConfig defines a structure of a project type

#### func  NewProjectTypeConfig

```go
func NewProjectTypeConfig(projecttype string) ProjectTypeConfig
```
NewProjectTypeConfig read the relevant configfile and return ProjectTypeConfig
object with relevant data.

#### func (ProjectTypeConfig) Describe

```go
func (ptc ProjectTypeConfig) Describe()
```
Describe describe

#### func (ProjectTypeConfig) Exists

```go
func (ptc ProjectTypeConfig) Exists(targetpath string) bool
```

#### func (*ProjectTypeConfig) Init

```go
func (ptc *ProjectTypeConfig) Init(projtypeconfigdir, projecttype string) error
```

#### func (ProjectTypeConfig) UpdateConfigFile

```go
func (ptc ProjectTypeConfig) UpdateConfigFile(target string) error
```

#### func (ProjectTypeConfig) Write

```go
func (ptc ProjectTypeConfig) Write(boxname, target string) error
```

#### type ProjectTypeFile

```go
type ProjectTypeFile struct {
	Name        string `yaml:"name"`
	Destination string `yaml:"destination"`
	Mode        string `yaml:"mode"`
}
```

ProjectTypeFile defines a structure of a file
