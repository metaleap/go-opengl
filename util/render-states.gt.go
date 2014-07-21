package ugl

import (
	gl "github.com/go3d/go-opengl/core"
)

//	Encapsulates a particular combination of render-states that RenderStates can Apply().
type RenderStatesBag struct {
	Blending    bool
	DepthTest   bool
	FaceCulling bool
	StencilTest bool
	ClearColor  GlVec4
	Other       struct {
		FramebufferSrgb bool
		ClearBits       gl.Bitfield
	}
}

//	Encapsulates various states in the OpenGL state machine, also providing toggles for some.
//	
//	RenderStates is designed to facilitate lazy state changes by caching all render states:
//	so unless you call any of the ForceFoo() methods, state changes are only propagated to
//	OpenGL if there is an actual change from the last time that state was set (via any
//	RenderStates instance -- being 'singleton' in design, they all utilize the same backing cache).
type RenderStates struct {
}

//	Calls the applicable me.SetFoo() methods for all render states in bag (except those in bag.Other).
func (me RenderStates) Apply(bag *RenderStatesBag) {
	me.SetBlending(bag.Blending)
	me.SetDepthTest(bag.DepthTest)
	me.SetFaceCulling(bag.FaceCulling)
	me.SetStencilTest(bag.StencilTest)
	me.SetClearColor(bag.ClearColor)
}

//	Sets the OpenGL clear-color to the specified rgba.
func (_ RenderStates) ForceClearColor(rgba GlVec4) {
	Cache.renderStates.ClearColor = rgba
	gl.ClearColor(Cache.renderStates.ClearColor[0], Cache.renderStates.ClearColor[1], Cache.renderStates.ClearColor[2], Cache.renderStates.ClearColor[3])
}

//	Sets the OpenGL clear-color if rgba differs from the currently cached value.
func (me RenderStates) SetClearColor(rgba GlVec4) {
	if Cache.renderStates.ClearColor != rgba {
		me.ForceClearColor(rgba)
	}
}



//#begin-gt -gen-renderstates.gt GT_MULT_SEP:, N:Blending,DepthTest,FaceCulling,FramebufferSrgb,ScissorTest,StencilTest D:blending,depth-testing,face-culling,SRGB-framebuffer,scissor-testing,stencil-testing E:BLEND,DEPTH_TEST,CULL_FACE,FRAMEBUFFER_SRGB,SCISSOR_TEST,STENCIL_TEST

//	Disables blending only if it is currently enabled.
func (me RenderStates) DisableBlending() {
	if Cache.renderStates.Blending {
		Cache.renderStates.Blending = false
		gl.Disable(gl.BLEND)
	}
}

//	Enables blending only if it is currently disabled.
func (me RenderStates) EnableBlending() {
	if !Cache.renderStates.Blending {
		Cache.renderStates.Blending = true
		gl.Enable(gl.BLEND)
	}
}

//	Deactivates blending.
func (_ RenderStates) ForceDisableBlending() {
	Cache.renderStates.Blending = false
	gl.Disable(gl.BLEND)
}

//	Activates blending.
func (_ RenderStates) ForceEnableBlending() {
	Cache.renderStates.Blending = true
	gl.Enable(gl.BLEND)
}

//	Activates or deactivates blending.
func (me RenderStates) SetBlending(newBlending bool) {
	if Cache.renderStates.Blending != newBlending {
		if Cache.renderStates.Blending = newBlending; Cache.renderStates.Blending {
			gl.Enable(gl.BLEND)
		} else {
			gl.Disable(gl.BLEND)
		}
	}
}

//	Toggles blending.
func (me RenderStates) ToggleBlending() {
	if Cache.renderStates.Blending = !Cache.renderStates.Blending; Cache.renderStates.Blending {
		gl.Enable(gl.BLEND)
	} else {
		gl.Disable(gl.BLEND)
	}
}

//	Disables depth-testing only if it is currently enabled.
func (me RenderStates) DisableDepthTest() {
	if Cache.renderStates.DepthTest {
		Cache.renderStates.DepthTest = false
		gl.Disable(gl.DEPTH_TEST)
	}
}

//	Enables depth-testing only if it is currently disabled.
func (me RenderStates) EnableDepthTest() {
	if !Cache.renderStates.DepthTest {
		Cache.renderStates.DepthTest = true
		gl.Enable(gl.DEPTH_TEST)
	}
}

//	Deactivates depth-testing.
func (_ RenderStates) ForceDisableDepthTest() {
	Cache.renderStates.DepthTest = false
	gl.Disable(gl.DEPTH_TEST)
}

//	Activates depth-testing.
func (_ RenderStates) ForceEnableDepthTest() {
	Cache.renderStates.DepthTest = true
	gl.Enable(gl.DEPTH_TEST)
}

//	Activates or deactivates depth-testing.
func (me RenderStates) SetDepthTest(newDepthTest bool) {
	if Cache.renderStates.DepthTest != newDepthTest {
		if Cache.renderStates.DepthTest = newDepthTest; Cache.renderStates.DepthTest {
			gl.Enable(gl.DEPTH_TEST)
		} else {
			gl.Disable(gl.DEPTH_TEST)
		}
	}
}

//	Toggles depth-testing.
func (me RenderStates) ToggleDepthTest() {
	if Cache.renderStates.DepthTest = !Cache.renderStates.DepthTest; Cache.renderStates.DepthTest {
		gl.Enable(gl.DEPTH_TEST)
	} else {
		gl.Disable(gl.DEPTH_TEST)
	}
}

//	Disables face-culling only if it is currently enabled.
func (me RenderStates) DisableFaceCulling() {
	if Cache.renderStates.FaceCulling {
		Cache.renderStates.FaceCulling = false
		gl.Disable(gl.CULL_FACE)
	}
}

//	Enables face-culling only if it is currently disabled.
func (me RenderStates) EnableFaceCulling() {
	if !Cache.renderStates.FaceCulling {
		Cache.renderStates.FaceCulling = true
		gl.Enable(gl.CULL_FACE)
	}
}

//	Deactivates face-culling.
func (_ RenderStates) ForceDisableFaceCulling() {
	Cache.renderStates.FaceCulling = false
	gl.Disable(gl.CULL_FACE)
}

//	Activates face-culling.
func (_ RenderStates) ForceEnableFaceCulling() {
	Cache.renderStates.FaceCulling = true
	gl.Enable(gl.CULL_FACE)
}

//	Activates or deactivates face-culling.
func (me RenderStates) SetFaceCulling(newFaceCulling bool) {
	if Cache.renderStates.FaceCulling != newFaceCulling {
		if Cache.renderStates.FaceCulling = newFaceCulling; Cache.renderStates.FaceCulling {
			gl.Enable(gl.CULL_FACE)
		} else {
			gl.Disable(gl.CULL_FACE)
		}
	}
}

//	Toggles face-culling.
func (me RenderStates) ToggleFaceCulling() {
	if Cache.renderStates.FaceCulling = !Cache.renderStates.FaceCulling; Cache.renderStates.FaceCulling {
		gl.Enable(gl.CULL_FACE)
	} else {
		gl.Disable(gl.CULL_FACE)
	}
}

//	Disables SRGB-framebuffer only if it is currently enabled.
func (me RenderStates) DisableFramebufferSrgb() {
	if Cache.renderStates.FramebufferSrgb {
		Cache.renderStates.FramebufferSrgb = false
		gl.Disable(gl.FRAMEBUFFER_SRGB)
	}
}

//	Enables SRGB-framebuffer only if it is currently disabled.
func (me RenderStates) EnableFramebufferSrgb() {
	if !Cache.renderStates.FramebufferSrgb {
		Cache.renderStates.FramebufferSrgb = true
		gl.Enable(gl.FRAMEBUFFER_SRGB)
	}
}

//	Deactivates SRGB-framebuffer.
func (_ RenderStates) ForceDisableFramebufferSrgb() {
	Cache.renderStates.FramebufferSrgb = false
	gl.Disable(gl.FRAMEBUFFER_SRGB)
}

//	Activates SRGB-framebuffer.
func (_ RenderStates) ForceEnableFramebufferSrgb() {
	Cache.renderStates.FramebufferSrgb = true
	gl.Enable(gl.FRAMEBUFFER_SRGB)
}

//	Activates or deactivates SRGB-framebuffer.
func (me RenderStates) SetFramebufferSrgb(newFramebufferSrgb bool) {
	if Cache.renderStates.FramebufferSrgb != newFramebufferSrgb {
		if Cache.renderStates.FramebufferSrgb = newFramebufferSrgb; Cache.renderStates.FramebufferSrgb {
			gl.Enable(gl.FRAMEBUFFER_SRGB)
		} else {
			gl.Disable(gl.FRAMEBUFFER_SRGB)
		}
	}
}

//	Toggles SRGB-framebuffer.
func (me RenderStates) ToggleFramebufferSrgb() {
	if Cache.renderStates.FramebufferSrgb = !Cache.renderStates.FramebufferSrgb; Cache.renderStates.FramebufferSrgb {
		gl.Enable(gl.FRAMEBUFFER_SRGB)
	} else {
		gl.Disable(gl.FRAMEBUFFER_SRGB)
	}
}

//	Disables scissor-testing only if it is currently enabled.
func (me RenderStates) DisableScissorTest() {
	if Cache.renderStates.ScissorTest {
		Cache.renderStates.ScissorTest = false
		gl.Disable(gl.SCISSOR_TEST)
	}
}

//	Enables scissor-testing only if it is currently disabled.
func (me RenderStates) EnableScissorTest() {
	if !Cache.renderStates.ScissorTest {
		Cache.renderStates.ScissorTest = true
		gl.Enable(gl.SCISSOR_TEST)
	}
}

//	Deactivates scissor-testing.
func (_ RenderStates) ForceDisableScissorTest() {
	Cache.renderStates.ScissorTest = false
	gl.Disable(gl.SCISSOR_TEST)
}

//	Activates scissor-testing.
func (_ RenderStates) ForceEnableScissorTest() {
	Cache.renderStates.ScissorTest = true
	gl.Enable(gl.SCISSOR_TEST)
}

//	Activates or deactivates scissor-testing.
func (me RenderStates) SetScissorTest(newScissorTest bool) {
	if Cache.renderStates.ScissorTest != newScissorTest {
		if Cache.renderStates.ScissorTest = newScissorTest; Cache.renderStates.ScissorTest {
			gl.Enable(gl.SCISSOR_TEST)
		} else {
			gl.Disable(gl.SCISSOR_TEST)
		}
	}
}

//	Toggles scissor-testing.
func (me RenderStates) ToggleScissorTest() {
	if Cache.renderStates.ScissorTest = !Cache.renderStates.ScissorTest; Cache.renderStates.ScissorTest {
		gl.Enable(gl.SCISSOR_TEST)
	} else {
		gl.Disable(gl.SCISSOR_TEST)
	}
}

//	Disables stencil-testing only if it is currently enabled.
func (me RenderStates) DisableStencilTest() {
	if Cache.renderStates.StencilTest {
		Cache.renderStates.StencilTest = false
		gl.Disable(gl.STENCIL_TEST)
	}
}

//	Enables stencil-testing only if it is currently disabled.
func (me RenderStates) EnableStencilTest() {
	if !Cache.renderStates.StencilTest {
		Cache.renderStates.StencilTest = true
		gl.Enable(gl.STENCIL_TEST)
	}
}

//	Deactivates stencil-testing.
func (_ RenderStates) ForceDisableStencilTest() {
	Cache.renderStates.StencilTest = false
	gl.Disable(gl.STENCIL_TEST)
}

//	Activates stencil-testing.
func (_ RenderStates) ForceEnableStencilTest() {
	Cache.renderStates.StencilTest = true
	gl.Enable(gl.STENCIL_TEST)
}

//	Activates or deactivates stencil-testing.
func (me RenderStates) SetStencilTest(newStencilTest bool) {
	if Cache.renderStates.StencilTest != newStencilTest {
		if Cache.renderStates.StencilTest = newStencilTest; Cache.renderStates.StencilTest {
			gl.Enable(gl.STENCIL_TEST)
		} else {
			gl.Disable(gl.STENCIL_TEST)
		}
	}
}

//	Toggles stencil-testing.
func (me RenderStates) ToggleStencilTest() {
	if Cache.renderStates.StencilTest = !Cache.renderStates.StencilTest; Cache.renderStates.StencilTest {
		gl.Enable(gl.STENCIL_TEST)
	} else {
		gl.Disable(gl.STENCIL_TEST)
	}
}

//#end-gt
