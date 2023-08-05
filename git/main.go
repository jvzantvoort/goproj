package git

import (
	"fmt"
	"os"
	"os/exec"

	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

// GitCmd object for git
type GitCmd struct {
	Path       *Path
	Cwd        string
	Command    string
	CommandMap map[string]string
}

// Prefix missing godoc.
func (g GitCmd) Prefix() string {
	pc, _, _, _ := runtime.Caller(1)
	elements := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	return fmt.Sprintf("%s.%s", PackageName, elements[len(elements)-1])
}

// LogDebugf missing godoc.
func (g GitCmd) LogDebugf(format string, args ...interface{}) {
	pc, _, _, _ := runtime.Caller(2)
	elements := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	message := fmt.Sprintf(format, args...)

	prefix := fmt.Sprintf("%s.%s", PackageName, elements[len(elements)-1])
	log.Debug(prefix + message)
}

// LogFatalf missing godoc.
func (g GitCmd) LogFatalf(format string, args ...interface{}) {
	pc, _, _, _ := runtime.Caller(2)
	elements := strings.Split(runtime.FuncForPC(pc).Name(), ".")

	message := fmt.Sprintf(format, args...)

	prefix := fmt.Sprintf("%s.%s", PackageName, elements[len(elements)-1])
	log.Fatal(prefix + message)
}

// Execute missing godoc.
func (g GitCmd) Execute(args ...string) ([]string, []string, error) {

	// set log prefix and log start and end
	g.LogDebugf("start")
	defer g.LogDebugf("end")

	cmnd := []string{}
	cmnd = append(cmnd, args...)

	g.Debugf("command %s %s", g.Command, strings.Join(cmnd, " "))

	cmd := exec.Command(g.Command, cmnd...)

	stdout, err := cmd.StdoutPipe()
	PanicOnError("stdout pipe failed, %v", err)

	stderr, err := cmd.StderrPipe()
	PanicOnError("stderr pipe failed, %v", err)

	cmd.Dir = g.Cwd
	cmd.Start()

	stdout_list := Buffer2Slice(stdout)
	stderr_list := Buffer2Slice(stderr)

	eerror := cmd.Wait()
	PrintError("command failed, %v", eerror)

	return stdout_list, stderr_list, eerror
}

// NewGitCmd create a new git object
func NewGitCmd() *GitCmd {
	retv := &GitCmd{}

	retv.Debugf("start")
	defer retv.Debugf("end")

	retv.Path = NewPath("PATH")

	retv.CommandMap = map[string]string{
		"windows": "git.exe",
		"linux":   "git",
		"default": "git",
	}

	dir, err := os.Getwd()
	PrintFatal("failed to get cwd: %v", err)
	if err == nil {
		retv.Cwd = dir
	}

	if result, err := retv.Path.LookupPlatform(retv.CommandMap); err == nil {
		retv.Command = result
	}

	return retv
}

// Aliasses
// URL function returning the git url
func (g GitCmd) URL() string {
	stdout, _, _ := g.Execute("config", "--get", "remote.origin.url")
	if len(stdout) == 0 {
		return ""
	}
	return string(stdout[0])
}

// Branch function returning the current git branch
func (g GitCmd) Branch() string {
	stdout, _, _ := g.Execute("rev-parse", "--abbrev-ref", "HEAD")
	if len(stdout) == 0 {
		return ""
	}
	return string(stdout[0])
}

// Root function returning the git root
func (g GitCmd) Root() string {
	stdout, _, _ := g.Execute("rev-parse", "--show-toplevel")
	if len(stdout) == 0 {
		return ""
	}
	return string(stdout[0])
}

// Commit missing godoc.
func (g GitCmd) Commit(message string, args ...string) ([]string, []string, error) {
	log_prefix := g.Prefix()
	log.Debugf("%s: start", log_prefix)
	defer log.Debugf("%s: end", log_prefix)

	arglist := []string{}
	arglist = append(arglist, "commit")
	arglist = append(arglist, "--message")
	arglist = append(arglist, message)
	arglist = append(arglist, args...)

	return g.Execute(arglist...)
}
