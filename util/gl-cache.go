package ugl

import (
	gl "github.com/go3d/go-opengl/core"
)

type GlCache struct {
	activeTexture   gl.Enum
	bindSampler     []gl.Uint
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
