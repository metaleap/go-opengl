package main

import (
	"path/filepath"

	uio "github.com/metaleap/go-util/io"
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
	for name, src := range me.sources {
		if err = uio.WriteTextFile(filepath.Join(cfg.outDirPath, "gl-"+name+".go"), "package gl\n\n"+string(src)); err != nil {
			break
		}
	}
	return
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
