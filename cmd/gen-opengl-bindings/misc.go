package main

import (
	"fmt"

	ugo "github.com/metaleap/go-util"
	usl "github.com/metaleap/go-util/slice"
)

type version struct {
	dotted, underscored string
	major, minor        int
}

var (
	cOrGoKeywords = []string{"func", "type", "near", "far", "string", "range"}
	vers          = map[string]*version{}
)

func parseVersion(dotted string) (ver *version) {
	majmin, _ := ugo.ParseVersion(dotted)
	ver = &version{dotted: dotted, major: majmin[0], minor: majmin[1]}
	ver.underscored = sfmt("%v_%v", ver.major, ver.minor)
	return
}

func saneName(name string) (sane string) {
	if sane = name; usl.StrHas(cOrGoKeywords, sane) {
		sane += "_"
	}
	return
}

func sfmt(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func ver(key string) (ver *version) {
	if ver = vers[key]; ver == nil {
		ver = parseVersion(key)
		vers[key] = ver
	}
	return
}
