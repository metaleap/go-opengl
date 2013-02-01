package main

import (
	xmlx "github.com/jteeuwen/go-pkg-xmlx"
	ustr "github.com/metaleap/go-util/str"
)

type nodeFunc func(*xmlx.Node)

func checkForUnknownAtts(xn *xmlx.Node, knownAttNames ...string) {
	for _, att := range xn.Attributes {
		if !ustr.IsInSlice(knownAttNames, att.Name.Local) {
			println("ENUM ATT: " + att.Name.Local)
		}
	}
}

func isLegacy(xn *xmlx.Node) bool {
	asd, asr := xas(xn, "deprecated"), xas(xn, "removed")
	return (len(asd) > 0 && asd <= cfg.minVer.dotted) || (len(asr) > 0 && asr <= cfg.minVer.dotted)
}

func walkDoc(nodeName string, onNode nodeFunc) {
	for _, xn := range specDoc.SelectNodesRecursive(xmlns, nodeName) {
		onNode(xn)
	}
}

func walkDocN(onNode nodeFunc, names ...string) {
	for _, name := range names {
		walkDoc(name, onNode)
	}
}

func walkEnums() {
	walkDoc("enum", func(xn *xmlx.Node) {
		// checkForUnknownAtts(xn, "name", "version", "value", "deprecated", "removed", "ref")
		if !isLegacy(xn) {
			enum := &glEnum{name: xas(xn, "name"), ref: xas(xn, "ref"), ver: xas(xn, "version"), val: xas(xn, "value")}
			if len(enum.ver) == 0 {
				for _, extNode := range xn.SelectNodes(xmlns, "ext") {
					if extName := xas(extNode, "name"); len(extName) > 0 {
						enum.exts = append(enum.exts, extName)
					}
				}
			}
			if len(enum.ver) > 0 || len(enum.ext) == 0 || ustr.IsInSlice(cfg.genExts, enum.ext) {
				allEnums[enum.name] = enum
				if len(enum.ext) > 0 {
					println(enum.ext)
				}
			}
		}
	})
}

func walkFuncs() {
	walkDoc("function", func(xn *xmlx.Node) {
		if !isLegacy(xn) {
		}
	})
}

func walkTypes() {
	walkDoc("type-def", func(xn *xmlx.Node) {
		// checkForUnknownAtts(xn, "typename", "C-lang")
		typeDef := &glType{name: xas(xn, "typename"), clang: xas(xn, "C-lang")}
		allTypes[typeDef.name] = typeDef
	})
}

func xas(xn *xmlx.Node, name string) string {
	return xn.As("", name)
}
