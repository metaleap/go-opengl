package ugl

import (
	"fmt"
	"strings"

	"github.com/metaleap/go-util"

	gl "github.com/metaleap/go-opengl/core"
)

type shaderCtor func() *Shader

type shaderJob struct {
	srcIn  string
	srcOut *string
	ctor   shaderCtor
}

type ShaderSources struct {
	Compute  string
	Fragment string
	Geometry string
	TessCtl  string
	TessEval string
	Vertex   string
}

//	Represents an OpenGL shader object.
type Shader struct {
	//	The OpenGL handle to the underlying shader object.
	//	This is 0 before calling Create() and after calling Dispose().
	GlHandle gl.Uint

	//	The shader stage provided by this shader,
	//	such as for example gl.FRAGMENT_SHADER, gl.VERTEX_SHADER, etc.
	Stage gl.Enum
}

//	Initializes and returns --but does not Create()-- a new Shader with the specified name and shader stage.
func NewShader(glStage gl.Enum) (me *Shader) {
	me = &Shader{Stage: glStage}
	return
}

//	If Support.Glsl.Shaders.ComputeStage is true, initializes and returns --but does not Create()-- a new Shader for the gl.COMPUTE_SHADER Stage with the specified name.
func NewComputeShader() (me *Shader) {
	if Support.Glsl.Shaders.ComputeStage {
		me = NewShader(gl.COMPUTE_SHADER)
	}
	return
}

//	Initializes and returns --but does not Create()-- a new Shader for the gl.FRAGMENT_SHADER Stage with the specified name.
func NewFragmentShader() *Shader {
	return NewShader(gl.FRAGMENT_SHADER)
}

//	Initializes and returns --but does not Create()-- a new Shader for the gl.GEOMETRY_SHADER Stage with the specified name.
func NewGeometryShader() *Shader {
	return NewShader(gl.GEOMETRY_SHADER)
}

//	If Support.Glsl.Shaders.TessellationStages is true, initializes and returns --but does not Create()-- a new Shader for the gl.TESS_CONTROL_SHADER ("Hull") Stage with the specified name.
func NewTessCtlShader() (me *Shader) {
	if Support.Glsl.Shaders.TessellationStages {
		me = NewShader(gl.TESS_CONTROL_SHADER)
	}
	return
}

//	If Support.Glsl.Shaders.TessellationStages is true, initializes and returns --but does not Create()-- a new Shader for the gl.TESS_EVALUATION_SHADER ("Domain") Stage with the specified name.
func NewTessEvalShader() (me *Shader) {
	if Support.Glsl.Shaders.TessellationStages {
		me = NewShader(gl.TESS_EVALUATION_SHADER)
	}
	return
}

//	Initializes and returns --but does not Create()-- a new Shader for the gl.VERTEX_SHADER Stage with the specified name.
func NewVertexShader() *Shader {
	return NewShader(gl.VERTEX_SHADER)
}

//	Attaches this shader object to the specified program object.
func (me *Shader) AttachTo(prog *Program) error {
	return Try.AttachShader(prog.GlHandle, me.GlHandle)
}

//	Returns the name of this shader's stage, for example "GL_VERTEX_SHADER", "GL_COMPUTE_SHADER", etc.
func (me *Shader) StageName() (sn string) {
	return Util.EnumName(me.Stage)
}

//	Compiles this shader object. This is a convenience short-hand for calling gl.CompileShader(),
//	checking gl.COMPILE_STATUS and obtaining gl.GetShaderInfoLog().
func (me *Shader) Compile(name string) (err error) {
	if gl.CompileShader(me.GlHandle); me.ParamInt(gl.COMPILE_STATUS) == gl.FALSE {
		err = errf("%s'%s'.Compile() error: %s\n", me.StageName(), name, me.InfoLog(name))
	}
	return
}

//	Creates this shader object in OpenGL.
func (me *Shader) Create() (err error) {
	me.Dispose()
	err, me.GlHandle = Try.CreateShader(me.Stage)
	return
}

//	Detaches this shader object from the specified program object.
func (me *Shader) DetachFrom(prog *Program) {
	gl.DetachShader(prog.GlHandle, me.GlHandle)
}

//	Deletes this shader object from OpenGL.
func (me *Shader) Dispose() {
	if me.GlHandle != 0 {
		gl.DeleteShader(me.GlHandle)
		me.GlHandle = 0
	}
}

//	Retrieves and returns the content of the current OpenGL info-log for this shader object.
func (me *Shader) InfoLog(name string) string {
	return shaderProgInfoLog(name, me.GlHandle, true)
}

//	Convenience short-hand for gl.GetShaderiv.
func (me *Shader) ParamInt(pname gl.Enum) (iv gl.Int) {
	gl.GetShaderiv(me.GlHandle, pname, &iv)
	return
}

//	Replaces the GLSL source code in this shader object with source,
//	prepending a #version (core) directive with the current Support.Glsl.Version.Num
//	and #define directives for all specified name-value pairs in defines.
func (me *Shader) SetSource(source string, defines map[string]interface{}) (finalRealSrc string, err error) {
	i, lines := 1, make([]string, (len(defines)*5)+3)
	lines[0] = strf("#version %v core\n", Support.Glsl.Version.Num)
	if len(defines) > 0 {
		for dk, dv := range defines {
			lines[i+0] = "#define "
			lines[i+1] = dk
			lines[i+2] = " "
			lines[i+3] = fmt.Sprintf("%v", dv)
			lines[i+4] = "\n"
			i = i + 5
		}
	}
	lines[i] = "#line 1\n"
	lines[i+1] = source
	finalRealSrc = strings.Join(lines, "")
	src := gl.Util.CStringArray(lines...)
	defer gl.Util.CStringArrayFree(src)
	err = Try.ShaderSource(me.GlHandle, gl.Sizei(len(src)), &src[0], nil)
	return
}

func shaderProgInfoLog(name string, glHandle gl.Uint, shader bool) (infoLog string) {
	const l gl.Sizei = 256
	s := gl.Util.CStringAlloc(l)
	defer gl.Util.CStringFree(s)
	if shader {
		gl.GetShaderInfoLog(glHandle, l, nil, s)
	} else {
		gl.GetProgramInfoLog(glHandle, l, nil, s)
	}
	if err := Util.LastError("%s'%s'.InfoLog()", ugo.Ifs(shader, "Shader", "Program"), name); err == nil {
		infoLog = gl.Util.StringFromChar(s)
	} else {
		infoLog = err.Error()
	}
	return
}
