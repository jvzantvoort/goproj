# utils
--
    import "github.com/jvzantvoort/goproj/utils"

Package utils provides utilitary functions.

## Usage

#### func  FileExists

```go
func FileExists(fpath string) (bool, os.FileInfo)
```
FileExists check if a target exists and is a file.

    check, info := utils.FileExists("/etc/passwd")
    if check {
       fmt.Printf("size: %d\n", info.Size())
    }

#### func  FileIsExecutable

```go
func FileIsExecutable(fpath string) bool
```
FileIsExecutable file exists and is executable

#### func  GetHomeDir

```go
func GetHomeDir() string
```
GetHomeDir get the user's homedir

#### func  MkdirP

```go
func MkdirP(dirname string, mode int) error
```
MkdirP

    err := utils.MkdirP("/lala", int(0755))
    if err != nil {
      panic(err)
    }

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
