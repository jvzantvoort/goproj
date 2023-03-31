// A session is a set of data describing a workspace in name, path and type. It
// also maintains timestamps on creation and activity.
//
package session

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

//
//

// Session struct representing a session.
type Session struct {
	Name     string `json:"name"` // name of the session
	Path     string `json:"path"` // path to the project
	Type     string `json:"type"` // type of project
	Activity int64  `json:"active"` // int time (unix time) of last time activated
	Created  int64  `json:"created"` // int time (unix time) of creation time
}

// Fields returns a list of columns of a struct.
//
// Fields returned
//   • Name
//   • Path
//   • Type
//   • Created ("YYYY-mm-dd HH:MM:SS")
//   • Activity ("YYYY-mm-dd HH:MM:SS")
//
func (s Session) Fields() []string {
	retv := []string{}

	ctime := time.Unix(s.Created, 0)
	atime := time.Unix(s.Activity, 0)

	retv = append(retv, s.Name)
	retv = append(retv, s.Path)
	retv = append(retv, s.Type)
	retv = append(retv, ctime.Format("2006-01-02 15:04:05"))
	retv = append(retv, atime.Format("2006-01-02 15:04:05"))
	return retv
}

// Write json output to an [io.Writer] compatible handle. It returns nil or the
// error of [json.MarshalIndent]
func (s Session) Write(writer io.Writer) error {
	content, err := json.MarshalIndent(s, "", "  ")
	if err == nil {
		fmt.Fprint(writer, string(content))
	}
	fmt.Fprintf(writer, "\n")
	return err
}

// Write session configuration to a sessionfile
func (s Session) WriteToFile(sessionfile string) error {
	file, err := os.OpenFile(sessionfile, os.O_CREATE, 0644)
	defer file.Close()
	s.Write(file)
	return err
}

// Read session content from a [io.Reader] object.
func (s *Session) Read(reader io.Reader) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) UpdateActivty() {
	now := time.Now()
	s.Activity = now.Unix()
}

// Return a string representation of [Session] object.
//
// Example output:
//  Name Created: 2006-01-02 15:04:05 Last active: 2006-01-02 15:04:05
//
func (s Session) String() string {
	ctime := time.Unix(s.Created, 0)
	atime := time.Unix(s.Activity, 0)

	var retv string
	retv += fmt.Sprintf("%s", s.Name)
	retv += " Created: " + ctime.Format("2006-01-02 15:04:05")
	retv += " Last active: " + atime.Format("2006-01-02 15:04:05")
	return retv
}

func NewSession(sess_name, sess_path, sess_type string) *Session {
	retv := &Session{}
	retv.Name = sess_name
	retv.Path = sess_path
	retv.Type = sess_type
	now := time.Now()
	retv.Created = now.Unix()
	retv.Activity = retv.Created
	return retv
}
