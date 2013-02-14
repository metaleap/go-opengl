package ugl

import (
	gl "github.com/go3d/go-opengl/core"
)

var (
	rsClearColor GlVec4
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
	rsClearColor = rgba
	gl.ClearColor(rsClearColor[0], rsClearColor[1], rsClearColor[2], rsClearColor[3])
}

//	Sets the OpenGL clear-color if rgba differs from the currently cached value.
func (me RenderStates) SetClearColor(rgba GlVec4) {
	if rsClearColor != rgba {
		me.ForceClearColor(rgba)
	}
}

//#begin-gt -gen-renderstates.gt GT_MULT_SEP:, N:Blending,DepthTest,FaceCulling,FramebufferSrgb,ScissorTest,StencilTest D:blending,depth-testing,face-culling,SRGB-framebuffer,scissor-testing,stencil-testing E:BLEND,DEPTH_TEST,CULL_FACE,FRAMEBUFFER_SRGB,SCISSOR_TEST,STENCIL_TEST

var rsBlending bool

//	Disables blending only if it is currently enabled.
func (me RenderStates) DisableBlending() {
	if rsBlending {
		rsBlending = false
		gl.Disable(gl.BLEND)
	}
}

//	Enables blending only if it is currently disabled.
func (me RenderStates) EnableBlending() {
	if !rsBlending {
		rsBlending = true
		gl.Enable(gl.BLEND)
	}
}

//	Deactivates blending.
func (_ RenderStates) ForceDisableBlending() {
	rsBlending = false
	gl.Disable(gl.BLEND)
}

//	Activates blending.
func (_ RenderStates) ForceEnableBlending() {
	rsBlending = true
	gl.Enable(gl.BLEND)
}

//	Activates or deactivates blending.
func (me RenderStates) SetBlending(newBlending bool) {
	if rsBlending != newBlending {
		if rsBlending = newBlending; rsBlending {
			gl.Enable(gl.BLEND)
		} else {
			gl.Disable(gl.BLEND)
		}
	}
}

//	Toggles blending.
func (me RenderStates) ToggleBlending() {
	if rsBlending = !rsBlending; rsBlending {
		gl.Enable(gl.BLEND)
	} else {
		gl.Disable(gl.BLEND)
	}
}

var rsDepthTest bool

//	Disables depth-testing only if it is currently enabled.
func (me RenderStates) DisableDepthTest() {
	if rsDepthTest {
		rsDepthTest = false
		gl.Disable(gl.DEPTH_TEST)
	}
}

//	Enables depth-testing only if it is currently disabled.
func (me RenderStates) EnableDepthTest() {
	if !rsDepthTest {
		rsDepthTest = true
		gl.Enable(gl.DEPTH_TEST)
	}
}

//	Deactivates depth-testing.
func (_ RenderStates) ForceDisableDepthTest() {
	rsDepthTest = false
	gl.Disable(gl.DEPTH_TEST)
}

//	Activates depth-testing.
func (_ RenderStates) ForceEnableDepthTest() {
	rsDepthTest = true
	gl.Enable(gl.DEPTH_TEST)
}

//	Activates or deactivates depth-testing.
func (me RenderStates) SetDepthTest(newDepthTest bool) {
	if rsDepthTest != newDepthTest {
		if rsDepthTest = newDepthTest; rsDepthTest {
			gl.Enable(gl.DEPTH_TEST)
		} else {
			gl.Disable(gl.DEPTH_TEST)
		}
	}
}

//	Toggles depth-testing.
func (me RenderStates) ToggleDepthTest() {
	if rsDepthTest = !rsDepthTest; rsDepthTest {
		gl.Enable(gl.DEPTH_TEST)
	} else {
		gl.Disable(gl.DEPTH_TEST)
	}
}

var rsFaceCulling bool

//	Disables face-culling only if it is currently enabled.
func (me RenderStates) DisableFaceCulling() {
	if rsFaceCulling {
		rsFaceCulling = false
		gl.Disable(gl.CULL_FACE)
	}
}

//	Enables face-culling only if it is currently disabled.
func (me RenderStates) EnableFaceCulling() {
	if !rsFaceCulling {
		rsFaceCulling = true
		gl.Enable(gl.CULL_FACE)
	}
}

//	Deactivates face-culling.
func (_ RenderStates) ForceDisableFaceCulling() {
	rsFaceCulling = false
	gl.Disable(gl.CULL_FACE)
}

//	Activates face-culling.
func (_ RenderStates) ForceEnableFaceCulling() {
	rsFaceCulling = true
	gl.Enable(gl.CULL_FACE)
}

//	Activates or deactivates face-culling.
func (me RenderStates) SetFaceCulling(newFaceCulling bool) {
	if rsFaceCulling != newFaceCulling {
		if rsFaceCulling = newFaceCulling; rsFaceCulling {
			gl.Enable(gl.CULL_FACE)
		} else {
			gl.Disable(gl.CULL_FACE)
		}
	}
}

//	Toggles face-culling.
func (me RenderStates) ToggleFaceCulling() {
	if rsFaceCulling = !rsFaceCulling; rsFaceCulling {
		gl.Enable(gl.CULL_FACE)
	} else {
		gl.Disable(gl.CULL_FACE)
	}
}

var rsFramebufferSrgb bool

//	Disables SRGB-framebuffer only if it is currently enabled.
func (me RenderStates) DisableFramebufferSrgb() {
	if rsFramebufferSrgb {
		rsFramebufferSrgb = false
		gl.Disable(gl.FRAMEBUFFER_SRGB)
	}
}

//	Enables SRGB-framebuffer only if it is currently disabled.
func (me RenderStates) EnableFramebufferSrgb() {
	if !rsFramebufferSrgb {
		rsFramebufferSrgb = true
		gl.Enable(gl.FRAMEBUFFER_SRGB)
	}
}

//	Deactivates SRGB-framebuffer.
func (_ RenderStates) ForceDisableFramebufferSrgb() {
	rsFramebufferSrgb = false
	gl.Disable(gl.FRAMEBUFFER_SRGB)
}

//	Activates SRGB-framebuffer.
func (_ RenderStates) ForceEnableFramebufferSrgb() {
	rsFramebufferSrgb = true
	gl.Enable(gl.FRAMEBUFFER_SRGB)
}

//	Activates or deactivates SRGB-framebuffer.
func (me RenderStates) SetFramebufferSrgb(newFramebufferSrgb bool) {
	if rsFramebufferSrgb != newFramebufferSrgb {
		if rsFramebufferSrgb = newFramebufferSrgb; rsFramebufferSrgb {
			gl.Enable(gl.FRAMEBUFFER_SRGB)
		} else {
			gl.Disable(gl.FRAMEBUFFER_SRGB)
		}
	}
}

//	Toggles SRGB-framebuffer.
func (me RenderStates) ToggleFramebufferSrgb() {
	if rsFramebufferSrgb = !rsFramebufferSrgb; rsFramebufferSrgb {
		gl.Enable(gl.FRAMEBUFFER_SRGB)
	} else {
		gl.Disable(gl.FRAMEBUFFER_SRGB)
	}
}

var rsScissorTest bool

//	Disables scissor-testing only if it is currently enabled.
func (me RenderStates) DisableScissorTest() {
	if rsScissorTest {
		rsScissorTest = false
		gl.Disable(gl.SCISSOR_TEST)
	}
}

//	Enables scissor-testing only if it is currently disabled.
func (me RenderStates) EnableScissorTest() {
	if !rsScissorTest {
		rsScissorTest = true
		gl.Enable(gl.SCISSOR_TEST)
	}
}

//	Deactivates scissor-testing.
func (_ RenderStates) ForceDisableScissorTest() {
	rsScissorTest = false
	gl.Disable(gl.SCISSOR_TEST)
}

//	Activates scissor-testing.
func (_ RenderStates) ForceEnableScissorTest() {
	rsScissorTest = true
	gl.Enable(gl.SCISSOR_TEST)
}

//	Activates or deactivates scissor-testing.
func (me RenderStates) SetScissorTest(newScissorTest bool) {
	if rsScissorTest != newScissorTest {
		if rsScissorTest = newScissorTest; rsScissorTest {
			gl.Enable(gl.SCISSOR_TEST)
		} else {
			gl.Disable(gl.SCISSOR_TEST)
		}
	}
}

//	Toggles scissor-testing.
func (me RenderStates) ToggleScissorTest() {
	if rsScissorTest = !rsScissorTest; rsScissorTest {
		gl.Enable(gl.SCISSOR_TEST)
	} else {
		gl.Disable(gl.SCISSOR_TEST)
	}
}

var rsStencilTest bool

//	Disables stencil-testing only if it is currently enabled.
func (me RenderStates) DisableStencilTest() {
	if rsStencilTest {
		rsStencilTest = false
		gl.Disable(gl.STENCIL_TEST)
	}
}

//	Enables stencil-testing only if it is currently disabled.
func (me RenderStates) EnableStencilTest() {
	if !rsStencilTest {
		rsStencilTest = true
		gl.Enable(gl.STENCIL_TEST)
	}
}

//	Deactivates stencil-testing.
func (_ RenderStates) ForceDisableStencilTest() {
	rsStencilTest = false
	gl.Disable(gl.STENCIL_TEST)
}

//	Activates stencil-testing.
func (_ RenderStates) ForceEnableStencilTest() {
	rsStencilTest = true
	gl.Enable(gl.STENCIL_TEST)
}

//	Activates or deactivates stencil-testing.
func (me RenderStates) SetStencilTest(newStencilTest bool) {
	if rsStencilTest != newStencilTest {
		if rsStencilTest = newStencilTest; rsStencilTest {
			gl.Enable(gl.STENCIL_TEST)
		} else {
			gl.Disable(gl.STENCIL_TEST)
		}
	}
}

//	Toggles stencil-testing.
func (me RenderStates) ToggleStencilTest() {
	if rsStencilTest = !rsStencilTest; rsStencilTest {
		gl.Enable(gl.STENCIL_TEST)
	} else {
		gl.Disable(gl.STENCIL_TEST)
	}
}

//#end-gt
