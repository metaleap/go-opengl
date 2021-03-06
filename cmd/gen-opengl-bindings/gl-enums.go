package main

import (
	"strconv"
	"strings"

	xmlx "github.com/go-forks/go-pkg-xmlx"
	"github.com/metaleap/go-util/slice"
	"github.com/metaleap/go-util/str"
)

type glEnum struct {
	legacy              bool
	name, ref, val, ver string
	exts                []string
}

func (me *glEnum) finalVal() (val string) {
	if val = me.val; len(val) == 0 && len(me.ref) > 0 {
		val = allEnums[me.ref].finalVal()
	}
	return
}

func (me *glPack) makeEnums() {
	var src glPackSrc
	if *flagTry {
		src.addLn("import \"fmt\"")
	}
	for name, enum := range allEnums {
		if (!enum.legacy) && (len(enum.exts) == 0 || len(enum.ver) > 0 || uslice.StrHasAny(cfg.genExts, enum.exts...)) {
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
	if *flagTry {
		src.addLn(`
//	Returns the name of the specified Enum. Not recommended in a real-time loop.
func (_ GlUtil) EnumName(enum Enum) (name string) {
	switch enum {`)
		for _, enum := range me.enums {
			if isEarlyEnum(enum.name) {
				enumsDone[toUi(enum.finalVal())] = true
				src.addLn("\tcase %s:\n\t\tname = \"GL_%s\"", enum.name, enum.name)
			}
		}
		for _, enum := range me.enums {
			if ui := toUi(enum.finalVal()); !enumsDone[ui] {
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
	}
	me.sources["enums"] = src
}

func xmlWalkEnums() {
	total, num := 0, 0
	xmlWalkDoc("enum", func(xn *xmlx.Node) {
		total++
		// checkForUnknownAtts(xn, "name", "version", "value", "deprecated", "removed", "ref")
		enum := &glEnum{name: xas(xn, "name"), ref: xas(xn, "ref"), ver: xas(xn, "version"), val: xas(xn, "value"), legacy: isLegacy(xn)}
		if !enum.legacy {
			num++
		}
		if len(enum.ver) == 0 {
			for _, extNode := range xn.SelectNodes(xmlns, "ext") {
				if extName := xas(extNode, "name"); len(extName) > 0 {
					enum.exts = append(enum.exts, extName)
				}
			}
		}
		allEnums[enum.name] = enum
	})
	println(sfmt("Picked %v/%v GL enums.", num, total))
}
