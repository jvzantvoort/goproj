package git

import (
	"bufio"
	"io"

	log "github.com/sirupsen/logrus"
)

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

func PrintError(fmtstr string, err error) error {
	if err == nil {
		return err
	}
	log.Errorf(fmtstr, err)
	return err
}

func PrintFatal(fmtstr string, err error) error {
	if err == nil {
		return err
	}
	log.Fatalf(fmtstr, err)
	return err
}


func PanicOnError(fmtstr string, err error) {
	PrintError(fmtstr, err)
	if err != nil {
		panic(err)
	}
}
