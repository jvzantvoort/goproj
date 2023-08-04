package project

import (
	"fmt"
	"os/exec"
	"path"

	"strings"

	log "github.com/sirupsen/logrus"
)

type Functions struct {
	Cwd     string
	ToolsPath string
	Locations Locations
}

type fn func(...string) ([]string, []string, error)

func NewFunctions(locations Locations) *Functions {
	retv := &Functions{}
	retv.Locations = locations
	retv.Cwd = retv.Locations.RootDir
	return retv
}

func (f Functions) Which(command string) string {
	return path.Join(f.Locations.ToolsPath(), command)
}

func (f Functions) LookupExt(command string) (string, error) {
	executable := NewExecutable(command)
	return executable.LookupExt()
}

func (f Functions) Execute(command string, args ...string) ([]string, []string, error) {
	var err error
	var msg string
	var cmndpath string
	var stdout_list []string
	var stderr_list []string
	var cmnd []string
	var cmd *exec.Cmd

	msg = fmt.Sprintf("command: %s", command)

	log.Debugf("%s, start", msg)
	defer log.Debugf("%s, end", msg)

	cmndpath, err = f.LookupExt(command)
	if err != nil {
		log.Errorf("cannot find path for %s\n", command)
	}

	cmnd = append(cmnd, args...)

	log.Debugf("%s: %s %s", msg, cmndpath, strings.Join(cmnd, " "))

	cmd = exec.Command(cmndpath, cmnd...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Errorf("stdout pipe failed, %v", err)
		return stdout_list, stderr_list, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Errorf("stderr pipe failed, %v", err)
		return stdout_list, stderr_list, err
	}

	log.Debugf("%s, pipes confirmed", msg)
	cmd.Dir = f.Cwd
	log.Debugf("%s, cwd %s", msg, cmd.Dir)
	cmd.Start()

	log.Debugf("%s, started", msg)
	stdout_list = Buffer2Slice(stdout)
	stderr_list = Buffer2Slice(stderr)

	err = cmd.Wait()
	if err != nil {
		log.Errorf("command failed, %v", err)
		return stdout_list, stderr_list, err
	}

	return stdout_list, stderr_list, err
}

func ExecNonFatal(f fn, args ...string) {
	stdout_list, stderr_list, err := f(args...)
	for _, line := range stdout_list {
		log.Infof("out>  %s", line)
	}

	for _, line := range stderr_list {
		log.Infof("err>  %s", line)
	}

	if err != nil {
		log.Infof("fail> %v", err)
	}

}

func (f Functions) Setup(args ...string) ([]string, []string, error) {
	return f.Execute(f.Which("setup"), args...)
}

func (f Functions) Teardown(args ...string) ([]string, []string, error) {
	return f.Execute(f.Which("teardown"), args...)
}

func (f Functions) Status(args ...string) ([]string, []string, error) {
	return f.Execute(f.Which("status"), args...)
}

func (f Functions) Backup(args ...string) ([]string, []string, error) {
	return f.Execute(f.Which("backup"), args...)
}

func (f Functions) Build(args ...string) ([]string, []string, error) {
	return f.Execute(f.Which("build"), args...)
}

func (f Functions) Package(args ...string) ([]string, []string, error) {
	return f.Execute(f.Which("package"), args...)
}

func (f Functions) Publish(args ...string) ([]string, []string, error) {
	return f.Execute(f.Which("publish"), args...)
}


func (f Functions) Test(args ...string) ([]string, []string, error) {
	return f.Execute(f.Which("test"), args...)
}

func (f Functions) SetupProject(args ...string) {
	MkdirAll(f.Locations.ToolsPath(), 0755)
	MkdirAll(f.Locations.BinPath(), 0755)
	ExecNonFatal(f.Setup, args...)
}

func (f Functions) BuildProject(args ...string) {
	MkdirAll(f.Locations.ToolsPath(), 0755)
	MkdirAll(f.Locations.BinPath(), 0755)
	ExecNonFatal(f.Build, args...)
	ExecNonFatal(f.Package, args...)
}


