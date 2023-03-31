// Sessions
//
// Sessions works with all the sessions as a collection
package session

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"

	"github.com/olekukonko/tablewriter"
)

// Sessions represents a list of sessions
type Sessions struct {
	CacheDir     string
	SessionsList []Session
}

func (S Sessions) OrderedByActive() {
	sort.Slice(S.SessionsList[:], func(i, j int) bool {
		return S.SessionsList[i].Activity > S.SessionsList[j].Activity
	})
}

// Writer writes a table representation of the sessions to a io.Writer
// compatible object.
func (S Sessions) Writer(writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Name", "Path", "Type", "Created", "Active"})
	S.OrderedByActive()
	for _, indx := range S.Fields() {
		table.Append(indx)
	}
	table.SetHeaderLine(true)
	table.SetBorder(false)
	table.Render()
}

// Return a slice of [Session.String] representations of each session.
func (S Sessions) List() []string {
	retv := []string{}
	for _, indx := range S.SessionsList {
		retv = append(retv, indx.String())
	}
	return retv
}

// Append a sessions to the list
func (S *Sessions) Append(sess *Session) {
	S.SessionsList = append(S.SessionsList, *sess)
}

// return a slice of slices of Fields
func (S Sessions) Fields() [][]string {
	retv := [][]string{}

	for _, indx := range S.SessionsList {
		retv = append(retv, indx.Fields())
	}
	return retv
}

// Internal sessions slices contains a slice name
func (S Sessions) SliceContainsSession(name string) bool {
	if len(S.SessionsList) == 0 {
		return false
	}
	for _, indx := range S.SessionsList {
		if name == indx.String() {
			return true
		}
	}
	return false
}

// load the session sessions
func (S *Sessions) Load() error {
	files, err := ioutil.ReadDir(S.CacheDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		fpath := path.Join(S.CacheDir, f.Name())
		fileh, _ := os.Open(fpath)
		defer fileh.Close()
		sess := &Session{}
		sess.Read(fileh)
		S.Append(sess)
	}
	return nil
}

// create a new Sessions object
func NewSessions() *Sessions {
	retv := &Sessions{}
	return retv
}
