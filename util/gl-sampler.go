package glutil

import (
	gl "github.com/go3d/go-opengl/gogl"
)

const (
	texture_max_anisotropy_ext     = 0x84FE
	max_texture_max_anisotropy_ext = 0x84FF
)

//	Stores sampling parameters for texture accesses.
type Sampler struct {
	//	The OpenGL handle to this sampler object.
	//	This is 0 before calling Create() and after calling Dispose().
	GlHandle gl.Uint
}

//	Binds this sampler object to the specified textureUnit.
func (me Sampler) Bind(textureUnit gl.Uint) {
	gl.BindSampler(textureUnit, me.GlHandle)
}

//	Creates this sampler object.
func (me *Sampler) Create() *Sampler {
	me.Dispose()
	gl.GenSamplers(1, &me.GlHandle)
	return me
}

//	Disables linear and anisotropic filtering on this sampler.
//	If allowMip is true, minified sampling still blends between MIP map levels (just without linear filtering).
func (me *Sampler) DisableAllFiltering(allowMip bool) *Sampler {
	me.SetFilterMag(gl.NEAREST)
	me.SetFilterMin(Typed.Ifi(allowMip, gl.NEAREST_MIPMAP_NEAREST, gl.NEAREST))
	me.SetFilterMaxAnisotropy(1)
	return me
}

//	Deletes this sampler object.
func (me *Sampler) Dispose() {
	if me.GlHandle != 0 {
		gl.DeleteSamplers(1, &me.GlHandle)
		me.GlHandle = 0
	}
}

//	Enables linear and anisotropic filtering on this sampler.
//	If allowMip is true, minified sampling blends between MIP map levels linearly filtered.
func (me *Sampler) EnableFullFiltering(allowMip bool, maxAniso gl.Float) *Sampler {
	me.SetFilterMag(gl.LINEAR)
	me.SetFilterMin(Typed.Ifi(allowMip, gl.LINEAR_MIPMAP_LINEAR, gl.LINEAR))
	me.SetFilterMaxAnisotropy(maxAniso)
	return me
}

//	A convenience short-hand to invoke gl.GetSamplerParameterfv for this sampler object.
func (me Sampler) ParamFloat(param gl.Enum) (val gl.Float) {
	gl.GetSamplerParameterfv(me.GlHandle, param, &val)
	return
}

//	A convenience short-hand to invoke gl.SamplerParameterf for this sampler object.
func (me Sampler) SetParamFloat(param gl.Enum, val gl.Float) {
	gl.SamplerParameterf(me.GlHandle, param, val)
}

//	A convenience short-hand to invoke gl.GetSamplerParameteriv for this sampler object.
func (me Sampler) ParamInt(param gl.Enum) (val gl.Int) {
	gl.GetSamplerParameteriv(me.GlHandle, param, &val)
	return
}

//	A convenience short-hand to invoke gl.SamplerParameteri for this sampler object.
func (me Sampler) SetParamInt(param gl.Enum, val gl.Int) {
	gl.SamplerParameteri(me.GlHandle, param, val)
}

//	Unbinds whatever sampler object is currently bound to the specified textureUnit.
func (_ Sampler) Unbind(textureUnit gl.Uint) {
	gl.BindSampler(textureUnit, 0)
}

//	Returns a single-valued texture comparison function, a symbolic constant. The initial value is GL_LEQUAL.
func (me Sampler) DepthCompareFunc() gl.Int {
	return me.ParamInt(gl.TEXTURE_COMPARE_FUNC)
}

//	Specifies the comparison operator used when DepthCompareMode() is set to gl.COMPARE_REF_TO_TEXTURE.
//	val may be one of: gl.LEQUAL, gl.GEQUAL, gl.LESS, gl.GREATER, gl.EQUAL, gl.NOTEQUAL, gl.ALWAYS, gl.NEVER.
func (me Sampler) SetDepthCompareFunc(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_COMPARE_FUNC, val)
}

//	Returns a single-valued texture comparison mode, a symbolic constant. The initial value is gl.NONE.
func (me Sampler) DepthCompareMode() gl.Int {
	return me.ParamInt(gl.TEXTURE_COMPARE_MODE)
}

//	Specifies the texture comparison mode for currently bound textures whose internal format is gl.DEPTH_COMPONENT_*.
//	val may be either gl.NONE or gl.COMPARE_REF_TO_TEXTURE.
func (me Sampler) SetDepthCompareMode(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_COMPARE_MODE, val)
}

//	Returns the single-valued texture magnification filter, a symbolic constant. The initial value is GL_LINEAR.
func (me Sampler) FilterMag() gl.Int {
	return me.ParamInt(gl.TEXTURE_MAG_FILTER)
}

//	Sets the texture magnification function to either GL_NEAREST or GL_LINEAR.
func (me Sampler) SetFilterMag(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_MAG_FILTER, val)
}

//	Returns filtering anisotropy. 0 if anisotropic filtering isn't supported, 1 if anisotropic filtering
//	isn't activated for this Sampler, else a value between 1 and Support.Textures.MaxFilterAnisotropy.
func (me Sampler) FilterMaxAnisotropy() (val gl.Float) {
	if Support.Textures.MaxFilterAnisotropy > 0 {
		val = me.ParamFloat(texture_max_anisotropy_ext)
	}
	return
}

//	Sets filtering anisotropy, if supported. val will be clamped between 1 and Support.Textures.MaxFilterAnisotropy. Only a val greater than 1 will activate anisotropic filtering.
func (me Sampler) SetFilterMaxAnisotropy(val gl.Float) {
	if Support.Textures.MaxFilterAnisotropy > 0 {
		me.SetParamFloat(texture_max_anisotropy_ext, Typed.Clamp(val, 1, Support.Textures.MaxFilterAnisotropy))
	}
}

//	Returns the single-valued texture minification filter, a symbolic constant. The initial value is GL_NEAREST_MIPMAP_LINEAR.
func (me Sampler) FilterMin() gl.Int {
	return me.ParamInt(gl.TEXTURE_MIN_FILTER)
}

//	Sets the texture magnification function. val may be one of: GL_NEAREST, GL_LINEAR,
//	gl.NEAREST_MIPMAP_NEAREST, gl.LINEAR_MIPMAP_NEAREST, gl.NEAREST_MIPMAP_LINEAR, gl.LINEAR_MIPMAP_LINEAR.
func (me Sampler) SetFilterMin(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_MIN_FILTER, val)
}

//	The mipmap image selection process can be adjusted coarsely by using this sampling parameter.
func (me Sampler) MipLodBias() gl.Float {
	return me.ParamFloat(gl.TEXTURE_LOD_BIAS)
}

//	This bias will be added to the mipmap LOD calculation (as well as added to the bias specified in one of the texture accessing functions in GLSL), which is used to select the image.
func (me Sampler) SetMipLodBias(val gl.Float) {
	me.SetParamFloat(gl.TEXTURE_LOD_BIAS, Typed.Clamp(val, -Support.Textures.MaxMipLoadBias, Support.Textures.MaxMipLoadBias))
}

//	Returns the single-valued texture maximum level-of-detail value. The initial value is 1000.
func (me Sampler) MipLodMax() gl.Float {
	return me.ParamFloat(gl.TEXTURE_MAX_LOD)
}

//	Sets the maximum level-of-detail parameter. This limits the selection of the lowest resolution mipmap (highest mipmap level).
func (me Sampler) SetMipLodMax(val gl.Float) {
	me.SetParamFloat(gl.TEXTURE_MAX_LOD, val)
}

//	Returns the single-valued texture minimum level-of-detail value. The initial value is -1000.
func (me Sampler) MipLodMin() gl.Float {
	return me.ParamFloat(gl.TEXTURE_MIN_LOD)
}

//	Sets the minimum level-of-detail parameter. This limits the selection of highest resolution mipmap (lowest mipmap level).
func (me Sampler) SetMipLodMin(val gl.Float) {
	me.SetParamFloat(gl.TEXTURE_MIN_LOD, val)
}

//	A convenience short-hand that calls SetWrapR(), SetWrapS() and SetWrapT() with the specified wrapVal.
func (me *Sampler) SetWrap(wrapVal gl.Int) *Sampler {
	me.SetWrapR(wrapVal)
	me.SetWrapS(wrapVal)
	me.SetWrapT(wrapVal)
	return me
}

//	Returns four floating-point numbers that comprise the RGBA color of the texture border.
//	The initial value is (0, 0, 0, 0).
func (me Sampler) WrapClampBorderColor() (val *GlVec4) {
	gl.GetSamplerParameterfv(me.GlHandle, gl.TEXTURE_BORDER_COLOR, &val[0])
	return
}

//	Specifies four values that define the border values that should be used for border texels.
func (me Sampler) SetWrapClampBorderColor(val *GlVec4) {
	gl.SamplerParameterfv(me.GlHandle, gl.TEXTURE_BORDER_COLOR, &val[0])
}

//	Returns the single-valued wrapping function for texture coordinate r,
//	a symbolic constant. The initial value is gl.REPEAT.
func (me Sampler) WrapR() gl.Int {
	return me.ParamInt(gl.TEXTURE_WRAP_R)
}

//	Sets the wrap parameter for texture coordinate r to
//	gl.CLAMP_TO_EDGE, gl.CLAMP_TO_BORDER, gl.MIRRORED_REPEAT, or gl.REPEAT.
func (me Sampler) SetWrapR(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_WRAP_R, val)
}

//	Returns the single-valued wrapping function for texture coordinate s,
//	a symbolic constant. The initial value is gl.REPEAT.
func (me Sampler) WrapS() gl.Int {
	return me.ParamInt(gl.TEXTURE_WRAP_S)
}

//	Sets the wrap parameter for texture coordinate s to
//	gl.CLAMP_TO_EDGE, gl.CLAMP_TO_BORDER, gl.MIRRORED_REPEAT, or gl.REPEAT.
func (me Sampler) SetWrapS(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_WRAP_S, val)
}

//	Returns the single-valued wrapping function for texture coordinate t,
//	a symbolic constant. The initial value is gl.REPEAT.
func (me Sampler) WrapT() gl.Int {
	return me.ParamInt(gl.TEXTURE_WRAP_T)
}

//	Sets the wrap parameter for texture coordinate t to
//	gl.CLAMP_TO_EDGE, gl.CLAMP_TO_BORDER, gl.MIRRORED_REPEAT, or gl.REPEAT.
func (me Sampler) SetWrapT(val gl.Int) {
	me.SetParamInt(gl.TEXTURE_WRAP_T, val)
}
