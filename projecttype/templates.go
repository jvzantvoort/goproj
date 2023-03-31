package projecttype

import "strings"

func ListTemplates() []string {
	retv := []string{}
	for _, target := range AssetNames() {
		retv = append(retv, strings.TrimPrefix(target, "templates/"))
	}
	return retv
}
