package ugl

import (
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

	Sources struct {
		In  ShaderSources
		Out ShaderSources
	}

	Tag interface{}

	//	All vertex-attributes and their locations mapped for this program object (after the SetAttrLocations() method has been called).
	AttrLocs map[string]gl.Uint

	//	All uniforms and their locations mapped for this program object (after the SetUnifLocations() method has been called).
	UnifLocs map[string]gl.Int

	unif struct {
		cache1i map[gl.Int]gl.Int
		cache1f map[gl.Int]gl.Float
	}
}

func NewProgram(name string) (me *Program) {
	me = new(Program)
	me.Init(name)
	return
}

func (me *Program) Init(name string) {
	const cap = 8
	me.Name, me.AttrLocs, me.UnifLocs = name, make(map[string]gl.Uint, cap), make(map[string]gl.Int, cap)
	me.unif.cache1i = make(map[gl.Int]gl.Int, cap)
	me.unif.cache1f = make(map[gl.Int]gl.Float, cap)
}

//	Creates and compiles Shader stages for the specified sources (and defines) and links them together in this Program.
//	All sources are string pointers and WILL be modified by this method to contain the final real
//	GLSL source string sent to GL as returned by Shader.SetSource(). This may aid diagnostics or be discarded,
//	but a source string should not be re-sent to a subsequent compilation in that modified form.
func (me *Program) CompileAndLinkShaders(defines map[string]interface{}) (err error) {
	var allJobs = [...]shaderJob{
		shaderJob{me.Sources.In.Compute, &me.Sources.Out.Compute, NewComputeShader},
		shaderJob{me.Sources.In.Fragment, &me.Sources.Out.Fragment, NewFragmentShader},
		shaderJob{me.Sources.In.Geometry, &me.Sources.Out.Geometry, NewGeometryShader},
		shaderJob{me.Sources.In.TessCtl, &me.Sources.Out.TessCtl, NewTessCtlShader},
		shaderJob{me.Sources.In.TessEval, &me.Sources.Out.TessEval, NewTessEvalShader},
		shaderJob{me.Sources.In.Vertex, &me.Sources.Out.Vertex, NewVertexShader},
	}
	var shader *Shader
	allShaders := make([]*Shader, 0, 6)
	for i := 0; i < len(allJobs); i++ {
		if shader, err = me.compileShader(&allJobs[i], defines); shader != nil {
			defer shader.Dispose()
			if err == nil {
				allShaders = append(allShaders, shader)
			} else {
				return
			}
		}

	}
	if len(allShaders) == 0 {
		err = errf("Program'%s'.CompileAndLinkShaders() -- no shaders specified.", me.Name)
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

func (me *Program) compileShader(job *shaderJob, defines map[string]interface{}) (shader *Shader, err error) {
	if len(job.srcIn) > 0 {
		if *job.srcOut, shader = job.srcIn, job.ctor(); shader != nil {
			if err = shader.Create(); err != nil {
				return
			}
			if *job.srcOut, err = shader.SetSource(job.srcIn, defines); err != nil {
				return
			}
			if err = shader.Compile(me.Name); err != nil {
				return
			}
		}
	}
	return
}

//	Creates this program object in OpenGL.
func (me *Program) Create() (err error) {
	if me.GlHandle == 0 {
		me.Dispose()
		err, me.GlHandle = Try.CreateProgram()
	}
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
		err = errf("Program'%s'.Link() error: %s\n", me.Name, me.InfoLog())
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

//	Convenience short-hand for gl.Uniform1i, if this Program contains
//	the specified uniform and v0 differs from its current value.
func (me *Program) Uniform1i(name string, v0 gl.Int) {
	if loc, ok := me.UnifLocs[name]; ok && me.unif.cache1i[loc] != v0 {
		me.unif.cache1i[loc] = v0
		gl.Uniform1i(loc, v0)
	}
}

//	Convenience short-hand for gl.Uniform1f, if this Program contains
//	the specified uniform and v0 differs from its current value.
func (me *Program) Uniform1f(name string, v0 gl.Float) {
	if loc, ok := me.UnifLocs[name]; ok && me.unif.cache1f[loc] != v0 {
		me.unif.cache1f[loc] = v0
		gl.Uniform1f(loc, v0)
	}
}

//	Convenience short-hand for gl.Uniform3fv, if this Program contains the specified uniform.
func (me *Program) Uniform3fv(name string, count gl.Sizei, value *gl.Float) {
	if loc, ok := me.UnifLocs[name]; ok {
		gl.Uniform3fv(loc, count, value)
	}
}

//	Convenience short-hand for gl.UniformMatrix4fv, if this Program contains the specified uniform.
func (me *Program) UniformMatrix4fv(name string, count gl.Sizei, transpose gl.Boolean, value *gl.Float) {
	if loc, ok := me.UnifLocs[name]; ok {
		gl.UniformMatrix4fv(loc, count, transpose, value)
	}
}

//	Installs this program object as part of current rendering state.
func (me *Program) Use() {
	Cache.UseProgram(me.GlHandle)
}

//	Validates this program object. This is a convenience short-hand for calling gl.ValidateProgram(),
//	checking gl.VALIDATE_STATUS and obtaining gl.GetProgramInfoLog().
func (me *Program) Validate() (err error) {
	if gl.ValidateProgram(me.GlHandle); me.ParamInt(gl.VALIDATE_STATUS) == gl.FALSE {
		if infoLog := me.InfoLog(); len(infoLog) > 0 {
			err = errf("Program'%s'.Validate() error: %s\n", me.Name, infoLog)
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
