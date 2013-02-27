package ugl

import (
	gl "github.com/go3d/go-opengl/core"
	ustr "github.com/metaleap/go-util/str"
)

//	A singleton type, only used for the package-global ugl.Util variable.
type GlUtil struct {
}

//	Returns an OpenGL connection-info string of the format "OpenGL {version} @ {vendor} {renderer} (GLSL: {version})"
//	Example: "OpenGL 4.2.0 @ NVIDIA Corporation Quadro 5010M/PCIe/SSE2 (GLSL: 4.20 NVIDIA via Cg compiler)"
func (_ GlUtil) ConnInfo() string {
	return strf("OpenGL %v @ %v %v (GLSL: %v)", Util.Str(gl.VERSION), Util.Str(gl.VENDOR), Util.Str(gl.RENDERER), Util.Str(gl.SHADING_LANGUAGE_VERSION))
}

//	Returns the name of the specified enum. Out of the over 1290 possible enum values, currently only supports
//	error codes and shader stages.
func (_ GlUtil) EnumName(enum gl.Enum) (name string) {
	switch enum {
	case gl.INVALID_ENUM:
		name = "GL_INVALID_ENUM"
	case gl.INVALID_VALUE:
		name = "GL_INVALID_VALUE"
	case gl.INVALID_OPERATION:
		name = "GL_INVALID_OPERATION"
	case gl.INVALID_FRAMEBUFFER_OPERATION:
		name = "GL_INVALID_FRAMEBUFFER_OPERATION"
	case gl.OUT_OF_MEMORY:
		name = "GL_OUT_OF_MEMORY"
	case gl.STACK_OVERFLOW:
		name = "GL_STACK_OVERFLOW"
	case gl.STACK_UNDERFLOW:
		name = "GL_STACK_UNDERFLOW"
	case gl.COMPUTE_SHADER:
		name = "GL_COMPUTE_SHADER"
	case gl.FRAGMENT_SHADER:
		name = "GL_FRAGMENT_SHADER"
	case gl.GEOMETRY_SHADER:
		name = "GL_GEOMETRY_SHADER"
	case gl.TESS_CONTROL_SHADER:
		name = "GL_TESS_CONTROL_SHADER"
	case gl.TESS_EVALUATION_SHADER:
		name = "GL_TESS_EVALUATION_SHADER"
	case gl.VERTEX_SHADER:
		name = "GL_VERTEX_SHADER"
	default:
		name = strf("GL_ENUM_%v", enum)
	}
	return
}

//	Returns true if the specified extension (ignoring upper/lower-case) is supported by the current OpenGL context, as per the global Support.Extensions variable.
//	If the extension uses a well-known prefix contained in KnownExtensionPrefixes, it can be omitted, ie. you can just specify "texture_filter_anisotropic" instead of "GL_EXT_texture_filter_anisotropic".
func (_ GlUtil) Extension(name string) bool {
	if ustr.IsInSliceIgnoreCase(Support.Extensions, name) {
		return true
	}
	for _, pref := range KnownExtensionPrefixes {
		if ustr.IsInSliceIgnoreCase(Support.Extensions, pref+name) {
			return true
		}
	}
	return false
}

//	If the GL error flag is currently set, returns an error with the specified
//	message and the GL error flag(s).
func (_ GlUtil) LastError(msgFmt string, fmtArgs ...interface{}) (err error) {
	if flags := gl.Util.ErrorFlags(); len(flags) > 0 {
		for _, e := range flags {
			msgFmt = Util.EnumName(e) + " " + msgFmt
		}
		err = errf(msgFmt, fmtArgs...)
	}
	return
}

//	Returns the specified OpenGL string.
//	Example: GlStr(gl.VENDOR) = "NVIDIA Corporation"
func (_ GlUtil) Str(name gl.Enum) string {
	return gl.Util.StringFromUbyte(gl.GetString(name))
}

//	Returns the specified OpenGL string at the specified index.
//	For example, GlStri(gl.EXTENSIONS, 0) returns the name of the "first" of all supported extensions.
func (_ GlUtil) Stri(name gl.Enum, i gl.Uint) string {
	return gl.Util.StringFromUbyte(gl.GetStringi(name, i))
}

//	Returns the specified integer as returned by gl.GetIntegerv().
//	Example: numExts := GlVal(gl.NUM_EXTENSIONS)
func (_ GlUtil) Val(name gl.Enum) (val gl.Int) {
	gl.GetIntegerv(name, &val)
	return
}

//	Returns the specified n integers as returned by gl.GetIntegerv().
func (_ GlUtil) Vals(name gl.Enum, n uint) (vals []gl.Int) {
	vals = make([]gl.Int, n)
	gl.GetIntegerv(name, &vals[0])
	return
}
