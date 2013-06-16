package main

import (
	"strings"

	xmlx "github.com/goforks/xmlx"
)

type glTypeRef struct {
	c, g string
}

func (me *glTypeRef) set(ref string) {
	if me.c = allTypes[ref]; me.c == "*" {
		me.c = "void"
	}
	me.g = goType(me.c)
}

var (
	typeMapCGo = map[string]string{
		"GLDEBUGPROC": "Ptr",
	}
	typeMapGoC = map[string]string{
		"Enum":     "C.GLenum",
		"Boolean":  "C.GLboolean",
		"Bitfield": "C.GLbitfield",
		"Byte":     "C.GLbyte",
		"Short":    "C.GLshort",
		"Int":      "C.GLint",
		"Sizei":    "C.GLsizei",
		"Ubyte":    "C.GLubyte",
		"Ushort":   "C.GLushort",
		"Uint":     "C.GLuint",
		"Half":     "C.GLhalf",
		"Float":    "C.GLfloat",
		"Clampf":   "C.GLclampf",
		"Double":   "C.GLdouble",
		"Clampd":   "C.GLclampd",
		"Char":     "C.GLchar",
		"Sync":     "C.GLsync",
		"Int64":    "C.GLint64",
		"Uint64":   "C.GLuint64",
		"Intptr":   "C.GLintptr",
		"Sizeiptr": "C.GLsizeiptr",
		"Ptr":      "unsafe.Pointer",
	}
)

func init() {
	for gt, ct := range typeMapGoC {
		if strings.HasPrefix(ct, "C.") {
			typeMapCGo[ct[2:]] = gt
		} else {
			typeMapCGo[ct] = gt
		}
	}
}

func goType(ctype string) (gt string) {
	if gt = typeMapCGo[ctype]; len(gt) == 0 {
		switch ctype {
		case "GLvoid", "void*":
			gt = "Ptr"
		case "GLvoid*", "GLvoid* const":
			gt = "*Ptr"
		case "GLchar* const":
			gt = "**Char"
		case "const GLubyte *":
			gt = "*Ubyte"
		case "void":
			return ""
		default:
			gt = ctype
		}
	}
	return
}

func xmlWalkTypes() {
	xmlWalkDoc("type-def", func(xn *xmlx.Node) {
		// checkForUnknownAtts(xn, "typename", "C-lang")
		allTypes[xas(xn, "typename")] = xas(xn, "C-lang")
	})
	println(sfmt("Parsed %v GL type-defs.", len(allTypes)))
}
