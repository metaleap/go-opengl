package ugl

import (
	gl "github.com/metaleap/go-opengl/core"
)

//	Returns the index of the lowest defined mipmap level. Defaults to 0.
func (me *TextureBase) ParamBaseLevel() gl.Int {
	return me.ParamInt(gl.TEXTURE_BASE_LEVEL)
}

//	Sets the index of the lowest defined mipmap level.
func (me *TextureBase) SetParamBaseLevel(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_BASE_LEVEL, val)
}

//	Returns the index of the highest defined mipmap level. Defaults to 1000.
func (me *TextureBase) ParamMaxLevel() gl.Int {
	return me.ParamInt(gl.TEXTURE_MAX_LEVEL)
}

//	Sets the index of the highest defined mipmap level.
func (me *TextureBase) SetParamMaxLevel(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_MAX_LEVEL, val)
}

//	Returns the swizzle that will be applied to the A component of a texel before it is returned to the shader. Defaults to gl.ALPHA.
func (me *TextureBase) ParamSwizzleA() gl.Int {
	return me.ParamInt(gl.TEXTURE_SWIZZLE_A)
}

//	Sets the swizzle that will be applied to the A component of a texel before it is returned to the shader.
func (me *TextureBase) SetParamSwizzleA(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_SWIZZLE_A, val)
}

//	Returns the swizzle that will be applied to the B component of a texel before it is returned to the shader. Defaults to gl.BLUE.
func (me *TextureBase) ParamSwizzleB() gl.Int {
	return me.ParamInt(gl.TEXTURE_SWIZZLE_B)
}

//	Sets the swizzle that will be applied to the B component of a texel before it is returned to the shader.
func (me *TextureBase) SetParamSwizzleB(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_SWIZZLE_B, val)
}

//	Returns the swizzle that will be applied to the G component of a texel before it is returned to the shader. Defaults to gl.GREEN.
func (me *TextureBase) ParamSwizzleG() gl.Int {
	return me.ParamInt(gl.TEXTURE_SWIZZLE_G)
}

//	Sets the swizzle that will be applied to the G component of a texel before it is returned to the shader.
func (me *TextureBase) SetParamSwizzleG(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_SWIZZLE_G, val)
}

//	Returns the swizzle that will be applied to the R component of a texel before it is returned to the shader. Defaults to gl.RED.
func (me *TextureBase) ParamSwizzleR() gl.Int {
	return me.ParamInt(gl.TEXTURE_SWIZZLE_R)
}

//	Sets the swizzle that will be applied to the R component of a texel before it is returned to the shader.
func (me *TextureBase) SetParamSwizzleR(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_SWIZZLE_R, val)
}

//	Returns the swizzle that will be applied to the R, G, B and A components of a texel before it is returned to the shader.
func (me *TextureBase) ParamSwizzleRgba() gl.Int {
	return me.ParamInt(gl.TEXTURE_SWIZZLE_RGBA)
}

//	Sets the swizzle that will be applied to the R, G, B and A components of a texel before it is returned to the shader.
func (me *TextureBase) SetParamSwizzleRgba(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_SWIZZLE_RGBA, val)
}

//	A convenience short-hand to invoke gl.GetTexParameterfv for the texture object currently bound to me.GlTarget.
func (me *TextureBase) ParamFloat(param gl.Enum) (val gl.Float) {
	gl.GetTexParameterfv(me.GlTarget, param, &val)
	return
}

//	A convenience short-hand to invoke gl.TexParameterf for the texture object currently bound to me.GlTarget.
func (me *TextureBase) SetParamFloat(param gl.Enum, val gl.Float) {
	gl.TexParameterf(me.GlTarget, param, val)
}

//	A convenience short-hand to invoke gl.GetTexParameteriv for the texture object currently bound to me.GlTarget.
func (me *TextureBase) ParamInt(param gl.Enum) (val gl.Int) {
	gl.GetTexParameteriv(me.GlTarget, param, &val)
	return
}

//	A convenience short-hand to invoke gl.TexParameteri for the texture object currently bound to me.GlTarget.
func (me *TextureBase) SetParamInt(param gl.Enum, val gl.Int) {
	gl.TexParameteri(me.GlTarget, param, val)
}
