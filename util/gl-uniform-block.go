package ugl

import (
	gl "github.com/metaleap/go-opengl/core"
)

type UniformBlock struct {
	Buffer Buffer
	Name   string
}

func NewUniformBlock(name string) (me *UniformBlock) {
	me = &UniformBlock{Name: name}
	me.Buffer.GlTarget = gl.UNIFORM_BUFFER
	return
}

func (me *UniformBlock) Create() {
	const sizeInBytes = 0
	me.Buffer.Recreate(0, sizeInBytes, PtrNil, gl.DYNAMIC_DRAW)
}

func (me *UniformBlock) Dispose() {
	me.Buffer.Dispose()
}
