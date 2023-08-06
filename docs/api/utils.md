# utils
--
    import "github.com/jvzantvoort/goproj/utils"

Package utils provides utilitary functions.

## Usage

```go
const (
	WIDTHSUBS          int             = 20
	TitleColor         color.Attribute = color.FgMagenta
	InfoNameColor      color.Attribute = color.Bold
	InfoValueColor     color.Attribute = color.FgYellow
	BranchDefaultColor color.Attribute = color.FgBlue
	BranchChangedColor color.Attribute = color.FgYellow

	SuccessColor color.Attribute = color.FgGreen
	FailureColor color.Attribute = color.FgRed
)
```

#### func  CenterLine

```go
func CenterLine(line string, width int) string
```

#### func  ErrorBox

```go
func ErrorBox(format string, args ...interface{})
```

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

#### func  PrintFailed

```go
func PrintFailed(format string, args ...interface{})
```

#### func  PrintFatal

```go
func PrintFatal(fmtstr string, err error) error
```
PrintFatal missing godoc.

#### func  PrintStatus

```go
func PrintStatus(colorattr color.Attribute, status, format string, args ...interface{})
```

#### func  PrintSuccess

```go
func PrintSuccess(format string, args ...interface{})
```

#### func  TextBox

```go
func TextBox(title, format string, args ...interface{})
```
