# ugl
--
    import "github.com/go3d/go-opengl/util"

OpenGL utilities and a few sane slim wrappers (without going full-on
"OO"-overboard) for Go (using the GL Core Profile 3.3+ [
http://github.com/go3d/go-opengl/core/ ] bindings).

## Usage

```go
const (
	//	The size of a gl.Float in bytes
	SizeOfGlFloat gl.Sizeiptr = 4

	//	The size of a gl.Uint in bytes
	SizeOfGlUint gl.Sizeiptr = 4
)
```

```go
var (
	Cache GlCache

	//	Provides access to miscellaneous GL functionality.
	Util GlUtil

	//	Provides access to miscellaneous type-specific GL functionality.
	Typed TypeUtils

	//	Provides methods that each wrap a GL function call and immediately check for GL errors to be returned as Go errors.
	Try GlTry
)
```

```go
var (
	//	All the known possible OpenGL core context versions that this package supports.
	KnownVersions = [5]float64{3.3, 4.0, 4.1, 4.2, 4.3}

	//	Contains the well-known extension prefixes that can be omitted
	//	when querying the current OpenGL context via Util.Extension().
	KnownExtensionPrefixes = [...]string{"GL_3DFX_", "GL_3DL_", "GL_AMD_", "GL_APPLE_", "GL_ARB_", "GL_ATI_", "GL_EXT_", "GL_GREMEDY_", "GL_HP_", "GL_I3D_", "GL_IBM_", "GL_INGR_", "GL_INTEL_", "GL_KHR_", "GL_KTX_", "GL_MESA_", "GL_MESAX_", "GL_NV_", "GL_NVX_", "GL_OES_", "GL_OML_", "GL_PGI_", "GL_REND__", "GL_S3_", "GL_SGI_", "GL_SGIS_", "GL_SGIX_", "GL_SUN_", "GL_SUNX_", "GL_WIN_", "WGL_EXT_"}

	//	After you have called ugl.Init() upon successful OpenGL context creation and GL binding initialization,
	//	contains information on what the currently active OpenGL profile supports.
	Support struct {
		//	Informs on whether certain buffer types are supported by the current OpenGL context.
		Buffers struct {
			//	True if OpenGL version is 4.2 or higher.
			AtomicCounter bool

			//	True if OpenGL version is 4.3 or higher.
			DispatchIndirect bool

			//	True if OpenGL version is 4.3 or higher.
			ShaderStorage bool
		}

		//	All extensions supported by the current OpenGL context.
		//	Useful for casual inspection, but to query for the presence of a specific extension, always use Util.Extension(name)
		Extensions []string

		//	Information on the OpenGL version of the current OpenGL context.
		GlVersion struct {
			//	Will be either 0 or any one of the values in KnownVersions.
			Num float64

			//	The major and minor parts of the version in Num.
			MajorMinor [2]int

			//	True if OpenGL version is 3.3 or higher.
			Is33 bool

			//	True if OpenGL version is 4.0 or higher.
			Is40 bool

			//	True if OpenGL version is 4.1 or higher.
			Is41 bool

			//	True if OpenGL version is 4.2 or higher.
			Is42 bool

			//	True if OpenGL version is 4.3 or higher.
			Is43 bool
		}

		//	Information on the GL Shading Language supported by the current OpenGL context.
		Glsl struct {
			//	Info on support for certain shader stages.
			Shaders struct {
				//	True if Support.GlSl.Version.Num is 430 or higher.
				ComputeStage bool

				//	True if Support.GlSl.Version.Num is 400 or higher.
				TessellationStages bool
			}

			//	Information on the GLSL version supported by the current OpenGL context.
			Version struct {
				//	Will be either 0, 330, 400, 410, 420 or 430.
				Num int

				//	True if OpenGL version is 3.3 or higher.
				Is330 bool

				//	True if OpenGL version is 4.0 or higher.
				Is400 bool

				//	True if OpenGL version is 4.1 or higher.
				Is410 bool

				//	True if OpenGL version is 4.2 or higher.
				Is420 bool

				//	True if OpenGL version is 4.3 or higher.
				Is430 bool
			}
		}

		//	Informs on texturing/sampling-related support by the current OpenGL context.
		Textures struct {
			//	True if OpenGL version is 4.2 or higher.
			Immutable bool

			//	The maximum supported texture filtering anisotropy.
			//	If the current OpenGL context does not support anisotropic texture filtering, the
			//	value will be 0; else the value is guaranteed to be at least 1, and is typically 16.
			MaxFilterAnisotropy gl.Float

			//	The maximum supported texture MIP-map LOD bias.
			//	(The minimum would be the negative equivalent of this value.)
			MaxMipLoadBias gl.Float

			ImageUnits struct {
				MaxCombined gl.Int
				MaxCompute  gl.Int
				MaxFragment gl.Int
				MaxGeometry gl.Int
				MaxVertex   gl.Int
			}

			//	Experimental.
			StreamUpdatesViaPixelBuffer bool
		}
	}
)
```

```go
var (
	PtrNil = gl.Ptr(nil)
)
```

#### func  Init

```go
func Init()
```
DO call this immediately upon successful OpenGL context initialization to
prepare this package for use. Fully populates all fields and sub-fields in the
Support package global.

#### func  LogLastError

```go
func LogLastError(stepFmt string, stepFmtArgs ...interface{})
```
log.Println()s the error returned by Util.Error(), if any.

#### func  VersionMajorMinor

```go
func VersionMajorMinor(v float64) (major, minor int)
```
Returns the major and minor parts of the specified version number.

#### func  VersionMatch

```go
func VersionMatch(v float64) bool
```
Returns true if GlVersion is greater than or equal to v.

#### type Buffer

```go
type Buffer struct {
	//	The OpenGL handle to this buffer object.
	//	This is 0 before calling Recreate() and after calling Dispose().
	GlHandle gl.Uint

	//	The current binding target for this buffer object
	//	(such as for example gl.ARRAY_BUFFER, gl.ELEMENT_ARRAY_BUFFER etc.).
	GlTarget gl.Enum
}
```

Represents an OpenGL buffer object.

#### func  NewArrayBuffer

```go
func NewArrayBuffer() *Buffer
```
Initializes and returns --but does not Recreate()-- a new Buffer initialized
with a GlTarget of gl.ARRAY_BUFFER.

#### func  NewAtomicCounterBuffer

```go
func NewAtomicCounterBuffer() (me *Buffer)
```
If Support.Buffers.AtomicCounter is true, initializes and returns --but does not
Recreate()-- a new Buffer initialized with a GlTarget of
gl.ATOMIC_COUNTER_BUFFER.

#### func  NewBuffer

```go
func NewBuffer(glTarget gl.Enum) (me *Buffer)
```
Initializes and returns --but does not Recreate()-- a new Buffer not yet
initialized with any valid GlTarget.

#### func  NewCopyReadBuffer

```go
func NewCopyReadBuffer() *Buffer
```
Initializes and returns --but does not Recreate()-- a new Buffer initialized
with a GlTarget of gl.COPY_READ_BUFFER.

#### func  NewCopyWriteBuffer

```go
func NewCopyWriteBuffer() *Buffer
```
Initializes and returns --but does not Recreate()-- a new Buffer initialized
with a GlTarget of gl.COPY_WRITE_BUFFER.

#### func  NewDispatchIndirectBuffer

```go
func NewDispatchIndirectBuffer() (me *Buffer)
```
If Support.Buffers.DispatchIndirect is true, initializes and returns --but does
not Recreate()-- a new Buffer initialized with a GlTarget of
gl.DISPATCH_INDIRECT_BUFFER.

#### func  NewDrawIndirectBuffer

```go
func NewDrawIndirectBuffer() *Buffer
```
Initializes and returns --but does not Recreate()-- a new Buffer initialized
with a GlTarget of gl.DRAW_INDIRECT_BUFFER.

#### func  NewElementArrayBuffer

```go
func NewElementArrayBuffer() *Buffer
```
Initializes and returns --but does not Recreate()-- a new Buffer initialized
with a GlTarget of gl.ELEMENT_ARRAY_BUFFER.

#### func  NewPixelPackBuffer

```go
func NewPixelPackBuffer() *Buffer
```
Initializes and returns --but does not Recreate()-- a new Buffer initialized
with a GlTarget of gl.PIXEL_PACK_BUFFER.

#### func  NewPixelUnpackBuffer

```go
func NewPixelUnpackBuffer() *Buffer
```
Initializes and returns --but does not Recreate()-- a new Buffer initialized
with a GlTarget of gl.PIXEL_UNPACK_BUFFER.

#### func  NewShaderStorageBuffer

```go
func NewShaderStorageBuffer() (me *Buffer)
```
If Support.Buffers.ShaderStorage is true, initializes and returns --but does not
Recreate()-- a new Buffer initialized with a GlTarget of
gl.SHADER_STORAGE_BUFFER.

#### func  NewTextureBuffer

```go
func NewTextureBuffer() *Buffer
```
Initializes and returns --but does not Recreate()-- a new Buffer initialized
with a GlTarget of gl.TEXTURE_BUFFER.

#### func  NewTransformFeedbackBuffer

```go
func NewTransformFeedbackBuffer() *Buffer
```
Initializes and returns --but does not Recreate()-- a new Buffer initialized
with a GlTarget of gl.TRANSFORM_FEEDBACK_BUFFER.

#### func  NewUniformBuffer

```go
func NewUniformBuffer() *Buffer
```
Initializes and returns --but does not Recreate()-- a new Buffer initialized
with a GlTarget of gl.UNIFORM_BUFFER.

#### func (*Buffer) Bind

```go
func (me *Buffer) Bind()
```
Binds this buffer object to its current GlTarget.

#### func (*Buffer) BindBase

```go
func (me *Buffer) BindBase(index gl.Uint)
```
Convenience short-hand for gl.BindBufferBase(). Only permissible if me.GlTarget
is gl.TRANSFORM_FEEDBACK_BUFFER or gl.UNIFORM_BUFFER.

#### func (*Buffer) BindRange

```go
func (me *Buffer) BindRange(index gl.Uint, offset gl.Intptr, size gl.Sizeiptr)
```
Convenience short-hand for gl.BindBufferRange(). Only permissible if me.GlTarget
is gl.TRANSFORM_FEEDBACK_BUFFER or gl.UNIFORM_BUFFER.

#### func (*Buffer) Dispose

```go
func (me *Buffer) Dispose()
```
Deletes this buffer object.

#### func (*Buffer) Map

```go
func (me *Buffer) Map(read, write bool) *gl.Ptr
```
Maps to the client's address space the entire data store of whatever buffer
object is currently bound to me.GlTarget.

#### func (*Buffer) Recreate

```go
func (me *Buffer) Recreate(newGlTarget gl.Enum, sizeInBytes gl.Sizeiptr, initialData gl.Ptr, usageHint gl.Enum) (err error)
```
Deletes this buffer object if applicable, generates it anew, and allocates its
data store as specified, populating it with the specified initialData if any.
GlTarget is changed to newGlTarget unless 0 is specified.

#### func (*Buffer) SubData

```go
func (me *Buffer) SubData(offset gl.Intptr, size gl.Sizeiptr, data gl.Ptr) (err error)
```
Buffers the specified data into this buffer object's data store at the specified
offset.

#### func (*Buffer) Unbind

```go
func (me *Buffer) Unbind()
```
Unbinds whatever buffer object is currently bound to me.GlTarget.

#### func (*Buffer) Unmap

```go
func (me *Buffer) Unmap() gl.Boolean
```
Convenience short-hand for gl.UnmapBuffer(me.GlTarget)

#### type Framebuffer

```go
type Framebuffer struct {
	//	Set in Create(), this is either gl.READ_FRAMEBUFFER or gl.FRAMEBUFFER
	GlTarget gl.Enum

	//	The OpenGL handle to this framebuffer object.
	//	This is 0 before calling Create() and after calling Dispose().
	GlHandle gl.Uint
}
```

Encapsulates an OpenGL framebuffer object.

#### func (*Framebuffer) AttachRenderbuffer

```go
func (me *Framebuffer) AttachRenderbuffer() (rb *FramebufferRenderbuffer)
```
Attaches the specified renderbuffer to this Framebuffer.

#### func (*Framebuffer) AttachRendertexture

```go
func (me *Framebuffer) AttachRendertexture() (tex *FramebufferRendertexture)
```
Attaches the specified render texture to this Framebuffer.

#### func (*Framebuffer) Bind

```go
func (me *Framebuffer) Bind()
```
Binds this framebuffer object.

#### func (*Framebuffer) Create

```go
func (me *Framebuffer) Create(width, height gl.Sizei, read bool)
```
Creates this Framebuffer in OpenGL sized as specified.

#### func (*Framebuffer) Dispose

```go
func (me *Framebuffer) Dispose()
```
Deletes this Framebuffer and all attached render textures and render buffers
from OpenGL.

#### func (*Framebuffer) RenderTextureHandle

```go
func (me *Framebuffer) RenderTextureHandle(index int) gl.Uint
```
Returns the FramebufferRendertexture attached to me at the specified index.

#### func (*Framebuffer) Resize

```go
func (me *Framebuffer) Resize(width, height gl.Sizei) (resized bool)
```
Resizes this Framebuffer and all attached render textures and render buffers.

#### func (*Framebuffer) Status

```go
func (me *Framebuffer) Status() gl.Enum
```
The current status of this Framebuffer as returned by
gl.CheckFramebufferStatus().

#### func (*Framebuffer) Unbind

```go
func (me *Framebuffer) Unbind()
```
Unbinds whatever Framebuffer is currently bound, and binds the "default"
framebuffer (typically, the window's drawing surface).

#### type FramebufferRenderbuffer

```go
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
```

Represents a Renderbuffer that can be attached to a Framebuffer.

#### func (*FramebufferRenderbuffer) Bind

```go
func (me *FramebufferRenderbuffer) Bind()
```
Binds this render buffer object.

#### func (*FramebufferRenderbuffer) Init

```go
func (me *FramebufferRenderbuffer) Init()
```

#### func (*FramebufferRenderbuffer) Unbind

```go
func (me *FramebufferRenderbuffer) Unbind()
```
Unbinds whatever render buffer object is currently bound, if any.

#### type FramebufferRendertexture

```go
type FramebufferRendertexture struct {
	//	The texture object being rendered to.
	Texture2D

	//	The attachment point for this render texture object.
	//	Defaults to gl.COLOR_ATTACHMENT0.
	Attachment gl.Enum
}
```

Represents a render-to-texture object that can be attached to a Framebuffer.

#### func (*FramebufferRendertexture) Init

```go
func (me *FramebufferRendertexture) Init()
```

#### type GlCache

```go
type GlCache struct {
}
```


#### func (*GlCache) ActiveTexture

```go
func (_ *GlCache) ActiveTexture(texture gl.Enum)
```

#### func (*GlCache) BindSampler

```go
func (_ *GlCache) BindSampler(textureUnit, glHandle gl.Uint)
```

#### func (*GlCache) BindTextureTo

```go
func (_ *GlCache) BindTextureTo(textureUnit, glHandle gl.Uint, glTarget gl.Enum)
```

#### func (*GlCache) BindVertexArray

```go
func (_ *GlCache) BindVertexArray(glHandle gl.Uint)
```

#### func (*GlCache) UseProgram

```go
func (_ *GlCache) UseProgram(glHandle gl.Uint)
```

#### type GlMat3

```go
type GlMat3 [9]gl.Float
```

Represents a 3x3 matrix of 9 32-bit gl.Float values

#### func  NewGlMat3

```go
func NewGlMat3(mat *unum.Mat3) (glMat *GlMat3)
```
Instantiates and returns a new 3x3 matrix (32-bit) from the specified 3x3 matrix
(64-bit).

#### func (*GlMat3) Load

```go
func (me *GlMat3) Load(mat *unum.Mat3)
```
Sets all values in this 3x3 matrix (32-bit) from the specified 3x3 matrix
(64-bit).

#### type GlMat4

```go
type GlMat4 [16]gl.Float
```

Represents a 3x3 matrix of 16 32-bit gl.Float values

#### func  NewGlMat4

```go
func NewGlMat4(mat *unum.Mat4) (glMat *GlMat4)
```
Instantiates and returns a new 4x4 matrix (32-bit) from the specified 4x4 matrix
(64-bit).

#### func (*GlMat4) Load

```go
func (me *GlMat4) Load(mat *unum.Mat4)
```
Sets all values in this 4x4 matrix (32-bit) from the specified 4x4 matrix
(64-bit).

#### type GlTry

```go
type GlTry struct {
}
```

A singleton type, only used for the package-global ugl.Try variable. (Methods
generated with github.com/go3d/go-opengl/cmd/gen-opengl-bindings.)

#### func (GlTry) AttachShader

```go
func (_ GlTry) AttachShader(program gl.Uint, shader gl.Uint) (err error)
```
Calls gl.AttachShader() and yields the err returned by Util.LastError(), if any.

#### func (GlTry) BufferData

```go
func (_ GlTry) BufferData(target gl.Enum, size gl.Sizeiptr, data gl.Ptr, usage gl.Enum) (err error)
```
Calls gl.BufferData() and yields the err returned by Util.LastError(), if any.

#### func (GlTry) BufferSubData

```go
func (_ GlTry) BufferSubData(target gl.Enum, offset gl.Intptr, size gl.Sizeiptr, data gl.Ptr) (err error)
```
Calls gl.BufferSubData() and yields the err returned by Util.LastError(), if
any.

#### func (GlTry) CreateProgram

```go
func (_ GlTry) CreateProgram() (err error, tryRetVal__ gl.Uint)
```
Calls gl.CreateProgram() and yields the err returned by Util.LastError(), if
any.

#### func (GlTry) CreateShader

```go
func (_ GlTry) CreateShader(type_ gl.Enum) (err error, tryRetVal__ gl.Uint)
```
Calls gl.CreateShader() and yields the err returned by Util.LastError(), if any.

#### func (GlTry) GenBuffers

```go
func (_ GlTry) GenBuffers(n gl.Sizei, buffers *gl.Uint) (err error)
```
Calls gl.GenBuffers() and yields the err returned by Util.LastError(), if any.

#### func (GlTry) GenTextures

```go
func (_ GlTry) GenTextures(n gl.Sizei, textures *gl.Uint) (err error)
```
Calls gl.GenTextures() and yields the err returned by Util.LastError(), if any.

#### func (GlTry) GenVertexArrays

```go
func (_ GlTry) GenVertexArrays(n gl.Sizei, arrays *gl.Uint) (err error)
```
Calls gl.GenVertexArrays() and yields the err returned by Util.LastError(), if
any.

#### func (GlTry) GenerateMipmap

```go
func (_ GlTry) GenerateMipmap(target gl.Enum) (err error)
```
Calls gl.GenerateMipmap() and yields the err returned by Util.LastError(), if
any.

#### func (GlTry) ShaderSource

```go
func (_ GlTry) ShaderSource(shader gl.Uint, count gl.Sizei, string_ **gl.Char, length *gl.Int) (err error)
```
Calls gl.ShaderSource() and yields the err returned by Util.LastError(), if any.

#### func (GlTry) TexImage2D

```go
func (_ GlTry) TexImage2D(target gl.Enum, level gl.Int, internalformat gl.Int, width gl.Sizei, height gl.Sizei, border gl.Int, format gl.Enum, type_ gl.Enum, pixels gl.Ptr) (err error)
```
Calls gl.TexImage2D() and yields the err returned by Util.LastError(), if any.

#### func (GlTry) TexStorage2D

```go
func (_ GlTry) TexStorage2D(target gl.Enum, levels gl.Sizei, internalformat gl.Enum, width gl.Sizei, height gl.Sizei) (err error)
```
Calls gl.TexStorage2D() and yields the err returned by Util.LastError(), if any.

#### func (GlTry) TexSubImage2D

```go
func (_ GlTry) TexSubImage2D(target gl.Enum, level gl.Int, xoffset gl.Int, yoffset gl.Int, width gl.Sizei, height gl.Sizei, format gl.Enum, type_ gl.Enum, pixels gl.Ptr) (err error)
```
Calls gl.TexSubImage2D() and yields the err returned by Util.LastError(), if
any.

#### type GlUtil

```go
type GlUtil struct {
}
```

A singleton type, only used for the package-global ugl.Util variable.

#### func (GlUtil) ConnInfo

```go
func (_ GlUtil) ConnInfo() string
```
Returns an OpenGL connection-info string of the format "OpenGL {version} @
{vendor} {renderer} (GLSL: {version})" Example: "OpenGL 4.2.0 @ NVIDIA
Corporation Quadro 5010M/PCIe/SSE2 (GLSL: 4.20 NVIDIA via Cg compiler)"

#### func (GlUtil) EnumName

```go
func (_ GlUtil) EnumName(enum gl.Enum) (name string)
```
Returns the name of the specified enum. Out of the over 1290 possible enum
values, currently only supports error codes and shader stages.

#### func (GlUtil) Extension

```go
func (_ GlUtil) Extension(name string) bool
```
Returns true if the specified extension (ignoring upper/lower-case) is supported
by the current OpenGL context, as per the global Support.Extensions variable. If
the extension uses a well-known prefix contained in KnownExtensionPrefixes, it
can be omitted, ie. you can just specify "texture_filter_anisotropic" instead of
"GL_EXT_texture_filter_anisotropic".

#### func (GlUtil) LastError

```go
func (_ GlUtil) LastError(msgFmt string, fmtArgs ...interface{}) (err error)
```
If the GL error flag is currently set, returns an error with the specified
message and the GL error flag(s).

#### func (GlUtil) Str

```go
func (_ GlUtil) Str(name gl.Enum) string
```
Returns the specified OpenGL string. Example: GlStr(gl.VENDOR) = "NVIDIA
Corporation"

#### func (GlUtil) Stri

```go
func (_ GlUtil) Stri(name gl.Enum, i gl.Uint) string
```
Returns the specified OpenGL string at the specified index. For example,
GlStri(gl.EXTENSIONS, 0) returns the name of the "first" of all supported
extensions.

#### func (GlUtil) Val

```go
func (_ GlUtil) Val(name gl.Enum) (val gl.Int)
```
Returns the specified integer as returned by gl.GetIntegerv(). Example: numExts
:= GlVal(gl.NUM_EXTENSIONS)

#### func (GlUtil) Vals

```go
func (_ GlUtil) Vals(name gl.Enum, n uint) (vals []gl.Int)
```
Returns the specified n integers as returned by gl.GetIntegerv().

#### type GlVec3

```go
type GlVec3 [3]gl.Float
```

Represents a 3-dimensional vector (32-bit per component)

#### func (*GlVec3) Set

```go
func (me *GlVec3) Set(vals ...gl.Float)
```
Sets me to the first 3-or-less vals.

#### type GlVec4

```go
type GlVec4 [4]gl.Float
```

Represents a quaternion or 4-dimensional vector (32-bit per component)

#### func (*GlVec4) Set

```go
func (me *GlVec4) Set(vals ...gl.Float)
```
Sets me to the first 4-or-less vals.

#### type Program

```go
type Program struct {
	//	The OpenGL handle to the underlying program object.
	//	This is 0 before calling Create() and after calling Dispose().
	GlHandle gl.Uint

	//	An arbitrary numeric identifier that is of no use to OpenGL.
	//	This is set by the parent ProgramManager, if any.
	Index int

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
}
```

Represents an OpenGL program object.

#### func  NewProgram

```go
func NewProgram(index int, name string) (me *Program)
```

#### func (*Program) CompileAndLinkShaders

```go
func (me *Program) CompileAndLinkShaders(defines map[string]interface{}) (err error)
```
Creates and compiles Shader stages for the specified sources (and defines) and
links them together in this Program. All sources are string pointers and WILL be
modified by this method to contain the final real GLSL source string sent to GL
as returned by Shader.SetSource(). This may aid diagnostics or be discarded, but
a source string should not be re-sent to a subsequent compilation in that
modified form.

#### func (*Program) Create

```go
func (me *Program) Create() (err error)
```
Creates this program object in OpenGL.

#### func (*Program) Dispose

```go
func (me *Program) Dispose()
```
Deletes this program object from OpenGL.

#### func (*Program) HasAttr

```go
func (me *Program) HasAttr(name string) (ok bool)
```
Returns true if the specified attribute name is present as a key in me.AttrLocs.

#### func (*Program) HasUnif

```go
func (me *Program) HasUnif(name string) (ok bool)
```
Returns true if the specified uniform name is present as a key in me.UnifLocs.

#### func (*Program) InfoLog

```go
func (me *Program) InfoLog() string
```
Retrieves and returns the content of the current OpenGL info-log for this shader
object.

#### func (*Program) Init

```go
func (me *Program) Init(index int, name string)
```

#### func (*Program) Link

```go
func (me *Program) Link() (err error)
```
Links this program object. This is a convenience short-hand for calling
gl.LinkProgram(), checking gl.LINK_STATUS and obtaining gl.GetProgramInfoLog().

#### func (*Program) ParamInt

```go
func (me *Program) ParamInt(pname gl.Enum) (iv gl.Int)
```
Convenience short-hand for gl.GetProgramiv.

#### func (*Program) SetAttrLocations

```go
func (me *Program) SetAttrLocations(attrNames ...string) (err error)
```
Stores the locations for the specified vertex-attributes in the AttrLocs field.
If an attribute in attrNames isn't found, it will not be added to AttrLocs but
this may not result in an error. err is the result of LastError() after each
call to gl.GetAttribLocation(), the method immediately returns on the first
error encountered.

#### func (*Program) SetUnifLocations

```go
func (me *Program) SetUnifLocations(unifNames ...string) (err error)
```
Stores the locations for the specified uniforms in the UnifLocs field. If a
uniform in unifNames isn't found, it will not be added to UnifLocs but this may
not result in an error. err is the result of LastError() after each call to
gl.GetUniformLocation(), the method immediately returns on the first error
encountered.

#### func (*Program) Uniform1f

```go
func (me *Program) Uniform1f(name string, v0 gl.Float)
```
Convenience short-hand for gl.Uniform1f, if this Program contains the specified
uniform and v0 differs from its current value.

#### func (*Program) Uniform1i

```go
func (me *Program) Uniform1i(name string, v0 gl.Int)
```
Convenience short-hand for gl.Uniform1i, if this Program contains the specified
uniform and v0 differs from its current value.

#### func (*Program) Uniform3fv

```go
func (me *Program) Uniform3fv(name string, count gl.Sizei, value *gl.Float)
```
Convenience short-hand for gl.Uniform3fv, if this Program contains the specified
uniform.

#### func (*Program) UniformMat4

```go
func (me *Program) UniformMat4(name string, mat4 *GlMat4)
```

#### func (*Program) UniformMatrix4fv

```go
func (me *Program) UniformMatrix4fv(name string, count gl.Sizei, transpose gl.Boolean, value *gl.Float)
```
Convenience short-hand for gl.UniformMatrix4fv, if this Program contains the
specified uniform.

#### func (*Program) UniformVec3

```go
func (me *Program) UniformVec3(name string, vec *GlVec3)
```

#### func (*Program) Use

```go
func (me *Program) Use()
```
Installs this program object as part of current rendering state.

#### func (*Program) Validate

```go
func (me *Program) Validate() (err error)
```
Validates this program object. This is a convenience short-hand for calling
gl.ValidateProgram(), checking gl.VALIDATE_STATUS and obtaining
gl.GetProgramInfoLog().

#### type ProgramManager

```go
type ProgramManager struct {
	All []Program
}
```


#### func (*ProgramManager) AddNew

```go
func (me *ProgramManager) AddNew(name string) (prog *Program)
```

#### func (*ProgramManager) Dispose

```go
func (me *ProgramManager) Dispose()
```

#### func (*ProgramManager) Get

```go
func (me *ProgramManager) Get(name string) (prog *Program)
```

#### func (*ProgramManager) Index

```go
func (me *ProgramManager) Index(name string) (index int)
```

#### func (*ProgramManager) Init

```go
func (me *ProgramManager) Init()
```

#### func (*ProgramManager) MakeProgramsFromRawSources

```go
func (me *ProgramManager) MakeProgramsFromRawSources(defines map[string]interface{}, forceAll bool, forceSome ...string) (dur time.Duration, progsMade []bool, err error)
```

#### type RenderStates

```go
type RenderStates struct {
}
```

Encapsulates various states in the OpenGL state machine, also providing toggles
for some.

RenderStates is designed to facilitate lazy state changes by caching all render
states: so unless you call any of the ForceFoo() methods, state changes are only
propagated to OpenGL if there is an actual change from the last time that state
was set (via any RenderStates instance -- being 'singleton' in design, they all
utilize the same backing cache).

#### func (RenderStates) Apply

```go
func (me RenderStates) Apply(bag *RenderStatesBag)
```
Calls the applicable me.SetFoo() methods for all render states in bag (except
those in bag.Other).

#### func (RenderStates) DisableBlending

```go
func (me RenderStates) DisableBlending()
```
Disables blending only if it is currently enabled.

#### func (RenderStates) DisableDepthTest

```go
func (me RenderStates) DisableDepthTest()
```
Disables depth-testing only if it is currently enabled.

#### func (RenderStates) DisableFaceCulling

```go
func (me RenderStates) DisableFaceCulling()
```
Disables face-culling only if it is currently enabled.

#### func (RenderStates) DisableFramebufferSrgb

```go
func (me RenderStates) DisableFramebufferSrgb()
```
Disables SRGB-framebuffer only if it is currently enabled.

#### func (RenderStates) DisableScissorTest

```go
func (me RenderStates) DisableScissorTest()
```
Disables scissor-testing only if it is currently enabled.

#### func (RenderStates) DisableStencilTest

```go
func (me RenderStates) DisableStencilTest()
```
Disables stencil-testing only if it is currently enabled.

#### func (RenderStates) EnableBlending

```go
func (me RenderStates) EnableBlending()
```
Enables blending only if it is currently disabled.

#### func (RenderStates) EnableDepthTest

```go
func (me RenderStates) EnableDepthTest()
```
Enables depth-testing only if it is currently disabled.

#### func (RenderStates) EnableFaceCulling

```go
func (me RenderStates) EnableFaceCulling()
```
Enables face-culling only if it is currently disabled.

#### func (RenderStates) EnableFramebufferSrgb

```go
func (me RenderStates) EnableFramebufferSrgb()
```
Enables SRGB-framebuffer only if it is currently disabled.

#### func (RenderStates) EnableScissorTest

```go
func (me RenderStates) EnableScissorTest()
```
Enables scissor-testing only if it is currently disabled.

#### func (RenderStates) EnableStencilTest

```go
func (me RenderStates) EnableStencilTest()
```
Enables stencil-testing only if it is currently disabled.

#### func (RenderStates) ForceClearColor

```go
func (_ RenderStates) ForceClearColor(rgba GlVec4)
```
Sets the OpenGL clear-color to the specified rgba.

#### func (RenderStates) ForceDisableBlending

```go
func (_ RenderStates) ForceDisableBlending()
```
Deactivates blending.

#### func (RenderStates) ForceDisableDepthTest

```go
func (_ RenderStates) ForceDisableDepthTest()
```
Deactivates depth-testing.

#### func (RenderStates) ForceDisableFaceCulling

```go
func (_ RenderStates) ForceDisableFaceCulling()
```
Deactivates face-culling.

#### func (RenderStates) ForceDisableFramebufferSrgb

```go
func (_ RenderStates) ForceDisableFramebufferSrgb()
```
Deactivates SRGB-framebuffer.

#### func (RenderStates) ForceDisableScissorTest

```go
func (_ RenderStates) ForceDisableScissorTest()
```
Deactivates scissor-testing.

#### func (RenderStates) ForceDisableStencilTest

```go
func (_ RenderStates) ForceDisableStencilTest()
```
Deactivates stencil-testing.

#### func (RenderStates) ForceEnableBlending

```go
func (_ RenderStates) ForceEnableBlending()
```
Activates blending.

#### func (RenderStates) ForceEnableDepthTest

```go
func (_ RenderStates) ForceEnableDepthTest()
```
Activates depth-testing.

#### func (RenderStates) ForceEnableFaceCulling

```go
func (_ RenderStates) ForceEnableFaceCulling()
```
Activates face-culling.

#### func (RenderStates) ForceEnableFramebufferSrgb

```go
func (_ RenderStates) ForceEnableFramebufferSrgb()
```
Activates SRGB-framebuffer.

#### func (RenderStates) ForceEnableScissorTest

```go
func (_ RenderStates) ForceEnableScissorTest()
```
Activates scissor-testing.

#### func (RenderStates) ForceEnableStencilTest

```go
func (_ RenderStates) ForceEnableStencilTest()
```
Activates stencil-testing.

#### func (RenderStates) SetBlending

```go
func (me RenderStates) SetBlending(newBlending bool)
```
Activates or deactivates blending.

#### func (RenderStates) SetClearColor

```go
func (me RenderStates) SetClearColor(rgba GlVec4)
```
Sets the OpenGL clear-color if rgba differs from the currently cached value.

#### func (RenderStates) SetDepthTest

```go
func (me RenderStates) SetDepthTest(newDepthTest bool)
```
Activates or deactivates depth-testing.

#### func (RenderStates) SetFaceCulling

```go
func (me RenderStates) SetFaceCulling(newFaceCulling bool)
```
Activates or deactivates face-culling.

#### func (RenderStates) SetFramebufferSrgb

```go
func (me RenderStates) SetFramebufferSrgb(newFramebufferSrgb bool)
```
Activates or deactivates SRGB-framebuffer.

#### func (RenderStates) SetScissorTest

```go
func (me RenderStates) SetScissorTest(newScissorTest bool)
```
Activates or deactivates scissor-testing.

#### func (RenderStates) SetStencilTest

```go
func (me RenderStates) SetStencilTest(newStencilTest bool)
```
Activates or deactivates stencil-testing.

#### func (RenderStates) ToggleBlending

```go
func (me RenderStates) ToggleBlending()
```
Toggles blending.

#### func (RenderStates) ToggleDepthTest

```go
func (me RenderStates) ToggleDepthTest()
```
Toggles depth-testing.

#### func (RenderStates) ToggleFaceCulling

```go
func (me RenderStates) ToggleFaceCulling()
```
Toggles face-culling.

#### func (RenderStates) ToggleFramebufferSrgb

```go
func (me RenderStates) ToggleFramebufferSrgb()
```
Toggles SRGB-framebuffer.

#### func (RenderStates) ToggleScissorTest

```go
func (me RenderStates) ToggleScissorTest()
```
Toggles scissor-testing.

#### func (RenderStates) ToggleStencilTest

```go
func (me RenderStates) ToggleStencilTest()
```
Toggles stencil-testing.

#### type RenderStatesBag

```go
type RenderStatesBag struct {
	Blending    bool
	DepthTest   bool
	FaceCulling bool
	StencilTest bool
	ClearColor  GlVec4
	Other       struct {
		FramebufferSrgb bool
		ClearBits       gl.Bitfield
	}
}
```

Encapsulates a particular combination of render-states that RenderStates can
Apply().

#### type Sampler

```go
type Sampler struct {
	//	The OpenGL handle to this sampler object.
	//	This is 0 before calling Create() and after calling Dispose().
	GlHandle gl.Uint
}
```

Stores sampling parameters for texture accesses.

#### func (Sampler) Bind

```go
func (me Sampler) Bind(textureUnit gl.Uint)
```
Binds this sampler object to the specified textureUnit.

#### func (*Sampler) Create

```go
func (me *Sampler) Create() *Sampler
```
Creates this sampler object.

#### func (Sampler) DepthCompareFunc

```go
func (me Sampler) DepthCompareFunc() gl.Int
```
Returns a single-valued texture comparison function, a symbolic constant. The
initial value is GL_LEQUAL.

#### func (Sampler) DepthCompareMode

```go
func (me Sampler) DepthCompareMode() gl.Int
```
Returns a single-valued texture comparison mode, a symbolic constant. The
initial value is gl.NONE.

#### func (*Sampler) DisableAllFiltering

```go
func (me *Sampler) DisableAllFiltering(allowMip bool) *Sampler
```
Disables linear and anisotropic filtering on this sampler. If allowMip is true,
minified sampling still blends between MIP map levels (just without linear
filtering).

#### func (*Sampler) Dispose

```go
func (me *Sampler) Dispose()
```
Deletes this sampler object.

#### func (*Sampler) EnableFullFiltering

```go
func (me *Sampler) EnableFullFiltering(allowMip bool, maxAniso gl.Float) *Sampler
```
Enables linear and anisotropic filtering on this sampler. If allowMip is true,
minified sampling blends between MIP map levels linearly filtered.

#### func (Sampler) FilterMag

```go
func (me Sampler) FilterMag() gl.Int
```
Returns the single-valued texture magnification filter, a symbolic constant. The
initial value is GL_LINEAR.

#### func (Sampler) FilterMaxAnisotropy

```go
func (me Sampler) FilterMaxAnisotropy() (val gl.Float)
```
Returns filtering anisotropy. 0 if anisotropic filtering isn't supported, 1 if
anisotropic filtering isn't activated for this Sampler, else a value between 1
and Support.Textures.MaxFilterAnisotropy.

#### func (Sampler) FilterMin

```go
func (me Sampler) FilterMin() gl.Int
```
Returns the single-valued texture minification filter, a symbolic constant. The
initial value is GL_NEAREST_MIPMAP_LINEAR.

#### func (Sampler) MipLodBias

```go
func (me Sampler) MipLodBias() gl.Float
```
The mipmap image selection process can be adjusted coarsely by using this
sampling parameter.

#### func (Sampler) MipLodMax

```go
func (me Sampler) MipLodMax() gl.Float
```
Returns the single-valued texture maximum level-of-detail value. The initial
value is 1000.

#### func (Sampler) MipLodMin

```go
func (me Sampler) MipLodMin() gl.Float
```
Returns the single-valued texture minimum level-of-detail value. The initial
value is -1000.

#### func (Sampler) ParamFloat

```go
func (me Sampler) ParamFloat(param gl.Enum) (val gl.Float)
```
A convenience short-hand to invoke gl.GetSamplerParameterfv for this sampler
object.

#### func (Sampler) ParamInt

```go
func (me Sampler) ParamInt(param gl.Enum) (val gl.Int)
```
A convenience short-hand to invoke gl.GetSamplerParameteriv for this sampler
object.

#### func (Sampler) SetDepthCompareFunc

```go
func (me Sampler) SetDepthCompareFunc(val gl.Int)
```
Specifies the comparison operator used when DepthCompareMode() is set to
gl.COMPARE_REF_TO_TEXTURE. val may be one of: gl.LEQUAL, gl.GEQUAL, gl.LESS,
gl.GREATER, gl.EQUAL, gl.NOTEQUAL, gl.ALWAYS, gl.NEVER.

#### func (Sampler) SetDepthCompareMode

```go
func (me Sampler) SetDepthCompareMode(val gl.Int)
```
Specifies the texture comparison mode for currently bound textures whose
internal format is gl.DEPTH_COMPONENT_*. val may be either gl.NONE or
gl.COMPARE_REF_TO_TEXTURE.

#### func (Sampler) SetFilterMag

```go
func (me Sampler) SetFilterMag(val gl.Int)
```
Sets the texture magnification function to either GL_NEAREST or GL_LINEAR.

#### func (Sampler) SetFilterMaxAnisotropy

```go
func (me Sampler) SetFilterMaxAnisotropy(val gl.Float)
```
Sets filtering anisotropy, if supported. val will be clamped between 1 and
Support.Textures.MaxFilterAnisotropy. Only a val greater than 1 will activate
anisotropic filtering.

#### func (Sampler) SetFilterMin

```go
func (me Sampler) SetFilterMin(val gl.Int)
```
Sets the texture magnification function. val may be one of: GL_NEAREST,
GL_LINEAR, gl.NEAREST_MIPMAP_NEAREST, gl.LINEAR_MIPMAP_NEAREST,
gl.NEAREST_MIPMAP_LINEAR, gl.LINEAR_MIPMAP_LINEAR.

#### func (Sampler) SetMipLodBias

```go
func (me Sampler) SetMipLodBias(val gl.Float)
```
This bias will be added to the mipmap LOD calculation (as well as added to the
bias specified in one of the texture accessing functions in GLSL), which is used
to select the image.

#### func (Sampler) SetMipLodMax

```go
func (me Sampler) SetMipLodMax(val gl.Float)
```
Sets the maximum level-of-detail parameter. This limits the selection of the
lowest resolution mipmap (highest mipmap level).

#### func (Sampler) SetMipLodMin

```go
func (me Sampler) SetMipLodMin(val gl.Float)
```
Sets the minimum level-of-detail parameter. This limits the selection of highest
resolution mipmap (lowest mipmap level).

#### func (Sampler) SetParamFloat

```go
func (me Sampler) SetParamFloat(param gl.Enum, val gl.Float)
```
A convenience short-hand to invoke gl.SamplerParameterf for this sampler object.

#### func (Sampler) SetParamInt

```go
func (me Sampler) SetParamInt(param gl.Enum, val gl.Int)
```
A convenience short-hand to invoke gl.SamplerParameteri for this sampler object.

#### func (*Sampler) SetWrap

```go
func (me *Sampler) SetWrap(wrapVal gl.Int) *Sampler
```
A convenience short-hand that calls SetWrapR(), SetWrapS() and SetWrapT() with
the specified wrapVal.

#### func (Sampler) SetWrapClampBorderColor

```go
func (me Sampler) SetWrapClampBorderColor(val *GlVec4)
```
Specifies four values that define the border values that should be used for
border texels.

#### func (Sampler) SetWrapR

```go
func (me Sampler) SetWrapR(val gl.Int)
```
Sets the wrap parameter for texture coordinate r to gl.CLAMP_TO_EDGE,
gl.CLAMP_TO_BORDER, gl.MIRRORED_REPEAT, or gl.REPEAT.

#### func (Sampler) SetWrapS

```go
func (me Sampler) SetWrapS(val gl.Int)
```
Sets the wrap parameter for texture coordinate s to gl.CLAMP_TO_EDGE,
gl.CLAMP_TO_BORDER, gl.MIRRORED_REPEAT, or gl.REPEAT.

#### func (Sampler) SetWrapT

```go
func (me Sampler) SetWrapT(val gl.Int)
```
Sets the wrap parameter for texture coordinate t to gl.CLAMP_TO_EDGE,
gl.CLAMP_TO_BORDER, gl.MIRRORED_REPEAT, or gl.REPEAT.

#### func (Sampler) Unbind

```go
func (_ Sampler) Unbind(textureUnit gl.Uint)
```
Unbinds whatever sampler object is currently bound to the specified textureUnit.

#### func (Sampler) WrapClampBorderColor

```go
func (me Sampler) WrapClampBorderColor() (val *GlVec4)
```
Returns four floating-point numbers that comprise the RGBA color of the texture
border. The initial value is (0, 0, 0, 0).

#### func (Sampler) WrapR

```go
func (me Sampler) WrapR() gl.Int
```
Returns the single-valued wrapping function for texture coordinate r, a symbolic
constant. The initial value is gl.REPEAT.

#### func (Sampler) WrapS

```go
func (me Sampler) WrapS() gl.Int
```
Returns the single-valued wrapping function for texture coordinate s, a symbolic
constant. The initial value is gl.REPEAT.

#### func (Sampler) WrapT

```go
func (me Sampler) WrapT() gl.Int
```
Returns the single-valued wrapping function for texture coordinate t, a symbolic
constant. The initial value is gl.REPEAT.

#### type Shader

```go
type Shader struct {
	//	The OpenGL handle to the underlying shader object.
	//	This is 0 before calling Create() and after calling Dispose().
	GlHandle gl.Uint

	//	The shader stage provided by this shader,
	//	such as for example gl.FRAGMENT_SHADER, gl.VERTEX_SHADER, etc.
	Stage gl.Enum
}
```

Represents an OpenGL shader object.

#### func  NewComputeShader

```go
func NewComputeShader() (me *Shader)
```
If Support.Glsl.Shaders.ComputeStage is true, initializes and returns --but does
not Create()-- a new Shader for the gl.COMPUTE_SHADER Stage with the specified
name.

#### func  NewFragmentShader

```go
func NewFragmentShader() *Shader
```
Initializes and returns --but does not Create()-- a new Shader for the
gl.FRAGMENT_SHADER Stage with the specified name.

#### func  NewGeometryShader

```go
func NewGeometryShader() *Shader
```
Initializes and returns --but does not Create()-- a new Shader for the
gl.GEOMETRY_SHADER Stage with the specified name.

#### func  NewShader

```go
func NewShader(glStage gl.Enum) (me *Shader)
```
Initializes and returns --but does not Create()-- a new Shader with the
specified name and shader stage.

#### func  NewTessCtlShader

```go
func NewTessCtlShader() (me *Shader)
```
If Support.Glsl.Shaders.TessellationStages is true, initializes and returns
--but does not Create()-- a new Shader for the gl.TESS_CONTROL_SHADER ("Hull")
Stage with the specified name.

#### func  NewTessEvalShader

```go
func NewTessEvalShader() (me *Shader)
```
If Support.Glsl.Shaders.TessellationStages is true, initializes and returns
--but does not Create()-- a new Shader for the gl.TESS_EVALUATION_SHADER
("Domain") Stage with the specified name.

#### func  NewVertexShader

```go
func NewVertexShader() *Shader
```
Initializes and returns --but does not Create()-- a new Shader for the
gl.VERTEX_SHADER Stage with the specified name.

#### func (*Shader) AttachTo

```go
func (me *Shader) AttachTo(prog *Program) error
```
Attaches this shader object to the specified program object.

#### func (*Shader) Compile

```go
func (me *Shader) Compile(name string) (err error)
```
Compiles this shader object. This is a convenience short-hand for calling
gl.CompileShader(), checking gl.COMPILE_STATUS and obtaining
gl.GetShaderInfoLog().

#### func (*Shader) Create

```go
func (me *Shader) Create() (err error)
```
Creates this shader object in OpenGL.

#### func (*Shader) DetachFrom

```go
func (me *Shader) DetachFrom(prog *Program)
```
Detaches this shader object from the specified program object.

#### func (*Shader) Dispose

```go
func (me *Shader) Dispose()
```
Deletes this shader object from OpenGL.

#### func (*Shader) InfoLog

```go
func (me *Shader) InfoLog(name string) string
```
Retrieves and returns the content of the current OpenGL info-log for this shader
object.

#### func (*Shader) ParamInt

```go
func (me *Shader) ParamInt(pname gl.Enum) (iv gl.Int)
```
Convenience short-hand for gl.GetShaderiv.

#### func (*Shader) SetSource

```go
func (me *Shader) SetSource(source string, defines map[string]interface{}) (finalRealSrc string, err error)
```
Replaces the GLSL source code in this shader object with source, prepending a
#version (core) directive with the current Support.Glsl.Version.Num and #define
directives for all specified name-value pairs in defines.

#### func (*Shader) StageName

```go
func (me *Shader) StageName() (sn string)
```
Returns the name of this shader's stage, for example "GL_VERTEX_SHADER",
"GL_COMPUTE_SHADER", etc.

#### type ShaderSources

```go
type ShaderSources struct {
	Compute  string
	Fragment string
	Geometry string
	TessCtl  string
	TessEval string
	Vertex   string
}
```


#### type Texture

```go
type Texture interface {
	Dispose()

	//	Deletes and (re)creates the texture object based on its current params.
	Recreate() error
}
```

Implemented by specialized texture types such as Texture2D.

#### type Texture2D

```go
type Texture2D struct {
	//	Common texture params
	TextureBase

	//	Width and height of mip-level 0 for this 2-dimensional texture object.
	Width, Height gl.Sizei
}
```

Represents a 2-dimensional texture image.

#### func  NewTexture2D

```go
func NewTexture2D() (me *Texture2D)
```
Initializes --but does not Recreate()-- a new Texture2D, calls its Init() method
and returns it.

#### func (*Texture2D) Init

```go
func (me *Texture2D) Init()
```
Sets GlTarget to gl.TEXTURE_2D and initializes TextureBase with defaults.

#### func (*Texture2D) MaxNumMipLevels

```go
func (me *Texture2D) MaxNumMipLevels() gl.Sizei
```
Returns the maximum number of possible MIP map levels for this 2-dimensional
texture image according to its current Width and Height.

#### func (*Texture2D) PrepFromImage

```go
func (me *Texture2D) PrepFromImage(bgra, uintRev bool, img image.Image) (err error)
```
Prepares this Texture2D for uploading the specified Image via Recreate(). This
sets all of the following fields to applicable values: me.PixelData.Type,
me.PixelData.Format, me.PixelData.Ptrs[0], me.Width, me.Height,
me.MipMap.NumLevels, me.SizedInternalFormat

#### func (*Texture2D) Recreate

```go
func (me *Texture2D) Recreate() error
```
Deletes and (re)creates the texture object based on its current params. Uploads
the image data at me.PixelData.Ptrs[0], if any. Generates the MIP map if
me.MipMap.AutoGen is true and me.MipMap.NumLevels isn't 1. If
me.MipMap.NumLevels is 0 (or just smaller than 1), then me.MaxNumMipLevels() is
used.

#### func (*Texture2D) SubImage

```go
func (me *Texture2D) SubImage(x, y gl.Int, width, height gl.Sizei, ptr gl.Ptr) error
```
Updates the specified portion of this texture image with the specified pixel
data.

#### type TextureBase

```go
type TextureBase struct {
	//	The OpenGL handle for this texture object.
	//	This is 0 before calling Recreate() and after calling Dispose().
	GlHandle gl.Uint

	//	The type of texture, such as for example gl.TEXTURE_2D.
	GlTarget gl.Enum

	//	Specifies the sized internal format to be used to store texture image data,
	//	as per gl.TexStorageNN(). Defaults to gl.RGBA8.
	SizedInternalFormat gl.Enum

	//	If true, then (if Support.Textures.Immutable is also true) whenever this Texture
	//	is Recreate()d it is declared immutable in OpenGL, meaning its dimensions are
	//	locked at creation time and cannot be changed subsequently.
	//	Defaults to Support.Textures.Immutable.
	Immutable bool

	//	Settings for this texture's MIP map, if any.
	MipMap struct {
		//	If true (the default) and NumLevels isn't 1, all MIP map
		//	levels are automatically generated by Recreate()
		AutoGen bool

		//	The maximum number of MIP map levels for this texture object.
		//	Set to 0 (the default) to have Recreate() determine this
		//	automatically, set to 1 for a texture object with no MIP map.
		NumLevels gl.Sizei
	}

	//	Information regarding the pixel data stored by this texture object.
	PixelData struct {
		//	Specifies the format of the pixel data, as per
		//	gl.TexSubImageNN(). Defaults to gl.RGBA.
		Format gl.Enum

		//	Specifies the data type of the pixel data, as per
		//	gl.TexSubImageNN(). Defaults to gl.UNSIGNED_BYTE.
		Type gl.Enum

		//	Pointers (one per sub-image) to the first pixel
		//	of the data stream to be uploaded by Recreate().
		//	Initially defaults to []gl.Ptr { PtrNil }
		Ptrs [6]gl.Ptr
	}
}
```

Embedded by specialized texture types such as Texture2D.

#### func (*TextureBase) Dispose

```go
func (me *TextureBase) Dispose()
```
Deletes this texture object from OpenGL.

#### func (*TextureBase) ParamBaseLevel

```go
func (me *TextureBase) ParamBaseLevel() gl.Int
```
Returns the index of the lowest defined mipmap level. Defaults to 0.

#### func (*TextureBase) ParamFloat

```go
func (me *TextureBase) ParamFloat(param gl.Enum) (val gl.Float)
```
A convenience short-hand to invoke gl.GetTexParameterfv for the texture object
currently bound to me.GlTarget.

#### func (*TextureBase) ParamInt

```go
func (me *TextureBase) ParamInt(param gl.Enum) (val gl.Int)
```
A convenience short-hand to invoke gl.GetTexParameteriv for the texture object
currently bound to me.GlTarget.

#### func (*TextureBase) ParamMaxLevel

```go
func (me *TextureBase) ParamMaxLevel() gl.Int
```
Returns the index of the highest defined mipmap level. Defaults to 1000.

#### func (*TextureBase) ParamSwizzleA

```go
func (me *TextureBase) ParamSwizzleA() gl.Int
```
Returns the swizzle that will be applied to the A component of a texel before it
is returned to the shader. Defaults to gl.ALPHA.

#### func (*TextureBase) ParamSwizzleB

```go
func (me *TextureBase) ParamSwizzleB() gl.Int
```
Returns the swizzle that will be applied to the B component of a texel before it
is returned to the shader. Defaults to gl.BLUE.

#### func (*TextureBase) ParamSwizzleG

```go
func (me *TextureBase) ParamSwizzleG() gl.Int
```
Returns the swizzle that will be applied to the G component of a texel before it
is returned to the shader. Defaults to gl.GREEN.

#### func (*TextureBase) ParamSwizzleR

```go
func (me *TextureBase) ParamSwizzleR() gl.Int
```
Returns the swizzle that will be applied to the R component of a texel before it
is returned to the shader. Defaults to gl.RED.

#### func (*TextureBase) ParamSwizzleRgba

```go
func (me *TextureBase) ParamSwizzleRgba() gl.Int
```
Returns the swizzle that will be applied to the R, G, B and A components of a
texel before it is returned to the shader.

#### func (*TextureBase) SetParamBaseLevel

```go
func (me *TextureBase) SetParamBaseLevel(val gl.Int)
```
Sets the index of the lowest defined mipmap level.

#### func (*TextureBase) SetParamFloat

```go
func (me *TextureBase) SetParamFloat(param gl.Enum, val gl.Float)
```
A convenience short-hand to invoke gl.TexParameterf for the texture object
currently bound to me.GlTarget.

#### func (*TextureBase) SetParamInt

```go
func (me *TextureBase) SetParamInt(param gl.Enum, val gl.Int)
```
A convenience short-hand to invoke gl.TexParameteri for the texture object
currently bound to me.GlTarget.

#### func (*TextureBase) SetParamMaxLevel

```go
func (me *TextureBase) SetParamMaxLevel(val gl.Int)
```
Sets the index of the highest defined mipmap level.

#### func (*TextureBase) SetParamSwizzleA

```go
func (me *TextureBase) SetParamSwizzleA(val gl.Int)
```
Sets the swizzle that will be applied to the A component of a texel before it is
returned to the shader.

#### func (*TextureBase) SetParamSwizzleB

```go
func (me *TextureBase) SetParamSwizzleB(val gl.Int)
```
Sets the swizzle that will be applied to the B component of a texel before it is
returned to the shader.

#### func (*TextureBase) SetParamSwizzleG

```go
func (me *TextureBase) SetParamSwizzleG(val gl.Int)
```
Sets the swizzle that will be applied to the G component of a texel before it is
returned to the shader.

#### func (*TextureBase) SetParamSwizzleR

```go
func (me *TextureBase) SetParamSwizzleR(val gl.Int)
```
Sets the swizzle that will be applied to the R component of a texel before it is
returned to the shader.

#### func (*TextureBase) SetParamSwizzleRgba

```go
func (me *TextureBase) SetParamSwizzleRgba(val gl.Int)
```
Sets the swizzle that will be applied to the R, G, B and A components of a texel
before it is returned to the shader.

#### type TextureCube

```go
type TextureCube struct {
	//	Common texture params
	TextureBase

	//	Width and height of mip-level 0 for all 6 cube-faces.
	Width, Height gl.Sizei
}
```

Represents six 2-dimensional texture images of equal dimensions.

#### func  NewTextureCube

```go
func NewTextureCube() (me *TextureCube)
```
Initializes --but does not Recreate()-- a new TextureCube, calls its Init()
method and returns it.

#### func (*TextureCube) Init

```go
func (me *TextureCube) Init()
```
Sets GlTarget to gl.TEXTURE_CUBE_MAP and initializes TextureBase with defaults.

#### func (*TextureCube) MaxNumMipLevels

```go
func (me *TextureCube) MaxNumMipLevels() gl.Sizei
```
Returns the maximum number of possible MIP map levels for any of this cube-map's
6 faces according to its current Width and Height.

#### func (*TextureCube) PrepFromImages

```go
func (me *TextureCube) PrepFromImages(bgra, uintRev bool, images ...image.Image) (err error)
```
Prepares this TextureCube for uploading the specified Images via Recreate().
This sets all of the following fields to applicable values: me.PixelData.Type,
me.PixelData.Format, me.PixelData.Ptrs, me.Width, me.Height,
me.MipMap.NumLevels, me.SizedInternalFormat

#### func (*TextureCube) Recreate

```go
func (me *TextureCube) Recreate() error
```
Deletes and (re)creates the texture object based on its current params. Uploads
the image data pointed to in me.PixelData.Ptrs, if any. Generates the MIP map if
me.MipMap.AutoGen is true and me.MipMap.NumLevels isn't 1. If
me.MipMap.NumLevels is 0 (or just smaller than 1), then me.MaxNumMipLevels() is
used.

#### func (*TextureCube) SubImage

```go
func (me *TextureCube) SubImage(glTargetFace gl.Enum, x, y gl.Int, width, height gl.Sizei, ptr gl.Ptr) error
```
Updates the specified portion of the specified cube-face image with the
specified pixel data.

#### type TypeUtils

```go
type TypeUtils struct {
}
```

A singleton type, only used for the package-global ugl.Typed variable.

#### func (TypeUtils) Clamp

```go
func (_ TypeUtils) Clamp(val, min, max gl.Float) gl.Float
```
Clamps the specified val between min and max.

#### func (TypeUtils) Ifb

```go
func (_ TypeUtils) Ifb(one, two bool, ifOne, ifTwo, ifBoth gl.Enum) (val gl.Enum)
```
If one but not two, return ifOne. If two but not one, return ifTwo. Else, return
ifBoth.

#### func (TypeUtils) Ife

```go
func (_ TypeUtils) Ife(cond bool, ifTrue, ifFalse gl.Enum) gl.Enum
```
Returns ifTrue if cond is true, otherwise returns ifFalse.

#### func (TypeUtils) Ifi

```go
func (_ TypeUtils) Ifi(cond bool, ifTrue, ifFalse gl.Int) gl.Int
```
Returns ifTrue if cond is true, otherwise returns ifFalse.

#### type UniformBlock

```go
type UniformBlock struct {
	Buffer Buffer
	Name   string
}
```


#### func  NewUniformBlock

```go
func NewUniformBlock(name string) (me *UniformBlock)
```

#### func (*UniformBlock) Create

```go
func (me *UniformBlock) Create()
```

#### func (*UniformBlock) Dispose

```go
func (me *UniformBlock) Dispose()
```

#### type VertexArray

```go
type VertexArray struct {
	//	The OpenGL handle to this vertex array object.
	//	This is 0 before calling Create() and after calling Dispose().
	GlHandle gl.Uint
}
```

Represents an OpenGL vertex array object (NOT the legacy OpenGL "vertex arrays",
but what is commonly abbreviated as VAO).

#### func (*VertexArray) Bind

```go
func (me *VertexArray) Bind()
```
Binds this vertex array object.

#### func (*VertexArray) Create

```go
func (me *VertexArray) Create() (err error)
```
Creates this vertex-array object.

#### func (*VertexArray) Dispose

```go
func (me *VertexArray) Dispose()
```
Deletes this vertex array object.

#### func (*VertexArray) Setup

```go
func (me *VertexArray) Setup(prog *Program, atts []VertexAttribPointer, bufs ...*Buffer) (err error)
```
Sets up this vertex array object. Unless atts are specified, prog can be be nil.
To specify element/index buffer and vertex/attribute buffers (if applicable),
pass them via bufs.

#### func (VertexArray) Unbind

```go
func (_ VertexArray) Unbind()
```
Unbinds whatever vertex array object is currently bound.

#### type VertexAttribPointer

```go
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
```

Encapsulates vertex attribute information used by VertexArray.Setup() to enable
that vertex attribute.

#### func  NewVertexAttribPointer

```go
func NewVertexAttribPointer(name string, size gl.Int, stride gl.Sizei, offset gl.Ptr) (me *VertexAttribPointer)
```
Initializes and returns a new VertexAttribPointer with the specified values.

#### func (*VertexAttribPointer) Init

```go
func (me *VertexAttribPointer) Init(name string, size gl.Int, stride gl.Sizei, offset gl.Ptr)
```

--
**godocdown** http://github.com/robertkrimen/godocdown
