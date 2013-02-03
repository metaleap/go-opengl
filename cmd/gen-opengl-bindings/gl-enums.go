package main

import (
	"strconv"
	"strings"

	xmlx "github.com/jteeuwen/go-pkg-xmlx"
	ustr "github.com/metaleap/go-util/str"
)

type glEnum struct {
	name, ref, val, ver string
	exts                []string
}

func (me *glPack) makeEnums() {
	var src glPackSrc
	src.addLn("import \"fmt\"")
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
	src.addLn(")\n")
	toUi := func(str string) (ui uint64) {
		ui, _ = strconv.ParseUint(str, 0, 64)
		return
	}
	enumsDone := map[uint64]bool{toUi("0xFFFFFFFFFFFFFFFF"): true}
	isEarlyEnum := func(name string) bool {
		return (strings.HasSuffix(name, "_SHADER") && strings.Count(name, "_") <= 2) || ustr.IsOneOf(name, "INVALID_ENUM", "INVALID_VALUE", "INVALID_OPERATION", "OUT_OF_MEMORY", "INVALID_FRAMEBUFFER_OPERATION", "STACK_OVERFLOW", "STACK_UNDERFLOW")
	}
	src.addLn(`
//	Returns the name of the specified Enum. Not recommended in a real-time loop.
func (_ GlUtil) EnumName(enum Enum) (name string) {
	switch enum {`)
	for _, enum := range me.enums {
		if isEarlyEnum(enum.name) {
			enumsDone[toUi(enum.val)] = true
			src.addLn("\tcase %s:\n\t\tname = \"GL_%s\"", enum.name, enum.name)
		}
	}
	for _, enum := range me.enums {
		if ui := toUi(enum.val); !enumsDone[ui] {
			enumsDone[ui] = true
			src.addLn("\tcase %s:\n\t\tname = \"GL_%s\"", enum.name, enum.name)
		}
	}
	src.addLn(`	default:
		name = fmt.Sprintf("GL_ENUM_%%v", enum)
	}
	return
}
`)
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
