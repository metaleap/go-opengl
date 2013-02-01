package glutil

import (
	gl "github.com/go3d/go-opengl/gogl"
)

var (
	rsBlending, rsFaceCulling, rsDepthTest, rsScissorTest, rsStencilTest bool
	rsClearColor                                                         GlVec4
)

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
	me.SetStencilTest(bag.StencilTest)
	me.SetClearColor(bag.ClearColor)
}

//	Disables blending only if it is currently enabled.
func (me RenderStates) DisableBlending() {
	if rsBlending {
		me.ForceDisableBlending()
	}
}

//	Disables depth-testing only if it is currently enabled.
func (me RenderStates) DisableDepthTest() {
	if rsDepthTest {
		me.ForceDisableDepthTest()
	}
}

//	Disables face-culling only if it is currently enabled.
func (me RenderStates) DisableFaceCulling() {
	if rsFaceCulling {
		me.ForceDisableFaceCulling()
	}
}

//	Disables scissor-testing only if it is currently enabled.
func (me RenderStates) DisableScissorTest() {
	if rsScissorTest {
		me.ForceDisableScissorTest()
	}
}

//	Disables stencil-testing only if it is currently enabled.
func (me RenderStates) DisableStencilTest() {
	if rsStencilTest {
		me.ForceDisableStencilTest()
	}
}

//	Enables blending only if it is currently disabled.
func (me RenderStates) EnableBlending() {
	if !rsBlending {
		me.ForceEnableBlending()
	}
}

//	Enables depth-testing only if it is currently disabled.
func (me RenderStates) EnableDepthTest() {
	if !rsDepthTest {
		me.ForceEnableDepthTest()
	}
}

//	Enables face-culling only if it is currently disabled.
func (me RenderStates) EnableFaceCulling() {
	if !rsFaceCulling {
		me.ForceEnableFaceCulling()
	}
}

//	Enables scissor-testing only if it is currently disabled.
func (me RenderStates) EnableScissorTest() {
	if !rsScissorTest {
		me.ForceEnableScissorTest()
	}
}

//	Enables stencil-testing only if it is currently disabled.
func (me RenderStates) EnableStencilTest() {
	if !rsStencilTest {
		me.ForceEnableStencilTest()
	}
}

//	Sets the OpenGL clear-color to the specified rgba.
func (_ RenderStates) ForceClearColor(rgba GlVec4) {
	rsClearColor = rgba
	gl.ClearColor(rsClearColor[0], rsClearColor[1], rsClearColor[2], rsClearColor[3])
}

//	Deactivates blending.
func (_ RenderStates) ForceDisableBlending() {
	rsBlending = false
	gl.Disable(gl.BLEND)
}

//	Deactivates depth-testing.
func (_ RenderStates) ForceDisableDepthTest() {
	rsDepthTest = false
	gl.Disable(gl.DEPTH_TEST)
}

//	Deactivates face-culling.
func (_ RenderStates) ForceDisableFaceCulling() {
	rsFaceCulling = false
	gl.Disable(gl.CULL_FACE)
}

//	Deactivates scissor-testing.
func (_ RenderStates) ForceDisableScissorTest() {
	rsScissorTest = false
	gl.Disable(gl.SCISSOR_TEST)
}

//	Deactivates stencil-testing.
func (_ RenderStates) ForceDisableStencilTest() {
	rsStencilTest = false
	gl.Disable(gl.STENCIL_TEST)
}

//	Activates blending.
func (_ RenderStates) ForceEnableBlending() {
	rsBlending = true
	gl.Enable(gl.BLEND)
}

//	Activates depth-testing.
func (_ RenderStates) ForceEnableDepthTest() {
	rsDepthTest = true
	gl.Enable(gl.DEPTH_TEST)
}

//	Activates face-culling.
func (_ RenderStates) ForceEnableFaceCulling() {
	rsFaceCulling = true
	gl.Enable(gl.CULL_FACE)
}

//	Activates scissor-testing.
func (_ RenderStates) ForceEnableScissorTest() {
	rsScissorTest = true
	gl.Enable(gl.SCISSOR_TEST)
}

//	Activates stencil-testing.
func (_ RenderStates) ForceEnableStencilTest() {
	rsStencilTest = true
	gl.Enable(gl.STENCIL_TEST)
}

//	Sets the OpenGL clear-color if rgba differs from the currently cached value.
func (me RenderStates) SetClearColor(rgba GlVec4) {
	if rsClearColor != rgba {
		me.ForceClearColor(rgba)
	}
}

//	Activates or deactivates blending.
func (me RenderStates) SetBlending(blending bool) {
	if blending {
		me.EnableBlending()
	} else {
		me.DisableBlending()
	}
}

//	Activates or deactivates depth-testing.
func (me RenderStates) SetDepthTest(depthTest bool) {
	if depthTest {
		me.EnableDepthTest()
	} else {
		me.DisableDepthTest()
	}
}

//	Activates or deactivates face-culling.
func (me RenderStates) SetFaceCulling(faceCulling bool) {
	if faceCulling {
		me.EnableFaceCulling()
	} else {
		me.DisableFaceCulling()
	}
}

//	Activates or deactivates scissor-testing.
func (me RenderStates) SetScissorTest(scissorTest bool) {
	if scissorTest {
		me.EnableScissorTest()
	} else {
		me.DisableScissorTest()
	}
}

//	Activates or deactivates stencil-testing.
func (me RenderStates) SetStencilTest(StencilTest bool) {
	if StencilTest {
		me.EnableStencilTest()
	} else {
		me.DisableStencilTest()
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

//	Toggles depth-testing.
func (me RenderStates) ToggleDepthTest() {
	if rsDepthTest {
		me.ForceDisableDepthTest()
	} else {
		me.ForceEnableDepthTest()
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

//	Toggles scissor-testing.
func (me RenderStates) ToggleScissorTest() {
	if rsScissorTest {
		me.ForceDisableScissorTest()
	} else {
		me.ForceEnableScissorTest()
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

//	Encapsulates a particular combination of render-states that RenderStates can Apply().
type RenderStatesBag struct {
	Blending    bool
	DepthTest   bool
	FaceCulling bool
	StencilTest bool
	ClearColor  GlVec4
	Other       struct {
		ClearBits gl.Bitfield
	}
}
