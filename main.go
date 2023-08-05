package goproj

import (
	"fmt"

	"github.com/hashicorp/go-version"
)

func AppTooOld(verstr string) error {
	appver, _ := version.NewVersion(Version)
	configver, _ := version.NewVersion(verstr)

	if appver.LessThan(configver) {
		return fmt.Errorf("Application version (%s) is too old for this config (%s)", appver, configver)
	}
	return nil
}
