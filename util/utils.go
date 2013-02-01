package glutil

import (
	gl "github.com/go3d/go-opengl/gogl"
	ustr "github.com/metaleap/go-util/str"
)

//	A singleton type, only used for the package-global glutil.Gl variable.
//	Provides access to miscellaneous OpenGL / GoGL functionality.
type Utils struct {
}

//	Returns an OpenGL connection-info string of the format "OpenGL {version} @ {vendor} {renderer} (GLSL: {version})"
//	Example: "OpenGL 4.2.0 @ NVIDIA Corporation Quadro 5010M/PCIe/SSE2 (GLSL: 4.20 NVIDIA via Cg compiler)"
func (_ Utils) ConnInfo() string {
	return sfmt("OpenGL %v @ %v %v (GLSL: %v)", Gl.Str(gl.VERSION), Gl.Str(gl.VENDOR), Gl.Str(gl.RENDERER), Gl.Str(gl.SHADING_LANGUAGE_VERSION))
}

//	Returns true if the specified extension (ignoring upper/lower-case) is supported by the current OpenGL context, as per the global Support.Extensions variable.
//	If the extension uses a well-known prefix contained in KnownExtensionPrefixes, it can be omitted, ie. you can just specify "texture_filter_anisotropic" instead of "GL_EXT_texture_filter_anisotropic".
func (_ Utils) Extension(name string) bool {
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

//	Returns the specified OpenGL string.
//	Example: GlStr(gl.VENDOR) = "NVIDIA Corporation"
func (_ Utils) Str(name gl.Enum) string {
	return gl.GoStringUb(gl.GetString(name))
}

//	Returns the specified OpenGL string at the specified index.
//	For example, GlStri(gl.EXTENSIONS, 0) returns the name of the "first" of all supported extensions.
func (_ Utils) Stri(name gl.Enum, i gl.Uint) string {
	return gl.GoStringUb(gl.GetStringi(name, i))
}

//	Returns the specified integer as returned by gl.GetIntegerv().
//	Example: numExts := GlVal(gl.NUM_EXTENSIONS)
func (_ Utils) Val(name gl.Enum) (val gl.Int) {
	gl.GetIntegerv(name, &val)
	return
}

//	Returns the specified n integers as returned by gl.GetIntegerv().
func (_ Utils) Vals(name gl.Enum, n uint) (vals []gl.Int) {
	vals = make([]gl.Int, n)
	gl.GetIntegerv(name, &vals[0])
	return
}
