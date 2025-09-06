# project_ui
--
    import "github.com/jvzantvoort/goproj/project_ui"


## Usage

#### type ProjectUI

```go
type ProjectUI struct {
	Project project.Project
}
```


#### func  NewProjectUI

```go
func NewProjectUI(p *project.Project) *ProjectUI
```

#### func (*ProjectUI) MainWindows

```go
func (ui *ProjectUI) MainWindows()
```
