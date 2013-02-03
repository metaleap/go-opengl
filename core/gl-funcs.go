package gl

// #cgo darwin  LDFLAGS: -framework OpenGL
// #cgo linux   LDFLAGS: -lGL
// #cgo windows LDFLAGS: -lopengl32
// 
// #include <stdlib.h>
// #if defined(__APPLE__)
// #include <dlfcn.h>
// #elif defined(_WIN32)
// #define WIN32_LEAN_AND_MEAN 1
// #include <windows.h>
// #else
// #include <X11/Xlib.h>
// #include <GL/glx.h>
// #endif
// 
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
// #ifndef GLAPI
// #define GLAPI extern
// #endif
// 
// typedef unsigned int GLenum;
// typedef unsigned char GLboolean;
// typedef unsigned int GLbitfield;
// typedef signed char GLbyte;
// typedef short GLshort;
// typedef int GLint;
// typedef int GLsizei;
// typedef unsigned char GLubyte;
// typedef unsigned short GLushort;
// typedef unsigned int GLuint;
// typedef unsigned short GLhalf;
// typedef float GLfloat;
// typedef float GLclampf;
// typedef double GLdouble;
// typedef double GLclampd;
// typedef void GLvoid;
// 
// #include <stddef.h>
// #ifndef GL_VERSION_2_0
// /* GL type for program/shader text */
// typedef char GLchar;
// #endif
// #ifndef GL_VERSION_1_5
// /* GL types for handling large vertex buffer objects */
// typedef ptrdiff_t GLintptr;
// typedef ptrdiff_t GLsizeiptr;
// #endif
// #ifndef GL_ARB_vertex_buffer_object
// /* GL types for handling large vertex buffer objects */
// typedef ptrdiff_t GLintptrARB;
// typedef ptrdiff_t GLsizeiptrARB;
// #endif
// #ifndef GL_ARB_shader_objects
// /* GL types for program/shader text and shader object handles */
// typedef char GLcharARB;
// typedef unsigned int GLhandleARB;
// #endif
// /* GL type for "half" precision (s10e5) float data in host memory */
// #ifndef GL_ARB_half_float_pixel
// typedef unsigned short GLhalfARB;
// #endif
// #ifndef GL_NV_half_float
// typedef unsigned short GLhalfNV;
// #endif
// #ifndef GLEXT_64_TYPES_DEFINED
// /* This code block is duplicated in glxext.h, so must be protected */
// #define GLEXT_64_TYPES_DEFINED
// /* Define int32_t, int64_t, and uint64_t types for UST/MSC */
// /* (as used in the GL_EXT_timer_query extension). */
// #if defined(__STDC_VERSION__) && __STDC_VERSION__ >= 199901L
// #include <inttypes.h>
// #elif defined(__sun__) || defined(__digital__)
// #include <inttypes.h>
// #if defined(__STDC__)
// #if defined(__arch64__) || defined(_LP64)
// typedef long int int64_t;
// typedef unsigned long int uint64_t;
// #else
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #endif /* __arch64__ */
// #endif /* __STDC__ */
// #elif defined( __VMS ) || defined(__sgi)
// #include <inttypes.h>
// #elif defined(__SCO__) || defined(__USLC__)
// #include <stdint.h>
// #elif defined(__UNIXOS2__) || defined(__SOL64__)
// typedef long int int32_t;
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #elif defined(_WIN32) && defined(__GNUC__)
// #include <stdint.h>
// #elif defined(_WIN32)
// typedef __int32 int32_t;
// typedef __int64 int64_t;
// typedef unsigned __int64 uint64_t;
// #else
// /* Fallback if nothing above works */
// #include <inttypes.h>
// #endif
// #endif
// #ifndef GL_EXT_timer_query
// typedef int64_t GLint64EXT;
// typedef uint64_t GLuint64EXT;
// #endif
// #ifndef GL_ARB_sync
// typedef int64_t GLint64;
// typedef uint64_t GLuint64;
// typedef struct __GLsync *GLsync;
// #endif
// #ifndef GL_ARB_cl_event
// /* These incomplete types let us declare types compatible with OpenCL's cl_context and cl_event */
// struct _cl_context;
// struct _cl_event;
// #endif
// #ifndef GL_ARB_debug_output
// typedef void (APIENTRY *GLDEBUGPROCARB)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,GLvoid *userParam);
// #endif
// #ifndef GL_AMD_debug_output
// typedef void (APIENTRY *GLDEBUGPROCAMD)(GLuint id,GLenum category,GLenum severity,GLsizei length,const GLchar *message,GLvoid *userParam);
// #endif
// #ifndef GL_KHR_debug
// typedef void (APIENTRY *GLDEBUGPROC)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,GLvoid *userParam);
// #endif
// #ifndef GL_NV_vdpau_interop
// typedef GLintptr GLvdpauSurfaceNV;
// #endif
// 
// #ifdef _WIN32
// static HMODULE opengl32 = NULL;
// #endif
// 
// static void* coreGetProcAddress(const char* name) { 
// #ifdef __APPLE__
// 	return dlsym(RTLD_DEFAULT, name);
// #elif _WIN32
// 	void* pf = wglGetProcAddress((LPCSTR)name);
// 	if(pf) {
// 		return pf;
// 	}
// 	if(opengl32 == NULL) {
// 		opengl32 = LoadLibraryA("opengl32.dll");
// 	}
// 	return GetProcAddress(opengl32, (LPCSTR)name);
// #else
// 	return glXGetProcAddress((const GLubyte*)name);
// #endif
// }
// 
// void (APIENTRYP ptrglDrawElementsInstancedBaseInstance)(GLenum mode, GLsizei count, GLenum type, void* indices, GLsizei instancecount, GLuint baseinstance);
// void (APIENTRYP ptrglDeleteTextures)(GLsizei n, GLuint* textures);
// void (APIENTRYP ptrglGetProgramiv)(GLuint program, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetRenderbufferParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglClear)(GLbitfield mask);
// void (APIENTRYP ptrglPolygonMode)(GLenum face, GLenum mode);
// void (APIENTRYP ptrglProgramUniform4uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// GLint (APIENTRYP ptrglGetSubroutineUniformLocation)(GLuint program, GLenum shadertype, GLchar* name);
// void (APIENTRYP ptrglDrawArraysInstancedBaseInstance)(GLenum mode, GLint first, GLsizei count, GLsizei instancecount, GLuint baseinstance);
// void (APIENTRYP ptrglGetVertexAttribPointerv)(GLuint index, GLenum pname, GLvoid** pointer);
// void (APIENTRYP ptrglEnableVertexAttribArray)(GLuint index);
// GLboolean (APIENTRYP ptrglIsQuery)(GLuint id);
// void (APIENTRYP ptrglProgramUniformMatrix4x2fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglStencilFuncSeparate)(GLenum face, GLenum func, GLint ref, GLuint mask);
// void (APIENTRYP ptrglPixelStorei)(GLenum pname, GLint param);
// void (APIENTRYP ptrglTexParameterIiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetActiveAtomicCounterBufferiv)(GLuint program, GLuint bufferIndex, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetActiveSubroutineUniformiv)(GLuint program, GLenum shadertype, GLuint index, GLenum pname, GLint* values);
// void (APIENTRYP ptrglColorMaski)(GLuint index, GLboolean r, GLboolean g, GLboolean b, GLboolean a);
// void (APIENTRYP ptrglViewportIndexedfv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglProgramUniform3ui)(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2);
// void (APIENTRYP ptrglMultiTexCoordP2ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglProgramUniform4f)(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3);
// void (APIENTRYP ptrglColorP3ui)(GLenum type, GLuint color);
// void (APIENTRYP ptrglBlendEquationSeparate)(GLenum modeRGB, GLenum modeAlpha);
// void (APIENTRYP ptrglProgramUniform4i)(GLuint program, GLint location, GLint v0, GLint v1, GLint v2, GLint v3);
// void (APIENTRYP ptrglFlush)();
// void (APIENTRYP ptrglFinish)();
// void (APIENTRYP ptrglUniformMatrix3dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglEndTransformFeedback)();
// void (APIENTRYP ptrglProgramUniform3d)(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2);
// void (APIENTRYP ptrglGetIntegerv)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglUniformMatrix4x3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglGetTransformFeedbackVarying)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLsizei* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrglProgramUniform4d)(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3);
// void (APIENTRYP ptrglGetProgramResourceName)(GLuint program, GLenum programInterface, GLuint index, GLsizei bufSize, GLsizei* length, GLchar* name);
// void (APIENTRYP ptrglDeleteProgramPipelines)(GLsizei n, GLuint* pipelines);
// void (APIENTRYP ptrglDisable)(GLenum cap);
// void (APIENTRYP ptrglVertexAttribLPointer)(GLuint index, GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglVertexAttribP3ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglProvokingVertex)(GLenum mode);
// void (APIENTRYP ptrglDrawArraysInstanced)(GLenum mode, GLint first, GLsizei count, GLsizei instancecount);
// void (APIENTRYP ptrglNormalP3uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglUniform2iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglProgramUniform1d)(GLuint program, GLint location, GLdouble v0);
// void (APIENTRYP ptrglPolygonOffset)(GLfloat factor, GLfloat units);
// void (APIENTRYP ptrglUniformMatrix3x4dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglTexStorage3D)(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth);
// void (APIENTRYP ptrglGetTexParameterIiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglProgramUniformMatrix4fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglVertexAttribP2uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrglCopyTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrglGetSynciv)(GLsync sync, GLenum pname, GLsizei bufSize, GLsizei* length, GLint* values);
// void (APIENTRYP ptrglUniform4f)(GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3);
// void (APIENTRYP ptrglProgramUniformMatrix2fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglDeleteShader)(GLuint shader);
// void (APIENTRYP ptrglBindBuffer)(GLenum target, GLuint buffer);
// void (APIENTRYP ptrglEnable)(GLenum cap);
// void (APIENTRYP ptrglWaitSync)(GLsync sync, GLbitfield flags, GLuint64 timeout);
// void (APIENTRYP ptrglBindAttribLocation)(GLuint program, GLuint index, GLchar* name);
// void (APIENTRYP ptrglVertexAttribI4iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglGetProgramStageiv)(GLuint program, GLenum shadertype, GLenum pname, GLint* values);
// void (APIENTRYP ptrglTexCoordP2uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglMultiTexCoordP4ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglUniform3iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglGetIntegeri_v)(GLenum target, GLuint index, GLint* data);
// void (APIENTRYP ptrglDrawElementsInstanced)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei instancecount);
// void (APIENTRYP ptrglMinSampleShading)(GLfloat value);
// void (APIENTRYP ptrglGetShaderSource)(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source);
// void (APIENTRYP ptrglProgramUniform2uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglStencilMaskSeparate)(GLenum face, GLuint mask);
// void (APIENTRYP ptrglScissor)(GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglGetShaderPrecisionFormat)(GLenum shadertype, GLenum precisiontype, GLint* range, GLint* precision);
// void (APIENTRYP ptrglVertexAttribI4ui)(GLuint index, GLuint x, GLuint y, GLuint z, GLuint w);
// void (APIENTRYP ptrglColorP3uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrglGetActiveUniformBlockName)(GLuint program, GLuint uniformBlockIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformBlockName);
// void (APIENTRYP ptrglTexStorage3DMultisample)(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrglVertexAttribI3uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexP2uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrglMultiDrawElementsIndirect)(GLenum mode, GLenum type, void* indirect, GLsizei drawcount, GLsizei stride);
// void (APIENTRYP ptrglUniform4uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglGetTexParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglVertexP2ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrglGenVertexArrays)(GLsizei n, GLuint* arrays);
// GLint (APIENTRYP ptrglGetProgramResourceLocationIndex)(GLuint program, GLenum programInterface, GLchar* name);
// void (APIENTRYP ptrglGetSamplerParameteriv)(GLuint sampler, GLenum pname, GLint* params);
// void (APIENTRYP ptrglDepthRangef)(GLfloat n, GLfloat f);
// void (APIENTRYP ptrglVertexAttribI1ui)(GLuint index, GLuint x);
// void (APIENTRYP ptrglClearColor)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglTexParameterIuiv)(GLenum target, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglProgramUniform2fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglVertexAttribBinding)(GLuint attribindex, GLuint bindingindex);
// void (APIENTRYP ptrglTexCoordP4uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglPatchParameterfv)(GLenum pname, GLfloat* values);
// void (APIENTRYP ptrglCompressedTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglActiveTexture)(GLenum texture);
// void (APIENTRYP ptrglGetProgramBinary)(GLuint program, GLsizei bufSize, GLsizei* length, GLenum* binaryFormat, GLvoid* binary);
// void (APIENTRYP ptrglUniform1ui)(GLint location, GLuint v0);
// void (APIENTRYP ptrglCopyTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border);
// void (APIENTRYP ptrglDepthRangeIndexed)(GLuint index, GLdouble n, GLdouble f);
// void (APIENTRYP ptrglCopyTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border);
// void (APIENTRYP ptrglScissorArrayv)(GLuint first, GLsizei count, GLint* v);
// void (APIENTRYP ptrglGetShaderInfoLog)(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrglGetMultisamplefv)(GLenum pname, GLuint index, GLfloat* val);
// void (APIENTRYP ptrglDrawArrays)(GLenum mode, GLint first, GLsizei count);
// void (APIENTRYP ptrglGetActiveUniformBlockiv)(GLuint program, GLuint uniformBlockIndex, GLenum pname, GLint* params);
// void (APIENTRYP ptrglUniformMatrix3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// GLuint (APIENTRYP ptrglCreateProgram)();
// void (APIENTRYP ptrglVertexAttribI1uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglGetQueryObjectuiv)(GLuint id, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglEndQuery)(GLenum target);
// void (APIENTRYP ptrglClearDepth)(GLdouble depth);
// void (APIENTRYP ptrglGetFloati_v)(GLenum target, GLuint index, GLfloat* data);
// void (APIENTRYP ptrglPointParameterfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetQueryObjectiv)(GLuint id, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetActiveUniformName)(GLuint program, GLuint uniformIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformName);
// void (APIENTRYP ptrglNormalP3ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglActiveShaderProgram)(GLuint pipeline, GLuint program);
// void (APIENTRYP ptrglVertexAttribI3iv)(GLuint index, GLint* v);
// GLboolean (APIENTRYP ptrglIsFramebuffer)(GLuint framebuffer);
// void (APIENTRYP ptrglBindImageTexture)(GLuint unit, GLuint texture, GLint level, GLboolean layered, GLint layer, GLenum access, GLenum format);
// void (APIENTRYP ptrglGetActiveUniform)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrglUniform1i)(GLint location, GLint v0);
// void (APIENTRYP ptrglObjectPtrLabel)(void* ptr, GLsizei length, GLchar* label);
// void (APIENTRYP ptrglUseProgramStages)(GLuint pipeline, GLbitfield stages, GLuint program);
// void (APIENTRYP ptrglGetTexLevelParameteriv)(GLenum target, GLint level, GLenum pname, GLint* params);
// void (APIENTRYP ptrglVertexAttribLFormat)(GLuint attribindex, GLint size, GLenum type, GLuint relativeoffset);
// void (APIENTRYP ptrglGetSamplerParameterIiv)(GLuint sampler, GLenum pname, GLint* params);
// void (APIENTRYP ptrglBufferSubData)(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data);
// void (APIENTRYP ptrglTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglUniform3ui)(GLint location, GLuint v0, GLuint v1, GLuint v2);
// void (APIENTRYP ptrglTexImage2DMultisample)(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrglRenderbufferStorageMultisample)(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
// GLboolean (APIENTRYP ptrglIsProgram)(GLuint program);
// void (APIENTRYP ptrglVertexAttribP2ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglVertexAttribI4bv)(GLuint index, GLbyte* v);
// void (APIENTRYP ptrglUniform2i)(GLint location, GLint v0, GLint v1);
// void (APIENTRYP ptrglDeleteRenderbuffers)(GLsizei n, GLuint* renderbuffers);
// void (APIENTRYP ptrglGetBufferPointerv)(GLenum target, GLenum pname, GLvoid** params);
// void (APIENTRYP ptrglDrawBuffer)(GLenum mode);
// GLvoid* (APIENTRYP ptrglMapBuffer)(GLenum target, GLenum access);
// void (APIENTRYP ptrglDepthRangeArrayv)(GLuint first, GLsizei count, GLdouble* v);
// void (APIENTRYP ptrglVertexP3uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrglFramebufferTexture3D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint zoffset);
// void (APIENTRYP ptrglVertexAttribI4ubv)(GLuint index, GLubyte* v);
// void (APIENTRYP ptrglUniform4i)(GLint location, GLint v0, GLint v1, GLint v2, GLint v3);
// void (APIENTRYP ptrglCompressedTexImage3D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglDeleteTransformFeedbacks)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglBeginQuery)(GLenum target, GLuint id);
// void (APIENTRYP ptrglReadBuffer)(GLenum mode);
// void (APIENTRYP ptrglDebugMessageControl)(GLenum source, GLenum type, GLenum severity, GLsizei count, GLuint* ids, GLboolean enabled);
// GLuint (APIENTRYP ptrglGetSubroutineIndex)(GLuint program, GLenum shadertype, GLchar* name);
// void (APIENTRYP ptrglVertexAttribI2uiv)(GLuint index, GLuint* v);
// GLvoid* (APIENTRYP ptrglMapBufferRange)(GLenum target, GLintptr offset, GLsizeiptr length, GLbitfield access);
// void (APIENTRYP ptrglDrawBuffers)(GLsizei n, GLenum* bufs);
// GLboolean (APIENTRYP ptrglUnmapBuffer)(GLenum target);
// void (APIENTRYP ptrglUniform3d)(GLint location, GLdouble x, GLdouble y, GLdouble z);
// GLint (APIENTRYP ptrglGetAttribLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglUseProgram)(GLuint program);
// void (APIENTRYP ptrglTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglUniformMatrix4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix3x4dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglBlendFuncSeparatei)(GLuint buf, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha);
// void (APIENTRYP ptrglVertexAttribP3uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// GLboolean (APIENTRYP ptrglIsRenderbuffer)(GLuint renderbuffer);
// GLboolean (APIENTRYP ptrglIsProgramPipeline)(GLuint pipeline);
// void (APIENTRYP ptrglScissorIndexed)(GLuint index, GLint left, GLint bottom, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglBlendColor)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglTexParameterf)(GLenum target, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglInvalidateTexSubImage)(GLuint texture, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth);
// GLint (APIENTRYP ptrglGetFragDataLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglUniformMatrix2x3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglUniform4d)(GLint location, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglClearBufferfv)(GLenum buffer, GLint drawbuffer, GLfloat* value);
// void (APIENTRYP ptrglGetQueryIndexediv)(GLenum target, GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrglProgramUniform3fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglProgramUniform1iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglUniform2fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglGetProgramInterfaceiv)(GLuint program, GLenum programInterface, GLenum pname, GLint* params);
// void (APIENTRYP ptrglUniformMatrix2dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// GLboolean (APIENTRYP ptrglIsSync)(GLsync sync);
// void (APIENTRYP ptrglGetBooleanv)(GLenum pname, GLboolean* params);
// void (APIENTRYP ptrglStencilOp)(GLenum fail, GLenum zfail, GLenum zpass);
// void (APIENTRYP ptrglGetQueryObjectui64v)(GLuint id, GLenum pname, GLuint64* params);
// void (APIENTRYP ptrglProgramUniformMatrix4x3dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglGetObjectPtrLabel)(void* ptr, GLsizei bufSize, GLsizei* length, GLchar* label);
// void (APIENTRYP ptrglClearBufferfi)(GLenum buffer, GLint drawbuffer, GLfloat depth, GLint stencil);
// void (APIENTRYP ptrglProgramUniformMatrix2x4fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglUniform3fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglVertexP4ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrglGetQueryObjecti64v)(GLuint id, GLenum pname, GLint64* params);
// void (APIENTRYP ptrglCompressedTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglPrimitiveRestartIndex)(GLuint index);
// void (APIENTRYP ptrglDrawTransformFeedback)(GLenum mode, GLuint id);
// void (APIENTRYP ptrglGetDoublev)(GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglBlendFuncSeparate)(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha);
// void (APIENTRYP ptrglVertexAttribP1ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglVertexAttribP4uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrglBeginTransformFeedback)(GLenum primitiveMode);
// void (APIENTRYP ptrglClearStencil)(GLint s);
// void (APIENTRYP ptrglDeleteFramebuffers)(GLsizei n, GLuint* framebuffers);
// void (APIENTRYP ptrglVertexAttribL1dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglPointParameterf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglVertexAttribI4sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglVertexAttribL2dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglBlendFunci)(GLuint buf, GLenum src, GLenum dst);
// void (APIENTRYP ptrglUniform2ui)(GLint location, GLuint v0, GLuint v1);
// void (APIENTRYP ptrglGetProgramPipelineiv)(GLuint pipeline, GLenum pname, GLint* params);
// void (APIENTRYP ptrglSampleMaski)(GLuint index, GLbitfield mask);
// void (APIENTRYP ptrglGetSamplerParameterIuiv)(GLuint sampler, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglCopyBufferSubData)(GLenum readTarget, GLenum writeTarget, GLintptr readOffset, GLintptr writeOffset, GLsizeiptr size);
// void (APIENTRYP ptrglGetActiveUniformsiv)(GLuint program, GLsizei uniformCount, GLuint* uniformIndices, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetBufferParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglStencilOpSeparate)(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass);
// void (APIENTRYP ptrglBlitFramebuffer)(GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
// void (APIENTRYP ptrglDepthFunc)(GLenum func);
// void (APIENTRYP ptrglDetachShader)(GLuint program, GLuint shader);
// void (APIENTRYP ptrglClearBufferuiv)(GLenum buffer, GLint drawbuffer, GLuint* value);
// void (APIENTRYP ptrglProgramBinary)(GLuint program, GLenum binaryFormat, GLvoid* binary, GLsizei length);
// void (APIENTRYP ptrglViewportIndexedf)(GLuint index, GLfloat x, GLfloat y, GLfloat w, GLfloat h);
// void (APIENTRYP ptrglProgramUniformMatrix3dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix2x4dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramParameteri)(GLuint program, GLenum pname, GLint value);
// void (APIENTRYP ptrglMultiTexCoordP3ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglUniformMatrix4dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglBindFragDataLocation)(GLuint program, GLuint color, GLchar* name);
// void (APIENTRYP ptrglGetAttachedShaders)(GLuint program, GLsizei maxCount, GLsizei* count, GLuint* obj);
// void (APIENTRYP ptrglVertexP4uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrglReadPixels)(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglVertexAttribIFormat)(GLuint attribindex, GLint size, GLenum type, GLuint relativeoffset);
// void (APIENTRYP ptrglTexStorage1D)(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width);
// void (APIENTRYP ptrglProgramUniform4dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglVertexBindingDivisor)(GLuint bindingindex, GLuint divisor);
// void (APIENTRYP ptrglDrawRangeElementsBaseVertex)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices, GLint basevertex);
// void (APIENTRYP ptrglGetUniformiv)(GLuint program, GLint location, GLint* params);
// void (APIENTRYP ptrglVertexAttribL3dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglInvalidateBufferData)(GLuint buffer);
// void (APIENTRYP ptrglResumeTransformFeedback)();
// void (APIENTRYP ptrglUniformMatrix2x3dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglEndQueryIndexed)(GLenum target, GLuint index);
// void (APIENTRYP ptrglGetVertexAttribIuiv)(GLuint index, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglBeginQueryIndexed)(GLenum target, GLuint index, GLuint id);
// void (APIENTRYP ptrglUniformMatrix4x3dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglTexBufferRange)(GLenum target, GLenum internalformat, GLuint buffer, GLintptr offset, GLsizeiptr size);
// void (APIENTRYP ptrglGetVertexAttribLdv)(GLuint index, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglDrawTransformFeedbackStream)(GLenum mode, GLuint id, GLuint stream);
// void (APIENTRYP ptrglBindFramebuffer)(GLenum target, GLuint framebuffer);
// void (APIENTRYP ptrglDeleteProgram)(GLuint program);
// void (APIENTRYP ptrglScissorIndexedv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglUniformMatrix3x2dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// GLboolean (APIENTRYP ptrglIsTexture)(GLuint texture);
// void (APIENTRYP ptrglSamplerParameterIiv)(GLuint sampler, GLenum pname, GLint* param);
// void (APIENTRYP ptrglProgramUniform4fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglBlendEquationi)(GLuint buf, GLenum mode);
// void (APIENTRYP ptrglDeleteSync)(GLsync sync);
// void (APIENTRYP ptrglProgramUniform3iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglInvalidateTexImage)(GLuint texture, GLint level);
// void (APIENTRYP ptrglViewportArrayv)(GLuint first, GLsizei count, GLfloat* v);
// void (APIENTRYP ptrglTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels);
// GLuint (APIENTRYP ptrglCreateShaderProgramv)(GLenum type, GLsizei count, GLchar* const* strings);
// void (APIENTRYP ptrglGenRenderbuffers)(GLsizei n, GLuint* renderbuffers);
// void (APIENTRYP ptrglProgramUniformMatrix2dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniform2ui)(GLuint program, GLint location, GLuint v0, GLuint v1);
// void (APIENTRYP ptrglCopyTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglMultiDrawElementsBaseVertex)(GLenum mode, GLsizei* count, GLenum type, GLvoid* const* indices, GLsizei drawcount, GLint* basevertex);
// void (APIENTRYP ptrglGetCompressedTexImage)(GLenum target, GLint level, GLvoid* img);
// void (APIENTRYP ptrglVertexAttribI3ui)(GLuint index, GLuint x, GLuint y, GLuint z);
// void (APIENTRYP ptrglUniformMatrix3x4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglVertexAttribI1iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglGetBufferSubData)(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data);
// void (APIENTRYP ptrglTexImage1D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglFramebufferTexture)(GLenum target, GLenum attachment, GLuint texture, GLint level);
// void (APIENTRYP ptrglVertexAttribI1i)(GLuint index, GLint x);
// void (APIENTRYP ptrglGetVertexAttribIiv)(GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetBufferParameteri64v)(GLenum target, GLenum pname, GLint64* params);
// void (APIENTRYP ptrglBeginConditionalRender)(GLuint id, GLenum mode);
// void (APIENTRYP ptrglGetInternalformati64v)(GLenum target, GLenum internalformat, GLenum pname, GLsizei bufSize, GLint64* params);
// void (APIENTRYP ptrglVertexAttribI2ui)(GLuint index, GLuint x, GLuint y);
// void (APIENTRYP ptrglBlendEquation)(GLenum mode);
// void (APIENTRYP ptrglEnablei)(GLenum target, GLuint index);
// const GLubyte * (APIENTRYP ptrglGetStringi)(GLenum name, GLuint index);
// void (APIENTRYP ptrglProgramUniformMatrix3x2dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglDrawTransformFeedbackStreamInstanced)(GLenum mode, GLuint id, GLuint stream, GLsizei instancecount);
// void (APIENTRYP ptrglGetActiveSubroutineUniformName)(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name);
// void (APIENTRYP ptrglGetInteger64i_v)(GLenum target, GLuint index, GLint64* data);
// void (APIENTRYP ptrglDebugMessageInsert)(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, GLchar* buf);
// void (APIENTRYP ptrglProgramUniform1f)(GLuint program, GLint location, GLfloat v0);
// void (APIENTRYP ptrglGetSamplerParameterfv)(GLuint sampler, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglTexCoordP4ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglProgramUniform1uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglDeleteSamplers)(GLsizei count, GLuint* samplers);
// GLboolean (APIENTRYP ptrglIsShader)(GLuint shader);
// void (APIENTRYP ptrglCopyTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglGetUniformuiv)(GLuint program, GLint location, GLuint* params);
// void (APIENTRYP ptrglVertexAttribL2d)(GLuint index, GLdouble x, GLdouble y);
// void (APIENTRYP ptrglSecondaryColorP3ui)(GLenum type, GLuint color);
// GLboolean (APIENTRYP ptrglIsSampler)(GLuint sampler);
// void (APIENTRYP ptrglUniform2dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglDepthMask)(GLboolean flag);
// void (APIENTRYP ptrglPushDebugGroup)(GLenum source, GLuint id, GLsizei length, GLchar* message);
// void (APIENTRYP ptrglProgramUniform1dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglMultiDrawArraysIndirect)(GLenum mode, void* indirect, GLsizei drawcount, GLsizei stride);
// void (APIENTRYP ptrglMultiTexCoordP2uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglInvalidateSubFramebuffer)(GLenum target, GLsizei numAttachments, GLenum* attachments, GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglUniform3i)(GLint location, GLint v0, GLint v1, GLint v2);
// void (APIENTRYP ptrglProgramUniform2i)(GLuint program, GLint location, GLint v0, GLint v1);
// void (APIENTRYP ptrglShaderSource)(GLuint shader, GLsizei count, GLchar* const* string, GLint* length);
// void (APIENTRYP ptrglTexCoordP1uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglUniformBlockBinding)(GLuint program, GLuint uniformBlockIndex, GLuint uniformBlockBinding);
// void (APIENTRYP ptrglUniform4fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglGetVertexAttribdv)(GLuint index, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglGenBuffers)(GLsizei n, GLuint* buffers);
// void (APIENTRYP ptrglVertexAttribP4ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglGetInternalformativ)(GLenum target, GLenum internalformat, GLenum pname, GLsizei bufSize, GLint* params);
// void (APIENTRYP ptrglProgramUniformMatrix4x2dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglCompileShader)(GLuint shader);
// void (APIENTRYP ptrglBindRenderbuffer)(GLenum target, GLuint renderbuffer);
// void (APIENTRYP ptrglCompressedTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglBlendFunc)(GLenum sfactor, GLenum dfactor);
// void (APIENTRYP ptrglTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglInvalidateFramebuffer)(GLenum target, GLsizei numAttachments, GLenum* attachments);
// void (APIENTRYP ptrglBindTransformFeedback)(GLenum target, GLuint id);
// GLint (APIENTRYP ptrglGetFragDataIndex)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglTransformFeedbackVaryings)(GLuint program, GLsizei count, GLchar* const* varyings, GLenum bufferMode);
// void (APIENTRYP ptrglRenderbufferStorage)(GLenum target, GLenum internalformat, GLsizei width, GLsizei height);
// GLuint (APIENTRYP ptrglGetUniformBlockIndex)(GLuint program, GLchar* uniformBlockName);
// void (APIENTRYP ptrglObjectLabel)(GLenum identifier, GLuint name, GLsizei length, GLchar* label);
// void (APIENTRYP ptrglGetTexLevelParameterfv)(GLenum target, GLint level, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglProgramUniformMatrix4x3fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglGetShaderiv)(GLuint shader, GLenum pname, GLint* params);
// void (APIENTRYP ptrglProgramUniform4ui)(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3);
// void (APIENTRYP ptrglMultiTexCoordP3uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglProgramUniform3uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglGetPointerv)(GLenum pname, GLvoid** params);
// void (APIENTRYP ptrglMultiTexCoordP4uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglUniformMatrix4x2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglShaderStorageBlockBinding)(GLuint program, GLuint storageBlockIndex, GLuint storageBlockBinding);
// void (APIENTRYP ptrglShaderBinary)(GLsizei count, GLuint* shaders, GLenum binaryformat, GLvoid* binary, GLsizei length);
// void (APIENTRYP ptrglUniform1d)(GLint location, GLdouble x);
// void (APIENTRYP ptrglBlendEquationSeparatei)(GLuint buf, GLenum modeRGB, GLenum modeAlpha);
// void (APIENTRYP ptrglGenQueries)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglClampColor)(GLenum target, GLenum clamp);
// void (APIENTRYP ptrglPauseTransformFeedback)();
// void (APIENTRYP ptrglLinkProgram)(GLuint program);
// GLenum (APIENTRYP ptrglCheckFramebufferStatus)(GLenum target);
// void (APIENTRYP ptrglPointParameteri)(GLenum pname, GLint param);
// void (APIENTRYP ptrglVertexAttribI4usv)(GLuint index, GLushort* v);
// void (APIENTRYP ptrglPointParameteriv)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglVertexAttribIPointer)(GLuint index, GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglGenTransformFeedbacks)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglUniform1uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglGenerateMipmap)(GLenum target);
// void (APIENTRYP ptrglGetTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglUniform2uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglUniformSubroutinesuiv)(GLenum shadertype, GLsizei count, GLuint* indices);
// void (APIENTRYP ptrglClearDepthf)(GLfloat d);
// void (APIENTRYP ptrglDeleteVertexArrays)(GLsizei n, GLuint* arrays);
// GLboolean (APIENTRYP ptrglIsEnabled)(GLenum cap);
// void (APIENTRYP ptrglUniform4ui)(GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3);
// void (APIENTRYP ptrglGetFramebufferAttachmentParameteriv)(GLenum target, GLenum attachment, GLenum pname, GLint* params);
// void (APIENTRYP ptrglUniform1dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglColorMask)(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
// const GLubyte * (APIENTRYP ptrglGetString)(GLenum name);
// void (APIENTRYP ptrglHint)(GLenum target, GLenum mode);
// void (APIENTRYP ptrglClearBufferData)(GLenum target, GLenum internalformat, GLenum format, GLenum type, void* data);
// void (APIENTRYP ptrglUniform2d)(GLint location, GLdouble x, GLdouble y);
// void (APIENTRYP ptrglGetInteger64v)(GLenum pname, GLint64* params);
// void (APIENTRYP ptrglUniform1f)(GLint location, GLfloat v0);
// void (APIENTRYP ptrglTexBuffer)(GLenum target, GLenum internalformat, GLuint buffer);
// void (APIENTRYP ptrglEndConditionalRender)();
// void (APIENTRYP ptrglVertexP3ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrglGetProgramInfoLog)(GLuint program, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrglVertexAttribI2i)(GLuint index, GLint x, GLint y);
// GLboolean (APIENTRYP ptrglIsTransformFeedback)(GLuint id);
// void (APIENTRYP ptrglLogicOp)(GLenum opcode);
// void (APIENTRYP ptrglDispatchComputeIndirect)(GLintptr indirect);
// void (APIENTRYP ptrglBindTexture)(GLenum target, GLuint texture);
// void (APIENTRYP ptrglDisableVertexAttribArray)(GLuint index);
// void (APIENTRYP ptrglQueryCounter)(GLuint id, GLenum target);
// void (APIENTRYP ptrglProgramUniform3f)(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2);
// GLboolean (APIENTRYP ptrglIsEnabledi)(GLenum target, GLuint index);
// void (APIENTRYP ptrglVertexAttribL3d)(GLuint index, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglUniform4dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglBindVertexArray)(GLuint array);
// void (APIENTRYP ptrglProgramUniform3dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglGetFramebufferParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglCullFace)(GLenum mode);
// void (APIENTRYP ptrglMultiDrawArrays)(GLenum mode, GLint* first, GLsizei* count, GLsizei drawcount);
// void (APIENTRYP ptrglProgramUniformMatrix2x3dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglValidateProgramPipeline)(GLuint pipeline);
// void (APIENTRYP ptrglProgramUniformMatrix3x2fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglVertexAttribI3i)(GLuint index, GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglGetProgramResourceiv)(GLuint program, GLenum programInterface, GLuint index, GLsizei propCount, GLenum* props, GLsizei bufSize, GLsizei* length, GLint* params);
// void (APIENTRYP ptrglGetUniformSubroutineuiv)(GLenum shadertype, GLint location, GLuint* params);
// void (APIENTRYP ptrglProgramUniform2iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglGenFramebuffers)(GLsizei n, GLuint* framebuffers);
// void (APIENTRYP ptrglDrawTransformFeedbackInstanced)(GLenum mode, GLuint id, GLsizei instancecount);
// GLboolean (APIENTRYP ptrglIsBuffer)(GLuint buffer);
// void (APIENTRYP ptrglAttachShader)(GLuint program, GLuint shader);
// void (APIENTRYP ptrglDrawElementsIndirect)(GLenum mode, GLenum type, GLvoid* indirect);
// void (APIENTRYP ptrglProgramUniform1ui)(GLuint program, GLint location, GLuint v0);
// void (APIENTRYP ptrglGetUniformIndices)(GLuint program, GLsizei uniformCount, GLchar* const* uniformNames, GLuint* uniformIndices);
// void (APIENTRYP ptrglUniform3f)(GLint location, GLfloat v0, GLfloat v1, GLfloat v2);
// void (APIENTRYP ptrglBindFragDataLocationIndexed)(GLuint program, GLuint colorNumber, GLuint index, GLchar* name);
// void (APIENTRYP ptrglTextureView)(GLuint texture, GLenum target, GLuint origtexture, GLenum internalformat, GLuint minlevel, GLuint numlevels, GLuint minlayer, GLuint numlayers);
// void (APIENTRYP ptrglGetQueryiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglFramebufferTexture1D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
// void (APIENTRYP ptrglPointSize)(GLfloat size);
// void (APIENTRYP ptrglVertexAttribI4i)(GLuint index, GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglPopDebugGroup)();
// void (APIENTRYP ptrglStencilFunc)(GLenum func, GLint ref, GLuint mask);
// void (APIENTRYP ptrglTexImage3DMultisample)(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrglProgramUniformMatrix2x3fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglVertexAttribL1d)(GLuint index, GLdouble x);
// void (APIENTRYP ptrglClearBufferiv)(GLenum buffer, GLint drawbuffer, GLint* value);
// void (APIENTRYP ptrglBindVertexBuffer)(GLuint bindingindex, GLuint buffer, GLintptr offset, GLsizei stride);
// void (APIENTRYP ptrglDisablei)(GLenum target, GLuint index);
// void (APIENTRYP ptrglVertexAttribP1uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrglUniform3dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglFramebufferTextureLayer)(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer);
// void (APIENTRYP ptrglFramebufferTexture2D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
// void (APIENTRYP ptrglReleaseShaderCompiler)();
// void (APIENTRYP ptrglProgramUniformMatrix3x4fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix2x4dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglPatchParameteri)(GLenum pname, GLint value);
// void (APIENTRYP ptrglGetProgramPipelineInfoLog)(GLuint pipeline, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrglGetActiveSubroutineName)(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name);
// void (APIENTRYP ptrglGetTexParameterIuiv)(GLenum target, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglProgramUniform1fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglTexCoordP3uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglGenProgramPipelines)(GLsizei n, GLuint* pipelines);
// GLuint (APIENTRYP ptrglGetDebugMessageLog)(GLuint count, GLsizei bufsize, GLenum* sources, GLenum* types, GLuint* ids, GLenum* severities, GLsizei* lengths, GLchar* messageLog);
// void (APIENTRYP ptrglSamplerParameteriv)(GLuint sampler, GLenum pname, GLint* param);
// void (APIENTRYP ptrglFrontFace)(GLenum mode);
// void (APIENTRYP ptrglDeleteBuffers)(GLsizei n, GLuint* buffers);
// void (APIENTRYP ptrglCompressedTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglCompressedTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglBindProgramPipeline)(GLuint pipeline);
// void (APIENTRYP ptrglUniform1iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglVertexAttribDivisor)(GLuint index, GLuint divisor);
// void (APIENTRYP ptrglBindBufferBase)(GLenum target, GLuint index, GLuint buffer);
// void (APIENTRYP ptrglColorP4ui)(GLenum type, GLuint color);
// GLuint (APIENTRYP ptrglGetProgramResourceIndex)(GLuint program, GLenum programInterface, GLchar* name);
// void (APIENTRYP ptrglFramebufferParameteri)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrglValidateProgram)(GLuint program);
// void (APIENTRYP ptrglDrawRangeElements)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices);
// void (APIENTRYP ptrglSamplerParameteri)(GLuint sampler, GLenum pname, GLint param);
// void (APIENTRYP ptrglGetActiveAttrib)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrglSecondaryColorP3uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrglSamplerParameterfv)(GLuint sampler, GLenum pname, GLfloat* param);
// void (APIENTRYP ptrglDepthRange)(GLdouble near, GLdouble far);
// void (APIENTRYP ptrglUniform3uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglSamplerParameterf)(GLuint sampler, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglUniform2f)(GLint location, GLfloat v0, GLfloat v1);
// void (APIENTRYP ptrglSamplerParameterIuiv)(GLuint sampler, GLenum pname, GLuint* param);
// void (APIENTRYP ptrglDrawArraysIndirect)(GLenum mode, GLvoid* indirect);
// void (APIENTRYP ptrglGenTextures)(GLsizei n, GLuint* textures);
// void (APIENTRYP ptrglBindBufferRange)(GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size);
// void (APIENTRYP ptrglUniformMatrix4x2dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglViewport)(GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglMultiTexCoordP1uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglUniformMatrix2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglProgramUniform2f)(GLuint program, GLint location, GLfloat v0, GLfloat v1);
// GLint (APIENTRYP ptrglGetUniformLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglTexStorage2DMultisample)(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrglGetFloatv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglUniformMatrix3x2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglBindSampler)(GLuint unit, GLuint sampler);
// void (APIENTRYP ptrglUniform4iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglMemoryBarrier)(GLbitfield barriers);
// void (APIENTRYP ptrglVertexAttribPointer)(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglBufferData)(GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage);
// void (APIENTRYP ptrglProgramUniform2dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglLineWidth)(GLfloat width);
// void (APIENTRYP ptrglColorP4uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrglProgramUniform3i)(GLuint program, GLint location, GLint v0, GLint v1, GLint v2);
// void (APIENTRYP ptrglDrawElements)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices);
// GLint (APIENTRYP ptrglGetProgramResourceLocation)(GLuint program, GLenum programInterface, GLchar* name);
// void (APIENTRYP ptrglInvalidateBufferSubData)(GLuint buffer, GLintptr offset, GLsizeiptr length);
// void (APIENTRYP ptrglTexStorage2D)(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglDeleteQueries)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglGetTexImage)(GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels);
// GLenum (APIENTRYP ptrglClientWaitSync)(GLsync sync, GLbitfield flags, GLuint64 timeout);
// void (APIENTRYP ptrglDrawElementsInstancedBaseVertex)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei instancecount, GLint basevertex);
// void (APIENTRYP ptrglPixelStoref)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglTexParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglFramebufferRenderbuffer)(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer);
// void (APIENTRYP ptrglClearBufferSubData)(GLenum target, GLenum internalformat, GLintptr offset, GLsizeiptr size, GLenum format, GLenum type, void* data);
// void (APIENTRYP ptrglProgramUniform2d)(GLuint program, GLint location, GLdouble v0, GLdouble v1);
// void (APIENTRYP ptrglSampleCoverage)(GLfloat value, GLboolean invert);
// GLboolean (APIENTRYP ptrglIsVertexArray)(GLuint array);
// void (APIENTRYP ptrglDispatchCompute)(GLuint num_groups_x, GLuint num_groups_y, GLuint num_groups_z);
// GLsync (APIENTRYP ptrglFenceSync)(GLenum condition, GLbitfield flags);
// void (APIENTRYP ptrglGetVertexAttribiv)(GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrglFlushMappedBufferRange)(GLenum target, GLintptr offset, GLsizeiptr length);
// void (APIENTRYP ptrglGetVertexAttribfv)(GLuint index, GLenum pname, GLfloat* params);
// GLuint (APIENTRYP ptrglCreateShader)(GLenum type);
// void (APIENTRYP ptrglVertexAttribFormat)(GLuint attribindex, GLint size, GLenum type, GLboolean normalized, GLuint relativeoffset);
// void (APIENTRYP ptrglMultiDrawElements)(GLenum mode, GLsizei* count, GLenum type, GLvoid* const* indices, GLsizei drawcount);
// void (APIENTRYP ptrglCopyImageSubData)(GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei srcWidth, GLsizei srcHeight, GLsizei srcDepth);
// void (APIENTRYP ptrglDrawElementsBaseVertex)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLint basevertex);
// void (APIENTRYP ptrglStencilMask)(GLuint mask);
// void (APIENTRYP ptrglVertexAttribL4dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglTexParameteri)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrglTexCoordP1ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglVertexAttribL4d)(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglTexCoordP2ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglGetUniformfv)(GLuint program, GLint location, GLfloat* params);
// void (APIENTRYP ptrglVertexAttribI2iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglGetBooleani_v)(GLenum target, GLuint index, GLboolean* data);
// GLenum (APIENTRYP ptrglGetError)();
// void (APIENTRYP ptrglUniformMatrix2x4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglGetObjectLabel)(GLenum identifier, GLuint name, GLsizei bufSize, GLsizei* length, GLchar* label);
// void (APIENTRYP ptrglTexImage2D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglProgramUniformMatrix3fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglTexCoordP3ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglGetDoublei_v)(GLenum target, GLuint index, GLdouble* data);
// void (APIENTRYP ptrglTexImage3D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglGenSamplers)(GLsizei count, GLuint* samplers);
// void (APIENTRYP ptrglVertexAttribI4uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglGetUniformdv)(GLuint program, GLint location, GLdouble* params);
// void (APIENTRYP ptrglProgramUniform4iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglProgramUniform1i)(GLuint program, GLint location, GLint v0);
// void (APIENTRYP ptrglMultiTexCoordP1ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglProgramUniformMatrix4dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniform1fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglDrawElementsInstancedBaseVertexBaseInstance)(GLenum mode, GLsizei count, GLenum type, void* indices, GLsizei instancecount, GLint basevertex, GLuint baseinstance);
// 
// void callUniformMatrix4x3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix4x3fv)(location, count, transpose, value); }
// int cancallUniformMatrix4x3fv() { return ((ptrglUniformMatrix4x3fv == NULL) ? 0 : 1); }
// void callGetTransformFeedbackVarying(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLsizei* size, GLenum* type_, GLchar* name) {  (*ptrglGetTransformFeedbackVarying)(program, index, bufSize, length, size, type_, name); }
// int cancallGetTransformFeedbackVarying() { return ((ptrglGetTransformFeedbackVarying == NULL) ? 0 : 1); }
// void callProgramUniform4d(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {  (*ptrglProgramUniform4d)(program, location, v0, v1, v2, v3); }
// int cancallProgramUniform4d() { return ((ptrglProgramUniform4d == NULL) ? 0 : 1); }
// void callGetProgramResourceName(GLuint program, GLenum programInterface, GLuint index, GLsizei bufSize, GLsizei* length, GLchar* name) {  (*ptrglGetProgramResourceName)(program, programInterface, index, bufSize, length, name); }
// int cancallGetProgramResourceName() { return ((ptrglGetProgramResourceName == NULL) ? 0 : 1); }
// void callDeleteProgramPipelines(GLsizei n, GLuint* pipelines) {  (*ptrglDeleteProgramPipelines)(n, pipelines); }
// int cancallDeleteProgramPipelines() { return ((ptrglDeleteProgramPipelines == NULL) ? 0 : 1); }
// void callDisable(GLenum cap) {  (*ptrglDisable)(cap); }
// int cancallDisable() { return ((ptrglDisable == NULL) ? 0 : 1); }
// void callVertexAttribLPointer(GLuint index, GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {  (*ptrglVertexAttribLPointer)(index, size, type_, stride, pointer); }
// int cancallVertexAttribLPointer() { return ((ptrglVertexAttribLPointer == NULL) ? 0 : 1); }
// void callVertexAttribP3ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {  (*ptrglVertexAttribP3ui)(index, type_, normalized, value); }
// int cancallVertexAttribP3ui() { return ((ptrglVertexAttribP3ui == NULL) ? 0 : 1); }
// void callProvokingVertex(GLenum mode) {  (*ptrglProvokingVertex)(mode); }
// int cancallProvokingVertex() { return ((ptrglProvokingVertex == NULL) ? 0 : 1); }
// void callDrawArraysInstanced(GLenum mode, GLint first, GLsizei count, GLsizei instancecount) {  (*ptrglDrawArraysInstanced)(mode, first, count, instancecount); }
// int cancallDrawArraysInstanced() { return ((ptrglDrawArraysInstanced == NULL) ? 0 : 1); }
// void callNormalP3uiv(GLenum type_, GLuint* coords) {  (*ptrglNormalP3uiv)(type_, coords); }
// int cancallNormalP3uiv() { return ((ptrglNormalP3uiv == NULL) ? 0 : 1); }
// void callUniform2iv(GLint location, GLsizei count, GLint* value) {  (*ptrglUniform2iv)(location, count, value); }
// int cancallUniform2iv() { return ((ptrglUniform2iv == NULL) ? 0 : 1); }
// void callProgramUniform1d(GLuint program, GLint location, GLdouble v0) {  (*ptrglProgramUniform1d)(program, location, v0); }
// int cancallProgramUniform1d() { return ((ptrglProgramUniform1d == NULL) ? 0 : 1); }
// void callPolygonOffset(GLfloat factor, GLfloat units) {  (*ptrglPolygonOffset)(factor, units); }
// int cancallPolygonOffset() { return ((ptrglPolygonOffset == NULL) ? 0 : 1); }
// void callUniformMatrix3x4dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix3x4dv)(location, count, transpose, value); }
// int cancallUniformMatrix3x4dv() { return ((ptrglUniformMatrix3x4dv == NULL) ? 0 : 1); }
// void callTexStorage3D(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth) {  (*ptrglTexStorage3D)(target, levels, internalformat, width, height, depth); }
// int cancallTexStorage3D() { return ((ptrglTexStorage3D == NULL) ? 0 : 1); }
// void callGetTexParameterIiv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetTexParameterIiv)(target, pname, params); }
// int cancallGetTexParameterIiv() { return ((ptrglGetTexParameterIiv == NULL) ? 0 : 1); }
// void callProgramUniformMatrix4fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix4fv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix4fv() { return ((ptrglProgramUniformMatrix4fv == NULL) ? 0 : 1); }
// void callVertexAttribP2uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {  (*ptrglVertexAttribP2uiv)(index, type_, normalized, value); }
// int cancallVertexAttribP2uiv() { return ((ptrglVertexAttribP2uiv == NULL) ? 0 : 1); }
// void callCopyTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {  (*ptrglCopyTexSubImage1D)(target, level, xoffset, x, y, width); }
// int cancallCopyTexSubImage1D() { return ((ptrglCopyTexSubImage1D == NULL) ? 0 : 1); }
// void callGetSynciv(GLsync sync, GLenum pname, GLsizei bufSize, GLsizei* length, GLint* values) {  (*ptrglGetSynciv)(sync, pname, bufSize, length, values); }
// int cancallGetSynciv() { return ((ptrglGetSynciv == NULL) ? 0 : 1); }
// void callUniform4f(GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {  (*ptrglUniform4f)(location, v0, v1, v2, v3); }
// int cancallUniform4f() { return ((ptrglUniform4f == NULL) ? 0 : 1); }
// void callProgramUniformMatrix2fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix2fv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix2fv() { return ((ptrglProgramUniformMatrix2fv == NULL) ? 0 : 1); }
// void callDeleteShader(GLuint shader) {  (*ptrglDeleteShader)(shader); }
// int cancallDeleteShader() { return ((ptrglDeleteShader == NULL) ? 0 : 1); }
// void callBindBuffer(GLenum target, GLuint buffer) {  (*ptrglBindBuffer)(target, buffer); }
// int cancallBindBuffer() { return ((ptrglBindBuffer == NULL) ? 0 : 1); }
// void callEnable(GLenum cap) {  (*ptrglEnable)(cap); }
// int cancallEnable() { return ((ptrglEnable == NULL) ? 0 : 1); }
// void callWaitSync(GLsync sync, GLbitfield flags, GLuint64 timeout) {  (*ptrglWaitSync)(sync, flags, timeout); }
// int cancallWaitSync() { return ((ptrglWaitSync == NULL) ? 0 : 1); }
// void callBindAttribLocation(GLuint program, GLuint index, GLchar* name) {  (*ptrglBindAttribLocation)(program, index, name); }
// int cancallBindAttribLocation() { return ((ptrglBindAttribLocation == NULL) ? 0 : 1); }
// void callVertexAttribI4iv(GLuint index, GLint* v) {  (*ptrglVertexAttribI4iv)(index, v); }
// int cancallVertexAttribI4iv() { return ((ptrglVertexAttribI4iv == NULL) ? 0 : 1); }
// void callGetProgramStageiv(GLuint program, GLenum shadertype, GLenum pname, GLint* values) {  (*ptrglGetProgramStageiv)(program, shadertype, pname, values); }
// int cancallGetProgramStageiv() { return ((ptrglGetProgramStageiv == NULL) ? 0 : 1); }
// void callTexCoordP2uiv(GLenum type_, GLuint* coords) {  (*ptrglTexCoordP2uiv)(type_, coords); }
// int cancallTexCoordP2uiv() { return ((ptrglTexCoordP2uiv == NULL) ? 0 : 1); }
// void callMultiTexCoordP4ui(GLenum texture, GLenum type_, GLuint coords) {  (*ptrglMultiTexCoordP4ui)(texture, type_, coords); }
// int cancallMultiTexCoordP4ui() { return ((ptrglMultiTexCoordP4ui == NULL) ? 0 : 1); }
// void callUniform3iv(GLint location, GLsizei count, GLint* value) {  (*ptrglUniform3iv)(location, count, value); }
// int cancallUniform3iv() { return ((ptrglUniform3iv == NULL) ? 0 : 1); }
// void callGetIntegeri_v(GLenum target, GLuint index, GLint* data) {  (*ptrglGetIntegeri_v)(target, index, data); }
// int cancallGetIntegeri_v() { return ((ptrglGetIntegeri_v == NULL) ? 0 : 1); }
// void callDrawElementsInstanced(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLsizei instancecount) {  (*ptrglDrawElementsInstanced)(mode, count, type_, indices, instancecount); }
// int cancallDrawElementsInstanced() { return ((ptrglDrawElementsInstanced == NULL) ? 0 : 1); }
// void callMinSampleShading(GLfloat value) {  (*ptrglMinSampleShading)(value); }
// int cancallMinSampleShading() { return ((ptrglMinSampleShading == NULL) ? 0 : 1); }
// void callGetShaderSource(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {  (*ptrglGetShaderSource)(shader, bufSize, length, source); }
// int cancallGetShaderSource() { return ((ptrglGetShaderSource == NULL) ? 0 : 1); }
// void callProgramUniform2uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {  (*ptrglProgramUniform2uiv)(program, location, count, value); }
// int cancallProgramUniform2uiv() { return ((ptrglProgramUniform2uiv == NULL) ? 0 : 1); }
// void callStencilMaskSeparate(GLenum face, GLuint mask) {  (*ptrglStencilMaskSeparate)(face, mask); }
// int cancallStencilMaskSeparate() { return ((ptrglStencilMaskSeparate == NULL) ? 0 : 1); }
// void callScissor(GLint x, GLint y, GLsizei width, GLsizei height) {  (*ptrglScissor)(x, y, width, height); }
// int cancallScissor() { return ((ptrglScissor == NULL) ? 0 : 1); }
// void callGetShaderPrecisionFormat(GLenum shadertype, GLenum precisiontype, GLint* range_, GLint* precision) {  (*ptrglGetShaderPrecisionFormat)(shadertype, precisiontype, range_, precision); }
// int cancallGetShaderPrecisionFormat() { return ((ptrglGetShaderPrecisionFormat == NULL) ? 0 : 1); }
// void callVertexAttribI4ui(GLuint index, GLuint x, GLuint y, GLuint z, GLuint w) {  (*ptrglVertexAttribI4ui)(index, x, y, z, w); }
// int cancallVertexAttribI4ui() { return ((ptrglVertexAttribI4ui == NULL) ? 0 : 1); }
// void callColorP3uiv(GLenum type_, GLuint* color) {  (*ptrglColorP3uiv)(type_, color); }
// int cancallColorP3uiv() { return ((ptrglColorP3uiv == NULL) ? 0 : 1); }
// void callGetActiveUniformBlockName(GLuint program, GLuint uniformBlockIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformBlockName) {  (*ptrglGetActiveUniformBlockName)(program, uniformBlockIndex, bufSize, length, uniformBlockName); }
// int cancallGetActiveUniformBlockName() { return ((ptrglGetActiveUniformBlockName == NULL) ? 0 : 1); }
// void callTexStorage3DMultisample(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations) {  (*ptrglTexStorage3DMultisample)(target, samples, internalformat, width, height, depth, fixedsamplelocations); }
// int cancallTexStorage3DMultisample() { return ((ptrglTexStorage3DMultisample == NULL) ? 0 : 1); }
// void callVertexAttribI3uiv(GLuint index, GLuint* v) {  (*ptrglVertexAttribI3uiv)(index, v); }
// int cancallVertexAttribI3uiv() { return ((ptrglVertexAttribI3uiv == NULL) ? 0 : 1); }
// void callVertexP2uiv(GLenum type_, GLuint* value) {  (*ptrglVertexP2uiv)(type_, value); }
// int cancallVertexP2uiv() { return ((ptrglVertexP2uiv == NULL) ? 0 : 1); }
// void callMultiDrawElementsIndirect(GLenum mode, GLenum type_, void* indirect, GLsizei drawcount, GLsizei stride) {  (*ptrglMultiDrawElementsIndirect)(mode, type_, indirect, drawcount, stride); }
// int cancallMultiDrawElementsIndirect() { return ((ptrglMultiDrawElementsIndirect == NULL) ? 0 : 1); }
// void callUniform4uiv(GLint location, GLsizei count, GLuint* value) {  (*ptrglUniform4uiv)(location, count, value); }
// int cancallUniform4uiv() { return ((ptrglUniform4uiv == NULL) ? 0 : 1); }
// void callGetTexParameteriv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetTexParameteriv)(target, pname, params); }
// int cancallGetTexParameteriv() { return ((ptrglGetTexParameteriv == NULL) ? 0 : 1); }
// void callVertexP2ui(GLenum type_, GLuint value) {  (*ptrglVertexP2ui)(type_, value); }
// int cancallVertexP2ui() { return ((ptrglVertexP2ui == NULL) ? 0 : 1); }
// void callGenVertexArrays(GLsizei n, GLuint* arrays) {  (*ptrglGenVertexArrays)(n, arrays); }
// int cancallGenVertexArrays() { return ((ptrglGenVertexArrays == NULL) ? 0 : 1); }
// GLint callGetProgramResourceLocationIndex(GLuint program, GLenum programInterface, GLchar* name) { return (*ptrglGetProgramResourceLocationIndex)(program, programInterface, name); }
// int cancallGetProgramResourceLocationIndex() { return ((ptrglGetProgramResourceLocationIndex == NULL) ? 0 : 1); }
// void callGetSamplerParameteriv(GLuint sampler, GLenum pname, GLint* params) {  (*ptrglGetSamplerParameteriv)(sampler, pname, params); }
// int cancallGetSamplerParameteriv() { return ((ptrglGetSamplerParameteriv == NULL) ? 0 : 1); }
// void callDepthRangef(GLfloat n, GLfloat f) {  (*ptrglDepthRangef)(n, f); }
// int cancallDepthRangef() { return ((ptrglDepthRangef == NULL) ? 0 : 1); }
// void callVertexAttribI1ui(GLuint index, GLuint x) {  (*ptrglVertexAttribI1ui)(index, x); }
// int cancallVertexAttribI1ui() { return ((ptrglVertexAttribI1ui == NULL) ? 0 : 1); }
// void callClearColor(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {  (*ptrglClearColor)(red, green, blue, alpha); }
// int cancallClearColor() { return ((ptrglClearColor == NULL) ? 0 : 1); }
// void callTexParameterIuiv(GLenum target, GLenum pname, GLuint* params) {  (*ptrglTexParameterIuiv)(target, pname, params); }
// int cancallTexParameterIuiv() { return ((ptrglTexParameterIuiv == NULL) ? 0 : 1); }
// void callProgramUniform2fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {  (*ptrglProgramUniform2fv)(program, location, count, value); }
// int cancallProgramUniform2fv() { return ((ptrglProgramUniform2fv == NULL) ? 0 : 1); }
// void callVertexAttribBinding(GLuint attribindex, GLuint bindingindex) {  (*ptrglVertexAttribBinding)(attribindex, bindingindex); }
// int cancallVertexAttribBinding() { return ((ptrglVertexAttribBinding == NULL) ? 0 : 1); }
// void callTexCoordP4uiv(GLenum type_, GLuint* coords) {  (*ptrglTexCoordP4uiv)(type_, coords); }
// int cancallTexCoordP4uiv() { return ((ptrglTexCoordP4uiv == NULL) ? 0 : 1); }
// void callPatchParameterfv(GLenum pname, GLfloat* values) {  (*ptrglPatchParameterfv)(pname, values); }
// int cancallPatchParameterfv() { return ((ptrglPatchParameterfv == NULL) ? 0 : 1); }
// void callCompressedTexImage2D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexImage2D)(target, level, internalformat, width, height, border, imageSize, data); }
// int cancallCompressedTexImage2D() { return ((ptrglCompressedTexImage2D == NULL) ? 0 : 1); }
// void callActiveTexture(GLenum texture) {  (*ptrglActiveTexture)(texture); }
// int cancallActiveTexture() { return ((ptrglActiveTexture == NULL) ? 0 : 1); }
// void callGetProgramBinary(GLuint program, GLsizei bufSize, GLsizei* length, GLenum* binaryFormat, GLvoid* binary) {  (*ptrglGetProgramBinary)(program, bufSize, length, binaryFormat, binary); }
// int cancallGetProgramBinary() { return ((ptrglGetProgramBinary == NULL) ? 0 : 1); }
// void callUniform1ui(GLint location, GLuint v0) {  (*ptrglUniform1ui)(location, v0); }
// int cancallUniform1ui() { return ((ptrglUniform1ui == NULL) ? 0 : 1); }
// void callCopyTexImage1D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border) {  (*ptrglCopyTexImage1D)(target, level, internalformat, x, y, width, border); }
// int cancallCopyTexImage1D() { return ((ptrglCopyTexImage1D == NULL) ? 0 : 1); }
// void callDepthRangeIndexed(GLuint index, GLdouble n, GLdouble f) {  (*ptrglDepthRangeIndexed)(index, n, f); }
// int cancallDepthRangeIndexed() { return ((ptrglDepthRangeIndexed == NULL) ? 0 : 1); }
// void callCopyTexImage2D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {  (*ptrglCopyTexImage2D)(target, level, internalformat, x, y, width, height, border); }
// int cancallCopyTexImage2D() { return ((ptrglCopyTexImage2D == NULL) ? 0 : 1); }
// void callScissorArrayv(GLuint first, GLsizei count, GLint* v) {  (*ptrglScissorArrayv)(first, count, v); }
// int cancallScissorArrayv() { return ((ptrglScissorArrayv == NULL) ? 0 : 1); }
// void callGetShaderInfoLog(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {  (*ptrglGetShaderInfoLog)(shader, bufSize, length, infoLog); }
// int cancallGetShaderInfoLog() { return ((ptrglGetShaderInfoLog == NULL) ? 0 : 1); }
// void callGetMultisamplefv(GLenum pname, GLuint index, GLfloat* val) {  (*ptrglGetMultisamplefv)(pname, index, val); }
// int cancallGetMultisamplefv() { return ((ptrglGetMultisamplefv == NULL) ? 0 : 1); }
// void callDrawArrays(GLenum mode, GLint first, GLsizei count) {  (*ptrglDrawArrays)(mode, first, count); }
// int cancallDrawArrays() { return ((ptrglDrawArrays == NULL) ? 0 : 1); }
// void callGetActiveUniformBlockiv(GLuint program, GLuint uniformBlockIndex, GLenum pname, GLint* params) {  (*ptrglGetActiveUniformBlockiv)(program, uniformBlockIndex, pname, params); }
// int cancallGetActiveUniformBlockiv() { return ((ptrglGetActiveUniformBlockiv == NULL) ? 0 : 1); }
// void callUniformMatrix3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix3fv)(location, count, transpose, value); }
// int cancallUniformMatrix3fv() { return ((ptrglUniformMatrix3fv == NULL) ? 0 : 1); }
// GLuint callCreateProgram() { return (*ptrglCreateProgram)(); }
// int cancallCreateProgram() { return ((ptrglCreateProgram == NULL) ? 0 : 1); }
// void callVertexAttribI1uiv(GLuint index, GLuint* v) {  (*ptrglVertexAttribI1uiv)(index, v); }
// int cancallVertexAttribI1uiv() { return ((ptrglVertexAttribI1uiv == NULL) ? 0 : 1); }
// void callGetQueryObjectuiv(GLuint id, GLenum pname, GLuint* params) {  (*ptrglGetQueryObjectuiv)(id, pname, params); }
// int cancallGetQueryObjectuiv() { return ((ptrglGetQueryObjectuiv == NULL) ? 0 : 1); }
// void callEndQuery(GLenum target) {  (*ptrglEndQuery)(target); }
// int cancallEndQuery() { return ((ptrglEndQuery == NULL) ? 0 : 1); }
// void callClearDepth(GLdouble depth) {  (*ptrglClearDepth)(depth); }
// int cancallClearDepth() { return ((ptrglClearDepth == NULL) ? 0 : 1); }
// void callGetFloati_v(GLenum target, GLuint index, GLfloat* data) {  (*ptrglGetFloati_v)(target, index, data); }
// int cancallGetFloati_v() { return ((ptrglGetFloati_v == NULL) ? 0 : 1); }
// void callPointParameterfv(GLenum pname, GLfloat* params) {  (*ptrglPointParameterfv)(pname, params); }
// int cancallPointParameterfv() { return ((ptrglPointParameterfv == NULL) ? 0 : 1); }
// void callGetQueryObjectiv(GLuint id, GLenum pname, GLint* params) {  (*ptrglGetQueryObjectiv)(id, pname, params); }
// int cancallGetQueryObjectiv() { return ((ptrglGetQueryObjectiv == NULL) ? 0 : 1); }
// void callGetActiveUniformName(GLuint program, GLuint uniformIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformName) {  (*ptrglGetActiveUniformName)(program, uniformIndex, bufSize, length, uniformName); }
// int cancallGetActiveUniformName() { return ((ptrglGetActiveUniformName == NULL) ? 0 : 1); }
// void callNormalP3ui(GLenum type_, GLuint coords) {  (*ptrglNormalP3ui)(type_, coords); }
// int cancallNormalP3ui() { return ((ptrglNormalP3ui == NULL) ? 0 : 1); }
// void callActiveShaderProgram(GLuint pipeline, GLuint program) {  (*ptrglActiveShaderProgram)(pipeline, program); }
// int cancallActiveShaderProgram() { return ((ptrglActiveShaderProgram == NULL) ? 0 : 1); }
// void callVertexAttribI3iv(GLuint index, GLint* v) {  (*ptrglVertexAttribI3iv)(index, v); }
// int cancallVertexAttribI3iv() { return ((ptrglVertexAttribI3iv == NULL) ? 0 : 1); }
// GLboolean callIsFramebuffer(GLuint framebuffer) { return (*ptrglIsFramebuffer)(framebuffer); }
// int cancallIsFramebuffer() { return ((ptrglIsFramebuffer == NULL) ? 0 : 1); }
// void callBindImageTexture(GLuint unit, GLuint texture, GLint level, GLboolean layered, GLint layer, GLenum access, GLenum format) {  (*ptrglBindImageTexture)(unit, texture, level, layered, layer, access, format); }
// int cancallBindImageTexture() { return ((ptrglBindImageTexture == NULL) ? 0 : 1); }
// void callGetActiveUniform(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type_, GLchar* name) {  (*ptrglGetActiveUniform)(program, index, bufSize, length, size, type_, name); }
// int cancallGetActiveUniform() { return ((ptrglGetActiveUniform == NULL) ? 0 : 1); }
// void callUniform1i(GLint location, GLint v0) {  (*ptrglUniform1i)(location, v0); }
// int cancallUniform1i() { return ((ptrglUniform1i == NULL) ? 0 : 1); }
// void callObjectPtrLabel(void* ptr, GLsizei length, GLchar* label) {  (*ptrglObjectPtrLabel)(ptr, length, label); }
// int cancallObjectPtrLabel() { return ((ptrglObjectPtrLabel == NULL) ? 0 : 1); }
// void callUseProgramStages(GLuint pipeline, GLbitfield stages, GLuint program) {  (*ptrglUseProgramStages)(pipeline, stages, program); }
// int cancallUseProgramStages() { return ((ptrglUseProgramStages == NULL) ? 0 : 1); }
// void callGetTexLevelParameteriv(GLenum target, GLint level, GLenum pname, GLint* params) {  (*ptrglGetTexLevelParameteriv)(target, level, pname, params); }
// int cancallGetTexLevelParameteriv() { return ((ptrglGetTexLevelParameteriv == NULL) ? 0 : 1); }
// void callVertexAttribLFormat(GLuint attribindex, GLint size, GLenum type_, GLuint relativeoffset) {  (*ptrglVertexAttribLFormat)(attribindex, size, type_, relativeoffset); }
// int cancallVertexAttribLFormat() { return ((ptrglVertexAttribLFormat == NULL) ? 0 : 1); }
// void callGetSamplerParameterIiv(GLuint sampler, GLenum pname, GLint* params) {  (*ptrglGetSamplerParameterIiv)(sampler, pname, params); }
// int cancallGetSamplerParameterIiv() { return ((ptrglGetSamplerParameterIiv == NULL) ? 0 : 1); }
// void callBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {  (*ptrglBufferSubData)(target, offset, size, data); }
// int cancallBufferSubData() { return ((ptrglBufferSubData == NULL) ? 0 : 1); }
// void callTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexSubImage2D)(target, level, xoffset, yoffset, width, height, format, type_, pixels); }
// int cancallTexSubImage2D() { return ((ptrglTexSubImage2D == NULL) ? 0 : 1); }
// void callUniform3ui(GLint location, GLuint v0, GLuint v1, GLuint v2) {  (*ptrglUniform3ui)(location, v0, v1, v2); }
// int cancallUniform3ui() { return ((ptrglUniform3ui == NULL) ? 0 : 1); }
// void callTexImage2DMultisample(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations) {  (*ptrglTexImage2DMultisample)(target, samples, internalformat, width, height, fixedsamplelocations); }
// int cancallTexImage2DMultisample() { return ((ptrglTexImage2DMultisample == NULL) ? 0 : 1); }
// void callRenderbufferStorageMultisample(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height) {  (*ptrglRenderbufferStorageMultisample)(target, samples, internalformat, width, height); }
// int cancallRenderbufferStorageMultisample() { return ((ptrglRenderbufferStorageMultisample == NULL) ? 0 : 1); }
// GLboolean callIsProgram(GLuint program) { return (*ptrglIsProgram)(program); }
// int cancallIsProgram() { return ((ptrglIsProgram == NULL) ? 0 : 1); }
// void callVertexAttribP2ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {  (*ptrglVertexAttribP2ui)(index, type_, normalized, value); }
// int cancallVertexAttribP2ui() { return ((ptrglVertexAttribP2ui == NULL) ? 0 : 1); }
// void callVertexAttribI4bv(GLuint index, GLbyte* v) {  (*ptrglVertexAttribI4bv)(index, v); }
// int cancallVertexAttribI4bv() { return ((ptrglVertexAttribI4bv == NULL) ? 0 : 1); }
// void callUniform2i(GLint location, GLint v0, GLint v1) {  (*ptrglUniform2i)(location, v0, v1); }
// int cancallUniform2i() { return ((ptrglUniform2i == NULL) ? 0 : 1); }
// void callDeleteRenderbuffers(GLsizei n, GLuint* renderbuffers) {  (*ptrglDeleteRenderbuffers)(n, renderbuffers); }
// int cancallDeleteRenderbuffers() { return ((ptrglDeleteRenderbuffers == NULL) ? 0 : 1); }
// void callGetBufferPointerv(GLenum target, GLenum pname, GLvoid** params) {  (*ptrglGetBufferPointerv)(target, pname, params); }
// int cancallGetBufferPointerv() { return ((ptrglGetBufferPointerv == NULL) ? 0 : 1); }
// void callDrawBuffer(GLenum mode) {  (*ptrglDrawBuffer)(mode); }
// int cancallDrawBuffer() { return ((ptrglDrawBuffer == NULL) ? 0 : 1); }
// GLvoid* callMapBuffer(GLenum target, GLenum access) { return (*ptrglMapBuffer)(target, access); }
// int cancallMapBuffer() { return ((ptrglMapBuffer == NULL) ? 0 : 1); }
// void callDepthRangeArrayv(GLuint first, GLsizei count, GLdouble* v) {  (*ptrglDepthRangeArrayv)(first, count, v); }
// int cancallDepthRangeArrayv() { return ((ptrglDepthRangeArrayv == NULL) ? 0 : 1); }
// void callVertexP3uiv(GLenum type_, GLuint* value) {  (*ptrglVertexP3uiv)(type_, value); }
// int cancallVertexP3uiv() { return ((ptrglVertexP3uiv == NULL) ? 0 : 1); }
// void callFramebufferTexture3D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint zoffset) {  (*ptrglFramebufferTexture3D)(target, attachment, textarget, texture, level, zoffset); }
// int cancallFramebufferTexture3D() { return ((ptrglFramebufferTexture3D == NULL) ? 0 : 1); }
// void callVertexAttribI4ubv(GLuint index, GLubyte* v) {  (*ptrglVertexAttribI4ubv)(index, v); }
// int cancallVertexAttribI4ubv() { return ((ptrglVertexAttribI4ubv == NULL) ? 0 : 1); }
// void callUniform4i(GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {  (*ptrglUniform4i)(location, v0, v1, v2, v3); }
// int cancallUniform4i() { return ((ptrglUniform4i == NULL) ? 0 : 1); }
// void callCompressedTexImage3D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexImage3D)(target, level, internalformat, width, height, depth, border, imageSize, data); }
// int cancallCompressedTexImage3D() { return ((ptrglCompressedTexImage3D == NULL) ? 0 : 1); }
// void callDeleteTransformFeedbacks(GLsizei n, GLuint* ids) {  (*ptrglDeleteTransformFeedbacks)(n, ids); }
// int cancallDeleteTransformFeedbacks() { return ((ptrglDeleteTransformFeedbacks == NULL) ? 0 : 1); }
// void callBeginQuery(GLenum target, GLuint id) {  (*ptrglBeginQuery)(target, id); }
// int cancallBeginQuery() { return ((ptrglBeginQuery == NULL) ? 0 : 1); }
// void callReadBuffer(GLenum mode) {  (*ptrglReadBuffer)(mode); }
// int cancallReadBuffer() { return ((ptrglReadBuffer == NULL) ? 0 : 1); }
// void callDebugMessageControl(GLenum source, GLenum type_, GLenum severity, GLsizei count, GLuint* ids, GLboolean enabled) {  (*ptrglDebugMessageControl)(source, type_, severity, count, ids, enabled); }
// int cancallDebugMessageControl() { return ((ptrglDebugMessageControl == NULL) ? 0 : 1); }
// GLuint callGetSubroutineIndex(GLuint program, GLenum shadertype, GLchar* name) { return (*ptrglGetSubroutineIndex)(program, shadertype, name); }
// int cancallGetSubroutineIndex() { return ((ptrglGetSubroutineIndex == NULL) ? 0 : 1); }
// void callVertexAttribI2uiv(GLuint index, GLuint* v) {  (*ptrglVertexAttribI2uiv)(index, v); }
// int cancallVertexAttribI2uiv() { return ((ptrglVertexAttribI2uiv == NULL) ? 0 : 1); }
// GLvoid* callMapBufferRange(GLenum target, GLintptr offset, GLsizeiptr length, GLbitfield access) { return (*ptrglMapBufferRange)(target, offset, length, access); }
// int cancallMapBufferRange() { return ((ptrglMapBufferRange == NULL) ? 0 : 1); }
// void callDrawBuffers(GLsizei n, GLenum* bufs) {  (*ptrglDrawBuffers)(n, bufs); }
// int cancallDrawBuffers() { return ((ptrglDrawBuffers == NULL) ? 0 : 1); }
// GLboolean callUnmapBuffer(GLenum target) { return (*ptrglUnmapBuffer)(target); }
// int cancallUnmapBuffer() { return ((ptrglUnmapBuffer == NULL) ? 0 : 1); }
// void callUniform3d(GLint location, GLdouble x, GLdouble y, GLdouble z) {  (*ptrglUniform3d)(location, x, y, z); }
// int cancallUniform3d() { return ((ptrglUniform3d == NULL) ? 0 : 1); }
// GLint callGetAttribLocation(GLuint program, GLchar* name) { return (*ptrglGetAttribLocation)(program, name); }
// int cancallGetAttribLocation() { return ((ptrglGetAttribLocation == NULL) ? 0 : 1); }
// void callUseProgram(GLuint program) {  (*ptrglUseProgram)(program); }
// int cancallUseProgram() { return ((ptrglUseProgram == NULL) ? 0 : 1); }
// void callTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexSubImage3D)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type_, pixels); }
// int cancallTexSubImage3D() { return ((ptrglTexSubImage3D == NULL) ? 0 : 1); }
// void callUniformMatrix4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix4fv)(location, count, transpose, value); }
// int cancallUniformMatrix4fv() { return ((ptrglUniformMatrix4fv == NULL) ? 0 : 1); }
// void callProgramUniformMatrix3x4dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix3x4dv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix3x4dv() { return ((ptrglProgramUniformMatrix3x4dv == NULL) ? 0 : 1); }
// void callBlendFuncSeparatei(GLuint buf, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {  (*ptrglBlendFuncSeparatei)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha); }
// int cancallBlendFuncSeparatei() { return ((ptrglBlendFuncSeparatei == NULL) ? 0 : 1); }
// void callVertexAttribP3uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {  (*ptrglVertexAttribP3uiv)(index, type_, normalized, value); }
// int cancallVertexAttribP3uiv() { return ((ptrglVertexAttribP3uiv == NULL) ? 0 : 1); }
// GLboolean callIsRenderbuffer(GLuint renderbuffer) { return (*ptrglIsRenderbuffer)(renderbuffer); }
// int cancallIsRenderbuffer() { return ((ptrglIsRenderbuffer == NULL) ? 0 : 1); }
// GLboolean callIsProgramPipeline(GLuint pipeline) { return (*ptrglIsProgramPipeline)(pipeline); }
// int cancallIsProgramPipeline() { return ((ptrglIsProgramPipeline == NULL) ? 0 : 1); }
// void callScissorIndexed(GLuint index, GLint left, GLint bottom, GLsizei width, GLsizei height) {  (*ptrglScissorIndexed)(index, left, bottom, width, height); }
// int cancallScissorIndexed() { return ((ptrglScissorIndexed == NULL) ? 0 : 1); }
// void callBlendColor(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {  (*ptrglBlendColor)(red, green, blue, alpha); }
// int cancallBlendColor() { return ((ptrglBlendColor == NULL) ? 0 : 1); }
// void callTexParameterf(GLenum target, GLenum pname, GLfloat param) {  (*ptrglTexParameterf)(target, pname, param); }
// int cancallTexParameterf() { return ((ptrglTexParameterf == NULL) ? 0 : 1); }
// void callInvalidateTexSubImage(GLuint texture, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth) {  (*ptrglInvalidateTexSubImage)(texture, level, xoffset, yoffset, zoffset, width, height, depth); }
// int cancallInvalidateTexSubImage() { return ((ptrglInvalidateTexSubImage == NULL) ? 0 : 1); }
// GLint callGetFragDataLocation(GLuint program, GLchar* name) { return (*ptrglGetFragDataLocation)(program, name); }
// int cancallGetFragDataLocation() { return ((ptrglGetFragDataLocation == NULL) ? 0 : 1); }
// void callUniformMatrix2x3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix2x3fv)(location, count, transpose, value); }
// int cancallUniformMatrix2x3fv() { return ((ptrglUniformMatrix2x3fv == NULL) ? 0 : 1); }
// void callUniform4d(GLint location, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {  (*ptrglUniform4d)(location, x, y, z, w); }
// int cancallUniform4d() { return ((ptrglUniform4d == NULL) ? 0 : 1); }
// void callClearBufferfv(GLenum buffer, GLint drawbuffer, GLfloat* value) {  (*ptrglClearBufferfv)(buffer, drawbuffer, value); }
// int cancallClearBufferfv() { return ((ptrglClearBufferfv == NULL) ? 0 : 1); }
// void callGetQueryIndexediv(GLenum target, GLuint index, GLenum pname, GLint* params) {  (*ptrglGetQueryIndexediv)(target, index, pname, params); }
// int cancallGetQueryIndexediv() { return ((ptrglGetQueryIndexediv == NULL) ? 0 : 1); }
// void callProgramUniform3fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {  (*ptrglProgramUniform3fv)(program, location, count, value); }
// int cancallProgramUniform3fv() { return ((ptrglProgramUniform3fv == NULL) ? 0 : 1); }
// void callProgramUniform1iv(GLuint program, GLint location, GLsizei count, GLint* value) {  (*ptrglProgramUniform1iv)(program, location, count, value); }
// int cancallProgramUniform1iv() { return ((ptrglProgramUniform1iv == NULL) ? 0 : 1); }
// void callUniform2fv(GLint location, GLsizei count, GLfloat* value) {  (*ptrglUniform2fv)(location, count, value); }
// int cancallUniform2fv() { return ((ptrglUniform2fv == NULL) ? 0 : 1); }
// void callGetProgramInterfaceiv(GLuint program, GLenum programInterface, GLenum pname, GLint* params) {  (*ptrglGetProgramInterfaceiv)(program, programInterface, pname, params); }
// int cancallGetProgramInterfaceiv() { return ((ptrglGetProgramInterfaceiv == NULL) ? 0 : 1); }
// void callUniformMatrix2dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix2dv)(location, count, transpose, value); }
// int cancallUniformMatrix2dv() { return ((ptrglUniformMatrix2dv == NULL) ? 0 : 1); }
// GLboolean callIsSync(GLsync sync) { return (*ptrglIsSync)(sync); }
// int cancallIsSync() { return ((ptrglIsSync == NULL) ? 0 : 1); }
// void callGetBooleanv(GLenum pname, GLboolean* params) {  (*ptrglGetBooleanv)(pname, params); }
// int cancallGetBooleanv() { return ((ptrglGetBooleanv == NULL) ? 0 : 1); }
// void callStencilOp(GLenum fail, GLenum zfail, GLenum zpass) {  (*ptrglStencilOp)(fail, zfail, zpass); }
// int cancallStencilOp() { return ((ptrglStencilOp == NULL) ? 0 : 1); }
// void callGetQueryObjectui64v(GLuint id, GLenum pname, GLuint64* params) {  (*ptrglGetQueryObjectui64v)(id, pname, params); }
// int cancallGetQueryObjectui64v() { return ((ptrglGetQueryObjectui64v == NULL) ? 0 : 1); }
// void callProgramUniformMatrix4x3dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix4x3dv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix4x3dv() { return ((ptrglProgramUniformMatrix4x3dv == NULL) ? 0 : 1); }
// void callGetObjectPtrLabel(void* ptr, GLsizei bufSize, GLsizei* length, GLchar* label) {  (*ptrglGetObjectPtrLabel)(ptr, bufSize, length, label); }
// int cancallGetObjectPtrLabel() { return ((ptrglGetObjectPtrLabel == NULL) ? 0 : 1); }
// void callClearBufferfi(GLenum buffer, GLint drawbuffer, GLfloat depth, GLint stencil) {  (*ptrglClearBufferfi)(buffer, drawbuffer, depth, stencil); }
// int cancallClearBufferfi() { return ((ptrglClearBufferfi == NULL) ? 0 : 1); }
// void callProgramUniformMatrix2x4fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix2x4fv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix2x4fv() { return ((ptrglProgramUniformMatrix2x4fv == NULL) ? 0 : 1); }
// void callUniform3fv(GLint location, GLsizei count, GLfloat* value) {  (*ptrglUniform3fv)(location, count, value); }
// int cancallUniform3fv() { return ((ptrglUniform3fv == NULL) ? 0 : 1); }
// void callVertexP4ui(GLenum type_, GLuint value) {  (*ptrglVertexP4ui)(type_, value); }
// int cancallVertexP4ui() { return ((ptrglVertexP4ui == NULL) ? 0 : 1); }
// void callGetQueryObjecti64v(GLuint id, GLenum pname, GLint64* params) {  (*ptrglGetQueryObjecti64v)(id, pname, params); }
// int cancallGetQueryObjecti64v() { return ((ptrglGetQueryObjecti64v == NULL) ? 0 : 1); }
// void callCompressedTexImage1D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexImage1D)(target, level, internalformat, width, border, imageSize, data); }
// int cancallCompressedTexImage1D() { return ((ptrglCompressedTexImage1D == NULL) ? 0 : 1); }
// void callPrimitiveRestartIndex(GLuint index) {  (*ptrglPrimitiveRestartIndex)(index); }
// int cancallPrimitiveRestartIndex() { return ((ptrglPrimitiveRestartIndex == NULL) ? 0 : 1); }
// void callDrawTransformFeedback(GLenum mode, GLuint id) {  (*ptrglDrawTransformFeedback)(mode, id); }
// int cancallDrawTransformFeedback() { return ((ptrglDrawTransformFeedback == NULL) ? 0 : 1); }
// void callGetDoublev(GLenum pname, GLdouble* params) {  (*ptrglGetDoublev)(pname, params); }
// int cancallGetDoublev() { return ((ptrglGetDoublev == NULL) ? 0 : 1); }
// void callBlendFuncSeparate(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha) {  (*ptrglBlendFuncSeparate)(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha); }
// int cancallBlendFuncSeparate() { return ((ptrglBlendFuncSeparate == NULL) ? 0 : 1); }
// void callVertexAttribP1ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {  (*ptrglVertexAttribP1ui)(index, type_, normalized, value); }
// int cancallVertexAttribP1ui() { return ((ptrglVertexAttribP1ui == NULL) ? 0 : 1); }
// void callVertexAttribP4uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {  (*ptrglVertexAttribP4uiv)(index, type_, normalized, value); }
// int cancallVertexAttribP4uiv() { return ((ptrglVertexAttribP4uiv == NULL) ? 0 : 1); }
// void callBeginTransformFeedback(GLenum primitiveMode) {  (*ptrglBeginTransformFeedback)(primitiveMode); }
// int cancallBeginTransformFeedback() { return ((ptrglBeginTransformFeedback == NULL) ? 0 : 1); }
// void callClearStencil(GLint s) {  (*ptrglClearStencil)(s); }
// int cancallClearStencil() { return ((ptrglClearStencil == NULL) ? 0 : 1); }
// void callDeleteFramebuffers(GLsizei n, GLuint* framebuffers) {  (*ptrglDeleteFramebuffers)(n, framebuffers); }
// int cancallDeleteFramebuffers() { return ((ptrglDeleteFramebuffers == NULL) ? 0 : 1); }
// void callVertexAttribL1dv(GLuint index, GLdouble* v) {  (*ptrglVertexAttribL1dv)(index, v); }
// int cancallVertexAttribL1dv() { return ((ptrglVertexAttribL1dv == NULL) ? 0 : 1); }
// void callPointParameterf(GLenum pname, GLfloat param) {  (*ptrglPointParameterf)(pname, param); }
// int cancallPointParameterf() { return ((ptrglPointParameterf == NULL) ? 0 : 1); }
// void callVertexAttribI4sv(GLuint index, GLshort* v) {  (*ptrglVertexAttribI4sv)(index, v); }
// int cancallVertexAttribI4sv() { return ((ptrglVertexAttribI4sv == NULL) ? 0 : 1); }
// void callVertexAttribL2dv(GLuint index, GLdouble* v) {  (*ptrglVertexAttribL2dv)(index, v); }
// int cancallVertexAttribL2dv() { return ((ptrglVertexAttribL2dv == NULL) ? 0 : 1); }
// void callBlendFunci(GLuint buf, GLenum src, GLenum dst) {  (*ptrglBlendFunci)(buf, src, dst); }
// int cancallBlendFunci() { return ((ptrglBlendFunci == NULL) ? 0 : 1); }
// void callUniform2ui(GLint location, GLuint v0, GLuint v1) {  (*ptrglUniform2ui)(location, v0, v1); }
// int cancallUniform2ui() { return ((ptrglUniform2ui == NULL) ? 0 : 1); }
// void callGetProgramPipelineiv(GLuint pipeline, GLenum pname, GLint* params) {  (*ptrglGetProgramPipelineiv)(pipeline, pname, params); }
// int cancallGetProgramPipelineiv() { return ((ptrglGetProgramPipelineiv == NULL) ? 0 : 1); }
// void callSampleMaski(GLuint index, GLbitfield mask) {  (*ptrglSampleMaski)(index, mask); }
// int cancallSampleMaski() { return ((ptrglSampleMaski == NULL) ? 0 : 1); }
// void callGetSamplerParameterIuiv(GLuint sampler, GLenum pname, GLuint* params) {  (*ptrglGetSamplerParameterIuiv)(sampler, pname, params); }
// int cancallGetSamplerParameterIuiv() { return ((ptrglGetSamplerParameterIuiv == NULL) ? 0 : 1); }
// void callCopyBufferSubData(GLenum readTarget, GLenum writeTarget, GLintptr readOffset, GLintptr writeOffset, GLsizeiptr size) {  (*ptrglCopyBufferSubData)(readTarget, writeTarget, readOffset, writeOffset, size); }
// int cancallCopyBufferSubData() { return ((ptrglCopyBufferSubData == NULL) ? 0 : 1); }
// void callGetActiveUniformsiv(GLuint program, GLsizei uniformCount, GLuint* uniformIndices, GLenum pname, GLint* params) {  (*ptrglGetActiveUniformsiv)(program, uniformCount, uniformIndices, pname, params); }
// int cancallGetActiveUniformsiv() { return ((ptrglGetActiveUniformsiv == NULL) ? 0 : 1); }
// void callGetBufferParameteriv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetBufferParameteriv)(target, pname, params); }
// int cancallGetBufferParameteriv() { return ((ptrglGetBufferParameteriv == NULL) ? 0 : 1); }
// void callStencilOpSeparate(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {  (*ptrglStencilOpSeparate)(face, sfail, dpfail, dppass); }
// int cancallStencilOpSeparate() { return ((ptrglStencilOpSeparate == NULL) ? 0 : 1); }
// void callBlitFramebuffer(GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter) {  (*ptrglBlitFramebuffer)(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter); }
// int cancallBlitFramebuffer() { return ((ptrglBlitFramebuffer == NULL) ? 0 : 1); }
// void callDepthFunc(GLenum func_) {  (*ptrglDepthFunc)(func_); }
// int cancallDepthFunc() { return ((ptrglDepthFunc == NULL) ? 0 : 1); }
// void callDetachShader(GLuint program, GLuint shader) {  (*ptrglDetachShader)(program, shader); }
// int cancallDetachShader() { return ((ptrglDetachShader == NULL) ? 0 : 1); }
// void callClearBufferuiv(GLenum buffer, GLint drawbuffer, GLuint* value) {  (*ptrglClearBufferuiv)(buffer, drawbuffer, value); }
// int cancallClearBufferuiv() { return ((ptrglClearBufferuiv == NULL) ? 0 : 1); }
// void callProgramBinary(GLuint program, GLenum binaryFormat, GLvoid* binary, GLsizei length) {  (*ptrglProgramBinary)(program, binaryFormat, binary, length); }
// int cancallProgramBinary() { return ((ptrglProgramBinary == NULL) ? 0 : 1); }
// void callViewportIndexedf(GLuint index, GLfloat x, GLfloat y, GLfloat w, GLfloat h) {  (*ptrglViewportIndexedf)(index, x, y, w, h); }
// int cancallViewportIndexedf() { return ((ptrglViewportIndexedf == NULL) ? 0 : 1); }
// void callProgramUniformMatrix3dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix3dv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix3dv() { return ((ptrglProgramUniformMatrix3dv == NULL) ? 0 : 1); }
// void callUniformMatrix2x4dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix2x4dv)(location, count, transpose, value); }
// int cancallUniformMatrix2x4dv() { return ((ptrglUniformMatrix2x4dv == NULL) ? 0 : 1); }
// void callProgramParameteri(GLuint program, GLenum pname, GLint value) {  (*ptrglProgramParameteri)(program, pname, value); }
// int cancallProgramParameteri() { return ((ptrglProgramParameteri == NULL) ? 0 : 1); }
// void callMultiTexCoordP3ui(GLenum texture, GLenum type_, GLuint coords) {  (*ptrglMultiTexCoordP3ui)(texture, type_, coords); }
// int cancallMultiTexCoordP3ui() { return ((ptrglMultiTexCoordP3ui == NULL) ? 0 : 1); }
// void callUniformMatrix4dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix4dv)(location, count, transpose, value); }
// int cancallUniformMatrix4dv() { return ((ptrglUniformMatrix4dv == NULL) ? 0 : 1); }
// void callBindFragDataLocation(GLuint program, GLuint color, GLchar* name) {  (*ptrglBindFragDataLocation)(program, color, name); }
// int cancallBindFragDataLocation() { return ((ptrglBindFragDataLocation == NULL) ? 0 : 1); }
// void callGetAttachedShaders(GLuint program, GLsizei maxCount, GLsizei* count, GLuint* obj) {  (*ptrglGetAttachedShaders)(program, maxCount, count, obj); }
// int cancallGetAttachedShaders() { return ((ptrglGetAttachedShaders == NULL) ? 0 : 1); }
// void callVertexP4uiv(GLenum type_, GLuint* value) {  (*ptrglVertexP4uiv)(type_, value); }
// int cancallVertexP4uiv() { return ((ptrglVertexP4uiv == NULL) ? 0 : 1); }
// void callReadPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglReadPixels)(x, y, width, height, format, type_, pixels); }
// int cancallReadPixels() { return ((ptrglReadPixels == NULL) ? 0 : 1); }
// void callVertexAttribIFormat(GLuint attribindex, GLint size, GLenum type_, GLuint relativeoffset) {  (*ptrglVertexAttribIFormat)(attribindex, size, type_, relativeoffset); }
// int cancallVertexAttribIFormat() { return ((ptrglVertexAttribIFormat == NULL) ? 0 : 1); }
// void callTexStorage1D(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width) {  (*ptrglTexStorage1D)(target, levels, internalformat, width); }
// int cancallTexStorage1D() { return ((ptrglTexStorage1D == NULL) ? 0 : 1); }
// void callProgramUniform4dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {  (*ptrglProgramUniform4dv)(program, location, count, value); }
// int cancallProgramUniform4dv() { return ((ptrglProgramUniform4dv == NULL) ? 0 : 1); }
// void callVertexBindingDivisor(GLuint bindingindex, GLuint divisor) {  (*ptrglVertexBindingDivisor)(bindingindex, divisor); }
// int cancallVertexBindingDivisor() { return ((ptrglVertexBindingDivisor == NULL) ? 0 : 1); }
// void callDrawRangeElementsBaseVertex(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type_, GLvoid* indices, GLint basevertex) {  (*ptrglDrawRangeElementsBaseVertex)(mode, start, end, count, type_, indices, basevertex); }
// int cancallDrawRangeElementsBaseVertex() { return ((ptrglDrawRangeElementsBaseVertex == NULL) ? 0 : 1); }
// void callGetUniformiv(GLuint program, GLint location, GLint* params) {  (*ptrglGetUniformiv)(program, location, params); }
// int cancallGetUniformiv() { return ((ptrglGetUniformiv == NULL) ? 0 : 1); }
// void callVertexAttribL3dv(GLuint index, GLdouble* v) {  (*ptrglVertexAttribL3dv)(index, v); }
// int cancallVertexAttribL3dv() { return ((ptrglVertexAttribL3dv == NULL) ? 0 : 1); }
// void callInvalidateBufferData(GLuint buffer) {  (*ptrglInvalidateBufferData)(buffer); }
// int cancallInvalidateBufferData() { return ((ptrglInvalidateBufferData == NULL) ? 0 : 1); }
// void callResumeTransformFeedback() {  (*ptrglResumeTransformFeedback)(); }
// int cancallResumeTransformFeedback() { return ((ptrglResumeTransformFeedback == NULL) ? 0 : 1); }
// void callUniformMatrix2x3dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix2x3dv)(location, count, transpose, value); }
// int cancallUniformMatrix2x3dv() { return ((ptrglUniformMatrix2x3dv == NULL) ? 0 : 1); }
// void callEndQueryIndexed(GLenum target, GLuint index) {  (*ptrglEndQueryIndexed)(target, index); }
// int cancallEndQueryIndexed() { return ((ptrglEndQueryIndexed == NULL) ? 0 : 1); }
// void callGetVertexAttribIuiv(GLuint index, GLenum pname, GLuint* params) {  (*ptrglGetVertexAttribIuiv)(index, pname, params); }
// int cancallGetVertexAttribIuiv() { return ((ptrglGetVertexAttribIuiv == NULL) ? 0 : 1); }
// void callBeginQueryIndexed(GLenum target, GLuint index, GLuint id) {  (*ptrglBeginQueryIndexed)(target, index, id); }
// int cancallBeginQueryIndexed() { return ((ptrglBeginQueryIndexed == NULL) ? 0 : 1); }
// void callUniformMatrix4x3dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix4x3dv)(location, count, transpose, value); }
// int cancallUniformMatrix4x3dv() { return ((ptrglUniformMatrix4x3dv == NULL) ? 0 : 1); }
// void callTexBufferRange(GLenum target, GLenum internalformat, GLuint buffer, GLintptr offset, GLsizeiptr size) {  (*ptrglTexBufferRange)(target, internalformat, buffer, offset, size); }
// int cancallTexBufferRange() { return ((ptrglTexBufferRange == NULL) ? 0 : 1); }
// void callGetVertexAttribLdv(GLuint index, GLenum pname, GLdouble* params) {  (*ptrglGetVertexAttribLdv)(index, pname, params); }
// int cancallGetVertexAttribLdv() { return ((ptrglGetVertexAttribLdv == NULL) ? 0 : 1); }
// void callDrawTransformFeedbackStream(GLenum mode, GLuint id, GLuint stream) {  (*ptrglDrawTransformFeedbackStream)(mode, id, stream); }
// int cancallDrawTransformFeedbackStream() { return ((ptrglDrawTransformFeedbackStream == NULL) ? 0 : 1); }
// void callBindFramebuffer(GLenum target, GLuint framebuffer) {  (*ptrglBindFramebuffer)(target, framebuffer); }
// int cancallBindFramebuffer() { return ((ptrglBindFramebuffer == NULL) ? 0 : 1); }
// void callDeleteProgram(GLuint program) {  (*ptrglDeleteProgram)(program); }
// int cancallDeleteProgram() { return ((ptrglDeleteProgram == NULL) ? 0 : 1); }
// void callScissorIndexedv(GLuint index, GLint* v) {  (*ptrglScissorIndexedv)(index, v); }
// int cancallScissorIndexedv() { return ((ptrglScissorIndexedv == NULL) ? 0 : 1); }
// void callUniformMatrix3x2dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix3x2dv)(location, count, transpose, value); }
// int cancallUniformMatrix3x2dv() { return ((ptrglUniformMatrix3x2dv == NULL) ? 0 : 1); }
// GLboolean callIsTexture(GLuint texture) { return (*ptrglIsTexture)(texture); }
// int cancallIsTexture() { return ((ptrglIsTexture == NULL) ? 0 : 1); }
// void callSamplerParameterIiv(GLuint sampler, GLenum pname, GLint* param) {  (*ptrglSamplerParameterIiv)(sampler, pname, param); }
// int cancallSamplerParameterIiv() { return ((ptrglSamplerParameterIiv == NULL) ? 0 : 1); }
// void callProgramUniform4fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {  (*ptrglProgramUniform4fv)(program, location, count, value); }
// int cancallProgramUniform4fv() { return ((ptrglProgramUniform4fv == NULL) ? 0 : 1); }
// void callBlendEquationi(GLuint buf, GLenum mode) {  (*ptrglBlendEquationi)(buf, mode); }
// int cancallBlendEquationi() { return ((ptrglBlendEquationi == NULL) ? 0 : 1); }
// void callDeleteSync(GLsync sync) {  (*ptrglDeleteSync)(sync); }
// int cancallDeleteSync() { return ((ptrglDeleteSync == NULL) ? 0 : 1); }
// void callProgramUniform3iv(GLuint program, GLint location, GLsizei count, GLint* value) {  (*ptrglProgramUniform3iv)(program, location, count, value); }
// int cancallProgramUniform3iv() { return ((ptrglProgramUniform3iv == NULL) ? 0 : 1); }
// void callInvalidateTexImage(GLuint texture, GLint level) {  (*ptrglInvalidateTexImage)(texture, level); }
// int cancallInvalidateTexImage() { return ((ptrglInvalidateTexImage == NULL) ? 0 : 1); }
// void callViewportArrayv(GLuint first, GLsizei count, GLfloat* v) {  (*ptrglViewportArrayv)(first, count, v); }
// int cancallViewportArrayv() { return ((ptrglViewportArrayv == NULL) ? 0 : 1); }
// void callTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexSubImage1D)(target, level, xoffset, width, format, type_, pixels); }
// int cancallTexSubImage1D() { return ((ptrglTexSubImage1D == NULL) ? 0 : 1); }
// GLuint callCreateShaderProgramv(GLenum type_, GLsizei count, GLchar* const* strings) { return (*ptrglCreateShaderProgramv)(type_, count, strings); }
// int cancallCreateShaderProgramv() { return ((ptrglCreateShaderProgramv == NULL) ? 0 : 1); }
// void callGenRenderbuffers(GLsizei n, GLuint* renderbuffers) {  (*ptrglGenRenderbuffers)(n, renderbuffers); }
// int cancallGenRenderbuffers() { return ((ptrglGenRenderbuffers == NULL) ? 0 : 1); }
// void callProgramUniformMatrix2dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix2dv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix2dv() { return ((ptrglProgramUniformMatrix2dv == NULL) ? 0 : 1); }
// void callProgramUniform2ui(GLuint program, GLint location, GLuint v0, GLuint v1) {  (*ptrglProgramUniform2ui)(program, location, v0, v1); }
// int cancallProgramUniform2ui() { return ((ptrglProgramUniform2ui == NULL) ? 0 : 1); }
// void callCopyTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {  (*ptrglCopyTexSubImage2D)(target, level, xoffset, yoffset, x, y, width, height); }
// int cancallCopyTexSubImage2D() { return ((ptrglCopyTexSubImage2D == NULL) ? 0 : 1); }
// void callMultiDrawElementsBaseVertex(GLenum mode, GLsizei* count, GLenum type_, GLvoid* const* indices, GLsizei drawcount, GLint* basevertex) {  (*ptrglMultiDrawElementsBaseVertex)(mode, count, type_, indices, drawcount, basevertex); }
// int cancallMultiDrawElementsBaseVertex() { return ((ptrglMultiDrawElementsBaseVertex == NULL) ? 0 : 1); }
// void callGetCompressedTexImage(GLenum target, GLint level, GLvoid* img) {  (*ptrglGetCompressedTexImage)(target, level, img); }
// int cancallGetCompressedTexImage() { return ((ptrglGetCompressedTexImage == NULL) ? 0 : 1); }
// void callVertexAttribI3ui(GLuint index, GLuint x, GLuint y, GLuint z) {  (*ptrglVertexAttribI3ui)(index, x, y, z); }
// int cancallVertexAttribI3ui() { return ((ptrglVertexAttribI3ui == NULL) ? 0 : 1); }
// void callUniformMatrix3x4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix3x4fv)(location, count, transpose, value); }
// int cancallUniformMatrix3x4fv() { return ((ptrglUniformMatrix3x4fv == NULL) ? 0 : 1); }
// void callVertexAttribI1iv(GLuint index, GLint* v) {  (*ptrglVertexAttribI1iv)(index, v); }
// int cancallVertexAttribI1iv() { return ((ptrglVertexAttribI1iv == NULL) ? 0 : 1); }
// void callGetBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {  (*ptrglGetBufferSubData)(target, offset, size, data); }
// int cancallGetBufferSubData() { return ((ptrglGetBufferSubData == NULL) ? 0 : 1); }
// void callTexImage1D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexImage1D)(target, level, internalformat, width, border, format, type_, pixels); }
// int cancallTexImage1D() { return ((ptrglTexImage1D == NULL) ? 0 : 1); }
// void callFramebufferTexture(GLenum target, GLenum attachment, GLuint texture, GLint level) {  (*ptrglFramebufferTexture)(target, attachment, texture, level); }
// int cancallFramebufferTexture() { return ((ptrglFramebufferTexture == NULL) ? 0 : 1); }
// void callVertexAttribI1i(GLuint index, GLint x) {  (*ptrglVertexAttribI1i)(index, x); }
// int cancallVertexAttribI1i() { return ((ptrglVertexAttribI1i == NULL) ? 0 : 1); }
// void callGetVertexAttribIiv(GLuint index, GLenum pname, GLint* params) {  (*ptrglGetVertexAttribIiv)(index, pname, params); }
// int cancallGetVertexAttribIiv() { return ((ptrglGetVertexAttribIiv == NULL) ? 0 : 1); }
// void callGetBufferParameteri64v(GLenum target, GLenum pname, GLint64* params) {  (*ptrglGetBufferParameteri64v)(target, pname, params); }
// int cancallGetBufferParameteri64v() { return ((ptrglGetBufferParameteri64v == NULL) ? 0 : 1); }
// void callBeginConditionalRender(GLuint id, GLenum mode) {  (*ptrglBeginConditionalRender)(id, mode); }
// int cancallBeginConditionalRender() { return ((ptrglBeginConditionalRender == NULL) ? 0 : 1); }
// void callGetInternalformati64v(GLenum target, GLenum internalformat, GLenum pname, GLsizei bufSize, GLint64* params) {  (*ptrglGetInternalformati64v)(target, internalformat, pname, bufSize, params); }
// int cancallGetInternalformati64v() { return ((ptrglGetInternalformati64v == NULL) ? 0 : 1); }
// void callVertexAttribI2ui(GLuint index, GLuint x, GLuint y) {  (*ptrglVertexAttribI2ui)(index, x, y); }
// int cancallVertexAttribI2ui() { return ((ptrglVertexAttribI2ui == NULL) ? 0 : 1); }
// void callBlendEquation(GLenum mode) {  (*ptrglBlendEquation)(mode); }
// int cancallBlendEquation() { return ((ptrglBlendEquation == NULL) ? 0 : 1); }
// void callEnablei(GLenum target, GLuint index) {  (*ptrglEnablei)(target, index); }
// int cancallEnablei() { return ((ptrglEnablei == NULL) ? 0 : 1); }
// const GLubyte * callGetStringi(GLenum name, GLuint index) { return (*ptrglGetStringi)(name, index); }
// int cancallGetStringi() { return ((ptrglGetStringi == NULL) ? 0 : 1); }
// void callProgramUniformMatrix3x2dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix3x2dv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix3x2dv() { return ((ptrglProgramUniformMatrix3x2dv == NULL) ? 0 : 1); }
// void callDrawTransformFeedbackStreamInstanced(GLenum mode, GLuint id, GLuint stream, GLsizei instancecount) {  (*ptrglDrawTransformFeedbackStreamInstanced)(mode, id, stream, instancecount); }
// int cancallDrawTransformFeedbackStreamInstanced() { return ((ptrglDrawTransformFeedbackStreamInstanced == NULL) ? 0 : 1); }
// void callGetActiveSubroutineUniformName(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name) {  (*ptrglGetActiveSubroutineUniformName)(program, shadertype, index, bufsize, length, name); }
// int cancallGetActiveSubroutineUniformName() { return ((ptrglGetActiveSubroutineUniformName == NULL) ? 0 : 1); }
// void callGetInteger64i_v(GLenum target, GLuint index, GLint64* data) {  (*ptrglGetInteger64i_v)(target, index, data); }
// int cancallGetInteger64i_v() { return ((ptrglGetInteger64i_v == NULL) ? 0 : 1); }
// void callDebugMessageInsert(GLenum source, GLenum type_, GLuint id, GLenum severity, GLsizei length, GLchar* buf) {  (*ptrglDebugMessageInsert)(source, type_, id, severity, length, buf); }
// int cancallDebugMessageInsert() { return ((ptrglDebugMessageInsert == NULL) ? 0 : 1); }
// void callProgramUniform1f(GLuint program, GLint location, GLfloat v0) {  (*ptrglProgramUniform1f)(program, location, v0); }
// int cancallProgramUniform1f() { return ((ptrglProgramUniform1f == NULL) ? 0 : 1); }
// void callGetSamplerParameterfv(GLuint sampler, GLenum pname, GLfloat* params) {  (*ptrglGetSamplerParameterfv)(sampler, pname, params); }
// int cancallGetSamplerParameterfv() { return ((ptrglGetSamplerParameterfv == NULL) ? 0 : 1); }
// void callTexCoordP4ui(GLenum type_, GLuint coords) {  (*ptrglTexCoordP4ui)(type_, coords); }
// int cancallTexCoordP4ui() { return ((ptrglTexCoordP4ui == NULL) ? 0 : 1); }
// void callProgramUniform1uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {  (*ptrglProgramUniform1uiv)(program, location, count, value); }
// int cancallProgramUniform1uiv() { return ((ptrglProgramUniform1uiv == NULL) ? 0 : 1); }
// void callDeleteSamplers(GLsizei count, GLuint* samplers) {  (*ptrglDeleteSamplers)(count, samplers); }
// int cancallDeleteSamplers() { return ((ptrglDeleteSamplers == NULL) ? 0 : 1); }
// GLboolean callIsShader(GLuint shader) { return (*ptrglIsShader)(shader); }
// int cancallIsShader() { return ((ptrglIsShader == NULL) ? 0 : 1); }
// void callCopyTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {  (*ptrglCopyTexSubImage3D)(target, level, xoffset, yoffset, zoffset, x, y, width, height); }
// int cancallCopyTexSubImage3D() { return ((ptrglCopyTexSubImage3D == NULL) ? 0 : 1); }
// void callGetUniformuiv(GLuint program, GLint location, GLuint* params) {  (*ptrglGetUniformuiv)(program, location, params); }
// int cancallGetUniformuiv() { return ((ptrglGetUniformuiv == NULL) ? 0 : 1); }
// void callVertexAttribL2d(GLuint index, GLdouble x, GLdouble y) {  (*ptrglVertexAttribL2d)(index, x, y); }
// int cancallVertexAttribL2d() { return ((ptrglVertexAttribL2d == NULL) ? 0 : 1); }
// void callSecondaryColorP3ui(GLenum type_, GLuint color) {  (*ptrglSecondaryColorP3ui)(type_, color); }
// int cancallSecondaryColorP3ui() { return ((ptrglSecondaryColorP3ui == NULL) ? 0 : 1); }
// GLboolean callIsSampler(GLuint sampler) { return (*ptrglIsSampler)(sampler); }
// int cancallIsSampler() { return ((ptrglIsSampler == NULL) ? 0 : 1); }
// void callUniform2dv(GLint location, GLsizei count, GLdouble* value) {  (*ptrglUniform2dv)(location, count, value); }
// int cancallUniform2dv() { return ((ptrglUniform2dv == NULL) ? 0 : 1); }
// void callDepthMask(GLboolean flag) {  (*ptrglDepthMask)(flag); }
// int cancallDepthMask() { return ((ptrglDepthMask == NULL) ? 0 : 1); }
// void callPushDebugGroup(GLenum source, GLuint id, GLsizei length, GLchar* message) {  (*ptrglPushDebugGroup)(source, id, length, message); }
// int cancallPushDebugGroup() { return ((ptrglPushDebugGroup == NULL) ? 0 : 1); }
// void callProgramUniform1dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {  (*ptrglProgramUniform1dv)(program, location, count, value); }
// int cancallProgramUniform1dv() { return ((ptrglProgramUniform1dv == NULL) ? 0 : 1); }
// void callMultiDrawArraysIndirect(GLenum mode, void* indirect, GLsizei drawcount, GLsizei stride) {  (*ptrglMultiDrawArraysIndirect)(mode, indirect, drawcount, stride); }
// int cancallMultiDrawArraysIndirect() { return ((ptrglMultiDrawArraysIndirect == NULL) ? 0 : 1); }
// void callMultiTexCoordP2uiv(GLenum texture, GLenum type_, GLuint* coords) {  (*ptrglMultiTexCoordP2uiv)(texture, type_, coords); }
// int cancallMultiTexCoordP2uiv() { return ((ptrglMultiTexCoordP2uiv == NULL) ? 0 : 1); }
// void callInvalidateSubFramebuffer(GLenum target, GLsizei numAttachments, GLenum* attachments, GLint x, GLint y, GLsizei width, GLsizei height) {  (*ptrglInvalidateSubFramebuffer)(target, numAttachments, attachments, x, y, width, height); }
// int cancallInvalidateSubFramebuffer() { return ((ptrglInvalidateSubFramebuffer == NULL) ? 0 : 1); }
// void callUniform3i(GLint location, GLint v0, GLint v1, GLint v2) {  (*ptrglUniform3i)(location, v0, v1, v2); }
// int cancallUniform3i() { return ((ptrglUniform3i == NULL) ? 0 : 1); }
// void callProgramUniform2i(GLuint program, GLint location, GLint v0, GLint v1) {  (*ptrglProgramUniform2i)(program, location, v0, v1); }
// int cancallProgramUniform2i() { return ((ptrglProgramUniform2i == NULL) ? 0 : 1); }
// void callShaderSource(GLuint shader, GLsizei count, GLchar* const* string_, GLint* length) {  (*ptrglShaderSource)(shader, count, string_, length); }
// int cancallShaderSource() { return ((ptrglShaderSource == NULL) ? 0 : 1); }
// void callTexCoordP1uiv(GLenum type_, GLuint* coords) {  (*ptrglTexCoordP1uiv)(type_, coords); }
// int cancallTexCoordP1uiv() { return ((ptrglTexCoordP1uiv == NULL) ? 0 : 1); }
// void callUniformBlockBinding(GLuint program, GLuint uniformBlockIndex, GLuint uniformBlockBinding) {  (*ptrglUniformBlockBinding)(program, uniformBlockIndex, uniformBlockBinding); }
// int cancallUniformBlockBinding() { return ((ptrglUniformBlockBinding == NULL) ? 0 : 1); }
// void callUniform4fv(GLint location, GLsizei count, GLfloat* value) {  (*ptrglUniform4fv)(location, count, value); }
// int cancallUniform4fv() { return ((ptrglUniform4fv == NULL) ? 0 : 1); }
// void callGetVertexAttribdv(GLuint index, GLenum pname, GLdouble* params) {  (*ptrglGetVertexAttribdv)(index, pname, params); }
// int cancallGetVertexAttribdv() { return ((ptrglGetVertexAttribdv == NULL) ? 0 : 1); }
// void callGenBuffers(GLsizei n, GLuint* buffers) {  (*ptrglGenBuffers)(n, buffers); }
// int cancallGenBuffers() { return ((ptrglGenBuffers == NULL) ? 0 : 1); }
// void callVertexAttribP4ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {  (*ptrglVertexAttribP4ui)(index, type_, normalized, value); }
// int cancallVertexAttribP4ui() { return ((ptrglVertexAttribP4ui == NULL) ? 0 : 1); }
// void callGetInternalformativ(GLenum target, GLenum internalformat, GLenum pname, GLsizei bufSize, GLint* params) {  (*ptrglGetInternalformativ)(target, internalformat, pname, bufSize, params); }
// int cancallGetInternalformativ() { return ((ptrglGetInternalformativ == NULL) ? 0 : 1); }
// void callProgramUniformMatrix4x2dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix4x2dv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix4x2dv() { return ((ptrglProgramUniformMatrix4x2dv == NULL) ? 0 : 1); }
// void callCompileShader(GLuint shader) {  (*ptrglCompileShader)(shader); }
// int cancallCompileShader() { return ((ptrglCompileShader == NULL) ? 0 : 1); }
// void callBindRenderbuffer(GLenum target, GLuint renderbuffer) {  (*ptrglBindRenderbuffer)(target, renderbuffer); }
// int cancallBindRenderbuffer() { return ((ptrglBindRenderbuffer == NULL) ? 0 : 1); }
// void callCompressedTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexSubImage1D)(target, level, xoffset, width, format, imageSize, data); }
// int cancallCompressedTexSubImage1D() { return ((ptrglCompressedTexSubImage1D == NULL) ? 0 : 1); }
// void callBlendFunc(GLenum sfactor, GLenum dfactor) {  (*ptrglBlendFunc)(sfactor, dfactor); }
// int cancallBlendFunc() { return ((ptrglBlendFunc == NULL) ? 0 : 1); }
// void callTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {  (*ptrglTexParameterfv)(target, pname, params); }
// int cancallTexParameterfv() { return ((ptrglTexParameterfv == NULL) ? 0 : 1); }
// void callInvalidateFramebuffer(GLenum target, GLsizei numAttachments, GLenum* attachments) {  (*ptrglInvalidateFramebuffer)(target, numAttachments, attachments); }
// int cancallInvalidateFramebuffer() { return ((ptrglInvalidateFramebuffer == NULL) ? 0 : 1); }
// void callBindTransformFeedback(GLenum target, GLuint id) {  (*ptrglBindTransformFeedback)(target, id); }
// int cancallBindTransformFeedback() { return ((ptrglBindTransformFeedback == NULL) ? 0 : 1); }
// GLint callGetFragDataIndex(GLuint program, GLchar* name) { return (*ptrglGetFragDataIndex)(program, name); }
// int cancallGetFragDataIndex() { return ((ptrglGetFragDataIndex == NULL) ? 0 : 1); }
// void callTransformFeedbackVaryings(GLuint program, GLsizei count, GLchar* const* varyings, GLenum bufferMode) {  (*ptrglTransformFeedbackVaryings)(program, count, varyings, bufferMode); }
// int cancallTransformFeedbackVaryings() { return ((ptrglTransformFeedbackVaryings == NULL) ? 0 : 1); }
// void callRenderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height) {  (*ptrglRenderbufferStorage)(target, internalformat, width, height); }
// int cancallRenderbufferStorage() { return ((ptrglRenderbufferStorage == NULL) ? 0 : 1); }
// GLuint callGetUniformBlockIndex(GLuint program, GLchar* uniformBlockName) { return (*ptrglGetUniformBlockIndex)(program, uniformBlockName); }
// int cancallGetUniformBlockIndex() { return ((ptrglGetUniformBlockIndex == NULL) ? 0 : 1); }
// void callObjectLabel(GLenum identifier, GLuint name, GLsizei length, GLchar* label) {  (*ptrglObjectLabel)(identifier, name, length, label); }
// int cancallObjectLabel() { return ((ptrglObjectLabel == NULL) ? 0 : 1); }
// void callGetTexLevelParameterfv(GLenum target, GLint level, GLenum pname, GLfloat* params) {  (*ptrglGetTexLevelParameterfv)(target, level, pname, params); }
// int cancallGetTexLevelParameterfv() { return ((ptrglGetTexLevelParameterfv == NULL) ? 0 : 1); }
// void callProgramUniformMatrix4x3fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix4x3fv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix4x3fv() { return ((ptrglProgramUniformMatrix4x3fv == NULL) ? 0 : 1); }
// void callGetShaderiv(GLuint shader, GLenum pname, GLint* params) {  (*ptrglGetShaderiv)(shader, pname, params); }
// int cancallGetShaderiv() { return ((ptrglGetShaderiv == NULL) ? 0 : 1); }
// void callProgramUniform4ui(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3) {  (*ptrglProgramUniform4ui)(program, location, v0, v1, v2, v3); }
// int cancallProgramUniform4ui() { return ((ptrglProgramUniform4ui == NULL) ? 0 : 1); }
// void callMultiTexCoordP3uiv(GLenum texture, GLenum type_, GLuint* coords) {  (*ptrglMultiTexCoordP3uiv)(texture, type_, coords); }
// int cancallMultiTexCoordP3uiv() { return ((ptrglMultiTexCoordP3uiv == NULL) ? 0 : 1); }
// void callProgramUniform3uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {  (*ptrglProgramUniform3uiv)(program, location, count, value); }
// int cancallProgramUniform3uiv() { return ((ptrglProgramUniform3uiv == NULL) ? 0 : 1); }
// void callGetPointerv(GLenum pname, GLvoid** params) {  (*ptrglGetPointerv)(pname, params); }
// int cancallGetPointerv() { return ((ptrglGetPointerv == NULL) ? 0 : 1); }
// void callMultiTexCoordP4uiv(GLenum texture, GLenum type_, GLuint* coords) {  (*ptrglMultiTexCoordP4uiv)(texture, type_, coords); }
// int cancallMultiTexCoordP4uiv() { return ((ptrglMultiTexCoordP4uiv == NULL) ? 0 : 1); }
// void callUniformMatrix4x2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix4x2fv)(location, count, transpose, value); }
// int cancallUniformMatrix4x2fv() { return ((ptrglUniformMatrix4x2fv == NULL) ? 0 : 1); }
// void callShaderStorageBlockBinding(GLuint program, GLuint storageBlockIndex, GLuint storageBlockBinding) {  (*ptrglShaderStorageBlockBinding)(program, storageBlockIndex, storageBlockBinding); }
// int cancallShaderStorageBlockBinding() { return ((ptrglShaderStorageBlockBinding == NULL) ? 0 : 1); }
// void callShaderBinary(GLsizei count, GLuint* shaders, GLenum binaryformat, GLvoid* binary, GLsizei length) {  (*ptrglShaderBinary)(count, shaders, binaryformat, binary, length); }
// int cancallShaderBinary() { return ((ptrglShaderBinary == NULL) ? 0 : 1); }
// void callUniform1d(GLint location, GLdouble x) {  (*ptrglUniform1d)(location, x); }
// int cancallUniform1d() { return ((ptrglUniform1d == NULL) ? 0 : 1); }
// void callBlendEquationSeparatei(GLuint buf, GLenum modeRGB, GLenum modeAlpha) {  (*ptrglBlendEquationSeparatei)(buf, modeRGB, modeAlpha); }
// int cancallBlendEquationSeparatei() { return ((ptrglBlendEquationSeparatei == NULL) ? 0 : 1); }
// void callGenQueries(GLsizei n, GLuint* ids) {  (*ptrglGenQueries)(n, ids); }
// int cancallGenQueries() { return ((ptrglGenQueries == NULL) ? 0 : 1); }
// void callClampColor(GLenum target, GLenum clamp) {  (*ptrglClampColor)(target, clamp); }
// int cancallClampColor() { return ((ptrglClampColor == NULL) ? 0 : 1); }
// void callPauseTransformFeedback() {  (*ptrglPauseTransformFeedback)(); }
// int cancallPauseTransformFeedback() { return ((ptrglPauseTransformFeedback == NULL) ? 0 : 1); }
// void callLinkProgram(GLuint program) {  (*ptrglLinkProgram)(program); }
// int cancallLinkProgram() { return ((ptrglLinkProgram == NULL) ? 0 : 1); }
// GLenum callCheckFramebufferStatus(GLenum target) { return (*ptrglCheckFramebufferStatus)(target); }
// int cancallCheckFramebufferStatus() { return ((ptrglCheckFramebufferStatus == NULL) ? 0 : 1); }
// void callPointParameteri(GLenum pname, GLint param) {  (*ptrglPointParameteri)(pname, param); }
// int cancallPointParameteri() { return ((ptrglPointParameteri == NULL) ? 0 : 1); }
// void callVertexAttribI4usv(GLuint index, GLushort* v) {  (*ptrglVertexAttribI4usv)(index, v); }
// int cancallVertexAttribI4usv() { return ((ptrglVertexAttribI4usv == NULL) ? 0 : 1); }
// void callPointParameteriv(GLenum pname, GLint* params) {  (*ptrglPointParameteriv)(pname, params); }
// int cancallPointParameteriv() { return ((ptrglPointParameteriv == NULL) ? 0 : 1); }
// void callVertexAttribIPointer(GLuint index, GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {  (*ptrglVertexAttribIPointer)(index, size, type_, stride, pointer); }
// int cancallVertexAttribIPointer() { return ((ptrglVertexAttribIPointer == NULL) ? 0 : 1); }
// void callGenTransformFeedbacks(GLsizei n, GLuint* ids) {  (*ptrglGenTransformFeedbacks)(n, ids); }
// int cancallGenTransformFeedbacks() { return ((ptrglGenTransformFeedbacks == NULL) ? 0 : 1); }
// void callUniform1uiv(GLint location, GLsizei count, GLuint* value) {  (*ptrglUniform1uiv)(location, count, value); }
// int cancallUniform1uiv() { return ((ptrglUniform1uiv == NULL) ? 0 : 1); }
// void callGenerateMipmap(GLenum target) {  (*ptrglGenerateMipmap)(target); }
// int cancallGenerateMipmap() { return ((ptrglGenerateMipmap == NULL) ? 0 : 1); }
// void callGetTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {  (*ptrglGetTexParameterfv)(target, pname, params); }
// int cancallGetTexParameterfv() { return ((ptrglGetTexParameterfv == NULL) ? 0 : 1); }
// void callUniform2uiv(GLint location, GLsizei count, GLuint* value) {  (*ptrglUniform2uiv)(location, count, value); }
// int cancallUniform2uiv() { return ((ptrglUniform2uiv == NULL) ? 0 : 1); }
// void callUniformSubroutinesuiv(GLenum shadertype, GLsizei count, GLuint* indices) {  (*ptrglUniformSubroutinesuiv)(shadertype, count, indices); }
// int cancallUniformSubroutinesuiv() { return ((ptrglUniformSubroutinesuiv == NULL) ? 0 : 1); }
// void callClearDepthf(GLfloat d) {  (*ptrglClearDepthf)(d); }
// int cancallClearDepthf() { return ((ptrglClearDepthf == NULL) ? 0 : 1); }
// void callDeleteVertexArrays(GLsizei n, GLuint* arrays) {  (*ptrglDeleteVertexArrays)(n, arrays); }
// int cancallDeleteVertexArrays() { return ((ptrglDeleteVertexArrays == NULL) ? 0 : 1); }
// GLboolean callIsEnabled(GLenum cap) { return (*ptrglIsEnabled)(cap); }
// int cancallIsEnabled() { return ((ptrglIsEnabled == NULL) ? 0 : 1); }
// void callUniform4ui(GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3) {  (*ptrglUniform4ui)(location, v0, v1, v2, v3); }
// int cancallUniform4ui() { return ((ptrglUniform4ui == NULL) ? 0 : 1); }
// void callGetFramebufferAttachmentParameteriv(GLenum target, GLenum attachment, GLenum pname, GLint* params) {  (*ptrglGetFramebufferAttachmentParameteriv)(target, attachment, pname, params); }
// int cancallGetFramebufferAttachmentParameteriv() { return ((ptrglGetFramebufferAttachmentParameteriv == NULL) ? 0 : 1); }
// void callUniform1dv(GLint location, GLsizei count, GLdouble* value) {  (*ptrglUniform1dv)(location, count, value); }
// int cancallUniform1dv() { return ((ptrglUniform1dv == NULL) ? 0 : 1); }
// void callColorMask(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {  (*ptrglColorMask)(red, green, blue, alpha); }
// int cancallColorMask() { return ((ptrglColorMask == NULL) ? 0 : 1); }
// const GLubyte * callGetString(GLenum name) { return (*ptrglGetString)(name); }
// int cancallGetString() { return ((ptrglGetString == NULL) ? 0 : 1); }
// void callHint(GLenum target, GLenum mode) {  (*ptrglHint)(target, mode); }
// int cancallHint() { return ((ptrglHint == NULL) ? 0 : 1); }
// void callClearBufferData(GLenum target, GLenum internalformat, GLenum format, GLenum type_, void* data) {  (*ptrglClearBufferData)(target, internalformat, format, type_, data); }
// int cancallClearBufferData() { return ((ptrglClearBufferData == NULL) ? 0 : 1); }
// void callUniform2d(GLint location, GLdouble x, GLdouble y) {  (*ptrglUniform2d)(location, x, y); }
// int cancallUniform2d() { return ((ptrglUniform2d == NULL) ? 0 : 1); }
// void callGetInteger64v(GLenum pname, GLint64* params) {  (*ptrglGetInteger64v)(pname, params); }
// int cancallGetInteger64v() { return ((ptrglGetInteger64v == NULL) ? 0 : 1); }
// void callUniform1f(GLint location, GLfloat v0) {  (*ptrglUniform1f)(location, v0); }
// int cancallUniform1f() { return ((ptrglUniform1f == NULL) ? 0 : 1); }
// void callTexBuffer(GLenum target, GLenum internalformat, GLuint buffer) {  (*ptrglTexBuffer)(target, internalformat, buffer); }
// int cancallTexBuffer() { return ((ptrglTexBuffer == NULL) ? 0 : 1); }
// void callEndConditionalRender() {  (*ptrglEndConditionalRender)(); }
// int cancallEndConditionalRender() { return ((ptrglEndConditionalRender == NULL) ? 0 : 1); }
// void callVertexP3ui(GLenum type_, GLuint value) {  (*ptrglVertexP3ui)(type_, value); }
// int cancallVertexP3ui() { return ((ptrglVertexP3ui == NULL) ? 0 : 1); }
// void callGetProgramInfoLog(GLuint program, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {  (*ptrglGetProgramInfoLog)(program, bufSize, length, infoLog); }
// int cancallGetProgramInfoLog() { return ((ptrglGetProgramInfoLog == NULL) ? 0 : 1); }
// void callVertexAttribI2i(GLuint index, GLint x, GLint y) {  (*ptrglVertexAttribI2i)(index, x, y); }
// int cancallVertexAttribI2i() { return ((ptrglVertexAttribI2i == NULL) ? 0 : 1); }
// GLboolean callIsTransformFeedback(GLuint id) { return (*ptrglIsTransformFeedback)(id); }
// int cancallIsTransformFeedback() { return ((ptrglIsTransformFeedback == NULL) ? 0 : 1); }
// void callLogicOp(GLenum opcode) {  (*ptrglLogicOp)(opcode); }
// int cancallLogicOp() { return ((ptrglLogicOp == NULL) ? 0 : 1); }
// void callDispatchComputeIndirect(GLintptr indirect) {  (*ptrglDispatchComputeIndirect)(indirect); }
// int cancallDispatchComputeIndirect() { return ((ptrglDispatchComputeIndirect == NULL) ? 0 : 1); }
// void callBindTexture(GLenum target, GLuint texture) {  (*ptrglBindTexture)(target, texture); }
// int cancallBindTexture() { return ((ptrglBindTexture == NULL) ? 0 : 1); }
// void callDisableVertexAttribArray(GLuint index) {  (*ptrglDisableVertexAttribArray)(index); }
// int cancallDisableVertexAttribArray() { return ((ptrglDisableVertexAttribArray == NULL) ? 0 : 1); }
// void callQueryCounter(GLuint id, GLenum target) {  (*ptrglQueryCounter)(id, target); }
// int cancallQueryCounter() { return ((ptrglQueryCounter == NULL) ? 0 : 1); }
// void callProgramUniform3f(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {  (*ptrglProgramUniform3f)(program, location, v0, v1, v2); }
// int cancallProgramUniform3f() { return ((ptrglProgramUniform3f == NULL) ? 0 : 1); }
// GLboolean callIsEnabledi(GLenum target, GLuint index) { return (*ptrglIsEnabledi)(target, index); }
// int cancallIsEnabledi() { return ((ptrglIsEnabledi == NULL) ? 0 : 1); }
// void callVertexAttribL3d(GLuint index, GLdouble x, GLdouble y, GLdouble z) {  (*ptrglVertexAttribL3d)(index, x, y, z); }
// int cancallVertexAttribL3d() { return ((ptrglVertexAttribL3d == NULL) ? 0 : 1); }
// void callUniform4dv(GLint location, GLsizei count, GLdouble* value) {  (*ptrglUniform4dv)(location, count, value); }
// int cancallUniform4dv() { return ((ptrglUniform4dv == NULL) ? 0 : 1); }
// void callBindVertexArray(GLuint array) {  (*ptrglBindVertexArray)(array); }
// int cancallBindVertexArray() { return ((ptrglBindVertexArray == NULL) ? 0 : 1); }
// void callProgramUniform3dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {  (*ptrglProgramUniform3dv)(program, location, count, value); }
// int cancallProgramUniform3dv() { return ((ptrglProgramUniform3dv == NULL) ? 0 : 1); }
// void callGetFramebufferParameteriv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetFramebufferParameteriv)(target, pname, params); }
// int cancallGetFramebufferParameteriv() { return ((ptrglGetFramebufferParameteriv == NULL) ? 0 : 1); }
// void callCullFace(GLenum mode) {  (*ptrglCullFace)(mode); }
// int cancallCullFace() { return ((ptrglCullFace == NULL) ? 0 : 1); }
// void callMultiDrawArrays(GLenum mode, GLint* first, GLsizei* count, GLsizei drawcount) {  (*ptrglMultiDrawArrays)(mode, first, count, drawcount); }
// int cancallMultiDrawArrays() { return ((ptrglMultiDrawArrays == NULL) ? 0 : 1); }
// void callProgramUniformMatrix2x3dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix2x3dv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix2x3dv() { return ((ptrglProgramUniformMatrix2x3dv == NULL) ? 0 : 1); }
// void callValidateProgramPipeline(GLuint pipeline) {  (*ptrglValidateProgramPipeline)(pipeline); }
// int cancallValidateProgramPipeline() { return ((ptrglValidateProgramPipeline == NULL) ? 0 : 1); }
// void callProgramUniformMatrix3x2fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix3x2fv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix3x2fv() { return ((ptrglProgramUniformMatrix3x2fv == NULL) ? 0 : 1); }
// void callVertexAttribI3i(GLuint index, GLint x, GLint y, GLint z) {  (*ptrglVertexAttribI3i)(index, x, y, z); }
// int cancallVertexAttribI3i() { return ((ptrglVertexAttribI3i == NULL) ? 0 : 1); }
// void callGetProgramResourceiv(GLuint program, GLenum programInterface, GLuint index, GLsizei propCount, GLenum* props, GLsizei bufSize, GLsizei* length, GLint* params) {  (*ptrglGetProgramResourceiv)(program, programInterface, index, propCount, props, bufSize, length, params); }
// int cancallGetProgramResourceiv() { return ((ptrglGetProgramResourceiv == NULL) ? 0 : 1); }
// void callGetUniformSubroutineuiv(GLenum shadertype, GLint location, GLuint* params) {  (*ptrglGetUniformSubroutineuiv)(shadertype, location, params); }
// int cancallGetUniformSubroutineuiv() { return ((ptrglGetUniformSubroutineuiv == NULL) ? 0 : 1); }
// void callProgramUniform2iv(GLuint program, GLint location, GLsizei count, GLint* value) {  (*ptrglProgramUniform2iv)(program, location, count, value); }
// int cancallProgramUniform2iv() { return ((ptrglProgramUniform2iv == NULL) ? 0 : 1); }
// void callGenFramebuffers(GLsizei n, GLuint* framebuffers) {  (*ptrglGenFramebuffers)(n, framebuffers); }
// int cancallGenFramebuffers() { return ((ptrglGenFramebuffers == NULL) ? 0 : 1); }
// void callDrawTransformFeedbackInstanced(GLenum mode, GLuint id, GLsizei instancecount) {  (*ptrglDrawTransformFeedbackInstanced)(mode, id, instancecount); }
// int cancallDrawTransformFeedbackInstanced() { return ((ptrglDrawTransformFeedbackInstanced == NULL) ? 0 : 1); }
// GLboolean callIsBuffer(GLuint buffer) { return (*ptrglIsBuffer)(buffer); }
// int cancallIsBuffer() { return ((ptrglIsBuffer == NULL) ? 0 : 1); }
// void callAttachShader(GLuint program, GLuint shader) {  (*ptrglAttachShader)(program, shader); }
// int cancallAttachShader() { return ((ptrglAttachShader == NULL) ? 0 : 1); }
// void callDrawElementsIndirect(GLenum mode, GLenum type_, GLvoid* indirect) {  (*ptrglDrawElementsIndirect)(mode, type_, indirect); }
// int cancallDrawElementsIndirect() { return ((ptrglDrawElementsIndirect == NULL) ? 0 : 1); }
// void callProgramUniform1ui(GLuint program, GLint location, GLuint v0) {  (*ptrglProgramUniform1ui)(program, location, v0); }
// int cancallProgramUniform1ui() { return ((ptrglProgramUniform1ui == NULL) ? 0 : 1); }
// void callGetUniformIndices(GLuint program, GLsizei uniformCount, GLchar* const* uniformNames, GLuint* uniformIndices) {  (*ptrglGetUniformIndices)(program, uniformCount, uniformNames, uniformIndices); }
// int cancallGetUniformIndices() { return ((ptrglGetUniformIndices == NULL) ? 0 : 1); }
// void callUniform3f(GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {  (*ptrglUniform3f)(location, v0, v1, v2); }
// int cancallUniform3f() { return ((ptrglUniform3f == NULL) ? 0 : 1); }
// void callBindFragDataLocationIndexed(GLuint program, GLuint colorNumber, GLuint index, GLchar* name) {  (*ptrglBindFragDataLocationIndexed)(program, colorNumber, index, name); }
// int cancallBindFragDataLocationIndexed() { return ((ptrglBindFragDataLocationIndexed == NULL) ? 0 : 1); }
// void callTextureView(GLuint texture, GLenum target, GLuint origtexture, GLenum internalformat, GLuint minlevel, GLuint numlevels, GLuint minlayer, GLuint numlayers) {  (*ptrglTextureView)(texture, target, origtexture, internalformat, minlevel, numlevels, minlayer, numlayers); }
// int cancallTextureView() { return ((ptrglTextureView == NULL) ? 0 : 1); }
// void callGetQueryiv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetQueryiv)(target, pname, params); }
// int cancallGetQueryiv() { return ((ptrglGetQueryiv == NULL) ? 0 : 1); }
// void callFramebufferTexture1D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level) {  (*ptrglFramebufferTexture1D)(target, attachment, textarget, texture, level); }
// int cancallFramebufferTexture1D() { return ((ptrglFramebufferTexture1D == NULL) ? 0 : 1); }
// void callPointSize(GLfloat size) {  (*ptrglPointSize)(size); }
// int cancallPointSize() { return ((ptrglPointSize == NULL) ? 0 : 1); }
// void callVertexAttribI4i(GLuint index, GLint x, GLint y, GLint z, GLint w) {  (*ptrglVertexAttribI4i)(index, x, y, z, w); }
// int cancallVertexAttribI4i() { return ((ptrglVertexAttribI4i == NULL) ? 0 : 1); }
// void callPopDebugGroup() {  (*ptrglPopDebugGroup)(); }
// int cancallPopDebugGroup() { return ((ptrglPopDebugGroup == NULL) ? 0 : 1); }
// void callStencilFunc(GLenum func_, GLint ref, GLuint mask) {  (*ptrglStencilFunc)(func_, ref, mask); }
// int cancallStencilFunc() { return ((ptrglStencilFunc == NULL) ? 0 : 1); }
// void callTexImage3DMultisample(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations) {  (*ptrglTexImage3DMultisample)(target, samples, internalformat, width, height, depth, fixedsamplelocations); }
// int cancallTexImage3DMultisample() { return ((ptrglTexImage3DMultisample == NULL) ? 0 : 1); }
// void callProgramUniformMatrix2x3fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix2x3fv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix2x3fv() { return ((ptrglProgramUniformMatrix2x3fv == NULL) ? 0 : 1); }
// void callVertexAttribL1d(GLuint index, GLdouble x) {  (*ptrglVertexAttribL1d)(index, x); }
// int cancallVertexAttribL1d() { return ((ptrglVertexAttribL1d == NULL) ? 0 : 1); }
// void callClearBufferiv(GLenum buffer, GLint drawbuffer, GLint* value) {  (*ptrglClearBufferiv)(buffer, drawbuffer, value); }
// int cancallClearBufferiv() { return ((ptrglClearBufferiv == NULL) ? 0 : 1); }
// void callBindVertexBuffer(GLuint bindingindex, GLuint buffer, GLintptr offset, GLsizei stride) {  (*ptrglBindVertexBuffer)(bindingindex, buffer, offset, stride); }
// int cancallBindVertexBuffer() { return ((ptrglBindVertexBuffer == NULL) ? 0 : 1); }
// void callDisablei(GLenum target, GLuint index) {  (*ptrglDisablei)(target, index); }
// int cancallDisablei() { return ((ptrglDisablei == NULL) ? 0 : 1); }
// void callVertexAttribP1uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {  (*ptrglVertexAttribP1uiv)(index, type_, normalized, value); }
// int cancallVertexAttribP1uiv() { return ((ptrglVertexAttribP1uiv == NULL) ? 0 : 1); }
// void callUniform3dv(GLint location, GLsizei count, GLdouble* value) {  (*ptrglUniform3dv)(location, count, value); }
// int cancallUniform3dv() { return ((ptrglUniform3dv == NULL) ? 0 : 1); }
// void callFramebufferTextureLayer(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer) {  (*ptrglFramebufferTextureLayer)(target, attachment, texture, level, layer); }
// int cancallFramebufferTextureLayer() { return ((ptrglFramebufferTextureLayer == NULL) ? 0 : 1); }
// void callFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level) {  (*ptrglFramebufferTexture2D)(target, attachment, textarget, texture, level); }
// int cancallFramebufferTexture2D() { return ((ptrglFramebufferTexture2D == NULL) ? 0 : 1); }
// void callReleaseShaderCompiler() {  (*ptrglReleaseShaderCompiler)(); }
// int cancallReleaseShaderCompiler() { return ((ptrglReleaseShaderCompiler == NULL) ? 0 : 1); }
// void callProgramUniformMatrix3x4fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix3x4fv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix3x4fv() { return ((ptrglProgramUniformMatrix3x4fv == NULL) ? 0 : 1); }
// void callProgramUniformMatrix2x4dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix2x4dv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix2x4dv() { return ((ptrglProgramUniformMatrix2x4dv == NULL) ? 0 : 1); }
// void callPatchParameteri(GLenum pname, GLint value) {  (*ptrglPatchParameteri)(pname, value); }
// int cancallPatchParameteri() { return ((ptrglPatchParameteri == NULL) ? 0 : 1); }
// void callGetProgramPipelineInfoLog(GLuint pipeline, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {  (*ptrglGetProgramPipelineInfoLog)(pipeline, bufSize, length, infoLog); }
// int cancallGetProgramPipelineInfoLog() { return ((ptrglGetProgramPipelineInfoLog == NULL) ? 0 : 1); }
// void callGetActiveSubroutineName(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name) {  (*ptrglGetActiveSubroutineName)(program, shadertype, index, bufsize, length, name); }
// int cancallGetActiveSubroutineName() { return ((ptrglGetActiveSubroutineName == NULL) ? 0 : 1); }
// void callGetTexParameterIuiv(GLenum target, GLenum pname, GLuint* params) {  (*ptrglGetTexParameterIuiv)(target, pname, params); }
// int cancallGetTexParameterIuiv() { return ((ptrglGetTexParameterIuiv == NULL) ? 0 : 1); }
// void callProgramUniform1fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {  (*ptrglProgramUniform1fv)(program, location, count, value); }
// int cancallProgramUniform1fv() { return ((ptrglProgramUniform1fv == NULL) ? 0 : 1); }
// void callTexCoordP3uiv(GLenum type_, GLuint* coords) {  (*ptrglTexCoordP3uiv)(type_, coords); }
// int cancallTexCoordP3uiv() { return ((ptrglTexCoordP3uiv == NULL) ? 0 : 1); }
// void callGenProgramPipelines(GLsizei n, GLuint* pipelines) {  (*ptrglGenProgramPipelines)(n, pipelines); }
// int cancallGenProgramPipelines() { return ((ptrglGenProgramPipelines == NULL) ? 0 : 1); }
// GLuint callGetDebugMessageLog(GLuint count, GLsizei bufsize, GLenum* sources, GLenum* types, GLuint* ids, GLenum* severities, GLsizei* lengths, GLchar* messageLog) { return (*ptrglGetDebugMessageLog)(count, bufsize, sources, types, ids, severities, lengths, messageLog); }
// int cancallGetDebugMessageLog() { return ((ptrglGetDebugMessageLog == NULL) ? 0 : 1); }
// void callSamplerParameteriv(GLuint sampler, GLenum pname, GLint* param) {  (*ptrglSamplerParameteriv)(sampler, pname, param); }
// int cancallSamplerParameteriv() { return ((ptrglSamplerParameteriv == NULL) ? 0 : 1); }
// void callFrontFace(GLenum mode) {  (*ptrglFrontFace)(mode); }
// int cancallFrontFace() { return ((ptrglFrontFace == NULL) ? 0 : 1); }
// void callDeleteBuffers(GLsizei n, GLuint* buffers) {  (*ptrglDeleteBuffers)(n, buffers); }
// int cancallDeleteBuffers() { return ((ptrglDeleteBuffers == NULL) ? 0 : 1); }
// void callCompressedTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexSubImage3D)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data); }
// int cancallCompressedTexSubImage3D() { return ((ptrglCompressedTexSubImage3D == NULL) ? 0 : 1); }
// void callCompressedTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexSubImage2D)(target, level, xoffset, yoffset, width, height, format, imageSize, data); }
// int cancallCompressedTexSubImage2D() { return ((ptrglCompressedTexSubImage2D == NULL) ? 0 : 1); }
// void callBindProgramPipeline(GLuint pipeline) {  (*ptrglBindProgramPipeline)(pipeline); }
// int cancallBindProgramPipeline() { return ((ptrglBindProgramPipeline == NULL) ? 0 : 1); }
// void callUniform1iv(GLint location, GLsizei count, GLint* value) {  (*ptrglUniform1iv)(location, count, value); }
// int cancallUniform1iv() { return ((ptrglUniform1iv == NULL) ? 0 : 1); }
// void callVertexAttribDivisor(GLuint index, GLuint divisor) {  (*ptrglVertexAttribDivisor)(index, divisor); }
// int cancallVertexAttribDivisor() { return ((ptrglVertexAttribDivisor == NULL) ? 0 : 1); }
// void callBindBufferBase(GLenum target, GLuint index, GLuint buffer) {  (*ptrglBindBufferBase)(target, index, buffer); }
// int cancallBindBufferBase() { return ((ptrglBindBufferBase == NULL) ? 0 : 1); }
// void callColorP4ui(GLenum type_, GLuint color) {  (*ptrglColorP4ui)(type_, color); }
// int cancallColorP4ui() { return ((ptrglColorP4ui == NULL) ? 0 : 1); }
// GLuint callGetProgramResourceIndex(GLuint program, GLenum programInterface, GLchar* name) { return (*ptrglGetProgramResourceIndex)(program, programInterface, name); }
// int cancallGetProgramResourceIndex() { return ((ptrglGetProgramResourceIndex == NULL) ? 0 : 1); }
// void callFramebufferParameteri(GLenum target, GLenum pname, GLint param) {  (*ptrglFramebufferParameteri)(target, pname, param); }
// int cancallFramebufferParameteri() { return ((ptrglFramebufferParameteri == NULL) ? 0 : 1); }
// void callValidateProgram(GLuint program) {  (*ptrglValidateProgram)(program); }
// int cancallValidateProgram() { return ((ptrglValidateProgram == NULL) ? 0 : 1); }
// void callDrawRangeElements(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type_, GLvoid* indices) {  (*ptrglDrawRangeElements)(mode, start, end, count, type_, indices); }
// int cancallDrawRangeElements() { return ((ptrglDrawRangeElements == NULL) ? 0 : 1); }
// void callSamplerParameteri(GLuint sampler, GLenum pname, GLint param) {  (*ptrglSamplerParameteri)(sampler, pname, param); }
// int cancallSamplerParameteri() { return ((ptrglSamplerParameteri == NULL) ? 0 : 1); }
// void callGetActiveAttrib(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type_, GLchar* name) {  (*ptrglGetActiveAttrib)(program, index, bufSize, length, size, type_, name); }
// int cancallGetActiveAttrib() { return ((ptrglGetActiveAttrib == NULL) ? 0 : 1); }
// void callSecondaryColorP3uiv(GLenum type_, GLuint* color) {  (*ptrglSecondaryColorP3uiv)(type_, color); }
// int cancallSecondaryColorP3uiv() { return ((ptrglSecondaryColorP3uiv == NULL) ? 0 : 1); }
// void callSamplerParameterfv(GLuint sampler, GLenum pname, GLfloat* param) {  (*ptrglSamplerParameterfv)(sampler, pname, param); }
// int cancallSamplerParameterfv() { return ((ptrglSamplerParameterfv == NULL) ? 0 : 1); }
// void callDepthRange(GLdouble near_, GLdouble far_) {  (*ptrglDepthRange)(near_, far_); }
// int cancallDepthRange() { return ((ptrglDepthRange == NULL) ? 0 : 1); }
// void callUniform3uiv(GLint location, GLsizei count, GLuint* value) {  (*ptrglUniform3uiv)(location, count, value); }
// int cancallUniform3uiv() { return ((ptrglUniform3uiv == NULL) ? 0 : 1); }
// void callSamplerParameterf(GLuint sampler, GLenum pname, GLfloat param) {  (*ptrglSamplerParameterf)(sampler, pname, param); }
// int cancallSamplerParameterf() { return ((ptrglSamplerParameterf == NULL) ? 0 : 1); }
// void callUniform2f(GLint location, GLfloat v0, GLfloat v1) {  (*ptrglUniform2f)(location, v0, v1); }
// int cancallUniform2f() { return ((ptrglUniform2f == NULL) ? 0 : 1); }
// void callSamplerParameterIuiv(GLuint sampler, GLenum pname, GLuint* param) {  (*ptrglSamplerParameterIuiv)(sampler, pname, param); }
// int cancallSamplerParameterIuiv() { return ((ptrglSamplerParameterIuiv == NULL) ? 0 : 1); }
// void callDrawArraysIndirect(GLenum mode, GLvoid* indirect) {  (*ptrglDrawArraysIndirect)(mode, indirect); }
// int cancallDrawArraysIndirect() { return ((ptrglDrawArraysIndirect == NULL) ? 0 : 1); }
// void callGenTextures(GLsizei n, GLuint* textures) {  (*ptrglGenTextures)(n, textures); }
// int cancallGenTextures() { return ((ptrglGenTextures == NULL) ? 0 : 1); }
// void callBindBufferRange(GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size) {  (*ptrglBindBufferRange)(target, index, buffer, offset, size); }
// int cancallBindBufferRange() { return ((ptrglBindBufferRange == NULL) ? 0 : 1); }
// void callUniformMatrix4x2dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix4x2dv)(location, count, transpose, value); }
// int cancallUniformMatrix4x2dv() { return ((ptrglUniformMatrix4x2dv == NULL) ? 0 : 1); }
// void callViewport(GLint x, GLint y, GLsizei width, GLsizei height) {  (*ptrglViewport)(x, y, width, height); }
// int cancallViewport() { return ((ptrglViewport == NULL) ? 0 : 1); }
// void callMultiTexCoordP1uiv(GLenum texture, GLenum type_, GLuint* coords) {  (*ptrglMultiTexCoordP1uiv)(texture, type_, coords); }
// int cancallMultiTexCoordP1uiv() { return ((ptrglMultiTexCoordP1uiv == NULL) ? 0 : 1); }
// void callUniformMatrix2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix2fv)(location, count, transpose, value); }
// int cancallUniformMatrix2fv() { return ((ptrglUniformMatrix2fv == NULL) ? 0 : 1); }
// void callProgramUniform2f(GLuint program, GLint location, GLfloat v0, GLfloat v1) {  (*ptrglProgramUniform2f)(program, location, v0, v1); }
// int cancallProgramUniform2f() { return ((ptrglProgramUniform2f == NULL) ? 0 : 1); }
// GLint callGetUniformLocation(GLuint program, GLchar* name) { return (*ptrglGetUniformLocation)(program, name); }
// int cancallGetUniformLocation() { return ((ptrglGetUniformLocation == NULL) ? 0 : 1); }
// void callTexStorage2DMultisample(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations) {  (*ptrglTexStorage2DMultisample)(target, samples, internalformat, width, height, fixedsamplelocations); }
// int cancallTexStorage2DMultisample() { return ((ptrglTexStorage2DMultisample == NULL) ? 0 : 1); }
// void callGetFloatv(GLenum pname, GLfloat* params) {  (*ptrglGetFloatv)(pname, params); }
// int cancallGetFloatv() { return ((ptrglGetFloatv == NULL) ? 0 : 1); }
// void callUniformMatrix3x2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix3x2fv)(location, count, transpose, value); }
// int cancallUniformMatrix3x2fv() { return ((ptrglUniformMatrix3x2fv == NULL) ? 0 : 1); }
// void callBindSampler(GLuint unit, GLuint sampler) {  (*ptrglBindSampler)(unit, sampler); }
// int cancallBindSampler() { return ((ptrglBindSampler == NULL) ? 0 : 1); }
// void callUniform4iv(GLint location, GLsizei count, GLint* value) {  (*ptrglUniform4iv)(location, count, value); }
// int cancallUniform4iv() { return ((ptrglUniform4iv == NULL) ? 0 : 1); }
// void callMemoryBarrier(GLbitfield barriers) {  (*ptrglMemoryBarrier)(barriers); }
// int cancallMemoryBarrier() { return ((ptrglMemoryBarrier == NULL) ? 0 : 1); }
// void callVertexAttribPointer(GLuint index, GLint size, GLenum type_, GLboolean normalized, GLsizei stride, GLvoid* pointer) {  (*ptrglVertexAttribPointer)(index, size, type_, normalized, stride, pointer); }
// int cancallVertexAttribPointer() { return ((ptrglVertexAttribPointer == NULL) ? 0 : 1); }
// void callBufferData(GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {  (*ptrglBufferData)(target, size, data, usage); }
// int cancallBufferData() { return ((ptrglBufferData == NULL) ? 0 : 1); }
// void callProgramUniform2dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {  (*ptrglProgramUniform2dv)(program, location, count, value); }
// int cancallProgramUniform2dv() { return ((ptrglProgramUniform2dv == NULL) ? 0 : 1); }
// void callLineWidth(GLfloat width) {  (*ptrglLineWidth)(width); }
// int cancallLineWidth() { return ((ptrglLineWidth == NULL) ? 0 : 1); }
// void callColorP4uiv(GLenum type_, GLuint* color) {  (*ptrglColorP4uiv)(type_, color); }
// int cancallColorP4uiv() { return ((ptrglColorP4uiv == NULL) ? 0 : 1); }
// void callProgramUniform3i(GLuint program, GLint location, GLint v0, GLint v1, GLint v2) {  (*ptrglProgramUniform3i)(program, location, v0, v1, v2); }
// int cancallProgramUniform3i() { return ((ptrglProgramUniform3i == NULL) ? 0 : 1); }
// void callDrawElements(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices) {  (*ptrglDrawElements)(mode, count, type_, indices); }
// int cancallDrawElements() { return ((ptrglDrawElements == NULL) ? 0 : 1); }
// GLint callGetProgramResourceLocation(GLuint program, GLenum programInterface, GLchar* name) { return (*ptrglGetProgramResourceLocation)(program, programInterface, name); }
// int cancallGetProgramResourceLocation() { return ((ptrglGetProgramResourceLocation == NULL) ? 0 : 1); }
// void callInvalidateBufferSubData(GLuint buffer, GLintptr offset, GLsizeiptr length) {  (*ptrglInvalidateBufferSubData)(buffer, offset, length); }
// int cancallInvalidateBufferSubData() { return ((ptrglInvalidateBufferSubData == NULL) ? 0 : 1); }
// void callTexStorage2D(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height) {  (*ptrglTexStorage2D)(target, levels, internalformat, width, height); }
// int cancallTexStorage2D() { return ((ptrglTexStorage2D == NULL) ? 0 : 1); }
// void callDeleteQueries(GLsizei n, GLuint* ids) {  (*ptrglDeleteQueries)(n, ids); }
// int cancallDeleteQueries() { return ((ptrglDeleteQueries == NULL) ? 0 : 1); }
// void callGetTexImage(GLenum target, GLint level, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglGetTexImage)(target, level, format, type_, pixels); }
// int cancallGetTexImage() { return ((ptrglGetTexImage == NULL) ? 0 : 1); }
// GLenum callClientWaitSync(GLsync sync, GLbitfield flags, GLuint64 timeout) { return (*ptrglClientWaitSync)(sync, flags, timeout); }
// int cancallClientWaitSync() { return ((ptrglClientWaitSync == NULL) ? 0 : 1); }
// void callDrawElementsInstancedBaseVertex(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLsizei instancecount, GLint basevertex) {  (*ptrglDrawElementsInstancedBaseVertex)(mode, count, type_, indices, instancecount, basevertex); }
// int cancallDrawElementsInstancedBaseVertex() { return ((ptrglDrawElementsInstancedBaseVertex == NULL) ? 0 : 1); }
// void callPixelStoref(GLenum pname, GLfloat param) {  (*ptrglPixelStoref)(pname, param); }
// int cancallPixelStoref() { return ((ptrglPixelStoref == NULL) ? 0 : 1); }
// void callTexParameteriv(GLenum target, GLenum pname, GLint* params) {  (*ptrglTexParameteriv)(target, pname, params); }
// int cancallTexParameteriv() { return ((ptrglTexParameteriv == NULL) ? 0 : 1); }
// void callFramebufferRenderbuffer(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer) {  (*ptrglFramebufferRenderbuffer)(target, attachment, renderbuffertarget, renderbuffer); }
// int cancallFramebufferRenderbuffer() { return ((ptrglFramebufferRenderbuffer == NULL) ? 0 : 1); }
// void callClearBufferSubData(GLenum target, GLenum internalformat, GLintptr offset, GLsizeiptr size, GLenum format, GLenum type_, void* data) {  (*ptrglClearBufferSubData)(target, internalformat, offset, size, format, type_, data); }
// int cancallClearBufferSubData() { return ((ptrglClearBufferSubData == NULL) ? 0 : 1); }
// void callProgramUniform2d(GLuint program, GLint location, GLdouble v0, GLdouble v1) {  (*ptrglProgramUniform2d)(program, location, v0, v1); }
// int cancallProgramUniform2d() { return ((ptrglProgramUniform2d == NULL) ? 0 : 1); }
// void callSampleCoverage(GLfloat value, GLboolean invert) {  (*ptrglSampleCoverage)(value, invert); }
// int cancallSampleCoverage() { return ((ptrglSampleCoverage == NULL) ? 0 : 1); }
// GLboolean callIsVertexArray(GLuint array) { return (*ptrglIsVertexArray)(array); }
// int cancallIsVertexArray() { return ((ptrglIsVertexArray == NULL) ? 0 : 1); }
// void callDispatchCompute(GLuint num_groups_x, GLuint num_groups_y, GLuint num_groups_z) {  (*ptrglDispatchCompute)(num_groups_x, num_groups_y, num_groups_z); }
// int cancallDispatchCompute() { return ((ptrglDispatchCompute == NULL) ? 0 : 1); }
// GLsync callFenceSync(GLenum condition, GLbitfield flags) { return (*ptrglFenceSync)(condition, flags); }
// int cancallFenceSync() { return ((ptrglFenceSync == NULL) ? 0 : 1); }
// void callGetVertexAttribiv(GLuint index, GLenum pname, GLint* params) {  (*ptrglGetVertexAttribiv)(index, pname, params); }
// int cancallGetVertexAttribiv() { return ((ptrglGetVertexAttribiv == NULL) ? 0 : 1); }
// void callFlushMappedBufferRange(GLenum target, GLintptr offset, GLsizeiptr length) {  (*ptrglFlushMappedBufferRange)(target, offset, length); }
// int cancallFlushMappedBufferRange() { return ((ptrglFlushMappedBufferRange == NULL) ? 0 : 1); }
// void callGetVertexAttribfv(GLuint index, GLenum pname, GLfloat* params) {  (*ptrglGetVertexAttribfv)(index, pname, params); }
// int cancallGetVertexAttribfv() { return ((ptrglGetVertexAttribfv == NULL) ? 0 : 1); }
// GLuint callCreateShader(GLenum type_) { return (*ptrglCreateShader)(type_); }
// int cancallCreateShader() { return ((ptrglCreateShader == NULL) ? 0 : 1); }
// void callVertexAttribFormat(GLuint attribindex, GLint size, GLenum type_, GLboolean normalized, GLuint relativeoffset) {  (*ptrglVertexAttribFormat)(attribindex, size, type_, normalized, relativeoffset); }
// int cancallVertexAttribFormat() { return ((ptrglVertexAttribFormat == NULL) ? 0 : 1); }
// void callMultiDrawElements(GLenum mode, GLsizei* count, GLenum type_, GLvoid* const* indices, GLsizei drawcount) {  (*ptrglMultiDrawElements)(mode, count, type_, indices, drawcount); }
// int cancallMultiDrawElements() { return ((ptrglMultiDrawElements == NULL) ? 0 : 1); }
// void callCopyImageSubData(GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei srcWidth, GLsizei srcHeight, GLsizei srcDepth) {  (*ptrglCopyImageSubData)(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth); }
// int cancallCopyImageSubData() { return ((ptrglCopyImageSubData == NULL) ? 0 : 1); }
// void callDrawElementsBaseVertex(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLint basevertex) {  (*ptrglDrawElementsBaseVertex)(mode, count, type_, indices, basevertex); }
// int cancallDrawElementsBaseVertex() { return ((ptrglDrawElementsBaseVertex == NULL) ? 0 : 1); }
// void callStencilMask(GLuint mask) {  (*ptrglStencilMask)(mask); }
// int cancallStencilMask() { return ((ptrglStencilMask == NULL) ? 0 : 1); }
// void callVertexAttribL4dv(GLuint index, GLdouble* v) {  (*ptrglVertexAttribL4dv)(index, v); }
// int cancallVertexAttribL4dv() { return ((ptrglVertexAttribL4dv == NULL) ? 0 : 1); }
// void callTexParameteri(GLenum target, GLenum pname, GLint param) {  (*ptrglTexParameteri)(target, pname, param); }
// int cancallTexParameteri() { return ((ptrglTexParameteri == NULL) ? 0 : 1); }
// void callTexCoordP1ui(GLenum type_, GLuint coords) {  (*ptrglTexCoordP1ui)(type_, coords); }
// int cancallTexCoordP1ui() { return ((ptrglTexCoordP1ui == NULL) ? 0 : 1); }
// void callVertexAttribL4d(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {  (*ptrglVertexAttribL4d)(index, x, y, z, w); }
// int cancallVertexAttribL4d() { return ((ptrglVertexAttribL4d == NULL) ? 0 : 1); }
// void callTexCoordP2ui(GLenum type_, GLuint coords) {  (*ptrglTexCoordP2ui)(type_, coords); }
// int cancallTexCoordP2ui() { return ((ptrglTexCoordP2ui == NULL) ? 0 : 1); }
// void callGetUniformfv(GLuint program, GLint location, GLfloat* params) {  (*ptrglGetUniformfv)(program, location, params); }
// int cancallGetUniformfv() { return ((ptrglGetUniformfv == NULL) ? 0 : 1); }
// void callVertexAttribI2iv(GLuint index, GLint* v) {  (*ptrglVertexAttribI2iv)(index, v); }
// int cancallVertexAttribI2iv() { return ((ptrglVertexAttribI2iv == NULL) ? 0 : 1); }
// void callGetBooleani_v(GLenum target, GLuint index, GLboolean* data) {  (*ptrglGetBooleani_v)(target, index, data); }
// int cancallGetBooleani_v() { return ((ptrglGetBooleani_v == NULL) ? 0 : 1); }
// GLenum callGetError() { return (*ptrglGetError)(); }
// int cancallGetError() { return ((ptrglGetError == NULL) ? 0 : 1); }
// void callUniformMatrix2x4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix2x4fv)(location, count, transpose, value); }
// int cancallUniformMatrix2x4fv() { return ((ptrglUniformMatrix2x4fv == NULL) ? 0 : 1); }
// void callGetObjectLabel(GLenum identifier, GLuint name, GLsizei bufSize, GLsizei* length, GLchar* label) {  (*ptrglGetObjectLabel)(identifier, name, bufSize, length, label); }
// int cancallGetObjectLabel() { return ((ptrglGetObjectLabel == NULL) ? 0 : 1); }
// void callTexImage2D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexImage2D)(target, level, internalformat, width, height, border, format, type_, pixels); }
// int cancallTexImage2D() { return ((ptrglTexImage2D == NULL) ? 0 : 1); }
// void callProgramUniformMatrix3fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix3fv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix3fv() { return ((ptrglProgramUniformMatrix3fv == NULL) ? 0 : 1); }
// void callTexCoordP3ui(GLenum type_, GLuint coords) {  (*ptrglTexCoordP3ui)(type_, coords); }
// int cancallTexCoordP3ui() { return ((ptrglTexCoordP3ui == NULL) ? 0 : 1); }
// void callGetDoublei_v(GLenum target, GLuint index, GLdouble* data) {  (*ptrglGetDoublei_v)(target, index, data); }
// int cancallGetDoublei_v() { return ((ptrglGetDoublei_v == NULL) ? 0 : 1); }
// void callTexImage3D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexImage3D)(target, level, internalformat, width, height, depth, border, format, type_, pixels); }
// int cancallTexImage3D() { return ((ptrglTexImage3D == NULL) ? 0 : 1); }
// void callGenSamplers(GLsizei count, GLuint* samplers) {  (*ptrglGenSamplers)(count, samplers); }
// int cancallGenSamplers() { return ((ptrglGenSamplers == NULL) ? 0 : 1); }
// void callVertexAttribI4uiv(GLuint index, GLuint* v) {  (*ptrglVertexAttribI4uiv)(index, v); }
// int cancallVertexAttribI4uiv() { return ((ptrglVertexAttribI4uiv == NULL) ? 0 : 1); }
// void callGetUniformdv(GLuint program, GLint location, GLdouble* params) {  (*ptrglGetUniformdv)(program, location, params); }
// int cancallGetUniformdv() { return ((ptrglGetUniformdv == NULL) ? 0 : 1); }
// void callProgramUniform4iv(GLuint program, GLint location, GLsizei count, GLint* value) {  (*ptrglProgramUniform4iv)(program, location, count, value); }
// int cancallProgramUniform4iv() { return ((ptrglProgramUniform4iv == NULL) ? 0 : 1); }
// void callProgramUniform1i(GLuint program, GLint location, GLint v0) {  (*ptrglProgramUniform1i)(program, location, v0); }
// int cancallProgramUniform1i() { return ((ptrglProgramUniform1i == NULL) ? 0 : 1); }
// void callMultiTexCoordP1ui(GLenum texture, GLenum type_, GLuint coords) {  (*ptrglMultiTexCoordP1ui)(texture, type_, coords); }
// int cancallMultiTexCoordP1ui() { return ((ptrglMultiTexCoordP1ui == NULL) ? 0 : 1); }
// void callProgramUniformMatrix4dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix4dv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix4dv() { return ((ptrglProgramUniformMatrix4dv == NULL) ? 0 : 1); }
// void callUniform1fv(GLint location, GLsizei count, GLfloat* value) {  (*ptrglUniform1fv)(location, count, value); }
// int cancallUniform1fv() { return ((ptrglUniform1fv == NULL) ? 0 : 1); }
// void callDrawElementsInstancedBaseVertexBaseInstance(GLenum mode, GLsizei count, GLenum type_, void* indices, GLsizei instancecount, GLint basevertex, GLuint baseinstance) {  (*ptrglDrawElementsInstancedBaseVertexBaseInstance)(mode, count, type_, indices, instancecount, basevertex, baseinstance); }
// int cancallDrawElementsInstancedBaseVertexBaseInstance() { return ((ptrglDrawElementsInstancedBaseVertexBaseInstance == NULL) ? 0 : 1); }
// void callDrawElementsInstancedBaseInstance(GLenum mode, GLsizei count, GLenum type_, void* indices, GLsizei instancecount, GLuint baseinstance) {  (*ptrglDrawElementsInstancedBaseInstance)(mode, count, type_, indices, instancecount, baseinstance); }
// int cancallDrawElementsInstancedBaseInstance() { return ((ptrglDrawElementsInstancedBaseInstance == NULL) ? 0 : 1); }
// void callDeleteTextures(GLsizei n, GLuint* textures) {  (*ptrglDeleteTextures)(n, textures); }
// int cancallDeleteTextures() { return ((ptrglDeleteTextures == NULL) ? 0 : 1); }
// void callGetProgramiv(GLuint program, GLenum pname, GLint* params) {  (*ptrglGetProgramiv)(program, pname, params); }
// int cancallGetProgramiv() { return ((ptrglGetProgramiv == NULL) ? 0 : 1); }
// void callGetRenderbufferParameteriv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetRenderbufferParameteriv)(target, pname, params); }
// int cancallGetRenderbufferParameteriv() { return ((ptrglGetRenderbufferParameteriv == NULL) ? 0 : 1); }
// void callClear(GLbitfield mask) {  (*ptrglClear)(mask); }
// int cancallClear() { return ((ptrglClear == NULL) ? 0 : 1); }
// void callPolygonMode(GLenum face, GLenum mode) {  (*ptrglPolygonMode)(face, mode); }
// int cancallPolygonMode() { return ((ptrglPolygonMode == NULL) ? 0 : 1); }
// void callProgramUniform4uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {  (*ptrglProgramUniform4uiv)(program, location, count, value); }
// int cancallProgramUniform4uiv() { return ((ptrglProgramUniform4uiv == NULL) ? 0 : 1); }
// GLint callGetSubroutineUniformLocation(GLuint program, GLenum shadertype, GLchar* name) { return (*ptrglGetSubroutineUniformLocation)(program, shadertype, name); }
// int cancallGetSubroutineUniformLocation() { return ((ptrglGetSubroutineUniformLocation == NULL) ? 0 : 1); }
// void callDrawArraysInstancedBaseInstance(GLenum mode, GLint first, GLsizei count, GLsizei instancecount, GLuint baseinstance) {  (*ptrglDrawArraysInstancedBaseInstance)(mode, first, count, instancecount, baseinstance); }
// int cancallDrawArraysInstancedBaseInstance() { return ((ptrglDrawArraysInstancedBaseInstance == NULL) ? 0 : 1); }
// void callGetVertexAttribPointerv(GLuint index, GLenum pname, GLvoid** pointer) {  (*ptrglGetVertexAttribPointerv)(index, pname, pointer); }
// int cancallGetVertexAttribPointerv() { return ((ptrglGetVertexAttribPointerv == NULL) ? 0 : 1); }
// void callEnableVertexAttribArray(GLuint index) {  (*ptrglEnableVertexAttribArray)(index); }
// int cancallEnableVertexAttribArray() { return ((ptrglEnableVertexAttribArray == NULL) ? 0 : 1); }
// GLboolean callIsQuery(GLuint id) { return (*ptrglIsQuery)(id); }
// int cancallIsQuery() { return ((ptrglIsQuery == NULL) ? 0 : 1); }
// void callProgramUniformMatrix4x2fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix4x2fv)(program, location, count, transpose, value); }
// int cancallProgramUniformMatrix4x2fv() { return ((ptrglProgramUniformMatrix4x2fv == NULL) ? 0 : 1); }
// void callStencilFuncSeparate(GLenum face, GLenum func_, GLint ref, GLuint mask) {  (*ptrglStencilFuncSeparate)(face, func_, ref, mask); }
// int cancallStencilFuncSeparate() { return ((ptrglStencilFuncSeparate == NULL) ? 0 : 1); }
// void callPixelStorei(GLenum pname, GLint param) {  (*ptrglPixelStorei)(pname, param); }
// int cancallPixelStorei() { return ((ptrglPixelStorei == NULL) ? 0 : 1); }
// void callTexParameterIiv(GLenum target, GLenum pname, GLint* params) {  (*ptrglTexParameterIiv)(target, pname, params); }
// int cancallTexParameterIiv() { return ((ptrglTexParameterIiv == NULL) ? 0 : 1); }
// void callGetActiveAtomicCounterBufferiv(GLuint program, GLuint bufferIndex, GLenum pname, GLint* params) {  (*ptrglGetActiveAtomicCounterBufferiv)(program, bufferIndex, pname, params); }
// int cancallGetActiveAtomicCounterBufferiv() { return ((ptrglGetActiveAtomicCounterBufferiv == NULL) ? 0 : 1); }
// void callGetActiveSubroutineUniformiv(GLuint program, GLenum shadertype, GLuint index, GLenum pname, GLint* values) {  (*ptrglGetActiveSubroutineUniformiv)(program, shadertype, index, pname, values); }
// int cancallGetActiveSubroutineUniformiv() { return ((ptrglGetActiveSubroutineUniformiv == NULL) ? 0 : 1); }
// void callColorMaski(GLuint index, GLboolean r, GLboolean g, GLboolean b, GLboolean a) {  (*ptrglColorMaski)(index, r, g, b, a); }
// int cancallColorMaski() { return ((ptrglColorMaski == NULL) ? 0 : 1); }
// void callViewportIndexedfv(GLuint index, GLfloat* v) {  (*ptrglViewportIndexedfv)(index, v); }
// int cancallViewportIndexedfv() { return ((ptrglViewportIndexedfv == NULL) ? 0 : 1); }
// void callProgramUniform3ui(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2) {  (*ptrglProgramUniform3ui)(program, location, v0, v1, v2); }
// int cancallProgramUniform3ui() { return ((ptrglProgramUniform3ui == NULL) ? 0 : 1); }
// void callMultiTexCoordP2ui(GLenum texture, GLenum type_, GLuint coords) {  (*ptrglMultiTexCoordP2ui)(texture, type_, coords); }
// int cancallMultiTexCoordP2ui() { return ((ptrglMultiTexCoordP2ui == NULL) ? 0 : 1); }
// void callProgramUniform4f(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {  (*ptrglProgramUniform4f)(program, location, v0, v1, v2, v3); }
// int cancallProgramUniform4f() { return ((ptrglProgramUniform4f == NULL) ? 0 : 1); }
// void callColorP3ui(GLenum type_, GLuint color) {  (*ptrglColorP3ui)(type_, color); }
// int cancallColorP3ui() { return ((ptrglColorP3ui == NULL) ? 0 : 1); }
// void callBlendEquationSeparate(GLenum modeRGB, GLenum modeAlpha) {  (*ptrglBlendEquationSeparate)(modeRGB, modeAlpha); }
// int cancallBlendEquationSeparate() { return ((ptrglBlendEquationSeparate == NULL) ? 0 : 1); }
// void callProgramUniform4i(GLuint program, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {  (*ptrglProgramUniform4i)(program, location, v0, v1, v2, v3); }
// int cancallProgramUniform4i() { return ((ptrglProgramUniform4i == NULL) ? 0 : 1); }
// void callFlush() {  (*ptrglFlush)(); }
// int cancallFlush() { return ((ptrglFlush == NULL) ? 0 : 1); }
// void callFinish() {  (*ptrglFinish)(); }
// int cancallFinish() { return ((ptrglFinish == NULL) ? 0 : 1); }
// void callUniformMatrix3dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix3dv)(location, count, transpose, value); }
// int cancallUniformMatrix3dv() { return ((ptrglUniformMatrix3dv == NULL) ? 0 : 1); }
// void callEndTransformFeedback() {  (*ptrglEndTransformFeedback)(); }
// int cancallEndTransformFeedback() { return ((ptrglEndTransformFeedback == NULL) ? 0 : 1); }
// void callProgramUniform3d(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2) {  (*ptrglProgramUniform3d)(program, location, v0, v1, v2); }
// int cancallProgramUniform3d() { return ((ptrglProgramUniform3d == NULL) ? 0 : 1); }
// void callGetIntegerv(GLenum pname, GLint* params) {  (*ptrglGetIntegerv)(pname, params); }
// int cancallGetIntegerv() { return ((ptrglGetIntegerv == NULL) ? 0 : 1); }
// int coreGetProcAddresses() {
// 	ptrglProgramUniform1dv = coreGetProcAddress("glProgramUniform1dv");
// 	ptrglMultiDrawArraysIndirect = coreGetProcAddress("glMultiDrawArraysIndirect");
// 	ptrglMultiTexCoordP2uiv = coreGetProcAddress("glMultiTexCoordP2uiv");
// 	ptrglInvalidateSubFramebuffer = coreGetProcAddress("glInvalidateSubFramebuffer");
// 	ptrglUniform3i = coreGetProcAddress("glUniform3i");
// 	ptrglProgramUniform2i = coreGetProcAddress("glProgramUniform2i");
// 	ptrglShaderSource = coreGetProcAddress("glShaderSource");
// 	ptrglTexCoordP1uiv = coreGetProcAddress("glTexCoordP1uiv");
// 	ptrglUniformBlockBinding = coreGetProcAddress("glUniformBlockBinding");
// 	ptrglUniform4fv = coreGetProcAddress("glUniform4fv");
// 	ptrglGetVertexAttribdv = coreGetProcAddress("glGetVertexAttribdv");
// 	ptrglGenBuffers = coreGetProcAddress("glGenBuffers");
// 	ptrglVertexAttribP4ui = coreGetProcAddress("glVertexAttribP4ui");
// 	ptrglGetInternalformativ = coreGetProcAddress("glGetInternalformativ");
// 	ptrglProgramUniformMatrix4x2dv = coreGetProcAddress("glProgramUniformMatrix4x2dv");
// 	ptrglCompileShader = coreGetProcAddress("glCompileShader");
// 	ptrglBindRenderbuffer = coreGetProcAddress("glBindRenderbuffer");
// 	ptrglCompressedTexSubImage1D = coreGetProcAddress("glCompressedTexSubImage1D");
// 	ptrglBlendFunc = coreGetProcAddress("glBlendFunc");
// 	ptrglTexParameterfv = coreGetProcAddress("glTexParameterfv");
// 	ptrglInvalidateFramebuffer = coreGetProcAddress("glInvalidateFramebuffer");
// 	ptrglBindTransformFeedback = coreGetProcAddress("glBindTransformFeedback");
// 	ptrglGetFragDataIndex = coreGetProcAddress("glGetFragDataIndex");
// 	ptrglTransformFeedbackVaryings = coreGetProcAddress("glTransformFeedbackVaryings");
// 	ptrglRenderbufferStorage = coreGetProcAddress("glRenderbufferStorage");
// 	ptrglGetUniformBlockIndex = coreGetProcAddress("glGetUniformBlockIndex");
// 	ptrglObjectLabel = coreGetProcAddress("glObjectLabel");
// 	ptrglGetTexLevelParameterfv = coreGetProcAddress("glGetTexLevelParameterfv");
// 	ptrglProgramUniformMatrix4x3fv = coreGetProcAddress("glProgramUniformMatrix4x3fv");
// 	ptrglGetShaderiv = coreGetProcAddress("glGetShaderiv");
// 	ptrglProgramUniform4ui = coreGetProcAddress("glProgramUniform4ui");
// 	ptrglMultiTexCoordP3uiv = coreGetProcAddress("glMultiTexCoordP3uiv");
// 	ptrglProgramUniform3uiv = coreGetProcAddress("glProgramUniform3uiv");
// 	ptrglGetPointerv = coreGetProcAddress("glGetPointerv");
// 	ptrglMultiTexCoordP4uiv = coreGetProcAddress("glMultiTexCoordP4uiv");
// 	ptrglUniformMatrix4x2fv = coreGetProcAddress("glUniformMatrix4x2fv");
// 	ptrglShaderStorageBlockBinding = coreGetProcAddress("glShaderStorageBlockBinding");
// 	ptrglShaderBinary = coreGetProcAddress("glShaderBinary");
// 	ptrglUniform1d = coreGetProcAddress("glUniform1d");
// 	ptrglBlendEquationSeparatei = coreGetProcAddress("glBlendEquationSeparatei");
// 	ptrglGenQueries = coreGetProcAddress("glGenQueries");
// 	ptrglClampColor = coreGetProcAddress("glClampColor");
// 	ptrglPauseTransformFeedback = coreGetProcAddress("glPauseTransformFeedback");
// 	ptrglLinkProgram = coreGetProcAddress("glLinkProgram");
// 	ptrglCheckFramebufferStatus = coreGetProcAddress("glCheckFramebufferStatus");
// 	ptrglPointParameteri = coreGetProcAddress("glPointParameteri");
// 	ptrglVertexAttribI4usv = coreGetProcAddress("glVertexAttribI4usv");
// 	ptrglPointParameteriv = coreGetProcAddress("glPointParameteriv");
// 	ptrglVertexAttribIPointer = coreGetProcAddress("glVertexAttribIPointer");
// 	ptrglGenTransformFeedbacks = coreGetProcAddress("glGenTransformFeedbacks");
// 	ptrglUniform1uiv = coreGetProcAddress("glUniform1uiv");
// 	ptrglGenerateMipmap = coreGetProcAddress("glGenerateMipmap");
// 	ptrglGetTexParameterfv = coreGetProcAddress("glGetTexParameterfv");
// 	ptrglUniform2uiv = coreGetProcAddress("glUniform2uiv");
// 	ptrglUniformSubroutinesuiv = coreGetProcAddress("glUniformSubroutinesuiv");
// 	ptrglClearDepthf = coreGetProcAddress("glClearDepthf");
// 	ptrglDeleteVertexArrays = coreGetProcAddress("glDeleteVertexArrays");
// 	ptrglIsEnabled = coreGetProcAddress("glIsEnabled");
// 	ptrglUniform4ui = coreGetProcAddress("glUniform4ui");
// 	ptrglGetFramebufferAttachmentParameteriv = coreGetProcAddress("glGetFramebufferAttachmentParameteriv");
// 	ptrglUniform1dv = coreGetProcAddress("glUniform1dv");
// 	ptrglColorMask = coreGetProcAddress("glColorMask");
// 	ptrglGetString = coreGetProcAddress("glGetString");
// 	ptrglHint = coreGetProcAddress("glHint");
// 	ptrglClearBufferData = coreGetProcAddress("glClearBufferData");
// 	ptrglUniform2d = coreGetProcAddress("glUniform2d");
// 	ptrglGetInteger64v = coreGetProcAddress("glGetInteger64v");
// 	ptrglUniform1f = coreGetProcAddress("glUniform1f");
// 	ptrglTexBuffer = coreGetProcAddress("glTexBuffer");
// 	ptrglEndConditionalRender = coreGetProcAddress("glEndConditionalRender");
// 	ptrglVertexP3ui = coreGetProcAddress("glVertexP3ui");
// 	ptrglGetProgramInfoLog = coreGetProcAddress("glGetProgramInfoLog");
// 	ptrglVertexAttribI2i = coreGetProcAddress("glVertexAttribI2i");
// 	ptrglIsTransformFeedback = coreGetProcAddress("glIsTransformFeedback");
// 	ptrglLogicOp = coreGetProcAddress("glLogicOp");
// 	ptrglDispatchComputeIndirect = coreGetProcAddress("glDispatchComputeIndirect");
// 	ptrglBindTexture = coreGetProcAddress("glBindTexture");
// 	ptrglDisableVertexAttribArray = coreGetProcAddress("glDisableVertexAttribArray");
// 	ptrglQueryCounter = coreGetProcAddress("glQueryCounter");
// 	ptrglProgramUniform3f = coreGetProcAddress("glProgramUniform3f");
// 	ptrglIsEnabledi = coreGetProcAddress("glIsEnabledi");
// 	ptrglVertexAttribL3d = coreGetProcAddress("glVertexAttribL3d");
// 	ptrglUniform4dv = coreGetProcAddress("glUniform4dv");
// 	ptrglBindVertexArray = coreGetProcAddress("glBindVertexArray");
// 	ptrglProgramUniform3dv = coreGetProcAddress("glProgramUniform3dv");
// 	ptrglGetFramebufferParameteriv = coreGetProcAddress("glGetFramebufferParameteriv");
// 	ptrglCullFace = coreGetProcAddress("glCullFace");
// 	ptrglMultiDrawArrays = coreGetProcAddress("glMultiDrawArrays");
// 	ptrglProgramUniformMatrix2x3dv = coreGetProcAddress("glProgramUniformMatrix2x3dv");
// 	ptrglValidateProgramPipeline = coreGetProcAddress("glValidateProgramPipeline");
// 	ptrglProgramUniformMatrix3x2fv = coreGetProcAddress("glProgramUniformMatrix3x2fv");
// 	ptrglVertexAttribI3i = coreGetProcAddress("glVertexAttribI3i");
// 	ptrglGetProgramResourceiv = coreGetProcAddress("glGetProgramResourceiv");
// 	ptrglGetUniformSubroutineuiv = coreGetProcAddress("glGetUniformSubroutineuiv");
// 	ptrglProgramUniform2iv = coreGetProcAddress("glProgramUniform2iv");
// 	ptrglGenFramebuffers = coreGetProcAddress("glGenFramebuffers");
// 	ptrglDrawTransformFeedbackInstanced = coreGetProcAddress("glDrawTransformFeedbackInstanced");
// 	ptrglIsBuffer = coreGetProcAddress("glIsBuffer");
// 	ptrglAttachShader = coreGetProcAddress("glAttachShader");
// 	ptrglDrawElementsIndirect = coreGetProcAddress("glDrawElementsIndirect");
// 	ptrglProgramUniform1ui = coreGetProcAddress("glProgramUniform1ui");
// 	ptrglGetUniformIndices = coreGetProcAddress("glGetUniformIndices");
// 	ptrglUniform3f = coreGetProcAddress("glUniform3f");
// 	ptrglBindFragDataLocationIndexed = coreGetProcAddress("glBindFragDataLocationIndexed");
// 	ptrglTextureView = coreGetProcAddress("glTextureView");
// 	ptrglGetQueryiv = coreGetProcAddress("glGetQueryiv");
// 	ptrglFramebufferTexture1D = coreGetProcAddress("glFramebufferTexture1D");
// 	ptrglPointSize = coreGetProcAddress("glPointSize");
// 	ptrglVertexAttribI4i = coreGetProcAddress("glVertexAttribI4i");
// 	ptrglPopDebugGroup = coreGetProcAddress("glPopDebugGroup");
// 	ptrglStencilFunc = coreGetProcAddress("glStencilFunc");
// 	ptrglTexImage3DMultisample = coreGetProcAddress("glTexImage3DMultisample");
// 	ptrglProgramUniformMatrix2x3fv = coreGetProcAddress("glProgramUniformMatrix2x3fv");
// 	ptrglVertexAttribL1d = coreGetProcAddress("glVertexAttribL1d");
// 	ptrglClearBufferiv = coreGetProcAddress("glClearBufferiv");
// 	ptrglBindVertexBuffer = coreGetProcAddress("glBindVertexBuffer");
// 	ptrglDisablei = coreGetProcAddress("glDisablei");
// 	ptrglVertexAttribP1uiv = coreGetProcAddress("glVertexAttribP1uiv");
// 	ptrglUniform3dv = coreGetProcAddress("glUniform3dv");
// 	ptrglFramebufferTextureLayer = coreGetProcAddress("glFramebufferTextureLayer");
// 	ptrglFramebufferTexture2D = coreGetProcAddress("glFramebufferTexture2D");
// 	ptrglReleaseShaderCompiler = coreGetProcAddress("glReleaseShaderCompiler");
// 	ptrglProgramUniformMatrix3x4fv = coreGetProcAddress("glProgramUniformMatrix3x4fv");
// 	ptrglProgramUniformMatrix2x4dv = coreGetProcAddress("glProgramUniformMatrix2x4dv");
// 	ptrglPatchParameteri = coreGetProcAddress("glPatchParameteri");
// 	ptrglGetProgramPipelineInfoLog = coreGetProcAddress("glGetProgramPipelineInfoLog");
// 	ptrglGetActiveSubroutineName = coreGetProcAddress("glGetActiveSubroutineName");
// 	ptrglGetTexParameterIuiv = coreGetProcAddress("glGetTexParameterIuiv");
// 	ptrglProgramUniform1fv = coreGetProcAddress("glProgramUniform1fv");
// 	ptrglTexCoordP3uiv = coreGetProcAddress("glTexCoordP3uiv");
// 	ptrglGenProgramPipelines = coreGetProcAddress("glGenProgramPipelines");
// 	ptrglGetDebugMessageLog = coreGetProcAddress("glGetDebugMessageLog");
// 	ptrglSamplerParameteriv = coreGetProcAddress("glSamplerParameteriv");
// 	ptrglFrontFace = coreGetProcAddress("glFrontFace");
// 	ptrglDeleteBuffers = coreGetProcAddress("glDeleteBuffers");
// 	ptrglCompressedTexSubImage3D = coreGetProcAddress("glCompressedTexSubImage3D");
// 	ptrglCompressedTexSubImage2D = coreGetProcAddress("glCompressedTexSubImage2D");
// 	ptrglBindProgramPipeline = coreGetProcAddress("glBindProgramPipeline");
// 	ptrglUniform1iv = coreGetProcAddress("glUniform1iv");
// 	ptrglVertexAttribDivisor = coreGetProcAddress("glVertexAttribDivisor");
// 	ptrglBindBufferBase = coreGetProcAddress("glBindBufferBase");
// 	ptrglColorP4ui = coreGetProcAddress("glColorP4ui");
// 	ptrglGetProgramResourceIndex = coreGetProcAddress("glGetProgramResourceIndex");
// 	ptrglFramebufferParameteri = coreGetProcAddress("glFramebufferParameteri");
// 	ptrglValidateProgram = coreGetProcAddress("glValidateProgram");
// 	ptrglDrawRangeElements = coreGetProcAddress("glDrawRangeElements");
// 	ptrglSamplerParameteri = coreGetProcAddress("glSamplerParameteri");
// 	ptrglGetActiveAttrib = coreGetProcAddress("glGetActiveAttrib");
// 	ptrglSecondaryColorP3uiv = coreGetProcAddress("glSecondaryColorP3uiv");
// 	ptrglSamplerParameterfv = coreGetProcAddress("glSamplerParameterfv");
// 	ptrglDepthRange = coreGetProcAddress("glDepthRange");
// 	ptrglUniform3uiv = coreGetProcAddress("glUniform3uiv");
// 	ptrglSamplerParameterf = coreGetProcAddress("glSamplerParameterf");
// 	ptrglUniform2f = coreGetProcAddress("glUniform2f");
// 	ptrglSamplerParameterIuiv = coreGetProcAddress("glSamplerParameterIuiv");
// 	ptrglDrawArraysIndirect = coreGetProcAddress("glDrawArraysIndirect");
// 	ptrglGenTextures = coreGetProcAddress("glGenTextures");
// 	ptrglBindBufferRange = coreGetProcAddress("glBindBufferRange");
// 	ptrglUniformMatrix4x2dv = coreGetProcAddress("glUniformMatrix4x2dv");
// 	ptrglViewport = coreGetProcAddress("glViewport");
// 	ptrglMultiTexCoordP1uiv = coreGetProcAddress("glMultiTexCoordP1uiv");
// 	ptrglUniformMatrix2fv = coreGetProcAddress("glUniformMatrix2fv");
// 	ptrglProgramUniform2f = coreGetProcAddress("glProgramUniform2f");
// 	ptrglGetUniformLocation = coreGetProcAddress("glGetUniformLocation");
// 	ptrglTexStorage2DMultisample = coreGetProcAddress("glTexStorage2DMultisample");
// 	ptrglGetFloatv = coreGetProcAddress("glGetFloatv");
// 	ptrglUniformMatrix3x2fv = coreGetProcAddress("glUniformMatrix3x2fv");
// 	ptrglBindSampler = coreGetProcAddress("glBindSampler");
// 	ptrglUniform4iv = coreGetProcAddress("glUniform4iv");
// 	ptrglMemoryBarrier = coreGetProcAddress("glMemoryBarrier");
// 	ptrglVertexAttribPointer = coreGetProcAddress("glVertexAttribPointer");
// 	ptrglBufferData = coreGetProcAddress("glBufferData");
// 	ptrglProgramUniform2dv = coreGetProcAddress("glProgramUniform2dv");
// 	ptrglLineWidth = coreGetProcAddress("glLineWidth");
// 	ptrglColorP4uiv = coreGetProcAddress("glColorP4uiv");
// 	ptrglProgramUniform3i = coreGetProcAddress("glProgramUniform3i");
// 	ptrglDrawElements = coreGetProcAddress("glDrawElements");
// 	ptrglGetProgramResourceLocation = coreGetProcAddress("glGetProgramResourceLocation");
// 	ptrglInvalidateBufferSubData = coreGetProcAddress("glInvalidateBufferSubData");
// 	ptrglTexStorage2D = coreGetProcAddress("glTexStorage2D");
// 	ptrglDeleteQueries = coreGetProcAddress("glDeleteQueries");
// 	ptrglGetTexImage = coreGetProcAddress("glGetTexImage");
// 	ptrglClientWaitSync = coreGetProcAddress("glClientWaitSync");
// 	ptrglDrawElementsInstancedBaseVertex = coreGetProcAddress("glDrawElementsInstancedBaseVertex");
// 	ptrglPixelStoref = coreGetProcAddress("glPixelStoref");
// 	ptrglTexParameteriv = coreGetProcAddress("glTexParameteriv");
// 	ptrglFramebufferRenderbuffer = coreGetProcAddress("glFramebufferRenderbuffer");
// 	ptrglClearBufferSubData = coreGetProcAddress("glClearBufferSubData");
// 	ptrglProgramUniform2d = coreGetProcAddress("glProgramUniform2d");
// 	ptrglSampleCoverage = coreGetProcAddress("glSampleCoverage");
// 	ptrglIsVertexArray = coreGetProcAddress("glIsVertexArray");
// 	ptrglDispatchCompute = coreGetProcAddress("glDispatchCompute");
// 	ptrglFenceSync = coreGetProcAddress("glFenceSync");
// 	ptrglGetVertexAttribiv = coreGetProcAddress("glGetVertexAttribiv");
// 	ptrglFlushMappedBufferRange = coreGetProcAddress("glFlushMappedBufferRange");
// 	ptrglGetVertexAttribfv = coreGetProcAddress("glGetVertexAttribfv");
// 	ptrglCreateShader = coreGetProcAddress("glCreateShader");
// 	ptrglVertexAttribFormat = coreGetProcAddress("glVertexAttribFormat");
// 	ptrglMultiDrawElements = coreGetProcAddress("glMultiDrawElements");
// 	ptrglCopyImageSubData = coreGetProcAddress("glCopyImageSubData");
// 	ptrglDrawElementsBaseVertex = coreGetProcAddress("glDrawElementsBaseVertex");
// 	ptrglStencilMask = coreGetProcAddress("glStencilMask");
// 	ptrglVertexAttribL4dv = coreGetProcAddress("glVertexAttribL4dv");
// 	ptrglTexParameteri = coreGetProcAddress("glTexParameteri");
// 	ptrglTexCoordP1ui = coreGetProcAddress("glTexCoordP1ui");
// 	ptrglVertexAttribL4d = coreGetProcAddress("glVertexAttribL4d");
// 	ptrglTexCoordP2ui = coreGetProcAddress("glTexCoordP2ui");
// 	ptrglGetUniformfv = coreGetProcAddress("glGetUniformfv");
// 	ptrglVertexAttribI2iv = coreGetProcAddress("glVertexAttribI2iv");
// 	ptrglGetBooleani_v = coreGetProcAddress("glGetBooleani_v");
// 	ptrglGetError = coreGetProcAddress("glGetError");
// 	ptrglUniformMatrix2x4fv = coreGetProcAddress("glUniformMatrix2x4fv");
// 	ptrglGetObjectLabel = coreGetProcAddress("glGetObjectLabel");
// 	ptrglTexImage2D = coreGetProcAddress("glTexImage2D");
// 	ptrglProgramUniformMatrix3fv = coreGetProcAddress("glProgramUniformMatrix3fv");
// 	ptrglTexCoordP3ui = coreGetProcAddress("glTexCoordP3ui");
// 	ptrglGetDoublei_v = coreGetProcAddress("glGetDoublei_v");
// 	ptrglTexImage3D = coreGetProcAddress("glTexImage3D");
// 	ptrglGenSamplers = coreGetProcAddress("glGenSamplers");
// 	ptrglVertexAttribI4uiv = coreGetProcAddress("glVertexAttribI4uiv");
// 	ptrglGetUniformdv = coreGetProcAddress("glGetUniformdv");
// 	ptrglProgramUniform4iv = coreGetProcAddress("glProgramUniform4iv");
// 	ptrglProgramUniform1i = coreGetProcAddress("glProgramUniform1i");
// 	ptrglMultiTexCoordP1ui = coreGetProcAddress("glMultiTexCoordP1ui");
// 	ptrglProgramUniformMatrix4dv = coreGetProcAddress("glProgramUniformMatrix4dv");
// 	ptrglUniform1fv = coreGetProcAddress("glUniform1fv");
// 	ptrglDrawElementsInstancedBaseVertexBaseInstance = coreGetProcAddress("glDrawElementsInstancedBaseVertexBaseInstance");
// 	ptrglDrawElementsInstancedBaseInstance = coreGetProcAddress("glDrawElementsInstancedBaseInstance");
// 	ptrglDeleteTextures = coreGetProcAddress("glDeleteTextures");
// 	ptrglGetProgramiv = coreGetProcAddress("glGetProgramiv");
// 	ptrglGetRenderbufferParameteriv = coreGetProcAddress("glGetRenderbufferParameteriv");
// 	ptrglClear = coreGetProcAddress("glClear");
// 	ptrglPolygonMode = coreGetProcAddress("glPolygonMode");
// 	ptrglProgramUniform4uiv = coreGetProcAddress("glProgramUniform4uiv");
// 	ptrglGetSubroutineUniformLocation = coreGetProcAddress("glGetSubroutineUniformLocation");
// 	ptrglDrawArraysInstancedBaseInstance = coreGetProcAddress("glDrawArraysInstancedBaseInstance");
// 	ptrglGetVertexAttribPointerv = coreGetProcAddress("glGetVertexAttribPointerv");
// 	ptrglEnableVertexAttribArray = coreGetProcAddress("glEnableVertexAttribArray");
// 	ptrglIsQuery = coreGetProcAddress("glIsQuery");
// 	ptrglProgramUniformMatrix4x2fv = coreGetProcAddress("glProgramUniformMatrix4x2fv");
// 	ptrglStencilFuncSeparate = coreGetProcAddress("glStencilFuncSeparate");
// 	ptrglPixelStorei = coreGetProcAddress("glPixelStorei");
// 	ptrglTexParameterIiv = coreGetProcAddress("glTexParameterIiv");
// 	ptrglGetActiveAtomicCounterBufferiv = coreGetProcAddress("glGetActiveAtomicCounterBufferiv");
// 	ptrglGetActiveSubroutineUniformiv = coreGetProcAddress("glGetActiveSubroutineUniformiv");
// 	ptrglColorMaski = coreGetProcAddress("glColorMaski");
// 	ptrglViewportIndexedfv = coreGetProcAddress("glViewportIndexedfv");
// 	ptrglProgramUniform3ui = coreGetProcAddress("glProgramUniform3ui");
// 	ptrglMultiTexCoordP2ui = coreGetProcAddress("glMultiTexCoordP2ui");
// 	ptrglProgramUniform4f = coreGetProcAddress("glProgramUniform4f");
// 	ptrglColorP3ui = coreGetProcAddress("glColorP3ui");
// 	ptrglBlendEquationSeparate = coreGetProcAddress("glBlendEquationSeparate");
// 	ptrglProgramUniform4i = coreGetProcAddress("glProgramUniform4i");
// 	ptrglFlush = coreGetProcAddress("glFlush");
// 	ptrglFinish = coreGetProcAddress("glFinish");
// 	ptrglUniformMatrix3dv = coreGetProcAddress("glUniformMatrix3dv");
// 	ptrglEndTransformFeedback = coreGetProcAddress("glEndTransformFeedback");
// 	ptrglProgramUniform3d = coreGetProcAddress("glProgramUniform3d");
// 	ptrglGetIntegerv = coreGetProcAddress("glGetIntegerv");
// 	ptrglUniformMatrix4x3fv = coreGetProcAddress("glUniformMatrix4x3fv");
// 	ptrglGetTransformFeedbackVarying = coreGetProcAddress("glGetTransformFeedbackVarying");
// 	ptrglProgramUniform4d = coreGetProcAddress("glProgramUniform4d");
// 	ptrglGetProgramResourceName = coreGetProcAddress("glGetProgramResourceName");
// 	ptrglDeleteProgramPipelines = coreGetProcAddress("glDeleteProgramPipelines");
// 	ptrglDisable = coreGetProcAddress("glDisable");
// 	ptrglVertexAttribLPointer = coreGetProcAddress("glVertexAttribLPointer");
// 	ptrglVertexAttribP3ui = coreGetProcAddress("glVertexAttribP3ui");
// 	ptrglProvokingVertex = coreGetProcAddress("glProvokingVertex");
// 	ptrglDrawArraysInstanced = coreGetProcAddress("glDrawArraysInstanced");
// 	ptrglNormalP3uiv = coreGetProcAddress("glNormalP3uiv");
// 	ptrglUniform2iv = coreGetProcAddress("glUniform2iv");
// 	ptrglProgramUniform1d = coreGetProcAddress("glProgramUniform1d");
// 	ptrglPolygonOffset = coreGetProcAddress("glPolygonOffset");
// 	ptrglUniformMatrix3x4dv = coreGetProcAddress("glUniformMatrix3x4dv");
// 	ptrglTexStorage3D = coreGetProcAddress("glTexStorage3D");
// 	ptrglGetTexParameterIiv = coreGetProcAddress("glGetTexParameterIiv");
// 	ptrglProgramUniformMatrix4fv = coreGetProcAddress("glProgramUniformMatrix4fv");
// 	ptrglVertexAttribP2uiv = coreGetProcAddress("glVertexAttribP2uiv");
// 	ptrglCopyTexSubImage1D = coreGetProcAddress("glCopyTexSubImage1D");
// 	ptrglGetSynciv = coreGetProcAddress("glGetSynciv");
// 	ptrglUniform4f = coreGetProcAddress("glUniform4f");
// 	ptrglProgramUniformMatrix2fv = coreGetProcAddress("glProgramUniformMatrix2fv");
// 	ptrglDeleteShader = coreGetProcAddress("glDeleteShader");
// 	ptrglBindBuffer = coreGetProcAddress("glBindBuffer");
// 	ptrglEnable = coreGetProcAddress("glEnable");
// 	ptrglWaitSync = coreGetProcAddress("glWaitSync");
// 	ptrglBindAttribLocation = coreGetProcAddress("glBindAttribLocation");
// 	ptrglVertexAttribI4iv = coreGetProcAddress("glVertexAttribI4iv");
// 	ptrglGetProgramStageiv = coreGetProcAddress("glGetProgramStageiv");
// 	ptrglTexCoordP2uiv = coreGetProcAddress("glTexCoordP2uiv");
// 	ptrglMultiTexCoordP4ui = coreGetProcAddress("glMultiTexCoordP4ui");
// 	ptrglUniform3iv = coreGetProcAddress("glUniform3iv");
// 	ptrglGetIntegeri_v = coreGetProcAddress("glGetIntegeri_v");
// 	ptrglDrawElementsInstanced = coreGetProcAddress("glDrawElementsInstanced");
// 	ptrglMinSampleShading = coreGetProcAddress("glMinSampleShading");
// 	ptrglGetShaderSource = coreGetProcAddress("glGetShaderSource");
// 	ptrglProgramUniform2uiv = coreGetProcAddress("glProgramUniform2uiv");
// 	ptrglStencilMaskSeparate = coreGetProcAddress("glStencilMaskSeparate");
// 	ptrglScissor = coreGetProcAddress("glScissor");
// 	ptrglGetShaderPrecisionFormat = coreGetProcAddress("glGetShaderPrecisionFormat");
// 	ptrglVertexAttribI4ui = coreGetProcAddress("glVertexAttribI4ui");
// 	ptrglColorP3uiv = coreGetProcAddress("glColorP3uiv");
// 	ptrglGetActiveUniformBlockName = coreGetProcAddress("glGetActiveUniformBlockName");
// 	ptrglTexStorage3DMultisample = coreGetProcAddress("glTexStorage3DMultisample");
// 	ptrglVertexAttribI3uiv = coreGetProcAddress("glVertexAttribI3uiv");
// 	ptrglVertexP2uiv = coreGetProcAddress("glVertexP2uiv");
// 	ptrglMultiDrawElementsIndirect = coreGetProcAddress("glMultiDrawElementsIndirect");
// 	ptrglUniform4uiv = coreGetProcAddress("glUniform4uiv");
// 	ptrglGetTexParameteriv = coreGetProcAddress("glGetTexParameteriv");
// 	ptrglVertexP2ui = coreGetProcAddress("glVertexP2ui");
// 	ptrglGenVertexArrays = coreGetProcAddress("glGenVertexArrays");
// 	ptrglGetProgramResourceLocationIndex = coreGetProcAddress("glGetProgramResourceLocationIndex");
// 	ptrglGetSamplerParameteriv = coreGetProcAddress("glGetSamplerParameteriv");
// 	ptrglDepthRangef = coreGetProcAddress("glDepthRangef");
// 	ptrglVertexAttribI1ui = coreGetProcAddress("glVertexAttribI1ui");
// 	ptrglClearColor = coreGetProcAddress("glClearColor");
// 	ptrglTexParameterIuiv = coreGetProcAddress("glTexParameterIuiv");
// 	ptrglProgramUniform2fv = coreGetProcAddress("glProgramUniform2fv");
// 	ptrglVertexAttribBinding = coreGetProcAddress("glVertexAttribBinding");
// 	ptrglTexCoordP4uiv = coreGetProcAddress("glTexCoordP4uiv");
// 	ptrglPatchParameterfv = coreGetProcAddress("glPatchParameterfv");
// 	ptrglCompressedTexImage2D = coreGetProcAddress("glCompressedTexImage2D");
// 	ptrglActiveTexture = coreGetProcAddress("glActiveTexture");
// 	ptrglGetProgramBinary = coreGetProcAddress("glGetProgramBinary");
// 	ptrglUniform1ui = coreGetProcAddress("glUniform1ui");
// 	ptrglCopyTexImage1D = coreGetProcAddress("glCopyTexImage1D");
// 	ptrglDepthRangeIndexed = coreGetProcAddress("glDepthRangeIndexed");
// 	ptrglCopyTexImage2D = coreGetProcAddress("glCopyTexImage2D");
// 	ptrglScissorArrayv = coreGetProcAddress("glScissorArrayv");
// 	ptrglGetShaderInfoLog = coreGetProcAddress("glGetShaderInfoLog");
// 	ptrglGetMultisamplefv = coreGetProcAddress("glGetMultisamplefv");
// 	ptrglDrawArrays = coreGetProcAddress("glDrawArrays");
// 	ptrglGetActiveUniformBlockiv = coreGetProcAddress("glGetActiveUniformBlockiv");
// 	ptrglUniformMatrix3fv = coreGetProcAddress("glUniformMatrix3fv");
// 	ptrglCreateProgram = coreGetProcAddress("glCreateProgram");
// 	ptrglVertexAttribI1uiv = coreGetProcAddress("glVertexAttribI1uiv");
// 	ptrglGetQueryObjectuiv = coreGetProcAddress("glGetQueryObjectuiv");
// 	ptrglEndQuery = coreGetProcAddress("glEndQuery");
// 	ptrglClearDepth = coreGetProcAddress("glClearDepth");
// 	ptrglGetFloati_v = coreGetProcAddress("glGetFloati_v");
// 	ptrglPointParameterfv = coreGetProcAddress("glPointParameterfv");
// 	ptrglGetQueryObjectiv = coreGetProcAddress("glGetQueryObjectiv");
// 	ptrglGetActiveUniformName = coreGetProcAddress("glGetActiveUniformName");
// 	ptrglNormalP3ui = coreGetProcAddress("glNormalP3ui");
// 	ptrglActiveShaderProgram = coreGetProcAddress("glActiveShaderProgram");
// 	ptrglVertexAttribI3iv = coreGetProcAddress("glVertexAttribI3iv");
// 	ptrglIsFramebuffer = coreGetProcAddress("glIsFramebuffer");
// 	ptrglBindImageTexture = coreGetProcAddress("glBindImageTexture");
// 	ptrglGetActiveUniform = coreGetProcAddress("glGetActiveUniform");
// 	ptrglUniform1i = coreGetProcAddress("glUniform1i");
// 	ptrglObjectPtrLabel = coreGetProcAddress("glObjectPtrLabel");
// 	ptrglUseProgramStages = coreGetProcAddress("glUseProgramStages");
// 	ptrglGetTexLevelParameteriv = coreGetProcAddress("glGetTexLevelParameteriv");
// 	ptrglVertexAttribLFormat = coreGetProcAddress("glVertexAttribLFormat");
// 	ptrglGetSamplerParameterIiv = coreGetProcAddress("glGetSamplerParameterIiv");
// 	ptrglBufferSubData = coreGetProcAddress("glBufferSubData");
// 	ptrglTexSubImage2D = coreGetProcAddress("glTexSubImage2D");
// 	ptrglUniform3ui = coreGetProcAddress("glUniform3ui");
// 	ptrglTexImage2DMultisample = coreGetProcAddress("glTexImage2DMultisample");
// 	ptrglRenderbufferStorageMultisample = coreGetProcAddress("glRenderbufferStorageMultisample");
// 	ptrglIsProgram = coreGetProcAddress("glIsProgram");
// 	ptrglVertexAttribP2ui = coreGetProcAddress("glVertexAttribP2ui");
// 	ptrglVertexAttribI4bv = coreGetProcAddress("glVertexAttribI4bv");
// 	ptrglUniform2i = coreGetProcAddress("glUniform2i");
// 	ptrglDeleteRenderbuffers = coreGetProcAddress("glDeleteRenderbuffers");
// 	ptrglGetBufferPointerv = coreGetProcAddress("glGetBufferPointerv");
// 	ptrglDrawBuffer = coreGetProcAddress("glDrawBuffer");
// 	ptrglMapBuffer = coreGetProcAddress("glMapBuffer");
// 	ptrglDepthRangeArrayv = coreGetProcAddress("glDepthRangeArrayv");
// 	ptrglVertexP3uiv = coreGetProcAddress("glVertexP3uiv");
// 	ptrglFramebufferTexture3D = coreGetProcAddress("glFramebufferTexture3D");
// 	ptrglVertexAttribI4ubv = coreGetProcAddress("glVertexAttribI4ubv");
// 	ptrglUniform4i = coreGetProcAddress("glUniform4i");
// 	ptrglCompressedTexImage3D = coreGetProcAddress("glCompressedTexImage3D");
// 	ptrglDeleteTransformFeedbacks = coreGetProcAddress("glDeleteTransformFeedbacks");
// 	ptrglBeginQuery = coreGetProcAddress("glBeginQuery");
// 	ptrglReadBuffer = coreGetProcAddress("glReadBuffer");
// 	ptrglDebugMessageControl = coreGetProcAddress("glDebugMessageControl");
// 	ptrglGetSubroutineIndex = coreGetProcAddress("glGetSubroutineIndex");
// 	ptrglVertexAttribI2uiv = coreGetProcAddress("glVertexAttribI2uiv");
// 	ptrglMapBufferRange = coreGetProcAddress("glMapBufferRange");
// 	ptrglDrawBuffers = coreGetProcAddress("glDrawBuffers");
// 	ptrglUnmapBuffer = coreGetProcAddress("glUnmapBuffer");
// 	ptrglUniform3d = coreGetProcAddress("glUniform3d");
// 	ptrglGetAttribLocation = coreGetProcAddress("glGetAttribLocation");
// 	ptrglUseProgram = coreGetProcAddress("glUseProgram");
// 	ptrglTexSubImage3D = coreGetProcAddress("glTexSubImage3D");
// 	ptrglUniformMatrix4fv = coreGetProcAddress("glUniformMatrix4fv");
// 	ptrglProgramUniformMatrix3x4dv = coreGetProcAddress("glProgramUniformMatrix3x4dv");
// 	ptrglBlendFuncSeparatei = coreGetProcAddress("glBlendFuncSeparatei");
// 	ptrglVertexAttribP3uiv = coreGetProcAddress("glVertexAttribP3uiv");
// 	ptrglIsRenderbuffer = coreGetProcAddress("glIsRenderbuffer");
// 	ptrglIsProgramPipeline = coreGetProcAddress("glIsProgramPipeline");
// 	ptrglScissorIndexed = coreGetProcAddress("glScissorIndexed");
// 	ptrglBlendColor = coreGetProcAddress("glBlendColor");
// 	ptrglTexParameterf = coreGetProcAddress("glTexParameterf");
// 	ptrglInvalidateTexSubImage = coreGetProcAddress("glInvalidateTexSubImage");
// 	ptrglGetFragDataLocation = coreGetProcAddress("glGetFragDataLocation");
// 	ptrglUniformMatrix2x3fv = coreGetProcAddress("glUniformMatrix2x3fv");
// 	ptrglUniform4d = coreGetProcAddress("glUniform4d");
// 	ptrglClearBufferfv = coreGetProcAddress("glClearBufferfv");
// 	ptrglGetQueryIndexediv = coreGetProcAddress("glGetQueryIndexediv");
// 	ptrglProgramUniform3fv = coreGetProcAddress("glProgramUniform3fv");
// 	ptrglProgramUniform1iv = coreGetProcAddress("glProgramUniform1iv");
// 	ptrglUniform2fv = coreGetProcAddress("glUniform2fv");
// 	ptrglGetProgramInterfaceiv = coreGetProcAddress("glGetProgramInterfaceiv");
// 	ptrglUniformMatrix2dv = coreGetProcAddress("glUniformMatrix2dv");
// 	ptrglIsSync = coreGetProcAddress("glIsSync");
// 	ptrglGetBooleanv = coreGetProcAddress("glGetBooleanv");
// 	ptrglStencilOp = coreGetProcAddress("glStencilOp");
// 	ptrglGetQueryObjectui64v = coreGetProcAddress("glGetQueryObjectui64v");
// 	ptrglProgramUniformMatrix4x3dv = coreGetProcAddress("glProgramUniformMatrix4x3dv");
// 	ptrglGetObjectPtrLabel = coreGetProcAddress("glGetObjectPtrLabel");
// 	ptrglClearBufferfi = coreGetProcAddress("glClearBufferfi");
// 	ptrglProgramUniformMatrix2x4fv = coreGetProcAddress("glProgramUniformMatrix2x4fv");
// 	ptrglUniform3fv = coreGetProcAddress("glUniform3fv");
// 	ptrglVertexP4ui = coreGetProcAddress("glVertexP4ui");
// 	ptrglGetQueryObjecti64v = coreGetProcAddress("glGetQueryObjecti64v");
// 	ptrglCompressedTexImage1D = coreGetProcAddress("glCompressedTexImage1D");
// 	ptrglPrimitiveRestartIndex = coreGetProcAddress("glPrimitiveRestartIndex");
// 	ptrglDrawTransformFeedback = coreGetProcAddress("glDrawTransformFeedback");
// 	ptrglGetDoublev = coreGetProcAddress("glGetDoublev");
// 	ptrglBlendFuncSeparate = coreGetProcAddress("glBlendFuncSeparate");
// 	ptrglVertexAttribP1ui = coreGetProcAddress("glVertexAttribP1ui");
// 	ptrglVertexAttribP4uiv = coreGetProcAddress("glVertexAttribP4uiv");
// 	ptrglBeginTransformFeedback = coreGetProcAddress("glBeginTransformFeedback");
// 	ptrglClearStencil = coreGetProcAddress("glClearStencil");
// 	ptrglDeleteFramebuffers = coreGetProcAddress("glDeleteFramebuffers");
// 	ptrglVertexAttribL1dv = coreGetProcAddress("glVertexAttribL1dv");
// 	ptrglPointParameterf = coreGetProcAddress("glPointParameterf");
// 	ptrglVertexAttribI4sv = coreGetProcAddress("glVertexAttribI4sv");
// 	ptrglVertexAttribL2dv = coreGetProcAddress("glVertexAttribL2dv");
// 	ptrglBlendFunci = coreGetProcAddress("glBlendFunci");
// 	ptrglUniform2ui = coreGetProcAddress("glUniform2ui");
// 	ptrglGetProgramPipelineiv = coreGetProcAddress("glGetProgramPipelineiv");
// 	ptrglSampleMaski = coreGetProcAddress("glSampleMaski");
// 	ptrglGetSamplerParameterIuiv = coreGetProcAddress("glGetSamplerParameterIuiv");
// 	ptrglCopyBufferSubData = coreGetProcAddress("glCopyBufferSubData");
// 	ptrglGetActiveUniformsiv = coreGetProcAddress("glGetActiveUniformsiv");
// 	ptrglGetBufferParameteriv = coreGetProcAddress("glGetBufferParameteriv");
// 	ptrglStencilOpSeparate = coreGetProcAddress("glStencilOpSeparate");
// 	ptrglBlitFramebuffer = coreGetProcAddress("glBlitFramebuffer");
// 	ptrglDepthFunc = coreGetProcAddress("glDepthFunc");
// 	ptrglDetachShader = coreGetProcAddress("glDetachShader");
// 	ptrglClearBufferuiv = coreGetProcAddress("glClearBufferuiv");
// 	ptrglProgramBinary = coreGetProcAddress("glProgramBinary");
// 	ptrglViewportIndexedf = coreGetProcAddress("glViewportIndexedf");
// 	ptrglProgramUniformMatrix3dv = coreGetProcAddress("glProgramUniformMatrix3dv");
// 	ptrglUniformMatrix2x4dv = coreGetProcAddress("glUniformMatrix2x4dv");
// 	ptrglProgramParameteri = coreGetProcAddress("glProgramParameteri");
// 	ptrglMultiTexCoordP3ui = coreGetProcAddress("glMultiTexCoordP3ui");
// 	ptrglUniformMatrix4dv = coreGetProcAddress("glUniformMatrix4dv");
// 	ptrglBindFragDataLocation = coreGetProcAddress("glBindFragDataLocation");
// 	ptrglGetAttachedShaders = coreGetProcAddress("glGetAttachedShaders");
// 	ptrglVertexP4uiv = coreGetProcAddress("glVertexP4uiv");
// 	ptrglReadPixels = coreGetProcAddress("glReadPixels");
// 	ptrglVertexAttribIFormat = coreGetProcAddress("glVertexAttribIFormat");
// 	ptrglTexStorage1D = coreGetProcAddress("glTexStorage1D");
// 	ptrglProgramUniform4dv = coreGetProcAddress("glProgramUniform4dv");
// 	ptrglVertexBindingDivisor = coreGetProcAddress("glVertexBindingDivisor");
// 	ptrglDrawRangeElementsBaseVertex = coreGetProcAddress("glDrawRangeElementsBaseVertex");
// 	ptrglGetUniformiv = coreGetProcAddress("glGetUniformiv");
// 	ptrglVertexAttribL3dv = coreGetProcAddress("glVertexAttribL3dv");
// 	ptrglInvalidateBufferData = coreGetProcAddress("glInvalidateBufferData");
// 	ptrglResumeTransformFeedback = coreGetProcAddress("glResumeTransformFeedback");
// 	ptrglUniformMatrix2x3dv = coreGetProcAddress("glUniformMatrix2x3dv");
// 	ptrglEndQueryIndexed = coreGetProcAddress("glEndQueryIndexed");
// 	ptrglGetVertexAttribIuiv = coreGetProcAddress("glGetVertexAttribIuiv");
// 	ptrglBeginQueryIndexed = coreGetProcAddress("glBeginQueryIndexed");
// 	ptrglUniformMatrix4x3dv = coreGetProcAddress("glUniformMatrix4x3dv");
// 	ptrglTexBufferRange = coreGetProcAddress("glTexBufferRange");
// 	ptrglGetVertexAttribLdv = coreGetProcAddress("glGetVertexAttribLdv");
// 	ptrglDrawTransformFeedbackStream = coreGetProcAddress("glDrawTransformFeedbackStream");
// 	ptrglBindFramebuffer = coreGetProcAddress("glBindFramebuffer");
// 	ptrglDeleteProgram = coreGetProcAddress("glDeleteProgram");
// 	ptrglScissorIndexedv = coreGetProcAddress("glScissorIndexedv");
// 	ptrglUniformMatrix3x2dv = coreGetProcAddress("glUniformMatrix3x2dv");
// 	ptrglIsTexture = coreGetProcAddress("glIsTexture");
// 	ptrglSamplerParameterIiv = coreGetProcAddress("glSamplerParameterIiv");
// 	ptrglProgramUniform4fv = coreGetProcAddress("glProgramUniform4fv");
// 	ptrglBlendEquationi = coreGetProcAddress("glBlendEquationi");
// 	ptrglDeleteSync = coreGetProcAddress("glDeleteSync");
// 	ptrglProgramUniform3iv = coreGetProcAddress("glProgramUniform3iv");
// 	ptrglInvalidateTexImage = coreGetProcAddress("glInvalidateTexImage");
// 	ptrglViewportArrayv = coreGetProcAddress("glViewportArrayv");
// 	ptrglTexSubImage1D = coreGetProcAddress("glTexSubImage1D");
// 	ptrglCreateShaderProgramv = coreGetProcAddress("glCreateShaderProgramv");
// 	ptrglGenRenderbuffers = coreGetProcAddress("glGenRenderbuffers");
// 	ptrglProgramUniformMatrix2dv = coreGetProcAddress("glProgramUniformMatrix2dv");
// 	ptrglProgramUniform2ui = coreGetProcAddress("glProgramUniform2ui");
// 	ptrglCopyTexSubImage2D = coreGetProcAddress("glCopyTexSubImage2D");
// 	ptrglMultiDrawElementsBaseVertex = coreGetProcAddress("glMultiDrawElementsBaseVertex");
// 	ptrglGetCompressedTexImage = coreGetProcAddress("glGetCompressedTexImage");
// 	ptrglVertexAttribI3ui = coreGetProcAddress("glVertexAttribI3ui");
// 	ptrglUniformMatrix3x4fv = coreGetProcAddress("glUniformMatrix3x4fv");
// 	ptrglVertexAttribI1iv = coreGetProcAddress("glVertexAttribI1iv");
// 	ptrglGetBufferSubData = coreGetProcAddress("glGetBufferSubData");
// 	ptrglTexImage1D = coreGetProcAddress("glTexImage1D");
// 	ptrglFramebufferTexture = coreGetProcAddress("glFramebufferTexture");
// 	ptrglVertexAttribI1i = coreGetProcAddress("glVertexAttribI1i");
// 	ptrglGetVertexAttribIiv = coreGetProcAddress("glGetVertexAttribIiv");
// 	ptrglGetBufferParameteri64v = coreGetProcAddress("glGetBufferParameteri64v");
// 	ptrglBeginConditionalRender = coreGetProcAddress("glBeginConditionalRender");
// 	ptrglGetInternalformati64v = coreGetProcAddress("glGetInternalformati64v");
// 	ptrglVertexAttribI2ui = coreGetProcAddress("glVertexAttribI2ui");
// 	ptrglBlendEquation = coreGetProcAddress("glBlendEquation");
// 	ptrglEnablei = coreGetProcAddress("glEnablei");
// 	ptrglGetStringi = coreGetProcAddress("glGetStringi");
// 	ptrglProgramUniformMatrix3x2dv = coreGetProcAddress("glProgramUniformMatrix3x2dv");
// 	ptrglDrawTransformFeedbackStreamInstanced = coreGetProcAddress("glDrawTransformFeedbackStreamInstanced");
// 	ptrglGetActiveSubroutineUniformName = coreGetProcAddress("glGetActiveSubroutineUniformName");
// 	ptrglGetInteger64i_v = coreGetProcAddress("glGetInteger64i_v");
// 	ptrglDebugMessageInsert = coreGetProcAddress("glDebugMessageInsert");
// 	ptrglProgramUniform1f = coreGetProcAddress("glProgramUniform1f");
// 	ptrglGetSamplerParameterfv = coreGetProcAddress("glGetSamplerParameterfv");
// 	ptrglTexCoordP4ui = coreGetProcAddress("glTexCoordP4ui");
// 	ptrglProgramUniform1uiv = coreGetProcAddress("glProgramUniform1uiv");
// 	ptrglDeleteSamplers = coreGetProcAddress("glDeleteSamplers");
// 	ptrglIsShader = coreGetProcAddress("glIsShader");
// 	ptrglCopyTexSubImage3D = coreGetProcAddress("glCopyTexSubImage3D");
// 	ptrglGetUniformuiv = coreGetProcAddress("glGetUniformuiv");
// 	ptrglVertexAttribL2d = coreGetProcAddress("glVertexAttribL2d");
// 	ptrglSecondaryColorP3ui = coreGetProcAddress("glSecondaryColorP3ui");
// 	ptrglIsSampler = coreGetProcAddress("glIsSampler");
// 	ptrglUniform2dv = coreGetProcAddress("glUniform2dv");
// 	ptrglDepthMask = coreGetProcAddress("glDepthMask");
// 	ptrglPushDebugGroup = coreGetProcAddress("glPushDebugGroup");
// /* a cheap way to determine if (at least) v3.3 core profile (or higher) seems to be supported: */
// 	return ((ptrglBindSampler == NULL) ? 0 : 1);
// }
import "C"
import "fmt"
import "unsafe"

type (
	Ptr      unsafe.Pointer
	Clampf   C.GLclampf
	Char     C.GLchar
	Clampd   C.GLclampd
	Int      C.GLint
	Uint     C.GLuint
	Short    C.GLshort
	Sync     C.GLsync
	Ubyte    C.GLubyte
	Double   C.GLdouble
	Int64    C.GLint64
	Intptr   C.GLintptr
	Ushort   C.GLushort
	Sizeiptr C.GLsizeiptr
	Boolean  C.GLboolean
	Enum     C.GLenum
	Uint64   C.GLuint64
	Sizei    C.GLsizei
	Half     C.GLhalf
	Byte     C.GLbyte
	Bitfield C.GLbitfield
	Float    C.GLfloat
)

//	Provides methods indicating whether a given GL function can be called within the current GL context.
type GlSupport struct{}

//	Provides method wrappers that call GL functions and check for any errors immediately afterwards.
//	Not recommended for use in a real-time loop, but useful for "dirty debugging" or one-off/setup tasks.
type GlTry struct{}

//	Provides utility methods for working with this OpenGL binding package.
type GlUtil struct{}

//	Provides methods indicating whether a given GL function can be called within the current GL context.
var Supports GlSupport

//	Provides method wrappers that call GL functions and check for any errors immediately afterwards.
//	Not recommended for use in a real-time loop, but useful for "dirty debugging" or one-off/setup tasks.
var Try GlTry

//	Provides utility methods for working with this OpenGL binding package.
var Util GlUtil

//	Go string to GL string.
func (_ GlUtil) CString(str string) *Char {
	return (*Char)(C.CString(str))
}

//	Allocates a GL string.
func (_ GlUtil) CStringAlloc(length Sizei) *Char {
	return (*Char)(C.malloc(C.size_t(length)))
}

//	Frees GL string.
func (_ GlUtil) CStringFree(str *Char) {
	C.free(unsafe.Pointer(str))
}

//	GL string (GLchar*) to Go string.
func (_ GlUtil) StringFromChar(str *Char) string {
	return C.GoString((*C.char)(str))
}

//	GL string (GLubyte*) to Go string.
func (_ GlUtil) StringFromUbyte(str *Ubyte) string {
	return C.GoString((*C.char)(unsafe.Pointer(str)))
}

//	GL string (GLchar*) with length to Go string.
func (_ GlUtil) StringN(str *Char, length Sizei) string {
	return C.GoStringN((*C.char)(str), C.int(length))
}

//	Converts a list of Go strings to a slice of GL strings.
func (_ GlUtil) CStringArray(strs ...string) []*Char {
	strSlice := make([]*Char, len(strs))
	for i, s := range strs {
		strSlice[i] = (*Char)(C.CString(s))
	}
	return strSlice
}

//	Free GL string slice allocated by GLStringArray().
func (_ GlUtil) CStringArrayFree(strs []*Char) {
	for _, s := range strs {
		C.free(unsafe.Pointer(s))
	}
}

//	Add offset to a pointer. Useful for VertexAttribPointer and friends.
func (_ GlUtil) PtrOffset(p Ptr, o uintptr) Ptr {
	return Ptr(uintptr(p) + o)
}

//	Returns all error flags collected from GetError() until the latter yields NO_ERROR.
func (_ GlUtil) ErrorFlags() (flags []Enum) {
	for e := GetError(); e != NO_ERROR; e = GetError() {
		flags = append(flags, e)
	}
	return
}

//	If any errors have been recorded since the last (direct or indirect) GetError() call, returns a single error value informing on those and including the specified message, if any.
func (_ GlUtil) Error(msgFmt string, fmtArgs ...interface{}) (err error) {
	if flags := Util.ErrorFlags(); len(flags) > 0 {
		for _, e := range flags {
			msgFmt = Util.EnumName(e) + " " + msgFmt
		}
		err = fmt.Errorf(msgFmt, fmtArgs)
	}
	return
}

//	Initializes this OpenGL binding after having created your GL context (with SDL, GLFW or in some other way).
func (_ GlUtil) Init() bool {
	return (C.coreGetProcAddresses() == 1)
}

//	Returns whether the current GL context supports calling the ObjectPtrLabel() function.
func (_ GlSupport) ObjectPtrLabel() bool {
	return C.cancallObjectPtrLabel() == 1
}

//	If Supports.ObjectPtrLabel() is true, calls ObjectPtrLabel() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ObjectPtrLabel(ptr Ptr, length Sizei, label *Char) (err error) {
	if !Supports.ObjectPtrLabel() {
		err = fmt.Errorf("ObjectPtrLabel() function call not supported by current GL context.")
	} else {
		ObjectPtrLabel(ptr, length, label)
		err = Util.Error("ObjectPtrLabel()")
	}
	return
}

//	Returns whether the current GL context supports calling the UseProgramStages() function.
func (_ GlSupport) UseProgramStages() bool {
	return C.cancallUseProgramStages() == 1
}

//	If Supports.UseProgramStages() is true, calls UseProgramStages() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UseProgramStages(pipeline Uint, stages Bitfield, program Uint) (err error) {
	if !Supports.UseProgramStages() {
		err = fmt.Errorf("UseProgramStages() function call not supported by current GL context.")
	} else {
		UseProgramStages(pipeline, stages, program)
		err = Util.Error("UseProgramStages()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetTexLevelParameteriv() function.
func (_ GlSupport) GetTexLevelParameteriv() bool {
	return C.cancallGetTexLevelParameteriv() == 1
}

//	If Supports.GetTexLevelParameteriv() is true, calls GetTexLevelParameteriv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetTexLevelParameteriv(target Enum, level Int, pname Enum, params *Int) (err error) {
	if !Supports.GetTexLevelParameteriv() {
		err = fmt.Errorf("GetTexLevelParameteriv() function call not supported by current GL context.")
	} else {
		GetTexLevelParameteriv(target, level, pname, params)
		err = Util.Error("GetTexLevelParameteriv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribLFormat() function.
func (_ GlSupport) VertexAttribLFormat() bool {
	return C.cancallVertexAttribLFormat() == 1
}

//	If Supports.VertexAttribLFormat() is true, calls VertexAttribLFormat() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribLFormat(attribindex Uint, size Int, type_ Enum, relativeoffset Uint) (err error) {
	if !Supports.VertexAttribLFormat() {
		err = fmt.Errorf("VertexAttribLFormat() function call not supported by current GL context.")
	} else {
		VertexAttribLFormat(attribindex, size, type_, relativeoffset)
		err = Util.Error("VertexAttribLFormat()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetSamplerParameterIiv() function.
func (_ GlSupport) GetSamplerParameterIiv() bool {
	return C.cancallGetSamplerParameterIiv() == 1
}

//	If Supports.GetSamplerParameterIiv() is true, calls GetSamplerParameterIiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetSamplerParameterIiv(sampler Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetSamplerParameterIiv() {
		err = fmt.Errorf("GetSamplerParameterIiv() function call not supported by current GL context.")
	} else {
		GetSamplerParameterIiv(sampler, pname, params)
		err = Util.Error("GetSamplerParameterIiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the BufferSubData() function.
func (_ GlSupport) BufferSubData() bool {
	return C.cancallBufferSubData() == 1
}

//	If Supports.BufferSubData() is true, calls BufferSubData() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BufferSubData(target Enum, offset Intptr, size Sizeiptr, data Ptr) (err error) {
	if !Supports.BufferSubData() {
		err = fmt.Errorf("BufferSubData() function call not supported by current GL context.")
	} else {
		BufferSubData(target, offset, size, data)
		err = Util.Error("BufferSubData()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexSubImage2D() function.
func (_ GlSupport) TexSubImage2D() bool {
	return C.cancallTexSubImage2D() == 1
}

//	If Supports.TexSubImage2D() is true, calls TexSubImage2D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Ptr) (err error) {
	if !Supports.TexSubImage2D() {
		err = fmt.Errorf("TexSubImage2D() function call not supported by current GL context.")
	} else {
		TexSubImage2D(target, level, xoffset, yoffset, width, height, format, type_, pixels)
		err = Util.Error("TexSubImage2D()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform3ui() function.
func (_ GlSupport) Uniform3ui() bool {
	return C.cancallUniform3ui() == 1
}

//	If Supports.Uniform3ui() is true, calls Uniform3ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform3ui(location Int, v0 Uint, v1 Uint, v2 Uint) (err error) {
	if !Supports.Uniform3ui() {
		err = fmt.Errorf("Uniform3ui() function call not supported by current GL context.")
	} else {
		Uniform3ui(location, v0, v1, v2)
		err = Util.Error("Uniform3ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexImage2DMultisample() function.
func (_ GlSupport) TexImage2DMultisample() bool {
	return C.cancallTexImage2DMultisample() == 1
}

//	If Supports.TexImage2DMultisample() is true, calls TexImage2DMultisample() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexImage2DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, fixedsamplelocations Boolean) (err error) {
	if !Supports.TexImage2DMultisample() {
		err = fmt.Errorf("TexImage2DMultisample() function call not supported by current GL context.")
	} else {
		TexImage2DMultisample(target, samples, internalformat, width, height, fixedsamplelocations)
		err = Util.Error("TexImage2DMultisample()")
	}
	return
}

//	Returns whether the current GL context supports calling the RenderbufferStorageMultisample() function.
func (_ GlSupport) RenderbufferStorageMultisample() bool {
	return C.cancallRenderbufferStorageMultisample() == 1
}

//	If Supports.RenderbufferStorageMultisample() is true, calls RenderbufferStorageMultisample() and yields the err returned by Util.Error(), if any.
func (_ GlTry) RenderbufferStorageMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei) (err error) {
	if !Supports.RenderbufferStorageMultisample() {
		err = fmt.Errorf("RenderbufferStorageMultisample() function call not supported by current GL context.")
	} else {
		RenderbufferStorageMultisample(target, samples, internalformat, width, height)
		err = Util.Error("RenderbufferStorageMultisample()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsProgram() function.
func (_ GlSupport) IsProgram() bool {
	return C.cancallIsProgram() == 1
}

//	If Supports.IsProgram() is true, calls IsProgram() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsProgram(program Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsProgram() {
		err = fmt.Errorf("IsProgram() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsProgram(program)
		err = Util.Error("IsProgram()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribP2ui() function.
func (_ GlSupport) VertexAttribP2ui() bool {
	return C.cancallVertexAttribP2ui() == 1
}

//	If Supports.VertexAttribP2ui() is true, calls VertexAttribP2ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribP2ui(index Uint, type_ Enum, normalized Boolean, value Uint) (err error) {
	if !Supports.VertexAttribP2ui() {
		err = fmt.Errorf("VertexAttribP2ui() function call not supported by current GL context.")
	} else {
		VertexAttribP2ui(index, type_, normalized, value)
		err = Util.Error("VertexAttribP2ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI4bv() function.
func (_ GlSupport) VertexAttribI4bv() bool {
	return C.cancallVertexAttribI4bv() == 1
}

//	If Supports.VertexAttribI4bv() is true, calls VertexAttribI4bv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI4bv(index Uint, v *Byte) (err error) {
	if !Supports.VertexAttribI4bv() {
		err = fmt.Errorf("VertexAttribI4bv() function call not supported by current GL context.")
	} else {
		VertexAttribI4bv(index, v)
		err = Util.Error("VertexAttribI4bv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform2i() function.
func (_ GlSupport) Uniform2i() bool {
	return C.cancallUniform2i() == 1
}

//	If Supports.Uniform2i() is true, calls Uniform2i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform2i(location Int, v0 Int, v1 Int) (err error) {
	if !Supports.Uniform2i() {
		err = fmt.Errorf("Uniform2i() function call not supported by current GL context.")
	} else {
		Uniform2i(location, v0, v1)
		err = Util.Error("Uniform2i()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteRenderbuffers() function.
func (_ GlSupport) DeleteRenderbuffers() bool {
	return C.cancallDeleteRenderbuffers() == 1
}

//	If Supports.DeleteRenderbuffers() is true, calls DeleteRenderbuffers() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteRenderbuffers(n Sizei, renderbuffers *Uint) (err error) {
	if !Supports.DeleteRenderbuffers() {
		err = fmt.Errorf("DeleteRenderbuffers() function call not supported by current GL context.")
	} else {
		DeleteRenderbuffers(n, renderbuffers)
		err = Util.Error("DeleteRenderbuffers()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetBufferPointerv() function.
func (_ GlSupport) GetBufferPointerv() bool {
	return C.cancallGetBufferPointerv() == 1
}

//	If Supports.GetBufferPointerv() is true, calls GetBufferPointerv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetBufferPointerv(target Enum, pname Enum, params *Ptr) (err error) {
	if !Supports.GetBufferPointerv() {
		err = fmt.Errorf("GetBufferPointerv() function call not supported by current GL context.")
	} else {
		GetBufferPointerv(target, pname, params)
		err = Util.Error("GetBufferPointerv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawBuffer() function.
func (_ GlSupport) DrawBuffer() bool {
	return C.cancallDrawBuffer() == 1
}

//	If Supports.DrawBuffer() is true, calls DrawBuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawBuffer(mode Enum) (err error) {
	if !Supports.DrawBuffer() {
		err = fmt.Errorf("DrawBuffer() function call not supported by current GL context.")
	} else {
		DrawBuffer(mode)
		err = Util.Error("DrawBuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the MapBuffer() function.
func (_ GlSupport) MapBuffer() bool {
	return C.cancallMapBuffer() == 1
}

//	If Supports.MapBuffer() is true, calls MapBuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MapBuffer(target Enum, access Enum) (err error, tryRetVal__ *Ptr) {
	if !Supports.MapBuffer() {
		err = fmt.Errorf("MapBuffer() function call not supported by current GL context.")
	} else {
		tryRetVal__ = MapBuffer(target, access)
		err = Util.Error("MapBuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the DepthRangeArrayv() function.
func (_ GlSupport) DepthRangeArrayv() bool {
	return C.cancallDepthRangeArrayv() == 1
}

//	If Supports.DepthRangeArrayv() is true, calls DepthRangeArrayv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DepthRangeArrayv(first Uint, count Sizei, v *Double) (err error) {
	if !Supports.DepthRangeArrayv() {
		err = fmt.Errorf("DepthRangeArrayv() function call not supported by current GL context.")
	} else {
		DepthRangeArrayv(first, count, v)
		err = Util.Error("DepthRangeArrayv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexP3uiv() function.
func (_ GlSupport) VertexP3uiv() bool {
	return C.cancallVertexP3uiv() == 1
}

//	If Supports.VertexP3uiv() is true, calls VertexP3uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexP3uiv(type_ Enum, value *Uint) (err error) {
	if !Supports.VertexP3uiv() {
		err = fmt.Errorf("VertexP3uiv() function call not supported by current GL context.")
	} else {
		VertexP3uiv(type_, value)
		err = Util.Error("VertexP3uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the FramebufferTexture3D() function.
func (_ GlSupport) FramebufferTexture3D() bool {
	return C.cancallFramebufferTexture3D() == 1
}

//	If Supports.FramebufferTexture3D() is true, calls FramebufferTexture3D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) FramebufferTexture3D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int, zoffset Int) (err error) {
	if !Supports.FramebufferTexture3D() {
		err = fmt.Errorf("FramebufferTexture3D() function call not supported by current GL context.")
	} else {
		FramebufferTexture3D(target, attachment, textarget, texture, level, zoffset)
		err = Util.Error("FramebufferTexture3D()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI4ubv() function.
func (_ GlSupport) VertexAttribI4ubv() bool {
	return C.cancallVertexAttribI4ubv() == 1
}

//	If Supports.VertexAttribI4ubv() is true, calls VertexAttribI4ubv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI4ubv(index Uint, v *Ubyte) (err error) {
	if !Supports.VertexAttribI4ubv() {
		err = fmt.Errorf("VertexAttribI4ubv() function call not supported by current GL context.")
	} else {
		VertexAttribI4ubv(index, v)
		err = Util.Error("VertexAttribI4ubv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform4i() function.
func (_ GlSupport) Uniform4i() bool {
	return C.cancallUniform4i() == 1
}

//	If Supports.Uniform4i() is true, calls Uniform4i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform4i(location Int, v0 Int, v1 Int, v2 Int, v3 Int) (err error) {
	if !Supports.Uniform4i() {
		err = fmt.Errorf("Uniform4i() function call not supported by current GL context.")
	} else {
		Uniform4i(location, v0, v1, v2, v3)
		err = Util.Error("Uniform4i()")
	}
	return
}

//	Returns whether the current GL context supports calling the CompressedTexImage3D() function.
func (_ GlSupport) CompressedTexImage3D() bool {
	return C.cancallCompressedTexImage3D() == 1
}

//	If Supports.CompressedTexImage3D() is true, calls CompressedTexImage3D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CompressedTexImage3D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, depth Sizei, border Int, imageSize Sizei, data Ptr) (err error) {
	if !Supports.CompressedTexImage3D() {
		err = fmt.Errorf("CompressedTexImage3D() function call not supported by current GL context.")
	} else {
		CompressedTexImage3D(target, level, internalformat, width, height, depth, border, imageSize, data)
		err = Util.Error("CompressedTexImage3D()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteTransformFeedbacks() function.
func (_ GlSupport) DeleteTransformFeedbacks() bool {
	return C.cancallDeleteTransformFeedbacks() == 1
}

//	If Supports.DeleteTransformFeedbacks() is true, calls DeleteTransformFeedbacks() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteTransformFeedbacks(n Sizei, ids *Uint) (err error) {
	if !Supports.DeleteTransformFeedbacks() {
		err = fmt.Errorf("DeleteTransformFeedbacks() function call not supported by current GL context.")
	} else {
		DeleteTransformFeedbacks(n, ids)
		err = Util.Error("DeleteTransformFeedbacks()")
	}
	return
}

//	Returns whether the current GL context supports calling the BeginQuery() function.
func (_ GlSupport) BeginQuery() bool {
	return C.cancallBeginQuery() == 1
}

//	If Supports.BeginQuery() is true, calls BeginQuery() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BeginQuery(target Enum, id Uint) (err error) {
	if !Supports.BeginQuery() {
		err = fmt.Errorf("BeginQuery() function call not supported by current GL context.")
	} else {
		BeginQuery(target, id)
		err = Util.Error("BeginQuery()")
	}
	return
}

//	Returns whether the current GL context supports calling the ReadBuffer() function.
func (_ GlSupport) ReadBuffer() bool {
	return C.cancallReadBuffer() == 1
}

//	If Supports.ReadBuffer() is true, calls ReadBuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ReadBuffer(mode Enum) (err error) {
	if !Supports.ReadBuffer() {
		err = fmt.Errorf("ReadBuffer() function call not supported by current GL context.")
	} else {
		ReadBuffer(mode)
		err = Util.Error("ReadBuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the DebugMessageControl() function.
func (_ GlSupport) DebugMessageControl() bool {
	return C.cancallDebugMessageControl() == 1
}

//	If Supports.DebugMessageControl() is true, calls DebugMessageControl() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DebugMessageControl(source Enum, type_ Enum, severity Enum, count Sizei, ids *Uint, enabled Boolean) (err error) {
	if !Supports.DebugMessageControl() {
		err = fmt.Errorf("DebugMessageControl() function call not supported by current GL context.")
	} else {
		DebugMessageControl(source, type_, severity, count, ids, enabled)
		err = Util.Error("DebugMessageControl()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetSubroutineIndex() function.
func (_ GlSupport) GetSubroutineIndex() bool {
	return C.cancallGetSubroutineIndex() == 1
}

//	If Supports.GetSubroutineIndex() is true, calls GetSubroutineIndex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetSubroutineIndex(program Uint, shadertype Enum, name *Char) (err error, tryRetVal__ Uint) {
	if !Supports.GetSubroutineIndex() {
		err = fmt.Errorf("GetSubroutineIndex() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetSubroutineIndex(program, shadertype, name)
		err = Util.Error("GetSubroutineIndex()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI2uiv() function.
func (_ GlSupport) VertexAttribI2uiv() bool {
	return C.cancallVertexAttribI2uiv() == 1
}

//	If Supports.VertexAttribI2uiv() is true, calls VertexAttribI2uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI2uiv(index Uint, v *Uint) (err error) {
	if !Supports.VertexAttribI2uiv() {
		err = fmt.Errorf("VertexAttribI2uiv() function call not supported by current GL context.")
	} else {
		VertexAttribI2uiv(index, v)
		err = Util.Error("VertexAttribI2uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the MapBufferRange() function.
func (_ GlSupport) MapBufferRange() bool {
	return C.cancallMapBufferRange() == 1
}

//	If Supports.MapBufferRange() is true, calls MapBufferRange() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MapBufferRange(target Enum, offset Intptr, length Sizeiptr, access Bitfield) (err error, tryRetVal__ *Ptr) {
	if !Supports.MapBufferRange() {
		err = fmt.Errorf("MapBufferRange() function call not supported by current GL context.")
	} else {
		tryRetVal__ = MapBufferRange(target, offset, length, access)
		err = Util.Error("MapBufferRange()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawBuffers() function.
func (_ GlSupport) DrawBuffers() bool {
	return C.cancallDrawBuffers() == 1
}

//	If Supports.DrawBuffers() is true, calls DrawBuffers() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawBuffers(n Sizei, bufs *Enum) (err error) {
	if !Supports.DrawBuffers() {
		err = fmt.Errorf("DrawBuffers() function call not supported by current GL context.")
	} else {
		DrawBuffers(n, bufs)
		err = Util.Error("DrawBuffers()")
	}
	return
}

//	Returns whether the current GL context supports calling the UnmapBuffer() function.
func (_ GlSupport) UnmapBuffer() bool {
	return C.cancallUnmapBuffer() == 1
}

//	If Supports.UnmapBuffer() is true, calls UnmapBuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UnmapBuffer(target Enum) (err error, tryRetVal__ Boolean) {
	if !Supports.UnmapBuffer() {
		err = fmt.Errorf("UnmapBuffer() function call not supported by current GL context.")
	} else {
		tryRetVal__ = UnmapBuffer(target)
		err = Util.Error("UnmapBuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform3d() function.
func (_ GlSupport) Uniform3d() bool {
	return C.cancallUniform3d() == 1
}

//	If Supports.Uniform3d() is true, calls Uniform3d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform3d(location Int, x Double, y Double, z Double) (err error) {
	if !Supports.Uniform3d() {
		err = fmt.Errorf("Uniform3d() function call not supported by current GL context.")
	} else {
		Uniform3d(location, x, y, z)
		err = Util.Error("Uniform3d()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetAttribLocation() function.
func (_ GlSupport) GetAttribLocation() bool {
	return C.cancallGetAttribLocation() == 1
}

//	If Supports.GetAttribLocation() is true, calls GetAttribLocation() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetAttribLocation(program Uint, name *Char) (err error, tryRetVal__ Int) {
	if !Supports.GetAttribLocation() {
		err = fmt.Errorf("GetAttribLocation() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetAttribLocation(program, name)
		err = Util.Error("GetAttribLocation()")
	}
	return
}

//	Returns whether the current GL context supports calling the UseProgram() function.
func (_ GlSupport) UseProgram() bool {
	return C.cancallUseProgram() == 1
}

//	If Supports.UseProgram() is true, calls UseProgram() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UseProgram(program Uint) (err error) {
	if !Supports.UseProgram() {
		err = fmt.Errorf("UseProgram() function call not supported by current GL context.")
	} else {
		UseProgram(program)
		err = Util.Error("UseProgram()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexSubImage3D() function.
func (_ GlSupport) TexSubImage3D() bool {
	return C.cancallTexSubImage3D() == 1
}

//	If Supports.TexSubImage3D() is true, calls TexSubImage3D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, type_ Enum, pixels Ptr) (err error) {
	if !Supports.TexSubImage3D() {
		err = fmt.Errorf("TexSubImage3D() function call not supported by current GL context.")
	} else {
		TexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type_, pixels)
		err = Util.Error("TexSubImage3D()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix4fv() function.
func (_ GlSupport) UniformMatrix4fv() bool {
	return C.cancallUniformMatrix4fv() == 1
}

//	If Supports.UniformMatrix4fv() is true, calls UniformMatrix4fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix4fv(location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.UniformMatrix4fv() {
		err = fmt.Errorf("UniformMatrix4fv() function call not supported by current GL context.")
	} else {
		UniformMatrix4fv(location, count, transpose, value)
		err = Util.Error("UniformMatrix4fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix3x4dv() function.
func (_ GlSupport) ProgramUniformMatrix3x4dv() bool {
	return C.cancallProgramUniformMatrix3x4dv() == 1
}

//	If Supports.ProgramUniformMatrix3x4dv() is true, calls ProgramUniformMatrix3x4dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix3x4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.ProgramUniformMatrix3x4dv() {
		err = fmt.Errorf("ProgramUniformMatrix3x4dv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix3x4dv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix3x4dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the BlendFuncSeparatei() function.
func (_ GlSupport) BlendFuncSeparatei() bool {
	return C.cancallBlendFuncSeparatei() == 1
}

//	If Supports.BlendFuncSeparatei() is true, calls BlendFuncSeparatei() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BlendFuncSeparatei(buf Uint, srcRGB Enum, dstRGB Enum, srcAlpha Enum, dstAlpha Enum) (err error) {
	if !Supports.BlendFuncSeparatei() {
		err = fmt.Errorf("BlendFuncSeparatei() function call not supported by current GL context.")
	} else {
		BlendFuncSeparatei(buf, srcRGB, dstRGB, srcAlpha, dstAlpha)
		err = Util.Error("BlendFuncSeparatei()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribP3uiv() function.
func (_ GlSupport) VertexAttribP3uiv() bool {
	return C.cancallVertexAttribP3uiv() == 1
}

//	If Supports.VertexAttribP3uiv() is true, calls VertexAttribP3uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribP3uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) (err error) {
	if !Supports.VertexAttribP3uiv() {
		err = fmt.Errorf("VertexAttribP3uiv() function call not supported by current GL context.")
	} else {
		VertexAttribP3uiv(index, type_, normalized, value)
		err = Util.Error("VertexAttribP3uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsRenderbuffer() function.
func (_ GlSupport) IsRenderbuffer() bool {
	return C.cancallIsRenderbuffer() == 1
}

//	If Supports.IsRenderbuffer() is true, calls IsRenderbuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsRenderbuffer(renderbuffer Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsRenderbuffer() {
		err = fmt.Errorf("IsRenderbuffer() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsRenderbuffer(renderbuffer)
		err = Util.Error("IsRenderbuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsProgramPipeline() function.
func (_ GlSupport) IsProgramPipeline() bool {
	return C.cancallIsProgramPipeline() == 1
}

//	If Supports.IsProgramPipeline() is true, calls IsProgramPipeline() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsProgramPipeline(pipeline Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsProgramPipeline() {
		err = fmt.Errorf("IsProgramPipeline() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsProgramPipeline(pipeline)
		err = Util.Error("IsProgramPipeline()")
	}
	return
}

//	Returns whether the current GL context supports calling the ScissorIndexed() function.
func (_ GlSupport) ScissorIndexed() bool {
	return C.cancallScissorIndexed() == 1
}

//	If Supports.ScissorIndexed() is true, calls ScissorIndexed() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ScissorIndexed(index Uint, left Int, bottom Int, width Sizei, height Sizei) (err error) {
	if !Supports.ScissorIndexed() {
		err = fmt.Errorf("ScissorIndexed() function call not supported by current GL context.")
	} else {
		ScissorIndexed(index, left, bottom, width, height)
		err = Util.Error("ScissorIndexed()")
	}
	return
}

//	Returns whether the current GL context supports calling the BlendColor() function.
func (_ GlSupport) BlendColor() bool {
	return C.cancallBlendColor() == 1
}

//	If Supports.BlendColor() is true, calls BlendColor() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BlendColor(red Float, green Float, blue Float, alpha Float) (err error) {
	if !Supports.BlendColor() {
		err = fmt.Errorf("BlendColor() function call not supported by current GL context.")
	} else {
		BlendColor(red, green, blue, alpha)
		err = Util.Error("BlendColor()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexParameterf() function.
func (_ GlSupport) TexParameterf() bool {
	return C.cancallTexParameterf() == 1
}

//	If Supports.TexParameterf() is true, calls TexParameterf() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexParameterf(target Enum, pname Enum, param Float) (err error) {
	if !Supports.TexParameterf() {
		err = fmt.Errorf("TexParameterf() function call not supported by current GL context.")
	} else {
		TexParameterf(target, pname, param)
		err = Util.Error("TexParameterf()")
	}
	return
}

//	Returns whether the current GL context supports calling the InvalidateTexSubImage() function.
func (_ GlSupport) InvalidateTexSubImage() bool {
	return C.cancallInvalidateTexSubImage() == 1
}

//	If Supports.InvalidateTexSubImage() is true, calls InvalidateTexSubImage() and yields the err returned by Util.Error(), if any.
func (_ GlTry) InvalidateTexSubImage(texture Uint, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei) (err error) {
	if !Supports.InvalidateTexSubImage() {
		err = fmt.Errorf("InvalidateTexSubImage() function call not supported by current GL context.")
	} else {
		InvalidateTexSubImage(texture, level, xoffset, yoffset, zoffset, width, height, depth)
		err = Util.Error("InvalidateTexSubImage()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetFragDataLocation() function.
func (_ GlSupport) GetFragDataLocation() bool {
	return C.cancallGetFragDataLocation() == 1
}

//	If Supports.GetFragDataLocation() is true, calls GetFragDataLocation() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetFragDataLocation(program Uint, name *Char) (err error, tryRetVal__ Int) {
	if !Supports.GetFragDataLocation() {
		err = fmt.Errorf("GetFragDataLocation() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetFragDataLocation(program, name)
		err = Util.Error("GetFragDataLocation()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix2x3fv() function.
func (_ GlSupport) UniformMatrix2x3fv() bool {
	return C.cancallUniformMatrix2x3fv() == 1
}

//	If Supports.UniformMatrix2x3fv() is true, calls UniformMatrix2x3fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix2x3fv(location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.UniformMatrix2x3fv() {
		err = fmt.Errorf("UniformMatrix2x3fv() function call not supported by current GL context.")
	} else {
		UniformMatrix2x3fv(location, count, transpose, value)
		err = Util.Error("UniformMatrix2x3fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform4d() function.
func (_ GlSupport) Uniform4d() bool {
	return C.cancallUniform4d() == 1
}

//	If Supports.Uniform4d() is true, calls Uniform4d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform4d(location Int, x Double, y Double, z Double, w Double) (err error) {
	if !Supports.Uniform4d() {
		err = fmt.Errorf("Uniform4d() function call not supported by current GL context.")
	} else {
		Uniform4d(location, x, y, z, w)
		err = Util.Error("Uniform4d()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClearBufferfv() function.
func (_ GlSupport) ClearBufferfv() bool {
	return C.cancallClearBufferfv() == 1
}

//	If Supports.ClearBufferfv() is true, calls ClearBufferfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClearBufferfv(buffer Enum, drawbuffer Int, value *Float) (err error) {
	if !Supports.ClearBufferfv() {
		err = fmt.Errorf("ClearBufferfv() function call not supported by current GL context.")
	} else {
		ClearBufferfv(buffer, drawbuffer, value)
		err = Util.Error("ClearBufferfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetQueryIndexediv() function.
func (_ GlSupport) GetQueryIndexediv() bool {
	return C.cancallGetQueryIndexediv() == 1
}

//	If Supports.GetQueryIndexediv() is true, calls GetQueryIndexediv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetQueryIndexediv(target Enum, index Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetQueryIndexediv() {
		err = fmt.Errorf("GetQueryIndexediv() function call not supported by current GL context.")
	} else {
		GetQueryIndexediv(target, index, pname, params)
		err = Util.Error("GetQueryIndexediv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform3fv() function.
func (_ GlSupport) ProgramUniform3fv() bool {
	return C.cancallProgramUniform3fv() == 1
}

//	If Supports.ProgramUniform3fv() is true, calls ProgramUniform3fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform3fv(program Uint, location Int, count Sizei, value *Float) (err error) {
	if !Supports.ProgramUniform3fv() {
		err = fmt.Errorf("ProgramUniform3fv() function call not supported by current GL context.")
	} else {
		ProgramUniform3fv(program, location, count, value)
		err = Util.Error("ProgramUniform3fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform1iv() function.
func (_ GlSupport) ProgramUniform1iv() bool {
	return C.cancallProgramUniform1iv() == 1
}

//	If Supports.ProgramUniform1iv() is true, calls ProgramUniform1iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform1iv(program Uint, location Int, count Sizei, value *Int) (err error) {
	if !Supports.ProgramUniform1iv() {
		err = fmt.Errorf("ProgramUniform1iv() function call not supported by current GL context.")
	} else {
		ProgramUniform1iv(program, location, count, value)
		err = Util.Error("ProgramUniform1iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform2fv() function.
func (_ GlSupport) Uniform2fv() bool {
	return C.cancallUniform2fv() == 1
}

//	If Supports.Uniform2fv() is true, calls Uniform2fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform2fv(location Int, count Sizei, value *Float) (err error) {
	if !Supports.Uniform2fv() {
		err = fmt.Errorf("Uniform2fv() function call not supported by current GL context.")
	} else {
		Uniform2fv(location, count, value)
		err = Util.Error("Uniform2fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramInterfaceiv() function.
func (_ GlSupport) GetProgramInterfaceiv() bool {
	return C.cancallGetProgramInterfaceiv() == 1
}

//	If Supports.GetProgramInterfaceiv() is true, calls GetProgramInterfaceiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramInterfaceiv(program Uint, programInterface Enum, pname Enum, params *Int) (err error) {
	if !Supports.GetProgramInterfaceiv() {
		err = fmt.Errorf("GetProgramInterfaceiv() function call not supported by current GL context.")
	} else {
		GetProgramInterfaceiv(program, programInterface, pname, params)
		err = Util.Error("GetProgramInterfaceiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix2dv() function.
func (_ GlSupport) UniformMatrix2dv() bool {
	return C.cancallUniformMatrix2dv() == 1
}

//	If Supports.UniformMatrix2dv() is true, calls UniformMatrix2dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix2dv(location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.UniformMatrix2dv() {
		err = fmt.Errorf("UniformMatrix2dv() function call not supported by current GL context.")
	} else {
		UniformMatrix2dv(location, count, transpose, value)
		err = Util.Error("UniformMatrix2dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsSync() function.
func (_ GlSupport) IsSync() bool {
	return C.cancallIsSync() == 1
}

//	If Supports.IsSync() is true, calls IsSync() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsSync(sync Sync) (err error, tryRetVal__ Boolean) {
	if !Supports.IsSync() {
		err = fmt.Errorf("IsSync() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsSync(sync)
		err = Util.Error("IsSync()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetBooleanv() function.
func (_ GlSupport) GetBooleanv() bool {
	return C.cancallGetBooleanv() == 1
}

//	If Supports.GetBooleanv() is true, calls GetBooleanv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetBooleanv(pname Enum, params *Boolean) (err error) {
	if !Supports.GetBooleanv() {
		err = fmt.Errorf("GetBooleanv() function call not supported by current GL context.")
	} else {
		GetBooleanv(pname, params)
		err = Util.Error("GetBooleanv()")
	}
	return
}

//	Returns whether the current GL context supports calling the StencilOp() function.
func (_ GlSupport) StencilOp() bool {
	return C.cancallStencilOp() == 1
}

//	If Supports.StencilOp() is true, calls StencilOp() and yields the err returned by Util.Error(), if any.
func (_ GlTry) StencilOp(fail Enum, zfail Enum, zpass Enum) (err error) {
	if !Supports.StencilOp() {
		err = fmt.Errorf("StencilOp() function call not supported by current GL context.")
	} else {
		StencilOp(fail, zfail, zpass)
		err = Util.Error("StencilOp()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetQueryObjectui64v() function.
func (_ GlSupport) GetQueryObjectui64v() bool {
	return C.cancallGetQueryObjectui64v() == 1
}

//	If Supports.GetQueryObjectui64v() is true, calls GetQueryObjectui64v() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetQueryObjectui64v(id Uint, pname Enum, params *Uint64) (err error) {
	if !Supports.GetQueryObjectui64v() {
		err = fmt.Errorf("GetQueryObjectui64v() function call not supported by current GL context.")
	} else {
		GetQueryObjectui64v(id, pname, params)
		err = Util.Error("GetQueryObjectui64v()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix4x3dv() function.
func (_ GlSupport) ProgramUniformMatrix4x3dv() bool {
	return C.cancallProgramUniformMatrix4x3dv() == 1
}

//	If Supports.ProgramUniformMatrix4x3dv() is true, calls ProgramUniformMatrix4x3dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix4x3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.ProgramUniformMatrix4x3dv() {
		err = fmt.Errorf("ProgramUniformMatrix4x3dv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix4x3dv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix4x3dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetObjectPtrLabel() function.
func (_ GlSupport) GetObjectPtrLabel() bool {
	return C.cancallGetObjectPtrLabel() == 1
}

//	If Supports.GetObjectPtrLabel() is true, calls GetObjectPtrLabel() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetObjectPtrLabel(ptr Ptr, bufSize Sizei, length *Sizei, label *Char) (err error) {
	if !Supports.GetObjectPtrLabel() {
		err = fmt.Errorf("GetObjectPtrLabel() function call not supported by current GL context.")
	} else {
		GetObjectPtrLabel(ptr, bufSize, length, label)
		err = Util.Error("GetObjectPtrLabel()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClearBufferfi() function.
func (_ GlSupport) ClearBufferfi() bool {
	return C.cancallClearBufferfi() == 1
}

//	If Supports.ClearBufferfi() is true, calls ClearBufferfi() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClearBufferfi(buffer Enum, drawbuffer Int, depth Float, stencil Int) (err error) {
	if !Supports.ClearBufferfi() {
		err = fmt.Errorf("ClearBufferfi() function call not supported by current GL context.")
	} else {
		ClearBufferfi(buffer, drawbuffer, depth, stencil)
		err = Util.Error("ClearBufferfi()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix2x4fv() function.
func (_ GlSupport) ProgramUniformMatrix2x4fv() bool {
	return C.cancallProgramUniformMatrix2x4fv() == 1
}

//	If Supports.ProgramUniformMatrix2x4fv() is true, calls ProgramUniformMatrix2x4fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix2x4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.ProgramUniformMatrix2x4fv() {
		err = fmt.Errorf("ProgramUniformMatrix2x4fv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix2x4fv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix2x4fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform3fv() function.
func (_ GlSupport) Uniform3fv() bool {
	return C.cancallUniform3fv() == 1
}

//	If Supports.Uniform3fv() is true, calls Uniform3fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform3fv(location Int, count Sizei, value *Float) (err error) {
	if !Supports.Uniform3fv() {
		err = fmt.Errorf("Uniform3fv() function call not supported by current GL context.")
	} else {
		Uniform3fv(location, count, value)
		err = Util.Error("Uniform3fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexP4ui() function.
func (_ GlSupport) VertexP4ui() bool {
	return C.cancallVertexP4ui() == 1
}

//	If Supports.VertexP4ui() is true, calls VertexP4ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexP4ui(type_ Enum, value Uint) (err error) {
	if !Supports.VertexP4ui() {
		err = fmt.Errorf("VertexP4ui() function call not supported by current GL context.")
	} else {
		VertexP4ui(type_, value)
		err = Util.Error("VertexP4ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetQueryObjecti64v() function.
func (_ GlSupport) GetQueryObjecti64v() bool {
	return C.cancallGetQueryObjecti64v() == 1
}

//	If Supports.GetQueryObjecti64v() is true, calls GetQueryObjecti64v() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetQueryObjecti64v(id Uint, pname Enum, params *Int64) (err error) {
	if !Supports.GetQueryObjecti64v() {
		err = fmt.Errorf("GetQueryObjecti64v() function call not supported by current GL context.")
	} else {
		GetQueryObjecti64v(id, pname, params)
		err = Util.Error("GetQueryObjecti64v()")
	}
	return
}

//	Returns whether the current GL context supports calling the CompressedTexImage1D() function.
func (_ GlSupport) CompressedTexImage1D() bool {
	return C.cancallCompressedTexImage1D() == 1
}

//	If Supports.CompressedTexImage1D() is true, calls CompressedTexImage1D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CompressedTexImage1D(target Enum, level Int, internalformat Enum, width Sizei, border Int, imageSize Sizei, data Ptr) (err error) {
	if !Supports.CompressedTexImage1D() {
		err = fmt.Errorf("CompressedTexImage1D() function call not supported by current GL context.")
	} else {
		CompressedTexImage1D(target, level, internalformat, width, border, imageSize, data)
		err = Util.Error("CompressedTexImage1D()")
	}
	return
}

//	Returns whether the current GL context supports calling the PrimitiveRestartIndex() function.
func (_ GlSupport) PrimitiveRestartIndex() bool {
	return C.cancallPrimitiveRestartIndex() == 1
}

//	If Supports.PrimitiveRestartIndex() is true, calls PrimitiveRestartIndex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PrimitiveRestartIndex(index Uint) (err error) {
	if !Supports.PrimitiveRestartIndex() {
		err = fmt.Errorf("PrimitiveRestartIndex() function call not supported by current GL context.")
	} else {
		PrimitiveRestartIndex(index)
		err = Util.Error("PrimitiveRestartIndex()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawTransformFeedback() function.
func (_ GlSupport) DrawTransformFeedback() bool {
	return C.cancallDrawTransformFeedback() == 1
}

//	If Supports.DrawTransformFeedback() is true, calls DrawTransformFeedback() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawTransformFeedback(mode Enum, id Uint) (err error) {
	if !Supports.DrawTransformFeedback() {
		err = fmt.Errorf("DrawTransformFeedback() function call not supported by current GL context.")
	} else {
		DrawTransformFeedback(mode, id)
		err = Util.Error("DrawTransformFeedback()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetDoublev() function.
func (_ GlSupport) GetDoublev() bool {
	return C.cancallGetDoublev() == 1
}

//	If Supports.GetDoublev() is true, calls GetDoublev() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetDoublev(pname Enum, params *Double) (err error) {
	if !Supports.GetDoublev() {
		err = fmt.Errorf("GetDoublev() function call not supported by current GL context.")
	} else {
		GetDoublev(pname, params)
		err = Util.Error("GetDoublev()")
	}
	return
}

//	Returns whether the current GL context supports calling the BlendFuncSeparate() function.
func (_ GlSupport) BlendFuncSeparate() bool {
	return C.cancallBlendFuncSeparate() == 1
}

//	If Supports.BlendFuncSeparate() is true, calls BlendFuncSeparate() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BlendFuncSeparate(sfactorRGB Enum, dfactorRGB Enum, sfactorAlpha Enum, dfactorAlpha Enum) (err error) {
	if !Supports.BlendFuncSeparate() {
		err = fmt.Errorf("BlendFuncSeparate() function call not supported by current GL context.")
	} else {
		BlendFuncSeparate(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha)
		err = Util.Error("BlendFuncSeparate()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribP1ui() function.
func (_ GlSupport) VertexAttribP1ui() bool {
	return C.cancallVertexAttribP1ui() == 1
}

//	If Supports.VertexAttribP1ui() is true, calls VertexAttribP1ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribP1ui(index Uint, type_ Enum, normalized Boolean, value Uint) (err error) {
	if !Supports.VertexAttribP1ui() {
		err = fmt.Errorf("VertexAttribP1ui() function call not supported by current GL context.")
	} else {
		VertexAttribP1ui(index, type_, normalized, value)
		err = Util.Error("VertexAttribP1ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribP4uiv() function.
func (_ GlSupport) VertexAttribP4uiv() bool {
	return C.cancallVertexAttribP4uiv() == 1
}

//	If Supports.VertexAttribP4uiv() is true, calls VertexAttribP4uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribP4uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) (err error) {
	if !Supports.VertexAttribP4uiv() {
		err = fmt.Errorf("VertexAttribP4uiv() function call not supported by current GL context.")
	} else {
		VertexAttribP4uiv(index, type_, normalized, value)
		err = Util.Error("VertexAttribP4uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the BeginTransformFeedback() function.
func (_ GlSupport) BeginTransformFeedback() bool {
	return C.cancallBeginTransformFeedback() == 1
}

//	If Supports.BeginTransformFeedback() is true, calls BeginTransformFeedback() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BeginTransformFeedback(primitiveMode Enum) (err error) {
	if !Supports.BeginTransformFeedback() {
		err = fmt.Errorf("BeginTransformFeedback() function call not supported by current GL context.")
	} else {
		BeginTransformFeedback(primitiveMode)
		err = Util.Error("BeginTransformFeedback()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClearStencil() function.
func (_ GlSupport) ClearStencil() bool {
	return C.cancallClearStencil() == 1
}

//	If Supports.ClearStencil() is true, calls ClearStencil() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClearStencil(s Int) (err error) {
	if !Supports.ClearStencil() {
		err = fmt.Errorf("ClearStencil() function call not supported by current GL context.")
	} else {
		ClearStencil(s)
		err = Util.Error("ClearStencil()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteFramebuffers() function.
func (_ GlSupport) DeleteFramebuffers() bool {
	return C.cancallDeleteFramebuffers() == 1
}

//	If Supports.DeleteFramebuffers() is true, calls DeleteFramebuffers() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteFramebuffers(n Sizei, framebuffers *Uint) (err error) {
	if !Supports.DeleteFramebuffers() {
		err = fmt.Errorf("DeleteFramebuffers() function call not supported by current GL context.")
	} else {
		DeleteFramebuffers(n, framebuffers)
		err = Util.Error("DeleteFramebuffers()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribL1dv() function.
func (_ GlSupport) VertexAttribL1dv() bool {
	return C.cancallVertexAttribL1dv() == 1
}

//	If Supports.VertexAttribL1dv() is true, calls VertexAttribL1dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribL1dv(index Uint, v *Double) (err error) {
	if !Supports.VertexAttribL1dv() {
		err = fmt.Errorf("VertexAttribL1dv() function call not supported by current GL context.")
	} else {
		VertexAttribL1dv(index, v)
		err = Util.Error("VertexAttribL1dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the PointParameterf() function.
func (_ GlSupport) PointParameterf() bool {
	return C.cancallPointParameterf() == 1
}

//	If Supports.PointParameterf() is true, calls PointParameterf() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PointParameterf(pname Enum, param Float) (err error) {
	if !Supports.PointParameterf() {
		err = fmt.Errorf("PointParameterf() function call not supported by current GL context.")
	} else {
		PointParameterf(pname, param)
		err = Util.Error("PointParameterf()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI4sv() function.
func (_ GlSupport) VertexAttribI4sv() bool {
	return C.cancallVertexAttribI4sv() == 1
}

//	If Supports.VertexAttribI4sv() is true, calls VertexAttribI4sv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI4sv(index Uint, v *Short) (err error) {
	if !Supports.VertexAttribI4sv() {
		err = fmt.Errorf("VertexAttribI4sv() function call not supported by current GL context.")
	} else {
		VertexAttribI4sv(index, v)
		err = Util.Error("VertexAttribI4sv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribL2dv() function.
func (_ GlSupport) VertexAttribL2dv() bool {
	return C.cancallVertexAttribL2dv() == 1
}

//	If Supports.VertexAttribL2dv() is true, calls VertexAttribL2dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribL2dv(index Uint, v *Double) (err error) {
	if !Supports.VertexAttribL2dv() {
		err = fmt.Errorf("VertexAttribL2dv() function call not supported by current GL context.")
	} else {
		VertexAttribL2dv(index, v)
		err = Util.Error("VertexAttribL2dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the BlendFunci() function.
func (_ GlSupport) BlendFunci() bool {
	return C.cancallBlendFunci() == 1
}

//	If Supports.BlendFunci() is true, calls BlendFunci() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BlendFunci(buf Uint, src Enum, dst Enum) (err error) {
	if !Supports.BlendFunci() {
		err = fmt.Errorf("BlendFunci() function call not supported by current GL context.")
	} else {
		BlendFunci(buf, src, dst)
		err = Util.Error("BlendFunci()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform2ui() function.
func (_ GlSupport) Uniform2ui() bool {
	return C.cancallUniform2ui() == 1
}

//	If Supports.Uniform2ui() is true, calls Uniform2ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform2ui(location Int, v0 Uint, v1 Uint) (err error) {
	if !Supports.Uniform2ui() {
		err = fmt.Errorf("Uniform2ui() function call not supported by current GL context.")
	} else {
		Uniform2ui(location, v0, v1)
		err = Util.Error("Uniform2ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramPipelineiv() function.
func (_ GlSupport) GetProgramPipelineiv() bool {
	return C.cancallGetProgramPipelineiv() == 1
}

//	If Supports.GetProgramPipelineiv() is true, calls GetProgramPipelineiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramPipelineiv(pipeline Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetProgramPipelineiv() {
		err = fmt.Errorf("GetProgramPipelineiv() function call not supported by current GL context.")
	} else {
		GetProgramPipelineiv(pipeline, pname, params)
		err = Util.Error("GetProgramPipelineiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the SampleMaski() function.
func (_ GlSupport) SampleMaski() bool {
	return C.cancallSampleMaski() == 1
}

//	If Supports.SampleMaski() is true, calls SampleMaski() and yields the err returned by Util.Error(), if any.
func (_ GlTry) SampleMaski(index Uint, mask Bitfield) (err error) {
	if !Supports.SampleMaski() {
		err = fmt.Errorf("SampleMaski() function call not supported by current GL context.")
	} else {
		SampleMaski(index, mask)
		err = Util.Error("SampleMaski()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetSamplerParameterIuiv() function.
func (_ GlSupport) GetSamplerParameterIuiv() bool {
	return C.cancallGetSamplerParameterIuiv() == 1
}

//	If Supports.GetSamplerParameterIuiv() is true, calls GetSamplerParameterIuiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetSamplerParameterIuiv(sampler Uint, pname Enum, params *Uint) (err error) {
	if !Supports.GetSamplerParameterIuiv() {
		err = fmt.Errorf("GetSamplerParameterIuiv() function call not supported by current GL context.")
	} else {
		GetSamplerParameterIuiv(sampler, pname, params)
		err = Util.Error("GetSamplerParameterIuiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the CopyBufferSubData() function.
func (_ GlSupport) CopyBufferSubData() bool {
	return C.cancallCopyBufferSubData() == 1
}

//	If Supports.CopyBufferSubData() is true, calls CopyBufferSubData() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CopyBufferSubData(readTarget Enum, writeTarget Enum, readOffset Intptr, writeOffset Intptr, size Sizeiptr) (err error) {
	if !Supports.CopyBufferSubData() {
		err = fmt.Errorf("CopyBufferSubData() function call not supported by current GL context.")
	} else {
		CopyBufferSubData(readTarget, writeTarget, readOffset, writeOffset, size)
		err = Util.Error("CopyBufferSubData()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetActiveUniformsiv() function.
func (_ GlSupport) GetActiveUniformsiv() bool {
	return C.cancallGetActiveUniformsiv() == 1
}

//	If Supports.GetActiveUniformsiv() is true, calls GetActiveUniformsiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetActiveUniformsiv(program Uint, uniformCount Sizei, uniformIndices *Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetActiveUniformsiv() {
		err = fmt.Errorf("GetActiveUniformsiv() function call not supported by current GL context.")
	} else {
		GetActiveUniformsiv(program, uniformCount, uniformIndices, pname, params)
		err = Util.Error("GetActiveUniformsiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetBufferParameteriv() function.
func (_ GlSupport) GetBufferParameteriv() bool {
	return C.cancallGetBufferParameteriv() == 1
}

//	If Supports.GetBufferParameteriv() is true, calls GetBufferParameteriv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetBufferParameteriv(target Enum, pname Enum, params *Int) (err error) {
	if !Supports.GetBufferParameteriv() {
		err = fmt.Errorf("GetBufferParameteriv() function call not supported by current GL context.")
	} else {
		GetBufferParameteriv(target, pname, params)
		err = Util.Error("GetBufferParameteriv()")
	}
	return
}

//	Returns whether the current GL context supports calling the StencilOpSeparate() function.
func (_ GlSupport) StencilOpSeparate() bool {
	return C.cancallStencilOpSeparate() == 1
}

//	If Supports.StencilOpSeparate() is true, calls StencilOpSeparate() and yields the err returned by Util.Error(), if any.
func (_ GlTry) StencilOpSeparate(face Enum, sfail Enum, dpfail Enum, dppass Enum) (err error) {
	if !Supports.StencilOpSeparate() {
		err = fmt.Errorf("StencilOpSeparate() function call not supported by current GL context.")
	} else {
		StencilOpSeparate(face, sfail, dpfail, dppass)
		err = Util.Error("StencilOpSeparate()")
	}
	return
}

//	Returns whether the current GL context supports calling the BlitFramebuffer() function.
func (_ GlSupport) BlitFramebuffer() bool {
	return C.cancallBlitFramebuffer() == 1
}

//	If Supports.BlitFramebuffer() is true, calls BlitFramebuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BlitFramebuffer(srcX0 Int, srcY0 Int, srcX1 Int, srcY1 Int, dstX0 Int, dstY0 Int, dstX1 Int, dstY1 Int, mask Bitfield, filter Enum) (err error) {
	if !Supports.BlitFramebuffer() {
		err = fmt.Errorf("BlitFramebuffer() function call not supported by current GL context.")
	} else {
		BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter)
		err = Util.Error("BlitFramebuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the DepthFunc() function.
func (_ GlSupport) DepthFunc() bool {
	return C.cancallDepthFunc() == 1
}

//	If Supports.DepthFunc() is true, calls DepthFunc() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DepthFunc(func_ Enum) (err error) {
	if !Supports.DepthFunc() {
		err = fmt.Errorf("DepthFunc() function call not supported by current GL context.")
	} else {
		DepthFunc(func_)
		err = Util.Error("DepthFunc()")
	}
	return
}

//	Returns whether the current GL context supports calling the DetachShader() function.
func (_ GlSupport) DetachShader() bool {
	return C.cancallDetachShader() == 1
}

//	If Supports.DetachShader() is true, calls DetachShader() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DetachShader(program Uint, shader Uint) (err error) {
	if !Supports.DetachShader() {
		err = fmt.Errorf("DetachShader() function call not supported by current GL context.")
	} else {
		DetachShader(program, shader)
		err = Util.Error("DetachShader()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClearBufferuiv() function.
func (_ GlSupport) ClearBufferuiv() bool {
	return C.cancallClearBufferuiv() == 1
}

//	If Supports.ClearBufferuiv() is true, calls ClearBufferuiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClearBufferuiv(buffer Enum, drawbuffer Int, value *Uint) (err error) {
	if !Supports.ClearBufferuiv() {
		err = fmt.Errorf("ClearBufferuiv() function call not supported by current GL context.")
	} else {
		ClearBufferuiv(buffer, drawbuffer, value)
		err = Util.Error("ClearBufferuiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramBinary() function.
func (_ GlSupport) ProgramBinary() bool {
	return C.cancallProgramBinary() == 1
}

//	If Supports.ProgramBinary() is true, calls ProgramBinary() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramBinary(program Uint, binaryFormat Enum, binary Ptr, length Sizei) (err error) {
	if !Supports.ProgramBinary() {
		err = fmt.Errorf("ProgramBinary() function call not supported by current GL context.")
	} else {
		ProgramBinary(program, binaryFormat, binary, length)
		err = Util.Error("ProgramBinary()")
	}
	return
}

//	Returns whether the current GL context supports calling the ViewportIndexedf() function.
func (_ GlSupport) ViewportIndexedf() bool {
	return C.cancallViewportIndexedf() == 1
}

//	If Supports.ViewportIndexedf() is true, calls ViewportIndexedf() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ViewportIndexedf(index Uint, x Float, y Float, w Float, h Float) (err error) {
	if !Supports.ViewportIndexedf() {
		err = fmt.Errorf("ViewportIndexedf() function call not supported by current GL context.")
	} else {
		ViewportIndexedf(index, x, y, w, h)
		err = Util.Error("ViewportIndexedf()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix3dv() function.
func (_ GlSupport) ProgramUniformMatrix3dv() bool {
	return C.cancallProgramUniformMatrix3dv() == 1
}

//	If Supports.ProgramUniformMatrix3dv() is true, calls ProgramUniformMatrix3dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.ProgramUniformMatrix3dv() {
		err = fmt.Errorf("ProgramUniformMatrix3dv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix3dv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix3dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix2x4dv() function.
func (_ GlSupport) UniformMatrix2x4dv() bool {
	return C.cancallUniformMatrix2x4dv() == 1
}

//	If Supports.UniformMatrix2x4dv() is true, calls UniformMatrix2x4dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix2x4dv(location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.UniformMatrix2x4dv() {
		err = fmt.Errorf("UniformMatrix2x4dv() function call not supported by current GL context.")
	} else {
		UniformMatrix2x4dv(location, count, transpose, value)
		err = Util.Error("UniformMatrix2x4dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramParameteri() function.
func (_ GlSupport) ProgramParameteri() bool {
	return C.cancallProgramParameteri() == 1
}

//	If Supports.ProgramParameteri() is true, calls ProgramParameteri() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramParameteri(program Uint, pname Enum, value Int) (err error) {
	if !Supports.ProgramParameteri() {
		err = fmt.Errorf("ProgramParameteri() function call not supported by current GL context.")
	} else {
		ProgramParameteri(program, pname, value)
		err = Util.Error("ProgramParameteri()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiTexCoordP3ui() function.
func (_ GlSupport) MultiTexCoordP3ui() bool {
	return C.cancallMultiTexCoordP3ui() == 1
}

//	If Supports.MultiTexCoordP3ui() is true, calls MultiTexCoordP3ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiTexCoordP3ui(texture Enum, type_ Enum, coords Uint) (err error) {
	if !Supports.MultiTexCoordP3ui() {
		err = fmt.Errorf("MultiTexCoordP3ui() function call not supported by current GL context.")
	} else {
		MultiTexCoordP3ui(texture, type_, coords)
		err = Util.Error("MultiTexCoordP3ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix4dv() function.
func (_ GlSupport) UniformMatrix4dv() bool {
	return C.cancallUniformMatrix4dv() == 1
}

//	If Supports.UniformMatrix4dv() is true, calls UniformMatrix4dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix4dv(location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.UniformMatrix4dv() {
		err = fmt.Errorf("UniformMatrix4dv() function call not supported by current GL context.")
	} else {
		UniformMatrix4dv(location, count, transpose, value)
		err = Util.Error("UniformMatrix4dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindFragDataLocation() function.
func (_ GlSupport) BindFragDataLocation() bool {
	return C.cancallBindFragDataLocation() == 1
}

//	If Supports.BindFragDataLocation() is true, calls BindFragDataLocation() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindFragDataLocation(program Uint, color Uint, name *Char) (err error) {
	if !Supports.BindFragDataLocation() {
		err = fmt.Errorf("BindFragDataLocation() function call not supported by current GL context.")
	} else {
		BindFragDataLocation(program, color, name)
		err = Util.Error("BindFragDataLocation()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetAttachedShaders() function.
func (_ GlSupport) GetAttachedShaders() bool {
	return C.cancallGetAttachedShaders() == 1
}

//	If Supports.GetAttachedShaders() is true, calls GetAttachedShaders() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetAttachedShaders(program Uint, maxCount Sizei, count *Sizei, obj *Uint) (err error) {
	if !Supports.GetAttachedShaders() {
		err = fmt.Errorf("GetAttachedShaders() function call not supported by current GL context.")
	} else {
		GetAttachedShaders(program, maxCount, count, obj)
		err = Util.Error("GetAttachedShaders()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexP4uiv() function.
func (_ GlSupport) VertexP4uiv() bool {
	return C.cancallVertexP4uiv() == 1
}

//	If Supports.VertexP4uiv() is true, calls VertexP4uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexP4uiv(type_ Enum, value *Uint) (err error) {
	if !Supports.VertexP4uiv() {
		err = fmt.Errorf("VertexP4uiv() function call not supported by current GL context.")
	} else {
		VertexP4uiv(type_, value)
		err = Util.Error("VertexP4uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ReadPixels() function.
func (_ GlSupport) ReadPixels() bool {
	return C.cancallReadPixels() == 1
}

//	If Supports.ReadPixels() is true, calls ReadPixels() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Ptr) (err error) {
	if !Supports.ReadPixels() {
		err = fmt.Errorf("ReadPixels() function call not supported by current GL context.")
	} else {
		ReadPixels(x, y, width, height, format, type_, pixels)
		err = Util.Error("ReadPixels()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribIFormat() function.
func (_ GlSupport) VertexAttribIFormat() bool {
	return C.cancallVertexAttribIFormat() == 1
}

//	If Supports.VertexAttribIFormat() is true, calls VertexAttribIFormat() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribIFormat(attribindex Uint, size Int, type_ Enum, relativeoffset Uint) (err error) {
	if !Supports.VertexAttribIFormat() {
		err = fmt.Errorf("VertexAttribIFormat() function call not supported by current GL context.")
	} else {
		VertexAttribIFormat(attribindex, size, type_, relativeoffset)
		err = Util.Error("VertexAttribIFormat()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexStorage1D() function.
func (_ GlSupport) TexStorage1D() bool {
	return C.cancallTexStorage1D() == 1
}

//	If Supports.TexStorage1D() is true, calls TexStorage1D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexStorage1D(target Enum, levels Sizei, internalformat Enum, width Sizei) (err error) {
	if !Supports.TexStorage1D() {
		err = fmt.Errorf("TexStorage1D() function call not supported by current GL context.")
	} else {
		TexStorage1D(target, levels, internalformat, width)
		err = Util.Error("TexStorage1D()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform4dv() function.
func (_ GlSupport) ProgramUniform4dv() bool {
	return C.cancallProgramUniform4dv() == 1
}

//	If Supports.ProgramUniform4dv() is true, calls ProgramUniform4dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform4dv(program Uint, location Int, count Sizei, value *Double) (err error) {
	if !Supports.ProgramUniform4dv() {
		err = fmt.Errorf("ProgramUniform4dv() function call not supported by current GL context.")
	} else {
		ProgramUniform4dv(program, location, count, value)
		err = Util.Error("ProgramUniform4dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexBindingDivisor() function.
func (_ GlSupport) VertexBindingDivisor() bool {
	return C.cancallVertexBindingDivisor() == 1
}

//	If Supports.VertexBindingDivisor() is true, calls VertexBindingDivisor() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexBindingDivisor(bindingindex Uint, divisor Uint) (err error) {
	if !Supports.VertexBindingDivisor() {
		err = fmt.Errorf("VertexBindingDivisor() function call not supported by current GL context.")
	} else {
		VertexBindingDivisor(bindingindex, divisor)
		err = Util.Error("VertexBindingDivisor()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawRangeElementsBaseVertex() function.
func (_ GlSupport) DrawRangeElementsBaseVertex() bool {
	return C.cancallDrawRangeElementsBaseVertex() == 1
}

//	If Supports.DrawRangeElementsBaseVertex() is true, calls DrawRangeElementsBaseVertex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawRangeElementsBaseVertex(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Ptr, basevertex Int) (err error) {
	if !Supports.DrawRangeElementsBaseVertex() {
		err = fmt.Errorf("DrawRangeElementsBaseVertex() function call not supported by current GL context.")
	} else {
		DrawRangeElementsBaseVertex(mode, start, end, count, type_, indices, basevertex)
		err = Util.Error("DrawRangeElementsBaseVertex()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetUniformiv() function.
func (_ GlSupport) GetUniformiv() bool {
	return C.cancallGetUniformiv() == 1
}

//	If Supports.GetUniformiv() is true, calls GetUniformiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetUniformiv(program Uint, location Int, params *Int) (err error) {
	if !Supports.GetUniformiv() {
		err = fmt.Errorf("GetUniformiv() function call not supported by current GL context.")
	} else {
		GetUniformiv(program, location, params)
		err = Util.Error("GetUniformiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribL3dv() function.
func (_ GlSupport) VertexAttribL3dv() bool {
	return C.cancallVertexAttribL3dv() == 1
}

//	If Supports.VertexAttribL3dv() is true, calls VertexAttribL3dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribL3dv(index Uint, v *Double) (err error) {
	if !Supports.VertexAttribL3dv() {
		err = fmt.Errorf("VertexAttribL3dv() function call not supported by current GL context.")
	} else {
		VertexAttribL3dv(index, v)
		err = Util.Error("VertexAttribL3dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the InvalidateBufferData() function.
func (_ GlSupport) InvalidateBufferData() bool {
	return C.cancallInvalidateBufferData() == 1
}

//	If Supports.InvalidateBufferData() is true, calls InvalidateBufferData() and yields the err returned by Util.Error(), if any.
func (_ GlTry) InvalidateBufferData(buffer Uint) (err error) {
	if !Supports.InvalidateBufferData() {
		err = fmt.Errorf("InvalidateBufferData() function call not supported by current GL context.")
	} else {
		InvalidateBufferData(buffer)
		err = Util.Error("InvalidateBufferData()")
	}
	return
}

//	Returns whether the current GL context supports calling the ResumeTransformFeedback() function.
func (_ GlSupport) ResumeTransformFeedback() bool {
	return C.cancallResumeTransformFeedback() == 1
}

//	If Supports.ResumeTransformFeedback() is true, calls ResumeTransformFeedback() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ResumeTransformFeedback() (err error) {
	if !Supports.ResumeTransformFeedback() {
		err = fmt.Errorf("ResumeTransformFeedback() function call not supported by current GL context.")
	} else {
		ResumeTransformFeedback()
		err = Util.Error("ResumeTransformFeedback()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix2x3dv() function.
func (_ GlSupport) UniformMatrix2x3dv() bool {
	return C.cancallUniformMatrix2x3dv() == 1
}

//	If Supports.UniformMatrix2x3dv() is true, calls UniformMatrix2x3dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix2x3dv(location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.UniformMatrix2x3dv() {
		err = fmt.Errorf("UniformMatrix2x3dv() function call not supported by current GL context.")
	} else {
		UniformMatrix2x3dv(location, count, transpose, value)
		err = Util.Error("UniformMatrix2x3dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the EndQueryIndexed() function.
func (_ GlSupport) EndQueryIndexed() bool {
	return C.cancallEndQueryIndexed() == 1
}

//	If Supports.EndQueryIndexed() is true, calls EndQueryIndexed() and yields the err returned by Util.Error(), if any.
func (_ GlTry) EndQueryIndexed(target Enum, index Uint) (err error) {
	if !Supports.EndQueryIndexed() {
		err = fmt.Errorf("EndQueryIndexed() function call not supported by current GL context.")
	} else {
		EndQueryIndexed(target, index)
		err = Util.Error("EndQueryIndexed()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetVertexAttribIuiv() function.
func (_ GlSupport) GetVertexAttribIuiv() bool {
	return C.cancallGetVertexAttribIuiv() == 1
}

//	If Supports.GetVertexAttribIuiv() is true, calls GetVertexAttribIuiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetVertexAttribIuiv(index Uint, pname Enum, params *Uint) (err error) {
	if !Supports.GetVertexAttribIuiv() {
		err = fmt.Errorf("GetVertexAttribIuiv() function call not supported by current GL context.")
	} else {
		GetVertexAttribIuiv(index, pname, params)
		err = Util.Error("GetVertexAttribIuiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the BeginQueryIndexed() function.
func (_ GlSupport) BeginQueryIndexed() bool {
	return C.cancallBeginQueryIndexed() == 1
}

//	If Supports.BeginQueryIndexed() is true, calls BeginQueryIndexed() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BeginQueryIndexed(target Enum, index Uint, id Uint) (err error) {
	if !Supports.BeginQueryIndexed() {
		err = fmt.Errorf("BeginQueryIndexed() function call not supported by current GL context.")
	} else {
		BeginQueryIndexed(target, index, id)
		err = Util.Error("BeginQueryIndexed()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix4x3dv() function.
func (_ GlSupport) UniformMatrix4x3dv() bool {
	return C.cancallUniformMatrix4x3dv() == 1
}

//	If Supports.UniformMatrix4x3dv() is true, calls UniformMatrix4x3dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix4x3dv(location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.UniformMatrix4x3dv() {
		err = fmt.Errorf("UniformMatrix4x3dv() function call not supported by current GL context.")
	} else {
		UniformMatrix4x3dv(location, count, transpose, value)
		err = Util.Error("UniformMatrix4x3dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexBufferRange() function.
func (_ GlSupport) TexBufferRange() bool {
	return C.cancallTexBufferRange() == 1
}

//	If Supports.TexBufferRange() is true, calls TexBufferRange() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexBufferRange(target Enum, internalformat Enum, buffer Uint, offset Intptr, size Sizeiptr) (err error) {
	if !Supports.TexBufferRange() {
		err = fmt.Errorf("TexBufferRange() function call not supported by current GL context.")
	} else {
		TexBufferRange(target, internalformat, buffer, offset, size)
		err = Util.Error("TexBufferRange()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetVertexAttribLdv() function.
func (_ GlSupport) GetVertexAttribLdv() bool {
	return C.cancallGetVertexAttribLdv() == 1
}

//	If Supports.GetVertexAttribLdv() is true, calls GetVertexAttribLdv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetVertexAttribLdv(index Uint, pname Enum, params *Double) (err error) {
	if !Supports.GetVertexAttribLdv() {
		err = fmt.Errorf("GetVertexAttribLdv() function call not supported by current GL context.")
	} else {
		GetVertexAttribLdv(index, pname, params)
		err = Util.Error("GetVertexAttribLdv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawTransformFeedbackStream() function.
func (_ GlSupport) DrawTransformFeedbackStream() bool {
	return C.cancallDrawTransformFeedbackStream() == 1
}

//	If Supports.DrawTransformFeedbackStream() is true, calls DrawTransformFeedbackStream() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawTransformFeedbackStream(mode Enum, id Uint, stream Uint) (err error) {
	if !Supports.DrawTransformFeedbackStream() {
		err = fmt.Errorf("DrawTransformFeedbackStream() function call not supported by current GL context.")
	} else {
		DrawTransformFeedbackStream(mode, id, stream)
		err = Util.Error("DrawTransformFeedbackStream()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindFramebuffer() function.
func (_ GlSupport) BindFramebuffer() bool {
	return C.cancallBindFramebuffer() == 1
}

//	If Supports.BindFramebuffer() is true, calls BindFramebuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindFramebuffer(target Enum, framebuffer Uint) (err error) {
	if !Supports.BindFramebuffer() {
		err = fmt.Errorf("BindFramebuffer() function call not supported by current GL context.")
	} else {
		BindFramebuffer(target, framebuffer)
		err = Util.Error("BindFramebuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteProgram() function.
func (_ GlSupport) DeleteProgram() bool {
	return C.cancallDeleteProgram() == 1
}

//	If Supports.DeleteProgram() is true, calls DeleteProgram() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteProgram(program Uint) (err error) {
	if !Supports.DeleteProgram() {
		err = fmt.Errorf("DeleteProgram() function call not supported by current GL context.")
	} else {
		DeleteProgram(program)
		err = Util.Error("DeleteProgram()")
	}
	return
}

//	Returns whether the current GL context supports calling the ScissorIndexedv() function.
func (_ GlSupport) ScissorIndexedv() bool {
	return C.cancallScissorIndexedv() == 1
}

//	If Supports.ScissorIndexedv() is true, calls ScissorIndexedv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ScissorIndexedv(index Uint, v *Int) (err error) {
	if !Supports.ScissorIndexedv() {
		err = fmt.Errorf("ScissorIndexedv() function call not supported by current GL context.")
	} else {
		ScissorIndexedv(index, v)
		err = Util.Error("ScissorIndexedv()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix3x2dv() function.
func (_ GlSupport) UniformMatrix3x2dv() bool {
	return C.cancallUniformMatrix3x2dv() == 1
}

//	If Supports.UniformMatrix3x2dv() is true, calls UniformMatrix3x2dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix3x2dv(location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.UniformMatrix3x2dv() {
		err = fmt.Errorf("UniformMatrix3x2dv() function call not supported by current GL context.")
	} else {
		UniformMatrix3x2dv(location, count, transpose, value)
		err = Util.Error("UniformMatrix3x2dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsTexture() function.
func (_ GlSupport) IsTexture() bool {
	return C.cancallIsTexture() == 1
}

//	If Supports.IsTexture() is true, calls IsTexture() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsTexture(texture Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsTexture() {
		err = fmt.Errorf("IsTexture() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsTexture(texture)
		err = Util.Error("IsTexture()")
	}
	return
}

//	Returns whether the current GL context supports calling the SamplerParameterIiv() function.
func (_ GlSupport) SamplerParameterIiv() bool {
	return C.cancallSamplerParameterIiv() == 1
}

//	If Supports.SamplerParameterIiv() is true, calls SamplerParameterIiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) SamplerParameterIiv(sampler Uint, pname Enum, param *Int) (err error) {
	if !Supports.SamplerParameterIiv() {
		err = fmt.Errorf("SamplerParameterIiv() function call not supported by current GL context.")
	} else {
		SamplerParameterIiv(sampler, pname, param)
		err = Util.Error("SamplerParameterIiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform4fv() function.
func (_ GlSupport) ProgramUniform4fv() bool {
	return C.cancallProgramUniform4fv() == 1
}

//	If Supports.ProgramUniform4fv() is true, calls ProgramUniform4fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform4fv(program Uint, location Int, count Sizei, value *Float) (err error) {
	if !Supports.ProgramUniform4fv() {
		err = fmt.Errorf("ProgramUniform4fv() function call not supported by current GL context.")
	} else {
		ProgramUniform4fv(program, location, count, value)
		err = Util.Error("ProgramUniform4fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the BlendEquationi() function.
func (_ GlSupport) BlendEquationi() bool {
	return C.cancallBlendEquationi() == 1
}

//	If Supports.BlendEquationi() is true, calls BlendEquationi() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BlendEquationi(buf Uint, mode Enum) (err error) {
	if !Supports.BlendEquationi() {
		err = fmt.Errorf("BlendEquationi() function call not supported by current GL context.")
	} else {
		BlendEquationi(buf, mode)
		err = Util.Error("BlendEquationi()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteSync() function.
func (_ GlSupport) DeleteSync() bool {
	return C.cancallDeleteSync() == 1
}

//	If Supports.DeleteSync() is true, calls DeleteSync() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteSync(sync Sync) (err error) {
	if !Supports.DeleteSync() {
		err = fmt.Errorf("DeleteSync() function call not supported by current GL context.")
	} else {
		DeleteSync(sync)
		err = Util.Error("DeleteSync()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform3iv() function.
func (_ GlSupport) ProgramUniform3iv() bool {
	return C.cancallProgramUniform3iv() == 1
}

//	If Supports.ProgramUniform3iv() is true, calls ProgramUniform3iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform3iv(program Uint, location Int, count Sizei, value *Int) (err error) {
	if !Supports.ProgramUniform3iv() {
		err = fmt.Errorf("ProgramUniform3iv() function call not supported by current GL context.")
	} else {
		ProgramUniform3iv(program, location, count, value)
		err = Util.Error("ProgramUniform3iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the InvalidateTexImage() function.
func (_ GlSupport) InvalidateTexImage() bool {
	return C.cancallInvalidateTexImage() == 1
}

//	If Supports.InvalidateTexImage() is true, calls InvalidateTexImage() and yields the err returned by Util.Error(), if any.
func (_ GlTry) InvalidateTexImage(texture Uint, level Int) (err error) {
	if !Supports.InvalidateTexImage() {
		err = fmt.Errorf("InvalidateTexImage() function call not supported by current GL context.")
	} else {
		InvalidateTexImage(texture, level)
		err = Util.Error("InvalidateTexImage()")
	}
	return
}

//	Returns whether the current GL context supports calling the ViewportArrayv() function.
func (_ GlSupport) ViewportArrayv() bool {
	return C.cancallViewportArrayv() == 1
}

//	If Supports.ViewportArrayv() is true, calls ViewportArrayv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ViewportArrayv(first Uint, count Sizei, v *Float) (err error) {
	if !Supports.ViewportArrayv() {
		err = fmt.Errorf("ViewportArrayv() function call not supported by current GL context.")
	} else {
		ViewportArrayv(first, count, v)
		err = Util.Error("ViewportArrayv()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexSubImage1D() function.
func (_ GlSupport) TexSubImage1D() bool {
	return C.cancallTexSubImage1D() == 1
}

//	If Supports.TexSubImage1D() is true, calls TexSubImage1D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, type_ Enum, pixels Ptr) (err error) {
	if !Supports.TexSubImage1D() {
		err = fmt.Errorf("TexSubImage1D() function call not supported by current GL context.")
	} else {
		TexSubImage1D(target, level, xoffset, width, format, type_, pixels)
		err = Util.Error("TexSubImage1D()")
	}
	return
}

//	Returns whether the current GL context supports calling the CreateShaderProgramv() function.
func (_ GlSupport) CreateShaderProgramv() bool {
	return C.cancallCreateShaderProgramv() == 1
}

//	If Supports.CreateShaderProgramv() is true, calls CreateShaderProgramv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CreateShaderProgramv(type_ Enum, count Sizei, strings **Char) (err error, tryRetVal__ Uint) {
	if !Supports.CreateShaderProgramv() {
		err = fmt.Errorf("CreateShaderProgramv() function call not supported by current GL context.")
	} else {
		tryRetVal__ = CreateShaderProgramv(type_, count, strings)
		err = Util.Error("CreateShaderProgramv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GenRenderbuffers() function.
func (_ GlSupport) GenRenderbuffers() bool {
	return C.cancallGenRenderbuffers() == 1
}

//	If Supports.GenRenderbuffers() is true, calls GenRenderbuffers() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GenRenderbuffers(n Sizei, renderbuffers *Uint) (err error) {
	if !Supports.GenRenderbuffers() {
		err = fmt.Errorf("GenRenderbuffers() function call not supported by current GL context.")
	} else {
		GenRenderbuffers(n, renderbuffers)
		err = Util.Error("GenRenderbuffers()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix2dv() function.
func (_ GlSupport) ProgramUniformMatrix2dv() bool {
	return C.cancallProgramUniformMatrix2dv() == 1
}

//	If Supports.ProgramUniformMatrix2dv() is true, calls ProgramUniformMatrix2dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.ProgramUniformMatrix2dv() {
		err = fmt.Errorf("ProgramUniformMatrix2dv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix2dv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix2dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform2ui() function.
func (_ GlSupport) ProgramUniform2ui() bool {
	return C.cancallProgramUniform2ui() == 1
}

//	If Supports.ProgramUniform2ui() is true, calls ProgramUniform2ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform2ui(program Uint, location Int, v0 Uint, v1 Uint) (err error) {
	if !Supports.ProgramUniform2ui() {
		err = fmt.Errorf("ProgramUniform2ui() function call not supported by current GL context.")
	} else {
		ProgramUniform2ui(program, location, v0, v1)
		err = Util.Error("ProgramUniform2ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the CopyTexSubImage2D() function.
func (_ GlSupport) CopyTexSubImage2D() bool {
	return C.cancallCopyTexSubImage2D() == 1
}

//	If Supports.CopyTexSubImage2D() is true, calls CopyTexSubImage2D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CopyTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, x Int, y Int, width Sizei, height Sizei) (err error) {
	if !Supports.CopyTexSubImage2D() {
		err = fmt.Errorf("CopyTexSubImage2D() function call not supported by current GL context.")
	} else {
		CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height)
		err = Util.Error("CopyTexSubImage2D()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiDrawElementsBaseVertex() function.
func (_ GlSupport) MultiDrawElementsBaseVertex() bool {
	return C.cancallMultiDrawElementsBaseVertex() == 1
}

//	If Supports.MultiDrawElementsBaseVertex() is true, calls MultiDrawElementsBaseVertex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiDrawElementsBaseVertex(mode Enum, count *Sizei, type_ Enum, indices *Ptr, drawcount Sizei, basevertex *Int) (err error) {
	if !Supports.MultiDrawElementsBaseVertex() {
		err = fmt.Errorf("MultiDrawElementsBaseVertex() function call not supported by current GL context.")
	} else {
		MultiDrawElementsBaseVertex(mode, count, type_, indices, drawcount, basevertex)
		err = Util.Error("MultiDrawElementsBaseVertex()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetCompressedTexImage() function.
func (_ GlSupport) GetCompressedTexImage() bool {
	return C.cancallGetCompressedTexImage() == 1
}

//	If Supports.GetCompressedTexImage() is true, calls GetCompressedTexImage() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetCompressedTexImage(target Enum, level Int, img Ptr) (err error) {
	if !Supports.GetCompressedTexImage() {
		err = fmt.Errorf("GetCompressedTexImage() function call not supported by current GL context.")
	} else {
		GetCompressedTexImage(target, level, img)
		err = Util.Error("GetCompressedTexImage()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI3ui() function.
func (_ GlSupport) VertexAttribI3ui() bool {
	return C.cancallVertexAttribI3ui() == 1
}

//	If Supports.VertexAttribI3ui() is true, calls VertexAttribI3ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI3ui(index Uint, x Uint, y Uint, z Uint) (err error) {
	if !Supports.VertexAttribI3ui() {
		err = fmt.Errorf("VertexAttribI3ui() function call not supported by current GL context.")
	} else {
		VertexAttribI3ui(index, x, y, z)
		err = Util.Error("VertexAttribI3ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix3x4fv() function.
func (_ GlSupport) UniformMatrix3x4fv() bool {
	return C.cancallUniformMatrix3x4fv() == 1
}

//	If Supports.UniformMatrix3x4fv() is true, calls UniformMatrix3x4fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix3x4fv(location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.UniformMatrix3x4fv() {
		err = fmt.Errorf("UniformMatrix3x4fv() function call not supported by current GL context.")
	} else {
		UniformMatrix3x4fv(location, count, transpose, value)
		err = Util.Error("UniformMatrix3x4fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI1iv() function.
func (_ GlSupport) VertexAttribI1iv() bool {
	return C.cancallVertexAttribI1iv() == 1
}

//	If Supports.VertexAttribI1iv() is true, calls VertexAttribI1iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI1iv(index Uint, v *Int) (err error) {
	if !Supports.VertexAttribI1iv() {
		err = fmt.Errorf("VertexAttribI1iv() function call not supported by current GL context.")
	} else {
		VertexAttribI1iv(index, v)
		err = Util.Error("VertexAttribI1iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetBufferSubData() function.
func (_ GlSupport) GetBufferSubData() bool {
	return C.cancallGetBufferSubData() == 1
}

//	If Supports.GetBufferSubData() is true, calls GetBufferSubData() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetBufferSubData(target Enum, offset Intptr, size Sizeiptr, data Ptr) (err error) {
	if !Supports.GetBufferSubData() {
		err = fmt.Errorf("GetBufferSubData() function call not supported by current GL context.")
	} else {
		GetBufferSubData(target, offset, size, data)
		err = Util.Error("GetBufferSubData()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexImage1D() function.
func (_ GlSupport) TexImage1D() bool {
	return C.cancallTexImage1D() == 1
}

//	If Supports.TexImage1D() is true, calls TexImage1D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexImage1D(target Enum, level Int, internalformat Int, width Sizei, border Int, format Enum, type_ Enum, pixels Ptr) (err error) {
	if !Supports.TexImage1D() {
		err = fmt.Errorf("TexImage1D() function call not supported by current GL context.")
	} else {
		TexImage1D(target, level, internalformat, width, border, format, type_, pixels)
		err = Util.Error("TexImage1D()")
	}
	return
}

//	Returns whether the current GL context supports calling the FramebufferTexture() function.
func (_ GlSupport) FramebufferTexture() bool {
	return C.cancallFramebufferTexture() == 1
}

//	If Supports.FramebufferTexture() is true, calls FramebufferTexture() and yields the err returned by Util.Error(), if any.
func (_ GlTry) FramebufferTexture(target Enum, attachment Enum, texture Uint, level Int) (err error) {
	if !Supports.FramebufferTexture() {
		err = fmt.Errorf("FramebufferTexture() function call not supported by current GL context.")
	} else {
		FramebufferTexture(target, attachment, texture, level)
		err = Util.Error("FramebufferTexture()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI1i() function.
func (_ GlSupport) VertexAttribI1i() bool {
	return C.cancallVertexAttribI1i() == 1
}

//	If Supports.VertexAttribI1i() is true, calls VertexAttribI1i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI1i(index Uint, x Int) (err error) {
	if !Supports.VertexAttribI1i() {
		err = fmt.Errorf("VertexAttribI1i() function call not supported by current GL context.")
	} else {
		VertexAttribI1i(index, x)
		err = Util.Error("VertexAttribI1i()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetVertexAttribIiv() function.
func (_ GlSupport) GetVertexAttribIiv() bool {
	return C.cancallGetVertexAttribIiv() == 1
}

//	If Supports.GetVertexAttribIiv() is true, calls GetVertexAttribIiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetVertexAttribIiv(index Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetVertexAttribIiv() {
		err = fmt.Errorf("GetVertexAttribIiv() function call not supported by current GL context.")
	} else {
		GetVertexAttribIiv(index, pname, params)
		err = Util.Error("GetVertexAttribIiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetBufferParameteri64v() function.
func (_ GlSupport) GetBufferParameteri64v() bool {
	return C.cancallGetBufferParameteri64v() == 1
}

//	If Supports.GetBufferParameteri64v() is true, calls GetBufferParameteri64v() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetBufferParameteri64v(target Enum, pname Enum, params *Int64) (err error) {
	if !Supports.GetBufferParameteri64v() {
		err = fmt.Errorf("GetBufferParameteri64v() function call not supported by current GL context.")
	} else {
		GetBufferParameteri64v(target, pname, params)
		err = Util.Error("GetBufferParameteri64v()")
	}
	return
}

//	Returns whether the current GL context supports calling the BeginConditionalRender() function.
func (_ GlSupport) BeginConditionalRender() bool {
	return C.cancallBeginConditionalRender() == 1
}

//	If Supports.BeginConditionalRender() is true, calls BeginConditionalRender() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BeginConditionalRender(id Uint, mode Enum) (err error) {
	if !Supports.BeginConditionalRender() {
		err = fmt.Errorf("BeginConditionalRender() function call not supported by current GL context.")
	} else {
		BeginConditionalRender(id, mode)
		err = Util.Error("BeginConditionalRender()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetInternalformati64v() function.
func (_ GlSupport) GetInternalformati64v() bool {
	return C.cancallGetInternalformati64v() == 1
}

//	If Supports.GetInternalformati64v() is true, calls GetInternalformati64v() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetInternalformati64v(target Enum, internalformat Enum, pname Enum, bufSize Sizei, params *Int64) (err error) {
	if !Supports.GetInternalformati64v() {
		err = fmt.Errorf("GetInternalformati64v() function call not supported by current GL context.")
	} else {
		GetInternalformati64v(target, internalformat, pname, bufSize, params)
		err = Util.Error("GetInternalformati64v()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI2ui() function.
func (_ GlSupport) VertexAttribI2ui() bool {
	return C.cancallVertexAttribI2ui() == 1
}

//	If Supports.VertexAttribI2ui() is true, calls VertexAttribI2ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI2ui(index Uint, x Uint, y Uint) (err error) {
	if !Supports.VertexAttribI2ui() {
		err = fmt.Errorf("VertexAttribI2ui() function call not supported by current GL context.")
	} else {
		VertexAttribI2ui(index, x, y)
		err = Util.Error("VertexAttribI2ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the BlendEquation() function.
func (_ GlSupport) BlendEquation() bool {
	return C.cancallBlendEquation() == 1
}

//	If Supports.BlendEquation() is true, calls BlendEquation() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BlendEquation(mode Enum) (err error) {
	if !Supports.BlendEquation() {
		err = fmt.Errorf("BlendEquation() function call not supported by current GL context.")
	} else {
		BlendEquation(mode)
		err = Util.Error("BlendEquation()")
	}
	return
}

//	Returns whether the current GL context supports calling the Enablei() function.
func (_ GlSupport) Enablei() bool {
	return C.cancallEnablei() == 1
}

//	If Supports.Enablei() is true, calls Enablei() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Enablei(target Enum, index Uint) (err error) {
	if !Supports.Enablei() {
		err = fmt.Errorf("Enablei() function call not supported by current GL context.")
	} else {
		Enablei(target, index)
		err = Util.Error("Enablei()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetStringi() function.
func (_ GlSupport) GetStringi() bool {
	return C.cancallGetStringi() == 1
}

//	If Supports.GetStringi() is true, calls GetStringi() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetStringi(name Enum, index Uint) (err error, tryRetVal__ *Ubyte) {
	if !Supports.GetStringi() {
		err = fmt.Errorf("GetStringi() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetStringi(name, index)
		err = Util.Error("GetStringi()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix3x2dv() function.
func (_ GlSupport) ProgramUniformMatrix3x2dv() bool {
	return C.cancallProgramUniformMatrix3x2dv() == 1
}

//	If Supports.ProgramUniformMatrix3x2dv() is true, calls ProgramUniformMatrix3x2dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix3x2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.ProgramUniformMatrix3x2dv() {
		err = fmt.Errorf("ProgramUniformMatrix3x2dv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix3x2dv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix3x2dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawTransformFeedbackStreamInstanced() function.
func (_ GlSupport) DrawTransformFeedbackStreamInstanced() bool {
	return C.cancallDrawTransformFeedbackStreamInstanced() == 1
}

//	If Supports.DrawTransformFeedbackStreamInstanced() is true, calls DrawTransformFeedbackStreamInstanced() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawTransformFeedbackStreamInstanced(mode Enum, id Uint, stream Uint, instancecount Sizei) (err error) {
	if !Supports.DrawTransformFeedbackStreamInstanced() {
		err = fmt.Errorf("DrawTransformFeedbackStreamInstanced() function call not supported by current GL context.")
	} else {
		DrawTransformFeedbackStreamInstanced(mode, id, stream, instancecount)
		err = Util.Error("DrawTransformFeedbackStreamInstanced()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetActiveSubroutineUniformName() function.
func (_ GlSupport) GetActiveSubroutineUniformName() bool {
	return C.cancallGetActiveSubroutineUniformName() == 1
}

//	If Supports.GetActiveSubroutineUniformName() is true, calls GetActiveSubroutineUniformName() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetActiveSubroutineUniformName(program Uint, shadertype Enum, index Uint, bufsize Sizei, length *Sizei, name *Char) (err error) {
	if !Supports.GetActiveSubroutineUniformName() {
		err = fmt.Errorf("GetActiveSubroutineUniformName() function call not supported by current GL context.")
	} else {
		GetActiveSubroutineUniformName(program, shadertype, index, bufsize, length, name)
		err = Util.Error("GetActiveSubroutineUniformName()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetInteger64i_v() function.
func (_ GlSupport) GetInteger64i_v() bool {
	return C.cancallGetInteger64i_v() == 1
}

//	If Supports.GetInteger64i_v() is true, calls GetInteger64i_v() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetInteger64i_v(target Enum, index Uint, data *Int64) (err error) {
	if !Supports.GetInteger64i_v() {
		err = fmt.Errorf("GetInteger64i_v() function call not supported by current GL context.")
	} else {
		GetInteger64i_v(target, index, data)
		err = Util.Error("GetInteger64i_v()")
	}
	return
}

//	Returns whether the current GL context supports calling the DebugMessageInsert() function.
func (_ GlSupport) DebugMessageInsert() bool {
	return C.cancallDebugMessageInsert() == 1
}

//	If Supports.DebugMessageInsert() is true, calls DebugMessageInsert() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DebugMessageInsert(source Enum, type_ Enum, id Uint, severity Enum, length Sizei, buf *Char) (err error) {
	if !Supports.DebugMessageInsert() {
		err = fmt.Errorf("DebugMessageInsert() function call not supported by current GL context.")
	} else {
		DebugMessageInsert(source, type_, id, severity, length, buf)
		err = Util.Error("DebugMessageInsert()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform1f() function.
func (_ GlSupport) ProgramUniform1f() bool {
	return C.cancallProgramUniform1f() == 1
}

//	If Supports.ProgramUniform1f() is true, calls ProgramUniform1f() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform1f(program Uint, location Int, v0 Float) (err error) {
	if !Supports.ProgramUniform1f() {
		err = fmt.Errorf("ProgramUniform1f() function call not supported by current GL context.")
	} else {
		ProgramUniform1f(program, location, v0)
		err = Util.Error("ProgramUniform1f()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetSamplerParameterfv() function.
func (_ GlSupport) GetSamplerParameterfv() bool {
	return C.cancallGetSamplerParameterfv() == 1
}

//	If Supports.GetSamplerParameterfv() is true, calls GetSamplerParameterfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetSamplerParameterfv(sampler Uint, pname Enum, params *Float) (err error) {
	if !Supports.GetSamplerParameterfv() {
		err = fmt.Errorf("GetSamplerParameterfv() function call not supported by current GL context.")
	} else {
		GetSamplerParameterfv(sampler, pname, params)
		err = Util.Error("GetSamplerParameterfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexCoordP4ui() function.
func (_ GlSupport) TexCoordP4ui() bool {
	return C.cancallTexCoordP4ui() == 1
}

//	If Supports.TexCoordP4ui() is true, calls TexCoordP4ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexCoordP4ui(type_ Enum, coords Uint) (err error) {
	if !Supports.TexCoordP4ui() {
		err = fmt.Errorf("TexCoordP4ui() function call not supported by current GL context.")
	} else {
		TexCoordP4ui(type_, coords)
		err = Util.Error("TexCoordP4ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform1uiv() function.
func (_ GlSupport) ProgramUniform1uiv() bool {
	return C.cancallProgramUniform1uiv() == 1
}

//	If Supports.ProgramUniform1uiv() is true, calls ProgramUniform1uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform1uiv(program Uint, location Int, count Sizei, value *Uint) (err error) {
	if !Supports.ProgramUniform1uiv() {
		err = fmt.Errorf("ProgramUniform1uiv() function call not supported by current GL context.")
	} else {
		ProgramUniform1uiv(program, location, count, value)
		err = Util.Error("ProgramUniform1uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteSamplers() function.
func (_ GlSupport) DeleteSamplers() bool {
	return C.cancallDeleteSamplers() == 1
}

//	If Supports.DeleteSamplers() is true, calls DeleteSamplers() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteSamplers(count Sizei, samplers *Uint) (err error) {
	if !Supports.DeleteSamplers() {
		err = fmt.Errorf("DeleteSamplers() function call not supported by current GL context.")
	} else {
		DeleteSamplers(count, samplers)
		err = Util.Error("DeleteSamplers()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsShader() function.
func (_ GlSupport) IsShader() bool {
	return C.cancallIsShader() == 1
}

//	If Supports.IsShader() is true, calls IsShader() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsShader(shader Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsShader() {
		err = fmt.Errorf("IsShader() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsShader(shader)
		err = Util.Error("IsShader()")
	}
	return
}

//	Returns whether the current GL context supports calling the CopyTexSubImage3D() function.
func (_ GlSupport) CopyTexSubImage3D() bool {
	return C.cancallCopyTexSubImage3D() == 1
}

//	If Supports.CopyTexSubImage3D() is true, calls CopyTexSubImage3D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CopyTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, x Int, y Int, width Sizei, height Sizei) (err error) {
	if !Supports.CopyTexSubImage3D() {
		err = fmt.Errorf("CopyTexSubImage3D() function call not supported by current GL context.")
	} else {
		CopyTexSubImage3D(target, level, xoffset, yoffset, zoffset, x, y, width, height)
		err = Util.Error("CopyTexSubImage3D()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetUniformuiv() function.
func (_ GlSupport) GetUniformuiv() bool {
	return C.cancallGetUniformuiv() == 1
}

//	If Supports.GetUniformuiv() is true, calls GetUniformuiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetUniformuiv(program Uint, location Int, params *Uint) (err error) {
	if !Supports.GetUniformuiv() {
		err = fmt.Errorf("GetUniformuiv() function call not supported by current GL context.")
	} else {
		GetUniformuiv(program, location, params)
		err = Util.Error("GetUniformuiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribL2d() function.
func (_ GlSupport) VertexAttribL2d() bool {
	return C.cancallVertexAttribL2d() == 1
}

//	If Supports.VertexAttribL2d() is true, calls VertexAttribL2d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribL2d(index Uint, x Double, y Double) (err error) {
	if !Supports.VertexAttribL2d() {
		err = fmt.Errorf("VertexAttribL2d() function call not supported by current GL context.")
	} else {
		VertexAttribL2d(index, x, y)
		err = Util.Error("VertexAttribL2d()")
	}
	return
}

//	Returns whether the current GL context supports calling the SecondaryColorP3ui() function.
func (_ GlSupport) SecondaryColorP3ui() bool {
	return C.cancallSecondaryColorP3ui() == 1
}

//	If Supports.SecondaryColorP3ui() is true, calls SecondaryColorP3ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) SecondaryColorP3ui(type_ Enum, color Uint) (err error) {
	if !Supports.SecondaryColorP3ui() {
		err = fmt.Errorf("SecondaryColorP3ui() function call not supported by current GL context.")
	} else {
		SecondaryColorP3ui(type_, color)
		err = Util.Error("SecondaryColorP3ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsSampler() function.
func (_ GlSupport) IsSampler() bool {
	return C.cancallIsSampler() == 1
}

//	If Supports.IsSampler() is true, calls IsSampler() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsSampler(sampler Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsSampler() {
		err = fmt.Errorf("IsSampler() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsSampler(sampler)
		err = Util.Error("IsSampler()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform2dv() function.
func (_ GlSupport) Uniform2dv() bool {
	return C.cancallUniform2dv() == 1
}

//	If Supports.Uniform2dv() is true, calls Uniform2dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform2dv(location Int, count Sizei, value *Double) (err error) {
	if !Supports.Uniform2dv() {
		err = fmt.Errorf("Uniform2dv() function call not supported by current GL context.")
	} else {
		Uniform2dv(location, count, value)
		err = Util.Error("Uniform2dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DepthMask() function.
func (_ GlSupport) DepthMask() bool {
	return C.cancallDepthMask() == 1
}

//	If Supports.DepthMask() is true, calls DepthMask() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DepthMask(flag Boolean) (err error) {
	if !Supports.DepthMask() {
		err = fmt.Errorf("DepthMask() function call not supported by current GL context.")
	} else {
		DepthMask(flag)
		err = Util.Error("DepthMask()")
	}
	return
}

//	Returns whether the current GL context supports calling the PushDebugGroup() function.
func (_ GlSupport) PushDebugGroup() bool {
	return C.cancallPushDebugGroup() == 1
}

//	If Supports.PushDebugGroup() is true, calls PushDebugGroup() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PushDebugGroup(source Enum, id Uint, length Sizei, message *Char) (err error) {
	if !Supports.PushDebugGroup() {
		err = fmt.Errorf("PushDebugGroup() function call not supported by current GL context.")
	} else {
		PushDebugGroup(source, id, length, message)
		err = Util.Error("PushDebugGroup()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform1dv() function.
func (_ GlSupport) ProgramUniform1dv() bool {
	return C.cancallProgramUniform1dv() == 1
}

//	If Supports.ProgramUniform1dv() is true, calls ProgramUniform1dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform1dv(program Uint, location Int, count Sizei, value *Double) (err error) {
	if !Supports.ProgramUniform1dv() {
		err = fmt.Errorf("ProgramUniform1dv() function call not supported by current GL context.")
	} else {
		ProgramUniform1dv(program, location, count, value)
		err = Util.Error("ProgramUniform1dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiDrawArraysIndirect() function.
func (_ GlSupport) MultiDrawArraysIndirect() bool {
	return C.cancallMultiDrawArraysIndirect() == 1
}

//	If Supports.MultiDrawArraysIndirect() is true, calls MultiDrawArraysIndirect() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiDrawArraysIndirect(mode Enum, indirect Ptr, drawcount Sizei, stride Sizei) (err error) {
	if !Supports.MultiDrawArraysIndirect() {
		err = fmt.Errorf("MultiDrawArraysIndirect() function call not supported by current GL context.")
	} else {
		MultiDrawArraysIndirect(mode, indirect, drawcount, stride)
		err = Util.Error("MultiDrawArraysIndirect()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiTexCoordP2uiv() function.
func (_ GlSupport) MultiTexCoordP2uiv() bool {
	return C.cancallMultiTexCoordP2uiv() == 1
}

//	If Supports.MultiTexCoordP2uiv() is true, calls MultiTexCoordP2uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiTexCoordP2uiv(texture Enum, type_ Enum, coords *Uint) (err error) {
	if !Supports.MultiTexCoordP2uiv() {
		err = fmt.Errorf("MultiTexCoordP2uiv() function call not supported by current GL context.")
	} else {
		MultiTexCoordP2uiv(texture, type_, coords)
		err = Util.Error("MultiTexCoordP2uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the InvalidateSubFramebuffer() function.
func (_ GlSupport) InvalidateSubFramebuffer() bool {
	return C.cancallInvalidateSubFramebuffer() == 1
}

//	If Supports.InvalidateSubFramebuffer() is true, calls InvalidateSubFramebuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) InvalidateSubFramebuffer(target Enum, numAttachments Sizei, attachments *Enum, x Int, y Int, width Sizei, height Sizei) (err error) {
	if !Supports.InvalidateSubFramebuffer() {
		err = fmt.Errorf("InvalidateSubFramebuffer() function call not supported by current GL context.")
	} else {
		InvalidateSubFramebuffer(target, numAttachments, attachments, x, y, width, height)
		err = Util.Error("InvalidateSubFramebuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform3i() function.
func (_ GlSupport) Uniform3i() bool {
	return C.cancallUniform3i() == 1
}

//	If Supports.Uniform3i() is true, calls Uniform3i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform3i(location Int, v0 Int, v1 Int, v2 Int) (err error) {
	if !Supports.Uniform3i() {
		err = fmt.Errorf("Uniform3i() function call not supported by current GL context.")
	} else {
		Uniform3i(location, v0, v1, v2)
		err = Util.Error("Uniform3i()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform2i() function.
func (_ GlSupport) ProgramUniform2i() bool {
	return C.cancallProgramUniform2i() == 1
}

//	If Supports.ProgramUniform2i() is true, calls ProgramUniform2i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform2i(program Uint, location Int, v0 Int, v1 Int) (err error) {
	if !Supports.ProgramUniform2i() {
		err = fmt.Errorf("ProgramUniform2i() function call not supported by current GL context.")
	} else {
		ProgramUniform2i(program, location, v0, v1)
		err = Util.Error("ProgramUniform2i()")
	}
	return
}

//	Returns whether the current GL context supports calling the ShaderSource() function.
func (_ GlSupport) ShaderSource() bool {
	return C.cancallShaderSource() == 1
}

//	If Supports.ShaderSource() is true, calls ShaderSource() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ShaderSource(shader Uint, count Sizei, string_ **Char, length *Int) (err error) {
	if !Supports.ShaderSource() {
		err = fmt.Errorf("ShaderSource() function call not supported by current GL context.")
	} else {
		ShaderSource(shader, count, string_, length)
		err = Util.Error("ShaderSource()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexCoordP1uiv() function.
func (_ GlSupport) TexCoordP1uiv() bool {
	return C.cancallTexCoordP1uiv() == 1
}

//	If Supports.TexCoordP1uiv() is true, calls TexCoordP1uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexCoordP1uiv(type_ Enum, coords *Uint) (err error) {
	if !Supports.TexCoordP1uiv() {
		err = fmt.Errorf("TexCoordP1uiv() function call not supported by current GL context.")
	} else {
		TexCoordP1uiv(type_, coords)
		err = Util.Error("TexCoordP1uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformBlockBinding() function.
func (_ GlSupport) UniformBlockBinding() bool {
	return C.cancallUniformBlockBinding() == 1
}

//	If Supports.UniformBlockBinding() is true, calls UniformBlockBinding() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformBlockBinding(program Uint, uniformBlockIndex Uint, uniformBlockBinding Uint) (err error) {
	if !Supports.UniformBlockBinding() {
		err = fmt.Errorf("UniformBlockBinding() function call not supported by current GL context.")
	} else {
		UniformBlockBinding(program, uniformBlockIndex, uniformBlockBinding)
		err = Util.Error("UniformBlockBinding()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform4fv() function.
func (_ GlSupport) Uniform4fv() bool {
	return C.cancallUniform4fv() == 1
}

//	If Supports.Uniform4fv() is true, calls Uniform4fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform4fv(location Int, count Sizei, value *Float) (err error) {
	if !Supports.Uniform4fv() {
		err = fmt.Errorf("Uniform4fv() function call not supported by current GL context.")
	} else {
		Uniform4fv(location, count, value)
		err = Util.Error("Uniform4fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetVertexAttribdv() function.
func (_ GlSupport) GetVertexAttribdv() bool {
	return C.cancallGetVertexAttribdv() == 1
}

//	If Supports.GetVertexAttribdv() is true, calls GetVertexAttribdv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetVertexAttribdv(index Uint, pname Enum, params *Double) (err error) {
	if !Supports.GetVertexAttribdv() {
		err = fmt.Errorf("GetVertexAttribdv() function call not supported by current GL context.")
	} else {
		GetVertexAttribdv(index, pname, params)
		err = Util.Error("GetVertexAttribdv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GenBuffers() function.
func (_ GlSupport) GenBuffers() bool {
	return C.cancallGenBuffers() == 1
}

//	If Supports.GenBuffers() is true, calls GenBuffers() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GenBuffers(n Sizei, buffers *Uint) (err error) {
	if !Supports.GenBuffers() {
		err = fmt.Errorf("GenBuffers() function call not supported by current GL context.")
	} else {
		GenBuffers(n, buffers)
		err = Util.Error("GenBuffers()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribP4ui() function.
func (_ GlSupport) VertexAttribP4ui() bool {
	return C.cancallVertexAttribP4ui() == 1
}

//	If Supports.VertexAttribP4ui() is true, calls VertexAttribP4ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribP4ui(index Uint, type_ Enum, normalized Boolean, value Uint) (err error) {
	if !Supports.VertexAttribP4ui() {
		err = fmt.Errorf("VertexAttribP4ui() function call not supported by current GL context.")
	} else {
		VertexAttribP4ui(index, type_, normalized, value)
		err = Util.Error("VertexAttribP4ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetInternalformativ() function.
func (_ GlSupport) GetInternalformativ() bool {
	return C.cancallGetInternalformativ() == 1
}

//	If Supports.GetInternalformativ() is true, calls GetInternalformativ() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetInternalformativ(target Enum, internalformat Enum, pname Enum, bufSize Sizei, params *Int) (err error) {
	if !Supports.GetInternalformativ() {
		err = fmt.Errorf("GetInternalformativ() function call not supported by current GL context.")
	} else {
		GetInternalformativ(target, internalformat, pname, bufSize, params)
		err = Util.Error("GetInternalformativ()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix4x2dv() function.
func (_ GlSupport) ProgramUniformMatrix4x2dv() bool {
	return C.cancallProgramUniformMatrix4x2dv() == 1
}

//	If Supports.ProgramUniformMatrix4x2dv() is true, calls ProgramUniformMatrix4x2dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix4x2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.ProgramUniformMatrix4x2dv() {
		err = fmt.Errorf("ProgramUniformMatrix4x2dv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix4x2dv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix4x2dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the CompileShader() function.
func (_ GlSupport) CompileShader() bool {
	return C.cancallCompileShader() == 1
}

//	If Supports.CompileShader() is true, calls CompileShader() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CompileShader(shader Uint) (err error) {
	if !Supports.CompileShader() {
		err = fmt.Errorf("CompileShader() function call not supported by current GL context.")
	} else {
		CompileShader(shader)
		err = Util.Error("CompileShader()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindRenderbuffer() function.
func (_ GlSupport) BindRenderbuffer() bool {
	return C.cancallBindRenderbuffer() == 1
}

//	If Supports.BindRenderbuffer() is true, calls BindRenderbuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindRenderbuffer(target Enum, renderbuffer Uint) (err error) {
	if !Supports.BindRenderbuffer() {
		err = fmt.Errorf("BindRenderbuffer() function call not supported by current GL context.")
	} else {
		BindRenderbuffer(target, renderbuffer)
		err = Util.Error("BindRenderbuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the CompressedTexSubImage1D() function.
func (_ GlSupport) CompressedTexSubImage1D() bool {
	return C.cancallCompressedTexSubImage1D() == 1
}

//	If Supports.CompressedTexSubImage1D() is true, calls CompressedTexSubImage1D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CompressedTexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, imageSize Sizei, data Ptr) (err error) {
	if !Supports.CompressedTexSubImage1D() {
		err = fmt.Errorf("CompressedTexSubImage1D() function call not supported by current GL context.")
	} else {
		CompressedTexSubImage1D(target, level, xoffset, width, format, imageSize, data)
		err = Util.Error("CompressedTexSubImage1D()")
	}
	return
}

//	Returns whether the current GL context supports calling the BlendFunc() function.
func (_ GlSupport) BlendFunc() bool {
	return C.cancallBlendFunc() == 1
}

//	If Supports.BlendFunc() is true, calls BlendFunc() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BlendFunc(sfactor Enum, dfactor Enum) (err error) {
	if !Supports.BlendFunc() {
		err = fmt.Errorf("BlendFunc() function call not supported by current GL context.")
	} else {
		BlendFunc(sfactor, dfactor)
		err = Util.Error("BlendFunc()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexParameterfv() function.
func (_ GlSupport) TexParameterfv() bool {
	return C.cancallTexParameterfv() == 1
}

//	If Supports.TexParameterfv() is true, calls TexParameterfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexParameterfv(target Enum, pname Enum, params *Float) (err error) {
	if !Supports.TexParameterfv() {
		err = fmt.Errorf("TexParameterfv() function call not supported by current GL context.")
	} else {
		TexParameterfv(target, pname, params)
		err = Util.Error("TexParameterfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the InvalidateFramebuffer() function.
func (_ GlSupport) InvalidateFramebuffer() bool {
	return C.cancallInvalidateFramebuffer() == 1
}

//	If Supports.InvalidateFramebuffer() is true, calls InvalidateFramebuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) InvalidateFramebuffer(target Enum, numAttachments Sizei, attachments *Enum) (err error) {
	if !Supports.InvalidateFramebuffer() {
		err = fmt.Errorf("InvalidateFramebuffer() function call not supported by current GL context.")
	} else {
		InvalidateFramebuffer(target, numAttachments, attachments)
		err = Util.Error("InvalidateFramebuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindTransformFeedback() function.
func (_ GlSupport) BindTransformFeedback() bool {
	return C.cancallBindTransformFeedback() == 1
}

//	If Supports.BindTransformFeedback() is true, calls BindTransformFeedback() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindTransformFeedback(target Enum, id Uint) (err error) {
	if !Supports.BindTransformFeedback() {
		err = fmt.Errorf("BindTransformFeedback() function call not supported by current GL context.")
	} else {
		BindTransformFeedback(target, id)
		err = Util.Error("BindTransformFeedback()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetFragDataIndex() function.
func (_ GlSupport) GetFragDataIndex() bool {
	return C.cancallGetFragDataIndex() == 1
}

//	If Supports.GetFragDataIndex() is true, calls GetFragDataIndex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetFragDataIndex(program Uint, name *Char) (err error, tryRetVal__ Int) {
	if !Supports.GetFragDataIndex() {
		err = fmt.Errorf("GetFragDataIndex() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetFragDataIndex(program, name)
		err = Util.Error("GetFragDataIndex()")
	}
	return
}

//	Returns whether the current GL context supports calling the TransformFeedbackVaryings() function.
func (_ GlSupport) TransformFeedbackVaryings() bool {
	return C.cancallTransformFeedbackVaryings() == 1
}

//	If Supports.TransformFeedbackVaryings() is true, calls TransformFeedbackVaryings() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TransformFeedbackVaryings(program Uint, count Sizei, varyings **Char, bufferMode Enum) (err error) {
	if !Supports.TransformFeedbackVaryings() {
		err = fmt.Errorf("TransformFeedbackVaryings() function call not supported by current GL context.")
	} else {
		TransformFeedbackVaryings(program, count, varyings, bufferMode)
		err = Util.Error("TransformFeedbackVaryings()")
	}
	return
}

//	Returns whether the current GL context supports calling the RenderbufferStorage() function.
func (_ GlSupport) RenderbufferStorage() bool {
	return C.cancallRenderbufferStorage() == 1
}

//	If Supports.RenderbufferStorage() is true, calls RenderbufferStorage() and yields the err returned by Util.Error(), if any.
func (_ GlTry) RenderbufferStorage(target Enum, internalformat Enum, width Sizei, height Sizei) (err error) {
	if !Supports.RenderbufferStorage() {
		err = fmt.Errorf("RenderbufferStorage() function call not supported by current GL context.")
	} else {
		RenderbufferStorage(target, internalformat, width, height)
		err = Util.Error("RenderbufferStorage()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetUniformBlockIndex() function.
func (_ GlSupport) GetUniformBlockIndex() bool {
	return C.cancallGetUniformBlockIndex() == 1
}

//	If Supports.GetUniformBlockIndex() is true, calls GetUniformBlockIndex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetUniformBlockIndex(program Uint, uniformBlockName *Char) (err error, tryRetVal__ Uint) {
	if !Supports.GetUniformBlockIndex() {
		err = fmt.Errorf("GetUniformBlockIndex() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetUniformBlockIndex(program, uniformBlockName)
		err = Util.Error("GetUniformBlockIndex()")
	}
	return
}

//	Returns whether the current GL context supports calling the ObjectLabel() function.
func (_ GlSupport) ObjectLabel() bool {
	return C.cancallObjectLabel() == 1
}

//	If Supports.ObjectLabel() is true, calls ObjectLabel() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ObjectLabel(identifier Enum, name Uint, length Sizei, label *Char) (err error) {
	if !Supports.ObjectLabel() {
		err = fmt.Errorf("ObjectLabel() function call not supported by current GL context.")
	} else {
		ObjectLabel(identifier, name, length, label)
		err = Util.Error("ObjectLabel()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetTexLevelParameterfv() function.
func (_ GlSupport) GetTexLevelParameterfv() bool {
	return C.cancallGetTexLevelParameterfv() == 1
}

//	If Supports.GetTexLevelParameterfv() is true, calls GetTexLevelParameterfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetTexLevelParameterfv(target Enum, level Int, pname Enum, params *Float) (err error) {
	if !Supports.GetTexLevelParameterfv() {
		err = fmt.Errorf("GetTexLevelParameterfv() function call not supported by current GL context.")
	} else {
		GetTexLevelParameterfv(target, level, pname, params)
		err = Util.Error("GetTexLevelParameterfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix4x3fv() function.
func (_ GlSupport) ProgramUniformMatrix4x3fv() bool {
	return C.cancallProgramUniformMatrix4x3fv() == 1
}

//	If Supports.ProgramUniformMatrix4x3fv() is true, calls ProgramUniformMatrix4x3fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix4x3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.ProgramUniformMatrix4x3fv() {
		err = fmt.Errorf("ProgramUniformMatrix4x3fv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix4x3fv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix4x3fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetShaderiv() function.
func (_ GlSupport) GetShaderiv() bool {
	return C.cancallGetShaderiv() == 1
}

//	If Supports.GetShaderiv() is true, calls GetShaderiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetShaderiv(shader Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetShaderiv() {
		err = fmt.Errorf("GetShaderiv() function call not supported by current GL context.")
	} else {
		GetShaderiv(shader, pname, params)
		err = Util.Error("GetShaderiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform4ui() function.
func (_ GlSupport) ProgramUniform4ui() bool {
	return C.cancallProgramUniform4ui() == 1
}

//	If Supports.ProgramUniform4ui() is true, calls ProgramUniform4ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform4ui(program Uint, location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint) (err error) {
	if !Supports.ProgramUniform4ui() {
		err = fmt.Errorf("ProgramUniform4ui() function call not supported by current GL context.")
	} else {
		ProgramUniform4ui(program, location, v0, v1, v2, v3)
		err = Util.Error("ProgramUniform4ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiTexCoordP3uiv() function.
func (_ GlSupport) MultiTexCoordP3uiv() bool {
	return C.cancallMultiTexCoordP3uiv() == 1
}

//	If Supports.MultiTexCoordP3uiv() is true, calls MultiTexCoordP3uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiTexCoordP3uiv(texture Enum, type_ Enum, coords *Uint) (err error) {
	if !Supports.MultiTexCoordP3uiv() {
		err = fmt.Errorf("MultiTexCoordP3uiv() function call not supported by current GL context.")
	} else {
		MultiTexCoordP3uiv(texture, type_, coords)
		err = Util.Error("MultiTexCoordP3uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform3uiv() function.
func (_ GlSupport) ProgramUniform3uiv() bool {
	return C.cancallProgramUniform3uiv() == 1
}

//	If Supports.ProgramUniform3uiv() is true, calls ProgramUniform3uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform3uiv(program Uint, location Int, count Sizei, value *Uint) (err error) {
	if !Supports.ProgramUniform3uiv() {
		err = fmt.Errorf("ProgramUniform3uiv() function call not supported by current GL context.")
	} else {
		ProgramUniform3uiv(program, location, count, value)
		err = Util.Error("ProgramUniform3uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetPointerv() function.
func (_ GlSupport) GetPointerv() bool {
	return C.cancallGetPointerv() == 1
}

//	If Supports.GetPointerv() is true, calls GetPointerv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetPointerv(pname Enum, params *Ptr) (err error) {
	if !Supports.GetPointerv() {
		err = fmt.Errorf("GetPointerv() function call not supported by current GL context.")
	} else {
		GetPointerv(pname, params)
		err = Util.Error("GetPointerv()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiTexCoordP4uiv() function.
func (_ GlSupport) MultiTexCoordP4uiv() bool {
	return C.cancallMultiTexCoordP4uiv() == 1
}

//	If Supports.MultiTexCoordP4uiv() is true, calls MultiTexCoordP4uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiTexCoordP4uiv(texture Enum, type_ Enum, coords *Uint) (err error) {
	if !Supports.MultiTexCoordP4uiv() {
		err = fmt.Errorf("MultiTexCoordP4uiv() function call not supported by current GL context.")
	} else {
		MultiTexCoordP4uiv(texture, type_, coords)
		err = Util.Error("MultiTexCoordP4uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix4x2fv() function.
func (_ GlSupport) UniformMatrix4x2fv() bool {
	return C.cancallUniformMatrix4x2fv() == 1
}

//	If Supports.UniformMatrix4x2fv() is true, calls UniformMatrix4x2fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix4x2fv(location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.UniformMatrix4x2fv() {
		err = fmt.Errorf("UniformMatrix4x2fv() function call not supported by current GL context.")
	} else {
		UniformMatrix4x2fv(location, count, transpose, value)
		err = Util.Error("UniformMatrix4x2fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ShaderStorageBlockBinding() function.
func (_ GlSupport) ShaderStorageBlockBinding() bool {
	return C.cancallShaderStorageBlockBinding() == 1
}

//	If Supports.ShaderStorageBlockBinding() is true, calls ShaderStorageBlockBinding() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ShaderStorageBlockBinding(program Uint, storageBlockIndex Uint, storageBlockBinding Uint) (err error) {
	if !Supports.ShaderStorageBlockBinding() {
		err = fmt.Errorf("ShaderStorageBlockBinding() function call not supported by current GL context.")
	} else {
		ShaderStorageBlockBinding(program, storageBlockIndex, storageBlockBinding)
		err = Util.Error("ShaderStorageBlockBinding()")
	}
	return
}

//	Returns whether the current GL context supports calling the ShaderBinary() function.
func (_ GlSupport) ShaderBinary() bool {
	return C.cancallShaderBinary() == 1
}

//	If Supports.ShaderBinary() is true, calls ShaderBinary() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ShaderBinary(count Sizei, shaders *Uint, binaryformat Enum, binary Ptr, length Sizei) (err error) {
	if !Supports.ShaderBinary() {
		err = fmt.Errorf("ShaderBinary() function call not supported by current GL context.")
	} else {
		ShaderBinary(count, shaders, binaryformat, binary, length)
		err = Util.Error("ShaderBinary()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform1d() function.
func (_ GlSupport) Uniform1d() bool {
	return C.cancallUniform1d() == 1
}

//	If Supports.Uniform1d() is true, calls Uniform1d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform1d(location Int, x Double) (err error) {
	if !Supports.Uniform1d() {
		err = fmt.Errorf("Uniform1d() function call not supported by current GL context.")
	} else {
		Uniform1d(location, x)
		err = Util.Error("Uniform1d()")
	}
	return
}

//	Returns whether the current GL context supports calling the BlendEquationSeparatei() function.
func (_ GlSupport) BlendEquationSeparatei() bool {
	return C.cancallBlendEquationSeparatei() == 1
}

//	If Supports.BlendEquationSeparatei() is true, calls BlendEquationSeparatei() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BlendEquationSeparatei(buf Uint, modeRGB Enum, modeAlpha Enum) (err error) {
	if !Supports.BlendEquationSeparatei() {
		err = fmt.Errorf("BlendEquationSeparatei() function call not supported by current GL context.")
	} else {
		BlendEquationSeparatei(buf, modeRGB, modeAlpha)
		err = Util.Error("BlendEquationSeparatei()")
	}
	return
}

//	Returns whether the current GL context supports calling the GenQueries() function.
func (_ GlSupport) GenQueries() bool {
	return C.cancallGenQueries() == 1
}

//	If Supports.GenQueries() is true, calls GenQueries() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GenQueries(n Sizei, ids *Uint) (err error) {
	if !Supports.GenQueries() {
		err = fmt.Errorf("GenQueries() function call not supported by current GL context.")
	} else {
		GenQueries(n, ids)
		err = Util.Error("GenQueries()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClampColor() function.
func (_ GlSupport) ClampColor() bool {
	return C.cancallClampColor() == 1
}

//	If Supports.ClampColor() is true, calls ClampColor() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClampColor(target Enum, clamp Enum) (err error) {
	if !Supports.ClampColor() {
		err = fmt.Errorf("ClampColor() function call not supported by current GL context.")
	} else {
		ClampColor(target, clamp)
		err = Util.Error("ClampColor()")
	}
	return
}

//	Returns whether the current GL context supports calling the PauseTransformFeedback() function.
func (_ GlSupport) PauseTransformFeedback() bool {
	return C.cancallPauseTransformFeedback() == 1
}

//	If Supports.PauseTransformFeedback() is true, calls PauseTransformFeedback() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PauseTransformFeedback() (err error) {
	if !Supports.PauseTransformFeedback() {
		err = fmt.Errorf("PauseTransformFeedback() function call not supported by current GL context.")
	} else {
		PauseTransformFeedback()
		err = Util.Error("PauseTransformFeedback()")
	}
	return
}

//	Returns whether the current GL context supports calling the LinkProgram() function.
func (_ GlSupport) LinkProgram() bool {
	return C.cancallLinkProgram() == 1
}

//	If Supports.LinkProgram() is true, calls LinkProgram() and yields the err returned by Util.Error(), if any.
func (_ GlTry) LinkProgram(program Uint) (err error) {
	if !Supports.LinkProgram() {
		err = fmt.Errorf("LinkProgram() function call not supported by current GL context.")
	} else {
		LinkProgram(program)
		err = Util.Error("LinkProgram()")
	}
	return
}

//	Returns whether the current GL context supports calling the CheckFramebufferStatus() function.
func (_ GlSupport) CheckFramebufferStatus() bool {
	return C.cancallCheckFramebufferStatus() == 1
}

//	If Supports.CheckFramebufferStatus() is true, calls CheckFramebufferStatus() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CheckFramebufferStatus(target Enum) (err error, tryRetVal__ Enum) {
	if !Supports.CheckFramebufferStatus() {
		err = fmt.Errorf("CheckFramebufferStatus() function call not supported by current GL context.")
	} else {
		tryRetVal__ = CheckFramebufferStatus(target)
		err = Util.Error("CheckFramebufferStatus()")
	}
	return
}

//	Returns whether the current GL context supports calling the PointParameteri() function.
func (_ GlSupport) PointParameteri() bool {
	return C.cancallPointParameteri() == 1
}

//	If Supports.PointParameteri() is true, calls PointParameteri() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PointParameteri(pname Enum, param Int) (err error) {
	if !Supports.PointParameteri() {
		err = fmt.Errorf("PointParameteri() function call not supported by current GL context.")
	} else {
		PointParameteri(pname, param)
		err = Util.Error("PointParameteri()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI4usv() function.
func (_ GlSupport) VertexAttribI4usv() bool {
	return C.cancallVertexAttribI4usv() == 1
}

//	If Supports.VertexAttribI4usv() is true, calls VertexAttribI4usv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI4usv(index Uint, v *Ushort) (err error) {
	if !Supports.VertexAttribI4usv() {
		err = fmt.Errorf("VertexAttribI4usv() function call not supported by current GL context.")
	} else {
		VertexAttribI4usv(index, v)
		err = Util.Error("VertexAttribI4usv()")
	}
	return
}

//	Returns whether the current GL context supports calling the PointParameteriv() function.
func (_ GlSupport) PointParameteriv() bool {
	return C.cancallPointParameteriv() == 1
}

//	If Supports.PointParameteriv() is true, calls PointParameteriv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PointParameteriv(pname Enum, params *Int) (err error) {
	if !Supports.PointParameteriv() {
		err = fmt.Errorf("PointParameteriv() function call not supported by current GL context.")
	} else {
		PointParameteriv(pname, params)
		err = Util.Error("PointParameteriv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribIPointer() function.
func (_ GlSupport) VertexAttribIPointer() bool {
	return C.cancallVertexAttribIPointer() == 1
}

//	If Supports.VertexAttribIPointer() is true, calls VertexAttribIPointer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribIPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Ptr) (err error) {
	if !Supports.VertexAttribIPointer() {
		err = fmt.Errorf("VertexAttribIPointer() function call not supported by current GL context.")
	} else {
		VertexAttribIPointer(index, size, type_, stride, pointer)
		err = Util.Error("VertexAttribIPointer()")
	}
	return
}

//	Returns whether the current GL context supports calling the GenTransformFeedbacks() function.
func (_ GlSupport) GenTransformFeedbacks() bool {
	return C.cancallGenTransformFeedbacks() == 1
}

//	If Supports.GenTransformFeedbacks() is true, calls GenTransformFeedbacks() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GenTransformFeedbacks(n Sizei, ids *Uint) (err error) {
	if !Supports.GenTransformFeedbacks() {
		err = fmt.Errorf("GenTransformFeedbacks() function call not supported by current GL context.")
	} else {
		GenTransformFeedbacks(n, ids)
		err = Util.Error("GenTransformFeedbacks()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform1uiv() function.
func (_ GlSupport) Uniform1uiv() bool {
	return C.cancallUniform1uiv() == 1
}

//	If Supports.Uniform1uiv() is true, calls Uniform1uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform1uiv(location Int, count Sizei, value *Uint) (err error) {
	if !Supports.Uniform1uiv() {
		err = fmt.Errorf("Uniform1uiv() function call not supported by current GL context.")
	} else {
		Uniform1uiv(location, count, value)
		err = Util.Error("Uniform1uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GenerateMipmap() function.
func (_ GlSupport) GenerateMipmap() bool {
	return C.cancallGenerateMipmap() == 1
}

//	If Supports.GenerateMipmap() is true, calls GenerateMipmap() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GenerateMipmap(target Enum) (err error) {
	if !Supports.GenerateMipmap() {
		err = fmt.Errorf("GenerateMipmap() function call not supported by current GL context.")
	} else {
		GenerateMipmap(target)
		err = Util.Error("GenerateMipmap()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetTexParameterfv() function.
func (_ GlSupport) GetTexParameterfv() bool {
	return C.cancallGetTexParameterfv() == 1
}

//	If Supports.GetTexParameterfv() is true, calls GetTexParameterfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetTexParameterfv(target Enum, pname Enum, params *Float) (err error) {
	if !Supports.GetTexParameterfv() {
		err = fmt.Errorf("GetTexParameterfv() function call not supported by current GL context.")
	} else {
		GetTexParameterfv(target, pname, params)
		err = Util.Error("GetTexParameterfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform2uiv() function.
func (_ GlSupport) Uniform2uiv() bool {
	return C.cancallUniform2uiv() == 1
}

//	If Supports.Uniform2uiv() is true, calls Uniform2uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform2uiv(location Int, count Sizei, value *Uint) (err error) {
	if !Supports.Uniform2uiv() {
		err = fmt.Errorf("Uniform2uiv() function call not supported by current GL context.")
	} else {
		Uniform2uiv(location, count, value)
		err = Util.Error("Uniform2uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformSubroutinesuiv() function.
func (_ GlSupport) UniformSubroutinesuiv() bool {
	return C.cancallUniformSubroutinesuiv() == 1
}

//	If Supports.UniformSubroutinesuiv() is true, calls UniformSubroutinesuiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformSubroutinesuiv(shadertype Enum, count Sizei, indices *Uint) (err error) {
	if !Supports.UniformSubroutinesuiv() {
		err = fmt.Errorf("UniformSubroutinesuiv() function call not supported by current GL context.")
	} else {
		UniformSubroutinesuiv(shadertype, count, indices)
		err = Util.Error("UniformSubroutinesuiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClearDepthf() function.
func (_ GlSupport) ClearDepthf() bool {
	return C.cancallClearDepthf() == 1
}

//	If Supports.ClearDepthf() is true, calls ClearDepthf() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClearDepthf(d Float) (err error) {
	if !Supports.ClearDepthf() {
		err = fmt.Errorf("ClearDepthf() function call not supported by current GL context.")
	} else {
		ClearDepthf(d)
		err = Util.Error("ClearDepthf()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteVertexArrays() function.
func (_ GlSupport) DeleteVertexArrays() bool {
	return C.cancallDeleteVertexArrays() == 1
}

//	If Supports.DeleteVertexArrays() is true, calls DeleteVertexArrays() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteVertexArrays(n Sizei, arrays *Uint) (err error) {
	if !Supports.DeleteVertexArrays() {
		err = fmt.Errorf("DeleteVertexArrays() function call not supported by current GL context.")
	} else {
		DeleteVertexArrays(n, arrays)
		err = Util.Error("DeleteVertexArrays()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsEnabled() function.
func (_ GlSupport) IsEnabled() bool {
	return C.cancallIsEnabled() == 1
}

//	If Supports.IsEnabled() is true, calls IsEnabled() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsEnabled(cap Enum) (err error, tryRetVal__ Boolean) {
	if !Supports.IsEnabled() {
		err = fmt.Errorf("IsEnabled() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsEnabled(cap)
		err = Util.Error("IsEnabled()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform4ui() function.
func (_ GlSupport) Uniform4ui() bool {
	return C.cancallUniform4ui() == 1
}

//	If Supports.Uniform4ui() is true, calls Uniform4ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform4ui(location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint) (err error) {
	if !Supports.Uniform4ui() {
		err = fmt.Errorf("Uniform4ui() function call not supported by current GL context.")
	} else {
		Uniform4ui(location, v0, v1, v2, v3)
		err = Util.Error("Uniform4ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetFramebufferAttachmentParameteriv() function.
func (_ GlSupport) GetFramebufferAttachmentParameteriv() bool {
	return C.cancallGetFramebufferAttachmentParameteriv() == 1
}

//	If Supports.GetFramebufferAttachmentParameteriv() is true, calls GetFramebufferAttachmentParameteriv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetFramebufferAttachmentParameteriv(target Enum, attachment Enum, pname Enum, params *Int) (err error) {
	if !Supports.GetFramebufferAttachmentParameteriv() {
		err = fmt.Errorf("GetFramebufferAttachmentParameteriv() function call not supported by current GL context.")
	} else {
		GetFramebufferAttachmentParameteriv(target, attachment, pname, params)
		err = Util.Error("GetFramebufferAttachmentParameteriv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform1dv() function.
func (_ GlSupport) Uniform1dv() bool {
	return C.cancallUniform1dv() == 1
}

//	If Supports.Uniform1dv() is true, calls Uniform1dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform1dv(location Int, count Sizei, value *Double) (err error) {
	if !Supports.Uniform1dv() {
		err = fmt.Errorf("Uniform1dv() function call not supported by current GL context.")
	} else {
		Uniform1dv(location, count, value)
		err = Util.Error("Uniform1dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ColorMask() function.
func (_ GlSupport) ColorMask() bool {
	return C.cancallColorMask() == 1
}

//	If Supports.ColorMask() is true, calls ColorMask() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean) (err error) {
	if !Supports.ColorMask() {
		err = fmt.Errorf("ColorMask() function call not supported by current GL context.")
	} else {
		ColorMask(red, green, blue, alpha)
		err = Util.Error("ColorMask()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetString() function.
func (_ GlSupport) GetString() bool {
	return C.cancallGetString() == 1
}

//	If Supports.GetString() is true, calls GetString() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetString(name Enum) (err error, tryRetVal__ *Ubyte) {
	if !Supports.GetString() {
		err = fmt.Errorf("GetString() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetString(name)
		err = Util.Error("GetString()")
	}
	return
}

//	Returns whether the current GL context supports calling the Hint() function.
func (_ GlSupport) Hint() bool {
	return C.cancallHint() == 1
}

//	If Supports.Hint() is true, calls Hint() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Hint(target Enum, mode Enum) (err error) {
	if !Supports.Hint() {
		err = fmt.Errorf("Hint() function call not supported by current GL context.")
	} else {
		Hint(target, mode)
		err = Util.Error("Hint()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClearBufferData() function.
func (_ GlSupport) ClearBufferData() bool {
	return C.cancallClearBufferData() == 1
}

//	If Supports.ClearBufferData() is true, calls ClearBufferData() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClearBufferData(target Enum, internalformat Enum, format Enum, type_ Enum, data Ptr) (err error) {
	if !Supports.ClearBufferData() {
		err = fmt.Errorf("ClearBufferData() function call not supported by current GL context.")
	} else {
		ClearBufferData(target, internalformat, format, type_, data)
		err = Util.Error("ClearBufferData()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform2d() function.
func (_ GlSupport) Uniform2d() bool {
	return C.cancallUniform2d() == 1
}

//	If Supports.Uniform2d() is true, calls Uniform2d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform2d(location Int, x Double, y Double) (err error) {
	if !Supports.Uniform2d() {
		err = fmt.Errorf("Uniform2d() function call not supported by current GL context.")
	} else {
		Uniform2d(location, x, y)
		err = Util.Error("Uniform2d()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetInteger64v() function.
func (_ GlSupport) GetInteger64v() bool {
	return C.cancallGetInteger64v() == 1
}

//	If Supports.GetInteger64v() is true, calls GetInteger64v() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetInteger64v(pname Enum, params *Int64) (err error) {
	if !Supports.GetInteger64v() {
		err = fmt.Errorf("GetInteger64v() function call not supported by current GL context.")
	} else {
		GetInteger64v(pname, params)
		err = Util.Error("GetInteger64v()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform1f() function.
func (_ GlSupport) Uniform1f() bool {
	return C.cancallUniform1f() == 1
}

//	If Supports.Uniform1f() is true, calls Uniform1f() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform1f(location Int, v0 Float) (err error) {
	if !Supports.Uniform1f() {
		err = fmt.Errorf("Uniform1f() function call not supported by current GL context.")
	} else {
		Uniform1f(location, v0)
		err = Util.Error("Uniform1f()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexBuffer() function.
func (_ GlSupport) TexBuffer() bool {
	return C.cancallTexBuffer() == 1
}

//	If Supports.TexBuffer() is true, calls TexBuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexBuffer(target Enum, internalformat Enum, buffer Uint) (err error) {
	if !Supports.TexBuffer() {
		err = fmt.Errorf("TexBuffer() function call not supported by current GL context.")
	} else {
		TexBuffer(target, internalformat, buffer)
		err = Util.Error("TexBuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the EndConditionalRender() function.
func (_ GlSupport) EndConditionalRender() bool {
	return C.cancallEndConditionalRender() == 1
}

//	If Supports.EndConditionalRender() is true, calls EndConditionalRender() and yields the err returned by Util.Error(), if any.
func (_ GlTry) EndConditionalRender() (err error) {
	if !Supports.EndConditionalRender() {
		err = fmt.Errorf("EndConditionalRender() function call not supported by current GL context.")
	} else {
		EndConditionalRender()
		err = Util.Error("EndConditionalRender()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexP3ui() function.
func (_ GlSupport) VertexP3ui() bool {
	return C.cancallVertexP3ui() == 1
}

//	If Supports.VertexP3ui() is true, calls VertexP3ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexP3ui(type_ Enum, value Uint) (err error) {
	if !Supports.VertexP3ui() {
		err = fmt.Errorf("VertexP3ui() function call not supported by current GL context.")
	} else {
		VertexP3ui(type_, value)
		err = Util.Error("VertexP3ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramInfoLog() function.
func (_ GlSupport) GetProgramInfoLog() bool {
	return C.cancallGetProgramInfoLog() == 1
}

//	If Supports.GetProgramInfoLog() is true, calls GetProgramInfoLog() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramInfoLog(program Uint, bufSize Sizei, length *Sizei, infoLog *Char) (err error) {
	if !Supports.GetProgramInfoLog() {
		err = fmt.Errorf("GetProgramInfoLog() function call not supported by current GL context.")
	} else {
		GetProgramInfoLog(program, bufSize, length, infoLog)
		err = Util.Error("GetProgramInfoLog()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI2i() function.
func (_ GlSupport) VertexAttribI2i() bool {
	return C.cancallVertexAttribI2i() == 1
}

//	If Supports.VertexAttribI2i() is true, calls VertexAttribI2i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI2i(index Uint, x Int, y Int) (err error) {
	if !Supports.VertexAttribI2i() {
		err = fmt.Errorf("VertexAttribI2i() function call not supported by current GL context.")
	} else {
		VertexAttribI2i(index, x, y)
		err = Util.Error("VertexAttribI2i()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsTransformFeedback() function.
func (_ GlSupport) IsTransformFeedback() bool {
	return C.cancallIsTransformFeedback() == 1
}

//	If Supports.IsTransformFeedback() is true, calls IsTransformFeedback() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsTransformFeedback(id Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsTransformFeedback() {
		err = fmt.Errorf("IsTransformFeedback() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsTransformFeedback(id)
		err = Util.Error("IsTransformFeedback()")
	}
	return
}

//	Returns whether the current GL context supports calling the LogicOp() function.
func (_ GlSupport) LogicOp() bool {
	return C.cancallLogicOp() == 1
}

//	If Supports.LogicOp() is true, calls LogicOp() and yields the err returned by Util.Error(), if any.
func (_ GlTry) LogicOp(opcode Enum) (err error) {
	if !Supports.LogicOp() {
		err = fmt.Errorf("LogicOp() function call not supported by current GL context.")
	} else {
		LogicOp(opcode)
		err = Util.Error("LogicOp()")
	}
	return
}

//	Returns whether the current GL context supports calling the DispatchComputeIndirect() function.
func (_ GlSupport) DispatchComputeIndirect() bool {
	return C.cancallDispatchComputeIndirect() == 1
}

//	If Supports.DispatchComputeIndirect() is true, calls DispatchComputeIndirect() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DispatchComputeIndirect(indirect Intptr) (err error) {
	if !Supports.DispatchComputeIndirect() {
		err = fmt.Errorf("DispatchComputeIndirect() function call not supported by current GL context.")
	} else {
		DispatchComputeIndirect(indirect)
		err = Util.Error("DispatchComputeIndirect()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindTexture() function.
func (_ GlSupport) BindTexture() bool {
	return C.cancallBindTexture() == 1
}

//	If Supports.BindTexture() is true, calls BindTexture() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindTexture(target Enum, texture Uint) (err error) {
	if !Supports.BindTexture() {
		err = fmt.Errorf("BindTexture() function call not supported by current GL context.")
	} else {
		BindTexture(target, texture)
		err = Util.Error("BindTexture()")
	}
	return
}

//	Returns whether the current GL context supports calling the DisableVertexAttribArray() function.
func (_ GlSupport) DisableVertexAttribArray() bool {
	return C.cancallDisableVertexAttribArray() == 1
}

//	If Supports.DisableVertexAttribArray() is true, calls DisableVertexAttribArray() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DisableVertexAttribArray(index Uint) (err error) {
	if !Supports.DisableVertexAttribArray() {
		err = fmt.Errorf("DisableVertexAttribArray() function call not supported by current GL context.")
	} else {
		DisableVertexAttribArray(index)
		err = Util.Error("DisableVertexAttribArray()")
	}
	return
}

//	Returns whether the current GL context supports calling the QueryCounter() function.
func (_ GlSupport) QueryCounter() bool {
	return C.cancallQueryCounter() == 1
}

//	If Supports.QueryCounter() is true, calls QueryCounter() and yields the err returned by Util.Error(), if any.
func (_ GlTry) QueryCounter(id Uint, target Enum) (err error) {
	if !Supports.QueryCounter() {
		err = fmt.Errorf("QueryCounter() function call not supported by current GL context.")
	} else {
		QueryCounter(id, target)
		err = Util.Error("QueryCounter()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform3f() function.
func (_ GlSupport) ProgramUniform3f() bool {
	return C.cancallProgramUniform3f() == 1
}

//	If Supports.ProgramUniform3f() is true, calls ProgramUniform3f() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform3f(program Uint, location Int, v0 Float, v1 Float, v2 Float) (err error) {
	if !Supports.ProgramUniform3f() {
		err = fmt.Errorf("ProgramUniform3f() function call not supported by current GL context.")
	} else {
		ProgramUniform3f(program, location, v0, v1, v2)
		err = Util.Error("ProgramUniform3f()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsEnabledi() function.
func (_ GlSupport) IsEnabledi() bool {
	return C.cancallIsEnabledi() == 1
}

//	If Supports.IsEnabledi() is true, calls IsEnabledi() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsEnabledi(target Enum, index Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsEnabledi() {
		err = fmt.Errorf("IsEnabledi() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsEnabledi(target, index)
		err = Util.Error("IsEnabledi()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribL3d() function.
func (_ GlSupport) VertexAttribL3d() bool {
	return C.cancallVertexAttribL3d() == 1
}

//	If Supports.VertexAttribL3d() is true, calls VertexAttribL3d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribL3d(index Uint, x Double, y Double, z Double) (err error) {
	if !Supports.VertexAttribL3d() {
		err = fmt.Errorf("VertexAttribL3d() function call not supported by current GL context.")
	} else {
		VertexAttribL3d(index, x, y, z)
		err = Util.Error("VertexAttribL3d()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform4dv() function.
func (_ GlSupport) Uniform4dv() bool {
	return C.cancallUniform4dv() == 1
}

//	If Supports.Uniform4dv() is true, calls Uniform4dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform4dv(location Int, count Sizei, value *Double) (err error) {
	if !Supports.Uniform4dv() {
		err = fmt.Errorf("Uniform4dv() function call not supported by current GL context.")
	} else {
		Uniform4dv(location, count, value)
		err = Util.Error("Uniform4dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindVertexArray() function.
func (_ GlSupport) BindVertexArray() bool {
	return C.cancallBindVertexArray() == 1
}

//	If Supports.BindVertexArray() is true, calls BindVertexArray() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindVertexArray(array Uint) (err error) {
	if !Supports.BindVertexArray() {
		err = fmt.Errorf("BindVertexArray() function call not supported by current GL context.")
	} else {
		BindVertexArray(array)
		err = Util.Error("BindVertexArray()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform3dv() function.
func (_ GlSupport) ProgramUniform3dv() bool {
	return C.cancallProgramUniform3dv() == 1
}

//	If Supports.ProgramUniform3dv() is true, calls ProgramUniform3dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform3dv(program Uint, location Int, count Sizei, value *Double) (err error) {
	if !Supports.ProgramUniform3dv() {
		err = fmt.Errorf("ProgramUniform3dv() function call not supported by current GL context.")
	} else {
		ProgramUniform3dv(program, location, count, value)
		err = Util.Error("ProgramUniform3dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetFramebufferParameteriv() function.
func (_ GlSupport) GetFramebufferParameteriv() bool {
	return C.cancallGetFramebufferParameteriv() == 1
}

//	If Supports.GetFramebufferParameteriv() is true, calls GetFramebufferParameteriv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetFramebufferParameteriv(target Enum, pname Enum, params *Int) (err error) {
	if !Supports.GetFramebufferParameteriv() {
		err = fmt.Errorf("GetFramebufferParameteriv() function call not supported by current GL context.")
	} else {
		GetFramebufferParameteriv(target, pname, params)
		err = Util.Error("GetFramebufferParameteriv()")
	}
	return
}

//	Returns whether the current GL context supports calling the CullFace() function.
func (_ GlSupport) CullFace() bool {
	return C.cancallCullFace() == 1
}

//	If Supports.CullFace() is true, calls CullFace() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CullFace(mode Enum) (err error) {
	if !Supports.CullFace() {
		err = fmt.Errorf("CullFace() function call not supported by current GL context.")
	} else {
		CullFace(mode)
		err = Util.Error("CullFace()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiDrawArrays() function.
func (_ GlSupport) MultiDrawArrays() bool {
	return C.cancallMultiDrawArrays() == 1
}

//	If Supports.MultiDrawArrays() is true, calls MultiDrawArrays() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiDrawArrays(mode Enum, first *Int, count *Sizei, drawcount Sizei) (err error) {
	if !Supports.MultiDrawArrays() {
		err = fmt.Errorf("MultiDrawArrays() function call not supported by current GL context.")
	} else {
		MultiDrawArrays(mode, first, count, drawcount)
		err = Util.Error("MultiDrawArrays()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix2x3dv() function.
func (_ GlSupport) ProgramUniformMatrix2x3dv() bool {
	return C.cancallProgramUniformMatrix2x3dv() == 1
}

//	If Supports.ProgramUniformMatrix2x3dv() is true, calls ProgramUniformMatrix2x3dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix2x3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.ProgramUniformMatrix2x3dv() {
		err = fmt.Errorf("ProgramUniformMatrix2x3dv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix2x3dv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix2x3dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ValidateProgramPipeline() function.
func (_ GlSupport) ValidateProgramPipeline() bool {
	return C.cancallValidateProgramPipeline() == 1
}

//	If Supports.ValidateProgramPipeline() is true, calls ValidateProgramPipeline() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ValidateProgramPipeline(pipeline Uint) (err error) {
	if !Supports.ValidateProgramPipeline() {
		err = fmt.Errorf("ValidateProgramPipeline() function call not supported by current GL context.")
	} else {
		ValidateProgramPipeline(pipeline)
		err = Util.Error("ValidateProgramPipeline()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix3x2fv() function.
func (_ GlSupport) ProgramUniformMatrix3x2fv() bool {
	return C.cancallProgramUniformMatrix3x2fv() == 1
}

//	If Supports.ProgramUniformMatrix3x2fv() is true, calls ProgramUniformMatrix3x2fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix3x2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.ProgramUniformMatrix3x2fv() {
		err = fmt.Errorf("ProgramUniformMatrix3x2fv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix3x2fv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix3x2fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI3i() function.
func (_ GlSupport) VertexAttribI3i() bool {
	return C.cancallVertexAttribI3i() == 1
}

//	If Supports.VertexAttribI3i() is true, calls VertexAttribI3i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI3i(index Uint, x Int, y Int, z Int) (err error) {
	if !Supports.VertexAttribI3i() {
		err = fmt.Errorf("VertexAttribI3i() function call not supported by current GL context.")
	} else {
		VertexAttribI3i(index, x, y, z)
		err = Util.Error("VertexAttribI3i()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramResourceiv() function.
func (_ GlSupport) GetProgramResourceiv() bool {
	return C.cancallGetProgramResourceiv() == 1
}

//	If Supports.GetProgramResourceiv() is true, calls GetProgramResourceiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramResourceiv(program Uint, programInterface Enum, index Uint, propCount Sizei, props *Enum, bufSize Sizei, length *Sizei, params *Int) (err error) {
	if !Supports.GetProgramResourceiv() {
		err = fmt.Errorf("GetProgramResourceiv() function call not supported by current GL context.")
	} else {
		GetProgramResourceiv(program, programInterface, index, propCount, props, bufSize, length, params)
		err = Util.Error("GetProgramResourceiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetUniformSubroutineuiv() function.
func (_ GlSupport) GetUniformSubroutineuiv() bool {
	return C.cancallGetUniformSubroutineuiv() == 1
}

//	If Supports.GetUniformSubroutineuiv() is true, calls GetUniformSubroutineuiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetUniformSubroutineuiv(shadertype Enum, location Int, params *Uint) (err error) {
	if !Supports.GetUniformSubroutineuiv() {
		err = fmt.Errorf("GetUniformSubroutineuiv() function call not supported by current GL context.")
	} else {
		GetUniformSubroutineuiv(shadertype, location, params)
		err = Util.Error("GetUniformSubroutineuiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform2iv() function.
func (_ GlSupport) ProgramUniform2iv() bool {
	return C.cancallProgramUniform2iv() == 1
}

//	If Supports.ProgramUniform2iv() is true, calls ProgramUniform2iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform2iv(program Uint, location Int, count Sizei, value *Int) (err error) {
	if !Supports.ProgramUniform2iv() {
		err = fmt.Errorf("ProgramUniform2iv() function call not supported by current GL context.")
	} else {
		ProgramUniform2iv(program, location, count, value)
		err = Util.Error("ProgramUniform2iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GenFramebuffers() function.
func (_ GlSupport) GenFramebuffers() bool {
	return C.cancallGenFramebuffers() == 1
}

//	If Supports.GenFramebuffers() is true, calls GenFramebuffers() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GenFramebuffers(n Sizei, framebuffers *Uint) (err error) {
	if !Supports.GenFramebuffers() {
		err = fmt.Errorf("GenFramebuffers() function call not supported by current GL context.")
	} else {
		GenFramebuffers(n, framebuffers)
		err = Util.Error("GenFramebuffers()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawTransformFeedbackInstanced() function.
func (_ GlSupport) DrawTransformFeedbackInstanced() bool {
	return C.cancallDrawTransformFeedbackInstanced() == 1
}

//	If Supports.DrawTransformFeedbackInstanced() is true, calls DrawTransformFeedbackInstanced() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawTransformFeedbackInstanced(mode Enum, id Uint, instancecount Sizei) (err error) {
	if !Supports.DrawTransformFeedbackInstanced() {
		err = fmt.Errorf("DrawTransformFeedbackInstanced() function call not supported by current GL context.")
	} else {
		DrawTransformFeedbackInstanced(mode, id, instancecount)
		err = Util.Error("DrawTransformFeedbackInstanced()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsBuffer() function.
func (_ GlSupport) IsBuffer() bool {
	return C.cancallIsBuffer() == 1
}

//	If Supports.IsBuffer() is true, calls IsBuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsBuffer(buffer Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsBuffer() {
		err = fmt.Errorf("IsBuffer() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsBuffer(buffer)
		err = Util.Error("IsBuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the AttachShader() function.
func (_ GlSupport) AttachShader() bool {
	return C.cancallAttachShader() == 1
}

//	If Supports.AttachShader() is true, calls AttachShader() and yields the err returned by Util.Error(), if any.
func (_ GlTry) AttachShader(program Uint, shader Uint) (err error) {
	if !Supports.AttachShader() {
		err = fmt.Errorf("AttachShader() function call not supported by current GL context.")
	} else {
		AttachShader(program, shader)
		err = Util.Error("AttachShader()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawElementsIndirect() function.
func (_ GlSupport) DrawElementsIndirect() bool {
	return C.cancallDrawElementsIndirect() == 1
}

//	If Supports.DrawElementsIndirect() is true, calls DrawElementsIndirect() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawElementsIndirect(mode Enum, type_ Enum, indirect Ptr) (err error) {
	if !Supports.DrawElementsIndirect() {
		err = fmt.Errorf("DrawElementsIndirect() function call not supported by current GL context.")
	} else {
		DrawElementsIndirect(mode, type_, indirect)
		err = Util.Error("DrawElementsIndirect()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform1ui() function.
func (_ GlSupport) ProgramUniform1ui() bool {
	return C.cancallProgramUniform1ui() == 1
}

//	If Supports.ProgramUniform1ui() is true, calls ProgramUniform1ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform1ui(program Uint, location Int, v0 Uint) (err error) {
	if !Supports.ProgramUniform1ui() {
		err = fmt.Errorf("ProgramUniform1ui() function call not supported by current GL context.")
	} else {
		ProgramUniform1ui(program, location, v0)
		err = Util.Error("ProgramUniform1ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetUniformIndices() function.
func (_ GlSupport) GetUniformIndices() bool {
	return C.cancallGetUniformIndices() == 1
}

//	If Supports.GetUniformIndices() is true, calls GetUniformIndices() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetUniformIndices(program Uint, uniformCount Sizei, uniformNames **Char, uniformIndices *Uint) (err error) {
	if !Supports.GetUniformIndices() {
		err = fmt.Errorf("GetUniformIndices() function call not supported by current GL context.")
	} else {
		GetUniformIndices(program, uniformCount, uniformNames, uniformIndices)
		err = Util.Error("GetUniformIndices()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform3f() function.
func (_ GlSupport) Uniform3f() bool {
	return C.cancallUniform3f() == 1
}

//	If Supports.Uniform3f() is true, calls Uniform3f() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform3f(location Int, v0 Float, v1 Float, v2 Float) (err error) {
	if !Supports.Uniform3f() {
		err = fmt.Errorf("Uniform3f() function call not supported by current GL context.")
	} else {
		Uniform3f(location, v0, v1, v2)
		err = Util.Error("Uniform3f()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindFragDataLocationIndexed() function.
func (_ GlSupport) BindFragDataLocationIndexed() bool {
	return C.cancallBindFragDataLocationIndexed() == 1
}

//	If Supports.BindFragDataLocationIndexed() is true, calls BindFragDataLocationIndexed() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindFragDataLocationIndexed(program Uint, colorNumber Uint, index Uint, name *Char) (err error) {
	if !Supports.BindFragDataLocationIndexed() {
		err = fmt.Errorf("BindFragDataLocationIndexed() function call not supported by current GL context.")
	} else {
		BindFragDataLocationIndexed(program, colorNumber, index, name)
		err = Util.Error("BindFragDataLocationIndexed()")
	}
	return
}

//	Returns whether the current GL context supports calling the TextureView() function.
func (_ GlSupport) TextureView() bool {
	return C.cancallTextureView() == 1
}

//	If Supports.TextureView() is true, calls TextureView() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TextureView(texture Uint, target Enum, origtexture Uint, internalformat Enum, minlevel Uint, numlevels Uint, minlayer Uint, numlayers Uint) (err error) {
	if !Supports.TextureView() {
		err = fmt.Errorf("TextureView() function call not supported by current GL context.")
	} else {
		TextureView(texture, target, origtexture, internalformat, minlevel, numlevels, minlayer, numlayers)
		err = Util.Error("TextureView()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetQueryiv() function.
func (_ GlSupport) GetQueryiv() bool {
	return C.cancallGetQueryiv() == 1
}

//	If Supports.GetQueryiv() is true, calls GetQueryiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetQueryiv(target Enum, pname Enum, params *Int) (err error) {
	if !Supports.GetQueryiv() {
		err = fmt.Errorf("GetQueryiv() function call not supported by current GL context.")
	} else {
		GetQueryiv(target, pname, params)
		err = Util.Error("GetQueryiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the FramebufferTexture1D() function.
func (_ GlSupport) FramebufferTexture1D() bool {
	return C.cancallFramebufferTexture1D() == 1
}

//	If Supports.FramebufferTexture1D() is true, calls FramebufferTexture1D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) FramebufferTexture1D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int) (err error) {
	if !Supports.FramebufferTexture1D() {
		err = fmt.Errorf("FramebufferTexture1D() function call not supported by current GL context.")
	} else {
		FramebufferTexture1D(target, attachment, textarget, texture, level)
		err = Util.Error("FramebufferTexture1D()")
	}
	return
}

//	Returns whether the current GL context supports calling the PointSize() function.
func (_ GlSupport) PointSize() bool {
	return C.cancallPointSize() == 1
}

//	If Supports.PointSize() is true, calls PointSize() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PointSize(size Float) (err error) {
	if !Supports.PointSize() {
		err = fmt.Errorf("PointSize() function call not supported by current GL context.")
	} else {
		PointSize(size)
		err = Util.Error("PointSize()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI4i() function.
func (_ GlSupport) VertexAttribI4i() bool {
	return C.cancallVertexAttribI4i() == 1
}

//	If Supports.VertexAttribI4i() is true, calls VertexAttribI4i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI4i(index Uint, x Int, y Int, z Int, w Int) (err error) {
	if !Supports.VertexAttribI4i() {
		err = fmt.Errorf("VertexAttribI4i() function call not supported by current GL context.")
	} else {
		VertexAttribI4i(index, x, y, z, w)
		err = Util.Error("VertexAttribI4i()")
	}
	return
}

//	Returns whether the current GL context supports calling the PopDebugGroup() function.
func (_ GlSupport) PopDebugGroup() bool {
	return C.cancallPopDebugGroup() == 1
}

//	If Supports.PopDebugGroup() is true, calls PopDebugGroup() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PopDebugGroup() (err error) {
	if !Supports.PopDebugGroup() {
		err = fmt.Errorf("PopDebugGroup() function call not supported by current GL context.")
	} else {
		PopDebugGroup()
		err = Util.Error("PopDebugGroup()")
	}
	return
}

//	Returns whether the current GL context supports calling the StencilFunc() function.
func (_ GlSupport) StencilFunc() bool {
	return C.cancallStencilFunc() == 1
}

//	If Supports.StencilFunc() is true, calls StencilFunc() and yields the err returned by Util.Error(), if any.
func (_ GlTry) StencilFunc(func_ Enum, ref Int, mask Uint) (err error) {
	if !Supports.StencilFunc() {
		err = fmt.Errorf("StencilFunc() function call not supported by current GL context.")
	} else {
		StencilFunc(func_, ref, mask)
		err = Util.Error("StencilFunc()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexImage3DMultisample() function.
func (_ GlSupport) TexImage3DMultisample() bool {
	return C.cancallTexImage3DMultisample() == 1
}

//	If Supports.TexImage3DMultisample() is true, calls TexImage3DMultisample() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexImage3DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean) (err error) {
	if !Supports.TexImage3DMultisample() {
		err = fmt.Errorf("TexImage3DMultisample() function call not supported by current GL context.")
	} else {
		TexImage3DMultisample(target, samples, internalformat, width, height, depth, fixedsamplelocations)
		err = Util.Error("TexImage3DMultisample()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix2x3fv() function.
func (_ GlSupport) ProgramUniformMatrix2x3fv() bool {
	return C.cancallProgramUniformMatrix2x3fv() == 1
}

//	If Supports.ProgramUniformMatrix2x3fv() is true, calls ProgramUniformMatrix2x3fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix2x3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.ProgramUniformMatrix2x3fv() {
		err = fmt.Errorf("ProgramUniformMatrix2x3fv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix2x3fv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix2x3fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribL1d() function.
func (_ GlSupport) VertexAttribL1d() bool {
	return C.cancallVertexAttribL1d() == 1
}

//	If Supports.VertexAttribL1d() is true, calls VertexAttribL1d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribL1d(index Uint, x Double) (err error) {
	if !Supports.VertexAttribL1d() {
		err = fmt.Errorf("VertexAttribL1d() function call not supported by current GL context.")
	} else {
		VertexAttribL1d(index, x)
		err = Util.Error("VertexAttribL1d()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClearBufferiv() function.
func (_ GlSupport) ClearBufferiv() bool {
	return C.cancallClearBufferiv() == 1
}

//	If Supports.ClearBufferiv() is true, calls ClearBufferiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClearBufferiv(buffer Enum, drawbuffer Int, value *Int) (err error) {
	if !Supports.ClearBufferiv() {
		err = fmt.Errorf("ClearBufferiv() function call not supported by current GL context.")
	} else {
		ClearBufferiv(buffer, drawbuffer, value)
		err = Util.Error("ClearBufferiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindVertexBuffer() function.
func (_ GlSupport) BindVertexBuffer() bool {
	return C.cancallBindVertexBuffer() == 1
}

//	If Supports.BindVertexBuffer() is true, calls BindVertexBuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindVertexBuffer(bindingindex Uint, buffer Uint, offset Intptr, stride Sizei) (err error) {
	if !Supports.BindVertexBuffer() {
		err = fmt.Errorf("BindVertexBuffer() function call not supported by current GL context.")
	} else {
		BindVertexBuffer(bindingindex, buffer, offset, stride)
		err = Util.Error("BindVertexBuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the Disablei() function.
func (_ GlSupport) Disablei() bool {
	return C.cancallDisablei() == 1
}

//	If Supports.Disablei() is true, calls Disablei() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Disablei(target Enum, index Uint) (err error) {
	if !Supports.Disablei() {
		err = fmt.Errorf("Disablei() function call not supported by current GL context.")
	} else {
		Disablei(target, index)
		err = Util.Error("Disablei()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribP1uiv() function.
func (_ GlSupport) VertexAttribP1uiv() bool {
	return C.cancallVertexAttribP1uiv() == 1
}

//	If Supports.VertexAttribP1uiv() is true, calls VertexAttribP1uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribP1uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) (err error) {
	if !Supports.VertexAttribP1uiv() {
		err = fmt.Errorf("VertexAttribP1uiv() function call not supported by current GL context.")
	} else {
		VertexAttribP1uiv(index, type_, normalized, value)
		err = Util.Error("VertexAttribP1uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform3dv() function.
func (_ GlSupport) Uniform3dv() bool {
	return C.cancallUniform3dv() == 1
}

//	If Supports.Uniform3dv() is true, calls Uniform3dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform3dv(location Int, count Sizei, value *Double) (err error) {
	if !Supports.Uniform3dv() {
		err = fmt.Errorf("Uniform3dv() function call not supported by current GL context.")
	} else {
		Uniform3dv(location, count, value)
		err = Util.Error("Uniform3dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the FramebufferTextureLayer() function.
func (_ GlSupport) FramebufferTextureLayer() bool {
	return C.cancallFramebufferTextureLayer() == 1
}

//	If Supports.FramebufferTextureLayer() is true, calls FramebufferTextureLayer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) FramebufferTextureLayer(target Enum, attachment Enum, texture Uint, level Int, layer Int) (err error) {
	if !Supports.FramebufferTextureLayer() {
		err = fmt.Errorf("FramebufferTextureLayer() function call not supported by current GL context.")
	} else {
		FramebufferTextureLayer(target, attachment, texture, level, layer)
		err = Util.Error("FramebufferTextureLayer()")
	}
	return
}

//	Returns whether the current GL context supports calling the FramebufferTexture2D() function.
func (_ GlSupport) FramebufferTexture2D() bool {
	return C.cancallFramebufferTexture2D() == 1
}

//	If Supports.FramebufferTexture2D() is true, calls FramebufferTexture2D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) FramebufferTexture2D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int) (err error) {
	if !Supports.FramebufferTexture2D() {
		err = fmt.Errorf("FramebufferTexture2D() function call not supported by current GL context.")
	} else {
		FramebufferTexture2D(target, attachment, textarget, texture, level)
		err = Util.Error("FramebufferTexture2D()")
	}
	return
}

//	Returns whether the current GL context supports calling the ReleaseShaderCompiler() function.
func (_ GlSupport) ReleaseShaderCompiler() bool {
	return C.cancallReleaseShaderCompiler() == 1
}

//	If Supports.ReleaseShaderCompiler() is true, calls ReleaseShaderCompiler() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ReleaseShaderCompiler() (err error) {
	if !Supports.ReleaseShaderCompiler() {
		err = fmt.Errorf("ReleaseShaderCompiler() function call not supported by current GL context.")
	} else {
		ReleaseShaderCompiler()
		err = Util.Error("ReleaseShaderCompiler()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix3x4fv() function.
func (_ GlSupport) ProgramUniformMatrix3x4fv() bool {
	return C.cancallProgramUniformMatrix3x4fv() == 1
}

//	If Supports.ProgramUniformMatrix3x4fv() is true, calls ProgramUniformMatrix3x4fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix3x4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.ProgramUniformMatrix3x4fv() {
		err = fmt.Errorf("ProgramUniformMatrix3x4fv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix3x4fv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix3x4fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix2x4dv() function.
func (_ GlSupport) ProgramUniformMatrix2x4dv() bool {
	return C.cancallProgramUniformMatrix2x4dv() == 1
}

//	If Supports.ProgramUniformMatrix2x4dv() is true, calls ProgramUniformMatrix2x4dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix2x4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.ProgramUniformMatrix2x4dv() {
		err = fmt.Errorf("ProgramUniformMatrix2x4dv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix2x4dv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix2x4dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the PatchParameteri() function.
func (_ GlSupport) PatchParameteri() bool {
	return C.cancallPatchParameteri() == 1
}

//	If Supports.PatchParameteri() is true, calls PatchParameteri() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PatchParameteri(pname Enum, value Int) (err error) {
	if !Supports.PatchParameteri() {
		err = fmt.Errorf("PatchParameteri() function call not supported by current GL context.")
	} else {
		PatchParameteri(pname, value)
		err = Util.Error("PatchParameteri()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramPipelineInfoLog() function.
func (_ GlSupport) GetProgramPipelineInfoLog() bool {
	return C.cancallGetProgramPipelineInfoLog() == 1
}

//	If Supports.GetProgramPipelineInfoLog() is true, calls GetProgramPipelineInfoLog() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramPipelineInfoLog(pipeline Uint, bufSize Sizei, length *Sizei, infoLog *Char) (err error) {
	if !Supports.GetProgramPipelineInfoLog() {
		err = fmt.Errorf("GetProgramPipelineInfoLog() function call not supported by current GL context.")
	} else {
		GetProgramPipelineInfoLog(pipeline, bufSize, length, infoLog)
		err = Util.Error("GetProgramPipelineInfoLog()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetActiveSubroutineName() function.
func (_ GlSupport) GetActiveSubroutineName() bool {
	return C.cancallGetActiveSubroutineName() == 1
}

//	If Supports.GetActiveSubroutineName() is true, calls GetActiveSubroutineName() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetActiveSubroutineName(program Uint, shadertype Enum, index Uint, bufsize Sizei, length *Sizei, name *Char) (err error) {
	if !Supports.GetActiveSubroutineName() {
		err = fmt.Errorf("GetActiveSubroutineName() function call not supported by current GL context.")
	} else {
		GetActiveSubroutineName(program, shadertype, index, bufsize, length, name)
		err = Util.Error("GetActiveSubroutineName()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetTexParameterIuiv() function.
func (_ GlSupport) GetTexParameterIuiv() bool {
	return C.cancallGetTexParameterIuiv() == 1
}

//	If Supports.GetTexParameterIuiv() is true, calls GetTexParameterIuiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetTexParameterIuiv(target Enum, pname Enum, params *Uint) (err error) {
	if !Supports.GetTexParameterIuiv() {
		err = fmt.Errorf("GetTexParameterIuiv() function call not supported by current GL context.")
	} else {
		GetTexParameterIuiv(target, pname, params)
		err = Util.Error("GetTexParameterIuiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform1fv() function.
func (_ GlSupport) ProgramUniform1fv() bool {
	return C.cancallProgramUniform1fv() == 1
}

//	If Supports.ProgramUniform1fv() is true, calls ProgramUniform1fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform1fv(program Uint, location Int, count Sizei, value *Float) (err error) {
	if !Supports.ProgramUniform1fv() {
		err = fmt.Errorf("ProgramUniform1fv() function call not supported by current GL context.")
	} else {
		ProgramUniform1fv(program, location, count, value)
		err = Util.Error("ProgramUniform1fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexCoordP3uiv() function.
func (_ GlSupport) TexCoordP3uiv() bool {
	return C.cancallTexCoordP3uiv() == 1
}

//	If Supports.TexCoordP3uiv() is true, calls TexCoordP3uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexCoordP3uiv(type_ Enum, coords *Uint) (err error) {
	if !Supports.TexCoordP3uiv() {
		err = fmt.Errorf("TexCoordP3uiv() function call not supported by current GL context.")
	} else {
		TexCoordP3uiv(type_, coords)
		err = Util.Error("TexCoordP3uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GenProgramPipelines() function.
func (_ GlSupport) GenProgramPipelines() bool {
	return C.cancallGenProgramPipelines() == 1
}

//	If Supports.GenProgramPipelines() is true, calls GenProgramPipelines() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GenProgramPipelines(n Sizei, pipelines *Uint) (err error) {
	if !Supports.GenProgramPipelines() {
		err = fmt.Errorf("GenProgramPipelines() function call not supported by current GL context.")
	} else {
		GenProgramPipelines(n, pipelines)
		err = Util.Error("GenProgramPipelines()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetDebugMessageLog() function.
func (_ GlSupport) GetDebugMessageLog() bool {
	return C.cancallGetDebugMessageLog() == 1
}

//	If Supports.GetDebugMessageLog() is true, calls GetDebugMessageLog() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetDebugMessageLog(count Uint, bufsize Sizei, sources *Enum, types *Enum, ids *Uint, severities *Enum, lengths *Sizei, messageLog *Char) (err error, tryRetVal__ Uint) {
	if !Supports.GetDebugMessageLog() {
		err = fmt.Errorf("GetDebugMessageLog() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetDebugMessageLog(count, bufsize, sources, types, ids, severities, lengths, messageLog)
		err = Util.Error("GetDebugMessageLog()")
	}
	return
}

//	Returns whether the current GL context supports calling the SamplerParameteriv() function.
func (_ GlSupport) SamplerParameteriv() bool {
	return C.cancallSamplerParameteriv() == 1
}

//	If Supports.SamplerParameteriv() is true, calls SamplerParameteriv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) SamplerParameteriv(sampler Uint, pname Enum, param *Int) (err error) {
	if !Supports.SamplerParameteriv() {
		err = fmt.Errorf("SamplerParameteriv() function call not supported by current GL context.")
	} else {
		SamplerParameteriv(sampler, pname, param)
		err = Util.Error("SamplerParameteriv()")
	}
	return
}

//	Returns whether the current GL context supports calling the FrontFace() function.
func (_ GlSupport) FrontFace() bool {
	return C.cancallFrontFace() == 1
}

//	If Supports.FrontFace() is true, calls FrontFace() and yields the err returned by Util.Error(), if any.
func (_ GlTry) FrontFace(mode Enum) (err error) {
	if !Supports.FrontFace() {
		err = fmt.Errorf("FrontFace() function call not supported by current GL context.")
	} else {
		FrontFace(mode)
		err = Util.Error("FrontFace()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteBuffers() function.
func (_ GlSupport) DeleteBuffers() bool {
	return C.cancallDeleteBuffers() == 1
}

//	If Supports.DeleteBuffers() is true, calls DeleteBuffers() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteBuffers(n Sizei, buffers *Uint) (err error) {
	if !Supports.DeleteBuffers() {
		err = fmt.Errorf("DeleteBuffers() function call not supported by current GL context.")
	} else {
		DeleteBuffers(n, buffers)
		err = Util.Error("DeleteBuffers()")
	}
	return
}

//	Returns whether the current GL context supports calling the CompressedTexSubImage3D() function.
func (_ GlSupport) CompressedTexSubImage3D() bool {
	return C.cancallCompressedTexSubImage3D() == 1
}

//	If Supports.CompressedTexSubImage3D() is true, calls CompressedTexSubImage3D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CompressedTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, imageSize Sizei, data Ptr) (err error) {
	if !Supports.CompressedTexSubImage3D() {
		err = fmt.Errorf("CompressedTexSubImage3D() function call not supported by current GL context.")
	} else {
		CompressedTexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data)
		err = Util.Error("CompressedTexSubImage3D()")
	}
	return
}

//	Returns whether the current GL context supports calling the CompressedTexSubImage2D() function.
func (_ GlSupport) CompressedTexSubImage2D() bool {
	return C.cancallCompressedTexSubImage2D() == 1
}

//	If Supports.CompressedTexSubImage2D() is true, calls CompressedTexSubImage2D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CompressedTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, imageSize Sizei, data Ptr) (err error) {
	if !Supports.CompressedTexSubImage2D() {
		err = fmt.Errorf("CompressedTexSubImage2D() function call not supported by current GL context.")
	} else {
		CompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, imageSize, data)
		err = Util.Error("CompressedTexSubImage2D()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindProgramPipeline() function.
func (_ GlSupport) BindProgramPipeline() bool {
	return C.cancallBindProgramPipeline() == 1
}

//	If Supports.BindProgramPipeline() is true, calls BindProgramPipeline() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindProgramPipeline(pipeline Uint) (err error) {
	if !Supports.BindProgramPipeline() {
		err = fmt.Errorf("BindProgramPipeline() function call not supported by current GL context.")
	} else {
		BindProgramPipeline(pipeline)
		err = Util.Error("BindProgramPipeline()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform1iv() function.
func (_ GlSupport) Uniform1iv() bool {
	return C.cancallUniform1iv() == 1
}

//	If Supports.Uniform1iv() is true, calls Uniform1iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform1iv(location Int, count Sizei, value *Int) (err error) {
	if !Supports.Uniform1iv() {
		err = fmt.Errorf("Uniform1iv() function call not supported by current GL context.")
	} else {
		Uniform1iv(location, count, value)
		err = Util.Error("Uniform1iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribDivisor() function.
func (_ GlSupport) VertexAttribDivisor() bool {
	return C.cancallVertexAttribDivisor() == 1
}

//	If Supports.VertexAttribDivisor() is true, calls VertexAttribDivisor() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribDivisor(index Uint, divisor Uint) (err error) {
	if !Supports.VertexAttribDivisor() {
		err = fmt.Errorf("VertexAttribDivisor() function call not supported by current GL context.")
	} else {
		VertexAttribDivisor(index, divisor)
		err = Util.Error("VertexAttribDivisor()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindBufferBase() function.
func (_ GlSupport) BindBufferBase() bool {
	return C.cancallBindBufferBase() == 1
}

//	If Supports.BindBufferBase() is true, calls BindBufferBase() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindBufferBase(target Enum, index Uint, buffer Uint) (err error) {
	if !Supports.BindBufferBase() {
		err = fmt.Errorf("BindBufferBase() function call not supported by current GL context.")
	} else {
		BindBufferBase(target, index, buffer)
		err = Util.Error("BindBufferBase()")
	}
	return
}

//	Returns whether the current GL context supports calling the ColorP4ui() function.
func (_ GlSupport) ColorP4ui() bool {
	return C.cancallColorP4ui() == 1
}

//	If Supports.ColorP4ui() is true, calls ColorP4ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ColorP4ui(type_ Enum, color Uint) (err error) {
	if !Supports.ColorP4ui() {
		err = fmt.Errorf("ColorP4ui() function call not supported by current GL context.")
	} else {
		ColorP4ui(type_, color)
		err = Util.Error("ColorP4ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramResourceIndex() function.
func (_ GlSupport) GetProgramResourceIndex() bool {
	return C.cancallGetProgramResourceIndex() == 1
}

//	If Supports.GetProgramResourceIndex() is true, calls GetProgramResourceIndex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramResourceIndex(program Uint, programInterface Enum, name *Char) (err error, tryRetVal__ Uint) {
	if !Supports.GetProgramResourceIndex() {
		err = fmt.Errorf("GetProgramResourceIndex() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetProgramResourceIndex(program, programInterface, name)
		err = Util.Error("GetProgramResourceIndex()")
	}
	return
}

//	Returns whether the current GL context supports calling the FramebufferParameteri() function.
func (_ GlSupport) FramebufferParameteri() bool {
	return C.cancallFramebufferParameteri() == 1
}

//	If Supports.FramebufferParameteri() is true, calls FramebufferParameteri() and yields the err returned by Util.Error(), if any.
func (_ GlTry) FramebufferParameteri(target Enum, pname Enum, param Int) (err error) {
	if !Supports.FramebufferParameteri() {
		err = fmt.Errorf("FramebufferParameteri() function call not supported by current GL context.")
	} else {
		FramebufferParameteri(target, pname, param)
		err = Util.Error("FramebufferParameteri()")
	}
	return
}

//	Returns whether the current GL context supports calling the ValidateProgram() function.
func (_ GlSupport) ValidateProgram() bool {
	return C.cancallValidateProgram() == 1
}

//	If Supports.ValidateProgram() is true, calls ValidateProgram() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ValidateProgram(program Uint) (err error) {
	if !Supports.ValidateProgram() {
		err = fmt.Errorf("ValidateProgram() function call not supported by current GL context.")
	} else {
		ValidateProgram(program)
		err = Util.Error("ValidateProgram()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawRangeElements() function.
func (_ GlSupport) DrawRangeElements() bool {
	return C.cancallDrawRangeElements() == 1
}

//	If Supports.DrawRangeElements() is true, calls DrawRangeElements() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawRangeElements(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Ptr) (err error) {
	if !Supports.DrawRangeElements() {
		err = fmt.Errorf("DrawRangeElements() function call not supported by current GL context.")
	} else {
		DrawRangeElements(mode, start, end, count, type_, indices)
		err = Util.Error("DrawRangeElements()")
	}
	return
}

//	Returns whether the current GL context supports calling the SamplerParameteri() function.
func (_ GlSupport) SamplerParameteri() bool {
	return C.cancallSamplerParameteri() == 1
}

//	If Supports.SamplerParameteri() is true, calls SamplerParameteri() and yields the err returned by Util.Error(), if any.
func (_ GlTry) SamplerParameteri(sampler Uint, pname Enum, param Int) (err error) {
	if !Supports.SamplerParameteri() {
		err = fmt.Errorf("SamplerParameteri() function call not supported by current GL context.")
	} else {
		SamplerParameteri(sampler, pname, param)
		err = Util.Error("SamplerParameteri()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetActiveAttrib() function.
func (_ GlSupport) GetActiveAttrib() bool {
	return C.cancallGetActiveAttrib() == 1
}

//	If Supports.GetActiveAttrib() is true, calls GetActiveAttrib() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetActiveAttrib(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char) (err error) {
	if !Supports.GetActiveAttrib() {
		err = fmt.Errorf("GetActiveAttrib() function call not supported by current GL context.")
	} else {
		GetActiveAttrib(program, index, bufSize, length, size, type_, name)
		err = Util.Error("GetActiveAttrib()")
	}
	return
}

//	Returns whether the current GL context supports calling the SecondaryColorP3uiv() function.
func (_ GlSupport) SecondaryColorP3uiv() bool {
	return C.cancallSecondaryColorP3uiv() == 1
}

//	If Supports.SecondaryColorP3uiv() is true, calls SecondaryColorP3uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) SecondaryColorP3uiv(type_ Enum, color *Uint) (err error) {
	if !Supports.SecondaryColorP3uiv() {
		err = fmt.Errorf("SecondaryColorP3uiv() function call not supported by current GL context.")
	} else {
		SecondaryColorP3uiv(type_, color)
		err = Util.Error("SecondaryColorP3uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the SamplerParameterfv() function.
func (_ GlSupport) SamplerParameterfv() bool {
	return C.cancallSamplerParameterfv() == 1
}

//	If Supports.SamplerParameterfv() is true, calls SamplerParameterfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) SamplerParameterfv(sampler Uint, pname Enum, param *Float) (err error) {
	if !Supports.SamplerParameterfv() {
		err = fmt.Errorf("SamplerParameterfv() function call not supported by current GL context.")
	} else {
		SamplerParameterfv(sampler, pname, param)
		err = Util.Error("SamplerParameterfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DepthRange() function.
func (_ GlSupport) DepthRange() bool {
	return C.cancallDepthRange() == 1
}

//	If Supports.DepthRange() is true, calls DepthRange() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DepthRange(near_ Double, far_ Double) (err error) {
	if !Supports.DepthRange() {
		err = fmt.Errorf("DepthRange() function call not supported by current GL context.")
	} else {
		DepthRange(near_, far_)
		err = Util.Error("DepthRange()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform3uiv() function.
func (_ GlSupport) Uniform3uiv() bool {
	return C.cancallUniform3uiv() == 1
}

//	If Supports.Uniform3uiv() is true, calls Uniform3uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform3uiv(location Int, count Sizei, value *Uint) (err error) {
	if !Supports.Uniform3uiv() {
		err = fmt.Errorf("Uniform3uiv() function call not supported by current GL context.")
	} else {
		Uniform3uiv(location, count, value)
		err = Util.Error("Uniform3uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the SamplerParameterf() function.
func (_ GlSupport) SamplerParameterf() bool {
	return C.cancallSamplerParameterf() == 1
}

//	If Supports.SamplerParameterf() is true, calls SamplerParameterf() and yields the err returned by Util.Error(), if any.
func (_ GlTry) SamplerParameterf(sampler Uint, pname Enum, param Float) (err error) {
	if !Supports.SamplerParameterf() {
		err = fmt.Errorf("SamplerParameterf() function call not supported by current GL context.")
	} else {
		SamplerParameterf(sampler, pname, param)
		err = Util.Error("SamplerParameterf()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform2f() function.
func (_ GlSupport) Uniform2f() bool {
	return C.cancallUniform2f() == 1
}

//	If Supports.Uniform2f() is true, calls Uniform2f() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform2f(location Int, v0 Float, v1 Float) (err error) {
	if !Supports.Uniform2f() {
		err = fmt.Errorf("Uniform2f() function call not supported by current GL context.")
	} else {
		Uniform2f(location, v0, v1)
		err = Util.Error("Uniform2f()")
	}
	return
}

//	Returns whether the current GL context supports calling the SamplerParameterIuiv() function.
func (_ GlSupport) SamplerParameterIuiv() bool {
	return C.cancallSamplerParameterIuiv() == 1
}

//	If Supports.SamplerParameterIuiv() is true, calls SamplerParameterIuiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) SamplerParameterIuiv(sampler Uint, pname Enum, param *Uint) (err error) {
	if !Supports.SamplerParameterIuiv() {
		err = fmt.Errorf("SamplerParameterIuiv() function call not supported by current GL context.")
	} else {
		SamplerParameterIuiv(sampler, pname, param)
		err = Util.Error("SamplerParameterIuiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawArraysIndirect() function.
func (_ GlSupport) DrawArraysIndirect() bool {
	return C.cancallDrawArraysIndirect() == 1
}

//	If Supports.DrawArraysIndirect() is true, calls DrawArraysIndirect() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawArraysIndirect(mode Enum, indirect Ptr) (err error) {
	if !Supports.DrawArraysIndirect() {
		err = fmt.Errorf("DrawArraysIndirect() function call not supported by current GL context.")
	} else {
		DrawArraysIndirect(mode, indirect)
		err = Util.Error("DrawArraysIndirect()")
	}
	return
}

//	Returns whether the current GL context supports calling the GenTextures() function.
func (_ GlSupport) GenTextures() bool {
	return C.cancallGenTextures() == 1
}

//	If Supports.GenTextures() is true, calls GenTextures() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GenTextures(n Sizei, textures *Uint) (err error) {
	if !Supports.GenTextures() {
		err = fmt.Errorf("GenTextures() function call not supported by current GL context.")
	} else {
		GenTextures(n, textures)
		err = Util.Error("GenTextures()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindBufferRange() function.
func (_ GlSupport) BindBufferRange() bool {
	return C.cancallBindBufferRange() == 1
}

//	If Supports.BindBufferRange() is true, calls BindBufferRange() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindBufferRange(target Enum, index Uint, buffer Uint, offset Intptr, size Sizeiptr) (err error) {
	if !Supports.BindBufferRange() {
		err = fmt.Errorf("BindBufferRange() function call not supported by current GL context.")
	} else {
		BindBufferRange(target, index, buffer, offset, size)
		err = Util.Error("BindBufferRange()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix4x2dv() function.
func (_ GlSupport) UniformMatrix4x2dv() bool {
	return C.cancallUniformMatrix4x2dv() == 1
}

//	If Supports.UniformMatrix4x2dv() is true, calls UniformMatrix4x2dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix4x2dv(location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.UniformMatrix4x2dv() {
		err = fmt.Errorf("UniformMatrix4x2dv() function call not supported by current GL context.")
	} else {
		UniformMatrix4x2dv(location, count, transpose, value)
		err = Util.Error("UniformMatrix4x2dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Viewport() function.
func (_ GlSupport) Viewport() bool {
	return C.cancallViewport() == 1
}

//	If Supports.Viewport() is true, calls Viewport() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Viewport(x Int, y Int, width Sizei, height Sizei) (err error) {
	if !Supports.Viewport() {
		err = fmt.Errorf("Viewport() function call not supported by current GL context.")
	} else {
		Viewport(x, y, width, height)
		err = Util.Error("Viewport()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiTexCoordP1uiv() function.
func (_ GlSupport) MultiTexCoordP1uiv() bool {
	return C.cancallMultiTexCoordP1uiv() == 1
}

//	If Supports.MultiTexCoordP1uiv() is true, calls MultiTexCoordP1uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiTexCoordP1uiv(texture Enum, type_ Enum, coords *Uint) (err error) {
	if !Supports.MultiTexCoordP1uiv() {
		err = fmt.Errorf("MultiTexCoordP1uiv() function call not supported by current GL context.")
	} else {
		MultiTexCoordP1uiv(texture, type_, coords)
		err = Util.Error("MultiTexCoordP1uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix2fv() function.
func (_ GlSupport) UniformMatrix2fv() bool {
	return C.cancallUniformMatrix2fv() == 1
}

//	If Supports.UniformMatrix2fv() is true, calls UniformMatrix2fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix2fv(location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.UniformMatrix2fv() {
		err = fmt.Errorf("UniformMatrix2fv() function call not supported by current GL context.")
	} else {
		UniformMatrix2fv(location, count, transpose, value)
		err = Util.Error("UniformMatrix2fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform2f() function.
func (_ GlSupport) ProgramUniform2f() bool {
	return C.cancallProgramUniform2f() == 1
}

//	If Supports.ProgramUniform2f() is true, calls ProgramUniform2f() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform2f(program Uint, location Int, v0 Float, v1 Float) (err error) {
	if !Supports.ProgramUniform2f() {
		err = fmt.Errorf("ProgramUniform2f() function call not supported by current GL context.")
	} else {
		ProgramUniform2f(program, location, v0, v1)
		err = Util.Error("ProgramUniform2f()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetUniformLocation() function.
func (_ GlSupport) GetUniformLocation() bool {
	return C.cancallGetUniformLocation() == 1
}

//	If Supports.GetUniformLocation() is true, calls GetUniformLocation() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetUniformLocation(program Uint, name *Char) (err error, tryRetVal__ Int) {
	if !Supports.GetUniformLocation() {
		err = fmt.Errorf("GetUniformLocation() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetUniformLocation(program, name)
		err = Util.Error("GetUniformLocation()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexStorage2DMultisample() function.
func (_ GlSupport) TexStorage2DMultisample() bool {
	return C.cancallTexStorage2DMultisample() == 1
}

//	If Supports.TexStorage2DMultisample() is true, calls TexStorage2DMultisample() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexStorage2DMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei, fixedsamplelocations Boolean) (err error) {
	if !Supports.TexStorage2DMultisample() {
		err = fmt.Errorf("TexStorage2DMultisample() function call not supported by current GL context.")
	} else {
		TexStorage2DMultisample(target, samples, internalformat, width, height, fixedsamplelocations)
		err = Util.Error("TexStorage2DMultisample()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetFloatv() function.
func (_ GlSupport) GetFloatv() bool {
	return C.cancallGetFloatv() == 1
}

//	If Supports.GetFloatv() is true, calls GetFloatv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetFloatv(pname Enum, params *Float) (err error) {
	if !Supports.GetFloatv() {
		err = fmt.Errorf("GetFloatv() function call not supported by current GL context.")
	} else {
		GetFloatv(pname, params)
		err = Util.Error("GetFloatv()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix3x2fv() function.
func (_ GlSupport) UniformMatrix3x2fv() bool {
	return C.cancallUniformMatrix3x2fv() == 1
}

//	If Supports.UniformMatrix3x2fv() is true, calls UniformMatrix3x2fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix3x2fv(location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.UniformMatrix3x2fv() {
		err = fmt.Errorf("UniformMatrix3x2fv() function call not supported by current GL context.")
	} else {
		UniformMatrix3x2fv(location, count, transpose, value)
		err = Util.Error("UniformMatrix3x2fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindSampler() function.
func (_ GlSupport) BindSampler() bool {
	return C.cancallBindSampler() == 1
}

//	If Supports.BindSampler() is true, calls BindSampler() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindSampler(unit Uint, sampler Uint) (err error) {
	if !Supports.BindSampler() {
		err = fmt.Errorf("BindSampler() function call not supported by current GL context.")
	} else {
		BindSampler(unit, sampler)
		err = Util.Error("BindSampler()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform4iv() function.
func (_ GlSupport) Uniform4iv() bool {
	return C.cancallUniform4iv() == 1
}

//	If Supports.Uniform4iv() is true, calls Uniform4iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform4iv(location Int, count Sizei, value *Int) (err error) {
	if !Supports.Uniform4iv() {
		err = fmt.Errorf("Uniform4iv() function call not supported by current GL context.")
	} else {
		Uniform4iv(location, count, value)
		err = Util.Error("Uniform4iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the MemoryBarrier() function.
func (_ GlSupport) MemoryBarrier() bool {
	return C.cancallMemoryBarrier() == 1
}

//	If Supports.MemoryBarrier() is true, calls MemoryBarrier() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MemoryBarrier(barriers Bitfield) (err error) {
	if !Supports.MemoryBarrier() {
		err = fmt.Errorf("MemoryBarrier() function call not supported by current GL context.")
	} else {
		MemoryBarrier(barriers)
		err = Util.Error("MemoryBarrier()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribPointer() function.
func (_ GlSupport) VertexAttribPointer() bool {
	return C.cancallVertexAttribPointer() == 1
}

//	If Supports.VertexAttribPointer() is true, calls VertexAttribPointer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribPointer(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, pointer Ptr) (err error) {
	if !Supports.VertexAttribPointer() {
		err = fmt.Errorf("VertexAttribPointer() function call not supported by current GL context.")
	} else {
		VertexAttribPointer(index, size, type_, normalized, stride, pointer)
		err = Util.Error("VertexAttribPointer()")
	}
	return
}

//	Returns whether the current GL context supports calling the BufferData() function.
func (_ GlSupport) BufferData() bool {
	return C.cancallBufferData() == 1
}

//	If Supports.BufferData() is true, calls BufferData() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BufferData(target Enum, size Sizeiptr, data Ptr, usage Enum) (err error) {
	if !Supports.BufferData() {
		err = fmt.Errorf("BufferData() function call not supported by current GL context.")
	} else {
		BufferData(target, size, data, usage)
		err = Util.Error("BufferData()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform2dv() function.
func (_ GlSupport) ProgramUniform2dv() bool {
	return C.cancallProgramUniform2dv() == 1
}

//	If Supports.ProgramUniform2dv() is true, calls ProgramUniform2dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform2dv(program Uint, location Int, count Sizei, value *Double) (err error) {
	if !Supports.ProgramUniform2dv() {
		err = fmt.Errorf("ProgramUniform2dv() function call not supported by current GL context.")
	} else {
		ProgramUniform2dv(program, location, count, value)
		err = Util.Error("ProgramUniform2dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the LineWidth() function.
func (_ GlSupport) LineWidth() bool {
	return C.cancallLineWidth() == 1
}

//	If Supports.LineWidth() is true, calls LineWidth() and yields the err returned by Util.Error(), if any.
func (_ GlTry) LineWidth(width Float) (err error) {
	if !Supports.LineWidth() {
		err = fmt.Errorf("LineWidth() function call not supported by current GL context.")
	} else {
		LineWidth(width)
		err = Util.Error("LineWidth()")
	}
	return
}

//	Returns whether the current GL context supports calling the ColorP4uiv() function.
func (_ GlSupport) ColorP4uiv() bool {
	return C.cancallColorP4uiv() == 1
}

//	If Supports.ColorP4uiv() is true, calls ColorP4uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ColorP4uiv(type_ Enum, color *Uint) (err error) {
	if !Supports.ColorP4uiv() {
		err = fmt.Errorf("ColorP4uiv() function call not supported by current GL context.")
	} else {
		ColorP4uiv(type_, color)
		err = Util.Error("ColorP4uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform3i() function.
func (_ GlSupport) ProgramUniform3i() bool {
	return C.cancallProgramUniform3i() == 1
}

//	If Supports.ProgramUniform3i() is true, calls ProgramUniform3i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform3i(program Uint, location Int, v0 Int, v1 Int, v2 Int) (err error) {
	if !Supports.ProgramUniform3i() {
		err = fmt.Errorf("ProgramUniform3i() function call not supported by current GL context.")
	} else {
		ProgramUniform3i(program, location, v0, v1, v2)
		err = Util.Error("ProgramUniform3i()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawElements() function.
func (_ GlSupport) DrawElements() bool {
	return C.cancallDrawElements() == 1
}

//	If Supports.DrawElements() is true, calls DrawElements() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawElements(mode Enum, count Sizei, type_ Enum, indices Ptr) (err error) {
	if !Supports.DrawElements() {
		err = fmt.Errorf("DrawElements() function call not supported by current GL context.")
	} else {
		DrawElements(mode, count, type_, indices)
		err = Util.Error("DrawElements()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramResourceLocation() function.
func (_ GlSupport) GetProgramResourceLocation() bool {
	return C.cancallGetProgramResourceLocation() == 1
}

//	If Supports.GetProgramResourceLocation() is true, calls GetProgramResourceLocation() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramResourceLocation(program Uint, programInterface Enum, name *Char) (err error, tryRetVal__ Int) {
	if !Supports.GetProgramResourceLocation() {
		err = fmt.Errorf("GetProgramResourceLocation() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetProgramResourceLocation(program, programInterface, name)
		err = Util.Error("GetProgramResourceLocation()")
	}
	return
}

//	Returns whether the current GL context supports calling the InvalidateBufferSubData() function.
func (_ GlSupport) InvalidateBufferSubData() bool {
	return C.cancallInvalidateBufferSubData() == 1
}

//	If Supports.InvalidateBufferSubData() is true, calls InvalidateBufferSubData() and yields the err returned by Util.Error(), if any.
func (_ GlTry) InvalidateBufferSubData(buffer Uint, offset Intptr, length Sizeiptr) (err error) {
	if !Supports.InvalidateBufferSubData() {
		err = fmt.Errorf("InvalidateBufferSubData() function call not supported by current GL context.")
	} else {
		InvalidateBufferSubData(buffer, offset, length)
		err = Util.Error("InvalidateBufferSubData()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexStorage2D() function.
func (_ GlSupport) TexStorage2D() bool {
	return C.cancallTexStorage2D() == 1
}

//	If Supports.TexStorage2D() is true, calls TexStorage2D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexStorage2D(target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei) (err error) {
	if !Supports.TexStorage2D() {
		err = fmt.Errorf("TexStorage2D() function call not supported by current GL context.")
	} else {
		TexStorage2D(target, levels, internalformat, width, height)
		err = Util.Error("TexStorage2D()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteQueries() function.
func (_ GlSupport) DeleteQueries() bool {
	return C.cancallDeleteQueries() == 1
}

//	If Supports.DeleteQueries() is true, calls DeleteQueries() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteQueries(n Sizei, ids *Uint) (err error) {
	if !Supports.DeleteQueries() {
		err = fmt.Errorf("DeleteQueries() function call not supported by current GL context.")
	} else {
		DeleteQueries(n, ids)
		err = Util.Error("DeleteQueries()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetTexImage() function.
func (_ GlSupport) GetTexImage() bool {
	return C.cancallGetTexImage() == 1
}

//	If Supports.GetTexImage() is true, calls GetTexImage() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetTexImage(target Enum, level Int, format Enum, type_ Enum, pixels Ptr) (err error) {
	if !Supports.GetTexImage() {
		err = fmt.Errorf("GetTexImage() function call not supported by current GL context.")
	} else {
		GetTexImage(target, level, format, type_, pixels)
		err = Util.Error("GetTexImage()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClientWaitSync() function.
func (_ GlSupport) ClientWaitSync() bool {
	return C.cancallClientWaitSync() == 1
}

//	If Supports.ClientWaitSync() is true, calls ClientWaitSync() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClientWaitSync(sync Sync, flags Bitfield, timeout Uint64) (err error, tryRetVal__ Enum) {
	if !Supports.ClientWaitSync() {
		err = fmt.Errorf("ClientWaitSync() function call not supported by current GL context.")
	} else {
		tryRetVal__ = ClientWaitSync(sync, flags, timeout)
		err = Util.Error("ClientWaitSync()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawElementsInstancedBaseVertex() function.
func (_ GlSupport) DrawElementsInstancedBaseVertex() bool {
	return C.cancallDrawElementsInstancedBaseVertex() == 1
}

//	If Supports.DrawElementsInstancedBaseVertex() is true, calls DrawElementsInstancedBaseVertex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawElementsInstancedBaseVertex(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, basevertex Int) (err error) {
	if !Supports.DrawElementsInstancedBaseVertex() {
		err = fmt.Errorf("DrawElementsInstancedBaseVertex() function call not supported by current GL context.")
	} else {
		DrawElementsInstancedBaseVertex(mode, count, type_, indices, instancecount, basevertex)
		err = Util.Error("DrawElementsInstancedBaseVertex()")
	}
	return
}

//	Returns whether the current GL context supports calling the PixelStoref() function.
func (_ GlSupport) PixelStoref() bool {
	return C.cancallPixelStoref() == 1
}

//	If Supports.PixelStoref() is true, calls PixelStoref() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PixelStoref(pname Enum, param Float) (err error) {
	if !Supports.PixelStoref() {
		err = fmt.Errorf("PixelStoref() function call not supported by current GL context.")
	} else {
		PixelStoref(pname, param)
		err = Util.Error("PixelStoref()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexParameteriv() function.
func (_ GlSupport) TexParameteriv() bool {
	return C.cancallTexParameteriv() == 1
}

//	If Supports.TexParameteriv() is true, calls TexParameteriv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexParameteriv(target Enum, pname Enum, params *Int) (err error) {
	if !Supports.TexParameteriv() {
		err = fmt.Errorf("TexParameteriv() function call not supported by current GL context.")
	} else {
		TexParameteriv(target, pname, params)
		err = Util.Error("TexParameteriv()")
	}
	return
}

//	Returns whether the current GL context supports calling the FramebufferRenderbuffer() function.
func (_ GlSupport) FramebufferRenderbuffer() bool {
	return C.cancallFramebufferRenderbuffer() == 1
}

//	If Supports.FramebufferRenderbuffer() is true, calls FramebufferRenderbuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) FramebufferRenderbuffer(target Enum, attachment Enum, renderbuffertarget Enum, renderbuffer Uint) (err error) {
	if !Supports.FramebufferRenderbuffer() {
		err = fmt.Errorf("FramebufferRenderbuffer() function call not supported by current GL context.")
	} else {
		FramebufferRenderbuffer(target, attachment, renderbuffertarget, renderbuffer)
		err = Util.Error("FramebufferRenderbuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClearBufferSubData() function.
func (_ GlSupport) ClearBufferSubData() bool {
	return C.cancallClearBufferSubData() == 1
}

//	If Supports.ClearBufferSubData() is true, calls ClearBufferSubData() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClearBufferSubData(target Enum, internalformat Enum, offset Intptr, size Sizeiptr, format Enum, type_ Enum, data Ptr) (err error) {
	if !Supports.ClearBufferSubData() {
		err = fmt.Errorf("ClearBufferSubData() function call not supported by current GL context.")
	} else {
		ClearBufferSubData(target, internalformat, offset, size, format, type_, data)
		err = Util.Error("ClearBufferSubData()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform2d() function.
func (_ GlSupport) ProgramUniform2d() bool {
	return C.cancallProgramUniform2d() == 1
}

//	If Supports.ProgramUniform2d() is true, calls ProgramUniform2d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform2d(program Uint, location Int, v0 Double, v1 Double) (err error) {
	if !Supports.ProgramUniform2d() {
		err = fmt.Errorf("ProgramUniform2d() function call not supported by current GL context.")
	} else {
		ProgramUniform2d(program, location, v0, v1)
		err = Util.Error("ProgramUniform2d()")
	}
	return
}

//	Returns whether the current GL context supports calling the SampleCoverage() function.
func (_ GlSupport) SampleCoverage() bool {
	return C.cancallSampleCoverage() == 1
}

//	If Supports.SampleCoverage() is true, calls SampleCoverage() and yields the err returned by Util.Error(), if any.
func (_ GlTry) SampleCoverage(value Float, invert Boolean) (err error) {
	if !Supports.SampleCoverage() {
		err = fmt.Errorf("SampleCoverage() function call not supported by current GL context.")
	} else {
		SampleCoverage(value, invert)
		err = Util.Error("SampleCoverage()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsVertexArray() function.
func (_ GlSupport) IsVertexArray() bool {
	return C.cancallIsVertexArray() == 1
}

//	If Supports.IsVertexArray() is true, calls IsVertexArray() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsVertexArray(array Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsVertexArray() {
		err = fmt.Errorf("IsVertexArray() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsVertexArray(array)
		err = Util.Error("IsVertexArray()")
	}
	return
}

//	Returns whether the current GL context supports calling the DispatchCompute() function.
func (_ GlSupport) DispatchCompute() bool {
	return C.cancallDispatchCompute() == 1
}

//	If Supports.DispatchCompute() is true, calls DispatchCompute() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DispatchCompute(num_groups_x Uint, num_groups_y Uint, num_groups_z Uint) (err error) {
	if !Supports.DispatchCompute() {
		err = fmt.Errorf("DispatchCompute() function call not supported by current GL context.")
	} else {
		DispatchCompute(num_groups_x, num_groups_y, num_groups_z)
		err = Util.Error("DispatchCompute()")
	}
	return
}

//	Returns whether the current GL context supports calling the FenceSync() function.
func (_ GlSupport) FenceSync() bool {
	return C.cancallFenceSync() == 1
}

//	If Supports.FenceSync() is true, calls FenceSync() and yields the err returned by Util.Error(), if any.
func (_ GlTry) FenceSync(condition Enum, flags Bitfield) (err error, tryRetVal__ Sync) {
	if !Supports.FenceSync() {
		err = fmt.Errorf("FenceSync() function call not supported by current GL context.")
	} else {
		tryRetVal__ = FenceSync(condition, flags)
		err = Util.Error("FenceSync()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetVertexAttribiv() function.
func (_ GlSupport) GetVertexAttribiv() bool {
	return C.cancallGetVertexAttribiv() == 1
}

//	If Supports.GetVertexAttribiv() is true, calls GetVertexAttribiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetVertexAttribiv(index Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetVertexAttribiv() {
		err = fmt.Errorf("GetVertexAttribiv() function call not supported by current GL context.")
	} else {
		GetVertexAttribiv(index, pname, params)
		err = Util.Error("GetVertexAttribiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the FlushMappedBufferRange() function.
func (_ GlSupport) FlushMappedBufferRange() bool {
	return C.cancallFlushMappedBufferRange() == 1
}

//	If Supports.FlushMappedBufferRange() is true, calls FlushMappedBufferRange() and yields the err returned by Util.Error(), if any.
func (_ GlTry) FlushMappedBufferRange(target Enum, offset Intptr, length Sizeiptr) (err error) {
	if !Supports.FlushMappedBufferRange() {
		err = fmt.Errorf("FlushMappedBufferRange() function call not supported by current GL context.")
	} else {
		FlushMappedBufferRange(target, offset, length)
		err = Util.Error("FlushMappedBufferRange()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetVertexAttribfv() function.
func (_ GlSupport) GetVertexAttribfv() bool {
	return C.cancallGetVertexAttribfv() == 1
}

//	If Supports.GetVertexAttribfv() is true, calls GetVertexAttribfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetVertexAttribfv(index Uint, pname Enum, params *Float) (err error) {
	if !Supports.GetVertexAttribfv() {
		err = fmt.Errorf("GetVertexAttribfv() function call not supported by current GL context.")
	} else {
		GetVertexAttribfv(index, pname, params)
		err = Util.Error("GetVertexAttribfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the CreateShader() function.
func (_ GlSupport) CreateShader() bool {
	return C.cancallCreateShader() == 1
}

//	If Supports.CreateShader() is true, calls CreateShader() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CreateShader(type_ Enum) (err error, tryRetVal__ Uint) {
	if !Supports.CreateShader() {
		err = fmt.Errorf("CreateShader() function call not supported by current GL context.")
	} else {
		tryRetVal__ = CreateShader(type_)
		err = Util.Error("CreateShader()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribFormat() function.
func (_ GlSupport) VertexAttribFormat() bool {
	return C.cancallVertexAttribFormat() == 1
}

//	If Supports.VertexAttribFormat() is true, calls VertexAttribFormat() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribFormat(attribindex Uint, size Int, type_ Enum, normalized Boolean, relativeoffset Uint) (err error) {
	if !Supports.VertexAttribFormat() {
		err = fmt.Errorf("VertexAttribFormat() function call not supported by current GL context.")
	} else {
		VertexAttribFormat(attribindex, size, type_, normalized, relativeoffset)
		err = Util.Error("VertexAttribFormat()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiDrawElements() function.
func (_ GlSupport) MultiDrawElements() bool {
	return C.cancallMultiDrawElements() == 1
}

//	If Supports.MultiDrawElements() is true, calls MultiDrawElements() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiDrawElements(mode Enum, count *Sizei, type_ Enum, indices *Ptr, drawcount Sizei) (err error) {
	if !Supports.MultiDrawElements() {
		err = fmt.Errorf("MultiDrawElements() function call not supported by current GL context.")
	} else {
		MultiDrawElements(mode, count, type_, indices, drawcount)
		err = Util.Error("MultiDrawElements()")
	}
	return
}

//	Returns whether the current GL context supports calling the CopyImageSubData() function.
func (_ GlSupport) CopyImageSubData() bool {
	return C.cancallCopyImageSubData() == 1
}

//	If Supports.CopyImageSubData() is true, calls CopyImageSubData() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CopyImageSubData(srcName Uint, srcTarget Enum, srcLevel Int, srcX Int, srcY Int, srcZ Int, dstName Uint, dstTarget Enum, dstLevel Int, dstX Int, dstY Int, dstZ Int, srcWidth Sizei, srcHeight Sizei, srcDepth Sizei) (err error) {
	if !Supports.CopyImageSubData() {
		err = fmt.Errorf("CopyImageSubData() function call not supported by current GL context.")
	} else {
		CopyImageSubData(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth)
		err = Util.Error("CopyImageSubData()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawElementsBaseVertex() function.
func (_ GlSupport) DrawElementsBaseVertex() bool {
	return C.cancallDrawElementsBaseVertex() == 1
}

//	If Supports.DrawElementsBaseVertex() is true, calls DrawElementsBaseVertex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawElementsBaseVertex(mode Enum, count Sizei, type_ Enum, indices Ptr, basevertex Int) (err error) {
	if !Supports.DrawElementsBaseVertex() {
		err = fmt.Errorf("DrawElementsBaseVertex() function call not supported by current GL context.")
	} else {
		DrawElementsBaseVertex(mode, count, type_, indices, basevertex)
		err = Util.Error("DrawElementsBaseVertex()")
	}
	return
}

//	Returns whether the current GL context supports calling the StencilMask() function.
func (_ GlSupport) StencilMask() bool {
	return C.cancallStencilMask() == 1
}

//	If Supports.StencilMask() is true, calls StencilMask() and yields the err returned by Util.Error(), if any.
func (_ GlTry) StencilMask(mask Uint) (err error) {
	if !Supports.StencilMask() {
		err = fmt.Errorf("StencilMask() function call not supported by current GL context.")
	} else {
		StencilMask(mask)
		err = Util.Error("StencilMask()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribL4dv() function.
func (_ GlSupport) VertexAttribL4dv() bool {
	return C.cancallVertexAttribL4dv() == 1
}

//	If Supports.VertexAttribL4dv() is true, calls VertexAttribL4dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribL4dv(index Uint, v *Double) (err error) {
	if !Supports.VertexAttribL4dv() {
		err = fmt.Errorf("VertexAttribL4dv() function call not supported by current GL context.")
	} else {
		VertexAttribL4dv(index, v)
		err = Util.Error("VertexAttribL4dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexParameteri() function.
func (_ GlSupport) TexParameteri() bool {
	return C.cancallTexParameteri() == 1
}

//	If Supports.TexParameteri() is true, calls TexParameteri() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexParameteri(target Enum, pname Enum, param Int) (err error) {
	if !Supports.TexParameteri() {
		err = fmt.Errorf("TexParameteri() function call not supported by current GL context.")
	} else {
		TexParameteri(target, pname, param)
		err = Util.Error("TexParameteri()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexCoordP1ui() function.
func (_ GlSupport) TexCoordP1ui() bool {
	return C.cancallTexCoordP1ui() == 1
}

//	If Supports.TexCoordP1ui() is true, calls TexCoordP1ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexCoordP1ui(type_ Enum, coords Uint) (err error) {
	if !Supports.TexCoordP1ui() {
		err = fmt.Errorf("TexCoordP1ui() function call not supported by current GL context.")
	} else {
		TexCoordP1ui(type_, coords)
		err = Util.Error("TexCoordP1ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribL4d() function.
func (_ GlSupport) VertexAttribL4d() bool {
	return C.cancallVertexAttribL4d() == 1
}

//	If Supports.VertexAttribL4d() is true, calls VertexAttribL4d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribL4d(index Uint, x Double, y Double, z Double, w Double) (err error) {
	if !Supports.VertexAttribL4d() {
		err = fmt.Errorf("VertexAttribL4d() function call not supported by current GL context.")
	} else {
		VertexAttribL4d(index, x, y, z, w)
		err = Util.Error("VertexAttribL4d()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexCoordP2ui() function.
func (_ GlSupport) TexCoordP2ui() bool {
	return C.cancallTexCoordP2ui() == 1
}

//	If Supports.TexCoordP2ui() is true, calls TexCoordP2ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexCoordP2ui(type_ Enum, coords Uint) (err error) {
	if !Supports.TexCoordP2ui() {
		err = fmt.Errorf("TexCoordP2ui() function call not supported by current GL context.")
	} else {
		TexCoordP2ui(type_, coords)
		err = Util.Error("TexCoordP2ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetUniformfv() function.
func (_ GlSupport) GetUniformfv() bool {
	return C.cancallGetUniformfv() == 1
}

//	If Supports.GetUniformfv() is true, calls GetUniformfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetUniformfv(program Uint, location Int, params *Float) (err error) {
	if !Supports.GetUniformfv() {
		err = fmt.Errorf("GetUniformfv() function call not supported by current GL context.")
	} else {
		GetUniformfv(program, location, params)
		err = Util.Error("GetUniformfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI2iv() function.
func (_ GlSupport) VertexAttribI2iv() bool {
	return C.cancallVertexAttribI2iv() == 1
}

//	If Supports.VertexAttribI2iv() is true, calls VertexAttribI2iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI2iv(index Uint, v *Int) (err error) {
	if !Supports.VertexAttribI2iv() {
		err = fmt.Errorf("VertexAttribI2iv() function call not supported by current GL context.")
	} else {
		VertexAttribI2iv(index, v)
		err = Util.Error("VertexAttribI2iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetBooleani_v() function.
func (_ GlSupport) GetBooleani_v() bool {
	return C.cancallGetBooleani_v() == 1
}

//	If Supports.GetBooleani_v() is true, calls GetBooleani_v() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetBooleani_v(target Enum, index Uint, data *Boolean) (err error) {
	if !Supports.GetBooleani_v() {
		err = fmt.Errorf("GetBooleani_v() function call not supported by current GL context.")
	} else {
		GetBooleani_v(target, index, data)
		err = Util.Error("GetBooleani_v()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetError() function.
func (_ GlSupport) GetError() bool {
	return C.cancallGetError() == 1
}

//	If Supports.GetError() is true, calls GetError() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetError() (err error, tryRetVal__ Enum) {
	if !Supports.GetError() {
		err = fmt.Errorf("GetError() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetError()
		err = Util.Error("GetError()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix2x4fv() function.
func (_ GlSupport) UniformMatrix2x4fv() bool {
	return C.cancallUniformMatrix2x4fv() == 1
}

//	If Supports.UniformMatrix2x4fv() is true, calls UniformMatrix2x4fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix2x4fv(location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.UniformMatrix2x4fv() {
		err = fmt.Errorf("UniformMatrix2x4fv() function call not supported by current GL context.")
	} else {
		UniformMatrix2x4fv(location, count, transpose, value)
		err = Util.Error("UniformMatrix2x4fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetObjectLabel() function.
func (_ GlSupport) GetObjectLabel() bool {
	return C.cancallGetObjectLabel() == 1
}

//	If Supports.GetObjectLabel() is true, calls GetObjectLabel() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetObjectLabel(identifier Enum, name Uint, bufSize Sizei, length *Sizei, label *Char) (err error) {
	if !Supports.GetObjectLabel() {
		err = fmt.Errorf("GetObjectLabel() function call not supported by current GL context.")
	} else {
		GetObjectLabel(identifier, name, bufSize, length, label)
		err = Util.Error("GetObjectLabel()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexImage2D() function.
func (_ GlSupport) TexImage2D() bool {
	return C.cancallTexImage2D() == 1
}

//	If Supports.TexImage2D() is true, calls TexImage2D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, type_ Enum, pixels Ptr) (err error) {
	if !Supports.TexImage2D() {
		err = fmt.Errorf("TexImage2D() function call not supported by current GL context.")
	} else {
		TexImage2D(target, level, internalformat, width, height, border, format, type_, pixels)
		err = Util.Error("TexImage2D()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix3fv() function.
func (_ GlSupport) ProgramUniformMatrix3fv() bool {
	return C.cancallProgramUniformMatrix3fv() == 1
}

//	If Supports.ProgramUniformMatrix3fv() is true, calls ProgramUniformMatrix3fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.ProgramUniformMatrix3fv() {
		err = fmt.Errorf("ProgramUniformMatrix3fv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix3fv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix3fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexCoordP3ui() function.
func (_ GlSupport) TexCoordP3ui() bool {
	return C.cancallTexCoordP3ui() == 1
}

//	If Supports.TexCoordP3ui() is true, calls TexCoordP3ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexCoordP3ui(type_ Enum, coords Uint) (err error) {
	if !Supports.TexCoordP3ui() {
		err = fmt.Errorf("TexCoordP3ui() function call not supported by current GL context.")
	} else {
		TexCoordP3ui(type_, coords)
		err = Util.Error("TexCoordP3ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetDoublei_v() function.
func (_ GlSupport) GetDoublei_v() bool {
	return C.cancallGetDoublei_v() == 1
}

//	If Supports.GetDoublei_v() is true, calls GetDoublei_v() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetDoublei_v(target Enum, index Uint, data *Double) (err error) {
	if !Supports.GetDoublei_v() {
		err = fmt.Errorf("GetDoublei_v() function call not supported by current GL context.")
	} else {
		GetDoublei_v(target, index, data)
		err = Util.Error("GetDoublei_v()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexImage3D() function.
func (_ GlSupport) TexImage3D() bool {
	return C.cancallTexImage3D() == 1
}

//	If Supports.TexImage3D() is true, calls TexImage3D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexImage3D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, depth Sizei, border Int, format Enum, type_ Enum, pixels Ptr) (err error) {
	if !Supports.TexImage3D() {
		err = fmt.Errorf("TexImage3D() function call not supported by current GL context.")
	} else {
		TexImage3D(target, level, internalformat, width, height, depth, border, format, type_, pixels)
		err = Util.Error("TexImage3D()")
	}
	return
}

//	Returns whether the current GL context supports calling the GenSamplers() function.
func (_ GlSupport) GenSamplers() bool {
	return C.cancallGenSamplers() == 1
}

//	If Supports.GenSamplers() is true, calls GenSamplers() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GenSamplers(count Sizei, samplers *Uint) (err error) {
	if !Supports.GenSamplers() {
		err = fmt.Errorf("GenSamplers() function call not supported by current GL context.")
	} else {
		GenSamplers(count, samplers)
		err = Util.Error("GenSamplers()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI4uiv() function.
func (_ GlSupport) VertexAttribI4uiv() bool {
	return C.cancallVertexAttribI4uiv() == 1
}

//	If Supports.VertexAttribI4uiv() is true, calls VertexAttribI4uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI4uiv(index Uint, v *Uint) (err error) {
	if !Supports.VertexAttribI4uiv() {
		err = fmt.Errorf("VertexAttribI4uiv() function call not supported by current GL context.")
	} else {
		VertexAttribI4uiv(index, v)
		err = Util.Error("VertexAttribI4uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetUniformdv() function.
func (_ GlSupport) GetUniformdv() bool {
	return C.cancallGetUniformdv() == 1
}

//	If Supports.GetUniformdv() is true, calls GetUniformdv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetUniformdv(program Uint, location Int, params *Double) (err error) {
	if !Supports.GetUniformdv() {
		err = fmt.Errorf("GetUniformdv() function call not supported by current GL context.")
	} else {
		GetUniformdv(program, location, params)
		err = Util.Error("GetUniformdv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform4iv() function.
func (_ GlSupport) ProgramUniform4iv() bool {
	return C.cancallProgramUniform4iv() == 1
}

//	If Supports.ProgramUniform4iv() is true, calls ProgramUniform4iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform4iv(program Uint, location Int, count Sizei, value *Int) (err error) {
	if !Supports.ProgramUniform4iv() {
		err = fmt.Errorf("ProgramUniform4iv() function call not supported by current GL context.")
	} else {
		ProgramUniform4iv(program, location, count, value)
		err = Util.Error("ProgramUniform4iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform1i() function.
func (_ GlSupport) ProgramUniform1i() bool {
	return C.cancallProgramUniform1i() == 1
}

//	If Supports.ProgramUniform1i() is true, calls ProgramUniform1i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform1i(program Uint, location Int, v0 Int) (err error) {
	if !Supports.ProgramUniform1i() {
		err = fmt.Errorf("ProgramUniform1i() function call not supported by current GL context.")
	} else {
		ProgramUniform1i(program, location, v0)
		err = Util.Error("ProgramUniform1i()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiTexCoordP1ui() function.
func (_ GlSupport) MultiTexCoordP1ui() bool {
	return C.cancallMultiTexCoordP1ui() == 1
}

//	If Supports.MultiTexCoordP1ui() is true, calls MultiTexCoordP1ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiTexCoordP1ui(texture Enum, type_ Enum, coords Uint) (err error) {
	if !Supports.MultiTexCoordP1ui() {
		err = fmt.Errorf("MultiTexCoordP1ui() function call not supported by current GL context.")
	} else {
		MultiTexCoordP1ui(texture, type_, coords)
		err = Util.Error("MultiTexCoordP1ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix4dv() function.
func (_ GlSupport) ProgramUniformMatrix4dv() bool {
	return C.cancallProgramUniformMatrix4dv() == 1
}

//	If Supports.ProgramUniformMatrix4dv() is true, calls ProgramUniformMatrix4dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.ProgramUniformMatrix4dv() {
		err = fmt.Errorf("ProgramUniformMatrix4dv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix4dv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix4dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform1fv() function.
func (_ GlSupport) Uniform1fv() bool {
	return C.cancallUniform1fv() == 1
}

//	If Supports.Uniform1fv() is true, calls Uniform1fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform1fv(location Int, count Sizei, value *Float) (err error) {
	if !Supports.Uniform1fv() {
		err = fmt.Errorf("Uniform1fv() function call not supported by current GL context.")
	} else {
		Uniform1fv(location, count, value)
		err = Util.Error("Uniform1fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawElementsInstancedBaseVertexBaseInstance() function.
func (_ GlSupport) DrawElementsInstancedBaseVertexBaseInstance() bool {
	return C.cancallDrawElementsInstancedBaseVertexBaseInstance() == 1
}

//	If Supports.DrawElementsInstancedBaseVertexBaseInstance() is true, calls DrawElementsInstancedBaseVertexBaseInstance() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawElementsInstancedBaseVertexBaseInstance(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, basevertex Int, baseinstance Uint) (err error) {
	if !Supports.DrawElementsInstancedBaseVertexBaseInstance() {
		err = fmt.Errorf("DrawElementsInstancedBaseVertexBaseInstance() function call not supported by current GL context.")
	} else {
		DrawElementsInstancedBaseVertexBaseInstance(mode, count, type_, indices, instancecount, basevertex, baseinstance)
		err = Util.Error("DrawElementsInstancedBaseVertexBaseInstance()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawElementsInstancedBaseInstance() function.
func (_ GlSupport) DrawElementsInstancedBaseInstance() bool {
	return C.cancallDrawElementsInstancedBaseInstance() == 1
}

//	If Supports.DrawElementsInstancedBaseInstance() is true, calls DrawElementsInstancedBaseInstance() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawElementsInstancedBaseInstance(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, baseinstance Uint) (err error) {
	if !Supports.DrawElementsInstancedBaseInstance() {
		err = fmt.Errorf("DrawElementsInstancedBaseInstance() function call not supported by current GL context.")
	} else {
		DrawElementsInstancedBaseInstance(mode, count, type_, indices, instancecount, baseinstance)
		err = Util.Error("DrawElementsInstancedBaseInstance()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteTextures() function.
func (_ GlSupport) DeleteTextures() bool {
	return C.cancallDeleteTextures() == 1
}

//	If Supports.DeleteTextures() is true, calls DeleteTextures() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteTextures(n Sizei, textures *Uint) (err error) {
	if !Supports.DeleteTextures() {
		err = fmt.Errorf("DeleteTextures() function call not supported by current GL context.")
	} else {
		DeleteTextures(n, textures)
		err = Util.Error("DeleteTextures()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramiv() function.
func (_ GlSupport) GetProgramiv() bool {
	return C.cancallGetProgramiv() == 1
}

//	If Supports.GetProgramiv() is true, calls GetProgramiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramiv(program Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetProgramiv() {
		err = fmt.Errorf("GetProgramiv() function call not supported by current GL context.")
	} else {
		GetProgramiv(program, pname, params)
		err = Util.Error("GetProgramiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetRenderbufferParameteriv() function.
func (_ GlSupport) GetRenderbufferParameteriv() bool {
	return C.cancallGetRenderbufferParameteriv() == 1
}

//	If Supports.GetRenderbufferParameteriv() is true, calls GetRenderbufferParameteriv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetRenderbufferParameteriv(target Enum, pname Enum, params *Int) (err error) {
	if !Supports.GetRenderbufferParameteriv() {
		err = fmt.Errorf("GetRenderbufferParameteriv() function call not supported by current GL context.")
	} else {
		GetRenderbufferParameteriv(target, pname, params)
		err = Util.Error("GetRenderbufferParameteriv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Clear() function.
func (_ GlSupport) Clear() bool {
	return C.cancallClear() == 1
}

//	If Supports.Clear() is true, calls Clear() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Clear(mask Bitfield) (err error) {
	if !Supports.Clear() {
		err = fmt.Errorf("Clear() function call not supported by current GL context.")
	} else {
		Clear(mask)
		err = Util.Error("Clear()")
	}
	return
}

//	Returns whether the current GL context supports calling the PolygonMode() function.
func (_ GlSupport) PolygonMode() bool {
	return C.cancallPolygonMode() == 1
}

//	If Supports.PolygonMode() is true, calls PolygonMode() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PolygonMode(face Enum, mode Enum) (err error) {
	if !Supports.PolygonMode() {
		err = fmt.Errorf("PolygonMode() function call not supported by current GL context.")
	} else {
		PolygonMode(face, mode)
		err = Util.Error("PolygonMode()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform4uiv() function.
func (_ GlSupport) ProgramUniform4uiv() bool {
	return C.cancallProgramUniform4uiv() == 1
}

//	If Supports.ProgramUniform4uiv() is true, calls ProgramUniform4uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform4uiv(program Uint, location Int, count Sizei, value *Uint) (err error) {
	if !Supports.ProgramUniform4uiv() {
		err = fmt.Errorf("ProgramUniform4uiv() function call not supported by current GL context.")
	} else {
		ProgramUniform4uiv(program, location, count, value)
		err = Util.Error("ProgramUniform4uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetSubroutineUniformLocation() function.
func (_ GlSupport) GetSubroutineUniformLocation() bool {
	return C.cancallGetSubroutineUniformLocation() == 1
}

//	If Supports.GetSubroutineUniformLocation() is true, calls GetSubroutineUniformLocation() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetSubroutineUniformLocation(program Uint, shadertype Enum, name *Char) (err error, tryRetVal__ Int) {
	if !Supports.GetSubroutineUniformLocation() {
		err = fmt.Errorf("GetSubroutineUniformLocation() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetSubroutineUniformLocation(program, shadertype, name)
		err = Util.Error("GetSubroutineUniformLocation()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawArraysInstancedBaseInstance() function.
func (_ GlSupport) DrawArraysInstancedBaseInstance() bool {
	return C.cancallDrawArraysInstancedBaseInstance() == 1
}

//	If Supports.DrawArraysInstancedBaseInstance() is true, calls DrawArraysInstancedBaseInstance() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawArraysInstancedBaseInstance(mode Enum, first Int, count Sizei, instancecount Sizei, baseinstance Uint) (err error) {
	if !Supports.DrawArraysInstancedBaseInstance() {
		err = fmt.Errorf("DrawArraysInstancedBaseInstance() function call not supported by current GL context.")
	} else {
		DrawArraysInstancedBaseInstance(mode, first, count, instancecount, baseinstance)
		err = Util.Error("DrawArraysInstancedBaseInstance()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetVertexAttribPointerv() function.
func (_ GlSupport) GetVertexAttribPointerv() bool {
	return C.cancallGetVertexAttribPointerv() == 1
}

//	If Supports.GetVertexAttribPointerv() is true, calls GetVertexAttribPointerv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetVertexAttribPointerv(index Uint, pname Enum, pointer *Ptr) (err error) {
	if !Supports.GetVertexAttribPointerv() {
		err = fmt.Errorf("GetVertexAttribPointerv() function call not supported by current GL context.")
	} else {
		GetVertexAttribPointerv(index, pname, pointer)
		err = Util.Error("GetVertexAttribPointerv()")
	}
	return
}

//	Returns whether the current GL context supports calling the EnableVertexAttribArray() function.
func (_ GlSupport) EnableVertexAttribArray() bool {
	return C.cancallEnableVertexAttribArray() == 1
}

//	If Supports.EnableVertexAttribArray() is true, calls EnableVertexAttribArray() and yields the err returned by Util.Error(), if any.
func (_ GlTry) EnableVertexAttribArray(index Uint) (err error) {
	if !Supports.EnableVertexAttribArray() {
		err = fmt.Errorf("EnableVertexAttribArray() function call not supported by current GL context.")
	} else {
		EnableVertexAttribArray(index)
		err = Util.Error("EnableVertexAttribArray()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsQuery() function.
func (_ GlSupport) IsQuery() bool {
	return C.cancallIsQuery() == 1
}

//	If Supports.IsQuery() is true, calls IsQuery() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsQuery(id Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsQuery() {
		err = fmt.Errorf("IsQuery() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsQuery(id)
		err = Util.Error("IsQuery()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix4x2fv() function.
func (_ GlSupport) ProgramUniformMatrix4x2fv() bool {
	return C.cancallProgramUniformMatrix4x2fv() == 1
}

//	If Supports.ProgramUniformMatrix4x2fv() is true, calls ProgramUniformMatrix4x2fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix4x2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.ProgramUniformMatrix4x2fv() {
		err = fmt.Errorf("ProgramUniformMatrix4x2fv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix4x2fv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix4x2fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the StencilFuncSeparate() function.
func (_ GlSupport) StencilFuncSeparate() bool {
	return C.cancallStencilFuncSeparate() == 1
}

//	If Supports.StencilFuncSeparate() is true, calls StencilFuncSeparate() and yields the err returned by Util.Error(), if any.
func (_ GlTry) StencilFuncSeparate(face Enum, func_ Enum, ref Int, mask Uint) (err error) {
	if !Supports.StencilFuncSeparate() {
		err = fmt.Errorf("StencilFuncSeparate() function call not supported by current GL context.")
	} else {
		StencilFuncSeparate(face, func_, ref, mask)
		err = Util.Error("StencilFuncSeparate()")
	}
	return
}

//	Returns whether the current GL context supports calling the PixelStorei() function.
func (_ GlSupport) PixelStorei() bool {
	return C.cancallPixelStorei() == 1
}

//	If Supports.PixelStorei() is true, calls PixelStorei() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PixelStorei(pname Enum, param Int) (err error) {
	if !Supports.PixelStorei() {
		err = fmt.Errorf("PixelStorei() function call not supported by current GL context.")
	} else {
		PixelStorei(pname, param)
		err = Util.Error("PixelStorei()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexParameterIiv() function.
func (_ GlSupport) TexParameterIiv() bool {
	return C.cancallTexParameterIiv() == 1
}

//	If Supports.TexParameterIiv() is true, calls TexParameterIiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexParameterIiv(target Enum, pname Enum, params *Int) (err error) {
	if !Supports.TexParameterIiv() {
		err = fmt.Errorf("TexParameterIiv() function call not supported by current GL context.")
	} else {
		TexParameterIiv(target, pname, params)
		err = Util.Error("TexParameterIiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetActiveAtomicCounterBufferiv() function.
func (_ GlSupport) GetActiveAtomicCounterBufferiv() bool {
	return C.cancallGetActiveAtomicCounterBufferiv() == 1
}

//	If Supports.GetActiveAtomicCounterBufferiv() is true, calls GetActiveAtomicCounterBufferiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetActiveAtomicCounterBufferiv(program Uint, bufferIndex Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetActiveAtomicCounterBufferiv() {
		err = fmt.Errorf("GetActiveAtomicCounterBufferiv() function call not supported by current GL context.")
	} else {
		GetActiveAtomicCounterBufferiv(program, bufferIndex, pname, params)
		err = Util.Error("GetActiveAtomicCounterBufferiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetActiveSubroutineUniformiv() function.
func (_ GlSupport) GetActiveSubroutineUniformiv() bool {
	return C.cancallGetActiveSubroutineUniformiv() == 1
}

//	If Supports.GetActiveSubroutineUniformiv() is true, calls GetActiveSubroutineUniformiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetActiveSubroutineUniformiv(program Uint, shadertype Enum, index Uint, pname Enum, values *Int) (err error) {
	if !Supports.GetActiveSubroutineUniformiv() {
		err = fmt.Errorf("GetActiveSubroutineUniformiv() function call not supported by current GL context.")
	} else {
		GetActiveSubroutineUniformiv(program, shadertype, index, pname, values)
		err = Util.Error("GetActiveSubroutineUniformiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ColorMaski() function.
func (_ GlSupport) ColorMaski() bool {
	return C.cancallColorMaski() == 1
}

//	If Supports.ColorMaski() is true, calls ColorMaski() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ColorMaski(index Uint, r Boolean, g Boolean, b Boolean, a Boolean) (err error) {
	if !Supports.ColorMaski() {
		err = fmt.Errorf("ColorMaski() function call not supported by current GL context.")
	} else {
		ColorMaski(index, r, g, b, a)
		err = Util.Error("ColorMaski()")
	}
	return
}

//	Returns whether the current GL context supports calling the ViewportIndexedfv() function.
func (_ GlSupport) ViewportIndexedfv() bool {
	return C.cancallViewportIndexedfv() == 1
}

//	If Supports.ViewportIndexedfv() is true, calls ViewportIndexedfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ViewportIndexedfv(index Uint, v *Float) (err error) {
	if !Supports.ViewportIndexedfv() {
		err = fmt.Errorf("ViewportIndexedfv() function call not supported by current GL context.")
	} else {
		ViewportIndexedfv(index, v)
		err = Util.Error("ViewportIndexedfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform3ui() function.
func (_ GlSupport) ProgramUniform3ui() bool {
	return C.cancallProgramUniform3ui() == 1
}

//	If Supports.ProgramUniform3ui() is true, calls ProgramUniform3ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform3ui(program Uint, location Int, v0 Uint, v1 Uint, v2 Uint) (err error) {
	if !Supports.ProgramUniform3ui() {
		err = fmt.Errorf("ProgramUniform3ui() function call not supported by current GL context.")
	} else {
		ProgramUniform3ui(program, location, v0, v1, v2)
		err = Util.Error("ProgramUniform3ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiTexCoordP2ui() function.
func (_ GlSupport) MultiTexCoordP2ui() bool {
	return C.cancallMultiTexCoordP2ui() == 1
}

//	If Supports.MultiTexCoordP2ui() is true, calls MultiTexCoordP2ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiTexCoordP2ui(texture Enum, type_ Enum, coords Uint) (err error) {
	if !Supports.MultiTexCoordP2ui() {
		err = fmt.Errorf("MultiTexCoordP2ui() function call not supported by current GL context.")
	} else {
		MultiTexCoordP2ui(texture, type_, coords)
		err = Util.Error("MultiTexCoordP2ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform4f() function.
func (_ GlSupport) ProgramUniform4f() bool {
	return C.cancallProgramUniform4f() == 1
}

//	If Supports.ProgramUniform4f() is true, calls ProgramUniform4f() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform4f(program Uint, location Int, v0 Float, v1 Float, v2 Float, v3 Float) (err error) {
	if !Supports.ProgramUniform4f() {
		err = fmt.Errorf("ProgramUniform4f() function call not supported by current GL context.")
	} else {
		ProgramUniform4f(program, location, v0, v1, v2, v3)
		err = Util.Error("ProgramUniform4f()")
	}
	return
}

//	Returns whether the current GL context supports calling the ColorP3ui() function.
func (_ GlSupport) ColorP3ui() bool {
	return C.cancallColorP3ui() == 1
}

//	If Supports.ColorP3ui() is true, calls ColorP3ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ColorP3ui(type_ Enum, color Uint) (err error) {
	if !Supports.ColorP3ui() {
		err = fmt.Errorf("ColorP3ui() function call not supported by current GL context.")
	} else {
		ColorP3ui(type_, color)
		err = Util.Error("ColorP3ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the BlendEquationSeparate() function.
func (_ GlSupport) BlendEquationSeparate() bool {
	return C.cancallBlendEquationSeparate() == 1
}

//	If Supports.BlendEquationSeparate() is true, calls BlendEquationSeparate() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BlendEquationSeparate(modeRGB Enum, modeAlpha Enum) (err error) {
	if !Supports.BlendEquationSeparate() {
		err = fmt.Errorf("BlendEquationSeparate() function call not supported by current GL context.")
	} else {
		BlendEquationSeparate(modeRGB, modeAlpha)
		err = Util.Error("BlendEquationSeparate()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform4i() function.
func (_ GlSupport) ProgramUniform4i() bool {
	return C.cancallProgramUniform4i() == 1
}

//	If Supports.ProgramUniform4i() is true, calls ProgramUniform4i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform4i(program Uint, location Int, v0 Int, v1 Int, v2 Int, v3 Int) (err error) {
	if !Supports.ProgramUniform4i() {
		err = fmt.Errorf("ProgramUniform4i() function call not supported by current GL context.")
	} else {
		ProgramUniform4i(program, location, v0, v1, v2, v3)
		err = Util.Error("ProgramUniform4i()")
	}
	return
}

//	Returns whether the current GL context supports calling the Flush() function.
func (_ GlSupport) Flush() bool {
	return C.cancallFlush() == 1
}

//	If Supports.Flush() is true, calls Flush() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Flush() (err error) {
	if !Supports.Flush() {
		err = fmt.Errorf("Flush() function call not supported by current GL context.")
	} else {
		Flush()
		err = Util.Error("Flush()")
	}
	return
}

//	Returns whether the current GL context supports calling the Finish() function.
func (_ GlSupport) Finish() bool {
	return C.cancallFinish() == 1
}

//	If Supports.Finish() is true, calls Finish() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Finish() (err error) {
	if !Supports.Finish() {
		err = fmt.Errorf("Finish() function call not supported by current GL context.")
	} else {
		Finish()
		err = Util.Error("Finish()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix3dv() function.
func (_ GlSupport) UniformMatrix3dv() bool {
	return C.cancallUniformMatrix3dv() == 1
}

//	If Supports.UniformMatrix3dv() is true, calls UniformMatrix3dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix3dv(location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.UniformMatrix3dv() {
		err = fmt.Errorf("UniformMatrix3dv() function call not supported by current GL context.")
	} else {
		UniformMatrix3dv(location, count, transpose, value)
		err = Util.Error("UniformMatrix3dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the EndTransformFeedback() function.
func (_ GlSupport) EndTransformFeedback() bool {
	return C.cancallEndTransformFeedback() == 1
}

//	If Supports.EndTransformFeedback() is true, calls EndTransformFeedback() and yields the err returned by Util.Error(), if any.
func (_ GlTry) EndTransformFeedback() (err error) {
	if !Supports.EndTransformFeedback() {
		err = fmt.Errorf("EndTransformFeedback() function call not supported by current GL context.")
	} else {
		EndTransformFeedback()
		err = Util.Error("EndTransformFeedback()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform3d() function.
func (_ GlSupport) ProgramUniform3d() bool {
	return C.cancallProgramUniform3d() == 1
}

//	If Supports.ProgramUniform3d() is true, calls ProgramUniform3d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform3d(program Uint, location Int, v0 Double, v1 Double, v2 Double) (err error) {
	if !Supports.ProgramUniform3d() {
		err = fmt.Errorf("ProgramUniform3d() function call not supported by current GL context.")
	} else {
		ProgramUniform3d(program, location, v0, v1, v2)
		err = Util.Error("ProgramUniform3d()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetIntegerv() function.
func (_ GlSupport) GetIntegerv() bool {
	return C.cancallGetIntegerv() == 1
}

//	If Supports.GetIntegerv() is true, calls GetIntegerv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetIntegerv(pname Enum, params *Int) (err error) {
	if !Supports.GetIntegerv() {
		err = fmt.Errorf("GetIntegerv() function call not supported by current GL context.")
	} else {
		GetIntegerv(pname, params)
		err = Util.Error("GetIntegerv()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix4x3fv() function.
func (_ GlSupport) UniformMatrix4x3fv() bool {
	return C.cancallUniformMatrix4x3fv() == 1
}

//	If Supports.UniformMatrix4x3fv() is true, calls UniformMatrix4x3fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix4x3fv(location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.UniformMatrix4x3fv() {
		err = fmt.Errorf("UniformMatrix4x3fv() function call not supported by current GL context.")
	} else {
		UniformMatrix4x3fv(location, count, transpose, value)
		err = Util.Error("UniformMatrix4x3fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetTransformFeedbackVarying() function.
func (_ GlSupport) GetTransformFeedbackVarying() bool {
	return C.cancallGetTransformFeedbackVarying() == 1
}

//	If Supports.GetTransformFeedbackVarying() is true, calls GetTransformFeedbackVarying() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetTransformFeedbackVarying(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Sizei, type_ *Enum, name *Char) (err error) {
	if !Supports.GetTransformFeedbackVarying() {
		err = fmt.Errorf("GetTransformFeedbackVarying() function call not supported by current GL context.")
	} else {
		GetTransformFeedbackVarying(program, index, bufSize, length, size, type_, name)
		err = Util.Error("GetTransformFeedbackVarying()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform4d() function.
func (_ GlSupport) ProgramUniform4d() bool {
	return C.cancallProgramUniform4d() == 1
}

//	If Supports.ProgramUniform4d() is true, calls ProgramUniform4d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform4d(program Uint, location Int, v0 Double, v1 Double, v2 Double, v3 Double) (err error) {
	if !Supports.ProgramUniform4d() {
		err = fmt.Errorf("ProgramUniform4d() function call not supported by current GL context.")
	} else {
		ProgramUniform4d(program, location, v0, v1, v2, v3)
		err = Util.Error("ProgramUniform4d()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramResourceName() function.
func (_ GlSupport) GetProgramResourceName() bool {
	return C.cancallGetProgramResourceName() == 1
}

//	If Supports.GetProgramResourceName() is true, calls GetProgramResourceName() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramResourceName(program Uint, programInterface Enum, index Uint, bufSize Sizei, length *Sizei, name *Char) (err error) {
	if !Supports.GetProgramResourceName() {
		err = fmt.Errorf("GetProgramResourceName() function call not supported by current GL context.")
	} else {
		GetProgramResourceName(program, programInterface, index, bufSize, length, name)
		err = Util.Error("GetProgramResourceName()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteProgramPipelines() function.
func (_ GlSupport) DeleteProgramPipelines() bool {
	return C.cancallDeleteProgramPipelines() == 1
}

//	If Supports.DeleteProgramPipelines() is true, calls DeleteProgramPipelines() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteProgramPipelines(n Sizei, pipelines *Uint) (err error) {
	if !Supports.DeleteProgramPipelines() {
		err = fmt.Errorf("DeleteProgramPipelines() function call not supported by current GL context.")
	} else {
		DeleteProgramPipelines(n, pipelines)
		err = Util.Error("DeleteProgramPipelines()")
	}
	return
}

//	Returns whether the current GL context supports calling the Disable() function.
func (_ GlSupport) Disable() bool {
	return C.cancallDisable() == 1
}

//	If Supports.Disable() is true, calls Disable() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Disable(cap Enum) (err error) {
	if !Supports.Disable() {
		err = fmt.Errorf("Disable() function call not supported by current GL context.")
	} else {
		Disable(cap)
		err = Util.Error("Disable()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribLPointer() function.
func (_ GlSupport) VertexAttribLPointer() bool {
	return C.cancallVertexAttribLPointer() == 1
}

//	If Supports.VertexAttribLPointer() is true, calls VertexAttribLPointer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribLPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Ptr) (err error) {
	if !Supports.VertexAttribLPointer() {
		err = fmt.Errorf("VertexAttribLPointer() function call not supported by current GL context.")
	} else {
		VertexAttribLPointer(index, size, type_, stride, pointer)
		err = Util.Error("VertexAttribLPointer()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribP3ui() function.
func (_ GlSupport) VertexAttribP3ui() bool {
	return C.cancallVertexAttribP3ui() == 1
}

//	If Supports.VertexAttribP3ui() is true, calls VertexAttribP3ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribP3ui(index Uint, type_ Enum, normalized Boolean, value Uint) (err error) {
	if !Supports.VertexAttribP3ui() {
		err = fmt.Errorf("VertexAttribP3ui() function call not supported by current GL context.")
	} else {
		VertexAttribP3ui(index, type_, normalized, value)
		err = Util.Error("VertexAttribP3ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProvokingVertex() function.
func (_ GlSupport) ProvokingVertex() bool {
	return C.cancallProvokingVertex() == 1
}

//	If Supports.ProvokingVertex() is true, calls ProvokingVertex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProvokingVertex(mode Enum) (err error) {
	if !Supports.ProvokingVertex() {
		err = fmt.Errorf("ProvokingVertex() function call not supported by current GL context.")
	} else {
		ProvokingVertex(mode)
		err = Util.Error("ProvokingVertex()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawArraysInstanced() function.
func (_ GlSupport) DrawArraysInstanced() bool {
	return C.cancallDrawArraysInstanced() == 1
}

//	If Supports.DrawArraysInstanced() is true, calls DrawArraysInstanced() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawArraysInstanced(mode Enum, first Int, count Sizei, instancecount Sizei) (err error) {
	if !Supports.DrawArraysInstanced() {
		err = fmt.Errorf("DrawArraysInstanced() function call not supported by current GL context.")
	} else {
		DrawArraysInstanced(mode, first, count, instancecount)
		err = Util.Error("DrawArraysInstanced()")
	}
	return
}

//	Returns whether the current GL context supports calling the NormalP3uiv() function.
func (_ GlSupport) NormalP3uiv() bool {
	return C.cancallNormalP3uiv() == 1
}

//	If Supports.NormalP3uiv() is true, calls NormalP3uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) NormalP3uiv(type_ Enum, coords *Uint) (err error) {
	if !Supports.NormalP3uiv() {
		err = fmt.Errorf("NormalP3uiv() function call not supported by current GL context.")
	} else {
		NormalP3uiv(type_, coords)
		err = Util.Error("NormalP3uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform2iv() function.
func (_ GlSupport) Uniform2iv() bool {
	return C.cancallUniform2iv() == 1
}

//	If Supports.Uniform2iv() is true, calls Uniform2iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform2iv(location Int, count Sizei, value *Int) (err error) {
	if !Supports.Uniform2iv() {
		err = fmt.Errorf("Uniform2iv() function call not supported by current GL context.")
	} else {
		Uniform2iv(location, count, value)
		err = Util.Error("Uniform2iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform1d() function.
func (_ GlSupport) ProgramUniform1d() bool {
	return C.cancallProgramUniform1d() == 1
}

//	If Supports.ProgramUniform1d() is true, calls ProgramUniform1d() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform1d(program Uint, location Int, v0 Double) (err error) {
	if !Supports.ProgramUniform1d() {
		err = fmt.Errorf("ProgramUniform1d() function call not supported by current GL context.")
	} else {
		ProgramUniform1d(program, location, v0)
		err = Util.Error("ProgramUniform1d()")
	}
	return
}

//	Returns whether the current GL context supports calling the PolygonOffset() function.
func (_ GlSupport) PolygonOffset() bool {
	return C.cancallPolygonOffset() == 1
}

//	If Supports.PolygonOffset() is true, calls PolygonOffset() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PolygonOffset(factor Float, units Float) (err error) {
	if !Supports.PolygonOffset() {
		err = fmt.Errorf("PolygonOffset() function call not supported by current GL context.")
	} else {
		PolygonOffset(factor, units)
		err = Util.Error("PolygonOffset()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix3x4dv() function.
func (_ GlSupport) UniformMatrix3x4dv() bool {
	return C.cancallUniformMatrix3x4dv() == 1
}

//	If Supports.UniformMatrix3x4dv() is true, calls UniformMatrix3x4dv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix3x4dv(location Int, count Sizei, transpose Boolean, value *Double) (err error) {
	if !Supports.UniformMatrix3x4dv() {
		err = fmt.Errorf("UniformMatrix3x4dv() function call not supported by current GL context.")
	} else {
		UniformMatrix3x4dv(location, count, transpose, value)
		err = Util.Error("UniformMatrix3x4dv()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexStorage3D() function.
func (_ GlSupport) TexStorage3D() bool {
	return C.cancallTexStorage3D() == 1
}

//	If Supports.TexStorage3D() is true, calls TexStorage3D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexStorage3D(target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei, depth Sizei) (err error) {
	if !Supports.TexStorage3D() {
		err = fmt.Errorf("TexStorage3D() function call not supported by current GL context.")
	} else {
		TexStorage3D(target, levels, internalformat, width, height, depth)
		err = Util.Error("TexStorage3D()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetTexParameterIiv() function.
func (_ GlSupport) GetTexParameterIiv() bool {
	return C.cancallGetTexParameterIiv() == 1
}

//	If Supports.GetTexParameterIiv() is true, calls GetTexParameterIiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetTexParameterIiv(target Enum, pname Enum, params *Int) (err error) {
	if !Supports.GetTexParameterIiv() {
		err = fmt.Errorf("GetTexParameterIiv() function call not supported by current GL context.")
	} else {
		GetTexParameterIiv(target, pname, params)
		err = Util.Error("GetTexParameterIiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix4fv() function.
func (_ GlSupport) ProgramUniformMatrix4fv() bool {
	return C.cancallProgramUniformMatrix4fv() == 1
}

//	If Supports.ProgramUniformMatrix4fv() is true, calls ProgramUniformMatrix4fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.ProgramUniformMatrix4fv() {
		err = fmt.Errorf("ProgramUniformMatrix4fv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix4fv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix4fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribP2uiv() function.
func (_ GlSupport) VertexAttribP2uiv() bool {
	return C.cancallVertexAttribP2uiv() == 1
}

//	If Supports.VertexAttribP2uiv() is true, calls VertexAttribP2uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribP2uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) (err error) {
	if !Supports.VertexAttribP2uiv() {
		err = fmt.Errorf("VertexAttribP2uiv() function call not supported by current GL context.")
	} else {
		VertexAttribP2uiv(index, type_, normalized, value)
		err = Util.Error("VertexAttribP2uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the CopyTexSubImage1D() function.
func (_ GlSupport) CopyTexSubImage1D() bool {
	return C.cancallCopyTexSubImage1D() == 1
}

//	If Supports.CopyTexSubImage1D() is true, calls CopyTexSubImage1D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CopyTexSubImage1D(target Enum, level Int, xoffset Int, x Int, y Int, width Sizei) (err error) {
	if !Supports.CopyTexSubImage1D() {
		err = fmt.Errorf("CopyTexSubImage1D() function call not supported by current GL context.")
	} else {
		CopyTexSubImage1D(target, level, xoffset, x, y, width)
		err = Util.Error("CopyTexSubImage1D()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetSynciv() function.
func (_ GlSupport) GetSynciv() bool {
	return C.cancallGetSynciv() == 1
}

//	If Supports.GetSynciv() is true, calls GetSynciv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetSynciv(sync Sync, pname Enum, bufSize Sizei, length *Sizei, values *Int) (err error) {
	if !Supports.GetSynciv() {
		err = fmt.Errorf("GetSynciv() function call not supported by current GL context.")
	} else {
		GetSynciv(sync, pname, bufSize, length, values)
		err = Util.Error("GetSynciv()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform4f() function.
func (_ GlSupport) Uniform4f() bool {
	return C.cancallUniform4f() == 1
}

//	If Supports.Uniform4f() is true, calls Uniform4f() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform4f(location Int, v0 Float, v1 Float, v2 Float, v3 Float) (err error) {
	if !Supports.Uniform4f() {
		err = fmt.Errorf("Uniform4f() function call not supported by current GL context.")
	} else {
		Uniform4f(location, v0, v1, v2, v3)
		err = Util.Error("Uniform4f()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniformMatrix2fv() function.
func (_ GlSupport) ProgramUniformMatrix2fv() bool {
	return C.cancallProgramUniformMatrix2fv() == 1
}

//	If Supports.ProgramUniformMatrix2fv() is true, calls ProgramUniformMatrix2fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniformMatrix2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.ProgramUniformMatrix2fv() {
		err = fmt.Errorf("ProgramUniformMatrix2fv() function call not supported by current GL context.")
	} else {
		ProgramUniformMatrix2fv(program, location, count, transpose, value)
		err = Util.Error("ProgramUniformMatrix2fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DeleteShader() function.
func (_ GlSupport) DeleteShader() bool {
	return C.cancallDeleteShader() == 1
}

//	If Supports.DeleteShader() is true, calls DeleteShader() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DeleteShader(shader Uint) (err error) {
	if !Supports.DeleteShader() {
		err = fmt.Errorf("DeleteShader() function call not supported by current GL context.")
	} else {
		DeleteShader(shader)
		err = Util.Error("DeleteShader()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindBuffer() function.
func (_ GlSupport) BindBuffer() bool {
	return C.cancallBindBuffer() == 1
}

//	If Supports.BindBuffer() is true, calls BindBuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindBuffer(target Enum, buffer Uint) (err error) {
	if !Supports.BindBuffer() {
		err = fmt.Errorf("BindBuffer() function call not supported by current GL context.")
	} else {
		BindBuffer(target, buffer)
		err = Util.Error("BindBuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the Enable() function.
func (_ GlSupport) Enable() bool {
	return C.cancallEnable() == 1
}

//	If Supports.Enable() is true, calls Enable() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Enable(cap Enum) (err error) {
	if !Supports.Enable() {
		err = fmt.Errorf("Enable() function call not supported by current GL context.")
	} else {
		Enable(cap)
		err = Util.Error("Enable()")
	}
	return
}

//	Returns whether the current GL context supports calling the WaitSync() function.
func (_ GlSupport) WaitSync() bool {
	return C.cancallWaitSync() == 1
}

//	If Supports.WaitSync() is true, calls WaitSync() and yields the err returned by Util.Error(), if any.
func (_ GlTry) WaitSync(sync Sync, flags Bitfield, timeout Uint64) (err error) {
	if !Supports.WaitSync() {
		err = fmt.Errorf("WaitSync() function call not supported by current GL context.")
	} else {
		WaitSync(sync, flags, timeout)
		err = Util.Error("WaitSync()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindAttribLocation() function.
func (_ GlSupport) BindAttribLocation() bool {
	return C.cancallBindAttribLocation() == 1
}

//	If Supports.BindAttribLocation() is true, calls BindAttribLocation() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindAttribLocation(program Uint, index Uint, name *Char) (err error) {
	if !Supports.BindAttribLocation() {
		err = fmt.Errorf("BindAttribLocation() function call not supported by current GL context.")
	} else {
		BindAttribLocation(program, index, name)
		err = Util.Error("BindAttribLocation()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI4iv() function.
func (_ GlSupport) VertexAttribI4iv() bool {
	return C.cancallVertexAttribI4iv() == 1
}

//	If Supports.VertexAttribI4iv() is true, calls VertexAttribI4iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI4iv(index Uint, v *Int) (err error) {
	if !Supports.VertexAttribI4iv() {
		err = fmt.Errorf("VertexAttribI4iv() function call not supported by current GL context.")
	} else {
		VertexAttribI4iv(index, v)
		err = Util.Error("VertexAttribI4iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramStageiv() function.
func (_ GlSupport) GetProgramStageiv() bool {
	return C.cancallGetProgramStageiv() == 1
}

//	If Supports.GetProgramStageiv() is true, calls GetProgramStageiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramStageiv(program Uint, shadertype Enum, pname Enum, values *Int) (err error) {
	if !Supports.GetProgramStageiv() {
		err = fmt.Errorf("GetProgramStageiv() function call not supported by current GL context.")
	} else {
		GetProgramStageiv(program, shadertype, pname, values)
		err = Util.Error("GetProgramStageiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexCoordP2uiv() function.
func (_ GlSupport) TexCoordP2uiv() bool {
	return C.cancallTexCoordP2uiv() == 1
}

//	If Supports.TexCoordP2uiv() is true, calls TexCoordP2uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexCoordP2uiv(type_ Enum, coords *Uint) (err error) {
	if !Supports.TexCoordP2uiv() {
		err = fmt.Errorf("TexCoordP2uiv() function call not supported by current GL context.")
	} else {
		TexCoordP2uiv(type_, coords)
		err = Util.Error("TexCoordP2uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiTexCoordP4ui() function.
func (_ GlSupport) MultiTexCoordP4ui() bool {
	return C.cancallMultiTexCoordP4ui() == 1
}

//	If Supports.MultiTexCoordP4ui() is true, calls MultiTexCoordP4ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiTexCoordP4ui(texture Enum, type_ Enum, coords Uint) (err error) {
	if !Supports.MultiTexCoordP4ui() {
		err = fmt.Errorf("MultiTexCoordP4ui() function call not supported by current GL context.")
	} else {
		MultiTexCoordP4ui(texture, type_, coords)
		err = Util.Error("MultiTexCoordP4ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform3iv() function.
func (_ GlSupport) Uniform3iv() bool {
	return C.cancallUniform3iv() == 1
}

//	If Supports.Uniform3iv() is true, calls Uniform3iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform3iv(location Int, count Sizei, value *Int) (err error) {
	if !Supports.Uniform3iv() {
		err = fmt.Errorf("Uniform3iv() function call not supported by current GL context.")
	} else {
		Uniform3iv(location, count, value)
		err = Util.Error("Uniform3iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetIntegeri_v() function.
func (_ GlSupport) GetIntegeri_v() bool {
	return C.cancallGetIntegeri_v() == 1
}

//	If Supports.GetIntegeri_v() is true, calls GetIntegeri_v() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetIntegeri_v(target Enum, index Uint, data *Int) (err error) {
	if !Supports.GetIntegeri_v() {
		err = fmt.Errorf("GetIntegeri_v() function call not supported by current GL context.")
	} else {
		GetIntegeri_v(target, index, data)
		err = Util.Error("GetIntegeri_v()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawElementsInstanced() function.
func (_ GlSupport) DrawElementsInstanced() bool {
	return C.cancallDrawElementsInstanced() == 1
}

//	If Supports.DrawElementsInstanced() is true, calls DrawElementsInstanced() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawElementsInstanced(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei) (err error) {
	if !Supports.DrawElementsInstanced() {
		err = fmt.Errorf("DrawElementsInstanced() function call not supported by current GL context.")
	} else {
		DrawElementsInstanced(mode, count, type_, indices, instancecount)
		err = Util.Error("DrawElementsInstanced()")
	}
	return
}

//	Returns whether the current GL context supports calling the MinSampleShading() function.
func (_ GlSupport) MinSampleShading() bool {
	return C.cancallMinSampleShading() == 1
}

//	If Supports.MinSampleShading() is true, calls MinSampleShading() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MinSampleShading(value Float) (err error) {
	if !Supports.MinSampleShading() {
		err = fmt.Errorf("MinSampleShading() function call not supported by current GL context.")
	} else {
		MinSampleShading(value)
		err = Util.Error("MinSampleShading()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetShaderSource() function.
func (_ GlSupport) GetShaderSource() bool {
	return C.cancallGetShaderSource() == 1
}

//	If Supports.GetShaderSource() is true, calls GetShaderSource() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetShaderSource(shader Uint, bufSize Sizei, length *Sizei, source *Char) (err error) {
	if !Supports.GetShaderSource() {
		err = fmt.Errorf("GetShaderSource() function call not supported by current GL context.")
	} else {
		GetShaderSource(shader, bufSize, length, source)
		err = Util.Error("GetShaderSource()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform2uiv() function.
func (_ GlSupport) ProgramUniform2uiv() bool {
	return C.cancallProgramUniform2uiv() == 1
}

//	If Supports.ProgramUniform2uiv() is true, calls ProgramUniform2uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform2uiv(program Uint, location Int, count Sizei, value *Uint) (err error) {
	if !Supports.ProgramUniform2uiv() {
		err = fmt.Errorf("ProgramUniform2uiv() function call not supported by current GL context.")
	} else {
		ProgramUniform2uiv(program, location, count, value)
		err = Util.Error("ProgramUniform2uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the StencilMaskSeparate() function.
func (_ GlSupport) StencilMaskSeparate() bool {
	return C.cancallStencilMaskSeparate() == 1
}

//	If Supports.StencilMaskSeparate() is true, calls StencilMaskSeparate() and yields the err returned by Util.Error(), if any.
func (_ GlTry) StencilMaskSeparate(face Enum, mask Uint) (err error) {
	if !Supports.StencilMaskSeparate() {
		err = fmt.Errorf("StencilMaskSeparate() function call not supported by current GL context.")
	} else {
		StencilMaskSeparate(face, mask)
		err = Util.Error("StencilMaskSeparate()")
	}
	return
}

//	Returns whether the current GL context supports calling the Scissor() function.
func (_ GlSupport) Scissor() bool {
	return C.cancallScissor() == 1
}

//	If Supports.Scissor() is true, calls Scissor() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Scissor(x Int, y Int, width Sizei, height Sizei) (err error) {
	if !Supports.Scissor() {
		err = fmt.Errorf("Scissor() function call not supported by current GL context.")
	} else {
		Scissor(x, y, width, height)
		err = Util.Error("Scissor()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetShaderPrecisionFormat() function.
func (_ GlSupport) GetShaderPrecisionFormat() bool {
	return C.cancallGetShaderPrecisionFormat() == 1
}

//	If Supports.GetShaderPrecisionFormat() is true, calls GetShaderPrecisionFormat() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetShaderPrecisionFormat(shadertype Enum, precisiontype Enum, range_ *Int, precision *Int) (err error) {
	if !Supports.GetShaderPrecisionFormat() {
		err = fmt.Errorf("GetShaderPrecisionFormat() function call not supported by current GL context.")
	} else {
		GetShaderPrecisionFormat(shadertype, precisiontype, range_, precision)
		err = Util.Error("GetShaderPrecisionFormat()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI4ui() function.
func (_ GlSupport) VertexAttribI4ui() bool {
	return C.cancallVertexAttribI4ui() == 1
}

//	If Supports.VertexAttribI4ui() is true, calls VertexAttribI4ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI4ui(index Uint, x Uint, y Uint, z Uint, w Uint) (err error) {
	if !Supports.VertexAttribI4ui() {
		err = fmt.Errorf("VertexAttribI4ui() function call not supported by current GL context.")
	} else {
		VertexAttribI4ui(index, x, y, z, w)
		err = Util.Error("VertexAttribI4ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the ColorP3uiv() function.
func (_ GlSupport) ColorP3uiv() bool {
	return C.cancallColorP3uiv() == 1
}

//	If Supports.ColorP3uiv() is true, calls ColorP3uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ColorP3uiv(type_ Enum, color *Uint) (err error) {
	if !Supports.ColorP3uiv() {
		err = fmt.Errorf("ColorP3uiv() function call not supported by current GL context.")
	} else {
		ColorP3uiv(type_, color)
		err = Util.Error("ColorP3uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetActiveUniformBlockName() function.
func (_ GlSupport) GetActiveUniformBlockName() bool {
	return C.cancallGetActiveUniformBlockName() == 1
}

//	If Supports.GetActiveUniformBlockName() is true, calls GetActiveUniformBlockName() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetActiveUniformBlockName(program Uint, uniformBlockIndex Uint, bufSize Sizei, length *Sizei, uniformBlockName *Char) (err error) {
	if !Supports.GetActiveUniformBlockName() {
		err = fmt.Errorf("GetActiveUniformBlockName() function call not supported by current GL context.")
	} else {
		GetActiveUniformBlockName(program, uniformBlockIndex, bufSize, length, uniformBlockName)
		err = Util.Error("GetActiveUniformBlockName()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexStorage3DMultisample() function.
func (_ GlSupport) TexStorage3DMultisample() bool {
	return C.cancallTexStorage3DMultisample() == 1
}

//	If Supports.TexStorage3DMultisample() is true, calls TexStorage3DMultisample() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexStorage3DMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean) (err error) {
	if !Supports.TexStorage3DMultisample() {
		err = fmt.Errorf("TexStorage3DMultisample() function call not supported by current GL context.")
	} else {
		TexStorage3DMultisample(target, samples, internalformat, width, height, depth, fixedsamplelocations)
		err = Util.Error("TexStorage3DMultisample()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI3uiv() function.
func (_ GlSupport) VertexAttribI3uiv() bool {
	return C.cancallVertexAttribI3uiv() == 1
}

//	If Supports.VertexAttribI3uiv() is true, calls VertexAttribI3uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI3uiv(index Uint, v *Uint) (err error) {
	if !Supports.VertexAttribI3uiv() {
		err = fmt.Errorf("VertexAttribI3uiv() function call not supported by current GL context.")
	} else {
		VertexAttribI3uiv(index, v)
		err = Util.Error("VertexAttribI3uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexP2uiv() function.
func (_ GlSupport) VertexP2uiv() bool {
	return C.cancallVertexP2uiv() == 1
}

//	If Supports.VertexP2uiv() is true, calls VertexP2uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexP2uiv(type_ Enum, value *Uint) (err error) {
	if !Supports.VertexP2uiv() {
		err = fmt.Errorf("VertexP2uiv() function call not supported by current GL context.")
	} else {
		VertexP2uiv(type_, value)
		err = Util.Error("VertexP2uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the MultiDrawElementsIndirect() function.
func (_ GlSupport) MultiDrawElementsIndirect() bool {
	return C.cancallMultiDrawElementsIndirect() == 1
}

//	If Supports.MultiDrawElementsIndirect() is true, calls MultiDrawElementsIndirect() and yields the err returned by Util.Error(), if any.
func (_ GlTry) MultiDrawElementsIndirect(mode Enum, type_ Enum, indirect Ptr, drawcount Sizei, stride Sizei) (err error) {
	if !Supports.MultiDrawElementsIndirect() {
		err = fmt.Errorf("MultiDrawElementsIndirect() function call not supported by current GL context.")
	} else {
		MultiDrawElementsIndirect(mode, type_, indirect, drawcount, stride)
		err = Util.Error("MultiDrawElementsIndirect()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform4uiv() function.
func (_ GlSupport) Uniform4uiv() bool {
	return C.cancallUniform4uiv() == 1
}

//	If Supports.Uniform4uiv() is true, calls Uniform4uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform4uiv(location Int, count Sizei, value *Uint) (err error) {
	if !Supports.Uniform4uiv() {
		err = fmt.Errorf("Uniform4uiv() function call not supported by current GL context.")
	} else {
		Uniform4uiv(location, count, value)
		err = Util.Error("Uniform4uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetTexParameteriv() function.
func (_ GlSupport) GetTexParameteriv() bool {
	return C.cancallGetTexParameteriv() == 1
}

//	If Supports.GetTexParameteriv() is true, calls GetTexParameteriv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetTexParameteriv(target Enum, pname Enum, params *Int) (err error) {
	if !Supports.GetTexParameteriv() {
		err = fmt.Errorf("GetTexParameteriv() function call not supported by current GL context.")
	} else {
		GetTexParameteriv(target, pname, params)
		err = Util.Error("GetTexParameteriv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexP2ui() function.
func (_ GlSupport) VertexP2ui() bool {
	return C.cancallVertexP2ui() == 1
}

//	If Supports.VertexP2ui() is true, calls VertexP2ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexP2ui(type_ Enum, value Uint) (err error) {
	if !Supports.VertexP2ui() {
		err = fmt.Errorf("VertexP2ui() function call not supported by current GL context.")
	} else {
		VertexP2ui(type_, value)
		err = Util.Error("VertexP2ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the GenVertexArrays() function.
func (_ GlSupport) GenVertexArrays() bool {
	return C.cancallGenVertexArrays() == 1
}

//	If Supports.GenVertexArrays() is true, calls GenVertexArrays() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GenVertexArrays(n Sizei, arrays *Uint) (err error) {
	if !Supports.GenVertexArrays() {
		err = fmt.Errorf("GenVertexArrays() function call not supported by current GL context.")
	} else {
		GenVertexArrays(n, arrays)
		err = Util.Error("GenVertexArrays()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramResourceLocationIndex() function.
func (_ GlSupport) GetProgramResourceLocationIndex() bool {
	return C.cancallGetProgramResourceLocationIndex() == 1
}

//	If Supports.GetProgramResourceLocationIndex() is true, calls GetProgramResourceLocationIndex() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramResourceLocationIndex(program Uint, programInterface Enum, name *Char) (err error, tryRetVal__ Int) {
	if !Supports.GetProgramResourceLocationIndex() {
		err = fmt.Errorf("GetProgramResourceLocationIndex() function call not supported by current GL context.")
	} else {
		tryRetVal__ = GetProgramResourceLocationIndex(program, programInterface, name)
		err = Util.Error("GetProgramResourceLocationIndex()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetSamplerParameteriv() function.
func (_ GlSupport) GetSamplerParameteriv() bool {
	return C.cancallGetSamplerParameteriv() == 1
}

//	If Supports.GetSamplerParameteriv() is true, calls GetSamplerParameteriv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetSamplerParameteriv(sampler Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetSamplerParameteriv() {
		err = fmt.Errorf("GetSamplerParameteriv() function call not supported by current GL context.")
	} else {
		GetSamplerParameteriv(sampler, pname, params)
		err = Util.Error("GetSamplerParameteriv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DepthRangef() function.
func (_ GlSupport) DepthRangef() bool {
	return C.cancallDepthRangef() == 1
}

//	If Supports.DepthRangef() is true, calls DepthRangef() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DepthRangef(n Float, f Float) (err error) {
	if !Supports.DepthRangef() {
		err = fmt.Errorf("DepthRangef() function call not supported by current GL context.")
	} else {
		DepthRangef(n, f)
		err = Util.Error("DepthRangef()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI1ui() function.
func (_ GlSupport) VertexAttribI1ui() bool {
	return C.cancallVertexAttribI1ui() == 1
}

//	If Supports.VertexAttribI1ui() is true, calls VertexAttribI1ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI1ui(index Uint, x Uint) (err error) {
	if !Supports.VertexAttribI1ui() {
		err = fmt.Errorf("VertexAttribI1ui() function call not supported by current GL context.")
	} else {
		VertexAttribI1ui(index, x)
		err = Util.Error("VertexAttribI1ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClearColor() function.
func (_ GlSupport) ClearColor() bool {
	return C.cancallClearColor() == 1
}

//	If Supports.ClearColor() is true, calls ClearColor() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClearColor(red Float, green Float, blue Float, alpha Float) (err error) {
	if !Supports.ClearColor() {
		err = fmt.Errorf("ClearColor() function call not supported by current GL context.")
	} else {
		ClearColor(red, green, blue, alpha)
		err = Util.Error("ClearColor()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexParameterIuiv() function.
func (_ GlSupport) TexParameterIuiv() bool {
	return C.cancallTexParameterIuiv() == 1
}

//	If Supports.TexParameterIuiv() is true, calls TexParameterIuiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexParameterIuiv(target Enum, pname Enum, params *Uint) (err error) {
	if !Supports.TexParameterIuiv() {
		err = fmt.Errorf("TexParameterIuiv() function call not supported by current GL context.")
	} else {
		TexParameterIuiv(target, pname, params)
		err = Util.Error("TexParameterIuiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the ProgramUniform2fv() function.
func (_ GlSupport) ProgramUniform2fv() bool {
	return C.cancallProgramUniform2fv() == 1
}

//	If Supports.ProgramUniform2fv() is true, calls ProgramUniform2fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ProgramUniform2fv(program Uint, location Int, count Sizei, value *Float) (err error) {
	if !Supports.ProgramUniform2fv() {
		err = fmt.Errorf("ProgramUniform2fv() function call not supported by current GL context.")
	} else {
		ProgramUniform2fv(program, location, count, value)
		err = Util.Error("ProgramUniform2fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribBinding() function.
func (_ GlSupport) VertexAttribBinding() bool {
	return C.cancallVertexAttribBinding() == 1
}

//	If Supports.VertexAttribBinding() is true, calls VertexAttribBinding() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribBinding(attribindex Uint, bindingindex Uint) (err error) {
	if !Supports.VertexAttribBinding() {
		err = fmt.Errorf("VertexAttribBinding() function call not supported by current GL context.")
	} else {
		VertexAttribBinding(attribindex, bindingindex)
		err = Util.Error("VertexAttribBinding()")
	}
	return
}

//	Returns whether the current GL context supports calling the TexCoordP4uiv() function.
func (_ GlSupport) TexCoordP4uiv() bool {
	return C.cancallTexCoordP4uiv() == 1
}

//	If Supports.TexCoordP4uiv() is true, calls TexCoordP4uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) TexCoordP4uiv(type_ Enum, coords *Uint) (err error) {
	if !Supports.TexCoordP4uiv() {
		err = fmt.Errorf("TexCoordP4uiv() function call not supported by current GL context.")
	} else {
		TexCoordP4uiv(type_, coords)
		err = Util.Error("TexCoordP4uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the PatchParameterfv() function.
func (_ GlSupport) PatchParameterfv() bool {
	return C.cancallPatchParameterfv() == 1
}

//	If Supports.PatchParameterfv() is true, calls PatchParameterfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PatchParameterfv(pname Enum, values *Float) (err error) {
	if !Supports.PatchParameterfv() {
		err = fmt.Errorf("PatchParameterfv() function call not supported by current GL context.")
	} else {
		PatchParameterfv(pname, values)
		err = Util.Error("PatchParameterfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the CompressedTexImage2D() function.
func (_ GlSupport) CompressedTexImage2D() bool {
	return C.cancallCompressedTexImage2D() == 1
}

//	If Supports.CompressedTexImage2D() is true, calls CompressedTexImage2D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CompressedTexImage2D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, border Int, imageSize Sizei, data Ptr) (err error) {
	if !Supports.CompressedTexImage2D() {
		err = fmt.Errorf("CompressedTexImage2D() function call not supported by current GL context.")
	} else {
		CompressedTexImage2D(target, level, internalformat, width, height, border, imageSize, data)
		err = Util.Error("CompressedTexImage2D()")
	}
	return
}

//	Returns whether the current GL context supports calling the ActiveTexture() function.
func (_ GlSupport) ActiveTexture() bool {
	return C.cancallActiveTexture() == 1
}

//	If Supports.ActiveTexture() is true, calls ActiveTexture() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ActiveTexture(texture Enum) (err error) {
	if !Supports.ActiveTexture() {
		err = fmt.Errorf("ActiveTexture() function call not supported by current GL context.")
	} else {
		ActiveTexture(texture)
		err = Util.Error("ActiveTexture()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetProgramBinary() function.
func (_ GlSupport) GetProgramBinary() bool {
	return C.cancallGetProgramBinary() == 1
}

//	If Supports.GetProgramBinary() is true, calls GetProgramBinary() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetProgramBinary(program Uint, bufSize Sizei, length *Sizei, binaryFormat *Enum, binary Ptr) (err error) {
	if !Supports.GetProgramBinary() {
		err = fmt.Errorf("GetProgramBinary() function call not supported by current GL context.")
	} else {
		GetProgramBinary(program, bufSize, length, binaryFormat, binary)
		err = Util.Error("GetProgramBinary()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform1ui() function.
func (_ GlSupport) Uniform1ui() bool {
	return C.cancallUniform1ui() == 1
}

//	If Supports.Uniform1ui() is true, calls Uniform1ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform1ui(location Int, v0 Uint) (err error) {
	if !Supports.Uniform1ui() {
		err = fmt.Errorf("Uniform1ui() function call not supported by current GL context.")
	} else {
		Uniform1ui(location, v0)
		err = Util.Error("Uniform1ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the CopyTexImage1D() function.
func (_ GlSupport) CopyTexImage1D() bool {
	return C.cancallCopyTexImage1D() == 1
}

//	If Supports.CopyTexImage1D() is true, calls CopyTexImage1D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CopyTexImage1D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, border Int) (err error) {
	if !Supports.CopyTexImage1D() {
		err = fmt.Errorf("CopyTexImage1D() function call not supported by current GL context.")
	} else {
		CopyTexImage1D(target, level, internalformat, x, y, width, border)
		err = Util.Error("CopyTexImage1D()")
	}
	return
}

//	Returns whether the current GL context supports calling the DepthRangeIndexed() function.
func (_ GlSupport) DepthRangeIndexed() bool {
	return C.cancallDepthRangeIndexed() == 1
}

//	If Supports.DepthRangeIndexed() is true, calls DepthRangeIndexed() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DepthRangeIndexed(index Uint, n Double, f Double) (err error) {
	if !Supports.DepthRangeIndexed() {
		err = fmt.Errorf("DepthRangeIndexed() function call not supported by current GL context.")
	} else {
		DepthRangeIndexed(index, n, f)
		err = Util.Error("DepthRangeIndexed()")
	}
	return
}

//	Returns whether the current GL context supports calling the CopyTexImage2D() function.
func (_ GlSupport) CopyTexImage2D() bool {
	return C.cancallCopyTexImage2D() == 1
}

//	If Supports.CopyTexImage2D() is true, calls CopyTexImage2D() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CopyTexImage2D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, height Sizei, border Int) (err error) {
	if !Supports.CopyTexImage2D() {
		err = fmt.Errorf("CopyTexImage2D() function call not supported by current GL context.")
	} else {
		CopyTexImage2D(target, level, internalformat, x, y, width, height, border)
		err = Util.Error("CopyTexImage2D()")
	}
	return
}

//	Returns whether the current GL context supports calling the ScissorArrayv() function.
func (_ GlSupport) ScissorArrayv() bool {
	return C.cancallScissorArrayv() == 1
}

//	If Supports.ScissorArrayv() is true, calls ScissorArrayv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ScissorArrayv(first Uint, count Sizei, v *Int) (err error) {
	if !Supports.ScissorArrayv() {
		err = fmt.Errorf("ScissorArrayv() function call not supported by current GL context.")
	} else {
		ScissorArrayv(first, count, v)
		err = Util.Error("ScissorArrayv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetShaderInfoLog() function.
func (_ GlSupport) GetShaderInfoLog() bool {
	return C.cancallGetShaderInfoLog() == 1
}

//	If Supports.GetShaderInfoLog() is true, calls GetShaderInfoLog() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetShaderInfoLog(shader Uint, bufSize Sizei, length *Sizei, infoLog *Char) (err error) {
	if !Supports.GetShaderInfoLog() {
		err = fmt.Errorf("GetShaderInfoLog() function call not supported by current GL context.")
	} else {
		GetShaderInfoLog(shader, bufSize, length, infoLog)
		err = Util.Error("GetShaderInfoLog()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetMultisamplefv() function.
func (_ GlSupport) GetMultisamplefv() bool {
	return C.cancallGetMultisamplefv() == 1
}

//	If Supports.GetMultisamplefv() is true, calls GetMultisamplefv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetMultisamplefv(pname Enum, index Uint, val *Float) (err error) {
	if !Supports.GetMultisamplefv() {
		err = fmt.Errorf("GetMultisamplefv() function call not supported by current GL context.")
	} else {
		GetMultisamplefv(pname, index, val)
		err = Util.Error("GetMultisamplefv()")
	}
	return
}

//	Returns whether the current GL context supports calling the DrawArrays() function.
func (_ GlSupport) DrawArrays() bool {
	return C.cancallDrawArrays() == 1
}

//	If Supports.DrawArrays() is true, calls DrawArrays() and yields the err returned by Util.Error(), if any.
func (_ GlTry) DrawArrays(mode Enum, first Int, count Sizei) (err error) {
	if !Supports.DrawArrays() {
		err = fmt.Errorf("DrawArrays() function call not supported by current GL context.")
	} else {
		DrawArrays(mode, first, count)
		err = Util.Error("DrawArrays()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetActiveUniformBlockiv() function.
func (_ GlSupport) GetActiveUniformBlockiv() bool {
	return C.cancallGetActiveUniformBlockiv() == 1
}

//	If Supports.GetActiveUniformBlockiv() is true, calls GetActiveUniformBlockiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetActiveUniformBlockiv(program Uint, uniformBlockIndex Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetActiveUniformBlockiv() {
		err = fmt.Errorf("GetActiveUniformBlockiv() function call not supported by current GL context.")
	} else {
		GetActiveUniformBlockiv(program, uniformBlockIndex, pname, params)
		err = Util.Error("GetActiveUniformBlockiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the UniformMatrix3fv() function.
func (_ GlSupport) UniformMatrix3fv() bool {
	return C.cancallUniformMatrix3fv() == 1
}

//	If Supports.UniformMatrix3fv() is true, calls UniformMatrix3fv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) UniformMatrix3fv(location Int, count Sizei, transpose Boolean, value *Float) (err error) {
	if !Supports.UniformMatrix3fv() {
		err = fmt.Errorf("UniformMatrix3fv() function call not supported by current GL context.")
	} else {
		UniformMatrix3fv(location, count, transpose, value)
		err = Util.Error("UniformMatrix3fv()")
	}
	return
}

//	Returns whether the current GL context supports calling the CreateProgram() function.
func (_ GlSupport) CreateProgram() bool {
	return C.cancallCreateProgram() == 1
}

//	If Supports.CreateProgram() is true, calls CreateProgram() and yields the err returned by Util.Error(), if any.
func (_ GlTry) CreateProgram() (err error, tryRetVal__ Uint) {
	if !Supports.CreateProgram() {
		err = fmt.Errorf("CreateProgram() function call not supported by current GL context.")
	} else {
		tryRetVal__ = CreateProgram()
		err = Util.Error("CreateProgram()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI1uiv() function.
func (_ GlSupport) VertexAttribI1uiv() bool {
	return C.cancallVertexAttribI1uiv() == 1
}

//	If Supports.VertexAttribI1uiv() is true, calls VertexAttribI1uiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI1uiv(index Uint, v *Uint) (err error) {
	if !Supports.VertexAttribI1uiv() {
		err = fmt.Errorf("VertexAttribI1uiv() function call not supported by current GL context.")
	} else {
		VertexAttribI1uiv(index, v)
		err = Util.Error("VertexAttribI1uiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetQueryObjectuiv() function.
func (_ GlSupport) GetQueryObjectuiv() bool {
	return C.cancallGetQueryObjectuiv() == 1
}

//	If Supports.GetQueryObjectuiv() is true, calls GetQueryObjectuiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetQueryObjectuiv(id Uint, pname Enum, params *Uint) (err error) {
	if !Supports.GetQueryObjectuiv() {
		err = fmt.Errorf("GetQueryObjectuiv() function call not supported by current GL context.")
	} else {
		GetQueryObjectuiv(id, pname, params)
		err = Util.Error("GetQueryObjectuiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the EndQuery() function.
func (_ GlSupport) EndQuery() bool {
	return C.cancallEndQuery() == 1
}

//	If Supports.EndQuery() is true, calls EndQuery() and yields the err returned by Util.Error(), if any.
func (_ GlTry) EndQuery(target Enum) (err error) {
	if !Supports.EndQuery() {
		err = fmt.Errorf("EndQuery() function call not supported by current GL context.")
	} else {
		EndQuery(target)
		err = Util.Error("EndQuery()")
	}
	return
}

//	Returns whether the current GL context supports calling the ClearDepth() function.
func (_ GlSupport) ClearDepth() bool {
	return C.cancallClearDepth() == 1
}

//	If Supports.ClearDepth() is true, calls ClearDepth() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ClearDepth(depth Double) (err error) {
	if !Supports.ClearDepth() {
		err = fmt.Errorf("ClearDepth() function call not supported by current GL context.")
	} else {
		ClearDepth(depth)
		err = Util.Error("ClearDepth()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetFloati_v() function.
func (_ GlSupport) GetFloati_v() bool {
	return C.cancallGetFloati_v() == 1
}

//	If Supports.GetFloati_v() is true, calls GetFloati_v() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetFloati_v(target Enum, index Uint, data *Float) (err error) {
	if !Supports.GetFloati_v() {
		err = fmt.Errorf("GetFloati_v() function call not supported by current GL context.")
	} else {
		GetFloati_v(target, index, data)
		err = Util.Error("GetFloati_v()")
	}
	return
}

//	Returns whether the current GL context supports calling the PointParameterfv() function.
func (_ GlSupport) PointParameterfv() bool {
	return C.cancallPointParameterfv() == 1
}

//	If Supports.PointParameterfv() is true, calls PointParameterfv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) PointParameterfv(pname Enum, params *Float) (err error) {
	if !Supports.PointParameterfv() {
		err = fmt.Errorf("PointParameterfv() function call not supported by current GL context.")
	} else {
		PointParameterfv(pname, params)
		err = Util.Error("PointParameterfv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetQueryObjectiv() function.
func (_ GlSupport) GetQueryObjectiv() bool {
	return C.cancallGetQueryObjectiv() == 1
}

//	If Supports.GetQueryObjectiv() is true, calls GetQueryObjectiv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetQueryObjectiv(id Uint, pname Enum, params *Int) (err error) {
	if !Supports.GetQueryObjectiv() {
		err = fmt.Errorf("GetQueryObjectiv() function call not supported by current GL context.")
	} else {
		GetQueryObjectiv(id, pname, params)
		err = Util.Error("GetQueryObjectiv()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetActiveUniformName() function.
func (_ GlSupport) GetActiveUniformName() bool {
	return C.cancallGetActiveUniformName() == 1
}

//	If Supports.GetActiveUniformName() is true, calls GetActiveUniformName() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetActiveUniformName(program Uint, uniformIndex Uint, bufSize Sizei, length *Sizei, uniformName *Char) (err error) {
	if !Supports.GetActiveUniformName() {
		err = fmt.Errorf("GetActiveUniformName() function call not supported by current GL context.")
	} else {
		GetActiveUniformName(program, uniformIndex, bufSize, length, uniformName)
		err = Util.Error("GetActiveUniformName()")
	}
	return
}

//	Returns whether the current GL context supports calling the NormalP3ui() function.
func (_ GlSupport) NormalP3ui() bool {
	return C.cancallNormalP3ui() == 1
}

//	If Supports.NormalP3ui() is true, calls NormalP3ui() and yields the err returned by Util.Error(), if any.
func (_ GlTry) NormalP3ui(type_ Enum, coords Uint) (err error) {
	if !Supports.NormalP3ui() {
		err = fmt.Errorf("NormalP3ui() function call not supported by current GL context.")
	} else {
		NormalP3ui(type_, coords)
		err = Util.Error("NormalP3ui()")
	}
	return
}

//	Returns whether the current GL context supports calling the ActiveShaderProgram() function.
func (_ GlSupport) ActiveShaderProgram() bool {
	return C.cancallActiveShaderProgram() == 1
}

//	If Supports.ActiveShaderProgram() is true, calls ActiveShaderProgram() and yields the err returned by Util.Error(), if any.
func (_ GlTry) ActiveShaderProgram(pipeline Uint, program Uint) (err error) {
	if !Supports.ActiveShaderProgram() {
		err = fmt.Errorf("ActiveShaderProgram() function call not supported by current GL context.")
	} else {
		ActiveShaderProgram(pipeline, program)
		err = Util.Error("ActiveShaderProgram()")
	}
	return
}

//	Returns whether the current GL context supports calling the VertexAttribI3iv() function.
func (_ GlSupport) VertexAttribI3iv() bool {
	return C.cancallVertexAttribI3iv() == 1
}

//	If Supports.VertexAttribI3iv() is true, calls VertexAttribI3iv() and yields the err returned by Util.Error(), if any.
func (_ GlTry) VertexAttribI3iv(index Uint, v *Int) (err error) {
	if !Supports.VertexAttribI3iv() {
		err = fmt.Errorf("VertexAttribI3iv() function call not supported by current GL context.")
	} else {
		VertexAttribI3iv(index, v)
		err = Util.Error("VertexAttribI3iv()")
	}
	return
}

//	Returns whether the current GL context supports calling the IsFramebuffer() function.
func (_ GlSupport) IsFramebuffer() bool {
	return C.cancallIsFramebuffer() == 1
}

//	If Supports.IsFramebuffer() is true, calls IsFramebuffer() and yields the err returned by Util.Error(), if any.
func (_ GlTry) IsFramebuffer(framebuffer Uint) (err error, tryRetVal__ Boolean) {
	if !Supports.IsFramebuffer() {
		err = fmt.Errorf("IsFramebuffer() function call not supported by current GL context.")
	} else {
		tryRetVal__ = IsFramebuffer(framebuffer)
		err = Util.Error("IsFramebuffer()")
	}
	return
}

//	Returns whether the current GL context supports calling the BindImageTexture() function.
func (_ GlSupport) BindImageTexture() bool {
	return C.cancallBindImageTexture() == 1
}

//	If Supports.BindImageTexture() is true, calls BindImageTexture() and yields the err returned by Util.Error(), if any.
func (_ GlTry) BindImageTexture(unit Uint, texture Uint, level Int, layered Boolean, layer Int, access Enum, format Enum) (err error) {
	if !Supports.BindImageTexture() {
		err = fmt.Errorf("BindImageTexture() function call not supported by current GL context.")
	} else {
		BindImageTexture(unit, texture, level, layered, layer, access, format)
		err = Util.Error("BindImageTexture()")
	}
	return
}

//	Returns whether the current GL context supports calling the GetActiveUniform() function.
func (_ GlSupport) GetActiveUniform() bool {
	return C.cancallGetActiveUniform() == 1
}

//	If Supports.GetActiveUniform() is true, calls GetActiveUniform() and yields the err returned by Util.Error(), if any.
func (_ GlTry) GetActiveUniform(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char) (err error) {
	if !Supports.GetActiveUniform() {
		err = fmt.Errorf("GetActiveUniform() function call not supported by current GL context.")
	} else {
		GetActiveUniform(program, index, bufSize, length, size, type_, name)
		err = Util.Error("GetActiveUniform()")
	}
	return
}

//	Returns whether the current GL context supports calling the Uniform1i() function.
func (_ GlSupport) Uniform1i() bool {
	return C.cancallUniform1i() == 1
}

//	If Supports.Uniform1i() is true, calls Uniform1i() and yields the err returned by Util.Error(), if any.
func (_ GlTry) Uniform1i(location Int, v0 Int) (err error) {
	if !Supports.Uniform1i() {
		err = fmt.Errorf("Uniform1i() function call not supported by current GL context.")
	} else {
		Uniform1i(location, v0)
		err = Util.Error("Uniform1i()")
	}
	return
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1f.xml
func ProgramUniform1f(program Uint, location Int, v0 Float) {
	C.callProgramUniform1f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameterfv.xml
func GetSamplerParameterfv(sampler Uint, pname Enum, params *Float) {
	C.callGetSamplerParameterfv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP4ui.xml
func TexCoordP4ui(type_ Enum, coords Uint) {
	C.callTexCoordP4ui((C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1uiv.xml
func ProgramUniform1uiv(program Uint, location Int, count Sizei, value *Uint) {
	C.callProgramUniform1uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteSamplers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteSamplers.xml
func DeleteSamplers(count Sizei, samplers *Uint) {
	C.callDeleteSamplers((C.GLsizei)(count), (*C.GLuint)(samplers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsShader.xml
func IsShader(shader Uint) Boolean {
	return (Boolean)(C.callIsShader((C.GLuint)(shader)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexSubImage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage3D.xml
func CopyTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, x Int, y Int, width Sizei, height Sizei) {
	C.callCopyTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformuiv.xml
func GetUniformuiv(program Uint, location Int, params *Uint) {
	C.callGetUniformuiv((C.GLuint)(program), (C.GLint)(location), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL2d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL2d.xml
func VertexAttribL2d(index Uint, x Double, y Double) {
	C.callVertexAttribL2d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSecondaryColorP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColorP3ui.xml
func SecondaryColorP3ui(type_ Enum, color Uint) {
	C.callSecondaryColorP3ui((C.GLenum)(type_), (C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsSampler
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsSampler.xml
func IsSampler(sampler Uint) Boolean {
	return (Boolean)(C.callIsSampler((C.GLuint)(sampler)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2dv.xml
func Uniform2dv(location Int, count Sizei, value *Double) {
	C.callUniform2dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthMask
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthMask.xml
func DepthMask(flag Boolean) {
	C.callDepthMask((C.GLboolean)(flag))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPushDebugGroup
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPushDebugGroup.xml
func PushDebugGroup(source Enum, id Uint, length Sizei, message *Char) {
	C.callPushDebugGroup((C.GLenum)(source), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(message))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1dv.xml
func ProgramUniform1dv(program Uint, location Int, count Sizei, value *Double) {
	C.callProgramUniform1dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawArraysIndirect
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawArraysIndirect.xml
func MultiDrawArraysIndirect(mode Enum, indirect Ptr, drawcount Sizei, stride Sizei) {
	C.callMultiDrawArraysIndirect((C.GLenum)(mode), (unsafe.Pointer)(indirect), (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP2uiv.xml
func MultiTexCoordP2uiv(texture Enum, type_ Enum, coords *Uint) {
	C.callMultiTexCoordP2uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateSubFramebuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateSubFramebuffer.xml
func InvalidateSubFramebuffer(target Enum, numAttachments Sizei, attachments *Enum, x Int, y Int, width Sizei, height Sizei) {
	C.callInvalidateSubFramebuffer((C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(attachments), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3i.xml
func Uniform3i(location Int, v0 Int, v1 Int, v2 Int) {
	C.callUniform3i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2i.xml
func ProgramUniform2i(program Uint, location Int, v0 Int, v1 Int) {
	C.callProgramUniform2i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glShaderSource
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glShaderSource.xml
func ShaderSource(shader Uint, count Sizei, string_ **Char, length *Int) {
	C.callShaderSource((C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(string_)), (*C.GLint)(length))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP1uiv.xml
func TexCoordP1uiv(type_ Enum, coords *Uint) {
	C.callTexCoordP1uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformBlockBinding
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformBlockBinding.xml
func UniformBlockBinding(program Uint, uniformBlockIndex Uint, uniformBlockBinding Uint) {
	C.callUniformBlockBinding((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLuint)(uniformBlockBinding))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4fv.xml
func Uniform4fv(location Int, count Sizei, value *Float) {
	C.callUniform4fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribdv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribdv.xml
func GetVertexAttribdv(index Uint, pname Enum, params *Double) {
	C.callGetVertexAttribdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenBuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenBuffers.xml
func GenBuffers(n Sizei, buffers *Uint) {
	C.callGenBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP4ui.xml
func VertexAttribP4ui(index Uint, type_ Enum, normalized Boolean, value Uint) {
	C.callVertexAttribP4ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInternalformativ
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetInternalformativ.xml
func GetInternalformativ(target Enum, internalformat Enum, pname Enum, bufSize Sizei, params *Int) {
	C.callGetInternalformativ((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x2dv.xml
func ProgramUniformMatrix4x2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix4x2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompileShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompileShader.xml
func CompileShader(shader Uint) {
	C.callCompileShader((C.GLuint)(shader))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindRenderbuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindRenderbuffer.xml
func BindRenderbuffer(target Enum, renderbuffer Uint) {
	C.callBindRenderbuffer((C.GLenum)(target), (C.GLuint)(renderbuffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexSubImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage1D.xml
func CompressedTexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, imageSize Sizei, data Ptr) {
	C.callCompressedTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFunc
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendFunc.xml
func BlendFunc(sfactor Enum, dfactor Enum) {
	C.callBlendFunc((C.GLenum)(sfactor), (C.GLenum)(dfactor))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterfv.xml
func TexParameterfv(target Enum, pname Enum, params *Float) {
	C.callTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateFramebuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateFramebuffer.xml
func InvalidateFramebuffer(target Enum, numAttachments Sizei, attachments *Enum) {
	C.callInvalidateFramebuffer((C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(attachments))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindTransformFeedback.xml
func BindTransformFeedback(target Enum, id Uint) {
	C.callBindTransformFeedback((C.GLenum)(target), (C.GLuint)(id))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFragDataIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFragDataIndex.xml
func GetFragDataIndex(program Uint, name *Char) Int {
	return (Int)(C.callGetFragDataIndex((C.GLuint)(program), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTransformFeedbackVaryings
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTransformFeedbackVaryings.xml
func TransformFeedbackVaryings(program Uint, count Sizei, varyings **Char, bufferMode Enum) {
	C.callTransformFeedbackVaryings((C.GLuint)(program), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(varyings)), (C.GLenum)(bufferMode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glRenderbufferStorage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glRenderbufferStorage.xml
func RenderbufferStorage(target Enum, internalformat Enum, width Sizei, height Sizei) {
	C.callRenderbufferStorage((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformBlockIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformBlockIndex.xml
func GetUniformBlockIndex(program Uint, uniformBlockName *Char) Uint {
	return (Uint)(C.callGetUniformBlockIndex((C.GLuint)(program), (*C.GLchar)(uniformBlockName)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glObjectLabel
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glObjectLabel.xml
func ObjectLabel(identifier Enum, name Uint, length Sizei, label *Char) {
	C.callObjectLabel((C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(length), (*C.GLchar)(label))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexLevelParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexLevelParameterfv.xml
func GetTexLevelParameterfv(target Enum, level Int, pname Enum, params *Float) {
	C.callGetTexLevelParameterfv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x3fv.xml
func ProgramUniformMatrix4x3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix4x3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderiv.xml
func GetShaderiv(shader Uint, pname Enum, params *Int) {
	C.callGetShaderiv((C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4ui.xml
func ProgramUniform4ui(program Uint, location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint) {
	C.callProgramUniform4ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP3uiv.xml
func MultiTexCoordP3uiv(texture Enum, type_ Enum, coords *Uint) {
	C.callMultiTexCoordP3uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3uiv.xml
func ProgramUniform3uiv(program Uint, location Int, count Sizei, value *Uint) {
	C.callProgramUniform3uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetPointerv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetPointerv.xml
func GetPointerv(pname Enum, params *Ptr) {
	C.callGetPointerv((C.GLenum)(pname), (*unsafe.Pointer)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP4uiv.xml
func MultiTexCoordP4uiv(texture Enum, type_ Enum, coords *Uint) {
	C.callMultiTexCoordP4uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x2fv.xml
func UniformMatrix4x2fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix4x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glShaderStorageBlockBinding
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glShaderStorageBlockBinding.xml
func ShaderStorageBlockBinding(program Uint, storageBlockIndex Uint, storageBlockBinding Uint) {
	C.callShaderStorageBlockBinding((C.GLuint)(program), (C.GLuint)(storageBlockIndex), (C.GLuint)(storageBlockBinding))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glShaderBinary
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glShaderBinary.xml
func ShaderBinary(count Sizei, shaders *Uint, binaryformat Enum, binary Ptr, length Sizei) {
	C.callShaderBinary((C.GLsizei)(count), (*C.GLuint)(shaders), (C.GLenum)(binaryformat), (unsafe.Pointer)(binary), (C.GLsizei)(length))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1d.xml
func Uniform1d(location Int, x Double) {
	C.callUniform1d((C.GLint)(location), (C.GLdouble)(x))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquationSeparatei
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquationSeparatei.xml
func BlendEquationSeparatei(buf Uint, modeRGB Enum, modeAlpha Enum) {
	C.callBlendEquationSeparatei((C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenQueries
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenQueries.xml
func GenQueries(n Sizei, ids *Uint) {
	C.callGenQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClampColor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClampColor.xml
func ClampColor(target Enum, clamp Enum) {
	C.callClampColor((C.GLenum)(target), (C.GLenum)(clamp))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPauseTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPauseTransformFeedback.xml
func PauseTransformFeedback() {
	C.callPauseTransformFeedback()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glLinkProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glLinkProgram.xml
func LinkProgram(program Uint) {
	C.callLinkProgram((C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCheckFramebufferStatus
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCheckFramebufferStatus.xml
func CheckFramebufferStatus(target Enum) Enum {
	return (Enum)(C.callCheckFramebufferStatus((C.GLenum)(target)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPointParameteri.xml
func PointParameteri(pname Enum, param Int) {
	C.callPointParameteri((C.GLenum)(pname), (C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4usv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4usv.xml
func VertexAttribI4usv(index Uint, v *Ushort) {
	C.callVertexAttribI4usv((C.GLuint)(index), (*C.GLushort)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPointParameteriv.xml
func PointParameteriv(pname Enum, params *Int) {
	C.callPointParameteriv((C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribIPointer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribIPointer.xml
func VertexAttribIPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Ptr) {
	C.callVertexAttribIPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenTransformFeedbacks
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenTransformFeedbacks.xml
func GenTransformFeedbacks(n Sizei, ids *Uint) {
	C.callGenTransformFeedbacks((C.GLsizei)(n), (*C.GLuint)(ids))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1uiv.xml
func Uniform1uiv(location Int, count Sizei, value *Uint) {
	C.callUniform1uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenerateMipmap
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenerateMipmap.xml
func GenerateMipmap(target Enum) {
	C.callGenerateMipmap((C.GLenum)(target))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterfv.xml
func GetTexParameterfv(target Enum, pname Enum, params *Float) {
	C.callGetTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2uiv.xml
func Uniform2uiv(location Int, count Sizei, value *Uint) {
	C.callUniform2uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformSubroutinesuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformSubroutinesuiv.xml
func UniformSubroutinesuiv(shadertype Enum, count Sizei, indices *Uint) {
	C.callUniformSubroutinesuiv((C.GLenum)(shadertype), (C.GLsizei)(count), (*C.GLuint)(indices))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearDepthf
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearDepthf.xml
func ClearDepthf(d Float) {
	C.callClearDepthf((C.GLfloat)(d))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteVertexArrays
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteVertexArrays.xml
func DeleteVertexArrays(n Sizei, arrays *Uint) {
	C.callDeleteVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsEnabled
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsEnabled.xml
func IsEnabled(cap Enum) Boolean {
	return (Boolean)(C.callIsEnabled((C.GLenum)(cap)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4ui.xml
func Uniform4ui(location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint) {
	C.callUniform4ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFramebufferAttachmentParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFramebufferAttachmentParameteriv.xml
func GetFramebufferAttachmentParameteriv(target Enum, attachment Enum, pname Enum, params *Int) {
	C.callGetFramebufferAttachmentParameteriv((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1dv.xml
func Uniform1dv(location Int, count Sizei, value *Double) {
	C.callUniform1dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorMask
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorMask.xml
func ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean) {
	C.callColorMask((C.GLboolean)(red), (C.GLboolean)(green), (C.GLboolean)(blue), (C.GLboolean)(alpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetString
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetString.xml
func GetString(name Enum) *Ubyte {
	return (*Ubyte)(C.callGetString((C.GLenum)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glHint
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glHint.xml
func Hint(target Enum, mode Enum) {
	C.callHint((C.GLenum)(target), (C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferData.xml
func ClearBufferData(target Enum, internalformat Enum, format Enum, type_ Enum, data Ptr) {
	C.callClearBufferData((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2d.xml
func Uniform2d(location Int, x Double, y Double) {
	C.callUniform2d((C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInteger64v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetInteger64v.xml
func GetInteger64v(pname Enum, params *Int64) {
	C.callGetInteger64v((C.GLenum)(pname), (*C.GLint64)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1f.xml
func Uniform1f(location Int, v0 Float) {
	C.callUniform1f((C.GLint)(location), (C.GLfloat)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexBuffer.xml
func TexBuffer(target Enum, internalformat Enum, buffer Uint) {
	C.callTexBuffer((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndConditionalRender
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEndConditionalRender.xml
func EndConditionalRender() {
	C.callEndConditionalRender()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP3ui.xml
func VertexP3ui(type_ Enum, value Uint) {
	C.callVertexP3ui((C.GLenum)(type_), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramInfoLog
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramInfoLog.xml
func GetProgramInfoLog(program Uint, bufSize Sizei, length *Sizei, infoLog *Char) {
	C.callGetProgramInfoLog((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2i.xml
func VertexAttribI2i(index Uint, x Int, y Int) {
	C.callVertexAttribI2i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsTransformFeedback.xml
func IsTransformFeedback(id Uint) Boolean {
	return (Boolean)(C.callIsTransformFeedback((C.GLuint)(id)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glLogicOp
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glLogicOp.xml
func LogicOp(opcode Enum) {
	C.callLogicOp((C.GLenum)(opcode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDispatchComputeIndirect
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDispatchComputeIndirect.xml
func DispatchComputeIndirect(indirect Intptr) {
	C.callDispatchComputeIndirect((C.GLintptr)(indirect))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindTexture
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindTexture.xml
func BindTexture(target Enum, texture Uint) {
	C.callBindTexture((C.GLenum)(target), (C.GLuint)(texture))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDisableVertexAttribArray
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDisableVertexAttribArray.xml
func DisableVertexAttribArray(index Uint) {
	C.callDisableVertexAttribArray((C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glQueryCounter
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glQueryCounter.xml
func QueryCounter(id Uint, target Enum) {
	C.callQueryCounter((C.GLuint)(id), (C.GLenum)(target))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3f.xml
func ProgramUniform3f(program Uint, location Int, v0 Float, v1 Float, v2 Float) {
	C.callProgramUniform3f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsEnabledi
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsEnabledi.xml
func IsEnabledi(target Enum, index Uint) Boolean {
	return (Boolean)(C.callIsEnabledi((C.GLenum)(target), (C.GLuint)(index)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL3d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL3d.xml
func VertexAttribL3d(index Uint, x Double, y Double, z Double) {
	C.callVertexAttribL3d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4dv.xml
func Uniform4dv(location Int, count Sizei, value *Double) {
	C.callUniform4dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindVertexArray
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindVertexArray.xml
func BindVertexArray(array Uint) {
	C.callBindVertexArray((C.GLuint)(array))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3dv.xml
func ProgramUniform3dv(program Uint, location Int, count Sizei, value *Double) {
	C.callProgramUniform3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFramebufferParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFramebufferParameteriv.xml
func GetFramebufferParameteriv(target Enum, pname Enum, params *Int) {
	C.callGetFramebufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCullFace
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCullFace.xml
func CullFace(mode Enum) {
	C.callCullFace((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawArrays
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawArrays.xml
func MultiDrawArrays(mode Enum, first *Int, count *Sizei, drawcount Sizei) {
	C.callMultiDrawArrays((C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(drawcount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x3dv.xml
func ProgramUniformMatrix2x3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix2x3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glValidateProgramPipeline
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glValidateProgramPipeline.xml
func ValidateProgramPipeline(pipeline Uint) {
	C.callValidateProgramPipeline((C.GLuint)(pipeline))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x2fv.xml
func ProgramUniformMatrix3x2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix3x2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3i.xml
func VertexAttribI3i(index Uint, x Int, y Int, z Int) {
	C.callVertexAttribI3i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceiv.xml
func GetProgramResourceiv(program Uint, programInterface Enum, index Uint, propCount Sizei, props *Enum, bufSize Sizei, length *Sizei, params *Int) {
	C.callGetProgramResourceiv((C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(propCount), (*C.GLenum)(props), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformSubroutineuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformSubroutineuiv.xml
func GetUniformSubroutineuiv(shadertype Enum, location Int, params *Uint) {
	C.callGetUniformSubroutineuiv((C.GLenum)(shadertype), (C.GLint)(location), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2iv.xml
func ProgramUniform2iv(program Uint, location Int, count Sizei, value *Int) {
	C.callProgramUniform2iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenFramebuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenFramebuffers.xml
func GenFramebuffers(n Sizei, framebuffers *Uint) {
	C.callGenFramebuffers((C.GLsizei)(n), (*C.GLuint)(framebuffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedbackInstanced
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedbackInstanced.xml
func DrawTransformFeedbackInstanced(mode Enum, id Uint, instancecount Sizei) {
	C.callDrawTransformFeedbackInstanced((C.GLenum)(mode), (C.GLuint)(id), (C.GLsizei)(instancecount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsBuffer.xml
func IsBuffer(buffer Uint) Boolean {
	return (Boolean)(C.callIsBuffer((C.GLuint)(buffer)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glAttachShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glAttachShader.xml
func AttachShader(program Uint, shader Uint) {
	C.callAttachShader((C.GLuint)(program), (C.GLuint)(shader))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsIndirect
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsIndirect.xml
func DrawElementsIndirect(mode Enum, type_ Enum, indirect Ptr) {
	C.callDrawElementsIndirect((C.GLenum)(mode), (C.GLenum)(type_), (unsafe.Pointer)(indirect))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1ui.xml
func ProgramUniform1ui(program Uint, location Int, v0 Uint) {
	C.callProgramUniform1ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformIndices
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformIndices.xml
func GetUniformIndices(program Uint, uniformCount Sizei, uniformNames **Char, uniformIndices *Uint) {
	C.callGetUniformIndices((C.GLuint)(program), (C.GLsizei)(uniformCount), (**C.GLchar)(unsafe.Pointer(uniformNames)), (*C.GLuint)(uniformIndices))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3f.xml
func Uniform3f(location Int, v0 Float, v1 Float, v2 Float) {
	C.callUniform3f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindFragDataLocationIndexed
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindFragDataLocationIndexed.xml
func BindFragDataLocationIndexed(program Uint, colorNumber Uint, index Uint, name *Char) {
	C.callBindFragDataLocationIndexed((C.GLuint)(program), (C.GLuint)(colorNumber), (C.GLuint)(index), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTextureView
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTextureView.xml
func TextureView(texture Uint, target Enum, origtexture Uint, internalformat Enum, minlevel Uint, numlevels Uint, minlayer Uint, numlayers Uint) {
	C.callTextureView((C.GLuint)(texture), (C.GLenum)(target), (C.GLuint)(origtexture), (C.GLenum)(internalformat), (C.GLuint)(minlevel), (C.GLuint)(numlevels), (C.GLuint)(minlayer), (C.GLuint)(numlayers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryiv.xml
func GetQueryiv(target Enum, pname Enum, params *Int) {
	C.callGetQueryiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture1D.xml
func FramebufferTexture1D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int) {
	C.callFramebufferTexture1D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointSize
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPointSize.xml
func PointSize(size Float) {
	C.callPointSize((C.GLfloat)(size))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4i.xml
func VertexAttribI4i(index Uint, x Int, y Int, z Int, w Int) {
	C.callVertexAttribI4i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPopDebugGroup
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPopDebugGroup.xml
func PopDebugGroup() {
	C.callPopDebugGroup()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilFunc
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilFunc.xml
func StencilFunc(func_ Enum, ref Int, mask Uint) {
	C.callStencilFunc((C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage3DMultisample
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexImage3DMultisample.xml
func TexImage3DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean) {
	C.callTexImage3DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(fixedsamplelocations))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x3fv.xml
func ProgramUniformMatrix2x3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix2x3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL1d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL1d.xml
func VertexAttribL1d(index Uint, x Double) {
	C.callVertexAttribL1d((C.GLuint)(index), (C.GLdouble)(x))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferiv.xml
func ClearBufferiv(buffer Enum, drawbuffer Int, value *Int) {
	C.callClearBufferiv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindVertexBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindVertexBuffer.xml
func BindVertexBuffer(bindingindex Uint, buffer Uint, offset Intptr, stride Sizei) {
	C.callBindVertexBuffer((C.GLuint)(bindingindex), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(stride))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDisablei
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDisablei.xml
func Disablei(target Enum, index Uint) {
	C.callDisablei((C.GLenum)(target), (C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP1uiv.xml
func VertexAttribP1uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) {
	C.callVertexAttribP1uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3dv.xml
func Uniform3dv(location Int, count Sizei, value *Double) {
	C.callUniform3dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTextureLayer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTextureLayer.xml
func FramebufferTextureLayer(target Enum, attachment Enum, texture Uint, level Int, layer Int) {
	C.callFramebufferTextureLayer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture2D.xml
func FramebufferTexture2D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int) {
	C.callFramebufferTexture2D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glReleaseShaderCompiler
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glReleaseShaderCompiler.xml
func ReleaseShaderCompiler() {
	C.callReleaseShaderCompiler()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x4fv.xml
func ProgramUniformMatrix3x4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix3x4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x4dv.xml
func ProgramUniformMatrix2x4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix2x4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPatchParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPatchParameteri.xml
func PatchParameteri(pname Enum, value Int) {
	C.callPatchParameteri((C.GLenum)(pname), (C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramPipelineInfoLog
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramPipelineInfoLog.xml
func GetProgramPipelineInfoLog(pipeline Uint, bufSize Sizei, length *Sizei, infoLog *Char) {
	C.callGetProgramPipelineInfoLog((C.GLuint)(pipeline), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveSubroutineName
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveSubroutineName.xml
func GetActiveSubroutineName(program Uint, shadertype Enum, index Uint, bufsize Sizei, length *Sizei, name *Char) {
	C.callGetActiveSubroutineName((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(length), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameterIuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterIuiv.xml
func GetTexParameterIuiv(target Enum, pname Enum, params *Uint) {
	C.callGetTexParameterIuiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1fv.xml
func ProgramUniform1fv(program Uint, location Int, count Sizei, value *Float) {
	C.callProgramUniform1fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP3uiv.xml
func TexCoordP3uiv(type_ Enum, coords *Uint) {
	C.callTexCoordP3uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenProgramPipelines
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenProgramPipelines.xml
func GenProgramPipelines(n Sizei, pipelines *Uint) {
	C.callGenProgramPipelines((C.GLsizei)(n), (*C.GLuint)(pipelines))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetDebugMessageLog
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetDebugMessageLog.xml
func GetDebugMessageLog(count Uint, bufsize Sizei, sources *Enum, types *Enum, ids *Uint, severities *Enum, lengths *Sizei, messageLog *Char) Uint {
	return (Uint)(C.callGetDebugMessageLog((C.GLuint)(count), (C.GLsizei)(bufsize), (*C.GLenum)(sources), (*C.GLenum)(types), (*C.GLuint)(ids), (*C.GLenum)(severities), (*C.GLsizei)(lengths), (*C.GLchar)(messageLog)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameteriv.xml
func SamplerParameteriv(sampler Uint, pname Enum, param *Int) {
	C.callSamplerParameteriv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFrontFace
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFrontFace.xml
func FrontFace(mode Enum) {
	C.callFrontFace((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteBuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteBuffers.xml
func DeleteBuffers(n Sizei, buffers *Uint) {
	C.callDeleteBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexSubImage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage3D.xml
func CompressedTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, imageSize Sizei, data Ptr) {
	C.callCompressedTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexSubImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage2D.xml
func CompressedTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, imageSize Sizei, data Ptr) {
	C.callCompressedTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindProgramPipeline
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindProgramPipeline.xml
func BindProgramPipeline(pipeline Uint) {
	C.callBindProgramPipeline((C.GLuint)(pipeline))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1iv.xml
func Uniform1iv(location Int, count Sizei, value *Int) {
	C.callUniform1iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribDivisor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribDivisor.xml
func VertexAttribDivisor(index Uint, divisor Uint) {
	C.callVertexAttribDivisor((C.GLuint)(index), (C.GLuint)(divisor))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindBufferBase
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindBufferBase.xml
func BindBufferBase(target Enum, index Uint, buffer Uint) {
	C.callBindBufferBase((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorP4ui.xml
func ColorP4ui(type_ Enum, color Uint) {
	C.callColorP4ui((C.GLenum)(type_), (C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceIndex.xml
func GetProgramResourceIndex(program Uint, programInterface Enum, name *Char) Uint {
	return (Uint)(C.callGetProgramResourceIndex((C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferParameteri.xml
func FramebufferParameteri(target Enum, pname Enum, param Int) {
	C.callFramebufferParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glValidateProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glValidateProgram.xml
func ValidateProgram(program Uint) {
	C.callValidateProgram((C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawRangeElements
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawRangeElements.xml
func DrawRangeElements(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Ptr) {
	C.callDrawRangeElements((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameteri.xml
func SamplerParameteri(sampler Uint, pname Enum, param Int) {
	C.callSamplerParameteri((C.GLuint)(sampler), (C.GLenum)(pname), (C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveAttrib
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveAttrib.xml
func GetActiveAttrib(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char) {
	C.callGetActiveAttrib((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSecondaryColorP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColorP3uiv.xml
func SecondaryColorP3uiv(type_ Enum, color *Uint) {
	C.callSecondaryColorP3uiv((C.GLenum)(type_), (*C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterfv.xml
func SamplerParameterfv(sampler Uint, pname Enum, param *Float) {
	C.callSamplerParameterfv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRange
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthRange.xml
func DepthRange(near_ Double, far_ Double) {
	C.callDepthRange((C.GLdouble)(near_), (C.GLdouble)(far_))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3uiv.xml
func Uniform3uiv(location Int, count Sizei, value *Uint) {
	C.callUniform3uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterf
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterf.xml
func SamplerParameterf(sampler Uint, pname Enum, param Float) {
	C.callSamplerParameterf((C.GLuint)(sampler), (C.GLenum)(pname), (C.GLfloat)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2f.xml
func Uniform2f(location Int, v0 Float, v1 Float) {
	C.callUniform2f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterIuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterIuiv.xml
func SamplerParameterIuiv(sampler Uint, pname Enum, param *Uint) {
	C.callSamplerParameterIuiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArraysIndirect
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawArraysIndirect.xml
func DrawArraysIndirect(mode Enum, indirect Ptr) {
	C.callDrawArraysIndirect((C.GLenum)(mode), (unsafe.Pointer)(indirect))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenTextures
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenTextures.xml
func GenTextures(n Sizei, textures *Uint) {
	C.callGenTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindBufferRange
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindBufferRange.xml
func BindBufferRange(target Enum, index Uint, buffer Uint, offset Intptr, size Sizeiptr) {
	C.callBindBufferRange((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x2dv.xml
func UniformMatrix4x2dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix4x2dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewport
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glViewport.xml
func Viewport(x Int, y Int, width Sizei, height Sizei) {
	C.callViewport((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP1uiv.xml
func MultiTexCoordP1uiv(texture Enum, type_ Enum, coords *Uint) {
	C.callMultiTexCoordP1uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2fv.xml
func UniformMatrix2fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2f.xml
func ProgramUniform2f(program Uint, location Int, v0 Float, v1 Float) {
	C.callProgramUniform2f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformLocation.xml
func GetUniformLocation(program Uint, name *Char) Int {
	return (Int)(C.callGetUniformLocation((C.GLuint)(program), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage2DMultisample
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage2DMultisample.xml
func TexStorage2DMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei, fixedsamplelocations Boolean) {
	C.callTexStorage2DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(fixedsamplelocations))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFloatv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFloatv.xml
func GetFloatv(pname Enum, params *Float) {
	C.callGetFloatv((C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x2fv.xml
func UniformMatrix3x2fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix3x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindSampler
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindSampler.xml
func BindSampler(unit Uint, sampler Uint) {
	C.callBindSampler((C.GLuint)(unit), (C.GLuint)(sampler))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4iv.xml
func Uniform4iv(location Int, count Sizei, value *Int) {
	C.callUniform4iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMemoryBarrier
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMemoryBarrier.xml
func MemoryBarrier(barriers Bitfield) {
	C.callMemoryBarrier((C.GLbitfield)(barriers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribPointer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribPointer.xml
func VertexAttribPointer(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, pointer Ptr) {
	C.callVertexAttribPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBufferData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBufferData.xml
func BufferData(target Enum, size Sizeiptr, data Ptr, usage Enum) {
	C.callBufferData((C.GLenum)(target), (C.GLsizeiptr)(size), (unsafe.Pointer)(data), (C.GLenum)(usage))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2dv.xml
func ProgramUniform2dv(program Uint, location Int, count Sizei, value *Double) {
	C.callProgramUniform2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glLineWidth
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glLineWidth.xml
func LineWidth(width Float) {
	C.callLineWidth((C.GLfloat)(width))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorP4uiv.xml
func ColorP4uiv(type_ Enum, color *Uint) {
	C.callColorP4uiv((C.GLenum)(type_), (*C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3i.xml
func ProgramUniform3i(program Uint, location Int, v0 Int, v1 Int, v2 Int) {
	C.callProgramUniform3i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElements
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElements.xml
func DrawElements(mode Enum, count Sizei, type_ Enum, indices Ptr) {
	C.callDrawElements((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceLocation.xml
func GetProgramResourceLocation(program Uint, programInterface Enum, name *Char) Int {
	return (Int)(C.callGetProgramResourceLocation((C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateBufferSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateBufferSubData.xml
func InvalidateBufferSubData(buffer Uint, offset Intptr, length Sizeiptr) {
	C.callInvalidateBufferSubData((C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage2D.xml
func TexStorage2D(target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei) {
	C.callTexStorage2D((C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteQueries
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteQueries.xml
func DeleteQueries(n Sizei, ids *Uint) {
	C.callDeleteQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexImage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexImage.xml
func GetTexImage(target Enum, level Int, format Enum, type_ Enum, pixels Ptr) {
	C.callGetTexImage((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClientWaitSync
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClientWaitSync.xml
func ClientWaitSync(sync Sync, flags Bitfield, timeout Uint64) Enum {
	return (Enum)(C.callClientWaitSync((C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstancedBaseVertex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstancedBaseVertex.xml
func DrawElementsInstancedBaseVertex(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, basevertex Int) {
	C.callDrawElementsInstancedBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount), (C.GLint)(basevertex))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPixelStoref
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPixelStoref.xml
func PixelStoref(pname Enum, param Float) {
	C.callPixelStoref((C.GLenum)(pname), (C.GLfloat)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameteriv.xml
func TexParameteriv(target Enum, pname Enum, params *Int) {
	C.callTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferRenderbuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferRenderbuffer.xml
func FramebufferRenderbuffer(target Enum, attachment Enum, renderbuffertarget Enum, renderbuffer Uint) {
	C.callFramebufferRenderbuffer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferSubData.xml
func ClearBufferSubData(target Enum, internalformat Enum, offset Intptr, size Sizeiptr, format Enum, type_ Enum, data Ptr) {
	C.callClearBufferSubData((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2d.xml
func ProgramUniform2d(program Uint, location Int, v0 Double, v1 Double) {
	C.callProgramUniform2d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSampleCoverage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSampleCoverage.xml
func SampleCoverage(value Float, invert Boolean) {
	C.callSampleCoverage((C.GLfloat)(value), (C.GLboolean)(invert))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsVertexArray
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsVertexArray.xml
func IsVertexArray(array Uint) Boolean {
	return (Boolean)(C.callIsVertexArray((C.GLuint)(array)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDispatchCompute
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDispatchCompute.xml
func DispatchCompute(num_groups_x Uint, num_groups_y Uint, num_groups_z Uint) {
	C.callDispatchCompute((C.GLuint)(num_groups_x), (C.GLuint)(num_groups_y), (C.GLuint)(num_groups_z))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFenceSync
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFenceSync.xml
func FenceSync(condition Enum, flags Bitfield) Sync {
	return (Sync)(C.callFenceSync((C.GLenum)(condition), (C.GLbitfield)(flags)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribiv.xml
func GetVertexAttribiv(index Uint, pname Enum, params *Int) {
	C.callGetVertexAttribiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFlushMappedBufferRange
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFlushMappedBufferRange.xml
func FlushMappedBufferRange(target Enum, offset Intptr, length Sizeiptr) {
	C.callFlushMappedBufferRange((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribfv.xml
func GetVertexAttribfv(index Uint, pname Enum, params *Float) {
	C.callGetVertexAttribfv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCreateShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCreateShader.xml
func CreateShader(type_ Enum) Uint {
	return (Uint)(C.callCreateShader((C.GLenum)(type_)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribFormat
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribFormat.xml
func VertexAttribFormat(attribindex Uint, size Int, type_ Enum, normalized Boolean, relativeoffset Uint) {
	C.callVertexAttribFormat((C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(relativeoffset))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawElements
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawElements.xml
func MultiDrawElements(mode Enum, count *Sizei, type_ Enum, indices *Ptr, drawcount Sizei) {
	C.callMultiDrawElements((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(drawcount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyImageSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyImageSubData.xml
func CopyImageSubData(srcName Uint, srcTarget Enum, srcLevel Int, srcX Int, srcY Int, srcZ Int, dstName Uint, dstTarget Enum, dstLevel Int, dstX Int, dstY Int, dstZ Int, srcWidth Sizei, srcHeight Sizei, srcDepth Sizei) {
	C.callCopyImageSubData((C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(srcWidth), (C.GLsizei)(srcHeight), (C.GLsizei)(srcDepth))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsBaseVertex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsBaseVertex.xml
func DrawElementsBaseVertex(mode Enum, count Sizei, type_ Enum, indices Ptr, basevertex Int) {
	C.callDrawElementsBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLint)(basevertex))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilMask
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilMask.xml
func StencilMask(mask Uint) {
	C.callStencilMask((C.GLuint)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL4dv.xml
func VertexAttribL4dv(index Uint, v *Double) {
	C.callVertexAttribL4dv((C.GLuint)(index), (*C.GLdouble)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameteri.xml
func TexParameteri(target Enum, pname Enum, param Int) {
	C.callTexParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP1ui.xml
func TexCoordP1ui(type_ Enum, coords Uint) {
	C.callTexCoordP1ui((C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL4d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL4d.xml
func VertexAttribL4d(index Uint, x Double, y Double, z Double, w Double) {
	C.callVertexAttribL4d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP2ui.xml
func TexCoordP2ui(type_ Enum, coords Uint) {
	C.callTexCoordP2ui((C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformfv.xml
func GetUniformfv(program Uint, location Int, params *Float) {
	C.callGetUniformfv((C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2iv.xml
func VertexAttribI2iv(index Uint, v *Int) {
	C.callVertexAttribI2iv((C.GLuint)(index), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBooleani_v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBooleani_v.xml
func GetBooleani_v(target Enum, index Uint, data *Boolean) {
	C.callGetBooleani_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLboolean)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetError
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetError.xml
func GetError() Enum {
	return (Enum)(C.callGetError())
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x4fv.xml
func UniformMatrix2x4fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix2x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetObjectLabel
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetObjectLabel.xml
func GetObjectLabel(identifier Enum, name Uint, bufSize Sizei, length *Sizei, label *Char) {
	C.callGetObjectLabel((C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(label))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexImage2D.xml
func TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, type_ Enum, pixels Ptr) {
	C.callTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3fv.xml
func ProgramUniformMatrix3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP3ui.xml
func TexCoordP3ui(type_ Enum, coords Uint) {
	C.callTexCoordP3ui((C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetDoublei_v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetDoublei_v.xml
func GetDoublei_v(target Enum, index Uint, data *Double) {
	C.callGetDoublei_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLdouble)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexImage3D.xml
func TexImage3D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, depth Sizei, border Int, format Enum, type_ Enum, pixels Ptr) {
	C.callTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenSamplers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenSamplers.xml
func GenSamplers(count Sizei, samplers *Uint) {
	C.callGenSamplers((C.GLsizei)(count), (*C.GLuint)(samplers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4uiv.xml
func VertexAttribI4uiv(index Uint, v *Uint) {
	C.callVertexAttribI4uiv((C.GLuint)(index), (*C.GLuint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformdv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformdv.xml
func GetUniformdv(program Uint, location Int, params *Double) {
	C.callGetUniformdv((C.GLuint)(program), (C.GLint)(location), (*C.GLdouble)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4iv.xml
func ProgramUniform4iv(program Uint, location Int, count Sizei, value *Int) {
	C.callProgramUniform4iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1i.xml
func ProgramUniform1i(program Uint, location Int, v0 Int) {
	C.callProgramUniform1i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP1ui.xml
func MultiTexCoordP1ui(texture Enum, type_ Enum, coords Uint) {
	C.callMultiTexCoordP1ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4dv.xml
func ProgramUniformMatrix4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1fv.xml
func Uniform1fv(location Int, count Sizei, value *Float) {
	C.callUniform1fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstancedBaseVertexBaseInstance
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstancedBaseVertexBaseInstance.xml
func DrawElementsInstancedBaseVertexBaseInstance(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, basevertex Int, baseinstance Uint) {
	C.callDrawElementsInstancedBaseVertexBaseInstance((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount), (C.GLint)(basevertex), (C.GLuint)(baseinstance))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstancedBaseInstance
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstancedBaseInstance.xml
func DrawElementsInstancedBaseInstance(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, baseinstance Uint) {
	C.callDrawElementsInstancedBaseInstance((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteTextures
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteTextures.xml
func DeleteTextures(n Sizei, textures *Uint) {
	C.callDeleteTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramiv.xml
func GetProgramiv(program Uint, pname Enum, params *Int) {
	C.callGetProgramiv((C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetRenderbufferParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetRenderbufferParameteriv.xml
func GetRenderbufferParameteriv(target Enum, pname Enum, params *Int) {
	C.callGetRenderbufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClear
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClear.xml
func Clear(mask Bitfield) {
	C.callClear((C.GLbitfield)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPolygonMode
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPolygonMode.xml
func PolygonMode(face Enum, mode Enum) {
	C.callPolygonMode((C.GLenum)(face), (C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4uiv.xml
func ProgramUniform4uiv(program Uint, location Int, count Sizei, value *Uint) {
	C.callProgramUniform4uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSubroutineUniformLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSubroutineUniformLocation.xml
func GetSubroutineUniformLocation(program Uint, shadertype Enum, name *Char) Int {
	return (Int)(C.callGetSubroutineUniformLocation((C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArraysInstancedBaseInstance
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawArraysInstancedBaseInstance.xml
func DrawArraysInstancedBaseInstance(mode Enum, first Int, count Sizei, instancecount Sizei, baseinstance Uint) {
	C.callDrawArraysInstancedBaseInstance((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribPointerv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribPointerv.xml
func GetVertexAttribPointerv(index Uint, pname Enum, pointer *Ptr) {
	C.callGetVertexAttribPointerv((C.GLuint)(index), (C.GLenum)(pname), (*unsafe.Pointer)(pointer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEnableVertexAttribArray
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEnableVertexAttribArray.xml
func EnableVertexAttribArray(index Uint) {
	C.callEnableVertexAttribArray((C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsQuery
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsQuery.xml
func IsQuery(id Uint) Boolean {
	return (Boolean)(C.callIsQuery((C.GLuint)(id)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x2fv.xml
func ProgramUniformMatrix4x2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix4x2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilFuncSeparate
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilFuncSeparate.xml
func StencilFuncSeparate(face Enum, func_ Enum, ref Int, mask Uint) {
	C.callStencilFuncSeparate((C.GLenum)(face), (C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPixelStorei
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPixelStorei.xml
func PixelStorei(pname Enum, param Int) {
	C.callPixelStorei((C.GLenum)(pname), (C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterIiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterIiv.xml
func TexParameterIiv(target Enum, pname Enum, params *Int) {
	C.callTexParameterIiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveAtomicCounterBufferiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveAtomicCounterBufferiv.xml
func GetActiveAtomicCounterBufferiv(program Uint, bufferIndex Uint, pname Enum, params *Int) {
	C.callGetActiveAtomicCounterBufferiv((C.GLuint)(program), (C.GLuint)(bufferIndex), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveSubroutineUniformiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveSubroutineUniformiv.xml
func GetActiveSubroutineUniformiv(program Uint, shadertype Enum, index Uint, pname Enum, values *Int) {
	C.callGetActiveSubroutineUniformiv((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(values))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorMaski
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorMaski.xml
func ColorMaski(index Uint, r Boolean, g Boolean, b Boolean, a Boolean) {
	C.callColorMaski((C.GLuint)(index), (C.GLboolean)(r), (C.GLboolean)(g), (C.GLboolean)(b), (C.GLboolean)(a))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewportIndexedfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glViewportIndexedfv.xml
func ViewportIndexedfv(index Uint, v *Float) {
	C.callViewportIndexedfv((C.GLuint)(index), (*C.GLfloat)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3ui.xml
func ProgramUniform3ui(program Uint, location Int, v0 Uint, v1 Uint, v2 Uint) {
	C.callProgramUniform3ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP2ui.xml
func MultiTexCoordP2ui(texture Enum, type_ Enum, coords Uint) {
	C.callMultiTexCoordP2ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4f.xml
func ProgramUniform4f(program Uint, location Int, v0 Float, v1 Float, v2 Float, v3 Float) {
	C.callProgramUniform4f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorP3ui.xml
func ColorP3ui(type_ Enum, color Uint) {
	C.callColorP3ui((C.GLenum)(type_), (C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquationSeparate
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquationSeparate.xml
func BlendEquationSeparate(modeRGB Enum, modeAlpha Enum) {
	C.callBlendEquationSeparate((C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4i.xml
func ProgramUniform4i(program Uint, location Int, v0 Int, v1 Int, v2 Int, v3 Int) {
	C.callProgramUniform4i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFlush
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFlush.xml
func Flush() {
	C.callFlush()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFinish
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFinish.xml
func Finish() {
	C.callFinish()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3dv.xml
func UniformMatrix3dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix3dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEndTransformFeedback.xml
func EndTransformFeedback() {
	C.callEndTransformFeedback()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3d.xml
func ProgramUniform3d(program Uint, location Int, v0 Double, v1 Double, v2 Double) {
	C.callProgramUniform3d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetIntegerv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetIntegerv.xml
func GetIntegerv(pname Enum, params *Int) {
	C.callGetIntegerv((C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x3fv.xml
func UniformMatrix4x3fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix4x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTransformFeedbackVarying
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTransformFeedbackVarying.xml
func GetTransformFeedbackVarying(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Sizei, type_ *Enum, name *Char) {
	C.callGetTransformFeedbackVarying((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLsizei)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4d.xml
func ProgramUniform4d(program Uint, location Int, v0 Double, v1 Double, v2 Double, v3 Double) {
	C.callProgramUniform4d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLdouble)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceName
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceName.xml
func GetProgramResourceName(program Uint, programInterface Enum, index Uint, bufSize Sizei, length *Sizei, name *Char) {
	C.callGetProgramResourceName((C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteProgramPipelines
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteProgramPipelines.xml
func DeleteProgramPipelines(n Sizei, pipelines *Uint) {
	C.callDeleteProgramPipelines((C.GLsizei)(n), (*C.GLuint)(pipelines))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDisable
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDisable.xml
func Disable(cap Enum) {
	C.callDisable((C.GLenum)(cap))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribLPointer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribLPointer.xml
func VertexAttribLPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Ptr) {
	C.callVertexAttribLPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP3ui.xml
func VertexAttribP3ui(index Uint, type_ Enum, normalized Boolean, value Uint) {
	C.callVertexAttribP3ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProvokingVertex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProvokingVertex.xml
func ProvokingVertex(mode Enum) {
	C.callProvokingVertex((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArraysInstanced
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawArraysInstanced.xml
func DrawArraysInstanced(mode Enum, first Int, count Sizei, instancecount Sizei) {
	C.callDrawArraysInstanced((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glNormalP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glNormalP3uiv.xml
func NormalP3uiv(type_ Enum, coords *Uint) {
	C.callNormalP3uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2iv.xml
func Uniform2iv(location Int, count Sizei, value *Int) {
	C.callUniform2iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1d.xml
func ProgramUniform1d(program Uint, location Int, v0 Double) {
	C.callProgramUniform1d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPolygonOffset
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPolygonOffset.xml
func PolygonOffset(factor Float, units Float) {
	C.callPolygonOffset((C.GLfloat)(factor), (C.GLfloat)(units))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x4dv.xml
func UniformMatrix3x4dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix3x4dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage3D.xml
func TexStorage3D(target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei, depth Sizei) {
	C.callTexStorage3D((C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameterIiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterIiv.xml
func GetTexParameterIiv(target Enum, pname Enum, params *Int) {
	C.callGetTexParameterIiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4fv.xml
func ProgramUniformMatrix4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP2uiv.xml
func VertexAttribP2uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) {
	C.callVertexAttribP2uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexSubImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage1D.xml
func CopyTexSubImage1D(target Enum, level Int, xoffset Int, x Int, y Int, width Sizei) {
	C.callCopyTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSynciv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSynciv.xml
func GetSynciv(sync Sync, pname Enum, bufSize Sizei, length *Sizei, values *Int) {
	C.callGetSynciv((C.GLsync)(sync), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(values))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4f.xml
func Uniform4f(location Int, v0 Float, v1 Float, v2 Float, v3 Float) {
	C.callUniform4f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2fv.xml
func ProgramUniformMatrix2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteShader.xml
func DeleteShader(shader Uint) {
	C.callDeleteShader((C.GLuint)(shader))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindBuffer.xml
func BindBuffer(target Enum, buffer Uint) {
	C.callBindBuffer((C.GLenum)(target), (C.GLuint)(buffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEnable
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEnable.xml
func Enable(cap Enum) {
	C.callEnable((C.GLenum)(cap))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glWaitSync
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glWaitSync.xml
func WaitSync(sync Sync, flags Bitfield, timeout Uint64) {
	C.callWaitSync((C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindAttribLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindAttribLocation.xml
func BindAttribLocation(program Uint, index Uint, name *Char) {
	C.callBindAttribLocation((C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4iv.xml
func VertexAttribI4iv(index Uint, v *Int) {
	C.callVertexAttribI4iv((C.GLuint)(index), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramStageiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramStageiv.xml
func GetProgramStageiv(program Uint, shadertype Enum, pname Enum, values *Int) {
	C.callGetProgramStageiv((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLenum)(pname), (*C.GLint)(values))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP2uiv.xml
func TexCoordP2uiv(type_ Enum, coords *Uint) {
	C.callTexCoordP2uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP4ui.xml
func MultiTexCoordP4ui(texture Enum, type_ Enum, coords Uint) {
	C.callMultiTexCoordP4ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3iv.xml
func Uniform3iv(location Int, count Sizei, value *Int) {
	C.callUniform3iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetIntegeri_v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetIntegeri_v.xml
func GetIntegeri_v(target Enum, index Uint, data *Int) {
	C.callGetIntegeri_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstanced
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstanced.xml
func DrawElementsInstanced(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei) {
	C.callDrawElementsInstanced((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMinSampleShading
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMinSampleShading.xml
func MinSampleShading(value Float) {
	C.callMinSampleShading((C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderSource
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderSource.xml
func GetShaderSource(shader Uint, bufSize Sizei, length *Sizei, source *Char) {
	C.callGetShaderSource((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(source))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2uiv.xml
func ProgramUniform2uiv(program Uint, location Int, count Sizei, value *Uint) {
	C.callProgramUniform2uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilMaskSeparate
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilMaskSeparate.xml
func StencilMaskSeparate(face Enum, mask Uint) {
	C.callStencilMaskSeparate((C.GLenum)(face), (C.GLuint)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glScissor.xml
func Scissor(x Int, y Int, width Sizei, height Sizei) {
	C.callScissor((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderPrecisionFormat
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderPrecisionFormat.xml
func GetShaderPrecisionFormat(shadertype Enum, precisiontype Enum, range_ *Int, precision *Int) {
	C.callGetShaderPrecisionFormat((C.GLenum)(shadertype), (C.GLenum)(precisiontype), (*C.GLint)(range_), (*C.GLint)(precision))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4ui.xml
func VertexAttribI4ui(index Uint, x Uint, y Uint, z Uint, w Uint) {
	C.callVertexAttribI4ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorP3uiv.xml
func ColorP3uiv(type_ Enum, color *Uint) {
	C.callColorP3uiv((C.GLenum)(type_), (*C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformBlockName
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformBlockName.xml
func GetActiveUniformBlockName(program Uint, uniformBlockIndex Uint, bufSize Sizei, length *Sizei, uniformBlockName *Char) {
	C.callGetActiveUniformBlockName((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(uniformBlockName))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage3DMultisample
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage3DMultisample.xml
func TexStorage3DMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean) {
	C.callTexStorage3DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(fixedsamplelocations))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3uiv.xml
func VertexAttribI3uiv(index Uint, v *Uint) {
	C.callVertexAttribI3uiv((C.GLuint)(index), (*C.GLuint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP2uiv.xml
func VertexP2uiv(type_ Enum, value *Uint) {
	C.callVertexP2uiv((C.GLenum)(type_), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawElementsIndirect
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawElementsIndirect.xml
func MultiDrawElementsIndirect(mode Enum, type_ Enum, indirect Ptr, drawcount Sizei, stride Sizei) {
	C.callMultiDrawElementsIndirect((C.GLenum)(mode), (C.GLenum)(type_), (unsafe.Pointer)(indirect), (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4uiv.xml
func Uniform4uiv(location Int, count Sizei, value *Uint) {
	C.callUniform4uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameteriv.xml
func GetTexParameteriv(target Enum, pname Enum, params *Int) {
	C.callGetTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP2ui.xml
func VertexP2ui(type_ Enum, value Uint) {
	C.callVertexP2ui((C.GLenum)(type_), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenVertexArrays
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenVertexArrays.xml
func GenVertexArrays(n Sizei, arrays *Uint) {
	C.callGenVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceLocationIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceLocationIndex.xml
func GetProgramResourceLocationIndex(program Uint, programInterface Enum, name *Char) Int {
	return (Int)(C.callGetProgramResourceLocationIndex((C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameteriv.xml
func GetSamplerParameteriv(sampler Uint, pname Enum, params *Int) {
	C.callGetSamplerParameteriv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRangef
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthRangef.xml
func DepthRangef(n Float, f Float) {
	C.callDepthRangef((C.GLfloat)(n), (C.GLfloat)(f))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1ui.xml
func VertexAttribI1ui(index Uint, x Uint) {
	C.callVertexAttribI1ui((C.GLuint)(index), (C.GLuint)(x))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearColor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearColor.xml
func ClearColor(red Float, green Float, blue Float, alpha Float) {
	C.callClearColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterIuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterIuiv.xml
func TexParameterIuiv(target Enum, pname Enum, params *Uint) {
	C.callTexParameterIuiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2fv.xml
func ProgramUniform2fv(program Uint, location Int, count Sizei, value *Float) {
	C.callProgramUniform2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribBinding
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribBinding.xml
func VertexAttribBinding(attribindex Uint, bindingindex Uint) {
	C.callVertexAttribBinding((C.GLuint)(attribindex), (C.GLuint)(bindingindex))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP4uiv.xml
func TexCoordP4uiv(type_ Enum, coords *Uint) {
	C.callTexCoordP4uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPatchParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPatchParameterfv.xml
func PatchParameterfv(pname Enum, values *Float) {
	C.callPatchParameterfv((C.GLenum)(pname), (*C.GLfloat)(values))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage2D.xml
func CompressedTexImage2D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, border Int, imageSize Sizei, data Ptr) {
	C.callCompressedTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glActiveTexture
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glActiveTexture.xml
func ActiveTexture(texture Enum) {
	C.callActiveTexture((C.GLenum)(texture))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramBinary
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramBinary.xml
func GetProgramBinary(program Uint, bufSize Sizei, length *Sizei, binaryFormat *Enum, binary Ptr) {
	C.callGetProgramBinary((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLenum)(binaryFormat), (unsafe.Pointer)(binary))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1ui.xml
func Uniform1ui(location Int, v0 Uint) {
	C.callUniform1ui((C.GLint)(location), (C.GLuint)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexImage1D.xml
func CopyTexImage1D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, border Int) {
	C.callCopyTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRangeIndexed
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthRangeIndexed.xml
func DepthRangeIndexed(index Uint, n Double, f Double) {
	C.callDepthRangeIndexed((C.GLuint)(index), (C.GLdouble)(n), (C.GLdouble)(f))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexImage2D.xml
func CopyTexImage2D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, height Sizei, border Int) {
	C.callCopyTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissorArrayv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glScissorArrayv.xml
func ScissorArrayv(first Uint, count Sizei, v *Int) {
	C.callScissorArrayv((C.GLuint)(first), (C.GLsizei)(count), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderInfoLog
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderInfoLog.xml
func GetShaderInfoLog(shader Uint, bufSize Sizei, length *Sizei, infoLog *Char) {
	C.callGetShaderInfoLog((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetMultisamplefv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetMultisamplefv.xml
func GetMultisamplefv(pname Enum, index Uint, val *Float) {
	C.callGetMultisamplefv((C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(val))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArrays
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawArrays.xml
func DrawArrays(mode Enum, first Int, count Sizei) {
	C.callDrawArrays((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformBlockiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformBlockiv.xml
func GetActiveUniformBlockiv(program Uint, uniformBlockIndex Uint, pname Enum, params *Int) {
	C.callGetActiveUniformBlockiv((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3fv.xml
func UniformMatrix3fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCreateProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCreateProgram.xml
func CreateProgram() Uint {
	return (Uint)(C.callCreateProgram())
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1uiv.xml
func VertexAttribI1uiv(index Uint, v *Uint) {
	C.callVertexAttribI1uiv((C.GLuint)(index), (*C.GLuint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjectuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectuiv.xml
func GetQueryObjectuiv(id Uint, pname Enum, params *Uint) {
	C.callGetQueryObjectuiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndQuery
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEndQuery.xml
func EndQuery(target Enum) {
	C.callEndQuery((C.GLenum)(target))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearDepth
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearDepth.xml
func ClearDepth(depth Double) {
	C.callClearDepth((C.GLdouble)(depth))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFloati_v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFloati_v.xml
func GetFloati_v(target Enum, index Uint, data *Float) {
	C.callGetFloati_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLfloat)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPointParameterfv.xml
func PointParameterfv(pname Enum, params *Float) {
	C.callPointParameterfv((C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjectiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectiv.xml
func GetQueryObjectiv(id Uint, pname Enum, params *Int) {
	C.callGetQueryObjectiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformName
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformName.xml
func GetActiveUniformName(program Uint, uniformIndex Uint, bufSize Sizei, length *Sizei, uniformName *Char) {
	C.callGetActiveUniformName((C.GLuint)(program), (C.GLuint)(uniformIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(uniformName))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glNormalP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glNormalP3ui.xml
func NormalP3ui(type_ Enum, coords Uint) {
	C.callNormalP3ui((C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glActiveShaderProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glActiveShaderProgram.xml
func ActiveShaderProgram(pipeline Uint, program Uint) {
	C.callActiveShaderProgram((C.GLuint)(pipeline), (C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3iv.xml
func VertexAttribI3iv(index Uint, v *Int) {
	C.callVertexAttribI3iv((C.GLuint)(index), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsFramebuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsFramebuffer.xml
func IsFramebuffer(framebuffer Uint) Boolean {
	return (Boolean)(C.callIsFramebuffer((C.GLuint)(framebuffer)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindImageTexture
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindImageTexture.xml
func BindImageTexture(unit Uint, texture Uint, level Int, layered Boolean, layer Int, access Enum, format Enum) {
	C.callBindImageTexture((C.GLuint)(unit), (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(layered), (C.GLint)(layer), (C.GLenum)(access), (C.GLenum)(format))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniform
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniform.xml
func GetActiveUniform(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char) {
	C.callGetActiveUniform((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1i.xml
func Uniform1i(location Int, v0 Int) {
	C.callUniform1i((C.GLint)(location), (C.GLint)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glObjectPtrLabel
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glObjectPtrLabel.xml
func ObjectPtrLabel(ptr Ptr, length Sizei, label *Char) {
	C.callObjectPtrLabel((unsafe.Pointer)(ptr), (C.GLsizei)(length), (*C.GLchar)(label))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUseProgramStages
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUseProgramStages.xml
func UseProgramStages(pipeline Uint, stages Bitfield, program Uint) {
	C.callUseProgramStages((C.GLuint)(pipeline), (C.GLbitfield)(stages), (C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexLevelParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexLevelParameteriv.xml
func GetTexLevelParameteriv(target Enum, level Int, pname Enum, params *Int) {
	C.callGetTexLevelParameteriv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribLFormat
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribLFormat.xml
func VertexAttribLFormat(attribindex Uint, size Int, type_ Enum, relativeoffset Uint) {
	C.callVertexAttribLFormat((C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(type_), (C.GLuint)(relativeoffset))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameterIiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameterIiv.xml
func GetSamplerParameterIiv(sampler Uint, pname Enum, params *Int) {
	C.callGetSamplerParameterIiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBufferSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBufferSubData.xml
func BufferSubData(target Enum, offset Intptr, size Sizeiptr, data Ptr) {
	C.callBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexSubImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage2D.xml
func TexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Ptr) {
	C.callTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3ui.xml
func Uniform3ui(location Int, v0 Uint, v1 Uint, v2 Uint) {
	C.callUniform3ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage2DMultisample
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexImage2DMultisample.xml
func TexImage2DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, fixedsamplelocations Boolean) {
	C.callTexImage2DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(fixedsamplelocations))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glRenderbufferStorageMultisample
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glRenderbufferStorageMultisample.xml
func RenderbufferStorageMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei) {
	C.callRenderbufferStorageMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsProgram.xml
func IsProgram(program Uint) Boolean {
	return (Boolean)(C.callIsProgram((C.GLuint)(program)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP2ui.xml
func VertexAttribP2ui(index Uint, type_ Enum, normalized Boolean, value Uint) {
	C.callVertexAttribP2ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4bv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4bv.xml
func VertexAttribI4bv(index Uint, v *Byte) {
	C.callVertexAttribI4bv((C.GLuint)(index), (*C.GLbyte)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2i.xml
func Uniform2i(location Int, v0 Int, v1 Int) {
	C.callUniform2i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteRenderbuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteRenderbuffers.xml
func DeleteRenderbuffers(n Sizei, renderbuffers *Uint) {
	C.callDeleteRenderbuffers((C.GLsizei)(n), (*C.GLuint)(renderbuffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferPointerv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferPointerv.xml
func GetBufferPointerv(target Enum, pname Enum, params *Ptr) {
	C.callGetBufferPointerv((C.GLenum)(target), (C.GLenum)(pname), (*unsafe.Pointer)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawBuffer.xml
func DrawBuffer(mode Enum) {
	C.callDrawBuffer((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMapBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMapBuffer.xml
func MapBuffer(target Enum, access Enum) *Ptr {
	return (*Ptr)(C.callMapBuffer((C.GLenum)(target), (C.GLenum)(access)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRangeArrayv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthRangeArrayv.xml
func DepthRangeArrayv(first Uint, count Sizei, v *Double) {
	C.callDepthRangeArrayv((C.GLuint)(first), (C.GLsizei)(count), (*C.GLdouble)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP3uiv.xml
func VertexP3uiv(type_ Enum, value *Uint) {
	C.callVertexP3uiv((C.GLenum)(type_), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture3D.xml
func FramebufferTexture3D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int, zoffset Int) {
	C.callFramebufferTexture3D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(zoffset))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4ubv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4ubv.xml
func VertexAttribI4ubv(index Uint, v *Ubyte) {
	C.callVertexAttribI4ubv((C.GLuint)(index), (*C.GLubyte)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4i.xml
func Uniform4i(location Int, v0 Int, v1 Int, v2 Int, v3 Int) {
	C.callUniform4i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexImage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage3D.xml
func CompressedTexImage3D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, depth Sizei, border Int, imageSize Sizei, data Ptr) {
	C.callCompressedTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteTransformFeedbacks
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteTransformFeedbacks.xml
func DeleteTransformFeedbacks(n Sizei, ids *Uint) {
	C.callDeleteTransformFeedbacks((C.GLsizei)(n), (*C.GLuint)(ids))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginQuery
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBeginQuery.xml
func BeginQuery(target Enum, id Uint) {
	C.callBeginQuery((C.GLenum)(target), (C.GLuint)(id))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glReadBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glReadBuffer.xml
func ReadBuffer(mode Enum) {
	C.callReadBuffer((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDebugMessageControl
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDebugMessageControl.xml
func DebugMessageControl(source Enum, type_ Enum, severity Enum, count Sizei, ids *Uint, enabled Boolean) {
	C.callDebugMessageControl((C.GLenum)(source), (C.GLenum)(type_), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(ids), (C.GLboolean)(enabled))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSubroutineIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSubroutineIndex.xml
func GetSubroutineIndex(program Uint, shadertype Enum, name *Char) Uint {
	return (Uint)(C.callGetSubroutineIndex((C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2uiv.xml
func VertexAttribI2uiv(index Uint, v *Uint) {
	C.callVertexAttribI2uiv((C.GLuint)(index), (*C.GLuint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMapBufferRange
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMapBufferRange.xml
func MapBufferRange(target Enum, offset Intptr, length Sizeiptr, access Bitfield) *Ptr {
	return (*Ptr)(C.callMapBufferRange((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length), (C.GLbitfield)(access)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawBuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawBuffers.xml
func DrawBuffers(n Sizei, bufs *Enum) {
	C.callDrawBuffers((C.GLsizei)(n), (*C.GLenum)(bufs))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUnmapBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUnmapBuffer.xml
func UnmapBuffer(target Enum) Boolean {
	return (Boolean)(C.callUnmapBuffer((C.GLenum)(target)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3d.xml
func Uniform3d(location Int, x Double, y Double, z Double) {
	C.callUniform3d((C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetAttribLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetAttribLocation.xml
func GetAttribLocation(program Uint, name *Char) Int {
	return (Int)(C.callGetAttribLocation((C.GLuint)(program), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUseProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUseProgram.xml
func UseProgram(program Uint) {
	C.callUseProgram((C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexSubImage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage3D.xml
func TexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, type_ Enum, pixels Ptr) {
	C.callTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4fv.xml
func UniformMatrix4fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x4dv.xml
func ProgramUniformMatrix3x4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix3x4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFuncSeparatei
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendFuncSeparatei.xml
func BlendFuncSeparatei(buf Uint, srcRGB Enum, dstRGB Enum, srcAlpha Enum, dstAlpha Enum) {
	C.callBlendFuncSeparatei((C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP3uiv.xml
func VertexAttribP3uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) {
	C.callVertexAttribP3uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsRenderbuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsRenderbuffer.xml
func IsRenderbuffer(renderbuffer Uint) Boolean {
	return (Boolean)(C.callIsRenderbuffer((C.GLuint)(renderbuffer)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsProgramPipeline
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsProgramPipeline.xml
func IsProgramPipeline(pipeline Uint) Boolean {
	return (Boolean)(C.callIsProgramPipeline((C.GLuint)(pipeline)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissorIndexed
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glScissorIndexed.xml
func ScissorIndexed(index Uint, left Int, bottom Int, width Sizei, height Sizei) {
	C.callScissorIndexed((C.GLuint)(index), (C.GLint)(left), (C.GLint)(bottom), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendColor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendColor.xml
func BlendColor(red Float, green Float, blue Float, alpha Float) {
	C.callBlendColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterf
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterf.xml
func TexParameterf(target Enum, pname Enum, param Float) {
	C.callTexParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateTexSubImage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateTexSubImage.xml
func InvalidateTexSubImage(texture Uint, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei) {
	C.callInvalidateTexSubImage((C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFragDataLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFragDataLocation.xml
func GetFragDataLocation(program Uint, name *Char) Int {
	return (Int)(C.callGetFragDataLocation((C.GLuint)(program), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x3fv.xml
func UniformMatrix2x3fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix2x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4d.xml
func Uniform4d(location Int, x Double, y Double, z Double, w Double) {
	C.callUniform4d((C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferfv.xml
func ClearBufferfv(buffer Enum, drawbuffer Int, value *Float) {
	C.callClearBufferfv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryIndexediv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryIndexediv.xml
func GetQueryIndexediv(target Enum, index Uint, pname Enum, params *Int) {
	C.callGetQueryIndexediv((C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3fv.xml
func ProgramUniform3fv(program Uint, location Int, count Sizei, value *Float) {
	C.callProgramUniform3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1iv.xml
func ProgramUniform1iv(program Uint, location Int, count Sizei, value *Int) {
	C.callProgramUniform1iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2fv.xml
func Uniform2fv(location Int, count Sizei, value *Float) {
	C.callUniform2fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramInterfaceiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramInterfaceiv.xml
func GetProgramInterfaceiv(program Uint, programInterface Enum, pname Enum, params *Int) {
	C.callGetProgramInterfaceiv((C.GLuint)(program), (C.GLenum)(programInterface), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2dv.xml
func UniformMatrix2dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix2dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsSync
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsSync.xml
func IsSync(sync Sync) Boolean {
	return (Boolean)(C.callIsSync((C.GLsync)(sync)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBooleanv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBooleanv.xml
func GetBooleanv(pname Enum, params *Boolean) {
	C.callGetBooleanv((C.GLenum)(pname), (*C.GLboolean)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilOp
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilOp.xml
func StencilOp(fail Enum, zfail Enum, zpass Enum) {
	C.callStencilOp((C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjectui64v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectui64v.xml
func GetQueryObjectui64v(id Uint, pname Enum, params *Uint64) {
	C.callGetQueryObjectui64v((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint64)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x3dv.xml
func ProgramUniformMatrix4x3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix4x3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetObjectPtrLabel
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetObjectPtrLabel.xml
func GetObjectPtrLabel(ptr Ptr, bufSize Sizei, length *Sizei, label *Char) {
	C.callGetObjectPtrLabel((unsafe.Pointer)(ptr), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(label))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferfi
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferfi.xml
func ClearBufferfi(buffer Enum, drawbuffer Int, depth Float, stencil Int) {
	C.callClearBufferfi((C.GLenum)(buffer), (C.GLint)(drawbuffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x4fv.xml
func ProgramUniformMatrix2x4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix2x4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3fv.xml
func Uniform3fv(location Int, count Sizei, value *Float) {
	C.callUniform3fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP4ui.xml
func VertexP4ui(type_ Enum, value Uint) {
	C.callVertexP4ui((C.GLenum)(type_), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjecti64v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjecti64v.xml
func GetQueryObjecti64v(id Uint, pname Enum, params *Int64) {
	C.callGetQueryObjecti64v((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint64)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage1D.xml
func CompressedTexImage1D(target Enum, level Int, internalformat Enum, width Sizei, border Int, imageSize Sizei, data Ptr) {
	C.callCompressedTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPrimitiveRestartIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPrimitiveRestartIndex.xml
func PrimitiveRestartIndex(index Uint) {
	C.callPrimitiveRestartIndex((C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedback.xml
func DrawTransformFeedback(mode Enum, id Uint) {
	C.callDrawTransformFeedback((C.GLenum)(mode), (C.GLuint)(id))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetDoublev
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetDoublev.xml
func GetDoublev(pname Enum, params *Double) {
	C.callGetDoublev((C.GLenum)(pname), (*C.GLdouble)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFuncSeparate
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendFuncSeparate.xml
func BlendFuncSeparate(sfactorRGB Enum, dfactorRGB Enum, sfactorAlpha Enum, dfactorAlpha Enum) {
	C.callBlendFuncSeparate((C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP1ui.xml
func VertexAttribP1ui(index Uint, type_ Enum, normalized Boolean, value Uint) {
	C.callVertexAttribP1ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP4uiv.xml
func VertexAttribP4uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) {
	C.callVertexAttribP4uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBeginTransformFeedback.xml
func BeginTransformFeedback(primitiveMode Enum) {
	C.callBeginTransformFeedback((C.GLenum)(primitiveMode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearStencil
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearStencil.xml
func ClearStencil(s Int) {
	C.callClearStencil((C.GLint)(s))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteFramebuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteFramebuffers.xml
func DeleteFramebuffers(n Sizei, framebuffers *Uint) {
	C.callDeleteFramebuffers((C.GLsizei)(n), (*C.GLuint)(framebuffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL1dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL1dv.xml
func VertexAttribL1dv(index Uint, v *Double) {
	C.callVertexAttribL1dv((C.GLuint)(index), (*C.GLdouble)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameterf
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPointParameterf.xml
func PointParameterf(pname Enum, param Float) {
	C.callPointParameterf((C.GLenum)(pname), (C.GLfloat)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4sv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4sv.xml
func VertexAttribI4sv(index Uint, v *Short) {
	C.callVertexAttribI4sv((C.GLuint)(index), (*C.GLshort)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL2dv.xml
func VertexAttribL2dv(index Uint, v *Double) {
	C.callVertexAttribL2dv((C.GLuint)(index), (*C.GLdouble)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFunci
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendFunci.xml
func BlendFunci(buf Uint, src Enum, dst Enum) {
	C.callBlendFunci((C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2ui.xml
func Uniform2ui(location Int, v0 Uint, v1 Uint) {
	C.callUniform2ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramPipelineiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramPipelineiv.xml
func GetProgramPipelineiv(pipeline Uint, pname Enum, params *Int) {
	C.callGetProgramPipelineiv((C.GLuint)(pipeline), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSampleMaski
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSampleMaski.xml
func SampleMaski(index Uint, mask Bitfield) {
	C.callSampleMaski((C.GLuint)(index), (C.GLbitfield)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameterIuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameterIuiv.xml
func GetSamplerParameterIuiv(sampler Uint, pname Enum, params *Uint) {
	C.callGetSamplerParameterIuiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyBufferSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyBufferSubData.xml
func CopyBufferSubData(readTarget Enum, writeTarget Enum, readOffset Intptr, writeOffset Intptr, size Sizeiptr) {
	C.callCopyBufferSubData((C.GLenum)(readTarget), (C.GLenum)(writeTarget), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizeiptr)(size))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformsiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformsiv.xml
func GetActiveUniformsiv(program Uint, uniformCount Sizei, uniformIndices *Uint, pname Enum, params *Int) {
	C.callGetActiveUniformsiv((C.GLuint)(program), (C.GLsizei)(uniformCount), (*C.GLuint)(uniformIndices), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferParameteriv.xml
func GetBufferParameteriv(target Enum, pname Enum, params *Int) {
	C.callGetBufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilOpSeparate
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilOpSeparate.xml
func StencilOpSeparate(face Enum, sfail Enum, dpfail Enum, dppass Enum) {
	C.callStencilOpSeparate((C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlitFramebuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlitFramebuffer.xml
func BlitFramebuffer(srcX0 Int, srcY0 Int, srcX1 Int, srcY1 Int, dstX0 Int, dstY0 Int, dstX1 Int, dstY1 Int, mask Bitfield, filter Enum) {
	C.callBlitFramebuffer((C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthFunc
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthFunc.xml
func DepthFunc(func_ Enum) {
	C.callDepthFunc((C.GLenum)(func_))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDetachShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDetachShader.xml
func DetachShader(program Uint, shader Uint) {
	C.callDetachShader((C.GLuint)(program), (C.GLuint)(shader))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferuiv.xml
func ClearBufferuiv(buffer Enum, drawbuffer Int, value *Uint) {
	C.callClearBufferuiv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramBinary
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramBinary.xml
func ProgramBinary(program Uint, binaryFormat Enum, binary Ptr, length Sizei) {
	C.callProgramBinary((C.GLuint)(program), (C.GLenum)(binaryFormat), (unsafe.Pointer)(binary), (C.GLsizei)(length))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewportIndexedf
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glViewportIndexedf.xml
func ViewportIndexedf(index Uint, x Float, y Float, w Float, h Float) {
	C.callViewportIndexedf((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(w), (C.GLfloat)(h))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3dv.xml
func ProgramUniformMatrix3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x4dv.xml
func UniformMatrix2x4dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix2x4dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramParameteri.xml
func ProgramParameteri(program Uint, pname Enum, value Int) {
	C.callProgramParameteri((C.GLuint)(program), (C.GLenum)(pname), (C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP3ui.xml
func MultiTexCoordP3ui(texture Enum, type_ Enum, coords Uint) {
	C.callMultiTexCoordP3ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4dv.xml
func UniformMatrix4dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix4dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindFragDataLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindFragDataLocation.xml
func BindFragDataLocation(program Uint, color Uint, name *Char) {
	C.callBindFragDataLocation((C.GLuint)(program), (C.GLuint)(color), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetAttachedShaders
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetAttachedShaders.xml
func GetAttachedShaders(program Uint, maxCount Sizei, count *Sizei, obj *Uint) {
	C.callGetAttachedShaders((C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(count), (*C.GLuint)(obj))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP4uiv.xml
func VertexP4uiv(type_ Enum, value *Uint) {
	C.callVertexP4uiv((C.GLenum)(type_), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glReadPixels
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glReadPixels.xml
func ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Ptr) {
	C.callReadPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribIFormat
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribIFormat.xml
func VertexAttribIFormat(attribindex Uint, size Int, type_ Enum, relativeoffset Uint) {
	C.callVertexAttribIFormat((C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(type_), (C.GLuint)(relativeoffset))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage1D.xml
func TexStorage1D(target Enum, levels Sizei, internalformat Enum, width Sizei) {
	C.callTexStorage1D((C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4dv.xml
func ProgramUniform4dv(program Uint, location Int, count Sizei, value *Double) {
	C.callProgramUniform4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexBindingDivisor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexBindingDivisor.xml
func VertexBindingDivisor(bindingindex Uint, divisor Uint) {
	C.callVertexBindingDivisor((C.GLuint)(bindingindex), (C.GLuint)(divisor))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawRangeElementsBaseVertex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawRangeElementsBaseVertex.xml
func DrawRangeElementsBaseVertex(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Ptr, basevertex Int) {
	C.callDrawRangeElementsBaseVertex((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLint)(basevertex))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformiv.xml
func GetUniformiv(program Uint, location Int, params *Int) {
	C.callGetUniformiv((C.GLuint)(program), (C.GLint)(location), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL3dv.xml
func VertexAttribL3dv(index Uint, v *Double) {
	C.callVertexAttribL3dv((C.GLuint)(index), (*C.GLdouble)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateBufferData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateBufferData.xml
func InvalidateBufferData(buffer Uint) {
	C.callInvalidateBufferData((C.GLuint)(buffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glResumeTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glResumeTransformFeedback.xml
func ResumeTransformFeedback() {
	C.callResumeTransformFeedback()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x3dv.xml
func UniformMatrix2x3dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix2x3dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndQueryIndexed
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEndQueryIndexed.xml
func EndQueryIndexed(target Enum, index Uint) {
	C.callEndQueryIndexed((C.GLenum)(target), (C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribIuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribIuiv.xml
func GetVertexAttribIuiv(index Uint, pname Enum, params *Uint) {
	C.callGetVertexAttribIuiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginQueryIndexed
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBeginQueryIndexed.xml
func BeginQueryIndexed(target Enum, index Uint, id Uint) {
	C.callBeginQueryIndexed((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(id))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x3dv.xml
func UniformMatrix4x3dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix4x3dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexBufferRange
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexBufferRange.xml
func TexBufferRange(target Enum, internalformat Enum, buffer Uint, offset Intptr, size Sizeiptr) {
	C.callTexBufferRange((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribLdv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribLdv.xml
func GetVertexAttribLdv(index Uint, pname Enum, params *Double) {
	C.callGetVertexAttribLdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedbackStream
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedbackStream.xml
func DrawTransformFeedbackStream(mode Enum, id Uint, stream Uint) {
	C.callDrawTransformFeedbackStream((C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindFramebuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindFramebuffer.xml
func BindFramebuffer(target Enum, framebuffer Uint) {
	C.callBindFramebuffer((C.GLenum)(target), (C.GLuint)(framebuffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteProgram.xml
func DeleteProgram(program Uint) {
	C.callDeleteProgram((C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissorIndexedv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glScissorIndexedv.xml
func ScissorIndexedv(index Uint, v *Int) {
	C.callScissorIndexedv((C.GLuint)(index), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x2dv.xml
func UniformMatrix3x2dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix3x2dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsTexture
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsTexture.xml
func IsTexture(texture Uint) Boolean {
	return (Boolean)(C.callIsTexture((C.GLuint)(texture)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterIiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterIiv.xml
func SamplerParameterIiv(sampler Uint, pname Enum, param *Int) {
	C.callSamplerParameterIiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4fv.xml
func ProgramUniform4fv(program Uint, location Int, count Sizei, value *Float) {
	C.callProgramUniform4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquationi
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquationi.xml
func BlendEquationi(buf Uint, mode Enum) {
	C.callBlendEquationi((C.GLuint)(buf), (C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteSync
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteSync.xml
func DeleteSync(sync Sync) {
	C.callDeleteSync((C.GLsync)(sync))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3iv.xml
func ProgramUniform3iv(program Uint, location Int, count Sizei, value *Int) {
	C.callProgramUniform3iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateTexImage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateTexImage.xml
func InvalidateTexImage(texture Uint, level Int) {
	C.callInvalidateTexImage((C.GLuint)(texture), (C.GLint)(level))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewportArrayv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glViewportArrayv.xml
func ViewportArrayv(first Uint, count Sizei, v *Float) {
	C.callViewportArrayv((C.GLuint)(first), (C.GLsizei)(count), (*C.GLfloat)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexSubImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage1D.xml
func TexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, type_ Enum, pixels Ptr) {
	C.callTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCreateShaderProgramv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCreateShaderProgramv.xml
func CreateShaderProgramv(type_ Enum, count Sizei, strings **Char) Uint {
	return (Uint)(C.callCreateShaderProgramv((C.GLenum)(type_), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(strings))))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenRenderbuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenRenderbuffers.xml
func GenRenderbuffers(n Sizei, renderbuffers *Uint) {
	C.callGenRenderbuffers((C.GLsizei)(n), (*C.GLuint)(renderbuffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2dv.xml
func ProgramUniformMatrix2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2ui.xml
func ProgramUniform2ui(program Uint, location Int, v0 Uint, v1 Uint) {
	C.callProgramUniform2ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexSubImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage2D.xml
func CopyTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, x Int, y Int, width Sizei, height Sizei) {
	C.callCopyTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawElementsBaseVertex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawElementsBaseVertex.xml
func MultiDrawElementsBaseVertex(mode Enum, count *Sizei, type_ Enum, indices *Ptr, drawcount Sizei, basevertex *Int) {
	C.callMultiDrawElementsBaseVertex((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(drawcount), (*C.GLint)(basevertex))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetCompressedTexImage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetCompressedTexImage.xml
func GetCompressedTexImage(target Enum, level Int, img Ptr) {
	C.callGetCompressedTexImage((C.GLenum)(target), (C.GLint)(level), (unsafe.Pointer)(img))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3ui.xml
func VertexAttribI3ui(index Uint, x Uint, y Uint, z Uint) {
	C.callVertexAttribI3ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x4fv.xml
func UniformMatrix3x4fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix3x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1iv.xml
func VertexAttribI1iv(index Uint, v *Int) {
	C.callVertexAttribI1iv((C.GLuint)(index), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferSubData.xml
func GetBufferSubData(target Enum, offset Intptr, size Sizeiptr, data Ptr) {
	C.callGetBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexImage1D.xml
func TexImage1D(target Enum, level Int, internalformat Int, width Sizei, border Int, format Enum, type_ Enum, pixels Ptr) {
	C.callTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture.xml
func FramebufferTexture(target Enum, attachment Enum, texture Uint, level Int) {
	C.callFramebufferTexture((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1i.xml
func VertexAttribI1i(index Uint, x Int) {
	C.callVertexAttribI1i((C.GLuint)(index), (C.GLint)(x))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribIiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribIiv.xml
func GetVertexAttribIiv(index Uint, pname Enum, params *Int) {
	C.callGetVertexAttribIiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferParameteri64v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferParameteri64v.xml
func GetBufferParameteri64v(target Enum, pname Enum, params *Int64) {
	C.callGetBufferParameteri64v((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint64)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginConditionalRender
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBeginConditionalRender.xml
func BeginConditionalRender(id Uint, mode Enum) {
	C.callBeginConditionalRender((C.GLuint)(id), (C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInternalformati64v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetInternalformati64v.xml
func GetInternalformati64v(target Enum, internalformat Enum, pname Enum, bufSize Sizei, params *Int64) {
	C.callGetInternalformati64v((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint64)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2ui.xml
func VertexAttribI2ui(index Uint, x Uint, y Uint) {
	C.callVertexAttribI2ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquation.xml
func BlendEquation(mode Enum) {
	C.callBlendEquation((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEnablei
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEnablei.xml
func Enablei(target Enum, index Uint) {
	C.callEnablei((C.GLenum)(target), (C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetStringi
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetStringi.xml
func GetStringi(name Enum, index Uint) *Ubyte {
	return (*Ubyte)(C.callGetStringi((C.GLenum)(name), (C.GLuint)(index)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x2dv.xml
func ProgramUniformMatrix3x2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix3x2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedbackStreamInstanced
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedbackStreamInstanced.xml
func DrawTransformFeedbackStreamInstanced(mode Enum, id Uint, stream Uint, instancecount Sizei) {
	C.callDrawTransformFeedbackStreamInstanced((C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream), (C.GLsizei)(instancecount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveSubroutineUniformName
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveSubroutineUniformName.xml
func GetActiveSubroutineUniformName(program Uint, shadertype Enum, index Uint, bufsize Sizei, length *Sizei, name *Char) {
	C.callGetActiveSubroutineUniformName((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(length), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInteger64i_v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetInteger64i_v.xml
func GetInteger64i_v(target Enum, index Uint, data *Int64) {
	C.callGetInteger64i_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLint64)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDebugMessageInsert
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDebugMessageInsert.xml
func DebugMessageInsert(source Enum, type_ Enum, id Uint, severity Enum, length Sizei, buf *Char) {
	C.callDebugMessageInsert((C.GLenum)(source), (C.GLenum)(type_), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(buf))
}
