package main

import (
	"flag"
	"strings"

	xmlx "github.com/jteeuwen/go-pkg-xmlx"
	ugo "github.com/metaleap/go-util"
)

const (
	xmlns = "http://www.opengl.org/registry/"
)

var (
	flagOutDirPath = flag.String("outdir", ugo.GopathSrcGithub("go3d", "go-opengl", "api"), "Full directory path where gl.go (and/or extension sub-dirs if any) will be (re)generated.")
	flagGenExts    = flag.String("xgen", "EXT_texture_filter_anisotropic", "Space-separated list of fully-named extensions to be generated under outdirpath/x/prefix/name.")
	flagMinVer     = flag.String("minver", "3.2", "Minimum GL version. Whatever was deprecated in that version or earlier will be skipped.")

	cfg struct {
		minVer  *version
		genExts []string
	}

	err     error
	specDoc = xmlx.New()

	allEnums = map[string]*glEnum{}
	allFuncs = map[string]*glFunc{}
	allTypes = map[string]*glType{}
)

func main() {
	flag.Parse()
	if len(*flagGenExts) > 0 {
		cfg.genExts = strings.Split(*flagGenExts, " ")
	}
	if cfg.minVer = parseVersion(*flagMinVer); cfg.minVer.major < 1 {
		panic("What kind of minver is that?")
	}
	filePath := ugo.GopathSrcGithub("go3d", "go-opengl", "api", "gen-opengl-api", "xmlspecs", "opengl.xml")
	if err = specDoc.LoadFile(filePath, nil); err != nil {
		panic(err)
	}
	walkEnums()
	walkTypes()
	walkFuncs()
}
