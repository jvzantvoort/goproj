package project

import (
	"bufio"
	"io"
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// OneOrLess missing godoc.
func OneOrLess(args ...string) (string, error) {
	if len(args) == 0 {
		return "", ErrListEmpty
	}
	if len(args) == 1 {
		return args[0], nil
	}

	return "", ErrListTooLong
}

// MkdirAll missing godoc.
func MkdirAll(path string, mode int) {

	log.Debugf("MkdirAll: start")
	defer log.Debugf("MkdirAll: end")

	stat, err := os.Stat(path)

	// we found something
	if err == nil {
		// already exists
		if stat.IsDir() {
			log.Debugf("found dir: %s", path)
			return
		} else {
			log.Errorf("found target: %s but it is not a directory", path)
		}
	}

	mode_oct := os.FileMode(mode)
	os.MkdirAll(path, mode_oct)

}

// Buffer2Slice missing godoc.
func Buffer2Slice(stream io.ReadCloser) []string {
	retv := []string{}

	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		msg := scanner.Text()
		retv = append(retv, msg)
	}
	return retv
}

// PrintError missing godoc.
func PrintError(fmtstr string, err error) error {
	if err == nil {
		return err
	}
	log.Errorf(fmtstr, err)
	return err
}

// PrintFatal missing godoc.
func PrintFatal(fmtstr string, err error) error {
	if err == nil {
		return err
	}
	log.Fatalf(fmtstr, err)
	return err
}

// PanicOnError missing godoc.
func PanicOnError(fmtstr string, err error) {
	PrintError(fmtstr, err)
	if err != nil {
		panic(err)
	}
}

// FileExists missing godoc.
func FileExists(fpath string) (bool, os.FileInfo) {
	info, err := os.Stat(fpath)
	if err != nil {
		return false, info
	}

	// is a directory
	if info.IsDir() {
		return false, info
	}

	return true, info
}

// FileIsExecutable missing godoc.
func FileIsExecutable(fpath string) bool {
	exists, info := FileExists(fpath)
	if !exists {
		return false
	}

	goos := runtime.GOOS

	// windows doesn't do that
	if goos == "windows" {
		return true
	}

	mode := info.Mode()

	// Exec owner
	if mode&0100 != 0 {
		return true
	}

	// Exec group
	if mode&0010 != 0 {
		return true
	}

	// Exec other
	if mode&0001 != 0 {
		return true
	}
	return false
}

// Reverse missing godoc.
func Reverse(input []string) []string {
	var output []string

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}
