package main

import (
	"path/filepath"
	"strings"

	"github.com/go-utils/ufs"
	"github.com/go-utils/ustr"
)

type glPack struct {
	enums   map[string]*glEnum
	funcs   map[string]*glFunc
	sources map[string]glPackSrc
}

func newGlPack() (me *glPack) {
	me = &glPack{enums: map[string]*glEnum{}, funcs: map[string]*glFunc{}}
	return
}

func (me *glPack) makeAllFiles() (err error) {
	me.makeAllSources()
	if !cfg.altTryFile.only {
		for name, src := range me.sources {
			if err = ufs.WriteTextFile(filepath.Join(cfg.outDirPath, "gl-"+name+".go"), "package gl\n\n"+string(src)); err != nil {
				return
			}
		}
	}
	if len(cfg.altTryFile.funcs) > 0 {
		var (
			src glPackSrc
			fun *glFunc
		)
		src.addLn(`package %s

import (
	gl "github.com/go3d/go-opengl/core"
)

//	Used only for the exported package-global variable 'Try'.
//	Provides methods that each wrap a GL function call and immediately check for GL errors to be returned as Go errors.
type GlTry struct {
}
`, cfg.altTryFile.pkgName)
		for _, try := range cfg.altTryFile.funcs {
			if fun = me.funcs[try]; fun != nil {
				me.makeTryFunc(&src, fun, "gl.", "Util.")
			}
		}
		err = ufs.WriteTextFile(cfg.altTryFile.outPath, string(src))
	}
	return
}

func (me *glPack) makeTryFunc(src *glPackSrc, fun *glFunc, glPrefix, errPrefix string) {
	src.addLn("//\tCalls %s%s() and yields the err returned by %sError(), if any.", glPrefix, fun.name, errPrefix)
	src.add("func (_ GlTry) %s(", fun.name)
	args, ret := "", ""
	for i, fp := range fun.params {
		fpgtPtr := strings.HasSuffix(fp.typeRef.g, "Ptr")
		tmp := strings.Replace(ustr.Ifs(ustr.IsOneOf(fp.kind, "array", "reference") && !fpgtPtr, "*"+fp.typeRef.g, fp.typeRef.g), "***", "**", -1)
		if pos := strings.LastIndex(tmp, "*"); pos < 0 {
			tmp = glPrefix + tmp
		} else {
			tmp = tmp[:pos+1] + glPrefix + tmp[pos+1:]
		}
		src.add("%s %s", fp.saneName, tmp)
		args += fp.saneName
		if i < len(fun.params)-1 {
			src.add(", ")
			args += ", "
		}
	}
	src.add(") (err error")
	if len(fun.retType.g) > 0 {
		ret = "tryRetVal__ = "
		src.add(", tryRetVal__ %v", glPrefix+fun.retType.g)
	}
	src.add(") {")
	src.addLn(`
	%s%s(%s)
	err = %sLastError("%s()")
	return`, ret, glPrefix+fun.name, args, errPrefix, glPrefix+fun.name)
	src.addLn("}\n")
}

func (me *glPack) makeAllSources() {
	me.sources = map[string]glPackSrc{}
	me.makeEnums()
	me.makeFuncs()
	me.makeCgo()
	return
}

type glPackSrc string

func (me *glPackSrc) add(fmt string, args ...interface{}) {
	*me += glPackSrc(sfmt(fmt, args...))
}

func (me *glPackSrc) addLn(fmt string, args ...interface{}) {
	me.add(fmt+"\n", args...)
}
