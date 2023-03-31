package session
import (
	"strings"
	"testing"
)

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


func TestSessionFields(t *testing.T) {
	sess := &Session{"Name", "Path", "Type", 1, 1}
	reference := []string{"Name", "Path", "Type", "1970-01-01 01:00:01", "1970-01-01 01:00:01"}
	fields := sess.Fields()
	for num, indx := range fields {
		if reference[num] != indx {
			t.Errorf("got %q, wanted %q", indx, reference[num])
		}
	}
}

func TestSessionString(t *testing.T) {
	sess := &Session{"Name", "Path", "Type", 1, 1}
	refstr := "Name Created: 1970-01-01 01:00:01 Last active: 1970-01-01 01:00:01"
	outstr := sess.String()
	if refstr != outstr {
			t.Errorf("got %q, wanted %q", refstr, outstr)
	}
}
