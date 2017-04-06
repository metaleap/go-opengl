package main

import (
	"flag"
	"path/filepath"
	"strings"

	xmlx "github.com/go-forks/go-pkg-xmlx"
	"github.com/metaleap/go-util-fs"
	"github.com/metaleap/go-util-misc"
)

const (
	xmlns = "http://www.opengl.org/registry/"
)

var (
	flagOutDir   = flag.String("outdir", ugo.GopathSrcGithub("go3d", "go-opengl"), "Full directory path where the ./core/ package folder will be generated.\n")
	flagGenExts  = flag.String("exts", "EXT_texture_filter_anisotropic EXT_texture_compression_s3tc EXT_texture_sRGB", "What extensions to include, space-separated. At a minimum you'll want the 'ubiquitous extensions' EXT_texture_filter_anisotropic EXT_texture_compression_s3tc EXT_texture_sRGB. Use asterisk * to include all extensions in the specpath file (pkg won't compile as of yet though).\n")
	flagMinVer   = flag.String("minver", "3.3", "Minimum GL version. Whatever was deprecated in that version or earlier will be skipped and not included in the generated binding. IF you specify a custom minver, you'll also need to specify a different requirefunc.\n")
	flagSpecFile = flag.String("specpath", ugo.GopathSrcGithub("go3d", "go-opengl", "cmd", "gen-opengl-bindings", "xmlspecs", "opengl.xml"), "Full path to the spec file.\n")
	flagProcAddr = flag.String("requirefunc", "BindSampler", "The name of a GL function (without gl prefix) that the binding's Init() method checks for to test initialization success. IF you specify a custom minver, you'll need a different func name here.\n")
	flagSupports = flag.Bool("supports", true, "Creates a Supports struct that allows runtime checking for individual function availability")
	flagTry      = flag.Bool("try", true, "Generate function wrappers that check for GL errors and return Go errors")

	cfg struct {
		altTryFile struct {
			only             bool
			funcs            []string
			outPath, pkgName string
		}
		outDirPath string
		minVer     *version
		genExts    []string
		genExtsAll bool
	}

	err     error
	specDoc = xmlx.New()

	allEnums = map[string]*glEnum{}
	allFuncs = map[string]*glFunc{}
	allTypes = map[string]string{}
)

func main() {
	if false {
		cfg.altTryFile.only, cfg.altTryFile.pkgName = true, "glutil"
		cfg.altTryFile.funcs = []string{"AttachShader", "BufferData", "BufferSubData", "CreateProgram", "CreateShader", "GenBuffers", "GenerateMipmap", "GenTextures", "GenVertexArrays", "ShaderSource", "TexImage2D", "TexStorage2D", "TexSubImage2D"}
		cfg.altTryFile.outPath = ugo.GopathSrcGithub("go3d", "go-opengl", "util", "-gen-try.go")
	}
	flag.Parse()
	if cfg.minVer = parseVersion(*flagMinVer); cfg.minVer.major < 1 {
		panic("What kind of minver is that?")
	}
	if len(*flagGenExts) > 0 {
		if *flagGenExts == "*" {
			cfg.genExtsAll = true
		} else {
			cfg.genExts = strings.Split(*flagGenExts, " ")
		}
	}
	if err = specDoc.LoadFile(*flagSpecFile, nil); err != nil {
		panic(err)
	}
	cfg.outDirPath = filepath.Join(*flagOutDir, "core")
	if err = ufs.EnsureDirExists(cfg.outDirPath); err != nil {
		panic(err)
	}
	if cfg.genExtsAll {
		xmlWalkExts()
	}
	xmlWalkTypes()
	xmlWalkEnums()
	xmlWalkFuncs()
	if err = newGlPack().makeAllFiles(); err != nil {
		panic(err)
	}
}
