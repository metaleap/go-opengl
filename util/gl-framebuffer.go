package ugl

import (
	gl "github.com/go3d/go-opengl/core"
)

//	Encapsulates an OpenGL framebuffer object.
type Framebuffer struct {
	//	Set in Create(), this is either gl.READ_FRAMEBUFFER or gl.FRAMEBUFFER
	GlTarget gl.Enum

	//	The OpenGL handle to this framebuffer object.
	//	This is 0 before calling Create() and after calling Dispose().
	GlHandle gl.Uint

	width, height gl.Sizei

	renderbuffers  []*FramebufferRenderbuffer
	rendertextures []*FramebufferRendertexture
}

//	Attaches the specified renderbuffer to this Framebuffer.
func (me *Framebuffer) AttachRenderbuffer(renderbuffer *FramebufferRenderbuffer) {
	me.renderbuffers = append(me.renderbuffers, renderbuffer)
	me.reinitRenderbuffer(renderbuffer)
}

//	Attaches the specified render texture to this Framebuffer.
func (me *Framebuffer) AttachRendertexture(texture *FramebufferRendertexture) {
	me.rendertextures = append(me.rendertextures, texture)
	me.reinitTexture(texture)
}

//	Binds this framebuffer object.
func (me *Framebuffer) Bind() {
	gl.BindFramebuffer(me.GlTarget, me.GlHandle)
}

//	Creates this Framebuffer in OpenGL sized as specified.
func (me *Framebuffer) Create(width, height gl.Sizei, read bool) {
	me.width, me.height = width, height
	me.GlTarget = Typed.Ife(read, gl.READ_FRAMEBUFFER, gl.DRAW_FRAMEBUFFER)
	gl.GenFramebuffers(1, &me.GlHandle)
	for _, tex := range me.rendertextures {
		me.reinitTexture(tex)
	}
	for _, rb := range me.renderbuffers {
		me.reinitRenderbuffer(rb)
	}
}

//	Deletes this Framebuffer and all attached render textures and render buffers from OpenGL.
func (me *Framebuffer) Dispose() {
	for _, tex := range me.rendertextures {
		tex.Dispose()
	}
	me.rendertextures = nil
	for _, rb := range me.renderbuffers {
		rb.dispose()
	}
	me.renderbuffers = nil
	if me.GlHandle != 0 {
		gl.DeleteFramebuffers(1, &me.GlHandle)
		me.GlHandle = 0
	}
}

func (me *Framebuffer) reinitRenderbuffer(rb *FramebufferRenderbuffer) {
	if (me.width > 0) && (me.height > 0) && (me.GlHandle != 0) {
		firstInit := (rb.GlHandle == 0)
		if firstInit {
			gl.GenRenderbuffers(1, &rb.GlHandle)
		}
		rb.Bind()
		gl.RenderbufferStorage(rb.GlTarget, rb.InternalFormat, me.width, me.height)
		if firstInit {
			me.Bind()
			gl.FramebufferRenderbuffer(me.GlTarget, rb.Attachment, rb.GlTarget, rb.GlHandle)
			me.Unbind()
		}
		rb.Unbind()
	}
}

func (me *Framebuffer) reinitTexture(tex *FramebufferRendertexture) {
	if (me.width > 0) && (me.height > 0) && (me.GlHandle != 0) {
		tex.Width, tex.Height = me.width, me.height
		tex.Recreate()
		me.Bind()
		gl.FramebufferTexture2D(me.GlTarget, tex.Attachment, tex.GlTarget, tex.GlHandle, 0)
		me.Unbind()
	}
}

func (me *Framebuffer) RenderTexture(index int) *FramebufferRendertexture {
	return me.rendertextures[index]
}

//	Resizes this Framebuffer and all attached render textures and render buffers.
func (me *Framebuffer) Resize(width, height gl.Sizei) (resized bool) {
	if resized = (me.width != width) || (me.height != height); resized {
		me.width, me.height = width, height
		for _, tex := range me.rendertextures {
			me.reinitTexture(tex)
		}
		for _, rb := range me.renderbuffers {
			me.reinitRenderbuffer(rb)
		}
	}
	return
}

//	The current status of this Framebuffer as returned by gl.CheckFramebufferStatus().
func (me *Framebuffer) Status() gl.Enum {
	return gl.CheckFramebufferStatus(me.GlTarget)
}

//	Unbinds whatever Framebuffer is currently bound, and binds the
//	"default" framebuffer (typically, the window's drawing surface).
func (me *Framebuffer) Unbind() {
	gl.BindFramebuffer(me.GlTarget, 0)
}

//	Represents a Renderbuffer that can be attached to a Framebuffer.
type FramebufferRenderbuffer struct {
	//	The attachment point of this render buffer.
	//	Defaults to gl.DEPTH_STENCIL_ATTACHMENT.
	Attachment gl.Enum

	//	Defaults to gl.RENDERBUFFER.
	GlTarget gl.Enum

	//	The OpenGL handle to this render buffer object.
	GlHandle gl.Uint

	//	Defaults to gl.DEPTH24_STENCIL8.
	InternalFormat gl.Enum
}

//	Initializes a new FramebufferRenderbuffer with default values and returns it.
func NewFramebufferRenderbuffer() (me *FramebufferRenderbuffer) {
	me = &FramebufferRenderbuffer{Attachment: gl.DEPTH_STENCIL_ATTACHMENT, GlTarget: gl.RENDERBUFFER, InternalFormat: gl.DEPTH24_STENCIL8}
	return
}

//	Binds this render buffer object.
func (me *FramebufferRenderbuffer) Bind() {
	gl.BindRenderbuffer(me.GlTarget, me.GlHandle)
}

func (me *FramebufferRenderbuffer) dispose() {
	if me.GlHandle != 0 {
		gl.DeleteRenderbuffers(1, &me.GlHandle)
		me.GlHandle = 0
	}
}

//	Unbinds whatever render buffer object is currently bound, if any.
func (me *FramebufferRenderbuffer) Unbind() {
	gl.BindRenderbuffer(me.GlTarget, 0)
}

//	Represents a render-to-texture object that can be attached to a Framebuffer.
type FramebufferRendertexture struct {
	//	The texture object being rendered to.
	Texture2D

	//	The attachment point for this render texture object.
	//	Defaults to gl.COLOR_ATTACHMENT0.
	Attachment gl.Enum
}

//	Initializes --but does not Recreate()-- a new FramebufferRendertexture with default values and returns it.
func NewFramebufferRendertexture() (me *FramebufferRendertexture) {
	me = &FramebufferRendertexture{Attachment: gl.COLOR_ATTACHMENT0}
	me.Texture2D.Init()
	me.Texture2D.SizedInternalFormat = gl.RGB8
	me.Texture2D.PixelData.Format = gl.RGB
	me.Texture2D.MipMap.AutoGen, me.Texture2D.MipMap.NumLevels = false, 1
	return
}
