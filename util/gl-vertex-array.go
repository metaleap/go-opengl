package glutil

import (
	gl "github.com/go3d/go-opengl/core"
)

//	Represents an OpenGL vertex array object (NOT the legacy OpenGL "vertex arrays", but what is commonly abbreviated as VAO).
type VertexArray struct {
	//	The OpenGL handle to this vertex array object.
	//	This is 0 before calling Create() and after calling Dispose().
	GlHandle gl.Uint
}

//	Binds this vertex array object.
func (me *VertexArray) Bind() {
	gl.BindVertexArray(me.GlHandle)
}

//	(Re-)Creates this vertex-array object.
func (me *VertexArray) Create() (err error) {
	me.Dispose()
	err = gl.Try.GenVertexArrays(1, &me.GlHandle)
	return
}

//	Deletes this vertex array object.
func (me *VertexArray) Dispose() {
	if me.GlHandle != 0 {
		gl.DeleteVertexArrays(1, &me.GlHandle)
		me.GlHandle = 0
	}
}

//	Sets up this vertex array object, associating it with the specified buffer objects and enabling the specified vertex attributes for it.
func (me *VertexArray) Setup(atts []*VertexAttribPointer, bufs ...*Buffer) (err error) {
	me.Bind()
	for _, buf := range bufs {
		buf.Bind()
	}
	for _, attr := range atts {
		gl.EnableVertexAttribArray(attr.Loc)
		gl.VertexAttribPointer(attr.Loc, attr.Size, attr.Type, attr.Normalized, attr.Stride, attr.Offset)
	}
	me.Unbind()
	for _, buf := range bufs {
		buf.Unbind()
	}
	err = gl.Util.Error("VertexArray.Setup()")
	return
}

//	Unbinds whatever vertex array object is currently bound.
func (_ VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

//	Encapsulates vertex attribute information used by VertexArray.Setup() to enable that vertex attribute.
type VertexAttribPointer struct {
	//	The name of the vertex attribute in a Program.AttrLocs hash-table.
	Name string

	//	The location of the vertex attribute in a Program.AttrLocs hash-table.
	Loc gl.Uint

	//	Specifies the number of components per generic vertex attribute. Must be 1, 2, 3, 4.
	//	Additionally, the symbolic constant GL_BGRA is accepted.
	Size gl.Int

	//	Defaults to gl.FLOAT
	Type gl.Enum

	//	Defaults to gl.FALSE
	Normalized gl.Boolean

	//	Specifies the byte offset between consecutive generic vertex attributes.
	//	If Stride is 0, the generic vertex attributes are understood to be tightly packed in the array.
	Stride gl.Sizei

	//	Specifies a offset of the first component of the first generic vertex attribute in the array
	//	in the data store of the Buffer currently bound to the gl.ARRAY_BUFFER target.
	Offset gl.Ptr
}

//	Initializes and returns a new VertexAttribPointer with the specified values.
func NewVertexAttribPointer(name string, loc gl.Uint, size gl.Int, stride gl.Sizei, offset gl.Ptr) (me *VertexAttribPointer) {
	me = &VertexAttribPointer{Type: gl.FLOAT, Normalized: gl.FALSE, Loc: loc, Name: name, Size: size, Stride: stride, Offset: offset}
	return
}
