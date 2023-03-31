package projecttype

import (
	"bytes"
	"text/template"

	"github.com/spf13/viper"
)

func (ptc *ProjectTypeConfig) readConfig(projtypeconfigdir string) error {

	viper.SetConfigName("config")
	viper.AddConfigPath(projtypeconfigdir)

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	return viper.Unmarshal(&ptc)
}

func (ptc ProjectTypeConfig) WriteTemplate(tmplname, target string) error {

	content, err := Asset("templates/" + tmplname)
	if err != nil {
		return err
	}

	text_template, err := template.New("tmpl").Parse(string(content))
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = text_template.Execute(buf, ptc)
	if err != nil {
		return err
	}

	return WriteContent(target, buf.String())
}
