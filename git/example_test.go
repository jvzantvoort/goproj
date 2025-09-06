package git_test

import "fmt"

func ExamplePrintError() {
	var err error
	var msg string
	msg = "error message %s"
	err = nil
	PrintError(msg, err)
	err = fmt.Errorf("error message")
	// Output:
	//
	// error message error message
}
