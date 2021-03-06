// Opens a GLFW3 window and renders a blue quad and triangle.
package main

import (
	"fmt"
	"runtime"
	"strings"

	gl "github.com/chsc/gogl/gl43"
	glfw "github.com/go-gl/glfw/v3.0/glfw"
)

type geometry struct {
	glVerts          []gl.Float
	glNumVerts       gl.Sizei
	glMode           gl.Enum
	glVertBuf, glVao gl.Uint
}

const (
	srcVertShader = `#version 150 core
in vec3 aPos;
void main() {
	gl_Position = vec4(aPos, 1.0);
}`

	srcFragShader = `#version 150 core
out vec4 oColor;
void main() {
	oColor = vec4(0.0, 0.6, 0.9, 1.0);
}`
)

var (
	faceTri, faceQuad          = &geometry{}, &geometry{}
	useStrictCoreProfile       = (runtime.GOOS == "darwin")
	isFirstLoop                = true
	lastErr                    error
	shaderProg, progPosAttrLoc gl.Uint
)

func compileShaders() (err error) {
	var glStatus gl.Int

	makeShader := func(stage gl.Enum, src string) (shader gl.Uint) {
		shader = gl.CreateShader(stage)
		logLastGlError("gl.CreateShader()")
		glStr := gl.GLStringArray(src)
		defer gl.GLStringArrayFree(glStr)
		gl.ShaderSource(shader, gl.Sizei(len(glStr)), &glStr[0], nil)
		logLastGlError("gl.ShaderSource()")
		gl.CompileShader(shader)
		if gl.GetShaderiv(shader, gl.COMPILE_STATUS, &glStatus); glStatus == 0 {
			panic("Shader compilation failed: " + src)
		}
		return
	}

	glVertShader := makeShader(gl.VERTEX_SHADER, srcVertShader)
	defer gl.DeleteShader(glVertShader)
	glFragShader := makeShader(gl.FRAGMENT_SHADER, srcFragShader)
	defer gl.DeleteShader(glFragShader)

	shaderProg = gl.CreateProgram()
	for stageName, shader := range map[string]gl.Uint{"vert": glVertShader, "frag": glFragShader} {
		gl.AttachShader(shaderProg, shader)
		logLastGlError("gl.AttachShader " + stageName)
		defer gl.DetachShader(shaderProg, shader)
	}

	gl.LinkProgram(shaderProg)
	if gl.GetProgramiv(shaderProg, gl.LINK_STATUS, &glStatus); glStatus == 0 {
		panic("Program linking failed...")
	}

	glAttName := gl.GLString("aPos")
	defer gl.GLStringFree(glAttName)
	progPosAttrLoc = gl.Uint(gl.GetAttribLocation(shaderProg, glAttName))
	logLastGlError("gl.GetAttribLocation()")

	return
}

func deleteGeometry() {
	gl.DeleteBuffers(1, &faceTri.glVertBuf)
	gl.DeleteBuffers(1, &faceQuad.glVertBuf)
	gl.DeleteVertexArrays(1, &faceTri.glVao)
	gl.DeleteVertexArrays(1, &faceQuad.glVao)
}

func onError(errCode glfw.ErrorCode, msg string) {
	panic(fmt.Errorf("GLFW error %v: %s", errCode, msg))
}

func renderGeometry(mesh *geometry) {
	gl.BindVertexArray(mesh.glVao)
	gl.DrawArrays(mesh.glMode, 0, mesh.glNumVerts)
	gl.BindVertexArray(0)
}

func setupGeometry() {
	var (
		triSize  gl.Float = 0.33
		quadSize gl.Float = 0.33
		xoff     gl.Float = 0.5
		z        gl.Float = -1
	)
	faceTri.glMode, faceTri.glNumVerts, faceTri.glVerts = gl.TRIANGLES, 3, []gl.Float{
		0 + xoff, triSize, z,
		-triSize + xoff, -triSize, z,
		triSize + xoff, -triSize, z,
	}
	uploadGeometry(faceTri)
	faceQuad.glMode, faceQuad.glNumVerts, faceQuad.glVerts = gl.TRIANGLE_STRIP, 4, []gl.Float{
		quadSize - xoff, quadSize, z,
		-quadSize - xoff, quadSize, z,
		quadSize - xoff, -quadSize, z,
		-quadSize - xoff, -quadSize, z,
	}
	uploadGeometry(faceQuad)
}

func uploadGeometry(mesh *geometry) {
	gl.GenVertexArrays(1, &mesh.glVao)
	gl.BindVertexArray(mesh.glVao)
	gl.GenBuffers(1, &mesh.glVertBuf)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.glVertBuf)
	gl.EnableVertexAttribArray(progPosAttrLoc)
	gl.VertexAttribPointer(progPosAttrLoc, 3, gl.FLOAT, gl.FALSE, 0, gl.Pointer(nil))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(4*len(mesh.glVerts)), gl.Pointer(&mesh.glVerts[0]), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

func logLastGlError(step string) {
	if errNum := gl.GetError(); errNum != 0 {
		fmt.Printf("ERR %v at step %v\n", errNum, step)
	}
}

func main() {
	runtime.LockOSThread()
	var looping = true
	defer fmt.Println("EXIT")
	// if err = glfw.Init(); err != nil {
	// 	panic(err)
	// }
	if !glfw.Init() {
		panic("glfw.Init() failed!")
	}
	defer glfw.Terminate()
	// glfw.OpenWindowHint(glfw.FsaaSamples, 0)
	glfw.WindowHint(glfw.Samples, 0)
	if useStrictCoreProfile {
		// glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
		// glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 2)
		// glfw.OpenWindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 2)
		glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
	}
	// if err = glfw.OpenWindow(1280, 720, 8, 8, 8, 0, 24, 8, glfw.Windowed); err != nil {
	// 	panic(err)
	// }
	// defer glfw.CloseWindow()
	glfw.WindowHint(glfw.RedBits, 8)
	glfw.WindowHint(glfw.BlueBits, 8)
	glfw.WindowHint(glfw.GreenBits, 8)
	glfw.WindowHint(glfw.AlphaBits, 0)
	glfw.WindowHint(glfw.DepthBits, 24)
	glfw.WindowHint(glfw.StencilBits, 8)
	win, err := glfw.CreateWindow(1280, 720, "Test", nil, nil)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()
	win.MakeContextCurrent()

	// glfw.Enable(glfw.StickyKeys)
	win.SetInputMode(glfw.StickyKeys, glfw.True)
	if err = gl.Init(); (err != nil) && (strings.Index(err.Error(), "VERSION_") < 0) {
		panic(err)
	}
	defer logLastGlError("(post loop)")
	gl.ClearColor(0.3, 0.1, 0.0, 1.0)
	gl.Enable(gl.DEPTH_TEST)
	gl.FrontFace(gl.CCW)
	gl.CullFace(gl.BACK)
	gl.Disable(gl.CULL_FACE)
	if err = compileShaders(); err != nil {
		panic(err)
	}
	setupGeometry()
	defer deleteGeometry()
	logLastGlError("(pre loop)")

	for looping {
		gl.UseProgram(shaderProg)
		if isFirstLoop {
			logLastGlError("gl.UseProgram")
		}

		gl.Viewport(0, 0, 1280, 720)
		if isFirstLoop {
			logLastGlError("gl.ViewPort")
		}

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		if isFirstLoop {
			logLastGlError("gl.Clear")
		}

		renderGeometry(faceTri)
		if isFirstLoop {
			logLastGlError("renderGeometry(faceTri)")
		}

		renderGeometry(faceQuad)
		if isFirstLoop {
			logLastGlError("renderGeometry(faceQuad)")
		}

		// if (glfw.WindowParam(glfw.Opened) != 1) || (glfw.Key(glfw.KeyEsc) == glfw.KeyPress) {
		if win.ShouldClose() || win.GetKey(glfw.KeyEscape) == glfw.Press {
			looping = false
		} else {
			// glfw.SwapBuffers()
			win.SwapBuffers()
			glfw.PollEvents()
		}
		isFirstLoop = false
	}
	logLastGlError("post-loop")
}
