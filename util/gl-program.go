package ugl

import (
	"fmt"

	gl "github.com/go3d/go-opengl/core"
)

const (
	progAttrInvalidLoc gl.Uint = 4294967295
)

//	Represents an OpenGL program object.
type Program struct {
	//	The OpenGL handle to the underlying program object.
	//	This is 0 before calling Create() and after calling Dispose().
	GlHandle gl.Uint

	//	An arbitrary application-defined name that is of no use to OpenGL, but may aid client-side diagnostics.
	Name string

	//	All vertex-attributes and their locations mapped for this program object (after the SetAttrLocations() method has been called).
	AttrLocs map[string]gl.Uint

	//	All uniforms and their locations mapped for this program object (after the SetUnifLocations() method has been called).
	UnifLocs map[string]gl.Int
}

//	Initializes and returns --but does not Create()-- a new Program with the specified name.
func NewProgram(name string) (me *Program) {
	me = &Program{Name: name, AttrLocs: map[string]gl.Uint{}, UnifLocs: map[string]gl.Int{}}
	return
}

//	Creates and compiles Shader stages for the specified sources (and defines) and links them together in this Program.
//	All sources are string pointers and WILL be modified by this method to contain the final real
//	GLSL source string sent to GL as returned by Shader.SetSource(). This may aid diagnostics or be discarded,
//	but a source string should not be re-sent to a subsequent compilation in that modified form.
func (me *Program) CompileAndLinkShaders(compute, fragment, geometry, tessCtl, tessEval, vertex *string, defines map[string]interface{}) (err error) {
	type shaderCtor func(string) *Shader
	var (
		ctor       shaderCtor
		source     *string
		shader     *Shader
		allShaders []*Shader
	)
	for source, ctor = range map[*string]shaderCtor{
		compute:  NewComputeShader,
		fragment: NewFragmentShader,
		geometry: NewGeometryShader,
		tessCtl:  NewTessCtlShader,
		tessEval: NewTessEvalShader,
		vertex:   NewVertexShader,
	} {
		if len(*source) > 0 {
			if shader = ctor(me.Name); shader != nil {
				if err = shader.Create(); err != nil {
					return
				}
				defer shader.Dispose()
				if *source, err = shader.SetSource(*source, defines); err != nil {
					return
				}
				if err = shader.Compile(); err != nil {
					return
				}
				allShaders = append(allShaders, shader)
			}
		}
	}
	if len(allShaders) == 0 {
		err = fmt.Errorf("Program'%s'.CompileAndLinkShaders() -- no shaders specified.", me.Name)
		return
	}
	for _, shader = range allShaders {
		if err = shader.AttachTo(me); err != nil {
			return
		}
		defer shader.DetachFrom(me)
	}
	if err = me.Link(); err != nil {
		return
	}
	return
}

//	Creates this program object in OpenGL.
func (me *Program) Create() (err error) {
	me.Dispose()
	err, me.GlHandle = Try.CreateProgram()
	return
}

//	Deletes this program object from OpenGL.
func (me *Program) Dispose() {
	if me.GlHandle != 0 {
		gl.DeleteProgram(me.GlHandle)
		me.GlHandle = 0
	}
}

//	Returns true if the specified attribute name is present as a key in me.AttrLocs.
func (me *Program) HasAttr(name string) (ok bool) {
	_, ok = me.AttrLocs[name]
	return
}

//	Returns true if the specified uniform name is present as a key in me.UnifLocs.
func (me *Program) HasUnif(name string) (ok bool) {
	_, ok = me.UnifLocs[name]
	return
}

//	Retrieves and returns the content of the current OpenGL info-log for this shader object.
func (me *Program) InfoLog() string {
	return shaderProgInfoLog(me.Name, me.GlHandle, false)
}

//	Links this program object. This is a convenience short-hand for calling gl.LinkProgram(),
//	checking gl.LINK_STATUS and obtaining gl.GetProgramInfoLog().
func (me *Program) Link() (err error) {
	if gl.LinkProgram(me.GlHandle); me.ParamInt(gl.LINK_STATUS) == gl.FALSE {
		err = fmt.Errorf("Program'%s'.Link() error: %s\n", me.Name, me.InfoLog())
	}
	return
}

func (me *Program) locationAttr(attrName string) gl.Uint {
	s := gl.Util.CString(attrName)
	defer gl.Util.CStringFree(s)
	l := gl.GetAttribLocation(me.GlHandle, s)
	if l < 0 {
		return progAttrInvalidLoc
	}
	return gl.Uint(l)
}

func (me *Program) locationUnif(unifName string) gl.Int {
	s := gl.Util.CString(unifName)
	defer gl.Util.CStringFree(s)
	return gl.GetUniformLocation(me.GlHandle, s)
}

//	Convenience short-hand for gl.GetProgramiv.
func (me *Program) ParamInt(pname gl.Enum) (iv gl.Int) {
	gl.GetProgramiv(me.GlHandle, pname, &iv)
	return
}

//	Stores the locations for the specified vertex-attributes in the AttrLocs field.
//	If an attribute in attrNames isn't found, it will not be added to AttrLocs but this may not result in an error.
//	err is the result of LastError() after each call to gl.GetAttribLocation(), the method immediately returns on the first error encountered.
func (me *Program) SetAttrLocations(attrNames ...string) (err error) {
	var loc gl.Uint
	for _, attrName := range attrNames {
		if len(attrName) > 0 {
			loc = me.locationAttr(attrName)
			if err = Util.LastError("Program'%s'.SetAttrLocations('%s')", me.Name, attrName); err != nil {
				return
			} else if progIsAttrLocation(loc) {
				me.AttrLocs[attrName] = loc
			}
		}
	}
	return
}

//	Stores the locations for the specified uniforms in the UnifLocs field.
//	If a uniform in unifNames isn't found, it will not be added to UnifLocs but this may not result in an error.
//	err is the result of LastError() after each call to gl.GetUniformLocation(), the method immediately returns on the first error encountered.
func (me *Program) SetUnifLocations(unifNames ...string) (err error) {
	var loc gl.Int
	for _, unifName := range unifNames {
		if len(unifName) > 0 {
			loc = me.locationUnif(unifName)
			if err = Util.LastError("Program'%s'.SetUnifLocations('%s')", me.Name, unifName); err != nil {
				return
			} else if progIsUnifLocation(loc) {
				me.UnifLocs[unifName] = loc
			}
		}
	}
	return
}

//	Installs this program object as part of current rendering state.
func (me *Program) Use() {
	gl.UseProgram(me.GlHandle)
}

//	Validates this program object. This is a convenience short-hand for calling gl.ValidateProgram(),
//	checking gl.VALIDATE_STATUS and obtaining gl.GetProgramInfoLog().
func (me *Program) Validate() (err error) {
	if gl.ValidateProgram(me.GlHandle); me.ParamInt(gl.VALIDATE_STATUS) == gl.FALSE {
		if infoLog := me.InfoLog(); len(infoLog) > 0 {
			err = fmt.Errorf("Program'%s'.Validate() error: %s\n", me.Name, infoLog)
		}
	}
	return
}

//	Returns true if the specified loc is a valid vertex-attribute location.
func progIsAttrLocation(loc gl.Uint) bool {
	return (loc >= 0) && (loc < progAttrInvalidLoc)
}

//	Returns true if the specified loc is a valid uniform location.
func progIsUnifLocation(loc gl.Int) bool {
	return loc >= 0
}
