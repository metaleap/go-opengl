package ugl

import (
	gl "github.com/metaleap/go-opengl/core"
)

//	A singleton type, only used for the package-global ugl.Try variable.
//	(Methods generated with github.com/metaleap/go-opengl/cmd/gen-opengl-bindings.)
type GlTry struct {
}

//	Calls gl.AttachShader() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) AttachShader(program gl.Uint, shader gl.Uint) (err error) {
	gl.AttachShader(program, shader)
	err = Util.LastError("gl.AttachShader()")
	return
}

//	Calls gl.BufferData() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) BufferData(target gl.Enum, size gl.Sizeiptr, data gl.Ptr, usage gl.Enum) (err error) {
	gl.BufferData(target, size, data, usage)
	err = Util.LastError("gl.BufferData()")
	return
}

//	Calls gl.BufferSubData() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) BufferSubData(target gl.Enum, offset gl.Intptr, size gl.Sizeiptr, data gl.Ptr) (err error) {
	gl.BufferSubData(target, offset, size, data)
	err = Util.LastError("gl.BufferSubData()")
	return
}

//	Calls gl.CreateProgram() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) CreateProgram() (err error, tryRetVal__ gl.Uint) {
	tryRetVal__ = gl.CreateProgram()
	err = Util.LastError("gl.CreateProgram()")
	return
}

//	Calls gl.CreateShader() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) CreateShader(type_ gl.Enum) (err error, tryRetVal__ gl.Uint) {
	tryRetVal__ = gl.CreateShader(type_)
	err = Util.LastError("gl.CreateShader()")
	return
}

//	Calls gl.GenBuffers() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) GenBuffers(n gl.Sizei, buffers *gl.Uint) (err error) {
	gl.GenBuffers(n, buffers)
	err = Util.LastError("gl.GenBuffers()")
	return
}

//	Calls gl.GenerateMipmap() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) GenerateMipmap(target gl.Enum) (err error) {
	gl.GenerateMipmap(target)
	err = Util.LastError("gl.GenerateMipmap()")
	return
}

//	Calls gl.GenTextures() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) GenTextures(n gl.Sizei, textures *gl.Uint) (err error) {
	gl.GenTextures(n, textures)
	err = Util.LastError("gl.GenTextures()")
	return
}

//	Calls gl.GenVertexArrays() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) GenVertexArrays(n gl.Sizei, arrays *gl.Uint) (err error) {
	gl.GenVertexArrays(n, arrays)
	err = Util.LastError("gl.GenVertexArrays()")
	return
}

//	Calls gl.ShaderSource() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) ShaderSource(shader gl.Uint, count gl.Sizei, string_ **gl.Char, length *gl.Int) (err error) {
	gl.ShaderSource(shader, count, string_, length)
	err = Util.LastError("gl.ShaderSource()")
	return
}

//	Calls gl.TexImage2D() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) TexImage2D(target gl.Enum, level gl.Int, internalformat gl.Int, width gl.Sizei, height gl.Sizei, border gl.Int, format gl.Enum, type_ gl.Enum, pixels gl.Ptr) (err error) {
	gl.TexImage2D(target, level, internalformat, width, height, border, format, type_, pixels)
	err = Util.LastError("gl.TexImage2D()")
	return
}

//	Calls gl.TexStorage2D() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) TexStorage2D(target gl.Enum, levels gl.Sizei, internalformat gl.Enum, width gl.Sizei, height gl.Sizei) (err error) {
	gl.TexStorage2D(target, levels, internalformat, width, height)
	err = Util.LastError("gl.TexStorage2D()")
	return
}

//	Calls gl.TexSubImage2D() and yields the err returned by Util.LastError(), if any.
func (_ GlTry) TexSubImage2D(target gl.Enum, level gl.Int, xoffset gl.Int, yoffset gl.Int, width gl.Sizei, height gl.Sizei, format gl.Enum, type_ gl.Enum, pixels gl.Ptr) (err error) {
	gl.TexSubImage2D(target, level, xoffset, yoffset, width, height, format, type_, pixels)
	err = Util.LastError("gl.TexSubImage2D()")
	return
}
