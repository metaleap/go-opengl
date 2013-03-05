package ugl

import (
	gl "github.com/go3d/go-opengl/core"
)

type GlCache struct {
	activeTexture   gl.Enum
	bindSampler     []gl.Uint
	bindTexture     []gl.Uint
	bindVertexArray gl.Uint
	useProgram      gl.Uint
	renderStates    struct {
		Blending, DepthTest, FaceCulling, FramebufferSrgb, ScissorTest, StencilTest bool
		ClearColor                                                                  GlVec4
	}
}

func init() {
	Cache.activeTexture = gl.TEXTURE0
}

func (_ *GlCache) onGlInit() {
	// texTargets := [...]gl.Enum{gl.TEXTURE_2D, gl.TEXTURE_CUBE_MAP, gl.TEXTURE_2D_ARRAY, gl.TEXTURE_1D, gl.TEXTURE_3D, gl.TEXTURE_1D_ARRAY, gl.TEXTURE_BUFFER, gl.TEXTURE_RECTANGLE, gl.TEXTURE_CUBE_MAP_ARRAY, gl.TEXTURE_2D_MULTISAMPLE, gl.TEXTURE_2D_MULTISAMPLE_ARRAY}
	Cache.bindSampler = make([]gl.Uint, Support.Textures.ImageUnits.MaxFragment)
	Cache.bindTexture = make([]gl.Uint, Support.Textures.ImageUnits.MaxCombined)
}

func (_ *GlCache) ActiveTexture(texture gl.Enum) {
	if texture != Cache.activeTexture {
		Cache.activeTexture = texture
		gl.ActiveTexture(texture)
	}
}

func (_ *GlCache) BindSampler(textureUnit, glHandle gl.Uint) {
	if Cache.bindSampler[textureUnit] != glHandle {
		Cache.bindSampler[textureUnit] = glHandle
		gl.BindSampler(textureUnit, glHandle)
	}
}

func (_ *GlCache) BindTextureTo(textureUnit, glHandle gl.Uint, glTarget gl.Enum) {
	if Cache.bindTexture[textureUnit] != glHandle {
		Cache.bindTexture[textureUnit] = glHandle
		Cache.ActiveTexture(gl.Enum(gl.TEXTURE0 + textureUnit))
		gl.BindTexture(glTarget, glHandle)
	}
}

func (_ *GlCache) BindVertexArray(glHandle gl.Uint) {
	if Cache.bindVertexArray != glHandle {
		Cache.bindVertexArray = glHandle
		gl.BindVertexArray(glHandle)
	}
}

func (_ *GlCache) UseProgram(glHandle gl.Uint) {
	if Cache.useProgram != glHandle {
		Cache.useProgram = glHandle
		gl.UseProgram(glHandle)
	}
}
