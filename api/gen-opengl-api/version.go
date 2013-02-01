package main

import (
	"fmt"

	ugo "github.com/metaleap/go-util"
)

type version struct {
	dotted, underscored string
	major, minor        int
}

func parseVersion(dotted string) (ver *version) {
	majmin, _ := ugo.ParseVersion(dotted)
	ver = &version{dotted: dotted, major: majmin[0], minor: majmin[1]}
	ver.underscored = sfmt("%v_%v", ver.major, ver.minor)
	return
}

func sfmt(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
