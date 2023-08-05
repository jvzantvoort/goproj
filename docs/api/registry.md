# registry
--
    import "github.com/jvzantvoort/goproj/registry"

Register configs

## Usage

#### type Registry

```go
type Registry struct {
	Projects    map[string]project.Project `json:"projects"`
	LastUpdated int64                      `json:"lastUpdated"`
}
```


#### func  NewRegistry

```go
func NewRegistry() *Registry
```

#### func (*Registry) Get

```go
func (r *Registry) Get(name string) (project.Project, bool)
```

#### func (Registry) GetConfigFile

```go
func (r Registry) GetConfigFile() string
```
GetConfigFile returns the config file path from the registry

#### func (*Registry) List

```go
func (r *Registry) List() []string
```

#### func (*Registry) Load

```go
func (r *Registry) Load() error
```
Load loads the registry data from the registry file

#### func (*Registry) Read

```go
func (r *Registry) Read(reader io.Reader) error
```
Read reads the registry data from the registry file

#### func (*Registry) Register

```go
func (r *Registry) Register(p project.Project)
```

#### func (*Registry) Remove

```go
func (r *Registry) Remove(name string)
```
Remove removes a project from the registry data

    reg := registry.NewRegistry()
    reg.Remove("foo")

#### func (*Registry) Save

```go
func (r *Registry) Save() error
```
Save saves the registry data to the registry file

#### func (*Registry) SetLastUpdated

```go
func (r *Registry) SetLastUpdated()
```
SetLastUpdated sets the last updated time in the registry data

#### func (Registry) Write

```go
func (r Registry) Write(writer io.Writer) error
```
Write writes the registry data to the registry file
