package git

import (
	"bufio"
	"io"

	log "github.com/sirupsen/logrus"
)

// Buffer2Slice translate a io stream into a slice of strings.
//
//	stdout_list := Buffer2Slice(stdout)
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

// PrintError if err is not nil print fmtstr as error.
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
