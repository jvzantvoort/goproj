package config

import (
	"strings"
	"testing"
)

type ticketExpandHome struct {
	inputstr  string
	outputstr string
	errmsg    string
}

var ticketExpandHomes = []ticketExpandHome{
	ticketExpandHome{"~/", "/lala", ""},
	ticketExpandHome{"~/foo", "/lala/foo", ""},
	ticketExpandHome{"/foo", "/foo", ""},
}

// ErrorContains checks if the error message in out contains the text in
// want.
//
// This is safe when out is nil. Use an empty string for want if you want to
// test that err is nil.
func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

func TestExpandHome(t *testing.T) {
	mc := &MainConfig{}

	mc.HomeDir = "/lala"

	for _, test := range ticketExpandHomes {
		gotstr, err := mc.ExpandHome(test.inputstr)

		if gotstr != test.outputstr {
			t.Errorf("got %q, wanted %q", gotstr, test.outputstr)
		}
		if !ErrorContains(err, test.errmsg) {
			t.Errorf("got %q, wanted %q", err, test.errmsg)
		}
	}

}
