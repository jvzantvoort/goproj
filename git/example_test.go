package git_test

import (
	"fmt"

	"github.com/jvzantvoort/goproj/git"
)

func ExamplePrintError() {
	var err error
	var msg string
	msg = "error message %s"
	err = nil
	git.PrintError(msg, err)
	err = fmt.Errorf("error message")
	// Output:
	//
	// error message error message
}
