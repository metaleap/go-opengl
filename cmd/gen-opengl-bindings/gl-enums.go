package main

import (
	xmlx "github.com/jteeuwen/go-pkg-xmlx"
	ustr "github.com/metaleap/go-util/str"
)

type glEnum struct {
	name, ref, val, ver string
	exts                []string
}

func (me *glPack) makeEnums() {
	var src glPackSrc
	for name, enum := range allEnums {
		if len(enum.exts) == 0 || len(enum.ver) > 0 || ustr.IsAnyInSlice(cfg.genExts, enum.exts...) {
			me.enums[name] = enum
		}
	}
	src.addLn("const (")
	for _, enum := range me.enums {
		if len(enum.val) == 0 && me.enums[enum.ref] == nil {
			me.enums[enum.ref] = allEnums[enum.ref]
		}
	}
	for _, enum := range me.enums {
		src.addLn("\t%s = %s", enum.name, ustr.Ifs(len(enum.val) == 0, enum.ref, enum.val))
	}
	src.addLn(")")
	me.sources["enums"] = src
}

func xmlWalkEnums() {
	total := 0
	xmlWalkDoc("enum", func(xn *xmlx.Node) {
		total++
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
			allEnums[enum.name] = enum
		}
	})
	println(sfmt("Picked %v/%v GL enums.", len(allEnums), total))
}
