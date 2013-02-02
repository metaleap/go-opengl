package main

import (
	"flag"
	"path/filepath"
	"strings"

	xmlx "github.com/jteeuwen/go-pkg-xmlx"
	ugo "github.com/metaleap/go-util"
	uio "github.com/metaleap/go-util/io"
)

const (
	xmlns = "http://www.opengl.org/registry/"
)

var (
	flagOutDir   = flag.String("outdir", ugo.GopathSrcGithub("go3d", "go-opengl"), "Full directory path\nwhere the ./core/ package folder will be generated.\n")
	flagGenExts  = flag.String("exts", "EXT_texture_filter_anisotropic EXT_texture_compression_s3tc EXT_texture_sRGB", "What extensions to include, space-separated.\n")
	flagMinVer   = flag.String("minver", "3.3", "Minimum GL version. Whatever was deprecated in that version or earlier\nwill be skipped and not included in the generated binding.\n")
	flagSpecFile = flag.String("specpath", ugo.GopathSrcGithub("go3d", "go-opengl", "cmd", "gen-opengl-bindings", "xmlspecs", "opengl.xml"), "Full path to the spec file.")

	cfg struct {
		outDirPath string
		minVer     *version
		genExts    []string
	}

	err     error
	specDoc = xmlx.New()

	allEnums = map[string]*glEnum{}
	allFuncs = map[string]*glFunc{}
	allTypes = map[string]string{}
)

func main() {
	flag.Parse()
	if cfg.minVer = parseVersion(*flagMinVer); cfg.minVer.major < 1 {
		panic("What kind of minver is that?")
	}
	if len(*flagGenExts) > 0 {
		cfg.genExts = strings.Split(*flagGenExts, " ")
	} else {
		cfg.genExts = []string{"core"}
	}
	if err = specDoc.LoadFile(*flagSpecFile, nil); err != nil {
		panic(err)
	}
	cfg.outDirPath = filepath.Join(*flagOutDir, "core")
	if err = uio.EnsureDirExists(cfg.outDirPath); err != nil {
		panic(err)
	}
	xmlWalkTypes()
	xmlWalkEnums()
	xmlWalkFuncs()
	if err = newGlPack().makeAllFiles(); err != nil {
		panic(err)
	}
}