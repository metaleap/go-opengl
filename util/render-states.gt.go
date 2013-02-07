package ugl

import (
	gl "github.com/go3d/go-opengl/core"
)

var (
	rsClearColor GlVec4
)

//	Encapsulates a particular combination of render-states that RenderStates can Apply().
type RenderStatesBag struct {
	Blending        bool
	DepthTest       bool
	FaceCulling     bool
	FramebufferSrgb bool
	StencilTest     bool
	ClearColor      GlVec4
	Other           struct {
		ClearBits gl.Bitfield
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

//	Calls the applicable me.SetFoo() methods for all render states in bag (except those in Other).
func (me RenderStates) Apply(bag *RenderStatesBag) {
	me.SetBlending(bag.Blending)
	me.SetDepthTest(bag.DepthTest)
	me.SetFaceCulling(bag.FaceCulling)
	me.SetFramebufferSrgb(bag.FramebufferSrgb)
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

//#begin-gt -gen-renderstates.gt GT_MULT_SEP:, N:Blending,DepthTest,FaceCulling,FramebufferSrgb,ScissorTest,StencilTest D:blending,depth-testing,face-culling,framebuffer-SRGB,scissor-testing,stencil-testing E:BLEND,DEPTH_TEST,CULL_FACE,FRAMEBUFFER_SRGB,SCISSOR_TEST,STENCIL_TEST

var rsBlending bool

//	Disables blending only if it is currently enabled.
func (me RenderStates) DisableBlending() {
	if rsBlending {
		me.ForceDisableBlending()
	}
}

//	Enables blending only if it is currently disabled.
func (me RenderStates) EnableBlending() {
	if !rsBlending {
		me.ForceEnableBlending()
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
func (me RenderStates) SetBlending(Blending bool) {
	if Blending {
		me.EnableBlending()
	} else {
		me.DisableBlending()
	}
}

//	Toggles blending.
func (me RenderStates) ToggleBlending() {
	if rsBlending {
		me.ForceDisableBlending()
	} else {
		me.ForceEnableBlending()
	}
}



var rsDepthTest bool

//	Disables depth-testing only if it is currently enabled.
func (me RenderStates) DisableDepthTest() {
	if rsDepthTest {
		me.ForceDisableDepthTest()
	}
}

//	Enables depth-testing only if it is currently disabled.
func (me RenderStates) EnableDepthTest() {
	if !rsDepthTest {
		me.ForceEnableDepthTest()
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
func (me RenderStates) SetDepthTest(DepthTest bool) {
	if DepthTest {
		me.EnableDepthTest()
	} else {
		me.DisableDepthTest()
	}
}

//	Toggles depth-testing.
func (me RenderStates) ToggleDepthTest() {
	if rsDepthTest {
		me.ForceDisableDepthTest()
	} else {
		me.ForceEnableDepthTest()
	}
}



var rsFaceCulling bool

//	Disables face-culling only if it is currently enabled.
func (me RenderStates) DisableFaceCulling() {
	if rsFaceCulling {
		me.ForceDisableFaceCulling()
	}
}

//	Enables face-culling only if it is currently disabled.
func (me RenderStates) EnableFaceCulling() {
	if !rsFaceCulling {
		me.ForceEnableFaceCulling()
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
func (me RenderStates) SetFaceCulling(FaceCulling bool) {
	if FaceCulling {
		me.EnableFaceCulling()
	} else {
		me.DisableFaceCulling()
	}
}

//	Toggles face-culling.
func (me RenderStates) ToggleFaceCulling() {
	if rsFaceCulling {
		me.ForceDisableFaceCulling()
	} else {
		me.ForceEnableFaceCulling()
	}
}



var rsFramebufferSrgb bool

//	Disables framebuffer-SRGB only if it is currently enabled.
func (me RenderStates) DisableFramebufferSrgb() {
	if rsFramebufferSrgb {
		me.ForceDisableFramebufferSrgb()
	}
}

//	Enables framebuffer-SRGB only if it is currently disabled.
func (me RenderStates) EnableFramebufferSrgb() {
	if !rsFramebufferSrgb {
		me.ForceEnableFramebufferSrgb()
	}
}

//	Deactivates framebuffer-SRGB.
func (_ RenderStates) ForceDisableFramebufferSrgb() {
	rsFramebufferSrgb = false
	gl.Disable(gl.FRAMEBUFFER_SRGB)
}

//	Activates framebuffer-SRGB.
func (_ RenderStates) ForceEnableFramebufferSrgb() {
	rsFramebufferSrgb = true
	gl.Enable(gl.FRAMEBUFFER_SRGB)
}

//	Activates or deactivates framebuffer-SRGB.
func (me RenderStates) SetFramebufferSrgb(FramebufferSrgb bool) {
	if FramebufferSrgb {
		me.EnableFramebufferSrgb()
	} else {
		me.DisableFramebufferSrgb()
	}
}

//	Toggles framebuffer-SRGB.
func (me RenderStates) ToggleFramebufferSrgb() {
	if rsFramebufferSrgb {
		me.ForceDisableFramebufferSrgb()
	} else {
		me.ForceEnableFramebufferSrgb()
	}
}



var rsScissorTest bool

//	Disables scissor-testing only if it is currently enabled.
func (me RenderStates) DisableScissorTest() {
	if rsScissorTest {
		me.ForceDisableScissorTest()
	}
}

//	Enables scissor-testing only if it is currently disabled.
func (me RenderStates) EnableScissorTest() {
	if !rsScissorTest {
		me.ForceEnableScissorTest()
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
func (me RenderStates) SetScissorTest(ScissorTest bool) {
	if ScissorTest {
		me.EnableScissorTest()
	} else {
		me.DisableScissorTest()
	}
}

//	Toggles scissor-testing.
func (me RenderStates) ToggleScissorTest() {
	if rsScissorTest {
		me.ForceDisableScissorTest()
	} else {
		me.ForceEnableScissorTest()
	}
}



var rsStencilTest bool

//	Disables stencil-testing only if it is currently enabled.
func (me RenderStates) DisableStencilTest() {
	if rsStencilTest {
		me.ForceDisableStencilTest()
	}
}

//	Enables stencil-testing only if it is currently disabled.
func (me RenderStates) EnableStencilTest() {
	if !rsStencilTest {
		me.ForceEnableStencilTest()
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
func (me RenderStates) SetStencilTest(StencilTest bool) {
	if StencilTest {
		me.EnableStencilTest()
	} else {
		me.DisableStencilTest()
	}
}

//	Toggles stencil-testing.
func (me RenderStates) ToggleStencilTest() {
	if rsStencilTest {
		me.ForceDisableStencilTest()
	} else {
		me.ForceEnableStencilTest()
	}
}

//#end-gt
