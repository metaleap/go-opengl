package glutil

import (
	"fmt"
	"math"

	gl "github.com/go3d/go-opengl/gogl"

	ugo "github.com/metaleap/go-util"
)

var (
	//	All the known possible OpenGL core context versions that this package supports.
	//	(This does not imply they're all supported by the current client-side GPU.)
	KnownVersions = []float64{3.3, 4.0, 4.1, 4.2, 4.3}

	//	Provides access to miscellaneous OpenGL / GoGL functionality.
	Gl Utils

	//	Provides access to miscellaneous type-specific OpenGL / GoGL functionality.
	Typed TypeUtils

	//	After you have called glutil.Init() upon successful OpenGL context creation and GoGL initialization,
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
		//	Useful for casual inspection, but to query for the presence of a specific extension, always use Gl.Extension(name)
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
		}
	}

	//	Contains the well-known extension prefixes that can be omitted when querying the current OpenGL context via Extension()
	//	NOTE: if you find a need to append a prefix not currently included here, please post it on GitHub Issues to have it included by default.
	KnownExtensionPrefixes = []string{"GL_3DFX_", "GL_3DL_", "GL_AMD_", "GL_APPLE_", "GL_ARB_", "GL_ATI_", "GL_EXT_", "GL_GREMEDY_", "GL_HP_", "GL_I3D_", "GL_IBM_", "GL_INGR_", "GL_INTEL_", "GL_KHR_", "GL_KTX_", "GL_MESA_", "GL_MESAX_", "GL_NV_", "GL_NVX_", "GL_OES_", "GL_OML_", "GL_PGI_", "GL_REND__", "GL_S3_", "GL_SGI_", "GL_SGIS_", "GL_SGIX_", "GL_SUN_", "GL_SUNX_", "GL_WIN_", "WGL_EXT_"}
)

//	DO call this immediately upon successful OpenGL context initialization to prepare this package for use.
//	Fully populates all fields and sub-fields in the Support package global.
func Init() {
	setVersion()
	setSupportInfos()
}

//	If any OpenGL errors have accumulated since LastError() was last called, returns a Go error for the first in that queue of accumulated errors.
//	Unless nil, the returned err includes the specified formatted message plus a string representation of the GL error category, for example: "GL_OUT_OF_MEMORY".
func LastError(stepFmt string, stepFmtArgs ...interface{}) (err error) {
	if errEnum := gl.GetError(); errEnum != 0 {
		if len(stepFmtArgs) > 0 {
			stepFmt = sfmt(stepFmt, stepFmtArgs...)
		}
		errStr := "GL ERR: "
		switch errEnum {
		case gl.INVALID_ENUM:
			errStr += "GL_INVALID_ENUM"
		case gl.INVALID_VALUE:
			errStr += "GL_INVALID_VALUE"
		case gl.INVALID_OPERATION:
			errStr += "GL_INVALID_OPERATION"
		case gl.OUT_OF_MEMORY:
			errStr += "GL_OUT_OF_MEMORY"
		case gl.INVALID_FRAMEBUFFER_OPERATION:
			errStr += "GL_INVALID_FRAMEBUFFER_OPERATION"
		case gl.STACK_OVERFLOW:
			errStr += "GL_STACK_OVERFLOW"
		case gl.STACK_UNDERFLOW:
			errStr += "GL_STACK_UNDERFLOW"
		default:
			errStr += sfmt("GOGL_ERROR_%v", errEnum)
		}
		err = fmt.Errorf("%s\t%s", errStr, stepFmt)
	}
	return
}

//	log.Println()s the error returned by LastError(), if any.
func LogLastError(stepFmt string, stepFmtArgs ...interface{}) {
	ugo.LogError(LastError(stepFmt, stepFmtArgs...))
}

func setSupportInfos() {
	//	Extensions
	var num gl.Int
	gl.GetIntegerv(gl.NUM_EXTENSIONS, &num)
	Support.Extensions = make([]string, num)
	if num > 0 {
		for i := gl.Int(0); i < num; i++ {
			Support.Extensions[i] = Gl.Stri(gl.EXTENSIONS, gl.Uint(i))
		}
	}

	//	Sampling limits
	if Gl.Extension("texture_filter_anisotropic") {
		gl.GetFloatv(max_texture_max_anisotropy_ext, &Support.Textures.MaxFilterAnisotropy)
	}
	gl.GetFloatv(gl.MAX_TEXTURE_LOD_BIAS, &Support.Textures.MaxMipLoadBias)

	//	Version-specifics
	if Support.GlVersion.Is42 {
		Support.Buffers.AtomicCounter, Support.Textures.Immutable = true, true
	}
	if Support.GlVersion.Is43 {
		Support.Buffers.DispatchIndirect, Support.Buffers.ShaderStorage = true, true
	}
	if Support.Glsl.Version.Is400 {
		Support.Glsl.Shaders.TessellationStages = true
	}
	if Support.Glsl.Version.Is430 {
		Support.Glsl.Shaders.ComputeStage = true
	}
}

func setVersion() {
	glVer, glslVer := &Support.GlVersion, &Support.Glsl.Version
	glVer.Num, glslVer.Num = 0.0, 0
	glVer.Is33, glVer.Is40, glVer.Is41, glVer.Is42, glVer.Is43 = false, false, false, false, false
	glslVer.Is330, glslVer.Is400, glslVer.Is410, glslVer.Is420, glslVer.Is430 = false, false, false, false, false
	if glVer.MajorMinor, glVer.Num = ugo.ParseVersion(Gl.Str(gl.VERSION)); glVer.Num > 0 {
		if glVer.Num >= 3.3 {
			glVer.Is33, glslVer.Num, glslVer.Is330 = true, 330, true
		}
		if glVer.Num >= 4.0 {
			glVer.Is40, glslVer.Num, glslVer.Is400 = true, 400, true
		}
		if glVer.Num >= 4.1 {
			glVer.Is41, glslVer.Num, glslVer.Is410 = true, 410, true
		}
		if glVer.Num >= 4.2 {
			glVer.Is42, glslVer.Num, glslVer.Is420 = true, 420, true
		}
		if glVer.Num >= 4.3 {
			glVer.Is43, glslVer.Num, glslVer.Is430 = true, 430, true
		}
	}
}

func sfmt(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

//	Returns the major and minor parts of the specified version number.
func VersionMajorMinor(v float64) (major, minor int) {
	major = int(v)
	//	Some sad acrobatics to get the accurate fractional part of the version number...
	minor = int(math.Ceil((v-float64(int64(v)))*100) / 10)
	return
}

//	Returns true if GlVersion is greater than or equal to v.
func VersionMatch(v float64) bool {
	return Support.GlVersion.Num >= v
}
