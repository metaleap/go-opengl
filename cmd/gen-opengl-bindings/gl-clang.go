package main

import (
	"strings"
)

func (me *glPack) makeCgo() {
	var src glPackSrc
	src.addLn(`
// #cgo darwin  LDFLAGS: -framework OpenGL
// #cgo linux   LDFLAGS: -lGL
// #cgo windows LDFLAGS: -lopengl32
// 
// #include <stdlib.h>
// #if defined(__APPLE__)
// #include <dlfcn.h>
// #elif defined(_WIN32)
// #define WIN32_LEAN_AND_MEAN 1
// #include <windows.h>
// #else
// #include <X11/Xlib.h>
// #include <GL/glx.h>
// #endif
// 
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
// #ifndef GLAPI
// #define GLAPI extern
// #endif
// 
// typedef unsigned int GLenum;
// typedef unsigned char GLboolean;
// typedef unsigned int GLbitfield;
// typedef signed char GLbyte;
// typedef short GLshort;
// typedef int GLint;
// typedef int GLsizei;
// typedef unsigned char GLubyte;
// typedef unsigned short GLushort;
// typedef unsigned int GLuint;
// typedef unsigned short GLhalf;
// typedef float GLfloat;
// typedef float GLclampf;
// typedef double GLdouble;
// typedef double GLclampd;
// typedef void GLvoid;
// `)
	for _, ptNode := range specDoc.SelectNodesRecursive(xmlns, "passthru") {
		for _, ln := range strings.Split(ptNode.Value, "\n") {
			src.addLn("// %s", strings.Trim(ln, "\r"))
		}
	}
	src.addLn(`// 
// #ifdef _WIN32
// static HMODULE opengl32 = NULL;
// #endif
// 
// static void* coreGetProcAddress(const char* name) { 
// #ifdef __APPLE__
// 	return dlsym(RTLD_DEFAULT, name);
// #elif _WIN32
// 	void* pf = wglGetProcAddress((LPCSTR)name);
// 	if(pf) {
// 		return pf;
// 	}
// 	if(opengl32 == NULL) {
// 		opengl32 = LoadLibraryA("opengl32.dll");
// 	}
// 	return GetProcAddress(opengl32, (LPCSTR)name);
// #else
// 	return glXGetProcAddress((const GLubyte*)name);
// #endif
// }`)
	src.addLn("// ")
	for _, fun := range me.funcs {
		src.addLn("// %s", fun.clang.entry)
	}
	src.addLn("// ")
	for _, fun := range me.funcs {
		src.addLn("// %s", fun.clang.wrapper)
		if *flagSupports {
			src.addLn("// %s", fun.clang.supporter)
		}
	}
	src.addLn("// int coreGetProcAddresses() {")
	for _, fun := range me.funcs {
		src.addLn("// \tptrgl%s = coreGetProcAddress(\"gl%s\");", fun.name, fun.name)
	}
	src.addLn("// /* a cheap way to determine if (at least) v%s core profile (or higher) seems to be supported: */", *flagMinVer)
	src.addLn("// \treturn ((ptrgl%s == NULL) ? 0 : 1);", *flagProcAddr)
	src.addLn("// }")
	src.addLn(`import "C"
import "unsafe"

type (`)
	for gt, ct := range typeMapGoC {
		src.addLn("\t%s %s", gt, ct)
	}
	src.addLn(")")
	if *flagSupports {
		src.addLn(`
//	Provides methods indicating whether a given GL function can be called within the current GL context.
type GlSupport struct {}

//	Provides methods indicating whether a given GL function can be called within the current GL context.
var Supports GlSupport
`)
	}
	if *flagTry {
		src.addLn(`
//	Provides method wrappers that call GL functions and check for any errors immediately afterwards.
//	Not recommended for use in a real-time loop, but useful for "dirty debugging" or one-off/setup tasks.
type GlTry struct {}

//	Provides method wrappers that call GL functions and check for any errors immediately afterwards.
//	Not recommended for use in a real-time loop, but useful for "dirty debugging" or one-off/setup tasks.
var Try GlTry
`)
	}
	src.addLn(`
//	Provides utility methods for working with this OpenGL binding package.
type GlUtil struct {}

//	Provides utility methods for working with this OpenGL binding package.
var Util GlUtil

//	Go string to GL string.
func (_ GlUtil) CString(str string) *Char {
	return (*Char)(C.CString(str))
}

//	Allocates a GL string.
func (_ GlUtil) CStringAlloc(length Sizei) *Char {
	return (*Char)(C.malloc(C.size_t(length)))
}

//	Frees GL string.
func (_ GlUtil) CStringFree(str *Char) {
	C.free(unsafe.Pointer(str))
}

//	GL string (GLchar*) to Go string.
func (_ GlUtil) StringFromChar(str *Char) string {
	return C.GoString((*C.char)(str))
}

//	GL string (GLubyte*) to Go string.
func (_ GlUtil) StringFromUbyte(str *Ubyte) string {
	return C.GoString((*C.char)(unsafe.Pointer(str)))
}

//	GL string (GLchar*) with length to Go string.
func (_ GlUtil) StringN(str *Char, length Sizei) string {
	return C.GoStringN((*C.char)(str), C.int(length))
}

//	Converts a list of Go strings to a slice of GL strings.
func (_ GlUtil) CStringArray(strs ...string) []*Char {
	strSlice := make([]*Char, len(strs))
	for i, s := range strs {
		strSlice[i] = (*Char)(C.CString(s))
	}
	return strSlice
}

//	Free GL string slice allocated by GLStringArray().
func (_ GlUtil) CStringArrayFree(strs []*Char) {
	for _, s := range strs {
		C.free(unsafe.Pointer(s))
	}
}

//	Add offset to a pointer. Useful for VertexAttribPointer and friends.
func (_ GlUtil) PtrOffset(p Ptr, o uintptr) Ptr {
	return Ptr(uintptr(p) + o)
}

//	Returns all error flags collected from GetError() until the latter yields NO_ERROR.
func (_ GlUtil) ErrorFlags() (flags []Enum) {
	for e := GetError(); e != NO_ERROR; e = GetError() {
		flags = append(flags, e)
	}
	return
}

//	Initializes this OpenGL binding after having created your GL context (with SDL, GLFW or in some other way).
func (_ GlUtil) Init() bool {
	return (C.coreGetProcAddresses() == 1)
}
`)
	if *flagTry {
		src.addLn(`
import "fmt"

//	If any errors have been recorded since the last (direct or indirect) GetError() call, returns a single error value informing on those and including the specified message, if any.
func (_ GlUtil) Error(msgFmt string, fmtArgs ...interface{}) (err error) {
	if flags := Util.ErrorFlags(); len(flags) > 0 {
		for _, e := range flags {
			msgFmt = Util.EnumName(e) + " " + msgFmt
		}
		err = fmt.Errorf(msgFmt, fmtArgs...)
	}
	return
}
`)
	}
	if *flagSupports {
		for _, fun := range me.funcs {
			src.addLn("//\tReturns whether the current GL context supports calling the %s() function.", fun.name)
			src.addLn(`func (_ GlSupport) %s () bool {
	return C.cancall%s() == 1
}
`, fun.name, fun.name)
		}
	}
	if *flagTry {
		for _, fun := range me.funcs {
			me.makeTryFunc(&src, fun, "", "Util.")
		}
	}
	me.sources["funcs"] = src + me.sources["funcs"]
}
