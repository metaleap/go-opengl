package main

import (
	"strings"

	xmlx "github.com/jteeuwen/go-pkg-xmlx"
	ustr "github.com/metaleap/go-util/str"
)

type glFunc struct {
	cat, name, ext, ver string
	retType             glTypeRef
	params              []*glFuncParam
	clang               struct {
		entry, wrapper, supporter string
	}
}

type glFuncParam struct {
	name, kind, compute, saneName string
	typeRef                       glTypeRef
	input                         bool
}

func (me *glFunc) makeCgo() {
	var (
		i  int
		fp *glFuncParam
	)
	me.clang.entry = sfmt("%s (APIENTRYP ptrgl%s)(", me.retType.c, me.name)
	for i, fp = range me.params {
		me.clang.entry += sfmt("%s %s", ustr.Ifs(ustr.IsOneOf(fp.kind, "array", "reference"), fp.typeRef.c+"*", fp.typeRef.c), fp.name)
		if i < len(me.params)-1 {
			me.clang.entry += ", "
		}
	}
	me.clang.entry += ");"

	if *flagSupports {
		me.clang.supporter = sfmt("int cancall%s() { return ((ptrgl%s == NULL) ? 0 : 1); }", me.name, me.name)
	}

	me.clang.wrapper = sfmt("%s call%s(", me.retType.c, me.name)
	for i, fp = range me.params {
		me.clang.wrapper += sfmt("%s %s", ustr.Ifs(ustr.IsOneOf(fp.kind, "array", "reference"), fp.typeRef.c+"*", fp.typeRef.c), fp.saneName)
		if i < len(me.params)-1 {
			me.clang.wrapper += ", "
		}
	}
	me.clang.wrapper += sfmt(") { %s (*ptrgl%s)(", ustr.Ifs(me.retType.c == "void", "", "return"), me.name)
	for i, fp = range me.params {
		me.clang.wrapper += fp.saneName
		if i < len(me.params)-1 {
			me.clang.wrapper += ", "
		}
	}
	me.clang.wrapper += "); }"
}

func (me *glPack) makeFuncs() {
	var (
		i       int
		fpgtPtr bool
		src     glPackSrc
		fun     *glFunc
		fp      *glFuncParam
	)
	isBeast := func() bool {
		b := (!strings.HasPrefix(fun.cat, "VERSION_")) && (!strings.HasSuffix(fun.name, fun.cat[:strings.Index(fun.cat, "_")])) && !strings.HasSuffix(fun.name, "EXT")
		return b
	}
	for _, fun = range allFuncs {
		if strings.Contains(fun.cat, " ") {
			println(fun.cat)
		}
		if strings.HasPrefix(fun.cat, "VERSION_") || ustr.IsInSlice(cfg.genExts, fun.cat) || isBeast() {
			me.funcs[fun.name] = fun
			if ustr.IsInSlice(cfg.genExts, fun.cat) {
				println("EXT:" + fun.name)
			}
		}
	}

	for _, fun = range me.funcs {
		src.addLn("//\tGLAPI Wiki: http://www.opengl.org/wiki/GLAPI/gl%s", fun.name)
		src.addLn("//\tSDK doc: http://www.opengl.org/sdk/docs/man/xhtml/gl%s.xml", fun.name)
		src.add("func %s(", fun.name)
		for i, fp = range fun.params {
			fpgtPtr = strings.HasSuffix(fp.typeRef.g, "Ptr")
			src.add("%s %s", fp.saneName, strings.Replace(ustr.Ifs(ustr.IsOneOf(fp.kind, "array", "reference") && !fpgtPtr, "*"+fp.typeRef.g, fp.typeRef.g), "***", "**", -1))
			if i < len(fun.params)-1 {
				src.add(", ")
			}
		}
		src.add(") %s {\n\t", fun.retType.g)
		if len(fun.retType.g) > 0 {
			src.add("return (%s)(", fun.retType.g)
		}
		src.add("C.call%s(", fun.name)
		for i, fp = range fun.params {
			fpgtPtr = fp.typeRef.g == "Ptr"
			if fp.typeRef.c == "GLchar* const" {
				src.add("(**C.GLchar)(unsafe.Pointer(%s))", fp.saneName)
			} else if fp.typeRef.c == "GLDEBUGPROC" {
				src.add("(*[0]byte)(%s)", fp.saneName)
			} else {
				src.add("(%s%s)(%s)", ustr.Ifs(ustr.IsOneOf(fp.kind, "array", "reference") && !fpgtPtr, "*", ""), ustr.Ifs(strings.HasSuffix(fp.typeRef.g, "Ptr"), "unsafe.Pointer", "C."+fp.typeRef.c), fp.saneName)
			}
			if i < len(fun.params)-1 {
				src.add(", ")
			}
		}
		src.add(")")
		if len(fun.retType.g) > 0 {
			src.add(")")
		}
		src.addLn("\n}\n")
		fun.makeCgo()
	}
	me.sources["funcs"] = src
}

func xmlWalkFuncs() {
	var (
		total int
		fun   *glFunc
		param *glFuncParam
	)
	xmlWalkDoc("function", func(xn *xmlx.Node) {
		total++
		if !isLegacy(xn) {
			fun = &glFunc{name: xas(xn, "name"), ver: xas(xn, "version"), cat: xas(xn, "category"), ext: xas(xn, "extension")}
			fun.retType.set(xas(xn, "return"))
			for _, pn := range xn.SelectNodes(xmlns, "param") {
				param = &glFuncParam{name: xas(pn, "name"), kind: xas(pn, "kind"), compute: xas(pn, "compute"), input: xas(pn, "input") != "false"}
				param.saneName = saneName(param.name)
				if param.typeRef.set(xas(pn, "type")); len(param.typeRef.g) == 0 {
					param.typeRef.g = "Ptr"
				}
				fun.params = append(fun.params, param)
			}
			allFuncs[fun.name] = fun
		}
	})
	println(sfmt("Picked %v/%v GL functions.", len(allFuncs), total))
}
