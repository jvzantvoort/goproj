# projecttype
--
    import "github.com/jvzantvoort/goproj/projecttype"


## Usage

```go
const AssetDebug = false
```
AssetDebug is true if the assets were built with the debug flag enabled.

#### func  Asset

```go
func Asset(name string) ([]byte, error)
```
Asset loads and returns the asset for the given name. It returns an error if the
asset could not be found or could not be loaded.

#### func  AssetDigest

```go
func AssetDigest(name string) ([sha256.Size]byte, error)
```
AssetDigest returns the digest of the file with the given name. It returns an
error if the asset could not be found or the digest could not be loaded.

#### func  AssetDir

```go
func AssetDir(name string) ([]string, error)
```
AssetDir returns the file names below a certain directory embedded in the file
by go-bindata. For example if you run go-bindata on data/... and data contains
the following hierarchy:

    data/
      foo.txt
      img/
        a.png
        b.png

then AssetDir("data") would return []string{"foo.txt", "img"},
AssetDir("data/img") would return []string{"a.png", "b.png"},
AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
AssetDir("") will return []string{"data"}.

#### func  AssetInfo

```go
func AssetInfo(name string) (os.FileInfo, error)
```
AssetInfo loads and returns the asset info for the given name. It returns an
error if the asset could not be found or could not be loaded.

#### func  AssetNames

```go
func AssetNames() []string
```
AssetNames returns the names of the assets.

#### func  AssetString

```go
func AssetString(name string) (string, error)
```
AssetString returns the asset contents as a string (instead of a []byte).

#### func  CreateProjectType

```go
func CreateProjectType(projecttype string) error
```

#### func  Digests

```go
func Digests() (map[string][sha256.Size]byte, error)
```
Digests returns a map of all known files and their checksums.

#### func  MustAsset

```go
func MustAsset(name string) []byte
```
MustAsset is like Asset but panics when Asset would return an error. It
simplifies safe initialization of global variables.

#### func  MustAssetString

```go
func MustAssetString(name string) string
```
MustAssetString is like AssetString but panics when Asset would return an error.
It simplifies safe initialization of global variables.

#### func  RestoreAsset

```go
func RestoreAsset(dir, name string) error
```
RestoreAsset restores an asset under the given directory.

#### func  RestoreAssets

```go
func RestoreAssets(dir, name string) error
```
RestoreAssets restores an asset under the given directory recursively.

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
