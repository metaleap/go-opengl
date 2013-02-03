package glutil

import (
	gl "github.com/go3d/go-opengl/core"
)

//	Represents an OpenGL buffer object.
type Buffer struct {
	//	The OpenGL handle to this buffer object.
	//	This is 0 before calling Recreate() and after calling Dispose().
	GlHandle gl.Uint

	//	The current binding target for this buffer object
	//	(such as for example gl.ARRAY_BUFFER, gl.ELEMENT_ARRAY_BUFFER etc.).
	GlTarget gl.Enum
}

//	Initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.ARRAY_BUFFER.
func NewArrayBuffer() *Buffer {
	return NewBuffer(gl.ARRAY_BUFFER)
}

//	If Support.Buffers.AtomicCounter is true, initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.ATOMIC_COUNTER_BUFFER.
func NewAtomicCounterBuffer() (me *Buffer) {
	if Support.Buffers.AtomicCounter {
		me = NewBuffer(gl.ATOMIC_COUNTER_BUFFER)
	}
	return
}

//	Initializes and returns --but does not Recreate()-- a new Buffer not yet initialized with any valid GlTarget.
func NewBuffer(glTarget gl.Enum) (me *Buffer) {
	me = &Buffer{GlTarget: glTarget}
	return
}

//	Initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.COPY_READ_BUFFER.
func NewCopyReadBuffer() *Buffer {
	return NewBuffer(gl.COPY_READ_BUFFER)
}

//	Initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.COPY_WRITE_BUFFER.
func NewCopyWriteBuffer() *Buffer {
	return NewBuffer(gl.COPY_WRITE_BUFFER)
}

//	If Support.Buffers.DispatchIndirect is true, initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.DISPATCH_INDIRECT_BUFFER.
func NewDispatchIndirectBuffer() (me *Buffer) {
	if Support.Buffers.DispatchIndirect {
		me = NewBuffer(gl.DISPATCH_INDIRECT_BUFFER)
	}
	return
}

//	Initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.DRAW_INDIRECT_BUFFER.
func NewDrawIndirectBuffer() *Buffer {
	return NewBuffer(gl.DRAW_INDIRECT_BUFFER)
}

//	Initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.ELEMENT_ARRAY_BUFFER.
func NewElementArrayBuffer() *Buffer {
	return NewBuffer(gl.ELEMENT_ARRAY_BUFFER)
}

//	Initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.PIXEL_PACK_BUFFER.
func NewPixelPackBuffer() *Buffer {
	return NewBuffer(gl.PIXEL_PACK_BUFFER)
}

//	Initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.PIXEL_UNPACK_BUFFER.
func NewPixelUnpackBuffer() *Buffer {
	return NewBuffer(gl.PIXEL_UNPACK_BUFFER)
}

//	If Support.Buffers.ShaderStorage is true, initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.SHADER_STORAGE_BUFFER.
func NewShaderStorageBuffer() (me *Buffer) {
	if Support.Buffers.ShaderStorage {
		me = NewBuffer(gl.SHADER_STORAGE_BUFFER)
	}
	return
}

//	Initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.TEXTURE_BUFFER.
func NewTextureBuffer() *Buffer {
	return NewBuffer(gl.TEXTURE_BUFFER)
}

//	Initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.TRANSFORM_FEEDBACK_BUFFER.
func NewTransformFeedbackBuffer() *Buffer {
	return NewBuffer(gl.TRANSFORM_FEEDBACK_BUFFER)
}

//	Initializes and returns --but does not Recreate()-- a new Buffer initialized with a GlTarget of gl.UNIFORM_BUFFER.
func NewUniformBuffer() *Buffer {
	return NewBuffer(gl.UNIFORM_BUFFER)
}

//	Binds this buffer object to its current GlTarget.
func (me *Buffer) Bind() {
	gl.BindBuffer(me.GlTarget, me.GlHandle)
}

//	Convenience short-hand for gl.BindBufferBase().
//	Only permissible if me.GlTarget is gl.TRANSFORM_FEEDBACK_BUFFER or gl.UNIFORM_BUFFER.
func (me *Buffer) BindBase(index gl.Uint) {
	gl.BindBufferBase(me.GlTarget, index, me.GlHandle)
}

//	Convenience short-hand for gl.BindBufferRange().
//	Only permissible if me.GlTarget is gl.TRANSFORM_FEEDBACK_BUFFER or gl.UNIFORM_BUFFER.
func (me *Buffer) BindRange(index gl.Uint, offset gl.Intptr, size gl.Sizeiptr) {
	gl.BindBufferRange(me.GlTarget, index, me.GlHandle, offset, size)
}

//	Deletes this buffer object.
func (me *Buffer) Dispose() {
	if me.GlHandle != 0 {
		gl.DeleteBuffers(1, &me.GlHandle)
		me.GlHandle = 0
	}
}

//	Maps to the client's address space the entire data store of whatever buffer object is currently bound to me.GlTarget.
func (me *Buffer) Map(read, write bool) {
	gl.MapBuffer(me.GlTarget, Typed.Ifb(read, write, gl.READ_ONLY, gl.WRITE_ONLY, gl.READ_WRITE))
}

//	Deletes this buffer object if applicable, generates it anew, and allocates its data store as specified,
//	populating it with the specified initialData if any. GlTarget is changed to newGlTarget unless 0 is specified.
func (me *Buffer) Recreate(newGlTarget gl.Enum, sizeInBytes gl.Sizeiptr, initialData gl.Ptr, usageHint gl.Enum) {
	if me.GlTarget != 0 {
		me.Unbind()
	}
	if newGlTarget != 0 {
		me.GlTarget = newGlTarget
	}
	if me.GlTarget != 0 {
		me.Unbind()
		me.Dispose()
		gl.GenBuffers(1, &me.GlHandle)
		me.Bind()
		gl.BufferData(me.GlTarget, sizeInBytes, initialData, usageHint)
		me.Unbind()
	}
}

//	Unbinds whatever buffer object is currently bound to me.GlTarget.
func (me *Buffer) Unbind() {
	gl.BindBuffer(me.GlTarget, 0)
}

//	Convenience short-hand for gl.UnmapBuffer(me.GlTarget)
func (me *Buffer) Unmap() {
	gl.UnmapBuffer(me.GlTarget)
}

//	Buffers the specified data into this buffer object's data store at the specified offset.
func (me *Buffer) SubData(offset gl.Intptr, size gl.Sizeiptr, data gl.Ptr) {
	gl.BufferSubData(me.GlTarget, offset, size, data)
}
