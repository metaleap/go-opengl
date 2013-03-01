package ugl

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
	Cache.BindVertexArray(me.GlHandle)
}

//	Creates this vertex-array object.
func (me *VertexArray) Create() (err error) {
	if me.GlHandle == 0 {
		err = Try.GenVertexArrays(1, &me.GlHandle)
	}
	return
}

//	Deletes this vertex array object.
func (me *VertexArray) Dispose() {
	if me.GlHandle != 0 {
		gl.DeleteVertexArrays(1, &me.GlHandle)
		me.GlHandle = 0
	}
}

//	Sets up this vertex array object. Unless atts are specified, prog can be be nil.
//	To specify element/index buffer and vertex/attribute buffers (if applicable), pass them via bufs.
func (me *VertexArray) Setup(prog *Program, atts []VertexAttribPointer, bufs ...*Buffer) (err error) {
	var (
		buf  *Buffer
		aloc gl.Uint
	)
	me.Bind()
	for _, buf = range bufs {
		buf.Bind()
	}
	for i := 0; i < len(atts); i++ {
		if aloc = prog.AttrLocs[atts[i].Name]; prog.HasAttr(atts[i].Name) {
			gl.EnableVertexAttribArray(aloc)
			gl.VertexAttribPointer(aloc, atts[i].Size, atts[i].Type, atts[i].Normalized, atts[i].Stride, atts[i].Offset)
		}
	}
	me.Unbind()
	for _, buf = range bufs {
		buf.Unbind()
	}
	err = Util.LastError("VertexArray.Setup()")
	return
}

//	Unbinds whatever vertex array object is currently bound.
func (_ VertexArray) Unbind() {
	Cache.BindVertexArray(0)
}

//	Encapsulates vertex attribute information used by VertexArray.Setup() to enable that vertex attribute.
type VertexAttribPointer struct {
	//	The name of the vertex attribute in a Program.AttrLocs hash-table.
	Name string

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

	//	Specifies an offset of the first component of the first generic vertex attribute in the array
	//	in the data store of the Buffer currently bound to the gl.ARRAY_BUFFER target.
	Offset gl.Ptr
}

//	Initializes and returns a new VertexAttribPointer with the specified values.
func NewVertexAttribPointer(name string, size gl.Int, stride gl.Sizei, offset gl.Ptr) (me *VertexAttribPointer) {
	me = &VertexAttribPointer{}
	me.Init(name, size, stride, offset)
	return
}

func (me *VertexAttribPointer) Init(name string, size gl.Int, stride gl.Sizei, offset gl.Ptr) {
	me.Name, me.Type, me.Normalized = name, gl.FLOAT, gl.FALSE
	me.Size, me.Stride, me.Offset = size, stride, offset
}
