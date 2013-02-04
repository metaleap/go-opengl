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
// void (APIENTRYP ptrglBindImageTexture)(GLuint unit, GLuint texture, GLint level, GLboolean layered, GLint layer, GLenum access, GLenum format);
// void (APIENTRYP ptrglProgramUniform4i)(GLuint program, GLint location, GLint v0, GLint v1, GLint v2, GLint v3);
// void (APIENTRYP ptrglGetBufferParameteri64v)(GLenum target, GLenum pname, GLint64* params);
// void (APIENTRYP ptrglRenderbufferStorage)(GLenum target, GLenum internalformat, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglDepthRangeArrayv)(GLuint first, GLsizei count, GLdouble* v);
// void (APIENTRYP ptrglFramebufferTexture)(GLenum target, GLenum attachment, GLuint texture, GLint level);
// void (APIENTRYP ptrglGetProgramResourceiv)(GLuint program, GLenum programInterface, GLuint index, GLsizei propCount, GLenum* props, GLsizei bufSize, GLsizei* length, GLint* params);
// void (APIENTRYP ptrglMultiTexCoordP2uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglBindSampler)(GLuint unit, GLuint sampler);
// void (APIENTRYP ptrglUniformMatrix4x3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglGetVertexAttribIiv)(GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrglMemoryBarrier)(GLbitfield barriers);
// void (APIENTRYP ptrglBeginQueryIndexed)(GLenum target, GLuint index, GLuint id);
// void (APIENTRYP ptrglUniform2d)(GLint location, GLdouble x, GLdouble y);
// void (APIENTRYP ptrglDeleteRenderbuffers)(GLsizei n, GLuint* renderbuffers);
// void (APIENTRYP ptrglUseProgramStages)(GLuint pipeline, GLbitfield stages, GLuint program);
// void (APIENTRYP ptrglDrawBuffers)(GLsizei n, GLenum* bufs);
// void (APIENTRYP ptrglGetSynciv)(GLsync sync, GLenum pname, GLsizei bufSize, GLsizei* length, GLint* values);
// void (APIENTRYP ptrglEndQuery)(GLenum target);
// void (APIENTRYP ptrglDepthRangeIndexed)(GLuint index, GLdouble n, GLdouble f);
// GLenum (APIENTRYP ptrglGetError)();
// void (APIENTRYP ptrglProgramUniform4d)(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3);
// void (APIENTRYP ptrglPushDebugGroup)(GLenum source, GLuint id, GLsizei length, GLchar* message);
// void (APIENTRYP ptrglGetQueryiv)(GLenum target, GLenum pname, GLint* params);
// GLboolean (APIENTRYP ptrglIsEnabled)(GLenum cap);
// void (APIENTRYP ptrglSecondaryColorP3uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrglTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglDisablei)(GLenum target, GLuint index);
// void (APIENTRYP ptrglDrawRangeElementsBaseVertex)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices, GLint basevertex);
// void (APIENTRYP ptrglBindTexture)(GLenum target, GLuint texture);
// void (APIENTRYP ptrglUniform4iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglMinSampleShading)(GLfloat value);
// void (APIENTRYP ptrglUniform1dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglVertexAttribP1ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglPixelStorei)(GLenum pname, GLint param);
// void (APIENTRYP ptrglFrontFace)(GLenum mode);
// void (APIENTRYP ptrglProgramUniform2d)(GLuint program, GLint location, GLdouble v0, GLdouble v1);
// void (APIENTRYP ptrglTexStorage1D)(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width);
// void (APIENTRYP ptrglUniform3i)(GLint location, GLint v0, GLint v1, GLint v2);
// void (APIENTRYP ptrglTexImage2DMultisample)(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrglUniformMatrix3x2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglTexImage2D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// GLint (APIENTRYP ptrglGetSubroutineUniformLocation)(GLuint program, GLenum shadertype, GLchar* name);
// void (APIENTRYP ptrglUniformMatrix3dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglBindRenderbuffer)(GLenum target, GLuint renderbuffer);
// void (APIENTRYP ptrglBindProgramPipeline)(GLuint pipeline);
// void (APIENTRYP ptrglUniform4dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglClearBufferuiv)(GLenum buffer, GLint drawbuffer, GLuint* value);
// void (APIENTRYP ptrglBeginTransformFeedback)(GLenum primitiveMode);
// void (APIENTRYP ptrglUniformMatrix3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglDispatchCompute)(GLuint num_groups_x, GLuint num_groups_y, GLuint num_groups_z);
// void (APIENTRYP ptrglBindVertexBuffer)(GLuint bindingindex, GLuint buffer, GLintptr offset, GLsizei stride);
// void (APIENTRYP ptrglVertexBindingDivisor)(GLuint bindingindex, GLuint divisor);
// void (APIENTRYP ptrglMultiTexCoordP4uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglProgramUniformMatrix2x3fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglGetTexLevelParameteriv)(GLenum target, GLint level, GLenum pname, GLint* params);
// void (APIENTRYP ptrglVertexP4uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrglGetTexParameterIuiv)(GLenum target, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglDebugMessageControl)(GLenum source, GLenum type, GLenum severity, GLsizei count, GLuint* ids, GLboolean enabled);
// void (APIENTRYP ptrglSamplerParameteri)(GLuint sampler, GLenum pname, GLint param);
// void (APIENTRYP ptrglProgramUniform2f)(GLuint program, GLint location, GLfloat v0, GLfloat v1);
// void (APIENTRYP ptrglGetVertexAttribfv)(GLuint index, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglVertexAttribP2uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrglScissorIndexed)(GLuint index, GLint left, GLint bottom, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglProgramUniform1d)(GLuint program, GLint location, GLdouble v0);
// void (APIENTRYP ptrglVertexP4ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrglDepthMask)(GLboolean flag);
// void (APIENTRYP ptrglUniformMatrix2dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglTexCoordP4uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglUniform4f)(GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3);
// void (APIENTRYP ptrglUniform4fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglClampColor)(GLenum target, GLenum clamp);
// void (APIENTRYP ptrglGenFramebuffers)(GLsizei n, GLuint* framebuffers);
// void (APIENTRYP ptrglTexCoordP3uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglColorMask)(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
// void (APIENTRYP ptrglSamplerParameterfv)(GLuint sampler, GLenum pname, GLfloat* param);
// void (APIENTRYP ptrglDrawElementsIndirect)(GLenum mode, GLenum type, GLvoid* indirect);
// void (APIENTRYP ptrglGetUniformiv)(GLuint program, GLint location, GLint* params);
// void (APIENTRYP ptrglGetRenderbufferParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglProgramParameteri)(GLuint program, GLenum pname, GLint value);
// void (APIENTRYP ptrglUniform2fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglProgramUniformMatrix2dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniform4fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglTexParameterf)(GLenum target, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglWaitSync)(GLsync sync, GLbitfield flags, GLuint64 timeout);
// void (APIENTRYP ptrglStencilOpSeparate)(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass);
// void (APIENTRYP ptrglVertexAttribP1uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrglVertexP2ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrglColorMaski)(GLuint index, GLboolean r, GLboolean g, GLboolean b, GLboolean a);
// void (APIENTRYP ptrglFramebufferTexture3D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint zoffset);
// void (APIENTRYP ptrglUniform2uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglGetActiveUniformBlockiv)(GLuint program, GLuint uniformBlockIndex, GLenum pname, GLint* params);
// void (APIENTRYP ptrglVertexAttribL4dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglStencilMask)(GLuint mask);
// void (APIENTRYP ptrglMultiDrawElementsIndirect)(GLenum mode, GLenum type, void* indirect, GLsizei drawcount, GLsizei stride);
// void (APIENTRYP ptrglGetActiveSubroutineName)(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name);
// void (APIENTRYP ptrglGetAttachedShaders)(GLuint program, GLsizei maxCount, GLsizei* count, GLuint* obj);
// void (APIENTRYP ptrglClearDepthf)(GLfloat d);
// void (APIENTRYP ptrglClearStencil)(GLint s);
// void (APIENTRYP ptrglGetCompressedTexImage)(GLenum target, GLint level, GLvoid* img);
// void (APIENTRYP ptrglPointParameteriv)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglProgramUniform2i)(GLuint program, GLint location, GLint v0, GLint v1);
// void (APIENTRYP ptrglGetTexImage)(GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglBindFragDataLocation)(GLuint program, GLuint color, GLchar* name);
// void (APIENTRYP ptrglUniformMatrix2x3dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglBlendFunci)(GLuint buf, GLenum src, GLenum dst);
// void (APIENTRYP ptrglBindAttribLocation)(GLuint program, GLuint index, GLchar* name);
// void (APIENTRYP ptrglCompressedTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglVertexAttribLPointer)(GLuint index, GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglTexCoordP4ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglPauseTransformFeedback)();
// void (APIENTRYP ptrglProgramUniformMatrix2x4dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglCompressedTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglEndTransformFeedback)();
// void (APIENTRYP ptrglUniform3dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglScissor)(GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglSampleMaski)(GLuint index, GLbitfield mask);
// void (APIENTRYP ptrglFinish)();
// void (APIENTRYP ptrglProgramUniform1iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglScissorIndexedv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglGetTexParameterIiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglUseProgram)(GLuint program);
// void (APIENTRYP ptrglInvalidateBufferData)(GLuint buffer);
// void (APIENTRYP ptrglScissorArrayv)(GLuint first, GLsizei count, GLint* v);
// void (APIENTRYP ptrglGetProgramiv)(GLuint program, GLenum pname, GLint* params);
// void (APIENTRYP ptrglDisable)(GLenum cap);
// GLuint (APIENTRYP ptrglGetUniformBlockIndex)(GLuint program, GLchar* uniformBlockName);
// void (APIENTRYP ptrglBlendEquation)(GLenum mode);
// void (APIENTRYP ptrglProgramUniformMatrix3fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglDrawArrays)(GLenum mode, GLint first, GLsizei count);
// GLuint (APIENTRYP ptrglCreateShaderProgramv)(GLenum type, GLsizei count, GLchar* const* strings);
// void (APIENTRYP ptrglGenTransformFeedbacks)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglUniform1i)(GLint location, GLint v0);
// void (APIENTRYP ptrglVertexAttribL4d)(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglVertexAttribI4iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglProgramUniformMatrix2fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglBlitFramebuffer)(GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
// void (APIENTRYP ptrglInvalidateSubFramebuffer)(GLenum target, GLsizei numAttachments, GLenum* attachments, GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglMultiTexCoordP3ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglShaderSource)(GLuint shader, GLsizei count, GLchar* const* string, GLint* length);
// void (APIENTRYP ptrglUniformMatrix4dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglProgramUniformMatrix4dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglDeleteSamplers)(GLsizei count, GLuint* samplers);
// void (APIENTRYP ptrglGenSamplers)(GLsizei count, GLuint* samplers);
// void (APIENTRYP ptrglTexParameterIiv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglAttachShader)(GLuint program, GLuint shader);
// void (APIENTRYP ptrglValidateProgramPipeline)(GLuint pipeline);
// void (APIENTRYP ptrglFramebufferRenderbuffer)(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer);
// void (APIENTRYP ptrglDetachShader)(GLuint program, GLuint shader);
// void (APIENTRYP ptrglUniform4d)(GLint location, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
// void (APIENTRYP ptrglDepthRange)(GLdouble near, GLdouble far);
// void (APIENTRYP ptrglTexBufferRange)(GLenum target, GLenum internalformat, GLuint buffer, GLintptr offset, GLsizeiptr size);
// void (APIENTRYP ptrglUniform4uiv)(GLint location, GLsizei count, GLuint* value);
// GLuint (APIENTRYP ptrglGetSubroutineIndex)(GLuint program, GLenum shadertype, GLchar* name);
// void (APIENTRYP ptrglVertexAttribI3iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglPatchParameteri)(GLenum pname, GLint value);
// void (APIENTRYP ptrglGetProgramInfoLog)(GLuint program, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrglFlushMappedBufferRange)(GLenum target, GLintptr offset, GLsizeiptr length);
// void (APIENTRYP ptrglBlendEquationSeparatei)(GLuint buf, GLenum modeRGB, GLenum modeAlpha);
// void (APIENTRYP ptrglUniform1ui)(GLint location, GLuint v0);
// void (APIENTRYP ptrglGetActiveSubroutineUniformName)(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name);
// void (APIENTRYP ptrglVertexAttribL2d)(GLuint index, GLdouble x, GLdouble y);
// GLboolean (APIENTRYP ptrglIsRenderbuffer)(GLuint renderbuffer);
// void (APIENTRYP ptrglProgramUniform3f)(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2);
// void (APIENTRYP ptrglViewportIndexedf)(GLuint index, GLfloat x, GLfloat y, GLfloat w, GLfloat h);
// void (APIENTRYP ptrglResumeTransformFeedback)();
// void (APIENTRYP ptrglMultiTexCoordP1ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglValidateProgram)(GLuint program);
// void (APIENTRYP ptrglPopDebugGroup)();
// void (APIENTRYP ptrglVertexAttribP4uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrglDrawArraysIndirect)(GLenum mode, GLvoid* indirect);
// void (APIENTRYP ptrglProgramUniformMatrix4x3dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglDrawElements)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices);
// void (APIENTRYP ptrglClearDepth)(GLdouble depth);
// void (APIENTRYP ptrglSamplerParameterf)(GLuint sampler, GLenum pname, GLfloat param);
// void (APIENTRYP ptrglBufferSubData)(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data);
// void (APIENTRYP ptrglBeginQuery)(GLenum target, GLuint id);
// void (APIENTRYP ptrglGetProgramPipelineiv)(GLuint pipeline, GLenum pname, GLint* params);
// void (APIENTRYP ptrglProgramBinary)(GLuint program, GLenum binaryFormat, GLvoid* binary, GLsizei length);
// void (APIENTRYP ptrglVertexAttribI4ubv)(GLuint index, GLubyte* v);
// void (APIENTRYP ptrglVertexAttribFormat)(GLuint attribindex, GLint size, GLenum type, GLboolean normalized, GLuint relativeoffset);
// GLboolean (APIENTRYP ptrglIsTransformFeedback)(GLuint id);
// GLint (APIENTRYP ptrglGetAttribLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglProgramUniform3iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// GLboolean (APIENTRYP ptrglUnmapBuffer)(GLenum target);
// void (APIENTRYP ptrglGenTextures)(GLsizei n, GLuint* textures);
// GLint (APIENTRYP ptrglGetProgramResourceLocationIndex)(GLuint program, GLenum programInterface, GLchar* name);
// void (APIENTRYP ptrglGetShaderSource)(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source);
// GLboolean (APIENTRYP ptrglIsFramebuffer)(GLuint framebuffer);
// void (APIENTRYP ptrglTexCoordP2uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglUniformMatrix4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglTexImage1D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglEnablei)(GLenum target, GLuint index);
// void (APIENTRYP ptrglProgramUniform2ui)(GLuint program, GLint location, GLuint v0, GLuint v1);
// void (APIENTRYP ptrglColorP3uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrglGetFramebufferParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglDeleteBuffers)(GLsizei n, GLuint* buffers);
// void (APIENTRYP ptrglMultiTexCoordP2ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglVertexAttribL2dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglDrawTransformFeedbackStream)(GLenum mode, GLuint id, GLuint stream);
// void (APIENTRYP ptrglDeleteProgramPipelines)(GLsizei n, GLuint* pipelines);
// void (APIENTRYP ptrglProgramUniform2dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglProgramUniform2fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglGetMultisamplefv)(GLenum pname, GLuint index, GLfloat* val);
// void (APIENTRYP ptrglVertexP3uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrglBlendFuncSeparatei)(GLuint buf, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha);
// void (APIENTRYP ptrglGetActiveUniformName)(GLuint program, GLuint uniformIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformName);
// void (APIENTRYP ptrglTexCoordP1ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglInvalidateTexImage)(GLuint texture, GLint level);
// GLenum (APIENTRYP ptrglCheckFramebufferStatus)(GLenum target);
// void (APIENTRYP ptrglMultiTexCoordP3uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglBindVertexArray)(GLuint array);
// void (APIENTRYP ptrglColorP3ui)(GLenum type, GLuint color);
// void (APIENTRYP ptrglUniform2i)(GLint location, GLint v0, GLint v1);
// void (APIENTRYP ptrglGetActiveAttrib)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrglBeginConditionalRender)(GLuint id, GLenum mode);
// void (APIENTRYP ptrglProgramUniform3d)(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2);
// void (APIENTRYP ptrglReadPixels)(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglVertexAttribL1d)(GLuint index, GLdouble x);
// void (APIENTRYP ptrglVertexAttribPointer)(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglGetVertexAttribIuiv)(GLuint index, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglProgramUniform3ui)(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2);
// void (APIENTRYP ptrglProgramUniform2iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglSamplerParameteriv)(GLuint sampler, GLenum pname, GLint* param);
// void (APIENTRYP ptrglVertexAttribI4sv)(GLuint index, GLshort* v);
// void (APIENTRYP ptrglProgramUniformMatrix4x2fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglBindFramebuffer)(GLenum target, GLuint framebuffer);
// void (APIENTRYP ptrglProgramUniformMatrix2x3dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniformMatrix2x4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglCopyTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border);
// void (APIENTRYP ptrglVertexP3ui)(GLenum type, GLuint value);
// void (APIENTRYP ptrglGetActiveSubroutineUniformiv)(GLuint program, GLenum shadertype, GLuint index, GLenum pname, GLint* values);
// void (APIENTRYP ptrglProgramUniformMatrix3x4dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglClearBufferfi)(GLenum buffer, GLint drawbuffer, GLfloat depth, GLint stencil);
// void (APIENTRYP ptrglProgramUniform1fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglBindTransformFeedback)(GLenum target, GLuint id);
// void (APIENTRYP ptrglGenerateMipmap)(GLenum target);
// void (APIENTRYP ptrglCompressedTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglGetSamplerParameteriv)(GLuint sampler, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGenProgramPipelines)(GLsizei n, GLuint* pipelines);
// GLvoid* (APIENTRYP ptrglMapBufferRange)(GLenum target, GLintptr offset, GLsizeiptr length, GLbitfield access);
// void (APIENTRYP ptrglMultiDrawArrays)(GLenum mode, GLint* first, GLsizei* count, GLsizei drawcount);
// void (APIENTRYP ptrglProgramUniform4f)(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3);
// void (APIENTRYP ptrglDeleteTransformFeedbacks)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglTextureView)(GLuint texture, GLenum target, GLuint origtexture, GLenum internalformat, GLuint minlevel, GLuint numlevels, GLuint minlayer, GLuint numlayers);
// void (APIENTRYP ptrglUniform1d)(GLint location, GLdouble x);
// void (APIENTRYP ptrglVertexAttribIFormat)(GLuint attribindex, GLint size, GLenum type, GLuint relativeoffset);
// void (APIENTRYP ptrglUniform4ui)(GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3);
// void (APIENTRYP ptrglUniform3iv)(GLint location, GLsizei count, GLint* value);
// GLboolean (APIENTRYP ptrglIsQuery)(GLuint id);
// void (APIENTRYP ptrglGetDoublei_v)(GLenum target, GLuint index, GLdouble* data);
// void (APIENTRYP ptrglViewport)(GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglBlendFunc)(GLenum sfactor, GLenum dfactor);
// void (APIENTRYP ptrglUniform4i)(GLint location, GLint v0, GLint v1, GLint v2, GLint v3);
// void (APIENTRYP ptrglDeleteTextures)(GLsizei n, GLuint* textures);
// void (APIENTRYP ptrglSecondaryColorP3ui)(GLenum type, GLuint color);
// void (APIENTRYP ptrglFramebufferTextureLayer)(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer);
// void (APIENTRYP ptrglGetTransformFeedbackVarying)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLsizei* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrglDebugMessageCallback)(GLDEBUGPROC callback, void* userParam);
// void (APIENTRYP ptrglCompressedTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglGetQueryObjectiv)(GLuint id, GLenum pname, GLint* params);
// void (APIENTRYP ptrglDeleteQueries)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglDrawTransformFeedbackStreamInstanced)(GLenum mode, GLuint id, GLuint stream, GLsizei instancecount);
// void (APIENTRYP ptrglVertexAttribI2iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglCopyTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border);
// void (APIENTRYP ptrglUniform1iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglStencilMaskSeparate)(GLenum face, GLuint mask);
// GLsync (APIENTRYP ptrglFenceSync)(GLenum condition, GLbitfield flags);
// void (APIENTRYP ptrglReleaseShaderCompiler)();
// void (APIENTRYP ptrglDrawElementsInstancedBaseInstance)(GLenum mode, GLsizei count, GLenum type, void* indices, GLsizei instancecount, GLuint baseinstance);
// void (APIENTRYP ptrglProgramUniformMatrix3dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglFlush)();
// void (APIENTRYP ptrglBindBufferBase)(GLenum target, GLuint index, GLuint buffer);
// void (APIENTRYP ptrglEnableVertexAttribArray)(GLuint index);
// GLint (APIENTRYP ptrglGetProgramResourceLocation)(GLuint program, GLenum programInterface, GLchar* name);
// void (APIENTRYP ptrglPointSize)(GLfloat size);
// void (APIENTRYP ptrglFramebufferTexture1D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
// void (APIENTRYP ptrglReadBuffer)(GLenum mode);
// void (APIENTRYP ptrglDrawArraysInstancedBaseInstance)(GLenum mode, GLint first, GLsizei count, GLsizei instancecount, GLuint baseinstance);
// void (APIENTRYP ptrglDepthRangef)(GLfloat n, GLfloat f);
// void (APIENTRYP ptrglHint)(GLenum target, GLenum mode);
// void (APIENTRYP ptrglVertexAttribP3ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglGetVertexAttribPointerv)(GLuint index, GLenum pname, GLvoid** pointer);
// void (APIENTRYP ptrglGetBooleani_v)(GLenum target, GLuint index, GLboolean* data);
// void (APIENTRYP ptrglTexParameteri)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrglVertexAttribI3i)(GLuint index, GLint x, GLint y, GLint z);
// void (APIENTRYP ptrglVertexAttribL3d)(GLuint index, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglGetUniformfv)(GLuint program, GLint location, GLfloat* params);
// void (APIENTRYP ptrglCullFace)(GLenum mode);
// GLboolean (APIENTRYP ptrglIsSync)(GLsync sync);
// void (APIENTRYP ptrglDrawElementsBaseVertex)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLint basevertex);
// void (APIENTRYP ptrglProgramUniform1ui)(GLuint program, GLint location, GLuint v0);
// void (APIENTRYP ptrglProgramUniformMatrix3x2fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglVertexAttribP4ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglGetSamplerParameterIuiv)(GLuint sampler, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglLinkProgram)(GLuint program);
// void (APIENTRYP ptrglGetBufferParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglTexParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglMultiDrawArraysIndirect)(GLenum mode, void* indirect, GLsizei drawcount, GLsizei stride);
// void (APIENTRYP ptrglGetQueryObjectui64v)(GLuint id, GLenum pname, GLuint64* params);
// void (APIENTRYP ptrglGenVertexArrays)(GLsizei n, GLuint* arrays);
// void (APIENTRYP ptrglVertexAttribBinding)(GLuint attribindex, GLuint bindingindex);
// void (APIENTRYP ptrglPointParameterf)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglProgramUniform1uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglUniformMatrix4x3dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// GLuint (APIENTRYP ptrglCreateShader)(GLenum type);
// void (APIENTRYP ptrglGetTexParameteriv)(GLenum target, GLenum pname, GLint* params);
// void (APIENTRYP ptrglUniformMatrix4x2dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglVertexAttribI1uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglPolygonMode)(GLenum face, GLenum mode);
// void (APIENTRYP ptrglGetProgramResourceName)(GLuint program, GLenum programInterface, GLuint index, GLsizei bufSize, GLsizei* length, GLchar* name);
// void (APIENTRYP ptrglDrawElementsInstancedBaseVertex)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei instancecount, GLint basevertex);
// void (APIENTRYP ptrglTexImage3DMultisample)(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations);
// GLboolean (APIENTRYP ptrglIsVertexArray)(GLuint array);
// void (APIENTRYP ptrglUniform3fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglTexCoordP2ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglProgramUniform1f)(GLuint program, GLint location, GLfloat v0);
// void (APIENTRYP ptrglProgramUniformMatrix4fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglGenQueries)(GLsizei n, GLuint* ids);
// void (APIENTRYP ptrglGetQueryObjectuiv)(GLuint id, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglGetObjectLabel)(GLenum identifier, GLuint name, GLsizei bufSize, GLsizei* length, GLchar* label);
// void (APIENTRYP ptrglProgramUniform4ui)(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3);
// GLboolean (APIENTRYP ptrglIsProgram)(GLuint program);
// void (APIENTRYP ptrglObjectPtrLabel)(void* ptr, GLsizei length, GLchar* label);
// void (APIENTRYP ptrglDeleteShader)(GLuint shader);
// void (APIENTRYP ptrglMultiDrawElementsBaseVertex)(GLenum mode, GLsizei* count, GLenum type, GLvoid* const* indices, GLsizei drawcount, GLint* basevertex);
// void (APIENTRYP ptrglInvalidateFramebuffer)(GLenum target, GLsizei numAttachments, GLenum* attachments);
// void (APIENTRYP ptrglGetActiveUniformsiv)(GLuint program, GLsizei uniformCount, GLuint* uniformIndices, GLenum pname, GLint* params);
// void (APIENTRYP ptrglUniform3ui)(GLint location, GLuint v0, GLuint v1, GLuint v2);
// void (APIENTRYP ptrglUniformMatrix3x2dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglGetShaderInfoLog)(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrglBlendColor)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglBindFragDataLocationIndexed)(GLuint program, GLuint colorNumber, GLuint index, GLchar* name);
// void (APIENTRYP ptrglGetFramebufferAttachmentParameteriv)(GLenum target, GLenum attachment, GLenum pname, GLint* params);
// void (APIENTRYP ptrglVertexAttribP3uiv)(GLuint index, GLenum type, GLboolean normalized, GLuint* value);
// void (APIENTRYP ptrglDrawElementsInstancedBaseVertexBaseInstance)(GLenum mode, GLsizei count, GLenum type, void* indices, GLsizei instancecount, GLint basevertex, GLuint baseinstance);
// void (APIENTRYP ptrglUniformSubroutinesuiv)(GLenum shadertype, GLsizei count, GLuint* indices);
// void (APIENTRYP ptrglUniform3uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglVertexAttribI4bv)(GLuint index, GLbyte* v);
// void (APIENTRYP ptrglProgramUniform3i)(GLuint program, GLint location, GLint v0, GLint v1, GLint v2);
// void (APIENTRYP ptrglProgramUniform3dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// const GLubyte * (APIENTRYP ptrglGetStringi)(GLenum name, GLuint index);
// void (APIENTRYP ptrglGetQueryObjecti64v)(GLuint id, GLenum pname, GLint64* params);
// void (APIENTRYP ptrglVertexAttribL1dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglGetInteger64i_v)(GLenum target, GLuint index, GLint64* data);
// void (APIENTRYP ptrglGetPointerv)(GLenum pname, GLvoid** params);
// void (APIENTRYP ptrglActiveShaderProgram)(GLuint pipeline, GLuint program);
// void (APIENTRYP ptrglUniformMatrix2x4dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglSampleCoverage)(GLfloat value, GLboolean invert);
// void (APIENTRYP ptrglVertexP2uiv)(GLenum type, GLuint* value);
// void (APIENTRYP ptrglDrawTransformFeedback)(GLenum mode, GLuint id);
// void (APIENTRYP ptrglGetIntegerv)(GLenum pname, GLint* params);
// void (APIENTRYP ptrglDeleteProgram)(GLuint program);
// void (APIENTRYP ptrglProgramUniform4uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglUniform1uiv)(GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglGetDoublev)(GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglProgramUniform3fv)(GLuint program, GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglGetActiveUniformBlockName)(GLuint program, GLuint uniformBlockIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformBlockName);
// void (APIENTRYP ptrglGetSamplerParameterIiv)(GLuint sampler, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetInternalformativ)(GLenum target, GLenum internalformat, GLenum pname, GLsizei bufSize, GLint* params);
// void (APIENTRYP ptrglGetBufferPointerv)(GLenum target, GLenum pname, GLvoid** params);
// void (APIENTRYP ptrglClear)(GLbitfield mask);
// void (APIENTRYP ptrglGetVertexAttribLdv)(GLuint index, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglProgramUniformMatrix2x4fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglPrimitiveRestartIndex)(GLuint index);
// void (APIENTRYP ptrglLineWidth)(GLfloat width);
// void (APIENTRYP ptrglGetTexLevelParameterfv)(GLenum target, GLint level, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetIntegeri_v)(GLenum target, GLuint index, GLint* data);
// void (APIENTRYP ptrglGetProgramBinary)(GLuint program, GLsizei bufSize, GLsizei* length, GLenum* binaryFormat, GLvoid* binary);
// GLboolean (APIENTRYP ptrglIsEnabledi)(GLenum target, GLuint index);
// void (APIENTRYP ptrglGetUniformIndices)(GLuint program, GLsizei uniformCount, GLchar* const* uniformNames, GLuint* uniformIndices);
// void (APIENTRYP ptrglClearBufferData)(GLenum target, GLenum internalformat, GLenum format, GLenum type, void* data);
// void (APIENTRYP ptrglCopyTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglProgramUniform3uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglTransformFeedbackVaryings)(GLuint program, GLsizei count, GLchar* const* varyings, GLenum bufferMode);
// void (APIENTRYP ptrglPointParameterfv)(GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglVertexAttribI1iv)(GLuint index, GLint* v);
// void (APIENTRYP ptrglTexStorage2D)(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglVertexAttribI4usv)(GLuint index, GLushort* v);
// void (APIENTRYP ptrglRenderbufferStorageMultisample)(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglBufferData)(GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage);
// void (APIENTRYP ptrglUniformBlockBinding)(GLuint program, GLuint uniformBlockIndex, GLuint uniformBlockBinding);
// void (APIENTRYP ptrglGetUniformuiv)(GLuint program, GLint location, GLuint* params);
// void (APIENTRYP ptrglGetInternalformati64v)(GLenum target, GLenum internalformat, GLenum pname, GLsizei bufSize, GLint64* params);
// void (APIENTRYP ptrglTexParameterIuiv)(GLenum target, GLenum pname, GLuint* params);
// void (APIENTRYP ptrglUniformMatrix3x4dv)(GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglGetFloati_v)(GLenum target, GLuint index, GLfloat* data);
// GLuint (APIENTRYP ptrglCreateProgram)();
// GLboolean (APIENTRYP ptrglIsProgramPipeline)(GLuint pipeline);
// void (APIENTRYP ptrglProgramUniform4dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglGetInteger64v)(GLenum pname, GLint64* params);
// void (APIENTRYP ptrglGetProgramInterfaceiv)(GLuint program, GLenum programInterface, GLenum pname, GLint* params);
// void (APIENTRYP ptrglNormalP3uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglShaderBinary)(GLsizei count, GLuint* shaders, GLenum binaryformat, GLvoid* binary, GLsizei length);
// void (APIENTRYP ptrglPatchParameterfv)(GLenum pname, GLfloat* values);
// void (APIENTRYP ptrglGetSamplerParameterfv)(GLuint sampler, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGetProgramPipelineInfoLog)(GLuint pipeline, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
// void (APIENTRYP ptrglObjectLabel)(GLenum identifier, GLuint name, GLsizei length, GLchar* label);
// void (APIENTRYP ptrglVertexAttribI4uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglCopyImageSubData)(GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei srcWidth, GLsizei srcHeight, GLsizei srcDepth);
// void (APIENTRYP ptrglBlendEquationi)(GLuint buf, GLenum mode);
// void (APIENTRYP ptrglProgramUniformMatrix3x2dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglUniform2iv)(GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglEnable)(GLenum cap);
// void (APIENTRYP ptrglUniform2dv)(GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglDrawArraysInstanced)(GLenum mode, GLint first, GLsizei count, GLsizei instancecount);
// void (APIENTRYP ptrglDepthFunc)(GLenum func);
// void (APIENTRYP ptrglColorP4uiv)(GLenum type, GLuint* color);
// void (APIENTRYP ptrglFramebufferParameteri)(GLenum target, GLenum pname, GLint param);
// void (APIENTRYP ptrglVertexAttribI3ui)(GLuint index, GLuint x, GLuint y, GLuint z);
// void (APIENTRYP ptrglQueryCounter)(GLuint id, GLenum target);
// void (APIENTRYP ptrglProgramUniformMatrix3x4fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglBlendFuncSeparate)(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha);
// void (APIENTRYP ptrglCopyBufferSubData)(GLenum readTarget, GLenum writeTarget, GLintptr readOffset, GLintptr writeOffset, GLsizeiptr size);
// void (APIENTRYP ptrglUniform3d)(GLint location, GLdouble x, GLdouble y, GLdouble z);
// void (APIENTRYP ptrglDrawRangeElements)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, GLvoid* indices);
// void (APIENTRYP ptrglVertexAttribIPointer)(GLuint index, GLint size, GLenum type, GLsizei stride, GLvoid* pointer);
// void (APIENTRYP ptrglSamplerParameterIuiv)(GLuint sampler, GLenum pname, GLuint* param);
// void (APIENTRYP ptrglGetActiveUniform)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
// void (APIENTRYP ptrglCopyTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height);
// void (APIENTRYP ptrglGetUniformdv)(GLuint program, GLint location, GLdouble* params);
// void (APIENTRYP ptrglVertexAttribI1ui)(GLuint index, GLuint x);
// void (APIENTRYP ptrglProgramUniform4iv)(GLuint program, GLint location, GLsizei count, GLint* value);
// void (APIENTRYP ptrglInvalidateTexSubImage)(GLuint texture, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth);
// void (APIENTRYP ptrglProgramUniformMatrix4x3fv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglColorP4ui)(GLenum type, GLuint color);
// void (APIENTRYP ptrglUniformMatrix2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglClearBufferfv)(GLenum buffer, GLint drawbuffer, GLfloat* value);
// void (APIENTRYP ptrglMultiDrawElements)(GLenum mode, GLsizei* count, GLenum type, GLvoid* const* indices, GLsizei drawcount);
// void (APIENTRYP ptrglTexImage3D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglClearBufferSubData)(GLenum target, GLenum internalformat, GLintptr offset, GLsizeiptr size, GLenum format, GLenum type, void* data);
// void (APIENTRYP ptrglGetActiveAtomicCounterBufferiv)(GLuint program, GLuint bufferIndex, GLenum pname, GLint* params);
// void (APIENTRYP ptrglViewportIndexedfv)(GLuint index, GLfloat* v);
// void (APIENTRYP ptrglFramebufferTexture2D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
// void (APIENTRYP ptrglBindBuffer)(GLenum target, GLuint buffer);
// void (APIENTRYP ptrglVertexAttribI3uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglVertexAttribI4i)(GLuint index, GLint x, GLint y, GLint z, GLint w);
// void (APIENTRYP ptrglCompressedTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglGetBufferSubData)(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data);
// void (APIENTRYP ptrglDeleteFramebuffers)(GLsizei n, GLuint* framebuffers);
// void (APIENTRYP ptrglVertexAttribLFormat)(GLuint attribindex, GLint size, GLenum type, GLuint relativeoffset);
// void (APIENTRYP ptrglUniformMatrix2x3fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// GLuint (APIENTRYP ptrglGetProgramResourceIndex)(GLuint program, GLenum programInterface, GLchar* name);
// void (APIENTRYP ptrglDisableVertexAttribArray)(GLuint index);
// void (APIENTRYP ptrglBindBufferRange)(GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size);
// void (APIENTRYP ptrglVertexAttribI2uiv)(GLuint index, GLuint* v);
// void (APIENTRYP ptrglGetShaderPrecisionFormat)(GLenum shadertype, GLenum precisiontype, GLint* range, GLint* precision);
// void (APIENTRYP ptrglVertexAttribP2ui)(GLuint index, GLenum type, GLboolean normalized, GLuint value);
// void (APIENTRYP ptrglProvokingVertex)(GLenum mode);
// void (APIENTRYP ptrglTexBuffer)(GLenum target, GLenum internalformat, GLuint buffer);
// const GLubyte * (APIENTRYP ptrglGetString)(GLenum name);
// void (APIENTRYP ptrglCompileShader)(GLuint shader);
// void (APIENTRYP ptrglUniform1f)(GLint location, GLfloat v0);
// void (APIENTRYP ptrglUniform3f)(GLint location, GLfloat v0, GLfloat v1, GLfloat v2);
// void (APIENTRYP ptrglGetObjectPtrLabel)(void* ptr, GLsizei bufSize, GLsizei* length, GLchar* label);
// void (APIENTRYP ptrglProgramUniform1i)(GLuint program, GLint location, GLint v0);
// void (APIENTRYP ptrglVertexAttribI2ui)(GLuint index, GLuint x, GLuint y);
// void (APIENTRYP ptrglProgramUniformMatrix4x2dv)(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value);
// void (APIENTRYP ptrglCompressedTexImage3D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data);
// void (APIENTRYP ptrglClearColor)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
// void (APIENTRYP ptrglTexCoordP3ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglSamplerParameterIiv)(GLuint sampler, GLenum pname, GLint* param);
// void (APIENTRYP ptrglDispatchComputeIndirect)(GLintptr indirect);
// void (APIENTRYP ptrglDeleteSync)(GLsync sync);
// void (APIENTRYP ptrglVertexAttribI2i)(GLuint index, GLint x, GLint y);
// void (APIENTRYP ptrglStencilFunc)(GLenum func, GLint ref, GLuint mask);
// GLuint (APIENTRYP ptrglGetDebugMessageLog)(GLuint count, GLsizei bufsize, GLenum* sources, GLenum* types, GLuint* ids, GLenum* severities, GLsizei* lengths, GLchar* messageLog);
// void (APIENTRYP ptrglGetUniformSubroutineuiv)(GLenum shadertype, GLint location, GLuint* params);
// GLint (APIENTRYP ptrglGetFragDataIndex)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglStencilFuncSeparate)(GLenum face, GLenum func, GLint ref, GLuint mask);
// void (APIENTRYP ptrglTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
// void (APIENTRYP ptrglPointParameteri)(GLenum pname, GLint param);
// GLboolean (APIENTRYP ptrglIsShader)(GLuint shader);
// void (APIENTRYP ptrglUniformMatrix4x2fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglGetQueryIndexediv)(GLenum target, GLuint index, GLenum pname, GLint* params);
// GLint (APIENTRYP ptrglGetUniformLocation)(GLuint program, GLchar* name);
// void (APIENTRYP ptrglActiveTexture)(GLenum texture);
// void (APIENTRYP ptrglMultiTexCoordP4ui)(GLenum texture, GLenum type, GLuint coords);
// void (APIENTRYP ptrglGetProgramStageiv)(GLuint program, GLenum shadertype, GLenum pname, GLint* values);
// void (APIENTRYP ptrglProgramUniform1dv)(GLuint program, GLint location, GLsizei count, GLdouble* value);
// void (APIENTRYP ptrglShaderStorageBlockBinding)(GLuint program, GLuint storageBlockIndex, GLuint storageBlockBinding);
// void (APIENTRYP ptrglTexStorage2DMultisample)(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrglBlendEquationSeparate)(GLenum modeRGB, GLenum modeAlpha);
// void (APIENTRYP ptrglNormalP3ui)(GLenum type, GLuint coords);
// void (APIENTRYP ptrglUniformMatrix3x4fv)(GLint location, GLsizei count, GLboolean transpose, GLfloat* value);
// void (APIENTRYP ptrglVertexAttribI1i)(GLuint index, GLint x);
// void (APIENTRYP ptrglUniform2ui)(GLint location, GLuint v0, GLuint v1);
// void (APIENTRYP ptrglGetTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
// void (APIENTRYP ptrglGenRenderbuffers)(GLsizei n, GLuint* renderbuffers);
// GLboolean (APIENTRYP ptrglIsSampler)(GLuint sampler);
// void (APIENTRYP ptrglEndQueryIndexed)(GLenum target, GLuint index);
// void (APIENTRYP ptrglPixelStoref)(GLenum pname, GLfloat param);
// void (APIENTRYP ptrglVertexAttribDivisor)(GLuint index, GLuint divisor);
// void (APIENTRYP ptrglStencilOp)(GLenum fail, GLenum zfail, GLenum zpass);
// void (APIENTRYP ptrglInvalidateBufferSubData)(GLuint buffer, GLintptr offset, GLsizeiptr length);
// void (APIENTRYP ptrglVertexAttribL3dv)(GLuint index, GLdouble* v);
// void (APIENTRYP ptrglClearBufferiv)(GLenum buffer, GLint drawbuffer, GLint* value);
// void (APIENTRYP ptrglLogicOp)(GLenum opcode);
// void (APIENTRYP ptrglCopyTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width);
// void (APIENTRYP ptrglGetVertexAttribdv)(GLuint index, GLenum pname, GLdouble* params);
// void (APIENTRYP ptrglDrawBuffer)(GLenum mode);
// void (APIENTRYP ptrglGenBuffers)(GLsizei n, GLuint* buffers);
// void (APIENTRYP ptrglTexStorage3D)(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth);
// void (APIENTRYP ptrglTexStorage3DMultisample)(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations);
// void (APIENTRYP ptrglDrawElementsInstanced)(GLenum mode, GLsizei count, GLenum type, GLvoid* indices, GLsizei instancecount);
// void (APIENTRYP ptrglPolygonOffset)(GLfloat factor, GLfloat units);
// GLvoid* (APIENTRYP ptrglMapBuffer)(GLenum target, GLenum access);
// GLint (APIENTRYP ptrglGetFragDataLocation)(GLuint program, GLchar* name);
// GLboolean (APIENTRYP ptrglIsTexture)(GLuint texture);
// void (APIENTRYP ptrglDeleteVertexArrays)(GLsizei n, GLuint* arrays);
// void (APIENTRYP ptrglViewportArrayv)(GLuint first, GLsizei count, GLfloat* v);
// void (APIENTRYP ptrglDrawTransformFeedbackInstanced)(GLenum mode, GLuint id, GLsizei instancecount);
// void (APIENTRYP ptrglDebugMessageInsert)(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, GLchar* buf);
// void (APIENTRYP ptrglGetVertexAttribiv)(GLuint index, GLenum pname, GLint* params);
// void (APIENTRYP ptrglGetFloatv)(GLenum pname, GLfloat* params);
// GLenum (APIENTRYP ptrglClientWaitSync)(GLsync sync, GLbitfield flags, GLuint64 timeout);
// void (APIENTRYP ptrglUniform2f)(GLint location, GLfloat v0, GLfloat v1);
// GLboolean (APIENTRYP ptrglIsBuffer)(GLuint buffer);
// void (APIENTRYP ptrglUniform1fv)(GLint location, GLsizei count, GLfloat* value);
// void (APIENTRYP ptrglMultiTexCoordP1uiv)(GLenum texture, GLenum type, GLuint* coords);
// void (APIENTRYP ptrglVertexAttribI4ui)(GLuint index, GLuint x, GLuint y, GLuint z, GLuint w);
// void (APIENTRYP ptrglGetBooleanv)(GLenum pname, GLboolean* params);
// void (APIENTRYP ptrglGetShaderiv)(GLuint shader, GLenum pname, GLint* params);
// void (APIENTRYP ptrglProgramUniform2uiv)(GLuint program, GLint location, GLsizei count, GLuint* value);
// void (APIENTRYP ptrglTexCoordP1uiv)(GLenum type, GLuint* coords);
// void (APIENTRYP ptrglEndConditionalRender)();
// 
// void callGetProgramBinary(GLuint program, GLsizei bufSize, GLsizei* length, GLenum* binaryFormat, GLvoid* binary) {  (*ptrglGetProgramBinary)(program, bufSize, length, binaryFormat, binary); }
// GLboolean callIsEnabledi(GLenum target, GLuint index) { return (*ptrglIsEnabledi)(target, index); }
// void callGetUniformIndices(GLuint program, GLsizei uniformCount, GLchar* const* uniformNames, GLuint* uniformIndices) {  (*ptrglGetUniformIndices)(program, uniformCount, uniformNames, uniformIndices); }
// void callClearBufferData(GLenum target, GLenum internalformat, GLenum format, GLenum type_, void* data) {  (*ptrglClearBufferData)(target, internalformat, format, type_, data); }
// void callCopyTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {  (*ptrglCopyTexSubImage3D)(target, level, xoffset, yoffset, zoffset, x, y, width, height); }
// void callProgramUniform3uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {  (*ptrglProgramUniform3uiv)(program, location, count, value); }
// void callTransformFeedbackVaryings(GLuint program, GLsizei count, GLchar* const* varyings, GLenum bufferMode) {  (*ptrglTransformFeedbackVaryings)(program, count, varyings, bufferMode); }
// void callPointParameterfv(GLenum pname, GLfloat* params) {  (*ptrglPointParameterfv)(pname, params); }
// void callVertexAttribI1iv(GLuint index, GLint* v) {  (*ptrglVertexAttribI1iv)(index, v); }
// void callTexStorage2D(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height) {  (*ptrglTexStorage2D)(target, levels, internalformat, width, height); }
// void callVertexAttribI4usv(GLuint index, GLushort* v) {  (*ptrglVertexAttribI4usv)(index, v); }
// void callRenderbufferStorageMultisample(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height) {  (*ptrglRenderbufferStorageMultisample)(target, samples, internalformat, width, height); }
// void callBufferData(GLenum target, GLsizeiptr size, GLvoid* data, GLenum usage) {  (*ptrglBufferData)(target, size, data, usage); }
// void callUniformBlockBinding(GLuint program, GLuint uniformBlockIndex, GLuint uniformBlockBinding) {  (*ptrglUniformBlockBinding)(program, uniformBlockIndex, uniformBlockBinding); }
// void callGetUniformuiv(GLuint program, GLint location, GLuint* params) {  (*ptrglGetUniformuiv)(program, location, params); }
// void callGetInternalformati64v(GLenum target, GLenum internalformat, GLenum pname, GLsizei bufSize, GLint64* params) {  (*ptrglGetInternalformati64v)(target, internalformat, pname, bufSize, params); }
// void callTexParameterIuiv(GLenum target, GLenum pname, GLuint* params) {  (*ptrglTexParameterIuiv)(target, pname, params); }
// void callUniformMatrix3x4dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix3x4dv)(location, count, transpose, value); }
// void callGetFloati_v(GLenum target, GLuint index, GLfloat* data) {  (*ptrglGetFloati_v)(target, index, data); }
// GLuint callCreateProgram() { return (*ptrglCreateProgram)(); }
// GLboolean callIsProgramPipeline(GLuint pipeline) { return (*ptrglIsProgramPipeline)(pipeline); }
// void callProgramUniform4dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {  (*ptrglProgramUniform4dv)(program, location, count, value); }
// void callGetInteger64v(GLenum pname, GLint64* params) {  (*ptrglGetInteger64v)(pname, params); }
// void callGetProgramInterfaceiv(GLuint program, GLenum programInterface, GLenum pname, GLint* params) {  (*ptrglGetProgramInterfaceiv)(program, programInterface, pname, params); }
// void callNormalP3uiv(GLenum type_, GLuint* coords) {  (*ptrglNormalP3uiv)(type_, coords); }
// void callShaderBinary(GLsizei count, GLuint* shaders, GLenum binaryformat, GLvoid* binary, GLsizei length) {  (*ptrglShaderBinary)(count, shaders, binaryformat, binary, length); }
// void callPatchParameterfv(GLenum pname, GLfloat* values) {  (*ptrglPatchParameterfv)(pname, values); }
// void callGetSamplerParameterfv(GLuint sampler, GLenum pname, GLfloat* params) {  (*ptrglGetSamplerParameterfv)(sampler, pname, params); }
// void callGetProgramPipelineInfoLog(GLuint pipeline, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {  (*ptrglGetProgramPipelineInfoLog)(pipeline, bufSize, length, infoLog); }
// void callObjectLabel(GLenum identifier, GLuint name, GLsizei length, GLchar* label) {  (*ptrglObjectLabel)(identifier, name, length, label); }
// void callVertexAttribI4uiv(GLuint index, GLuint* v) {  (*ptrglVertexAttribI4uiv)(index, v); }
// void callCopyImageSubData(GLuint srcName, GLenum srcTarget, GLint srcLevel, GLint srcX, GLint srcY, GLint srcZ, GLuint dstName, GLenum dstTarget, GLint dstLevel, GLint dstX, GLint dstY, GLint dstZ, GLsizei srcWidth, GLsizei srcHeight, GLsizei srcDepth) {  (*ptrglCopyImageSubData)(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth); }
// void callBlendEquationi(GLuint buf, GLenum mode) {  (*ptrglBlendEquationi)(buf, mode); }
// void callProgramUniformMatrix3x2dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix3x2dv)(program, location, count, transpose, value); }
// void callUniform2iv(GLint location, GLsizei count, GLint* value) {  (*ptrglUniform2iv)(location, count, value); }
// void callEnable(GLenum cap) {  (*ptrglEnable)(cap); }
// void callUniform2dv(GLint location, GLsizei count, GLdouble* value) {  (*ptrglUniform2dv)(location, count, value); }
// void callDrawArraysInstanced(GLenum mode, GLint first, GLsizei count, GLsizei instancecount) {  (*ptrglDrawArraysInstanced)(mode, first, count, instancecount); }
// void callDepthFunc(GLenum func_) {  (*ptrglDepthFunc)(func_); }
// void callColorP4uiv(GLenum type_, GLuint* color) {  (*ptrglColorP4uiv)(type_, color); }
// void callFramebufferParameteri(GLenum target, GLenum pname, GLint param) {  (*ptrglFramebufferParameteri)(target, pname, param); }
// void callVertexAttribI3ui(GLuint index, GLuint x, GLuint y, GLuint z) {  (*ptrglVertexAttribI3ui)(index, x, y, z); }
// void callQueryCounter(GLuint id, GLenum target) {  (*ptrglQueryCounter)(id, target); }
// void callProgramUniformMatrix3x4fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix3x4fv)(program, location, count, transpose, value); }
// void callBlendFuncSeparate(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha) {  (*ptrglBlendFuncSeparate)(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha); }
// void callCopyBufferSubData(GLenum readTarget, GLenum writeTarget, GLintptr readOffset, GLintptr writeOffset, GLsizeiptr size) {  (*ptrglCopyBufferSubData)(readTarget, writeTarget, readOffset, writeOffset, size); }
// void callUniform3d(GLint location, GLdouble x, GLdouble y, GLdouble z) {  (*ptrglUniform3d)(location, x, y, z); }
// void callDrawRangeElements(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type_, GLvoid* indices) {  (*ptrglDrawRangeElements)(mode, start, end, count, type_, indices); }
// void callVertexAttribIPointer(GLuint index, GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {  (*ptrglVertexAttribIPointer)(index, size, type_, stride, pointer); }
// void callSamplerParameterIuiv(GLuint sampler, GLenum pname, GLuint* param) {  (*ptrglSamplerParameterIuiv)(sampler, pname, param); }
// void callGetActiveUniform(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type_, GLchar* name) {  (*ptrglGetActiveUniform)(program, index, bufSize, length, size, type_, name); }
// void callCopyTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {  (*ptrglCopyTexSubImage2D)(target, level, xoffset, yoffset, x, y, width, height); }
// void callGetUniformdv(GLuint program, GLint location, GLdouble* params) {  (*ptrglGetUniformdv)(program, location, params); }
// void callVertexAttribI1ui(GLuint index, GLuint x) {  (*ptrglVertexAttribI1ui)(index, x); }
// void callProgramUniform4iv(GLuint program, GLint location, GLsizei count, GLint* value) {  (*ptrglProgramUniform4iv)(program, location, count, value); }
// void callInvalidateTexSubImage(GLuint texture, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth) {  (*ptrglInvalidateTexSubImage)(texture, level, xoffset, yoffset, zoffset, width, height, depth); }
// void callProgramUniformMatrix4x3fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix4x3fv)(program, location, count, transpose, value); }
// void callColorP4ui(GLenum type_, GLuint color) {  (*ptrglColorP4ui)(type_, color); }
// void callUniformMatrix2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix2fv)(location, count, transpose, value); }
// void callClearBufferfv(GLenum buffer, GLint drawbuffer, GLfloat* value) {  (*ptrglClearBufferfv)(buffer, drawbuffer, value); }
// void callMultiDrawElements(GLenum mode, GLsizei* count, GLenum type_, GLvoid* const* indices, GLsizei drawcount) {  (*ptrglMultiDrawElements)(mode, count, type_, indices, drawcount); }
// void callTexImage3D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexImage3D)(target, level, internalformat, width, height, depth, border, format, type_, pixels); }
// void callClearBufferSubData(GLenum target, GLenum internalformat, GLintptr offset, GLsizeiptr size, GLenum format, GLenum type_, void* data) {  (*ptrglClearBufferSubData)(target, internalformat, offset, size, format, type_, data); }
// void callGetActiveAtomicCounterBufferiv(GLuint program, GLuint bufferIndex, GLenum pname, GLint* params) {  (*ptrglGetActiveAtomicCounterBufferiv)(program, bufferIndex, pname, params); }
// void callViewportIndexedfv(GLuint index, GLfloat* v) {  (*ptrglViewportIndexedfv)(index, v); }
// void callFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level) {  (*ptrglFramebufferTexture2D)(target, attachment, textarget, texture, level); }
// void callBindBuffer(GLenum target, GLuint buffer) {  (*ptrglBindBuffer)(target, buffer); }
// void callVertexAttribI3uiv(GLuint index, GLuint* v) {  (*ptrglVertexAttribI3uiv)(index, v); }
// void callVertexAttribI4i(GLuint index, GLint x, GLint y, GLint z, GLint w) {  (*ptrglVertexAttribI4i)(index, x, y, z, w); }
// void callCompressedTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexSubImage1D)(target, level, xoffset, width, format, imageSize, data); }
// void callGetBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {  (*ptrglGetBufferSubData)(target, offset, size, data); }
// void callDeleteFramebuffers(GLsizei n, GLuint* framebuffers) {  (*ptrglDeleteFramebuffers)(n, framebuffers); }
// void callVertexAttribLFormat(GLuint attribindex, GLint size, GLenum type_, GLuint relativeoffset) {  (*ptrglVertexAttribLFormat)(attribindex, size, type_, relativeoffset); }
// void callUniformMatrix2x3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix2x3fv)(location, count, transpose, value); }
// GLuint callGetProgramResourceIndex(GLuint program, GLenum programInterface, GLchar* name) { return (*ptrglGetProgramResourceIndex)(program, programInterface, name); }
// void callDisableVertexAttribArray(GLuint index) {  (*ptrglDisableVertexAttribArray)(index); }
// void callBindBufferRange(GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size) {  (*ptrglBindBufferRange)(target, index, buffer, offset, size); }
// void callVertexAttribI2uiv(GLuint index, GLuint* v) {  (*ptrglVertexAttribI2uiv)(index, v); }
// void callGetShaderPrecisionFormat(GLenum shadertype, GLenum precisiontype, GLint* range_, GLint* precision) {  (*ptrglGetShaderPrecisionFormat)(shadertype, precisiontype, range_, precision); }
// void callVertexAttribP2ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {  (*ptrglVertexAttribP2ui)(index, type_, normalized, value); }
// void callProvokingVertex(GLenum mode) {  (*ptrglProvokingVertex)(mode); }
// void callTexBuffer(GLenum target, GLenum internalformat, GLuint buffer) {  (*ptrglTexBuffer)(target, internalformat, buffer); }
// const GLubyte * callGetString(GLenum name) { return (*ptrglGetString)(name); }
// void callCompileShader(GLuint shader) {  (*ptrglCompileShader)(shader); }
// void callUniform1f(GLint location, GLfloat v0) {  (*ptrglUniform1f)(location, v0); }
// void callUniform3f(GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {  (*ptrglUniform3f)(location, v0, v1, v2); }
// void callGetObjectPtrLabel(void* ptr, GLsizei bufSize, GLsizei* length, GLchar* label) {  (*ptrglGetObjectPtrLabel)(ptr, bufSize, length, label); }
// void callProgramUniform1i(GLuint program, GLint location, GLint v0) {  (*ptrglProgramUniform1i)(program, location, v0); }
// void callVertexAttribI2ui(GLuint index, GLuint x, GLuint y) {  (*ptrglVertexAttribI2ui)(index, x, y); }
// void callProgramUniformMatrix4x2dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix4x2dv)(program, location, count, transpose, value); }
// void callCompressedTexImage3D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexImage3D)(target, level, internalformat, width, height, depth, border, imageSize, data); }
// void callClearColor(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {  (*ptrglClearColor)(red, green, blue, alpha); }
// void callTexCoordP3ui(GLenum type_, GLuint coords) {  (*ptrglTexCoordP3ui)(type_, coords); }
// void callSamplerParameterIiv(GLuint sampler, GLenum pname, GLint* param) {  (*ptrglSamplerParameterIiv)(sampler, pname, param); }
// void callDispatchComputeIndirect(GLintptr indirect) {  (*ptrglDispatchComputeIndirect)(indirect); }
// void callDeleteSync(GLsync sync) {  (*ptrglDeleteSync)(sync); }
// void callVertexAttribI2i(GLuint index, GLint x, GLint y) {  (*ptrglVertexAttribI2i)(index, x, y); }
// void callStencilFunc(GLenum func_, GLint ref, GLuint mask) {  (*ptrglStencilFunc)(func_, ref, mask); }
// GLuint callGetDebugMessageLog(GLuint count, GLsizei bufsize, GLenum* sources, GLenum* types, GLuint* ids, GLenum* severities, GLsizei* lengths, GLchar* messageLog) { return (*ptrglGetDebugMessageLog)(count, bufsize, sources, types, ids, severities, lengths, messageLog); }
// void callGetUniformSubroutineuiv(GLenum shadertype, GLint location, GLuint* params) {  (*ptrglGetUniformSubroutineuiv)(shadertype, location, params); }
// GLint callGetFragDataIndex(GLuint program, GLchar* name) { return (*ptrglGetFragDataIndex)(program, name); }
// void callStencilFuncSeparate(GLenum face, GLenum func_, GLint ref, GLuint mask) {  (*ptrglStencilFuncSeparate)(face, func_, ref, mask); }
// void callTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexSubImage2D)(target, level, xoffset, yoffset, width, height, format, type_, pixels); }
// void callPointParameteri(GLenum pname, GLint param) {  (*ptrglPointParameteri)(pname, param); }
// GLboolean callIsShader(GLuint shader) { return (*ptrglIsShader)(shader); }
// void callUniformMatrix4x2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix4x2fv)(location, count, transpose, value); }
// void callGetQueryIndexediv(GLenum target, GLuint index, GLenum pname, GLint* params) {  (*ptrglGetQueryIndexediv)(target, index, pname, params); }
// GLint callGetUniformLocation(GLuint program, GLchar* name) { return (*ptrglGetUniformLocation)(program, name); }
// void callActiveTexture(GLenum texture) {  (*ptrglActiveTexture)(texture); }
// void callMultiTexCoordP4ui(GLenum texture, GLenum type_, GLuint coords) {  (*ptrglMultiTexCoordP4ui)(texture, type_, coords); }
// void callGetProgramStageiv(GLuint program, GLenum shadertype, GLenum pname, GLint* values) {  (*ptrglGetProgramStageiv)(program, shadertype, pname, values); }
// void callProgramUniform1dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {  (*ptrglProgramUniform1dv)(program, location, count, value); }
// void callShaderStorageBlockBinding(GLuint program, GLuint storageBlockIndex, GLuint storageBlockBinding) {  (*ptrglShaderStorageBlockBinding)(program, storageBlockIndex, storageBlockBinding); }
// void callTexStorage2DMultisample(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations) {  (*ptrglTexStorage2DMultisample)(target, samples, internalformat, width, height, fixedsamplelocations); }
// void callBlendEquationSeparate(GLenum modeRGB, GLenum modeAlpha) {  (*ptrglBlendEquationSeparate)(modeRGB, modeAlpha); }
// void callNormalP3ui(GLenum type_, GLuint coords) {  (*ptrglNormalP3ui)(type_, coords); }
// void callUniformMatrix3x4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix3x4fv)(location, count, transpose, value); }
// void callVertexAttribI1i(GLuint index, GLint x) {  (*ptrglVertexAttribI1i)(index, x); }
// void callUniform2ui(GLint location, GLuint v0, GLuint v1) {  (*ptrglUniform2ui)(location, v0, v1); }
// void callGetTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {  (*ptrglGetTexParameterfv)(target, pname, params); }
// void callGenRenderbuffers(GLsizei n, GLuint* renderbuffers) {  (*ptrglGenRenderbuffers)(n, renderbuffers); }
// GLboolean callIsSampler(GLuint sampler) { return (*ptrglIsSampler)(sampler); }
// void callEndQueryIndexed(GLenum target, GLuint index) {  (*ptrglEndQueryIndexed)(target, index); }
// void callPixelStoref(GLenum pname, GLfloat param) {  (*ptrglPixelStoref)(pname, param); }
// void callVertexAttribDivisor(GLuint index, GLuint divisor) {  (*ptrglVertexAttribDivisor)(index, divisor); }
// void callStencilOp(GLenum fail, GLenum zfail, GLenum zpass) {  (*ptrglStencilOp)(fail, zfail, zpass); }
// void callInvalidateBufferSubData(GLuint buffer, GLintptr offset, GLsizeiptr length) {  (*ptrglInvalidateBufferSubData)(buffer, offset, length); }
// void callVertexAttribL3dv(GLuint index, GLdouble* v) {  (*ptrglVertexAttribL3dv)(index, v); }
// void callClearBufferiv(GLenum buffer, GLint drawbuffer, GLint* value) {  (*ptrglClearBufferiv)(buffer, drawbuffer, value); }
// void callLogicOp(GLenum opcode) {  (*ptrglLogicOp)(opcode); }
// void callCopyTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {  (*ptrglCopyTexSubImage1D)(target, level, xoffset, x, y, width); }
// void callGetVertexAttribdv(GLuint index, GLenum pname, GLdouble* params) {  (*ptrglGetVertexAttribdv)(index, pname, params); }
// void callDrawBuffer(GLenum mode) {  (*ptrglDrawBuffer)(mode); }
// void callGenBuffers(GLsizei n, GLuint* buffers) {  (*ptrglGenBuffers)(n, buffers); }
// void callTexStorage3D(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth) {  (*ptrglTexStorage3D)(target, levels, internalformat, width, height, depth); }
// void callTexStorage3DMultisample(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations) {  (*ptrglTexStorage3DMultisample)(target, samples, internalformat, width, height, depth, fixedsamplelocations); }
// void callDrawElementsInstanced(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLsizei instancecount) {  (*ptrglDrawElementsInstanced)(mode, count, type_, indices, instancecount); }
// void callPolygonOffset(GLfloat factor, GLfloat units) {  (*ptrglPolygonOffset)(factor, units); }
// GLvoid* callMapBuffer(GLenum target, GLenum access) { return (*ptrglMapBuffer)(target, access); }
// GLint callGetFragDataLocation(GLuint program, GLchar* name) { return (*ptrglGetFragDataLocation)(program, name); }
// GLboolean callIsTexture(GLuint texture) { return (*ptrglIsTexture)(texture); }
// void callDeleteVertexArrays(GLsizei n, GLuint* arrays) {  (*ptrglDeleteVertexArrays)(n, arrays); }
// void callViewportArrayv(GLuint first, GLsizei count, GLfloat* v) {  (*ptrglViewportArrayv)(first, count, v); }
// void callDrawTransformFeedbackInstanced(GLenum mode, GLuint id, GLsizei instancecount) {  (*ptrglDrawTransformFeedbackInstanced)(mode, id, instancecount); }
// void callDebugMessageInsert(GLenum source, GLenum type_, GLuint id, GLenum severity, GLsizei length, GLchar* buf) {  (*ptrglDebugMessageInsert)(source, type_, id, severity, length, buf); }
// void callGetVertexAttribiv(GLuint index, GLenum pname, GLint* params) {  (*ptrglGetVertexAttribiv)(index, pname, params); }
// void callGetFloatv(GLenum pname, GLfloat* params) {  (*ptrglGetFloatv)(pname, params); }
// GLenum callClientWaitSync(GLsync sync, GLbitfield flags, GLuint64 timeout) { return (*ptrglClientWaitSync)(sync, flags, timeout); }
// void callUniform2f(GLint location, GLfloat v0, GLfloat v1) {  (*ptrglUniform2f)(location, v0, v1); }
// GLboolean callIsBuffer(GLuint buffer) { return (*ptrglIsBuffer)(buffer); }
// void callUniform1fv(GLint location, GLsizei count, GLfloat* value) {  (*ptrglUniform1fv)(location, count, value); }
// void callMultiTexCoordP1uiv(GLenum texture, GLenum type_, GLuint* coords) {  (*ptrglMultiTexCoordP1uiv)(texture, type_, coords); }
// void callVertexAttribI4ui(GLuint index, GLuint x, GLuint y, GLuint z, GLuint w) {  (*ptrglVertexAttribI4ui)(index, x, y, z, w); }
// void callGetBooleanv(GLenum pname, GLboolean* params) {  (*ptrglGetBooleanv)(pname, params); }
// void callGetShaderiv(GLuint shader, GLenum pname, GLint* params) {  (*ptrglGetShaderiv)(shader, pname, params); }
// void callProgramUniform2uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {  (*ptrglProgramUniform2uiv)(program, location, count, value); }
// void callTexCoordP1uiv(GLenum type_, GLuint* coords) {  (*ptrglTexCoordP1uiv)(type_, coords); }
// void callEndConditionalRender() {  (*ptrglEndConditionalRender)(); }
// void callBindImageTexture(GLuint unit, GLuint texture, GLint level, GLboolean layered, GLint layer, GLenum access, GLenum format) {  (*ptrglBindImageTexture)(unit, texture, level, layered, layer, access, format); }
// void callProgramUniform4i(GLuint program, GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {  (*ptrglProgramUniform4i)(program, location, v0, v1, v2, v3); }
// void callGetBufferParameteri64v(GLenum target, GLenum pname, GLint64* params) {  (*ptrglGetBufferParameteri64v)(target, pname, params); }
// void callRenderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height) {  (*ptrglRenderbufferStorage)(target, internalformat, width, height); }
// void callDepthRangeArrayv(GLuint first, GLsizei count, GLdouble* v) {  (*ptrglDepthRangeArrayv)(first, count, v); }
// void callFramebufferTexture(GLenum target, GLenum attachment, GLuint texture, GLint level) {  (*ptrglFramebufferTexture)(target, attachment, texture, level); }
// void callGetProgramResourceiv(GLuint program, GLenum programInterface, GLuint index, GLsizei propCount, GLenum* props, GLsizei bufSize, GLsizei* length, GLint* params) {  (*ptrglGetProgramResourceiv)(program, programInterface, index, propCount, props, bufSize, length, params); }
// void callMultiTexCoordP2uiv(GLenum texture, GLenum type_, GLuint* coords) {  (*ptrglMultiTexCoordP2uiv)(texture, type_, coords); }
// void callBindSampler(GLuint unit, GLuint sampler) {  (*ptrglBindSampler)(unit, sampler); }
// void callUniformMatrix4x3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix4x3fv)(location, count, transpose, value); }
// void callGetVertexAttribIiv(GLuint index, GLenum pname, GLint* params) {  (*ptrglGetVertexAttribIiv)(index, pname, params); }
// void callMemoryBarrier(GLbitfield barriers) {  (*ptrglMemoryBarrier)(barriers); }
// void callBeginQueryIndexed(GLenum target, GLuint index, GLuint id) {  (*ptrglBeginQueryIndexed)(target, index, id); }
// void callUniform2d(GLint location, GLdouble x, GLdouble y) {  (*ptrglUniform2d)(location, x, y); }
// void callDeleteRenderbuffers(GLsizei n, GLuint* renderbuffers) {  (*ptrglDeleteRenderbuffers)(n, renderbuffers); }
// void callUseProgramStages(GLuint pipeline, GLbitfield stages, GLuint program) {  (*ptrglUseProgramStages)(pipeline, stages, program); }
// void callDrawBuffers(GLsizei n, GLenum* bufs) {  (*ptrglDrawBuffers)(n, bufs); }
// void callGetSynciv(GLsync sync, GLenum pname, GLsizei bufSize, GLsizei* length, GLint* values) {  (*ptrglGetSynciv)(sync, pname, bufSize, length, values); }
// void callEndQuery(GLenum target) {  (*ptrglEndQuery)(target); }
// void callDepthRangeIndexed(GLuint index, GLdouble n, GLdouble f) {  (*ptrglDepthRangeIndexed)(index, n, f); }
// GLenum callGetError() { return (*ptrglGetError)(); }
// void callProgramUniform4d(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2, GLdouble v3) {  (*ptrglProgramUniform4d)(program, location, v0, v1, v2, v3); }
// void callPushDebugGroup(GLenum source, GLuint id, GLsizei length, GLchar* message) {  (*ptrglPushDebugGroup)(source, id, length, message); }
// void callGetQueryiv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetQueryiv)(target, pname, params); }
// GLboolean callIsEnabled(GLenum cap) { return (*ptrglIsEnabled)(cap); }
// void callSecondaryColorP3uiv(GLenum type_, GLuint* color) {  (*ptrglSecondaryColorP3uiv)(type_, color); }
// void callTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexSubImage3D)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type_, pixels); }
// void callDisablei(GLenum target, GLuint index) {  (*ptrglDisablei)(target, index); }
// void callDrawRangeElementsBaseVertex(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type_, GLvoid* indices, GLint basevertex) {  (*ptrglDrawRangeElementsBaseVertex)(mode, start, end, count, type_, indices, basevertex); }
// void callBindTexture(GLenum target, GLuint texture) {  (*ptrglBindTexture)(target, texture); }
// void callUniform4iv(GLint location, GLsizei count, GLint* value) {  (*ptrglUniform4iv)(location, count, value); }
// void callMinSampleShading(GLfloat value) {  (*ptrglMinSampleShading)(value); }
// void callUniform1dv(GLint location, GLsizei count, GLdouble* value) {  (*ptrglUniform1dv)(location, count, value); }
// void callVertexAttribP1ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {  (*ptrglVertexAttribP1ui)(index, type_, normalized, value); }
// void callPixelStorei(GLenum pname, GLint param) {  (*ptrglPixelStorei)(pname, param); }
// void callFrontFace(GLenum mode) {  (*ptrglFrontFace)(mode); }
// void callProgramUniform2d(GLuint program, GLint location, GLdouble v0, GLdouble v1) {  (*ptrglProgramUniform2d)(program, location, v0, v1); }
// void callTexStorage1D(GLenum target, GLsizei levels, GLenum internalformat, GLsizei width) {  (*ptrglTexStorage1D)(target, levels, internalformat, width); }
// void callUniform3i(GLint location, GLint v0, GLint v1, GLint v2) {  (*ptrglUniform3i)(location, v0, v1, v2); }
// void callTexImage2DMultisample(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLboolean fixedsamplelocations) {  (*ptrglTexImage2DMultisample)(target, samples, internalformat, width, height, fixedsamplelocations); }
// void callUniformMatrix3x2fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix3x2fv)(location, count, transpose, value); }
// void callTexImage2D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexImage2D)(target, level, internalformat, width, height, border, format, type_, pixels); }
// GLint callGetSubroutineUniformLocation(GLuint program, GLenum shadertype, GLchar* name) { return (*ptrglGetSubroutineUniformLocation)(program, shadertype, name); }
// void callUniformMatrix3dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix3dv)(location, count, transpose, value); }
// void callBindRenderbuffer(GLenum target, GLuint renderbuffer) {  (*ptrglBindRenderbuffer)(target, renderbuffer); }
// void callBindProgramPipeline(GLuint pipeline) {  (*ptrglBindProgramPipeline)(pipeline); }
// void callUniform4dv(GLint location, GLsizei count, GLdouble* value) {  (*ptrglUniform4dv)(location, count, value); }
// void callClearBufferuiv(GLenum buffer, GLint drawbuffer, GLuint* value) {  (*ptrglClearBufferuiv)(buffer, drawbuffer, value); }
// void callBeginTransformFeedback(GLenum primitiveMode) {  (*ptrglBeginTransformFeedback)(primitiveMode); }
// void callUniformMatrix3fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix3fv)(location, count, transpose, value); }
// void callDispatchCompute(GLuint num_groups_x, GLuint num_groups_y, GLuint num_groups_z) {  (*ptrglDispatchCompute)(num_groups_x, num_groups_y, num_groups_z); }
// void callBindVertexBuffer(GLuint bindingindex, GLuint buffer, GLintptr offset, GLsizei stride) {  (*ptrglBindVertexBuffer)(bindingindex, buffer, offset, stride); }
// void callVertexBindingDivisor(GLuint bindingindex, GLuint divisor) {  (*ptrglVertexBindingDivisor)(bindingindex, divisor); }
// void callMultiTexCoordP4uiv(GLenum texture, GLenum type_, GLuint* coords) {  (*ptrglMultiTexCoordP4uiv)(texture, type_, coords); }
// void callProgramUniformMatrix2x3fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix2x3fv)(program, location, count, transpose, value); }
// void callGetTexLevelParameteriv(GLenum target, GLint level, GLenum pname, GLint* params) {  (*ptrglGetTexLevelParameteriv)(target, level, pname, params); }
// void callVertexP4uiv(GLenum type_, GLuint* value) {  (*ptrglVertexP4uiv)(type_, value); }
// void callGetTexParameterIuiv(GLenum target, GLenum pname, GLuint* params) {  (*ptrglGetTexParameterIuiv)(target, pname, params); }
// void callDebugMessageControl(GLenum source, GLenum type_, GLenum severity, GLsizei count, GLuint* ids, GLboolean enabled) {  (*ptrglDebugMessageControl)(source, type_, severity, count, ids, enabled); }
// void callSamplerParameteri(GLuint sampler, GLenum pname, GLint param) {  (*ptrglSamplerParameteri)(sampler, pname, param); }
// void callProgramUniform2f(GLuint program, GLint location, GLfloat v0, GLfloat v1) {  (*ptrglProgramUniform2f)(program, location, v0, v1); }
// void callGetVertexAttribfv(GLuint index, GLenum pname, GLfloat* params) {  (*ptrglGetVertexAttribfv)(index, pname, params); }
// void callVertexAttribP2uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {  (*ptrglVertexAttribP2uiv)(index, type_, normalized, value); }
// void callScissorIndexed(GLuint index, GLint left, GLint bottom, GLsizei width, GLsizei height) {  (*ptrglScissorIndexed)(index, left, bottom, width, height); }
// void callProgramUniform1d(GLuint program, GLint location, GLdouble v0) {  (*ptrglProgramUniform1d)(program, location, v0); }
// void callVertexP4ui(GLenum type_, GLuint value) {  (*ptrglVertexP4ui)(type_, value); }
// void callDepthMask(GLboolean flag) {  (*ptrglDepthMask)(flag); }
// void callUniformMatrix2dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix2dv)(location, count, transpose, value); }
// void callTexCoordP4uiv(GLenum type_, GLuint* coords) {  (*ptrglTexCoordP4uiv)(type_, coords); }
// void callUniform4f(GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {  (*ptrglUniform4f)(location, v0, v1, v2, v3); }
// void callUniform4fv(GLint location, GLsizei count, GLfloat* value) {  (*ptrglUniform4fv)(location, count, value); }
// void callClampColor(GLenum target, GLenum clamp) {  (*ptrglClampColor)(target, clamp); }
// void callGenFramebuffers(GLsizei n, GLuint* framebuffers) {  (*ptrglGenFramebuffers)(n, framebuffers); }
// void callTexCoordP3uiv(GLenum type_, GLuint* coords) {  (*ptrglTexCoordP3uiv)(type_, coords); }
// void callColorMask(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {  (*ptrglColorMask)(red, green, blue, alpha); }
// void callSamplerParameterfv(GLuint sampler, GLenum pname, GLfloat* param) {  (*ptrglSamplerParameterfv)(sampler, pname, param); }
// void callDrawElementsIndirect(GLenum mode, GLenum type_, GLvoid* indirect) {  (*ptrglDrawElementsIndirect)(mode, type_, indirect); }
// void callGetUniformiv(GLuint program, GLint location, GLint* params) {  (*ptrglGetUniformiv)(program, location, params); }
// void callGetRenderbufferParameteriv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetRenderbufferParameteriv)(target, pname, params); }
// void callProgramParameteri(GLuint program, GLenum pname, GLint value) {  (*ptrglProgramParameteri)(program, pname, value); }
// void callUniform2fv(GLint location, GLsizei count, GLfloat* value) {  (*ptrglUniform2fv)(location, count, value); }
// void callProgramUniformMatrix2dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix2dv)(program, location, count, transpose, value); }
// void callProgramUniform4fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {  (*ptrglProgramUniform4fv)(program, location, count, value); }
// void callTexParameterf(GLenum target, GLenum pname, GLfloat param) {  (*ptrglTexParameterf)(target, pname, param); }
// void callWaitSync(GLsync sync, GLbitfield flags, GLuint64 timeout) {  (*ptrglWaitSync)(sync, flags, timeout); }
// void callStencilOpSeparate(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {  (*ptrglStencilOpSeparate)(face, sfail, dpfail, dppass); }
// void callVertexAttribP1uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {  (*ptrglVertexAttribP1uiv)(index, type_, normalized, value); }
// void callVertexP2ui(GLenum type_, GLuint value) {  (*ptrglVertexP2ui)(type_, value); }
// void callColorMaski(GLuint index, GLboolean r, GLboolean g, GLboolean b, GLboolean a) {  (*ptrglColorMaski)(index, r, g, b, a); }
// void callFramebufferTexture3D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint zoffset) {  (*ptrglFramebufferTexture3D)(target, attachment, textarget, texture, level, zoffset); }
// void callUniform2uiv(GLint location, GLsizei count, GLuint* value) {  (*ptrglUniform2uiv)(location, count, value); }
// void callGetActiveUniformBlockiv(GLuint program, GLuint uniformBlockIndex, GLenum pname, GLint* params) {  (*ptrglGetActiveUniformBlockiv)(program, uniformBlockIndex, pname, params); }
// void callVertexAttribL4dv(GLuint index, GLdouble* v) {  (*ptrglVertexAttribL4dv)(index, v); }
// void callStencilMask(GLuint mask) {  (*ptrglStencilMask)(mask); }
// void callMultiDrawElementsIndirect(GLenum mode, GLenum type_, void* indirect, GLsizei drawcount, GLsizei stride) {  (*ptrglMultiDrawElementsIndirect)(mode, type_, indirect, drawcount, stride); }
// void callGetActiveSubroutineName(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name) {  (*ptrglGetActiveSubroutineName)(program, shadertype, index, bufsize, length, name); }
// void callGetAttachedShaders(GLuint program, GLsizei maxCount, GLsizei* count, GLuint* obj) {  (*ptrglGetAttachedShaders)(program, maxCount, count, obj); }
// void callClearDepthf(GLfloat d) {  (*ptrglClearDepthf)(d); }
// void callClearStencil(GLint s) {  (*ptrglClearStencil)(s); }
// void callGetCompressedTexImage(GLenum target, GLint level, GLvoid* img) {  (*ptrglGetCompressedTexImage)(target, level, img); }
// void callPointParameteriv(GLenum pname, GLint* params) {  (*ptrglPointParameteriv)(pname, params); }
// void callProgramUniform2i(GLuint program, GLint location, GLint v0, GLint v1) {  (*ptrglProgramUniform2i)(program, location, v0, v1); }
// void callGetTexImage(GLenum target, GLint level, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglGetTexImage)(target, level, format, type_, pixels); }
// void callBindFragDataLocation(GLuint program, GLuint color, GLchar* name) {  (*ptrglBindFragDataLocation)(program, color, name); }
// void callUniformMatrix2x3dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix2x3dv)(location, count, transpose, value); }
// void callBlendFunci(GLuint buf, GLenum src, GLenum dst) {  (*ptrglBlendFunci)(buf, src, dst); }
// void callBindAttribLocation(GLuint program, GLuint index, GLchar* name) {  (*ptrglBindAttribLocation)(program, index, name); }
// void callCompressedTexImage2D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexImage2D)(target, level, internalformat, width, height, border, imageSize, data); }
// void callVertexAttribLPointer(GLuint index, GLint size, GLenum type_, GLsizei stride, GLvoid* pointer) {  (*ptrglVertexAttribLPointer)(index, size, type_, stride, pointer); }
// void callTexCoordP4ui(GLenum type_, GLuint coords) {  (*ptrglTexCoordP4ui)(type_, coords); }
// void callPauseTransformFeedback() {  (*ptrglPauseTransformFeedback)(); }
// void callProgramUniformMatrix2x4dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix2x4dv)(program, location, count, transpose, value); }
// void callCompressedTexImage1D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexImage1D)(target, level, internalformat, width, border, imageSize, data); }
// void callEndTransformFeedback() {  (*ptrglEndTransformFeedback)(); }
// void callUniform3dv(GLint location, GLsizei count, GLdouble* value) {  (*ptrglUniform3dv)(location, count, value); }
// void callScissor(GLint x, GLint y, GLsizei width, GLsizei height) {  (*ptrglScissor)(x, y, width, height); }
// void callTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexSubImage1D)(target, level, xoffset, width, format, type_, pixels); }
// void callSampleMaski(GLuint index, GLbitfield mask) {  (*ptrglSampleMaski)(index, mask); }
// void callFinish() {  (*ptrglFinish)(); }
// void callProgramUniform1iv(GLuint program, GLint location, GLsizei count, GLint* value) {  (*ptrglProgramUniform1iv)(program, location, count, value); }
// void callScissorIndexedv(GLuint index, GLint* v) {  (*ptrglScissorIndexedv)(index, v); }
// void callGetTexParameterIiv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetTexParameterIiv)(target, pname, params); }
// void callUseProgram(GLuint program) {  (*ptrglUseProgram)(program); }
// void callInvalidateBufferData(GLuint buffer) {  (*ptrglInvalidateBufferData)(buffer); }
// void callScissorArrayv(GLuint first, GLsizei count, GLint* v) {  (*ptrglScissorArrayv)(first, count, v); }
// void callGetProgramiv(GLuint program, GLenum pname, GLint* params) {  (*ptrglGetProgramiv)(program, pname, params); }
// void callDisable(GLenum cap) {  (*ptrglDisable)(cap); }
// GLuint callGetUniformBlockIndex(GLuint program, GLchar* uniformBlockName) { return (*ptrglGetUniformBlockIndex)(program, uniformBlockName); }
// void callBlendEquation(GLenum mode) {  (*ptrglBlendEquation)(mode); }
// void callProgramUniformMatrix3fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix3fv)(program, location, count, transpose, value); }
// void callDrawArrays(GLenum mode, GLint first, GLsizei count) {  (*ptrglDrawArrays)(mode, first, count); }
// GLuint callCreateShaderProgramv(GLenum type_, GLsizei count, GLchar* const* strings) { return (*ptrglCreateShaderProgramv)(type_, count, strings); }
// void callGenTransformFeedbacks(GLsizei n, GLuint* ids) {  (*ptrglGenTransformFeedbacks)(n, ids); }
// void callUniform1i(GLint location, GLint v0) {  (*ptrglUniform1i)(location, v0); }
// void callVertexAttribL4d(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {  (*ptrglVertexAttribL4d)(index, x, y, z, w); }
// void callVertexAttribI4iv(GLuint index, GLint* v) {  (*ptrglVertexAttribI4iv)(index, v); }
// void callProgramUniformMatrix2fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix2fv)(program, location, count, transpose, value); }
// void callBlitFramebuffer(GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter) {  (*ptrglBlitFramebuffer)(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter); }
// void callInvalidateSubFramebuffer(GLenum target, GLsizei numAttachments, GLenum* attachments, GLint x, GLint y, GLsizei width, GLsizei height) {  (*ptrglInvalidateSubFramebuffer)(target, numAttachments, attachments, x, y, width, height); }
// void callMultiTexCoordP3ui(GLenum texture, GLenum type_, GLuint coords) {  (*ptrglMultiTexCoordP3ui)(texture, type_, coords); }
// void callShaderSource(GLuint shader, GLsizei count, GLchar* const* string_, GLint* length) {  (*ptrglShaderSource)(shader, count, string_, length); }
// void callUniformMatrix4dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix4dv)(location, count, transpose, value); }
// void callProgramUniformMatrix4dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix4dv)(program, location, count, transpose, value); }
// void callDeleteSamplers(GLsizei count, GLuint* samplers) {  (*ptrglDeleteSamplers)(count, samplers); }
// void callGenSamplers(GLsizei count, GLuint* samplers) {  (*ptrglGenSamplers)(count, samplers); }
// void callTexParameterIiv(GLenum target, GLenum pname, GLint* params) {  (*ptrglTexParameterIiv)(target, pname, params); }
// void callAttachShader(GLuint program, GLuint shader) {  (*ptrglAttachShader)(program, shader); }
// void callValidateProgramPipeline(GLuint pipeline) {  (*ptrglValidateProgramPipeline)(pipeline); }
// void callFramebufferRenderbuffer(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer) {  (*ptrglFramebufferRenderbuffer)(target, attachment, renderbuffertarget, renderbuffer); }
// void callDetachShader(GLuint program, GLuint shader) {  (*ptrglDetachShader)(program, shader); }
// void callUniform4d(GLint location, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {  (*ptrglUniform4d)(location, x, y, z, w); }
// void callDepthRange(GLdouble near_, GLdouble far_) {  (*ptrglDepthRange)(near_, far_); }
// void callTexBufferRange(GLenum target, GLenum internalformat, GLuint buffer, GLintptr offset, GLsizeiptr size) {  (*ptrglTexBufferRange)(target, internalformat, buffer, offset, size); }
// void callUniform4uiv(GLint location, GLsizei count, GLuint* value) {  (*ptrglUniform4uiv)(location, count, value); }
// GLuint callGetSubroutineIndex(GLuint program, GLenum shadertype, GLchar* name) { return (*ptrglGetSubroutineIndex)(program, shadertype, name); }
// void callVertexAttribI3iv(GLuint index, GLint* v) {  (*ptrglVertexAttribI3iv)(index, v); }
// void callPatchParameteri(GLenum pname, GLint value) {  (*ptrglPatchParameteri)(pname, value); }
// void callGetProgramInfoLog(GLuint program, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {  (*ptrglGetProgramInfoLog)(program, bufSize, length, infoLog); }
// void callFlushMappedBufferRange(GLenum target, GLintptr offset, GLsizeiptr length) {  (*ptrglFlushMappedBufferRange)(target, offset, length); }
// void callBlendEquationSeparatei(GLuint buf, GLenum modeRGB, GLenum modeAlpha) {  (*ptrglBlendEquationSeparatei)(buf, modeRGB, modeAlpha); }
// void callUniform1ui(GLint location, GLuint v0) {  (*ptrglUniform1ui)(location, v0); }
// void callGetActiveSubroutineUniformName(GLuint program, GLenum shadertype, GLuint index, GLsizei bufsize, GLsizei* length, GLchar* name) {  (*ptrglGetActiveSubroutineUniformName)(program, shadertype, index, bufsize, length, name); }
// void callVertexAttribL2d(GLuint index, GLdouble x, GLdouble y) {  (*ptrglVertexAttribL2d)(index, x, y); }
// GLboolean callIsRenderbuffer(GLuint renderbuffer) { return (*ptrglIsRenderbuffer)(renderbuffer); }
// void callProgramUniform3f(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {  (*ptrglProgramUniform3f)(program, location, v0, v1, v2); }
// void callViewportIndexedf(GLuint index, GLfloat x, GLfloat y, GLfloat w, GLfloat h) {  (*ptrglViewportIndexedf)(index, x, y, w, h); }
// void callResumeTransformFeedback() {  (*ptrglResumeTransformFeedback)(); }
// void callMultiTexCoordP1ui(GLenum texture, GLenum type_, GLuint coords) {  (*ptrglMultiTexCoordP1ui)(texture, type_, coords); }
// void callValidateProgram(GLuint program) {  (*ptrglValidateProgram)(program); }
// void callPopDebugGroup() {  (*ptrglPopDebugGroup)(); }
// void callVertexAttribP4uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {  (*ptrglVertexAttribP4uiv)(index, type_, normalized, value); }
// void callDrawArraysIndirect(GLenum mode, GLvoid* indirect) {  (*ptrglDrawArraysIndirect)(mode, indirect); }
// void callProgramUniformMatrix4x3dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix4x3dv)(program, location, count, transpose, value); }
// void callDrawElements(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices) {  (*ptrglDrawElements)(mode, count, type_, indices); }
// void callClearDepth(GLdouble depth) {  (*ptrglClearDepth)(depth); }
// void callSamplerParameterf(GLuint sampler, GLenum pname, GLfloat param) {  (*ptrglSamplerParameterf)(sampler, pname, param); }
// void callBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {  (*ptrglBufferSubData)(target, offset, size, data); }
// void callBeginQuery(GLenum target, GLuint id) {  (*ptrglBeginQuery)(target, id); }
// void callGetProgramPipelineiv(GLuint pipeline, GLenum pname, GLint* params) {  (*ptrglGetProgramPipelineiv)(pipeline, pname, params); }
// void callProgramBinary(GLuint program, GLenum binaryFormat, GLvoid* binary, GLsizei length) {  (*ptrglProgramBinary)(program, binaryFormat, binary, length); }
// void callVertexAttribI4ubv(GLuint index, GLubyte* v) {  (*ptrglVertexAttribI4ubv)(index, v); }
// void callVertexAttribFormat(GLuint attribindex, GLint size, GLenum type_, GLboolean normalized, GLuint relativeoffset) {  (*ptrglVertexAttribFormat)(attribindex, size, type_, normalized, relativeoffset); }
// GLboolean callIsTransformFeedback(GLuint id) { return (*ptrglIsTransformFeedback)(id); }
// GLint callGetAttribLocation(GLuint program, GLchar* name) { return (*ptrglGetAttribLocation)(program, name); }
// void callProgramUniform3iv(GLuint program, GLint location, GLsizei count, GLint* value) {  (*ptrglProgramUniform3iv)(program, location, count, value); }
// GLboolean callUnmapBuffer(GLenum target) { return (*ptrglUnmapBuffer)(target); }
// void callGenTextures(GLsizei n, GLuint* textures) {  (*ptrglGenTextures)(n, textures); }
// GLint callGetProgramResourceLocationIndex(GLuint program, GLenum programInterface, GLchar* name) { return (*ptrglGetProgramResourceLocationIndex)(program, programInterface, name); }
// void callGetShaderSource(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {  (*ptrglGetShaderSource)(shader, bufSize, length, source); }
// GLboolean callIsFramebuffer(GLuint framebuffer) { return (*ptrglIsFramebuffer)(framebuffer); }
// void callTexCoordP2uiv(GLenum type_, GLuint* coords) {  (*ptrglTexCoordP2uiv)(type_, coords); }
// void callUniformMatrix4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix4fv)(location, count, transpose, value); }
// void callTexImage1D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglTexImage1D)(target, level, internalformat, width, border, format, type_, pixels); }
// void callEnablei(GLenum target, GLuint index) {  (*ptrglEnablei)(target, index); }
// void callProgramUniform2ui(GLuint program, GLint location, GLuint v0, GLuint v1) {  (*ptrglProgramUniform2ui)(program, location, v0, v1); }
// void callColorP3uiv(GLenum type_, GLuint* color) {  (*ptrglColorP3uiv)(type_, color); }
// void callGetFramebufferParameteriv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetFramebufferParameteriv)(target, pname, params); }
// void callDeleteBuffers(GLsizei n, GLuint* buffers) {  (*ptrglDeleteBuffers)(n, buffers); }
// void callMultiTexCoordP2ui(GLenum texture, GLenum type_, GLuint coords) {  (*ptrglMultiTexCoordP2ui)(texture, type_, coords); }
// void callTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {  (*ptrglTexParameterfv)(target, pname, params); }
// void callVertexAttribL2dv(GLuint index, GLdouble* v) {  (*ptrglVertexAttribL2dv)(index, v); }
// void callDrawTransformFeedbackStream(GLenum mode, GLuint id, GLuint stream) {  (*ptrglDrawTransformFeedbackStream)(mode, id, stream); }
// void callDeleteProgramPipelines(GLsizei n, GLuint* pipelines) {  (*ptrglDeleteProgramPipelines)(n, pipelines); }
// void callProgramUniform2dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {  (*ptrglProgramUniform2dv)(program, location, count, value); }
// void callProgramUniform2fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {  (*ptrglProgramUniform2fv)(program, location, count, value); }
// void callGetMultisamplefv(GLenum pname, GLuint index, GLfloat* val) {  (*ptrglGetMultisamplefv)(pname, index, val); }
// void callVertexP3uiv(GLenum type_, GLuint* value) {  (*ptrglVertexP3uiv)(type_, value); }
// void callBlendFuncSeparatei(GLuint buf, GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha) {  (*ptrglBlendFuncSeparatei)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha); }
// void callGetActiveUniformName(GLuint program, GLuint uniformIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformName) {  (*ptrglGetActiveUniformName)(program, uniformIndex, bufSize, length, uniformName); }
// void callTexCoordP1ui(GLenum type_, GLuint coords) {  (*ptrglTexCoordP1ui)(type_, coords); }
// void callInvalidateTexImage(GLuint texture, GLint level) {  (*ptrglInvalidateTexImage)(texture, level); }
// GLenum callCheckFramebufferStatus(GLenum target) { return (*ptrglCheckFramebufferStatus)(target); }
// void callMultiTexCoordP3uiv(GLenum texture, GLenum type_, GLuint* coords) {  (*ptrglMultiTexCoordP3uiv)(texture, type_, coords); }
// void callBindVertexArray(GLuint array) {  (*ptrglBindVertexArray)(array); }
// void callColorP3ui(GLenum type_, GLuint color) {  (*ptrglColorP3ui)(type_, color); }
// void callUniform2i(GLint location, GLint v0, GLint v1) {  (*ptrglUniform2i)(location, v0, v1); }
// void callGetActiveAttrib(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type_, GLchar* name) {  (*ptrglGetActiveAttrib)(program, index, bufSize, length, size, type_, name); }
// void callBeginConditionalRender(GLuint id, GLenum mode) {  (*ptrglBeginConditionalRender)(id, mode); }
// void callProgramUniform3d(GLuint program, GLint location, GLdouble v0, GLdouble v1, GLdouble v2) {  (*ptrglProgramUniform3d)(program, location, v0, v1, v2); }
// void callReadPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type_, GLvoid* pixels) {  (*ptrglReadPixels)(x, y, width, height, format, type_, pixels); }
// void callVertexAttribL1d(GLuint index, GLdouble x) {  (*ptrglVertexAttribL1d)(index, x); }
// void callVertexAttribPointer(GLuint index, GLint size, GLenum type_, GLboolean normalized, GLsizei stride, GLvoid* pointer) {  (*ptrglVertexAttribPointer)(index, size, type_, normalized, stride, pointer); }
// void callGetVertexAttribIuiv(GLuint index, GLenum pname, GLuint* params) {  (*ptrglGetVertexAttribIuiv)(index, pname, params); }
// void callProgramUniform3ui(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2) {  (*ptrglProgramUniform3ui)(program, location, v0, v1, v2); }
// void callProgramUniform2iv(GLuint program, GLint location, GLsizei count, GLint* value) {  (*ptrglProgramUniform2iv)(program, location, count, value); }
// void callSamplerParameteriv(GLuint sampler, GLenum pname, GLint* param) {  (*ptrglSamplerParameteriv)(sampler, pname, param); }
// void callVertexAttribI4sv(GLuint index, GLshort* v) {  (*ptrglVertexAttribI4sv)(index, v); }
// void callProgramUniformMatrix4x2fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix4x2fv)(program, location, count, transpose, value); }
// void callBindFramebuffer(GLenum target, GLuint framebuffer) {  (*ptrglBindFramebuffer)(target, framebuffer); }
// void callProgramUniformMatrix2x3dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix2x3dv)(program, location, count, transpose, value); }
// void callUniformMatrix2x4fv(GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglUniformMatrix2x4fv)(location, count, transpose, value); }
// void callCopyTexImage2D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {  (*ptrglCopyTexImage2D)(target, level, internalformat, x, y, width, height, border); }
// void callVertexP3ui(GLenum type_, GLuint value) {  (*ptrglVertexP3ui)(type_, value); }
// void callGetActiveSubroutineUniformiv(GLuint program, GLenum shadertype, GLuint index, GLenum pname, GLint* values) {  (*ptrglGetActiveSubroutineUniformiv)(program, shadertype, index, pname, values); }
// void callProgramUniformMatrix3x4dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix3x4dv)(program, location, count, transpose, value); }
// void callClearBufferfi(GLenum buffer, GLint drawbuffer, GLfloat depth, GLint stencil) {  (*ptrglClearBufferfi)(buffer, drawbuffer, depth, stencil); }
// void callProgramUniform1fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {  (*ptrglProgramUniform1fv)(program, location, count, value); }
// void callBindTransformFeedback(GLenum target, GLuint id) {  (*ptrglBindTransformFeedback)(target, id); }
// void callGenerateMipmap(GLenum target) {  (*ptrglGenerateMipmap)(target); }
// void callCompressedTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexSubImage3D)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data); }
// void callGetSamplerParameteriv(GLuint sampler, GLenum pname, GLint* params) {  (*ptrglGetSamplerParameteriv)(sampler, pname, params); }
// void callGenProgramPipelines(GLsizei n, GLuint* pipelines) {  (*ptrglGenProgramPipelines)(n, pipelines); }
// GLvoid* callMapBufferRange(GLenum target, GLintptr offset, GLsizeiptr length, GLbitfield access) { return (*ptrglMapBufferRange)(target, offset, length, access); }
// void callMultiDrawArrays(GLenum mode, GLint* first, GLsizei* count, GLsizei drawcount) {  (*ptrglMultiDrawArrays)(mode, first, count, drawcount); }
// void callProgramUniform4f(GLuint program, GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {  (*ptrglProgramUniform4f)(program, location, v0, v1, v2, v3); }
// void callDeleteTransformFeedbacks(GLsizei n, GLuint* ids) {  (*ptrglDeleteTransformFeedbacks)(n, ids); }
// void callTextureView(GLuint texture, GLenum target, GLuint origtexture, GLenum internalformat, GLuint minlevel, GLuint numlevels, GLuint minlayer, GLuint numlayers) {  (*ptrglTextureView)(texture, target, origtexture, internalformat, minlevel, numlevels, minlayer, numlayers); }
// void callUniform1d(GLint location, GLdouble x) {  (*ptrglUniform1d)(location, x); }
// void callVertexAttribIFormat(GLuint attribindex, GLint size, GLenum type_, GLuint relativeoffset) {  (*ptrglVertexAttribIFormat)(attribindex, size, type_, relativeoffset); }
// void callUniform4ui(GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3) {  (*ptrglUniform4ui)(location, v0, v1, v2, v3); }
// void callUniform3iv(GLint location, GLsizei count, GLint* value) {  (*ptrglUniform3iv)(location, count, value); }
// GLboolean callIsQuery(GLuint id) { return (*ptrglIsQuery)(id); }
// void callGetDoublei_v(GLenum target, GLuint index, GLdouble* data) {  (*ptrglGetDoublei_v)(target, index, data); }
// void callViewport(GLint x, GLint y, GLsizei width, GLsizei height) {  (*ptrglViewport)(x, y, width, height); }
// void callBlendFunc(GLenum sfactor, GLenum dfactor) {  (*ptrglBlendFunc)(sfactor, dfactor); }
// void callUniform4i(GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {  (*ptrglUniform4i)(location, v0, v1, v2, v3); }
// void callDeleteTextures(GLsizei n, GLuint* textures) {  (*ptrglDeleteTextures)(n, textures); }
// void callSecondaryColorP3ui(GLenum type_, GLuint color) {  (*ptrglSecondaryColorP3ui)(type_, color); }
// void callFramebufferTextureLayer(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer) {  (*ptrglFramebufferTextureLayer)(target, attachment, texture, level, layer); }
// void callGetTransformFeedbackVarying(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLsizei* size, GLenum* type_, GLchar* name) {  (*ptrglGetTransformFeedbackVarying)(program, index, bufSize, length, size, type_, name); }
// void callDebugMessageCallback(GLDEBUGPROC callback, void* userParam) {  (*ptrglDebugMessageCallback)(callback, userParam); }
// void callCompressedTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, GLvoid* data) {  (*ptrglCompressedTexSubImage2D)(target, level, xoffset, yoffset, width, height, format, imageSize, data); }
// void callGetQueryObjectiv(GLuint id, GLenum pname, GLint* params) {  (*ptrglGetQueryObjectiv)(id, pname, params); }
// void callDeleteQueries(GLsizei n, GLuint* ids) {  (*ptrglDeleteQueries)(n, ids); }
// void callDrawTransformFeedbackStreamInstanced(GLenum mode, GLuint id, GLuint stream, GLsizei instancecount) {  (*ptrglDrawTransformFeedbackStreamInstanced)(mode, id, stream, instancecount); }
// void callVertexAttribI2iv(GLuint index, GLint* v) {  (*ptrglVertexAttribI2iv)(index, v); }
// void callCopyTexImage1D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border) {  (*ptrglCopyTexImage1D)(target, level, internalformat, x, y, width, border); }
// void callUniform1iv(GLint location, GLsizei count, GLint* value) {  (*ptrglUniform1iv)(location, count, value); }
// void callStencilMaskSeparate(GLenum face, GLuint mask) {  (*ptrglStencilMaskSeparate)(face, mask); }
// GLsync callFenceSync(GLenum condition, GLbitfield flags) { return (*ptrglFenceSync)(condition, flags); }
// void callReleaseShaderCompiler() {  (*ptrglReleaseShaderCompiler)(); }
// void callDrawElementsInstancedBaseInstance(GLenum mode, GLsizei count, GLenum type_, void* indices, GLsizei instancecount, GLuint baseinstance) {  (*ptrglDrawElementsInstancedBaseInstance)(mode, count, type_, indices, instancecount, baseinstance); }
// void callProgramUniformMatrix3dv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglProgramUniformMatrix3dv)(program, location, count, transpose, value); }
// void callFlush() {  (*ptrglFlush)(); }
// void callBindBufferBase(GLenum target, GLuint index, GLuint buffer) {  (*ptrglBindBufferBase)(target, index, buffer); }
// void callEnableVertexAttribArray(GLuint index) {  (*ptrglEnableVertexAttribArray)(index); }
// GLint callGetProgramResourceLocation(GLuint program, GLenum programInterface, GLchar* name) { return (*ptrglGetProgramResourceLocation)(program, programInterface, name); }
// void callPointSize(GLfloat size) {  (*ptrglPointSize)(size); }
// void callFramebufferTexture1D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level) {  (*ptrglFramebufferTexture1D)(target, attachment, textarget, texture, level); }
// void callReadBuffer(GLenum mode) {  (*ptrglReadBuffer)(mode); }
// void callDrawArraysInstancedBaseInstance(GLenum mode, GLint first, GLsizei count, GLsizei instancecount, GLuint baseinstance) {  (*ptrglDrawArraysInstancedBaseInstance)(mode, first, count, instancecount, baseinstance); }
// void callDepthRangef(GLfloat n, GLfloat f) {  (*ptrglDepthRangef)(n, f); }
// void callHint(GLenum target, GLenum mode) {  (*ptrglHint)(target, mode); }
// void callVertexAttribP3ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {  (*ptrglVertexAttribP3ui)(index, type_, normalized, value); }
// void callGetVertexAttribPointerv(GLuint index, GLenum pname, GLvoid** pointer) {  (*ptrglGetVertexAttribPointerv)(index, pname, pointer); }
// void callGetBooleani_v(GLenum target, GLuint index, GLboolean* data) {  (*ptrglGetBooleani_v)(target, index, data); }
// void callTexParameteri(GLenum target, GLenum pname, GLint param) {  (*ptrglTexParameteri)(target, pname, param); }
// void callVertexAttribI3i(GLuint index, GLint x, GLint y, GLint z) {  (*ptrglVertexAttribI3i)(index, x, y, z); }
// void callVertexAttribL3d(GLuint index, GLdouble x, GLdouble y, GLdouble z) {  (*ptrglVertexAttribL3d)(index, x, y, z); }
// void callGetUniformfv(GLuint program, GLint location, GLfloat* params) {  (*ptrglGetUniformfv)(program, location, params); }
// void callCullFace(GLenum mode) {  (*ptrglCullFace)(mode); }
// GLboolean callIsSync(GLsync sync) { return (*ptrglIsSync)(sync); }
// void callDrawElementsBaseVertex(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLint basevertex) {  (*ptrglDrawElementsBaseVertex)(mode, count, type_, indices, basevertex); }
// void callProgramUniform1ui(GLuint program, GLint location, GLuint v0) {  (*ptrglProgramUniform1ui)(program, location, v0); }
// void callProgramUniformMatrix3x2fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix3x2fv)(program, location, count, transpose, value); }
// void callVertexAttribP4ui(GLuint index, GLenum type_, GLboolean normalized, GLuint value) {  (*ptrglVertexAttribP4ui)(index, type_, normalized, value); }
// void callGetSamplerParameterIuiv(GLuint sampler, GLenum pname, GLuint* params) {  (*ptrglGetSamplerParameterIuiv)(sampler, pname, params); }
// void callLinkProgram(GLuint program) {  (*ptrglLinkProgram)(program); }
// void callGetBufferParameteriv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetBufferParameteriv)(target, pname, params); }
// void callTexParameteriv(GLenum target, GLenum pname, GLint* params) {  (*ptrglTexParameteriv)(target, pname, params); }
// void callMultiDrawArraysIndirect(GLenum mode, void* indirect, GLsizei drawcount, GLsizei stride) {  (*ptrglMultiDrawArraysIndirect)(mode, indirect, drawcount, stride); }
// void callGetQueryObjectui64v(GLuint id, GLenum pname, GLuint64* params) {  (*ptrglGetQueryObjectui64v)(id, pname, params); }
// void callGenVertexArrays(GLsizei n, GLuint* arrays) {  (*ptrglGenVertexArrays)(n, arrays); }
// void callVertexAttribBinding(GLuint attribindex, GLuint bindingindex) {  (*ptrglVertexAttribBinding)(attribindex, bindingindex); }
// void callPointParameterf(GLenum pname, GLfloat param) {  (*ptrglPointParameterf)(pname, param); }
// void callProgramUniform1uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {  (*ptrglProgramUniform1uiv)(program, location, count, value); }
// void callUniformMatrix4x3dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix4x3dv)(location, count, transpose, value); }
// GLuint callCreateShader(GLenum type_) { return (*ptrglCreateShader)(type_); }
// void callGetTexParameteriv(GLenum target, GLenum pname, GLint* params) {  (*ptrglGetTexParameteriv)(target, pname, params); }
// void callUniformMatrix4x2dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix4x2dv)(location, count, transpose, value); }
// void callVertexAttribI1uiv(GLuint index, GLuint* v) {  (*ptrglVertexAttribI1uiv)(index, v); }
// void callPolygonMode(GLenum face, GLenum mode) {  (*ptrglPolygonMode)(face, mode); }
// void callGetProgramResourceName(GLuint program, GLenum programInterface, GLuint index, GLsizei bufSize, GLsizei* length, GLchar* name) {  (*ptrglGetProgramResourceName)(program, programInterface, index, bufSize, length, name); }
// void callDrawElementsInstancedBaseVertex(GLenum mode, GLsizei count, GLenum type_, GLvoid* indices, GLsizei instancecount, GLint basevertex) {  (*ptrglDrawElementsInstancedBaseVertex)(mode, count, type_, indices, instancecount, basevertex); }
// void callTexImage3DMultisample(GLenum target, GLsizei samples, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLboolean fixedsamplelocations) {  (*ptrglTexImage3DMultisample)(target, samples, internalformat, width, height, depth, fixedsamplelocations); }
// GLboolean callIsVertexArray(GLuint array) { return (*ptrglIsVertexArray)(array); }
// void callUniform3fv(GLint location, GLsizei count, GLfloat* value) {  (*ptrglUniform3fv)(location, count, value); }
// void callTexCoordP2ui(GLenum type_, GLuint coords) {  (*ptrglTexCoordP2ui)(type_, coords); }
// void callProgramUniform1f(GLuint program, GLint location, GLfloat v0) {  (*ptrglProgramUniform1f)(program, location, v0); }
// void callProgramUniformMatrix4fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix4fv)(program, location, count, transpose, value); }
// void callGenQueries(GLsizei n, GLuint* ids) {  (*ptrglGenQueries)(n, ids); }
// void callGetQueryObjectuiv(GLuint id, GLenum pname, GLuint* params) {  (*ptrglGetQueryObjectuiv)(id, pname, params); }
// void callGetObjectLabel(GLenum identifier, GLuint name, GLsizei bufSize, GLsizei* length, GLchar* label) {  (*ptrglGetObjectLabel)(identifier, name, bufSize, length, label); }
// void callProgramUniform4ui(GLuint program, GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3) {  (*ptrglProgramUniform4ui)(program, location, v0, v1, v2, v3); }
// GLboolean callIsProgram(GLuint program) { return (*ptrglIsProgram)(program); }
// void callObjectPtrLabel(void* ptr, GLsizei length, GLchar* label) {  (*ptrglObjectPtrLabel)(ptr, length, label); }
// void callDeleteShader(GLuint shader) {  (*ptrglDeleteShader)(shader); }
// void callMultiDrawElementsBaseVertex(GLenum mode, GLsizei* count, GLenum type_, GLvoid* const* indices, GLsizei drawcount, GLint* basevertex) {  (*ptrglMultiDrawElementsBaseVertex)(mode, count, type_, indices, drawcount, basevertex); }
// void callInvalidateFramebuffer(GLenum target, GLsizei numAttachments, GLenum* attachments) {  (*ptrglInvalidateFramebuffer)(target, numAttachments, attachments); }
// void callGetActiveUniformsiv(GLuint program, GLsizei uniformCount, GLuint* uniformIndices, GLenum pname, GLint* params) {  (*ptrglGetActiveUniformsiv)(program, uniformCount, uniformIndices, pname, params); }
// void callUniform3ui(GLint location, GLuint v0, GLuint v1, GLuint v2) {  (*ptrglUniform3ui)(location, v0, v1, v2); }
// void callUniformMatrix3x2dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix3x2dv)(location, count, transpose, value); }
// void callGetShaderInfoLog(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {  (*ptrglGetShaderInfoLog)(shader, bufSize, length, infoLog); }
// void callBlendColor(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {  (*ptrglBlendColor)(red, green, blue, alpha); }
// void callBindFragDataLocationIndexed(GLuint program, GLuint colorNumber, GLuint index, GLchar* name) {  (*ptrglBindFragDataLocationIndexed)(program, colorNumber, index, name); }
// void callGetFramebufferAttachmentParameteriv(GLenum target, GLenum attachment, GLenum pname, GLint* params) {  (*ptrglGetFramebufferAttachmentParameteriv)(target, attachment, pname, params); }
// void callVertexAttribP3uiv(GLuint index, GLenum type_, GLboolean normalized, GLuint* value) {  (*ptrglVertexAttribP3uiv)(index, type_, normalized, value); }
// void callDrawElementsInstancedBaseVertexBaseInstance(GLenum mode, GLsizei count, GLenum type_, void* indices, GLsizei instancecount, GLint basevertex, GLuint baseinstance) {  (*ptrglDrawElementsInstancedBaseVertexBaseInstance)(mode, count, type_, indices, instancecount, basevertex, baseinstance); }
// void callUniformSubroutinesuiv(GLenum shadertype, GLsizei count, GLuint* indices) {  (*ptrglUniformSubroutinesuiv)(shadertype, count, indices); }
// void callUniform3uiv(GLint location, GLsizei count, GLuint* value) {  (*ptrglUniform3uiv)(location, count, value); }
// void callVertexAttribI4bv(GLuint index, GLbyte* v) {  (*ptrglVertexAttribI4bv)(index, v); }
// void callProgramUniform3i(GLuint program, GLint location, GLint v0, GLint v1, GLint v2) {  (*ptrglProgramUniform3i)(program, location, v0, v1, v2); }
// void callProgramUniform3dv(GLuint program, GLint location, GLsizei count, GLdouble* value) {  (*ptrglProgramUniform3dv)(program, location, count, value); }
// const GLubyte * callGetStringi(GLenum name, GLuint index) { return (*ptrglGetStringi)(name, index); }
// void callGetQueryObjecti64v(GLuint id, GLenum pname, GLint64* params) {  (*ptrglGetQueryObjecti64v)(id, pname, params); }
// void callVertexAttribL1dv(GLuint index, GLdouble* v) {  (*ptrglVertexAttribL1dv)(index, v); }
// void callGetInteger64i_v(GLenum target, GLuint index, GLint64* data) {  (*ptrglGetInteger64i_v)(target, index, data); }
// void callGetPointerv(GLenum pname, GLvoid** params) {  (*ptrglGetPointerv)(pname, params); }
// void callActiveShaderProgram(GLuint pipeline, GLuint program) {  (*ptrglActiveShaderProgram)(pipeline, program); }
// void callUniformMatrix2x4dv(GLint location, GLsizei count, GLboolean transpose, GLdouble* value) {  (*ptrglUniformMatrix2x4dv)(location, count, transpose, value); }
// void callSampleCoverage(GLfloat value, GLboolean invert) {  (*ptrglSampleCoverage)(value, invert); }
// void callVertexP2uiv(GLenum type_, GLuint* value) {  (*ptrglVertexP2uiv)(type_, value); }
// void callDrawTransformFeedback(GLenum mode, GLuint id) {  (*ptrglDrawTransformFeedback)(mode, id); }
// void callGetIntegerv(GLenum pname, GLint* params) {  (*ptrglGetIntegerv)(pname, params); }
// void callDeleteProgram(GLuint program) {  (*ptrglDeleteProgram)(program); }
// void callProgramUniform4uiv(GLuint program, GLint location, GLsizei count, GLuint* value) {  (*ptrglProgramUniform4uiv)(program, location, count, value); }
// void callUniform1uiv(GLint location, GLsizei count, GLuint* value) {  (*ptrglUniform1uiv)(location, count, value); }
// void callGetDoublev(GLenum pname, GLdouble* params) {  (*ptrglGetDoublev)(pname, params); }
// void callProgramUniform3fv(GLuint program, GLint location, GLsizei count, GLfloat* value) {  (*ptrglProgramUniform3fv)(program, location, count, value); }
// void callGetActiveUniformBlockName(GLuint program, GLuint uniformBlockIndex, GLsizei bufSize, GLsizei* length, GLchar* uniformBlockName) {  (*ptrglGetActiveUniformBlockName)(program, uniformBlockIndex, bufSize, length, uniformBlockName); }
// void callGetSamplerParameterIiv(GLuint sampler, GLenum pname, GLint* params) {  (*ptrglGetSamplerParameterIiv)(sampler, pname, params); }
// void callGetInternalformativ(GLenum target, GLenum internalformat, GLenum pname, GLsizei bufSize, GLint* params) {  (*ptrglGetInternalformativ)(target, internalformat, pname, bufSize, params); }
// void callGetBufferPointerv(GLenum target, GLenum pname, GLvoid** params) {  (*ptrglGetBufferPointerv)(target, pname, params); }
// void callClear(GLbitfield mask) {  (*ptrglClear)(mask); }
// void callGetVertexAttribLdv(GLuint index, GLenum pname, GLdouble* params) {  (*ptrglGetVertexAttribLdv)(index, pname, params); }
// void callProgramUniformMatrix2x4fv(GLuint program, GLint location, GLsizei count, GLboolean transpose, GLfloat* value) {  (*ptrglProgramUniformMatrix2x4fv)(program, location, count, transpose, value); }
// void callPrimitiveRestartIndex(GLuint index) {  (*ptrglPrimitiveRestartIndex)(index); }
// void callLineWidth(GLfloat width) {  (*ptrglLineWidth)(width); }
// void callGetTexLevelParameterfv(GLenum target, GLint level, GLenum pname, GLfloat* params) {  (*ptrglGetTexLevelParameterfv)(target, level, pname, params); }
// void callGetIntegeri_v(GLenum target, GLuint index, GLint* data) {  (*ptrglGetIntegeri_v)(target, index, data); }
// int coreGetProcAddresses() {
// 	ptrglUniform1uiv = coreGetProcAddress("glUniform1uiv");
// 	ptrglGetDoublev = coreGetProcAddress("glGetDoublev");
// 	ptrglProgramUniform3fv = coreGetProcAddress("glProgramUniform3fv");
// 	ptrglGetActiveUniformBlockName = coreGetProcAddress("glGetActiveUniformBlockName");
// 	ptrglGetSamplerParameterIiv = coreGetProcAddress("glGetSamplerParameterIiv");
// 	ptrglGetInternalformativ = coreGetProcAddress("glGetInternalformativ");
// 	ptrglGetBufferPointerv = coreGetProcAddress("glGetBufferPointerv");
// 	ptrglClear = coreGetProcAddress("glClear");
// 	ptrglGetVertexAttribLdv = coreGetProcAddress("glGetVertexAttribLdv");
// 	ptrglProgramUniformMatrix2x4fv = coreGetProcAddress("glProgramUniformMatrix2x4fv");
// 	ptrglPrimitiveRestartIndex = coreGetProcAddress("glPrimitiveRestartIndex");
// 	ptrglLineWidth = coreGetProcAddress("glLineWidth");
// 	ptrglGetTexLevelParameterfv = coreGetProcAddress("glGetTexLevelParameterfv");
// 	ptrglGetIntegeri_v = coreGetProcAddress("glGetIntegeri_v");
// 	ptrglGetProgramBinary = coreGetProcAddress("glGetProgramBinary");
// 	ptrglIsEnabledi = coreGetProcAddress("glIsEnabledi");
// 	ptrglGetUniformIndices = coreGetProcAddress("glGetUniformIndices");
// 	ptrglClearBufferData = coreGetProcAddress("glClearBufferData");
// 	ptrglCopyTexSubImage3D = coreGetProcAddress("glCopyTexSubImage3D");
// 	ptrglProgramUniform3uiv = coreGetProcAddress("glProgramUniform3uiv");
// 	ptrglTransformFeedbackVaryings = coreGetProcAddress("glTransformFeedbackVaryings");
// 	ptrglPointParameterfv = coreGetProcAddress("glPointParameterfv");
// 	ptrglVertexAttribI1iv = coreGetProcAddress("glVertexAttribI1iv");
// 	ptrglTexStorage2D = coreGetProcAddress("glTexStorage2D");
// 	ptrglVertexAttribI4usv = coreGetProcAddress("glVertexAttribI4usv");
// 	ptrglRenderbufferStorageMultisample = coreGetProcAddress("glRenderbufferStorageMultisample");
// 	ptrglBufferData = coreGetProcAddress("glBufferData");
// 	ptrglUniformBlockBinding = coreGetProcAddress("glUniformBlockBinding");
// 	ptrglGetUniformuiv = coreGetProcAddress("glGetUniformuiv");
// 	ptrglGetInternalformati64v = coreGetProcAddress("glGetInternalformati64v");
// 	ptrglTexParameterIuiv = coreGetProcAddress("glTexParameterIuiv");
// 	ptrglUniformMatrix3x4dv = coreGetProcAddress("glUniformMatrix3x4dv");
// 	ptrglGetFloati_v = coreGetProcAddress("glGetFloati_v");
// 	ptrglCreateProgram = coreGetProcAddress("glCreateProgram");
// 	ptrglIsProgramPipeline = coreGetProcAddress("glIsProgramPipeline");
// 	ptrglProgramUniform4dv = coreGetProcAddress("glProgramUniform4dv");
// 	ptrglGetInteger64v = coreGetProcAddress("glGetInteger64v");
// 	ptrglGetProgramInterfaceiv = coreGetProcAddress("glGetProgramInterfaceiv");
// 	ptrglNormalP3uiv = coreGetProcAddress("glNormalP3uiv");
// 	ptrglShaderBinary = coreGetProcAddress("glShaderBinary");
// 	ptrglPatchParameterfv = coreGetProcAddress("glPatchParameterfv");
// 	ptrglGetSamplerParameterfv = coreGetProcAddress("glGetSamplerParameterfv");
// 	ptrglGetProgramPipelineInfoLog = coreGetProcAddress("glGetProgramPipelineInfoLog");
// 	ptrglObjectLabel = coreGetProcAddress("glObjectLabel");
// 	ptrglVertexAttribI4uiv = coreGetProcAddress("glVertexAttribI4uiv");
// 	ptrglCopyImageSubData = coreGetProcAddress("glCopyImageSubData");
// 	ptrglBlendEquationi = coreGetProcAddress("glBlendEquationi");
// 	ptrglProgramUniformMatrix3x2dv = coreGetProcAddress("glProgramUniformMatrix3x2dv");
// 	ptrglUniform2iv = coreGetProcAddress("glUniform2iv");
// 	ptrglEnable = coreGetProcAddress("glEnable");
// 	ptrglUniform2dv = coreGetProcAddress("glUniform2dv");
// 	ptrglDrawArraysInstanced = coreGetProcAddress("glDrawArraysInstanced");
// 	ptrglDepthFunc = coreGetProcAddress("glDepthFunc");
// 	ptrglColorP4uiv = coreGetProcAddress("glColorP4uiv");
// 	ptrglFramebufferParameteri = coreGetProcAddress("glFramebufferParameteri");
// 	ptrglVertexAttribI3ui = coreGetProcAddress("glVertexAttribI3ui");
// 	ptrglQueryCounter = coreGetProcAddress("glQueryCounter");
// 	ptrglProgramUniformMatrix3x4fv = coreGetProcAddress("glProgramUniformMatrix3x4fv");
// 	ptrglBlendFuncSeparate = coreGetProcAddress("glBlendFuncSeparate");
// 	ptrglCopyBufferSubData = coreGetProcAddress("glCopyBufferSubData");
// 	ptrglUniform3d = coreGetProcAddress("glUniform3d");
// 	ptrglDrawRangeElements = coreGetProcAddress("glDrawRangeElements");
// 	ptrglVertexAttribIPointer = coreGetProcAddress("glVertexAttribIPointer");
// 	ptrglSamplerParameterIuiv = coreGetProcAddress("glSamplerParameterIuiv");
// 	ptrglGetActiveUniform = coreGetProcAddress("glGetActiveUniform");
// 	ptrglCopyTexSubImage2D = coreGetProcAddress("glCopyTexSubImage2D");
// 	ptrglGetUniformdv = coreGetProcAddress("glGetUniformdv");
// 	ptrglVertexAttribI1ui = coreGetProcAddress("glVertexAttribI1ui");
// 	ptrglProgramUniform4iv = coreGetProcAddress("glProgramUniform4iv");
// 	ptrglInvalidateTexSubImage = coreGetProcAddress("glInvalidateTexSubImage");
// 	ptrglProgramUniformMatrix4x3fv = coreGetProcAddress("glProgramUniformMatrix4x3fv");
// 	ptrglColorP4ui = coreGetProcAddress("glColorP4ui");
// 	ptrglUniformMatrix2fv = coreGetProcAddress("glUniformMatrix2fv");
// 	ptrglClearBufferfv = coreGetProcAddress("glClearBufferfv");
// 	ptrglMultiDrawElements = coreGetProcAddress("glMultiDrawElements");
// 	ptrglTexImage3D = coreGetProcAddress("glTexImage3D");
// 	ptrglClearBufferSubData = coreGetProcAddress("glClearBufferSubData");
// 	ptrglGetActiveAtomicCounterBufferiv = coreGetProcAddress("glGetActiveAtomicCounterBufferiv");
// 	ptrglViewportIndexedfv = coreGetProcAddress("glViewportIndexedfv");
// 	ptrglFramebufferTexture2D = coreGetProcAddress("glFramebufferTexture2D");
// 	ptrglBindBuffer = coreGetProcAddress("glBindBuffer");
// 	ptrglVertexAttribI3uiv = coreGetProcAddress("glVertexAttribI3uiv");
// 	ptrglVertexAttribI4i = coreGetProcAddress("glVertexAttribI4i");
// 	ptrglCompressedTexSubImage1D = coreGetProcAddress("glCompressedTexSubImage1D");
// 	ptrglGetBufferSubData = coreGetProcAddress("glGetBufferSubData");
// 	ptrglDeleteFramebuffers = coreGetProcAddress("glDeleteFramebuffers");
// 	ptrglVertexAttribLFormat = coreGetProcAddress("glVertexAttribLFormat");
// 	ptrglUniformMatrix2x3fv = coreGetProcAddress("glUniformMatrix2x3fv");
// 	ptrglGetProgramResourceIndex = coreGetProcAddress("glGetProgramResourceIndex");
// 	ptrglDisableVertexAttribArray = coreGetProcAddress("glDisableVertexAttribArray");
// 	ptrglBindBufferRange = coreGetProcAddress("glBindBufferRange");
// 	ptrglVertexAttribI2uiv = coreGetProcAddress("glVertexAttribI2uiv");
// 	ptrglGetShaderPrecisionFormat = coreGetProcAddress("glGetShaderPrecisionFormat");
// 	ptrglVertexAttribP2ui = coreGetProcAddress("glVertexAttribP2ui");
// 	ptrglProvokingVertex = coreGetProcAddress("glProvokingVertex");
// 	ptrglTexBuffer = coreGetProcAddress("glTexBuffer");
// 	ptrglGetString = coreGetProcAddress("glGetString");
// 	ptrglCompileShader = coreGetProcAddress("glCompileShader");
// 	ptrglUniform1f = coreGetProcAddress("glUniform1f");
// 	ptrglUniform3f = coreGetProcAddress("glUniform3f");
// 	ptrglGetObjectPtrLabel = coreGetProcAddress("glGetObjectPtrLabel");
// 	ptrglProgramUniform1i = coreGetProcAddress("glProgramUniform1i");
// 	ptrglVertexAttribI2ui = coreGetProcAddress("glVertexAttribI2ui");
// 	ptrglProgramUniformMatrix4x2dv = coreGetProcAddress("glProgramUniformMatrix4x2dv");
// 	ptrglCompressedTexImage3D = coreGetProcAddress("glCompressedTexImage3D");
// 	ptrglClearColor = coreGetProcAddress("glClearColor");
// 	ptrglTexCoordP3ui = coreGetProcAddress("glTexCoordP3ui");
// 	ptrglSamplerParameterIiv = coreGetProcAddress("glSamplerParameterIiv");
// 	ptrglDispatchComputeIndirect = coreGetProcAddress("glDispatchComputeIndirect");
// 	ptrglDeleteSync = coreGetProcAddress("glDeleteSync");
// 	ptrglVertexAttribI2i = coreGetProcAddress("glVertexAttribI2i");
// 	ptrglStencilFunc = coreGetProcAddress("glStencilFunc");
// 	ptrglGetDebugMessageLog = coreGetProcAddress("glGetDebugMessageLog");
// 	ptrglGetUniformSubroutineuiv = coreGetProcAddress("glGetUniformSubroutineuiv");
// 	ptrglGetFragDataIndex = coreGetProcAddress("glGetFragDataIndex");
// 	ptrglStencilFuncSeparate = coreGetProcAddress("glStencilFuncSeparate");
// 	ptrglTexSubImage2D = coreGetProcAddress("glTexSubImage2D");
// 	ptrglPointParameteri = coreGetProcAddress("glPointParameteri");
// 	ptrglIsShader = coreGetProcAddress("glIsShader");
// 	ptrglUniformMatrix4x2fv = coreGetProcAddress("glUniformMatrix4x2fv");
// 	ptrglGetQueryIndexediv = coreGetProcAddress("glGetQueryIndexediv");
// 	ptrglGetUniformLocation = coreGetProcAddress("glGetUniformLocation");
// 	ptrglActiveTexture = coreGetProcAddress("glActiveTexture");
// 	ptrglMultiTexCoordP4ui = coreGetProcAddress("glMultiTexCoordP4ui");
// 	ptrglGetProgramStageiv = coreGetProcAddress("glGetProgramStageiv");
// 	ptrglProgramUniform1dv = coreGetProcAddress("glProgramUniform1dv");
// 	ptrglShaderStorageBlockBinding = coreGetProcAddress("glShaderStorageBlockBinding");
// 	ptrglTexStorage2DMultisample = coreGetProcAddress("glTexStorage2DMultisample");
// 	ptrglBlendEquationSeparate = coreGetProcAddress("glBlendEquationSeparate");
// 	ptrglNormalP3ui = coreGetProcAddress("glNormalP3ui");
// 	ptrglUniformMatrix3x4fv = coreGetProcAddress("glUniformMatrix3x4fv");
// 	ptrglVertexAttribI1i = coreGetProcAddress("glVertexAttribI1i");
// 	ptrglUniform2ui = coreGetProcAddress("glUniform2ui");
// 	ptrglGetTexParameterfv = coreGetProcAddress("glGetTexParameterfv");
// 	ptrglGenRenderbuffers = coreGetProcAddress("glGenRenderbuffers");
// 	ptrglIsSampler = coreGetProcAddress("glIsSampler");
// 	ptrglEndQueryIndexed = coreGetProcAddress("glEndQueryIndexed");
// 	ptrglPixelStoref = coreGetProcAddress("glPixelStoref");
// 	ptrglVertexAttribDivisor = coreGetProcAddress("glVertexAttribDivisor");
// 	ptrglStencilOp = coreGetProcAddress("glStencilOp");
// 	ptrglInvalidateBufferSubData = coreGetProcAddress("glInvalidateBufferSubData");
// 	ptrglVertexAttribL3dv = coreGetProcAddress("glVertexAttribL3dv");
// 	ptrglClearBufferiv = coreGetProcAddress("glClearBufferiv");
// 	ptrglLogicOp = coreGetProcAddress("glLogicOp");
// 	ptrglCopyTexSubImage1D = coreGetProcAddress("glCopyTexSubImage1D");
// 	ptrglGetVertexAttribdv = coreGetProcAddress("glGetVertexAttribdv");
// 	ptrglDrawBuffer = coreGetProcAddress("glDrawBuffer");
// 	ptrglGenBuffers = coreGetProcAddress("glGenBuffers");
// 	ptrglTexStorage3D = coreGetProcAddress("glTexStorage3D");
// 	ptrglTexStorage3DMultisample = coreGetProcAddress("glTexStorage3DMultisample");
// 	ptrglDrawElementsInstanced = coreGetProcAddress("glDrawElementsInstanced");
// 	ptrglPolygonOffset = coreGetProcAddress("glPolygonOffset");
// 	ptrglMapBuffer = coreGetProcAddress("glMapBuffer");
// 	ptrglGetFragDataLocation = coreGetProcAddress("glGetFragDataLocation");
// 	ptrglIsTexture = coreGetProcAddress("glIsTexture");
// 	ptrglDeleteVertexArrays = coreGetProcAddress("glDeleteVertexArrays");
// 	ptrglViewportArrayv = coreGetProcAddress("glViewportArrayv");
// 	ptrglDrawTransformFeedbackInstanced = coreGetProcAddress("glDrawTransformFeedbackInstanced");
// 	ptrglDebugMessageInsert = coreGetProcAddress("glDebugMessageInsert");
// 	ptrglGetVertexAttribiv = coreGetProcAddress("glGetVertexAttribiv");
// 	ptrglGetFloatv = coreGetProcAddress("glGetFloatv");
// 	ptrglClientWaitSync = coreGetProcAddress("glClientWaitSync");
// 	ptrglUniform2f = coreGetProcAddress("glUniform2f");
// 	ptrglIsBuffer = coreGetProcAddress("glIsBuffer");
// 	ptrglUniform1fv = coreGetProcAddress("glUniform1fv");
// 	ptrglMultiTexCoordP1uiv = coreGetProcAddress("glMultiTexCoordP1uiv");
// 	ptrglVertexAttribI4ui = coreGetProcAddress("glVertexAttribI4ui");
// 	ptrglGetBooleanv = coreGetProcAddress("glGetBooleanv");
// 	ptrglGetShaderiv = coreGetProcAddress("glGetShaderiv");
// 	ptrglProgramUniform2uiv = coreGetProcAddress("glProgramUniform2uiv");
// 	ptrglTexCoordP1uiv = coreGetProcAddress("glTexCoordP1uiv");
// 	ptrglEndConditionalRender = coreGetProcAddress("glEndConditionalRender");
// 	ptrglBindImageTexture = coreGetProcAddress("glBindImageTexture");
// 	ptrglProgramUniform4i = coreGetProcAddress("glProgramUniform4i");
// 	ptrglGetBufferParameteri64v = coreGetProcAddress("glGetBufferParameteri64v");
// 	ptrglRenderbufferStorage = coreGetProcAddress("glRenderbufferStorage");
// 	ptrglDepthRangeArrayv = coreGetProcAddress("glDepthRangeArrayv");
// 	ptrglFramebufferTexture = coreGetProcAddress("glFramebufferTexture");
// 	ptrglGetProgramResourceiv = coreGetProcAddress("glGetProgramResourceiv");
// 	ptrglMultiTexCoordP2uiv = coreGetProcAddress("glMultiTexCoordP2uiv");
// 	ptrglBindSampler = coreGetProcAddress("glBindSampler");
// 	ptrglUniformMatrix4x3fv = coreGetProcAddress("glUniformMatrix4x3fv");
// 	ptrglGetVertexAttribIiv = coreGetProcAddress("glGetVertexAttribIiv");
// 	ptrglMemoryBarrier = coreGetProcAddress("glMemoryBarrier");
// 	ptrglBeginQueryIndexed = coreGetProcAddress("glBeginQueryIndexed");
// 	ptrglUniform2d = coreGetProcAddress("glUniform2d");
// 	ptrglDeleteRenderbuffers = coreGetProcAddress("glDeleteRenderbuffers");
// 	ptrglUseProgramStages = coreGetProcAddress("glUseProgramStages");
// 	ptrglDrawBuffers = coreGetProcAddress("glDrawBuffers");
// 	ptrglGetSynciv = coreGetProcAddress("glGetSynciv");
// 	ptrglEndQuery = coreGetProcAddress("glEndQuery");
// 	ptrglDepthRangeIndexed = coreGetProcAddress("glDepthRangeIndexed");
// 	ptrglGetError = coreGetProcAddress("glGetError");
// 	ptrglProgramUniform4d = coreGetProcAddress("glProgramUniform4d");
// 	ptrglPushDebugGroup = coreGetProcAddress("glPushDebugGroup");
// 	ptrglGetQueryiv = coreGetProcAddress("glGetQueryiv");
// 	ptrglIsEnabled = coreGetProcAddress("glIsEnabled");
// 	ptrglSecondaryColorP3uiv = coreGetProcAddress("glSecondaryColorP3uiv");
// 	ptrglTexSubImage3D = coreGetProcAddress("glTexSubImage3D");
// 	ptrglDisablei = coreGetProcAddress("glDisablei");
// 	ptrglDrawRangeElementsBaseVertex = coreGetProcAddress("glDrawRangeElementsBaseVertex");
// 	ptrglBindTexture = coreGetProcAddress("glBindTexture");
// 	ptrglUniform4iv = coreGetProcAddress("glUniform4iv");
// 	ptrglMinSampleShading = coreGetProcAddress("glMinSampleShading");
// 	ptrglUniform1dv = coreGetProcAddress("glUniform1dv");
// 	ptrglVertexAttribP1ui = coreGetProcAddress("glVertexAttribP1ui");
// 	ptrglPixelStorei = coreGetProcAddress("glPixelStorei");
// 	ptrglFrontFace = coreGetProcAddress("glFrontFace");
// 	ptrglProgramUniform2d = coreGetProcAddress("glProgramUniform2d");
// 	ptrglTexStorage1D = coreGetProcAddress("glTexStorage1D");
// 	ptrglUniform3i = coreGetProcAddress("glUniform3i");
// 	ptrglTexImage2DMultisample = coreGetProcAddress("glTexImage2DMultisample");
// 	ptrglUniformMatrix3x2fv = coreGetProcAddress("glUniformMatrix3x2fv");
// 	ptrglTexImage2D = coreGetProcAddress("glTexImage2D");
// 	ptrglGetSubroutineUniformLocation = coreGetProcAddress("glGetSubroutineUniformLocation");
// 	ptrglUniformMatrix3dv = coreGetProcAddress("glUniformMatrix3dv");
// 	ptrglBindRenderbuffer = coreGetProcAddress("glBindRenderbuffer");
// 	ptrglBindProgramPipeline = coreGetProcAddress("glBindProgramPipeline");
// 	ptrglUniform4dv = coreGetProcAddress("glUniform4dv");
// 	ptrglClearBufferuiv = coreGetProcAddress("glClearBufferuiv");
// 	ptrglBeginTransformFeedback = coreGetProcAddress("glBeginTransformFeedback");
// 	ptrglUniformMatrix3fv = coreGetProcAddress("glUniformMatrix3fv");
// 	ptrglDispatchCompute = coreGetProcAddress("glDispatchCompute");
// 	ptrglBindVertexBuffer = coreGetProcAddress("glBindVertexBuffer");
// 	ptrglVertexBindingDivisor = coreGetProcAddress("glVertexBindingDivisor");
// 	ptrglMultiTexCoordP4uiv = coreGetProcAddress("glMultiTexCoordP4uiv");
// 	ptrglProgramUniformMatrix2x3fv = coreGetProcAddress("glProgramUniformMatrix2x3fv");
// 	ptrglGetTexLevelParameteriv = coreGetProcAddress("glGetTexLevelParameteriv");
// 	ptrglVertexP4uiv = coreGetProcAddress("glVertexP4uiv");
// 	ptrglGetTexParameterIuiv = coreGetProcAddress("glGetTexParameterIuiv");
// 	ptrglDebugMessageControl = coreGetProcAddress("glDebugMessageControl");
// 	ptrglSamplerParameteri = coreGetProcAddress("glSamplerParameteri");
// 	ptrglProgramUniform2f = coreGetProcAddress("glProgramUniform2f");
// 	ptrglGetVertexAttribfv = coreGetProcAddress("glGetVertexAttribfv");
// 	ptrglVertexAttribP2uiv = coreGetProcAddress("glVertexAttribP2uiv");
// 	ptrglScissorIndexed = coreGetProcAddress("glScissorIndexed");
// 	ptrglProgramUniform1d = coreGetProcAddress("glProgramUniform1d");
// 	ptrglVertexP4ui = coreGetProcAddress("glVertexP4ui");
// 	ptrglDepthMask = coreGetProcAddress("glDepthMask");
// 	ptrglUniformMatrix2dv = coreGetProcAddress("glUniformMatrix2dv");
// 	ptrglTexCoordP4uiv = coreGetProcAddress("glTexCoordP4uiv");
// 	ptrglUniform4f = coreGetProcAddress("glUniform4f");
// 	ptrglUniform4fv = coreGetProcAddress("glUniform4fv");
// 	ptrglClampColor = coreGetProcAddress("glClampColor");
// 	ptrglGenFramebuffers = coreGetProcAddress("glGenFramebuffers");
// 	ptrglTexCoordP3uiv = coreGetProcAddress("glTexCoordP3uiv");
// 	ptrglColorMask = coreGetProcAddress("glColorMask");
// 	ptrglSamplerParameterfv = coreGetProcAddress("glSamplerParameterfv");
// 	ptrglDrawElementsIndirect = coreGetProcAddress("glDrawElementsIndirect");
// 	ptrglGetUniformiv = coreGetProcAddress("glGetUniformiv");
// 	ptrglGetRenderbufferParameteriv = coreGetProcAddress("glGetRenderbufferParameteriv");
// 	ptrglProgramParameteri = coreGetProcAddress("glProgramParameteri");
// 	ptrglUniform2fv = coreGetProcAddress("glUniform2fv");
// 	ptrglProgramUniformMatrix2dv = coreGetProcAddress("glProgramUniformMatrix2dv");
// 	ptrglProgramUniform4fv = coreGetProcAddress("glProgramUniform4fv");
// 	ptrglTexParameterf = coreGetProcAddress("glTexParameterf");
// 	ptrglWaitSync = coreGetProcAddress("glWaitSync");
// 	ptrglStencilOpSeparate = coreGetProcAddress("glStencilOpSeparate");
// 	ptrglVertexAttribP1uiv = coreGetProcAddress("glVertexAttribP1uiv");
// 	ptrglVertexP2ui = coreGetProcAddress("glVertexP2ui");
// 	ptrglColorMaski = coreGetProcAddress("glColorMaski");
// 	ptrglFramebufferTexture3D = coreGetProcAddress("glFramebufferTexture3D");
// 	ptrglUniform2uiv = coreGetProcAddress("glUniform2uiv");
// 	ptrglGetActiveUniformBlockiv = coreGetProcAddress("glGetActiveUniformBlockiv");
// 	ptrglVertexAttribL4dv = coreGetProcAddress("glVertexAttribL4dv");
// 	ptrglStencilMask = coreGetProcAddress("glStencilMask");
// 	ptrglMultiDrawElementsIndirect = coreGetProcAddress("glMultiDrawElementsIndirect");
// 	ptrglGetActiveSubroutineName = coreGetProcAddress("glGetActiveSubroutineName");
// 	ptrglGetAttachedShaders = coreGetProcAddress("glGetAttachedShaders");
// 	ptrglClearDepthf = coreGetProcAddress("glClearDepthf");
// 	ptrglClearStencil = coreGetProcAddress("glClearStencil");
// 	ptrglGetCompressedTexImage = coreGetProcAddress("glGetCompressedTexImage");
// 	ptrglPointParameteriv = coreGetProcAddress("glPointParameteriv");
// 	ptrglProgramUniform2i = coreGetProcAddress("glProgramUniform2i");
// 	ptrglGetTexImage = coreGetProcAddress("glGetTexImage");
// 	ptrglBindFragDataLocation = coreGetProcAddress("glBindFragDataLocation");
// 	ptrglUniformMatrix2x3dv = coreGetProcAddress("glUniformMatrix2x3dv");
// 	ptrglBlendFunci = coreGetProcAddress("glBlendFunci");
// 	ptrglBindAttribLocation = coreGetProcAddress("glBindAttribLocation");
// 	ptrglCompressedTexImage2D = coreGetProcAddress("glCompressedTexImage2D");
// 	ptrglVertexAttribLPointer = coreGetProcAddress("glVertexAttribLPointer");
// 	ptrglTexCoordP4ui = coreGetProcAddress("glTexCoordP4ui");
// 	ptrglPauseTransformFeedback = coreGetProcAddress("glPauseTransformFeedback");
// 	ptrglProgramUniformMatrix2x4dv = coreGetProcAddress("glProgramUniformMatrix2x4dv");
// 	ptrglCompressedTexImage1D = coreGetProcAddress("glCompressedTexImage1D");
// 	ptrglEndTransformFeedback = coreGetProcAddress("glEndTransformFeedback");
// 	ptrglUniform3dv = coreGetProcAddress("glUniform3dv");
// 	ptrglScissor = coreGetProcAddress("glScissor");
// 	ptrglTexSubImage1D = coreGetProcAddress("glTexSubImage1D");
// 	ptrglSampleMaski = coreGetProcAddress("glSampleMaski");
// 	ptrglFinish = coreGetProcAddress("glFinish");
// 	ptrglProgramUniform1iv = coreGetProcAddress("glProgramUniform1iv");
// 	ptrglScissorIndexedv = coreGetProcAddress("glScissorIndexedv");
// 	ptrglGetTexParameterIiv = coreGetProcAddress("glGetTexParameterIiv");
// 	ptrglUseProgram = coreGetProcAddress("glUseProgram");
// 	ptrglInvalidateBufferData = coreGetProcAddress("glInvalidateBufferData");
// 	ptrglScissorArrayv = coreGetProcAddress("glScissorArrayv");
// 	ptrglGetProgramiv = coreGetProcAddress("glGetProgramiv");
// 	ptrglDisable = coreGetProcAddress("glDisable");
// 	ptrglGetUniformBlockIndex = coreGetProcAddress("glGetUniformBlockIndex");
// 	ptrglBlendEquation = coreGetProcAddress("glBlendEquation");
// 	ptrglProgramUniformMatrix3fv = coreGetProcAddress("glProgramUniformMatrix3fv");
// 	ptrglDrawArrays = coreGetProcAddress("glDrawArrays");
// 	ptrglCreateShaderProgramv = coreGetProcAddress("glCreateShaderProgramv");
// 	ptrglGenTransformFeedbacks = coreGetProcAddress("glGenTransformFeedbacks");
// 	ptrglUniform1i = coreGetProcAddress("glUniform1i");
// 	ptrglVertexAttribL4d = coreGetProcAddress("glVertexAttribL4d");
// 	ptrglVertexAttribI4iv = coreGetProcAddress("glVertexAttribI4iv");
// 	ptrglProgramUniformMatrix2fv = coreGetProcAddress("glProgramUniformMatrix2fv");
// 	ptrglBlitFramebuffer = coreGetProcAddress("glBlitFramebuffer");
// 	ptrglInvalidateSubFramebuffer = coreGetProcAddress("glInvalidateSubFramebuffer");
// 	ptrglMultiTexCoordP3ui = coreGetProcAddress("glMultiTexCoordP3ui");
// 	ptrglShaderSource = coreGetProcAddress("glShaderSource");
// 	ptrglUniformMatrix4dv = coreGetProcAddress("glUniformMatrix4dv");
// 	ptrglProgramUniformMatrix4dv = coreGetProcAddress("glProgramUniformMatrix4dv");
// 	ptrglDeleteSamplers = coreGetProcAddress("glDeleteSamplers");
// 	ptrglGenSamplers = coreGetProcAddress("glGenSamplers");
// 	ptrglTexParameterIiv = coreGetProcAddress("glTexParameterIiv");
// 	ptrglAttachShader = coreGetProcAddress("glAttachShader");
// 	ptrglValidateProgramPipeline = coreGetProcAddress("glValidateProgramPipeline");
// 	ptrglFramebufferRenderbuffer = coreGetProcAddress("glFramebufferRenderbuffer");
// 	ptrglDetachShader = coreGetProcAddress("glDetachShader");
// 	ptrglUniform4d = coreGetProcAddress("glUniform4d");
// 	ptrglDepthRange = coreGetProcAddress("glDepthRange");
// 	ptrglTexBufferRange = coreGetProcAddress("glTexBufferRange");
// 	ptrglUniform4uiv = coreGetProcAddress("glUniform4uiv");
// 	ptrglGetSubroutineIndex = coreGetProcAddress("glGetSubroutineIndex");
// 	ptrglVertexAttribI3iv = coreGetProcAddress("glVertexAttribI3iv");
// 	ptrglPatchParameteri = coreGetProcAddress("glPatchParameteri");
// 	ptrglGetProgramInfoLog = coreGetProcAddress("glGetProgramInfoLog");
// 	ptrglFlushMappedBufferRange = coreGetProcAddress("glFlushMappedBufferRange");
// 	ptrglBlendEquationSeparatei = coreGetProcAddress("glBlendEquationSeparatei");
// 	ptrglUniform1ui = coreGetProcAddress("glUniform1ui");
// 	ptrglGetActiveSubroutineUniformName = coreGetProcAddress("glGetActiveSubroutineUniformName");
// 	ptrglVertexAttribL2d = coreGetProcAddress("glVertexAttribL2d");
// 	ptrglIsRenderbuffer = coreGetProcAddress("glIsRenderbuffer");
// 	ptrglProgramUniform3f = coreGetProcAddress("glProgramUniform3f");
// 	ptrglViewportIndexedf = coreGetProcAddress("glViewportIndexedf");
// 	ptrglResumeTransformFeedback = coreGetProcAddress("glResumeTransformFeedback");
// 	ptrglMultiTexCoordP1ui = coreGetProcAddress("glMultiTexCoordP1ui");
// 	ptrglValidateProgram = coreGetProcAddress("glValidateProgram");
// 	ptrglPopDebugGroup = coreGetProcAddress("glPopDebugGroup");
// 	ptrglVertexAttribP4uiv = coreGetProcAddress("glVertexAttribP4uiv");
// 	ptrglDrawArraysIndirect = coreGetProcAddress("glDrawArraysIndirect");
// 	ptrglProgramUniformMatrix4x3dv = coreGetProcAddress("glProgramUniformMatrix4x3dv");
// 	ptrglDrawElements = coreGetProcAddress("glDrawElements");
// 	ptrglClearDepth = coreGetProcAddress("glClearDepth");
// 	ptrglSamplerParameterf = coreGetProcAddress("glSamplerParameterf");
// 	ptrglBufferSubData = coreGetProcAddress("glBufferSubData");
// 	ptrglBeginQuery = coreGetProcAddress("glBeginQuery");
// 	ptrglGetProgramPipelineiv = coreGetProcAddress("glGetProgramPipelineiv");
// 	ptrglProgramBinary = coreGetProcAddress("glProgramBinary");
// 	ptrglVertexAttribI4ubv = coreGetProcAddress("glVertexAttribI4ubv");
// 	ptrglVertexAttribFormat = coreGetProcAddress("glVertexAttribFormat");
// 	ptrglIsTransformFeedback = coreGetProcAddress("glIsTransformFeedback");
// 	ptrglGetAttribLocation = coreGetProcAddress("glGetAttribLocation");
// 	ptrglProgramUniform3iv = coreGetProcAddress("glProgramUniform3iv");
// 	ptrglUnmapBuffer = coreGetProcAddress("glUnmapBuffer");
// 	ptrglGenTextures = coreGetProcAddress("glGenTextures");
// 	ptrglGetProgramResourceLocationIndex = coreGetProcAddress("glGetProgramResourceLocationIndex");
// 	ptrglGetShaderSource = coreGetProcAddress("glGetShaderSource");
// 	ptrglIsFramebuffer = coreGetProcAddress("glIsFramebuffer");
// 	ptrglTexCoordP2uiv = coreGetProcAddress("glTexCoordP2uiv");
// 	ptrglUniformMatrix4fv = coreGetProcAddress("glUniformMatrix4fv");
// 	ptrglTexImage1D = coreGetProcAddress("glTexImage1D");
// 	ptrglEnablei = coreGetProcAddress("glEnablei");
// 	ptrglProgramUniform2ui = coreGetProcAddress("glProgramUniform2ui");
// 	ptrglColorP3uiv = coreGetProcAddress("glColorP3uiv");
// 	ptrglGetFramebufferParameteriv = coreGetProcAddress("glGetFramebufferParameteriv");
// 	ptrglDeleteBuffers = coreGetProcAddress("glDeleteBuffers");
// 	ptrglMultiTexCoordP2ui = coreGetProcAddress("glMultiTexCoordP2ui");
// 	ptrglTexParameterfv = coreGetProcAddress("glTexParameterfv");
// 	ptrglVertexAttribL2dv = coreGetProcAddress("glVertexAttribL2dv");
// 	ptrglDrawTransformFeedbackStream = coreGetProcAddress("glDrawTransformFeedbackStream");
// 	ptrglDeleteProgramPipelines = coreGetProcAddress("glDeleteProgramPipelines");
// 	ptrglProgramUniform2dv = coreGetProcAddress("glProgramUniform2dv");
// 	ptrglProgramUniform2fv = coreGetProcAddress("glProgramUniform2fv");
// 	ptrglGetMultisamplefv = coreGetProcAddress("glGetMultisamplefv");
// 	ptrglVertexP3uiv = coreGetProcAddress("glVertexP3uiv");
// 	ptrglBlendFuncSeparatei = coreGetProcAddress("glBlendFuncSeparatei");
// 	ptrglGetActiveUniformName = coreGetProcAddress("glGetActiveUniformName");
// 	ptrglTexCoordP1ui = coreGetProcAddress("glTexCoordP1ui");
// 	ptrglInvalidateTexImage = coreGetProcAddress("glInvalidateTexImage");
// 	ptrglCheckFramebufferStatus = coreGetProcAddress("glCheckFramebufferStatus");
// 	ptrglMultiTexCoordP3uiv = coreGetProcAddress("glMultiTexCoordP3uiv");
// 	ptrglBindVertexArray = coreGetProcAddress("glBindVertexArray");
// 	ptrglColorP3ui = coreGetProcAddress("glColorP3ui");
// 	ptrglUniform2i = coreGetProcAddress("glUniform2i");
// 	ptrglGetActiveAttrib = coreGetProcAddress("glGetActiveAttrib");
// 	ptrglBeginConditionalRender = coreGetProcAddress("glBeginConditionalRender");
// 	ptrglProgramUniform3d = coreGetProcAddress("glProgramUniform3d");
// 	ptrglReadPixels = coreGetProcAddress("glReadPixels");
// 	ptrglVertexAttribL1d = coreGetProcAddress("glVertexAttribL1d");
// 	ptrglVertexAttribPointer = coreGetProcAddress("glVertexAttribPointer");
// 	ptrglGetVertexAttribIuiv = coreGetProcAddress("glGetVertexAttribIuiv");
// 	ptrglProgramUniform3ui = coreGetProcAddress("glProgramUniform3ui");
// 	ptrglProgramUniform2iv = coreGetProcAddress("glProgramUniform2iv");
// 	ptrglSamplerParameteriv = coreGetProcAddress("glSamplerParameteriv");
// 	ptrglVertexAttribI4sv = coreGetProcAddress("glVertexAttribI4sv");
// 	ptrglProgramUniformMatrix4x2fv = coreGetProcAddress("glProgramUniformMatrix4x2fv");
// 	ptrglBindFramebuffer = coreGetProcAddress("glBindFramebuffer");
// 	ptrglProgramUniformMatrix2x3dv = coreGetProcAddress("glProgramUniformMatrix2x3dv");
// 	ptrglUniformMatrix2x4fv = coreGetProcAddress("glUniformMatrix2x4fv");
// 	ptrglCopyTexImage2D = coreGetProcAddress("glCopyTexImage2D");
// 	ptrglVertexP3ui = coreGetProcAddress("glVertexP3ui");
// 	ptrglGetActiveSubroutineUniformiv = coreGetProcAddress("glGetActiveSubroutineUniformiv");
// 	ptrglProgramUniformMatrix3x4dv = coreGetProcAddress("glProgramUniformMatrix3x4dv");
// 	ptrglClearBufferfi = coreGetProcAddress("glClearBufferfi");
// 	ptrglProgramUniform1fv = coreGetProcAddress("glProgramUniform1fv");
// 	ptrglBindTransformFeedback = coreGetProcAddress("glBindTransformFeedback");
// 	ptrglGenerateMipmap = coreGetProcAddress("glGenerateMipmap");
// 	ptrglCompressedTexSubImage3D = coreGetProcAddress("glCompressedTexSubImage3D");
// 	ptrglGetSamplerParameteriv = coreGetProcAddress("glGetSamplerParameteriv");
// 	ptrglGenProgramPipelines = coreGetProcAddress("glGenProgramPipelines");
// 	ptrglMapBufferRange = coreGetProcAddress("glMapBufferRange");
// 	ptrglMultiDrawArrays = coreGetProcAddress("glMultiDrawArrays");
// 	ptrglProgramUniform4f = coreGetProcAddress("glProgramUniform4f");
// 	ptrglDeleteTransformFeedbacks = coreGetProcAddress("glDeleteTransformFeedbacks");
// 	ptrglTextureView = coreGetProcAddress("glTextureView");
// 	ptrglUniform1d = coreGetProcAddress("glUniform1d");
// 	ptrglVertexAttribIFormat = coreGetProcAddress("glVertexAttribIFormat");
// 	ptrglUniform4ui = coreGetProcAddress("glUniform4ui");
// 	ptrglUniform3iv = coreGetProcAddress("glUniform3iv");
// 	ptrglIsQuery = coreGetProcAddress("glIsQuery");
// 	ptrglGetDoublei_v = coreGetProcAddress("glGetDoublei_v");
// 	ptrglViewport = coreGetProcAddress("glViewport");
// 	ptrglBlendFunc = coreGetProcAddress("glBlendFunc");
// 	ptrglUniform4i = coreGetProcAddress("glUniform4i");
// 	ptrglDeleteTextures = coreGetProcAddress("glDeleteTextures");
// 	ptrglSecondaryColorP3ui = coreGetProcAddress("glSecondaryColorP3ui");
// 	ptrglFramebufferTextureLayer = coreGetProcAddress("glFramebufferTextureLayer");
// 	ptrglGetTransformFeedbackVarying = coreGetProcAddress("glGetTransformFeedbackVarying");
// 	ptrglDebugMessageCallback = coreGetProcAddress("glDebugMessageCallback");
// 	ptrglCompressedTexSubImage2D = coreGetProcAddress("glCompressedTexSubImage2D");
// 	ptrglGetQueryObjectiv = coreGetProcAddress("glGetQueryObjectiv");
// 	ptrglDeleteQueries = coreGetProcAddress("glDeleteQueries");
// 	ptrglDrawTransformFeedbackStreamInstanced = coreGetProcAddress("glDrawTransformFeedbackStreamInstanced");
// 	ptrglVertexAttribI2iv = coreGetProcAddress("glVertexAttribI2iv");
// 	ptrglCopyTexImage1D = coreGetProcAddress("glCopyTexImage1D");
// 	ptrglUniform1iv = coreGetProcAddress("glUniform1iv");
// 	ptrglStencilMaskSeparate = coreGetProcAddress("glStencilMaskSeparate");
// 	ptrglFenceSync = coreGetProcAddress("glFenceSync");
// 	ptrglReleaseShaderCompiler = coreGetProcAddress("glReleaseShaderCompiler");
// 	ptrglDrawElementsInstancedBaseInstance = coreGetProcAddress("glDrawElementsInstancedBaseInstance");
// 	ptrglProgramUniformMatrix3dv = coreGetProcAddress("glProgramUniformMatrix3dv");
// 	ptrglFlush = coreGetProcAddress("glFlush");
// 	ptrglBindBufferBase = coreGetProcAddress("glBindBufferBase");
// 	ptrglEnableVertexAttribArray = coreGetProcAddress("glEnableVertexAttribArray");
// 	ptrglGetProgramResourceLocation = coreGetProcAddress("glGetProgramResourceLocation");
// 	ptrglPointSize = coreGetProcAddress("glPointSize");
// 	ptrglFramebufferTexture1D = coreGetProcAddress("glFramebufferTexture1D");
// 	ptrglReadBuffer = coreGetProcAddress("glReadBuffer");
// 	ptrglDrawArraysInstancedBaseInstance = coreGetProcAddress("glDrawArraysInstancedBaseInstance");
// 	ptrglDepthRangef = coreGetProcAddress("glDepthRangef");
// 	ptrglHint = coreGetProcAddress("glHint");
// 	ptrglVertexAttribP3ui = coreGetProcAddress("glVertexAttribP3ui");
// 	ptrglGetVertexAttribPointerv = coreGetProcAddress("glGetVertexAttribPointerv");
// 	ptrglGetBooleani_v = coreGetProcAddress("glGetBooleani_v");
// 	ptrglTexParameteri = coreGetProcAddress("glTexParameteri");
// 	ptrglVertexAttribI3i = coreGetProcAddress("glVertexAttribI3i");
// 	ptrglVertexAttribL3d = coreGetProcAddress("glVertexAttribL3d");
// 	ptrglGetUniformfv = coreGetProcAddress("glGetUniformfv");
// 	ptrglCullFace = coreGetProcAddress("glCullFace");
// 	ptrglIsSync = coreGetProcAddress("glIsSync");
// 	ptrglDrawElementsBaseVertex = coreGetProcAddress("glDrawElementsBaseVertex");
// 	ptrglProgramUniform1ui = coreGetProcAddress("glProgramUniform1ui");
// 	ptrglProgramUniformMatrix3x2fv = coreGetProcAddress("glProgramUniformMatrix3x2fv");
// 	ptrglVertexAttribP4ui = coreGetProcAddress("glVertexAttribP4ui");
// 	ptrglGetSamplerParameterIuiv = coreGetProcAddress("glGetSamplerParameterIuiv");
// 	ptrglLinkProgram = coreGetProcAddress("glLinkProgram");
// 	ptrglGetBufferParameteriv = coreGetProcAddress("glGetBufferParameteriv");
// 	ptrglTexParameteriv = coreGetProcAddress("glTexParameteriv");
// 	ptrglMultiDrawArraysIndirect = coreGetProcAddress("glMultiDrawArraysIndirect");
// 	ptrglGetQueryObjectui64v = coreGetProcAddress("glGetQueryObjectui64v");
// 	ptrglGenVertexArrays = coreGetProcAddress("glGenVertexArrays");
// 	ptrglVertexAttribBinding = coreGetProcAddress("glVertexAttribBinding");
// 	ptrglPointParameterf = coreGetProcAddress("glPointParameterf");
// 	ptrglProgramUniform1uiv = coreGetProcAddress("glProgramUniform1uiv");
// 	ptrglUniformMatrix4x3dv = coreGetProcAddress("glUniformMatrix4x3dv");
// 	ptrglCreateShader = coreGetProcAddress("glCreateShader");
// 	ptrglGetTexParameteriv = coreGetProcAddress("glGetTexParameteriv");
// 	ptrglUniformMatrix4x2dv = coreGetProcAddress("glUniformMatrix4x2dv");
// 	ptrglVertexAttribI1uiv = coreGetProcAddress("glVertexAttribI1uiv");
// 	ptrglPolygonMode = coreGetProcAddress("glPolygonMode");
// 	ptrglGetProgramResourceName = coreGetProcAddress("glGetProgramResourceName");
// 	ptrglDrawElementsInstancedBaseVertex = coreGetProcAddress("glDrawElementsInstancedBaseVertex");
// 	ptrglTexImage3DMultisample = coreGetProcAddress("glTexImage3DMultisample");
// 	ptrglIsVertexArray = coreGetProcAddress("glIsVertexArray");
// 	ptrglUniform3fv = coreGetProcAddress("glUniform3fv");
// 	ptrglTexCoordP2ui = coreGetProcAddress("glTexCoordP2ui");
// 	ptrglProgramUniform1f = coreGetProcAddress("glProgramUniform1f");
// 	ptrglProgramUniformMatrix4fv = coreGetProcAddress("glProgramUniformMatrix4fv");
// 	ptrglGenQueries = coreGetProcAddress("glGenQueries");
// 	ptrglGetQueryObjectuiv = coreGetProcAddress("glGetQueryObjectuiv");
// 	ptrglGetObjectLabel = coreGetProcAddress("glGetObjectLabel");
// 	ptrglProgramUniform4ui = coreGetProcAddress("glProgramUniform4ui");
// 	ptrglIsProgram = coreGetProcAddress("glIsProgram");
// 	ptrglObjectPtrLabel = coreGetProcAddress("glObjectPtrLabel");
// 	ptrglDeleteShader = coreGetProcAddress("glDeleteShader");
// 	ptrglMultiDrawElementsBaseVertex = coreGetProcAddress("glMultiDrawElementsBaseVertex");
// 	ptrglInvalidateFramebuffer = coreGetProcAddress("glInvalidateFramebuffer");
// 	ptrglGetActiveUniformsiv = coreGetProcAddress("glGetActiveUniformsiv");
// 	ptrglUniform3ui = coreGetProcAddress("glUniform3ui");
// 	ptrglUniformMatrix3x2dv = coreGetProcAddress("glUniformMatrix3x2dv");
// 	ptrglGetShaderInfoLog = coreGetProcAddress("glGetShaderInfoLog");
// 	ptrglBlendColor = coreGetProcAddress("glBlendColor");
// 	ptrglBindFragDataLocationIndexed = coreGetProcAddress("glBindFragDataLocationIndexed");
// 	ptrglGetFramebufferAttachmentParameteriv = coreGetProcAddress("glGetFramebufferAttachmentParameteriv");
// 	ptrglVertexAttribP3uiv = coreGetProcAddress("glVertexAttribP3uiv");
// 	ptrglDrawElementsInstancedBaseVertexBaseInstance = coreGetProcAddress("glDrawElementsInstancedBaseVertexBaseInstance");
// 	ptrglUniformSubroutinesuiv = coreGetProcAddress("glUniformSubroutinesuiv");
// 	ptrglUniform3uiv = coreGetProcAddress("glUniform3uiv");
// 	ptrglVertexAttribI4bv = coreGetProcAddress("glVertexAttribI4bv");
// 	ptrglProgramUniform3i = coreGetProcAddress("glProgramUniform3i");
// 	ptrglProgramUniform3dv = coreGetProcAddress("glProgramUniform3dv");
// 	ptrglGetStringi = coreGetProcAddress("glGetStringi");
// 	ptrglGetQueryObjecti64v = coreGetProcAddress("glGetQueryObjecti64v");
// 	ptrglVertexAttribL1dv = coreGetProcAddress("glVertexAttribL1dv");
// 	ptrglGetInteger64i_v = coreGetProcAddress("glGetInteger64i_v");
// 	ptrglGetPointerv = coreGetProcAddress("glGetPointerv");
// 	ptrglActiveShaderProgram = coreGetProcAddress("glActiveShaderProgram");
// 	ptrglUniformMatrix2x4dv = coreGetProcAddress("glUniformMatrix2x4dv");
// 	ptrglSampleCoverage = coreGetProcAddress("glSampleCoverage");
// 	ptrglVertexP2uiv = coreGetProcAddress("glVertexP2uiv");
// 	ptrglDrawTransformFeedback = coreGetProcAddress("glDrawTransformFeedback");
// 	ptrglGetIntegerv = coreGetProcAddress("glGetIntegerv");
// 	ptrglDeleteProgram = coreGetProcAddress("glDeleteProgram");
// 	ptrglProgramUniform4uiv = coreGetProcAddress("glProgramUniform4uiv");
// /* a cheap way to determine if (at least) v3.3 core profile (or higher) seems to be supported: */
// 	return ((ptrglBindSampler == NULL) ? 0 : 1);
// }
import "C"
import "unsafe"

type (
	Sizei    C.GLsizei
	Uint     C.GLuint
	Int64    C.GLint64
	Sync     C.GLsync
	Ushort   C.GLushort
	Ubyte    C.GLubyte
	Uint64   C.GLuint64
	Int      C.GLint
	Double   C.GLdouble
	Short    C.GLshort
	Enum     C.GLenum
	Float    C.GLfloat
	Half     C.GLhalf
	Ptr      unsafe.Pointer
	Byte     C.GLbyte
	Char     C.GLchar
	Clampf   C.GLclampf
	Intptr   C.GLintptr
	Clampd   C.GLclampd
	Bitfield C.GLbitfield
	Boolean  C.GLboolean
	Sizeiptr C.GLsizeiptr
)

//	Provides utility methods for working with this OpenGL binding package.
type GlUtil struct{}

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

//	Initializes this OpenGL binding after having created your GL context (with SDL, GLFW or in some other way).
func (_ GlUtil) Init() bool {
	return (C.coreGetProcAddresses() == 1)
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointSize
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPointSize.xml
func PointSize(size Float) {
	C.callPointSize((C.GLfloat)(size))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture1D.xml
func FramebufferTexture1D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int) {
	C.callFramebufferTexture1D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glReadBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glReadBuffer.xml
func ReadBuffer(mode Enum) {
	C.callReadBuffer((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArraysInstancedBaseInstance
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawArraysInstancedBaseInstance.xml
func DrawArraysInstancedBaseInstance(mode Enum, first Int, count Sizei, instancecount Sizei, baseinstance Uint) {
	C.callDrawArraysInstancedBaseInstance((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRangef
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthRangef.xml
func DepthRangef(n Float, f Float) {
	C.callDepthRangef((C.GLfloat)(n), (C.GLfloat)(f))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glHint
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glHint.xml
func Hint(target Enum, mode Enum) {
	C.callHint((C.GLenum)(target), (C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP3ui.xml
func VertexAttribP3ui(index Uint, type_ Enum, normalized Boolean, value Uint) {
	C.callVertexAttribP3ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribPointerv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribPointerv.xml
func GetVertexAttribPointerv(index Uint, pname Enum, pointer *Ptr) {
	C.callGetVertexAttribPointerv((C.GLuint)(index), (C.GLenum)(pname), (*unsafe.Pointer)(pointer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBooleani_v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBooleani_v.xml
func GetBooleani_v(target Enum, index Uint, data *Boolean) {
	C.callGetBooleani_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLboolean)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameteri.xml
func TexParameteri(target Enum, pname Enum, param Int) {
	C.callTexParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3i.xml
func VertexAttribI3i(index Uint, x Int, y Int, z Int) {
	C.callVertexAttribI3i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL3d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL3d.xml
func VertexAttribL3d(index Uint, x Double, y Double, z Double) {
	C.callVertexAttribL3d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformfv.xml
func GetUniformfv(program Uint, location Int, params *Float) {
	C.callGetUniformfv((C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCullFace
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCullFace.xml
func CullFace(mode Enum) {
	C.callCullFace((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsSync
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsSync.xml
func IsSync(sync Sync) Boolean {
	return (Boolean)(C.callIsSync((C.GLsync)(sync)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsBaseVertex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsBaseVertex.xml
func DrawElementsBaseVertex(mode Enum, count Sizei, type_ Enum, indices Ptr, basevertex Int) {
	C.callDrawElementsBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLint)(basevertex))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1ui.xml
func ProgramUniform1ui(program Uint, location Int, v0 Uint) {
	C.callProgramUniform1ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x2fv.xml
func ProgramUniformMatrix3x2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix3x2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP4ui.xml
func VertexAttribP4ui(index Uint, type_ Enum, normalized Boolean, value Uint) {
	C.callVertexAttribP4ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameterIuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameterIuiv.xml
func GetSamplerParameterIuiv(sampler Uint, pname Enum, params *Uint) {
	C.callGetSamplerParameterIuiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glLinkProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glLinkProgram.xml
func LinkProgram(program Uint) {
	C.callLinkProgram((C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferParameteriv.xml
func GetBufferParameteriv(target Enum, pname Enum, params *Int) {
	C.callGetBufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameteriv.xml
func TexParameteriv(target Enum, pname Enum, params *Int) {
	C.callTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawArraysIndirect
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawArraysIndirect.xml
func MultiDrawArraysIndirect(mode Enum, indirect Ptr, drawcount Sizei, stride Sizei) {
	C.callMultiDrawArraysIndirect((C.GLenum)(mode), (unsafe.Pointer)(indirect), (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjectui64v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectui64v.xml
func GetQueryObjectui64v(id Uint, pname Enum, params *Uint64) {
	C.callGetQueryObjectui64v((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint64)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenVertexArrays
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenVertexArrays.xml
func GenVertexArrays(n Sizei, arrays *Uint) {
	C.callGenVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribBinding
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribBinding.xml
func VertexAttribBinding(attribindex Uint, bindingindex Uint) {
	C.callVertexAttribBinding((C.GLuint)(attribindex), (C.GLuint)(bindingindex))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameterf
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPointParameterf.xml
func PointParameterf(pname Enum, param Float) {
	C.callPointParameterf((C.GLenum)(pname), (C.GLfloat)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1uiv.xml
func ProgramUniform1uiv(program Uint, location Int, count Sizei, value *Uint) {
	C.callProgramUniform1uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x3dv.xml
func UniformMatrix4x3dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix4x3dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCreateShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCreateShader.xml
func CreateShader(type_ Enum) Uint {
	return (Uint)(C.callCreateShader((C.GLenum)(type_)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameteriv.xml
func GetTexParameteriv(target Enum, pname Enum, params *Int) {
	C.callGetTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x2dv.xml
func UniformMatrix4x2dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix4x2dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1uiv.xml
func VertexAttribI1uiv(index Uint, v *Uint) {
	C.callVertexAttribI1uiv((C.GLuint)(index), (*C.GLuint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPolygonMode
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPolygonMode.xml
func PolygonMode(face Enum, mode Enum) {
	C.callPolygonMode((C.GLenum)(face), (C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceName
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceName.xml
func GetProgramResourceName(program Uint, programInterface Enum, index Uint, bufSize Sizei, length *Sizei, name *Char) {
	C.callGetProgramResourceName((C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstancedBaseVertex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstancedBaseVertex.xml
func DrawElementsInstancedBaseVertex(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, basevertex Int) {
	C.callDrawElementsInstancedBaseVertex((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount), (C.GLint)(basevertex))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage3DMultisample
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexImage3DMultisample.xml
func TexImage3DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean) {
	C.callTexImage3DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(fixedsamplelocations))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsVertexArray
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsVertexArray.xml
func IsVertexArray(array Uint) Boolean {
	return (Boolean)(C.callIsVertexArray((C.GLuint)(array)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3fv.xml
func Uniform3fv(location Int, count Sizei, value *Float) {
	C.callUniform3fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP2ui.xml
func TexCoordP2ui(type_ Enum, coords Uint) {
	C.callTexCoordP2ui((C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1f.xml
func ProgramUniform1f(program Uint, location Int, v0 Float) {
	C.callProgramUniform1f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4fv.xml
func ProgramUniformMatrix4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenQueries
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenQueries.xml
func GenQueries(n Sizei, ids *Uint) {
	C.callGenQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjectuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectuiv.xml
func GetQueryObjectuiv(id Uint, pname Enum, params *Uint) {
	C.callGetQueryObjectuiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetObjectLabel
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetObjectLabel.xml
func GetObjectLabel(identifier Enum, name Uint, bufSize Sizei, length *Sizei, label *Char) {
	C.callGetObjectLabel((C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(label))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4ui.xml
func ProgramUniform4ui(program Uint, location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint) {
	C.callProgramUniform4ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsProgram.xml
func IsProgram(program Uint) Boolean {
	return (Boolean)(C.callIsProgram((C.GLuint)(program)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glObjectPtrLabel
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glObjectPtrLabel.xml
func ObjectPtrLabel(ptr Ptr, length Sizei, label *Char) {
	C.callObjectPtrLabel((unsafe.Pointer)(ptr), (C.GLsizei)(length), (*C.GLchar)(label))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteShader.xml
func DeleteShader(shader Uint) {
	C.callDeleteShader((C.GLuint)(shader))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawElementsBaseVertex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawElementsBaseVertex.xml
func MultiDrawElementsBaseVertex(mode Enum, count *Sizei, type_ Enum, indices *Ptr, drawcount Sizei, basevertex *Int) {
	C.callMultiDrawElementsBaseVertex((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(drawcount), (*C.GLint)(basevertex))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateFramebuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateFramebuffer.xml
func InvalidateFramebuffer(target Enum, numAttachments Sizei, attachments *Enum) {
	C.callInvalidateFramebuffer((C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(attachments))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformsiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformsiv.xml
func GetActiveUniformsiv(program Uint, uniformCount Sizei, uniformIndices *Uint, pname Enum, params *Int) {
	C.callGetActiveUniformsiv((C.GLuint)(program), (C.GLsizei)(uniformCount), (*C.GLuint)(uniformIndices), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3ui.xml
func Uniform3ui(location Int, v0 Uint, v1 Uint, v2 Uint) {
	C.callUniform3ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x2dv.xml
func UniformMatrix3x2dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix3x2dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderInfoLog
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderInfoLog.xml
func GetShaderInfoLog(shader Uint, bufSize Sizei, length *Sizei, infoLog *Char) {
	C.callGetShaderInfoLog((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendColor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendColor.xml
func BlendColor(red Float, green Float, blue Float, alpha Float) {
	C.callBlendColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindFragDataLocationIndexed
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindFragDataLocationIndexed.xml
func BindFragDataLocationIndexed(program Uint, colorNumber Uint, index Uint, name *Char) {
	C.callBindFragDataLocationIndexed((C.GLuint)(program), (C.GLuint)(colorNumber), (C.GLuint)(index), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFramebufferAttachmentParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFramebufferAttachmentParameteriv.xml
func GetFramebufferAttachmentParameteriv(target Enum, attachment Enum, pname Enum, params *Int) {
	C.callGetFramebufferAttachmentParameteriv((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP3uiv.xml
func VertexAttribP3uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) {
	C.callVertexAttribP3uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstancedBaseVertexBaseInstance
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstancedBaseVertexBaseInstance.xml
func DrawElementsInstancedBaseVertexBaseInstance(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, basevertex Int, baseinstance Uint) {
	C.callDrawElementsInstancedBaseVertexBaseInstance((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount), (C.GLint)(basevertex), (C.GLuint)(baseinstance))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformSubroutinesuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformSubroutinesuiv.xml
func UniformSubroutinesuiv(shadertype Enum, count Sizei, indices *Uint) {
	C.callUniformSubroutinesuiv((C.GLenum)(shadertype), (C.GLsizei)(count), (*C.GLuint)(indices))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3uiv.xml
func Uniform3uiv(location Int, count Sizei, value *Uint) {
	C.callUniform3uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4bv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4bv.xml
func VertexAttribI4bv(index Uint, v *Byte) {
	C.callVertexAttribI4bv((C.GLuint)(index), (*C.GLbyte)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3i.xml
func ProgramUniform3i(program Uint, location Int, v0 Int, v1 Int, v2 Int) {
	C.callProgramUniform3i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3dv.xml
func ProgramUniform3dv(program Uint, location Int, count Sizei, value *Double) {
	C.callProgramUniform3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetStringi
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetStringi.xml
func GetStringi(name Enum, index Uint) *Ubyte {
	return (*Ubyte)(C.callGetStringi((C.GLenum)(name), (C.GLuint)(index)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjecti64v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjecti64v.xml
func GetQueryObjecti64v(id Uint, pname Enum, params *Int64) {
	C.callGetQueryObjecti64v((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint64)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL1dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL1dv.xml
func VertexAttribL1dv(index Uint, v *Double) {
	C.callVertexAttribL1dv((C.GLuint)(index), (*C.GLdouble)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInteger64i_v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetInteger64i_v.xml
func GetInteger64i_v(target Enum, index Uint, data *Int64) {
	C.callGetInteger64i_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLint64)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetPointerv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetPointerv.xml
func GetPointerv(pname Enum, params *Ptr) {
	C.callGetPointerv((C.GLenum)(pname), (*unsafe.Pointer)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glActiveShaderProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glActiveShaderProgram.xml
func ActiveShaderProgram(pipeline Uint, program Uint) {
	C.callActiveShaderProgram((C.GLuint)(pipeline), (C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x4dv.xml
func UniformMatrix2x4dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix2x4dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSampleCoverage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSampleCoverage.xml
func SampleCoverage(value Float, invert Boolean) {
	C.callSampleCoverage((C.GLfloat)(value), (C.GLboolean)(invert))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP2uiv.xml
func VertexP2uiv(type_ Enum, value *Uint) {
	C.callVertexP2uiv((C.GLenum)(type_), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedback.xml
func DrawTransformFeedback(mode Enum, id Uint) {
	C.callDrawTransformFeedback((C.GLenum)(mode), (C.GLuint)(id))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetIntegerv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetIntegerv.xml
func GetIntegerv(pname Enum, params *Int) {
	C.callGetIntegerv((C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteProgram.xml
func DeleteProgram(program Uint) {
	C.callDeleteProgram((C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4uiv.xml
func ProgramUniform4uiv(program Uint, location Int, count Sizei, value *Uint) {
	C.callProgramUniform4uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1uiv.xml
func Uniform1uiv(location Int, count Sizei, value *Uint) {
	C.callUniform1uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetDoublev
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetDoublev.xml
func GetDoublev(pname Enum, params *Double) {
	C.callGetDoublev((C.GLenum)(pname), (*C.GLdouble)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3fv.xml
func ProgramUniform3fv(program Uint, location Int, count Sizei, value *Float) {
	C.callProgramUniform3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformBlockName
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformBlockName.xml
func GetActiveUniformBlockName(program Uint, uniformBlockIndex Uint, bufSize Sizei, length *Sizei, uniformBlockName *Char) {
	C.callGetActiveUniformBlockName((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(uniformBlockName))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameterIiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameterIiv.xml
func GetSamplerParameterIiv(sampler Uint, pname Enum, params *Int) {
	C.callGetSamplerParameterIiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInternalformativ
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetInternalformativ.xml
func GetInternalformativ(target Enum, internalformat Enum, pname Enum, bufSize Sizei, params *Int) {
	C.callGetInternalformativ((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferPointerv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferPointerv.xml
func GetBufferPointerv(target Enum, pname Enum, params *Ptr) {
	C.callGetBufferPointerv((C.GLenum)(target), (C.GLenum)(pname), (*unsafe.Pointer)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClear
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClear.xml
func Clear(mask Bitfield) {
	C.callClear((C.GLbitfield)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribLdv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribLdv.xml
func GetVertexAttribLdv(index Uint, pname Enum, params *Double) {
	C.callGetVertexAttribLdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x4fv.xml
func ProgramUniformMatrix2x4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix2x4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPrimitiveRestartIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPrimitiveRestartIndex.xml
func PrimitiveRestartIndex(index Uint) {
	C.callPrimitiveRestartIndex((C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glLineWidth
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glLineWidth.xml
func LineWidth(width Float) {
	C.callLineWidth((C.GLfloat)(width))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexLevelParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexLevelParameterfv.xml
func GetTexLevelParameterfv(target Enum, level Int, pname Enum, params *Float) {
	C.callGetTexLevelParameterfv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetIntegeri_v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetIntegeri_v.xml
func GetIntegeri_v(target Enum, index Uint, data *Int) {
	C.callGetIntegeri_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramBinary
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramBinary.xml
func GetProgramBinary(program Uint, bufSize Sizei, length *Sizei, binaryFormat *Enum, binary Ptr) {
	C.callGetProgramBinary((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLenum)(binaryFormat), (unsafe.Pointer)(binary))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsEnabledi
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsEnabledi.xml
func IsEnabledi(target Enum, index Uint) Boolean {
	return (Boolean)(C.callIsEnabledi((C.GLenum)(target), (C.GLuint)(index)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformIndices
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformIndices.xml
func GetUniformIndices(program Uint, uniformCount Sizei, uniformNames **Char, uniformIndices *Uint) {
	C.callGetUniformIndices((C.GLuint)(program), (C.GLsizei)(uniformCount), (**C.GLchar)(unsafe.Pointer(uniformNames)), (*C.GLuint)(uniformIndices))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferData.xml
func ClearBufferData(target Enum, internalformat Enum, format Enum, type_ Enum, data Ptr) {
	C.callClearBufferData((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexSubImage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage3D.xml
func CopyTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, x Int, y Int, width Sizei, height Sizei) {
	C.callCopyTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3uiv.xml
func ProgramUniform3uiv(program Uint, location Int, count Sizei, value *Uint) {
	C.callProgramUniform3uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTransformFeedbackVaryings
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTransformFeedbackVaryings.xml
func TransformFeedbackVaryings(program Uint, count Sizei, varyings **Char, bufferMode Enum) {
	C.callTransformFeedbackVaryings((C.GLuint)(program), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(varyings)), (C.GLenum)(bufferMode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPointParameterfv.xml
func PointParameterfv(pname Enum, params *Float) {
	C.callPointParameterfv((C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1iv.xml
func VertexAttribI1iv(index Uint, v *Int) {
	C.callVertexAttribI1iv((C.GLuint)(index), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage2D.xml
func TexStorage2D(target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei) {
	C.callTexStorage2D((C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4usv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4usv.xml
func VertexAttribI4usv(index Uint, v *Ushort) {
	C.callVertexAttribI4usv((C.GLuint)(index), (*C.GLushort)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glRenderbufferStorageMultisample
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glRenderbufferStorageMultisample.xml
func RenderbufferStorageMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei) {
	C.callRenderbufferStorageMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBufferData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBufferData.xml
func BufferData(target Enum, size Sizeiptr, data Ptr, usage Enum) {
	C.callBufferData((C.GLenum)(target), (C.GLsizeiptr)(size), (unsafe.Pointer)(data), (C.GLenum)(usage))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformBlockBinding
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformBlockBinding.xml
func UniformBlockBinding(program Uint, uniformBlockIndex Uint, uniformBlockBinding Uint) {
	C.callUniformBlockBinding((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLuint)(uniformBlockBinding))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformuiv.xml
func GetUniformuiv(program Uint, location Int, params *Uint) {
	C.callGetUniformuiv((C.GLuint)(program), (C.GLint)(location), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInternalformati64v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetInternalformati64v.xml
func GetInternalformati64v(target Enum, internalformat Enum, pname Enum, bufSize Sizei, params *Int64) {
	C.callGetInternalformati64v((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint64)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterIuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterIuiv.xml
func TexParameterIuiv(target Enum, pname Enum, params *Uint) {
	C.callTexParameterIuiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x4dv.xml
func UniformMatrix3x4dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix3x4dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFloati_v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFloati_v.xml
func GetFloati_v(target Enum, index Uint, data *Float) {
	C.callGetFloati_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLfloat)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCreateProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCreateProgram.xml
func CreateProgram() Uint {
	return (Uint)(C.callCreateProgram())
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsProgramPipeline
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsProgramPipeline.xml
func IsProgramPipeline(pipeline Uint) Boolean {
	return (Boolean)(C.callIsProgramPipeline((C.GLuint)(pipeline)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4dv.xml
func ProgramUniform4dv(program Uint, location Int, count Sizei, value *Double) {
	C.callProgramUniform4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInteger64v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetInteger64v.xml
func GetInteger64v(pname Enum, params *Int64) {
	C.callGetInteger64v((C.GLenum)(pname), (*C.GLint64)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramInterfaceiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramInterfaceiv.xml
func GetProgramInterfaceiv(program Uint, programInterface Enum, pname Enum, params *Int) {
	C.callGetProgramInterfaceiv((C.GLuint)(program), (C.GLenum)(programInterface), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glNormalP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glNormalP3uiv.xml
func NormalP3uiv(type_ Enum, coords *Uint) {
	C.callNormalP3uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glShaderBinary
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glShaderBinary.xml
func ShaderBinary(count Sizei, shaders *Uint, binaryformat Enum, binary Ptr, length Sizei) {
	C.callShaderBinary((C.GLsizei)(count), (*C.GLuint)(shaders), (C.GLenum)(binaryformat), (unsafe.Pointer)(binary), (C.GLsizei)(length))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPatchParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPatchParameterfv.xml
func PatchParameterfv(pname Enum, values *Float) {
	C.callPatchParameterfv((C.GLenum)(pname), (*C.GLfloat)(values))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameterfv.xml
func GetSamplerParameterfv(sampler Uint, pname Enum, params *Float) {
	C.callGetSamplerParameterfv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramPipelineInfoLog
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramPipelineInfoLog.xml
func GetProgramPipelineInfoLog(pipeline Uint, bufSize Sizei, length *Sizei, infoLog *Char) {
	C.callGetProgramPipelineInfoLog((C.GLuint)(pipeline), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glObjectLabel
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glObjectLabel.xml
func ObjectLabel(identifier Enum, name Uint, length Sizei, label *Char) {
	C.callObjectLabel((C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(length), (*C.GLchar)(label))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4uiv.xml
func VertexAttribI4uiv(index Uint, v *Uint) {
	C.callVertexAttribI4uiv((C.GLuint)(index), (*C.GLuint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyImageSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyImageSubData.xml
func CopyImageSubData(srcName Uint, srcTarget Enum, srcLevel Int, srcX Int, srcY Int, srcZ Int, dstName Uint, dstTarget Enum, dstLevel Int, dstX Int, dstY Int, dstZ Int, srcWidth Sizei, srcHeight Sizei, srcDepth Sizei) {
	C.callCopyImageSubData((C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(srcWidth), (C.GLsizei)(srcHeight), (C.GLsizei)(srcDepth))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquationi
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquationi.xml
func BlendEquationi(buf Uint, mode Enum) {
	C.callBlendEquationi((C.GLuint)(buf), (C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x2dv.xml
func ProgramUniformMatrix3x2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix3x2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2iv.xml
func Uniform2iv(location Int, count Sizei, value *Int) {
	C.callUniform2iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEnable
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEnable.xml
func Enable(cap Enum) {
	C.callEnable((C.GLenum)(cap))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2dv.xml
func Uniform2dv(location Int, count Sizei, value *Double) {
	C.callUniform2dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArraysInstanced
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawArraysInstanced.xml
func DrawArraysInstanced(mode Enum, first Int, count Sizei, instancecount Sizei) {
	C.callDrawArraysInstanced((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthFunc
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthFunc.xml
func DepthFunc(func_ Enum) {
	C.callDepthFunc((C.GLenum)(func_))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorP4uiv.xml
func ColorP4uiv(type_ Enum, color *Uint) {
	C.callColorP4uiv((C.GLenum)(type_), (*C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferParameteri.xml
func FramebufferParameteri(target Enum, pname Enum, param Int) {
	C.callFramebufferParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3ui.xml
func VertexAttribI3ui(index Uint, x Uint, y Uint, z Uint) {
	C.callVertexAttribI3ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glQueryCounter
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glQueryCounter.xml
func QueryCounter(id Uint, target Enum) {
	C.callQueryCounter((C.GLuint)(id), (C.GLenum)(target))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x4fv.xml
func ProgramUniformMatrix3x4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix3x4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFuncSeparate
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendFuncSeparate.xml
func BlendFuncSeparate(sfactorRGB Enum, dfactorRGB Enum, sfactorAlpha Enum, dfactorAlpha Enum) {
	C.callBlendFuncSeparate((C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyBufferSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyBufferSubData.xml
func CopyBufferSubData(readTarget Enum, writeTarget Enum, readOffset Intptr, writeOffset Intptr, size Sizeiptr) {
	C.callCopyBufferSubData((C.GLenum)(readTarget), (C.GLenum)(writeTarget), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizeiptr)(size))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3d.xml
func Uniform3d(location Int, x Double, y Double, z Double) {
	C.callUniform3d((C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawRangeElements
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawRangeElements.xml
func DrawRangeElements(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Ptr) {
	C.callDrawRangeElements((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribIPointer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribIPointer.xml
func VertexAttribIPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Ptr) {
	C.callVertexAttribIPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterIuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterIuiv.xml
func SamplerParameterIuiv(sampler Uint, pname Enum, param *Uint) {
	C.callSamplerParameterIuiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniform
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniform.xml
func GetActiveUniform(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char) {
	C.callGetActiveUniform((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexSubImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage2D.xml
func CopyTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, x Int, y Int, width Sizei, height Sizei) {
	C.callCopyTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformdv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformdv.xml
func GetUniformdv(program Uint, location Int, params *Double) {
	C.callGetUniformdv((C.GLuint)(program), (C.GLint)(location), (*C.GLdouble)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1ui.xml
func VertexAttribI1ui(index Uint, x Uint) {
	C.callVertexAttribI1ui((C.GLuint)(index), (C.GLuint)(x))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4iv.xml
func ProgramUniform4iv(program Uint, location Int, count Sizei, value *Int) {
	C.callProgramUniform4iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateTexSubImage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateTexSubImage.xml
func InvalidateTexSubImage(texture Uint, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei) {
	C.callInvalidateTexSubImage((C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x3fv.xml
func ProgramUniformMatrix4x3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix4x3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorP4ui.xml
func ColorP4ui(type_ Enum, color Uint) {
	C.callColorP4ui((C.GLenum)(type_), (C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2fv.xml
func UniformMatrix2fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferfv.xml
func ClearBufferfv(buffer Enum, drawbuffer Int, value *Float) {
	C.callClearBufferfv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawElements
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawElements.xml
func MultiDrawElements(mode Enum, count *Sizei, type_ Enum, indices *Ptr, drawcount Sizei) {
	C.callMultiDrawElements((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(type_), (*unsafe.Pointer)(indices), (C.GLsizei)(drawcount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexImage3D.xml
func TexImage3D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, depth Sizei, border Int, format Enum, type_ Enum, pixels Ptr) {
	C.callTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferSubData.xml
func ClearBufferSubData(target Enum, internalformat Enum, offset Intptr, size Sizeiptr, format Enum, type_ Enum, data Ptr) {
	C.callClearBufferSubData((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveAtomicCounterBufferiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveAtomicCounterBufferiv.xml
func GetActiveAtomicCounterBufferiv(program Uint, bufferIndex Uint, pname Enum, params *Int) {
	C.callGetActiveAtomicCounterBufferiv((C.GLuint)(program), (C.GLuint)(bufferIndex), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewportIndexedfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glViewportIndexedfv.xml
func ViewportIndexedfv(index Uint, v *Float) {
	C.callViewportIndexedfv((C.GLuint)(index), (*C.GLfloat)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture2D.xml
func FramebufferTexture2D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int) {
	C.callFramebufferTexture2D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindBuffer.xml
func BindBuffer(target Enum, buffer Uint) {
	C.callBindBuffer((C.GLenum)(target), (C.GLuint)(buffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3uiv.xml
func VertexAttribI3uiv(index Uint, v *Uint) {
	C.callVertexAttribI3uiv((C.GLuint)(index), (*C.GLuint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4i.xml
func VertexAttribI4i(index Uint, x Int, y Int, z Int, w Int) {
	C.callVertexAttribI4i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexSubImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage1D.xml
func CompressedTexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, imageSize Sizei, data Ptr) {
	C.callCompressedTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferSubData.xml
func GetBufferSubData(target Enum, offset Intptr, size Sizeiptr, data Ptr) {
	C.callGetBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteFramebuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteFramebuffers.xml
func DeleteFramebuffers(n Sizei, framebuffers *Uint) {
	C.callDeleteFramebuffers((C.GLsizei)(n), (*C.GLuint)(framebuffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribLFormat
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribLFormat.xml
func VertexAttribLFormat(attribindex Uint, size Int, type_ Enum, relativeoffset Uint) {
	C.callVertexAttribLFormat((C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(type_), (C.GLuint)(relativeoffset))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x3fv.xml
func UniformMatrix2x3fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix2x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceIndex.xml
func GetProgramResourceIndex(program Uint, programInterface Enum, name *Char) Uint {
	return (Uint)(C.callGetProgramResourceIndex((C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDisableVertexAttribArray
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDisableVertexAttribArray.xml
func DisableVertexAttribArray(index Uint) {
	C.callDisableVertexAttribArray((C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindBufferRange
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindBufferRange.xml
func BindBufferRange(target Enum, index Uint, buffer Uint, offset Intptr, size Sizeiptr) {
	C.callBindBufferRange((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2uiv.xml
func VertexAttribI2uiv(index Uint, v *Uint) {
	C.callVertexAttribI2uiv((C.GLuint)(index), (*C.GLuint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderPrecisionFormat
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderPrecisionFormat.xml
func GetShaderPrecisionFormat(shadertype Enum, precisiontype Enum, range_ *Int, precision *Int) {
	C.callGetShaderPrecisionFormat((C.GLenum)(shadertype), (C.GLenum)(precisiontype), (*C.GLint)(range_), (*C.GLint)(precision))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP2ui.xml
func VertexAttribP2ui(index Uint, type_ Enum, normalized Boolean, value Uint) {
	C.callVertexAttribP2ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProvokingVertex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProvokingVertex.xml
func ProvokingVertex(mode Enum) {
	C.callProvokingVertex((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexBuffer.xml
func TexBuffer(target Enum, internalformat Enum, buffer Uint) {
	C.callTexBuffer((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetString
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetString.xml
func GetString(name Enum) *Ubyte {
	return (*Ubyte)(C.callGetString((C.GLenum)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompileShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompileShader.xml
func CompileShader(shader Uint) {
	C.callCompileShader((C.GLuint)(shader))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1f.xml
func Uniform1f(location Int, v0 Float) {
	C.callUniform1f((C.GLint)(location), (C.GLfloat)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3f.xml
func Uniform3f(location Int, v0 Float, v1 Float, v2 Float) {
	C.callUniform3f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetObjectPtrLabel
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetObjectPtrLabel.xml
func GetObjectPtrLabel(ptr Ptr, bufSize Sizei, length *Sizei, label *Char) {
	C.callGetObjectPtrLabel((unsafe.Pointer)(ptr), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(label))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1i.xml
func ProgramUniform1i(program Uint, location Int, v0 Int) {
	C.callProgramUniform1i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2ui.xml
func VertexAttribI2ui(index Uint, x Uint, y Uint) {
	C.callVertexAttribI2ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x2dv.xml
func ProgramUniformMatrix4x2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix4x2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexImage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage3D.xml
func CompressedTexImage3D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, depth Sizei, border Int, imageSize Sizei, data Ptr) {
	C.callCompressedTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearColor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearColor.xml
func ClearColor(red Float, green Float, blue Float, alpha Float) {
	C.callClearColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP3ui.xml
func TexCoordP3ui(type_ Enum, coords Uint) {
	C.callTexCoordP3ui((C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterIiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterIiv.xml
func SamplerParameterIiv(sampler Uint, pname Enum, param *Int) {
	C.callSamplerParameterIiv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDispatchComputeIndirect
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDispatchComputeIndirect.xml
func DispatchComputeIndirect(indirect Intptr) {
	C.callDispatchComputeIndirect((C.GLintptr)(indirect))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteSync
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteSync.xml
func DeleteSync(sync Sync) {
	C.callDeleteSync((C.GLsync)(sync))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2i.xml
func VertexAttribI2i(index Uint, x Int, y Int) {
	C.callVertexAttribI2i((C.GLuint)(index), (C.GLint)(x), (C.GLint)(y))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilFunc
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilFunc.xml
func StencilFunc(func_ Enum, ref Int, mask Uint) {
	C.callStencilFunc((C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetDebugMessageLog
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetDebugMessageLog.xml
func GetDebugMessageLog(count Uint, bufsize Sizei, sources *Enum, types *Enum, ids *Uint, severities *Enum, lengths *Sizei, messageLog *Char) Uint {
	return (Uint)(C.callGetDebugMessageLog((C.GLuint)(count), (C.GLsizei)(bufsize), (*C.GLenum)(sources), (*C.GLenum)(types), (*C.GLuint)(ids), (*C.GLenum)(severities), (*C.GLsizei)(lengths), (*C.GLchar)(messageLog)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformSubroutineuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformSubroutineuiv.xml
func GetUniformSubroutineuiv(shadertype Enum, location Int, params *Uint) {
	C.callGetUniformSubroutineuiv((C.GLenum)(shadertype), (C.GLint)(location), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFragDataIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFragDataIndex.xml
func GetFragDataIndex(program Uint, name *Char) Int {
	return (Int)(C.callGetFragDataIndex((C.GLuint)(program), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilFuncSeparate
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilFuncSeparate.xml
func StencilFuncSeparate(face Enum, func_ Enum, ref Int, mask Uint) {
	C.callStencilFuncSeparate((C.GLenum)(face), (C.GLenum)(func_), (C.GLint)(ref), (C.GLuint)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexSubImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage2D.xml
func TexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Ptr) {
	C.callTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPointParameteri.xml
func PointParameteri(pname Enum, param Int) {
	C.callPointParameteri((C.GLenum)(pname), (C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsShader.xml
func IsShader(shader Uint) Boolean {
	return (Boolean)(C.callIsShader((C.GLuint)(shader)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x2fv.xml
func UniformMatrix4x2fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix4x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryIndexediv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryIndexediv.xml
func GetQueryIndexediv(target Enum, index Uint, pname Enum, params *Int) {
	C.callGetQueryIndexediv((C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformLocation.xml
func GetUniformLocation(program Uint, name *Char) Int {
	return (Int)(C.callGetUniformLocation((C.GLuint)(program), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glActiveTexture
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glActiveTexture.xml
func ActiveTexture(texture Enum) {
	C.callActiveTexture((C.GLenum)(texture))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP4ui.xml
func MultiTexCoordP4ui(texture Enum, type_ Enum, coords Uint) {
	C.callMultiTexCoordP4ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramStageiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramStageiv.xml
func GetProgramStageiv(program Uint, shadertype Enum, pname Enum, values *Int) {
	C.callGetProgramStageiv((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLenum)(pname), (*C.GLint)(values))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1dv.xml
func ProgramUniform1dv(program Uint, location Int, count Sizei, value *Double) {
	C.callProgramUniform1dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glShaderStorageBlockBinding
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glShaderStorageBlockBinding.xml
func ShaderStorageBlockBinding(program Uint, storageBlockIndex Uint, storageBlockBinding Uint) {
	C.callShaderStorageBlockBinding((C.GLuint)(program), (C.GLuint)(storageBlockIndex), (C.GLuint)(storageBlockBinding))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage2DMultisample
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage2DMultisample.xml
func TexStorage2DMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei, fixedsamplelocations Boolean) {
	C.callTexStorage2DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(fixedsamplelocations))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquationSeparate
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquationSeparate.xml
func BlendEquationSeparate(modeRGB Enum, modeAlpha Enum) {
	C.callBlendEquationSeparate((C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glNormalP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glNormalP3ui.xml
func NormalP3ui(type_ Enum, coords Uint) {
	C.callNormalP3ui((C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x4fv.xml
func UniformMatrix3x4fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix3x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1i.xml
func VertexAttribI1i(index Uint, x Int) {
	C.callVertexAttribI1i((C.GLuint)(index), (C.GLint)(x))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2ui.xml
func Uniform2ui(location Int, v0 Uint, v1 Uint) {
	C.callUniform2ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterfv.xml
func GetTexParameterfv(target Enum, pname Enum, params *Float) {
	C.callGetTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenRenderbuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenRenderbuffers.xml
func GenRenderbuffers(n Sizei, renderbuffers *Uint) {
	C.callGenRenderbuffers((C.GLsizei)(n), (*C.GLuint)(renderbuffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsSampler
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsSampler.xml
func IsSampler(sampler Uint) Boolean {
	return (Boolean)(C.callIsSampler((C.GLuint)(sampler)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndQueryIndexed
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEndQueryIndexed.xml
func EndQueryIndexed(target Enum, index Uint) {
	C.callEndQueryIndexed((C.GLenum)(target), (C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPixelStoref
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPixelStoref.xml
func PixelStoref(pname Enum, param Float) {
	C.callPixelStoref((C.GLenum)(pname), (C.GLfloat)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribDivisor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribDivisor.xml
func VertexAttribDivisor(index Uint, divisor Uint) {
	C.callVertexAttribDivisor((C.GLuint)(index), (C.GLuint)(divisor))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilOp
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilOp.xml
func StencilOp(fail Enum, zfail Enum, zpass Enum) {
	C.callStencilOp((C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateBufferSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateBufferSubData.xml
func InvalidateBufferSubData(buffer Uint, offset Intptr, length Sizeiptr) {
	C.callInvalidateBufferSubData((C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL3dv.xml
func VertexAttribL3dv(index Uint, v *Double) {
	C.callVertexAttribL3dv((C.GLuint)(index), (*C.GLdouble)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferiv.xml
func ClearBufferiv(buffer Enum, drawbuffer Int, value *Int) {
	C.callClearBufferiv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glLogicOp
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glLogicOp.xml
func LogicOp(opcode Enum) {
	C.callLogicOp((C.GLenum)(opcode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexSubImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage1D.xml
func CopyTexSubImage1D(target Enum, level Int, xoffset Int, x Int, y Int, width Sizei) {
	C.callCopyTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribdv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribdv.xml
func GetVertexAttribdv(index Uint, pname Enum, params *Double) {
	C.callGetVertexAttribdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawBuffer.xml
func DrawBuffer(mode Enum) {
	C.callDrawBuffer((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenBuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenBuffers.xml
func GenBuffers(n Sizei, buffers *Uint) {
	C.callGenBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage3D.xml
func TexStorage3D(target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei, depth Sizei) {
	C.callTexStorage3D((C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage3DMultisample
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage3DMultisample.xml
func TexStorage3DMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean) {
	C.callTexStorage3DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(fixedsamplelocations))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstanced
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstanced.xml
func DrawElementsInstanced(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei) {
	C.callDrawElementsInstanced((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPolygonOffset
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPolygonOffset.xml
func PolygonOffset(factor Float, units Float) {
	C.callPolygonOffset((C.GLfloat)(factor), (C.GLfloat)(units))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMapBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMapBuffer.xml
func MapBuffer(target Enum, access Enum) *Ptr {
	return (*Ptr)(C.callMapBuffer((C.GLenum)(target), (C.GLenum)(access)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFragDataLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFragDataLocation.xml
func GetFragDataLocation(program Uint, name *Char) Int {
	return (Int)(C.callGetFragDataLocation((C.GLuint)(program), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsTexture
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsTexture.xml
func IsTexture(texture Uint) Boolean {
	return (Boolean)(C.callIsTexture((C.GLuint)(texture)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteVertexArrays
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteVertexArrays.xml
func DeleteVertexArrays(n Sizei, arrays *Uint) {
	C.callDeleteVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewportArrayv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glViewportArrayv.xml
func ViewportArrayv(first Uint, count Sizei, v *Float) {
	C.callViewportArrayv((C.GLuint)(first), (C.GLsizei)(count), (*C.GLfloat)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedbackInstanced
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedbackInstanced.xml
func DrawTransformFeedbackInstanced(mode Enum, id Uint, instancecount Sizei) {
	C.callDrawTransformFeedbackInstanced((C.GLenum)(mode), (C.GLuint)(id), (C.GLsizei)(instancecount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDebugMessageInsert
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDebugMessageInsert.xml
func DebugMessageInsert(source Enum, type_ Enum, id Uint, severity Enum, length Sizei, buf *Char) {
	C.callDebugMessageInsert((C.GLenum)(source), (C.GLenum)(type_), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(buf))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribiv.xml
func GetVertexAttribiv(index Uint, pname Enum, params *Int) {
	C.callGetVertexAttribiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFloatv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFloatv.xml
func GetFloatv(pname Enum, params *Float) {
	C.callGetFloatv((C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClientWaitSync
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClientWaitSync.xml
func ClientWaitSync(sync Sync, flags Bitfield, timeout Uint64) Enum {
	return (Enum)(C.callClientWaitSync((C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2f.xml
func Uniform2f(location Int, v0 Float, v1 Float) {
	C.callUniform2f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsBuffer.xml
func IsBuffer(buffer Uint) Boolean {
	return (Boolean)(C.callIsBuffer((C.GLuint)(buffer)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1fv.xml
func Uniform1fv(location Int, count Sizei, value *Float) {
	C.callUniform1fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP1uiv.xml
func MultiTexCoordP1uiv(texture Enum, type_ Enum, coords *Uint) {
	C.callMultiTexCoordP1uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4ui.xml
func VertexAttribI4ui(index Uint, x Uint, y Uint, z Uint, w Uint) {
	C.callVertexAttribI4ui((C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBooleanv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBooleanv.xml
func GetBooleanv(pname Enum, params *Boolean) {
	C.callGetBooleanv((C.GLenum)(pname), (*C.GLboolean)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderiv.xml
func GetShaderiv(shader Uint, pname Enum, params *Int) {
	C.callGetShaderiv((C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2uiv.xml
func ProgramUniform2uiv(program Uint, location Int, count Sizei, value *Uint) {
	C.callProgramUniform2uiv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP1uiv.xml
func TexCoordP1uiv(type_ Enum, coords *Uint) {
	C.callTexCoordP1uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndConditionalRender
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEndConditionalRender.xml
func EndConditionalRender() {
	C.callEndConditionalRender()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindImageTexture
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindImageTexture.xml
func BindImageTexture(unit Uint, texture Uint, level Int, layered Boolean, layer Int, access Enum, format Enum) {
	C.callBindImageTexture((C.GLuint)(unit), (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(layered), (C.GLint)(layer), (C.GLenum)(access), (C.GLenum)(format))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4i.xml
func ProgramUniform4i(program Uint, location Int, v0 Int, v1 Int, v2 Int, v3 Int) {
	C.callProgramUniform4i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferParameteri64v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferParameteri64v.xml
func GetBufferParameteri64v(target Enum, pname Enum, params *Int64) {
	C.callGetBufferParameteri64v((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint64)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glRenderbufferStorage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glRenderbufferStorage.xml
func RenderbufferStorage(target Enum, internalformat Enum, width Sizei, height Sizei) {
	C.callRenderbufferStorage((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRangeArrayv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthRangeArrayv.xml
func DepthRangeArrayv(first Uint, count Sizei, v *Double) {
	C.callDepthRangeArrayv((C.GLuint)(first), (C.GLsizei)(count), (*C.GLdouble)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture.xml
func FramebufferTexture(target Enum, attachment Enum, texture Uint, level Int) {
	C.callFramebufferTexture((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceiv.xml
func GetProgramResourceiv(program Uint, programInterface Enum, index Uint, propCount Sizei, props *Enum, bufSize Sizei, length *Sizei, params *Int) {
	C.callGetProgramResourceiv((C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(propCount), (*C.GLenum)(props), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP2uiv.xml
func MultiTexCoordP2uiv(texture Enum, type_ Enum, coords *Uint) {
	C.callMultiTexCoordP2uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindSampler
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindSampler.xml
func BindSampler(unit Uint, sampler Uint) {
	C.callBindSampler((C.GLuint)(unit), (C.GLuint)(sampler))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x3fv.xml
func UniformMatrix4x3fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix4x3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribIiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribIiv.xml
func GetVertexAttribIiv(index Uint, pname Enum, params *Int) {
	C.callGetVertexAttribIiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMemoryBarrier
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMemoryBarrier.xml
func MemoryBarrier(barriers Bitfield) {
	C.callMemoryBarrier((C.GLbitfield)(barriers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginQueryIndexed
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBeginQueryIndexed.xml
func BeginQueryIndexed(target Enum, index Uint, id Uint) {
	C.callBeginQueryIndexed((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(id))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2d.xml
func Uniform2d(location Int, x Double, y Double) {
	C.callUniform2d((C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteRenderbuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteRenderbuffers.xml
func DeleteRenderbuffers(n Sizei, renderbuffers *Uint) {
	C.callDeleteRenderbuffers((C.GLsizei)(n), (*C.GLuint)(renderbuffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUseProgramStages
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUseProgramStages.xml
func UseProgramStages(pipeline Uint, stages Bitfield, program Uint) {
	C.callUseProgramStages((C.GLuint)(pipeline), (C.GLbitfield)(stages), (C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawBuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawBuffers.xml
func DrawBuffers(n Sizei, bufs *Enum) {
	C.callDrawBuffers((C.GLsizei)(n), (*C.GLenum)(bufs))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSynciv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSynciv.xml
func GetSynciv(sync Sync, pname Enum, bufSize Sizei, length *Sizei, values *Int) {
	C.callGetSynciv((C.GLsync)(sync), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(values))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndQuery
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEndQuery.xml
func EndQuery(target Enum) {
	C.callEndQuery((C.GLenum)(target))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRangeIndexed
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthRangeIndexed.xml
func DepthRangeIndexed(index Uint, n Double, f Double) {
	C.callDepthRangeIndexed((C.GLuint)(index), (C.GLdouble)(n), (C.GLdouble)(f))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetError
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetError.xml
func GetError() Enum {
	return (Enum)(C.callGetError())
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4d.xml
func ProgramUniform4d(program Uint, location Int, v0 Double, v1 Double, v2 Double, v3 Double) {
	C.callProgramUniform4d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLdouble)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPushDebugGroup
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPushDebugGroup.xml
func PushDebugGroup(source Enum, id Uint, length Sizei, message *Char) {
	C.callPushDebugGroup((C.GLenum)(source), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(message))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryiv.xml
func GetQueryiv(target Enum, pname Enum, params *Int) {
	C.callGetQueryiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsEnabled
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsEnabled.xml
func IsEnabled(cap Enum) Boolean {
	return (Boolean)(C.callIsEnabled((C.GLenum)(cap)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSecondaryColorP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColorP3uiv.xml
func SecondaryColorP3uiv(type_ Enum, color *Uint) {
	C.callSecondaryColorP3uiv((C.GLenum)(type_), (*C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexSubImage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage3D.xml
func TexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, type_ Enum, pixels Ptr) {
	C.callTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDisablei
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDisablei.xml
func Disablei(target Enum, index Uint) {
	C.callDisablei((C.GLenum)(target), (C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawRangeElementsBaseVertex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawRangeElementsBaseVertex.xml
func DrawRangeElementsBaseVertex(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Ptr, basevertex Int) {
	C.callDrawRangeElementsBaseVertex((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLint)(basevertex))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindTexture
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindTexture.xml
func BindTexture(target Enum, texture Uint) {
	C.callBindTexture((C.GLenum)(target), (C.GLuint)(texture))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4iv.xml
func Uniform4iv(location Int, count Sizei, value *Int) {
	C.callUniform4iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMinSampleShading
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMinSampleShading.xml
func MinSampleShading(value Float) {
	C.callMinSampleShading((C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1dv.xml
func Uniform1dv(location Int, count Sizei, value *Double) {
	C.callUniform1dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP1ui.xml
func VertexAttribP1ui(index Uint, type_ Enum, normalized Boolean, value Uint) {
	C.callVertexAttribP1ui((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPixelStorei
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPixelStorei.xml
func PixelStorei(pname Enum, param Int) {
	C.callPixelStorei((C.GLenum)(pname), (C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFrontFace
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFrontFace.xml
func FrontFace(mode Enum) {
	C.callFrontFace((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2d.xml
func ProgramUniform2d(program Uint, location Int, v0 Double, v1 Double) {
	C.callProgramUniform2d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage1D.xml
func TexStorage1D(target Enum, levels Sizei, internalformat Enum, width Sizei) {
	C.callTexStorage1D((C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3i.xml
func Uniform3i(location Int, v0 Int, v1 Int, v2 Int) {
	C.callUniform3i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage2DMultisample
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexImage2DMultisample.xml
func TexImage2DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, fixedsamplelocations Boolean) {
	C.callTexImage2DMultisample((C.GLenum)(target), (C.GLsizei)(samples), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(fixedsamplelocations))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x2fv.xml
func UniformMatrix3x2fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix3x2fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexImage2D.xml
func TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, type_ Enum, pixels Ptr) {
	C.callTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSubroutineUniformLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSubroutineUniformLocation.xml
func GetSubroutineUniformLocation(program Uint, shadertype Enum, name *Char) Int {
	return (Int)(C.callGetSubroutineUniformLocation((C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3dv.xml
func UniformMatrix3dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix3dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindRenderbuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindRenderbuffer.xml
func BindRenderbuffer(target Enum, renderbuffer Uint) {
	C.callBindRenderbuffer((C.GLenum)(target), (C.GLuint)(renderbuffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindProgramPipeline
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindProgramPipeline.xml
func BindProgramPipeline(pipeline Uint) {
	C.callBindProgramPipeline((C.GLuint)(pipeline))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4dv.xml
func Uniform4dv(location Int, count Sizei, value *Double) {
	C.callUniform4dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferuiv.xml
func ClearBufferuiv(buffer Enum, drawbuffer Int, value *Uint) {
	C.callClearBufferuiv((C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBeginTransformFeedback.xml
func BeginTransformFeedback(primitiveMode Enum) {
	C.callBeginTransformFeedback((C.GLenum)(primitiveMode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3fv.xml
func UniformMatrix3fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix3fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDispatchCompute
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDispatchCompute.xml
func DispatchCompute(num_groups_x Uint, num_groups_y Uint, num_groups_z Uint) {
	C.callDispatchCompute((C.GLuint)(num_groups_x), (C.GLuint)(num_groups_y), (C.GLuint)(num_groups_z))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindVertexBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindVertexBuffer.xml
func BindVertexBuffer(bindingindex Uint, buffer Uint, offset Intptr, stride Sizei) {
	C.callBindVertexBuffer((C.GLuint)(bindingindex), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(stride))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexBindingDivisor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexBindingDivisor.xml
func VertexBindingDivisor(bindingindex Uint, divisor Uint) {
	C.callVertexBindingDivisor((C.GLuint)(bindingindex), (C.GLuint)(divisor))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP4uiv.xml
func MultiTexCoordP4uiv(texture Enum, type_ Enum, coords *Uint) {
	C.callMultiTexCoordP4uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x3fv.xml
func ProgramUniformMatrix2x3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix2x3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexLevelParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexLevelParameteriv.xml
func GetTexLevelParameteriv(target Enum, level Int, pname Enum, params *Int) {
	C.callGetTexLevelParameteriv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP4uiv.xml
func VertexP4uiv(type_ Enum, value *Uint) {
	C.callVertexP4uiv((C.GLenum)(type_), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameterIuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterIuiv.xml
func GetTexParameterIuiv(target Enum, pname Enum, params *Uint) {
	C.callGetTexParameterIuiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDebugMessageControl
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDebugMessageControl.xml
func DebugMessageControl(source Enum, type_ Enum, severity Enum, count Sizei, ids *Uint, enabled Boolean) {
	C.callDebugMessageControl((C.GLenum)(source), (C.GLenum)(type_), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(ids), (C.GLboolean)(enabled))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameteri.xml
func SamplerParameteri(sampler Uint, pname Enum, param Int) {
	C.callSamplerParameteri((C.GLuint)(sampler), (C.GLenum)(pname), (C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2f.xml
func ProgramUniform2f(program Uint, location Int, v0 Float, v1 Float) {
	C.callProgramUniform2f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribfv.xml
func GetVertexAttribfv(index Uint, pname Enum, params *Float) {
	C.callGetVertexAttribfv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP2uiv.xml
func VertexAttribP2uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) {
	C.callVertexAttribP2uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissorIndexed
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glScissorIndexed.xml
func ScissorIndexed(index Uint, left Int, bottom Int, width Sizei, height Sizei) {
	C.callScissorIndexed((C.GLuint)(index), (C.GLint)(left), (C.GLint)(bottom), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1d.xml
func ProgramUniform1d(program Uint, location Int, v0 Double) {
	C.callProgramUniform1d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP4ui.xml
func VertexP4ui(type_ Enum, value Uint) {
	C.callVertexP4ui((C.GLenum)(type_), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthMask
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthMask.xml
func DepthMask(flag Boolean) {
	C.callDepthMask((C.GLboolean)(flag))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2dv.xml
func UniformMatrix2dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix2dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP4uiv.xml
func TexCoordP4uiv(type_ Enum, coords *Uint) {
	C.callTexCoordP4uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4f.xml
func Uniform4f(location Int, v0 Float, v1 Float, v2 Float, v3 Float) {
	C.callUniform4f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4fv.xml
func Uniform4fv(location Int, count Sizei, value *Float) {
	C.callUniform4fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClampColor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClampColor.xml
func ClampColor(target Enum, clamp Enum) {
	C.callClampColor((C.GLenum)(target), (C.GLenum)(clamp))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenFramebuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenFramebuffers.xml
func GenFramebuffers(n Sizei, framebuffers *Uint) {
	C.callGenFramebuffers((C.GLsizei)(n), (*C.GLuint)(framebuffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP3uiv.xml
func TexCoordP3uiv(type_ Enum, coords *Uint) {
	C.callTexCoordP3uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorMask
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorMask.xml
func ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean) {
	C.callColorMask((C.GLboolean)(red), (C.GLboolean)(green), (C.GLboolean)(blue), (C.GLboolean)(alpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterfv.xml
func SamplerParameterfv(sampler Uint, pname Enum, param *Float) {
	C.callSamplerParameterfv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsIndirect
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsIndirect.xml
func DrawElementsIndirect(mode Enum, type_ Enum, indirect Ptr) {
	C.callDrawElementsIndirect((C.GLenum)(mode), (C.GLenum)(type_), (unsafe.Pointer)(indirect))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformiv.xml
func GetUniformiv(program Uint, location Int, params *Int) {
	C.callGetUniformiv((C.GLuint)(program), (C.GLint)(location), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetRenderbufferParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetRenderbufferParameteriv.xml
func GetRenderbufferParameteriv(target Enum, pname Enum, params *Int) {
	C.callGetRenderbufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramParameteri.xml
func ProgramParameteri(program Uint, pname Enum, value Int) {
	C.callProgramParameteri((C.GLuint)(program), (C.GLenum)(pname), (C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2fv.xml
func Uniform2fv(location Int, count Sizei, value *Float) {
	C.callUniform2fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2dv.xml
func ProgramUniformMatrix2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4fv.xml
func ProgramUniform4fv(program Uint, location Int, count Sizei, value *Float) {
	C.callProgramUniform4fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterf
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterf.xml
func TexParameterf(target Enum, pname Enum, param Float) {
	C.callTexParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glWaitSync
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glWaitSync.xml
func WaitSync(sync Sync, flags Bitfield, timeout Uint64) {
	C.callWaitSync((C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilOpSeparate
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilOpSeparate.xml
func StencilOpSeparate(face Enum, sfail Enum, dpfail Enum, dppass Enum) {
	C.callStencilOpSeparate((C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP1uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP1uiv.xml
func VertexAttribP1uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) {
	C.callVertexAttribP1uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP2ui.xml
func VertexP2ui(type_ Enum, value Uint) {
	C.callVertexP2ui((C.GLenum)(type_), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorMaski
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorMaski.xml
func ColorMaski(index Uint, r Boolean, g Boolean, b Boolean, a Boolean) {
	C.callColorMaski((C.GLuint)(index), (C.GLboolean)(r), (C.GLboolean)(g), (C.GLboolean)(b), (C.GLboolean)(a))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture3D.xml
func FramebufferTexture3D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int, zoffset Int) {
	C.callFramebufferTexture3D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(zoffset))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2uiv.xml
func Uniform2uiv(location Int, count Sizei, value *Uint) {
	C.callUniform2uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformBlockiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformBlockiv.xml
func GetActiveUniformBlockiv(program Uint, uniformBlockIndex Uint, pname Enum, params *Int) {
	C.callGetActiveUniformBlockiv((C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL4dv.xml
func VertexAttribL4dv(index Uint, v *Double) {
	C.callVertexAttribL4dv((C.GLuint)(index), (*C.GLdouble)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilMask
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilMask.xml
func StencilMask(mask Uint) {
	C.callStencilMask((C.GLuint)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawElementsIndirect
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawElementsIndirect.xml
func MultiDrawElementsIndirect(mode Enum, type_ Enum, indirect Ptr, drawcount Sizei, stride Sizei) {
	C.callMultiDrawElementsIndirect((C.GLenum)(mode), (C.GLenum)(type_), (unsafe.Pointer)(indirect), (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveSubroutineName
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveSubroutineName.xml
func GetActiveSubroutineName(program Uint, shadertype Enum, index Uint, bufsize Sizei, length *Sizei, name *Char) {
	C.callGetActiveSubroutineName((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(length), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetAttachedShaders
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetAttachedShaders.xml
func GetAttachedShaders(program Uint, maxCount Sizei, count *Sizei, obj *Uint) {
	C.callGetAttachedShaders((C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(count), (*C.GLuint)(obj))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearDepthf
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearDepthf.xml
func ClearDepthf(d Float) {
	C.callClearDepthf((C.GLfloat)(d))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearStencil
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearStencil.xml
func ClearStencil(s Int) {
	C.callClearStencil((C.GLint)(s))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetCompressedTexImage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetCompressedTexImage.xml
func GetCompressedTexImage(target Enum, level Int, img Ptr) {
	C.callGetCompressedTexImage((C.GLenum)(target), (C.GLint)(level), (unsafe.Pointer)(img))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPointParameteriv.xml
func PointParameteriv(pname Enum, params *Int) {
	C.callPointParameteriv((C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2i.xml
func ProgramUniform2i(program Uint, location Int, v0 Int, v1 Int) {
	C.callProgramUniform2i((C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexImage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexImage.xml
func GetTexImage(target Enum, level Int, format Enum, type_ Enum, pixels Ptr) {
	C.callGetTexImage((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindFragDataLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindFragDataLocation.xml
func BindFragDataLocation(program Uint, color Uint, name *Char) {
	C.callBindFragDataLocation((C.GLuint)(program), (C.GLuint)(color), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x3dv.xml
func UniformMatrix2x3dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix2x3dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFunci
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendFunci.xml
func BlendFunci(buf Uint, src Enum, dst Enum) {
	C.callBlendFunci((C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindAttribLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindAttribLocation.xml
func BindAttribLocation(program Uint, index Uint, name *Char) {
	C.callBindAttribLocation((C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage2D.xml
func CompressedTexImage2D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, border Int, imageSize Sizei, data Ptr) {
	C.callCompressedTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribLPointer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribLPointer.xml
func VertexAttribLPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Ptr) {
	C.callVertexAttribLPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP4ui.xml
func TexCoordP4ui(type_ Enum, coords Uint) {
	C.callTexCoordP4ui((C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPauseTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPauseTransformFeedback.xml
func PauseTransformFeedback() {
	C.callPauseTransformFeedback()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x4dv.xml
func ProgramUniformMatrix2x4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix2x4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage1D.xml
func CompressedTexImage1D(target Enum, level Int, internalformat Enum, width Sizei, border Int, imageSize Sizei, data Ptr) {
	C.callCompressedTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEndTransformFeedback.xml
func EndTransformFeedback() {
	C.callEndTransformFeedback()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3dv.xml
func Uniform3dv(location Int, count Sizei, value *Double) {
	C.callUniform3dv((C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissor
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glScissor.xml
func Scissor(x Int, y Int, width Sizei, height Sizei) {
	C.callScissor((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexSubImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage1D.xml
func TexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, type_ Enum, pixels Ptr) {
	C.callTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSampleMaski
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSampleMaski.xml
func SampleMaski(index Uint, mask Bitfield) {
	C.callSampleMaski((C.GLuint)(index), (C.GLbitfield)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFinish
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFinish.xml
func Finish() {
	C.callFinish()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1iv.xml
func ProgramUniform1iv(program Uint, location Int, count Sizei, value *Int) {
	C.callProgramUniform1iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissorIndexedv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glScissorIndexedv.xml
func ScissorIndexedv(index Uint, v *Int) {
	C.callScissorIndexedv((C.GLuint)(index), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameterIiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterIiv.xml
func GetTexParameterIiv(target Enum, pname Enum, params *Int) {
	C.callGetTexParameterIiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUseProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUseProgram.xml
func UseProgram(program Uint) {
	C.callUseProgram((C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateBufferData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateBufferData.xml
func InvalidateBufferData(buffer Uint) {
	C.callInvalidateBufferData((C.GLuint)(buffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissorArrayv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glScissorArrayv.xml
func ScissorArrayv(first Uint, count Sizei, v *Int) {
	C.callScissorArrayv((C.GLuint)(first), (C.GLsizei)(count), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramiv.xml
func GetProgramiv(program Uint, pname Enum, params *Int) {
	C.callGetProgramiv((C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDisable
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDisable.xml
func Disable(cap Enum) {
	C.callDisable((C.GLenum)(cap))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformBlockIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformBlockIndex.xml
func GetUniformBlockIndex(program Uint, uniformBlockName *Char) Uint {
	return (Uint)(C.callGetUniformBlockIndex((C.GLuint)(program), (*C.GLchar)(uniformBlockName)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquation.xml
func BlendEquation(mode Enum) {
	C.callBlendEquation((C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3fv.xml
func ProgramUniformMatrix3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix3fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArrays
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawArrays.xml
func DrawArrays(mode Enum, first Int, count Sizei) {
	C.callDrawArrays((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCreateShaderProgramv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCreateShaderProgramv.xml
func CreateShaderProgramv(type_ Enum, count Sizei, strings **Char) Uint {
	return (Uint)(C.callCreateShaderProgramv((C.GLenum)(type_), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(strings))))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenTransformFeedbacks
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenTransformFeedbacks.xml
func GenTransformFeedbacks(n Sizei, ids *Uint) {
	C.callGenTransformFeedbacks((C.GLsizei)(n), (*C.GLuint)(ids))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1i.xml
func Uniform1i(location Int, v0 Int) {
	C.callUniform1i((C.GLint)(location), (C.GLint)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL4d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL4d.xml
func VertexAttribL4d(index Uint, x Double, y Double, z Double, w Double) {
	C.callVertexAttribL4d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4iv.xml
func VertexAttribI4iv(index Uint, v *Int) {
	C.callVertexAttribI4iv((C.GLuint)(index), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2fv.xml
func ProgramUniformMatrix2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlitFramebuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlitFramebuffer.xml
func BlitFramebuffer(srcX0 Int, srcY0 Int, srcX1 Int, srcY1 Int, dstX0 Int, dstY0 Int, dstX1 Int, dstY1 Int, mask Bitfield, filter Enum) {
	C.callBlitFramebuffer((C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateSubFramebuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateSubFramebuffer.xml
func InvalidateSubFramebuffer(target Enum, numAttachments Sizei, attachments *Enum, x Int, y Int, width Sizei, height Sizei) {
	C.callInvalidateSubFramebuffer((C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(attachments), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP3ui.xml
func MultiTexCoordP3ui(texture Enum, type_ Enum, coords Uint) {
	C.callMultiTexCoordP3ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glShaderSource
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glShaderSource.xml
func ShaderSource(shader Uint, count Sizei, string_ **Char, length *Int) {
	C.callShaderSource((C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(string_)), (*C.GLint)(length))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4dv.xml
func UniformMatrix4dv(location Int, count Sizei, transpose Boolean, value *Double) {
	C.callUniformMatrix4dv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4dv.xml
func ProgramUniformMatrix4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteSamplers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteSamplers.xml
func DeleteSamplers(count Sizei, samplers *Uint) {
	C.callDeleteSamplers((C.GLsizei)(count), (*C.GLuint)(samplers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenSamplers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenSamplers.xml
func GenSamplers(count Sizei, samplers *Uint) {
	C.callGenSamplers((C.GLsizei)(count), (*C.GLuint)(samplers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterIiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterIiv.xml
func TexParameterIiv(target Enum, pname Enum, params *Int) {
	C.callTexParameterIiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glAttachShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glAttachShader.xml
func AttachShader(program Uint, shader Uint) {
	C.callAttachShader((C.GLuint)(program), (C.GLuint)(shader))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glValidateProgramPipeline
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glValidateProgramPipeline.xml
func ValidateProgramPipeline(pipeline Uint) {
	C.callValidateProgramPipeline((C.GLuint)(pipeline))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferRenderbuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferRenderbuffer.xml
func FramebufferRenderbuffer(target Enum, attachment Enum, renderbuffertarget Enum, renderbuffer Uint) {
	C.callFramebufferRenderbuffer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDetachShader
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDetachShader.xml
func DetachShader(program Uint, shader Uint) {
	C.callDetachShader((C.GLuint)(program), (C.GLuint)(shader))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4d.xml
func Uniform4d(location Int, x Double, y Double, z Double, w Double) {
	C.callUniform4d((C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRange
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDepthRange.xml
func DepthRange(near_ Double, far_ Double) {
	C.callDepthRange((C.GLdouble)(near_), (C.GLdouble)(far_))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexBufferRange
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexBufferRange.xml
func TexBufferRange(target Enum, internalformat Enum, buffer Uint, offset Intptr, size Sizeiptr) {
	C.callTexBufferRange((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4uiv.xml
func Uniform4uiv(location Int, count Sizei, value *Uint) {
	C.callUniform4uiv((C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSubroutineIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSubroutineIndex.xml
func GetSubroutineIndex(program Uint, shadertype Enum, name *Char) Uint {
	return (Uint)(C.callGetSubroutineIndex((C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3iv.xml
func VertexAttribI3iv(index Uint, v *Int) {
	C.callVertexAttribI3iv((C.GLuint)(index), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPatchParameteri
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPatchParameteri.xml
func PatchParameteri(pname Enum, value Int) {
	C.callPatchParameteri((C.GLenum)(pname), (C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramInfoLog
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramInfoLog.xml
func GetProgramInfoLog(program Uint, bufSize Sizei, length *Sizei, infoLog *Char) {
	C.callGetProgramInfoLog((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFlushMappedBufferRange
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFlushMappedBufferRange.xml
func FlushMappedBufferRange(target Enum, offset Intptr, length Sizeiptr) {
	C.callFlushMappedBufferRange((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquationSeparatei
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquationSeparatei.xml
func BlendEquationSeparatei(buf Uint, modeRGB Enum, modeAlpha Enum) {
	C.callBlendEquationSeparatei((C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1ui.xml
func Uniform1ui(location Int, v0 Uint) {
	C.callUniform1ui((C.GLint)(location), (C.GLuint)(v0))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveSubroutineUniformName
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveSubroutineUniformName.xml
func GetActiveSubroutineUniformName(program Uint, shadertype Enum, index Uint, bufsize Sizei, length *Sizei, name *Char) {
	C.callGetActiveSubroutineUniformName((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(length), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL2d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL2d.xml
func VertexAttribL2d(index Uint, x Double, y Double) {
	C.callVertexAttribL2d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsRenderbuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsRenderbuffer.xml
func IsRenderbuffer(renderbuffer Uint) Boolean {
	return (Boolean)(C.callIsRenderbuffer((C.GLuint)(renderbuffer)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3f.xml
func ProgramUniform3f(program Uint, location Int, v0 Float, v1 Float, v2 Float) {
	C.callProgramUniform3f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewportIndexedf
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glViewportIndexedf.xml
func ViewportIndexedf(index Uint, x Float, y Float, w Float, h Float) {
	C.callViewportIndexedf((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(w), (C.GLfloat)(h))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glResumeTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glResumeTransformFeedback.xml
func ResumeTransformFeedback() {
	C.callResumeTransformFeedback()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP1ui.xml
func MultiTexCoordP1ui(texture Enum, type_ Enum, coords Uint) {
	C.callMultiTexCoordP1ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glValidateProgram
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glValidateProgram.xml
func ValidateProgram(program Uint) {
	C.callValidateProgram((C.GLuint)(program))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPopDebugGroup
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glPopDebugGroup.xml
func PopDebugGroup() {
	C.callPopDebugGroup()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP4uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP4uiv.xml
func VertexAttribP4uiv(index Uint, type_ Enum, normalized Boolean, value *Uint) {
	C.callVertexAttribP4uiv((C.GLuint)(index), (C.GLenum)(type_), (C.GLboolean)(normalized), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArraysIndirect
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawArraysIndirect.xml
func DrawArraysIndirect(mode Enum, indirect Ptr) {
	C.callDrawArraysIndirect((C.GLenum)(mode), (unsafe.Pointer)(indirect))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x3dv.xml
func ProgramUniformMatrix4x3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix4x3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElements
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElements.xml
func DrawElements(mode Enum, count Sizei, type_ Enum, indices Ptr) {
	C.callDrawElements((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearDepth
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearDepth.xml
func ClearDepth(depth Double) {
	C.callClearDepth((C.GLdouble)(depth))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterf
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterf.xml
func SamplerParameterf(sampler Uint, pname Enum, param Float) {
	C.callSamplerParameterf((C.GLuint)(sampler), (C.GLenum)(pname), (C.GLfloat)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBufferSubData
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBufferSubData.xml
func BufferSubData(target Enum, offset Intptr, size Sizeiptr, data Ptr) {
	C.callBufferSubData((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginQuery
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBeginQuery.xml
func BeginQuery(target Enum, id Uint) {
	C.callBeginQuery((C.GLenum)(target), (C.GLuint)(id))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramPipelineiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramPipelineiv.xml
func GetProgramPipelineiv(pipeline Uint, pname Enum, params *Int) {
	C.callGetProgramPipelineiv((C.GLuint)(pipeline), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramBinary
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramBinary.xml
func ProgramBinary(program Uint, binaryFormat Enum, binary Ptr, length Sizei) {
	C.callProgramBinary((C.GLuint)(program), (C.GLenum)(binaryFormat), (unsafe.Pointer)(binary), (C.GLsizei)(length))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4ubv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4ubv.xml
func VertexAttribI4ubv(index Uint, v *Ubyte) {
	C.callVertexAttribI4ubv((C.GLuint)(index), (*C.GLubyte)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribFormat
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribFormat.xml
func VertexAttribFormat(attribindex Uint, size Int, type_ Enum, normalized Boolean, relativeoffset Uint) {
	C.callVertexAttribFormat((C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLuint)(relativeoffset))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsTransformFeedback.xml
func IsTransformFeedback(id Uint) Boolean {
	return (Boolean)(C.callIsTransformFeedback((C.GLuint)(id)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetAttribLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetAttribLocation.xml
func GetAttribLocation(program Uint, name *Char) Int {
	return (Int)(C.callGetAttribLocation((C.GLuint)(program), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3iv.xml
func ProgramUniform3iv(program Uint, location Int, count Sizei, value *Int) {
	C.callProgramUniform3iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUnmapBuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUnmapBuffer.xml
func UnmapBuffer(target Enum) Boolean {
	return (Boolean)(C.callUnmapBuffer((C.GLenum)(target)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenTextures
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenTextures.xml
func GenTextures(n Sizei, textures *Uint) {
	C.callGenTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceLocationIndex
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceLocationIndex.xml
func GetProgramResourceLocationIndex(program Uint, programInterface Enum, name *Char) Int {
	return (Int)(C.callGetProgramResourceLocationIndex((C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(name)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderSource
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderSource.xml
func GetShaderSource(shader Uint, bufSize Sizei, length *Sizei, source *Char) {
	C.callGetShaderSource((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(source))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsFramebuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsFramebuffer.xml
func IsFramebuffer(framebuffer Uint) Boolean {
	return (Boolean)(C.callIsFramebuffer((C.GLuint)(framebuffer)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP2uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP2uiv.xml
func TexCoordP2uiv(type_ Enum, coords *Uint) {
	C.callTexCoordP2uiv((C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4fv.xml
func UniformMatrix4fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexImage1D.xml
func TexImage1D(target Enum, level Int, internalformat Int, width Sizei, border Int, format Enum, type_ Enum, pixels Ptr) {
	C.callTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEnablei
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEnablei.xml
func Enablei(target Enum, index Uint) {
	C.callEnablei((C.GLenum)(target), (C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2ui.xml
func ProgramUniform2ui(program Uint, location Int, v0 Uint, v1 Uint) {
	C.callProgramUniform2ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorP3uiv.xml
func ColorP3uiv(type_ Enum, color *Uint) {
	C.callColorP3uiv((C.GLenum)(type_), (*C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFramebufferParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFramebufferParameteriv.xml
func GetFramebufferParameteriv(target Enum, pname Enum, params *Int) {
	C.callGetFramebufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteBuffers
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteBuffers.xml
func DeleteBuffers(n Sizei, buffers *Uint) {
	C.callDeleteBuffers((C.GLsizei)(n), (*C.GLuint)(buffers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP2ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP2ui.xml
func MultiTexCoordP2ui(texture Enum, type_ Enum, coords Uint) {
	C.callMultiTexCoordP2ui((C.GLenum)(texture), (C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterfv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterfv.xml
func TexParameterfv(target Enum, pname Enum, params *Float) {
	C.callTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL2dv.xml
func VertexAttribL2dv(index Uint, v *Double) {
	C.callVertexAttribL2dv((C.GLuint)(index), (*C.GLdouble)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedbackStream
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedbackStream.xml
func DrawTransformFeedbackStream(mode Enum, id Uint, stream Uint) {
	C.callDrawTransformFeedbackStream((C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteProgramPipelines
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteProgramPipelines.xml
func DeleteProgramPipelines(n Sizei, pipelines *Uint) {
	C.callDeleteProgramPipelines((C.GLsizei)(n), (*C.GLuint)(pipelines))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2dv.xml
func ProgramUniform2dv(program Uint, location Int, count Sizei, value *Double) {
	C.callProgramUniform2dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2fv.xml
func ProgramUniform2fv(program Uint, location Int, count Sizei, value *Float) {
	C.callProgramUniform2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetMultisamplefv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetMultisamplefv.xml
func GetMultisamplefv(pname Enum, index Uint, val *Float) {
	C.callGetMultisamplefv((C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(val))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP3uiv.xml
func VertexP3uiv(type_ Enum, value *Uint) {
	C.callVertexP3uiv((C.GLenum)(type_), (*C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFuncSeparatei
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendFuncSeparatei.xml
func BlendFuncSeparatei(buf Uint, srcRGB Enum, dstRGB Enum, srcAlpha Enum, dstAlpha Enum) {
	C.callBlendFuncSeparatei((C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformName
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformName.xml
func GetActiveUniformName(program Uint, uniformIndex Uint, bufSize Sizei, length *Sizei, uniformName *Char) {
	C.callGetActiveUniformName((C.GLuint)(program), (C.GLuint)(uniformIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(uniformName))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP1ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP1ui.xml
func TexCoordP1ui(type_ Enum, coords Uint) {
	C.callTexCoordP1ui((C.GLenum)(type_), (C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateTexImage
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateTexImage.xml
func InvalidateTexImage(texture Uint, level Int) {
	C.callInvalidateTexImage((C.GLuint)(texture), (C.GLint)(level))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCheckFramebufferStatus
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCheckFramebufferStatus.xml
func CheckFramebufferStatus(target Enum) Enum {
	return (Enum)(C.callCheckFramebufferStatus((C.GLenum)(target)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP3uiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP3uiv.xml
func MultiTexCoordP3uiv(texture Enum, type_ Enum, coords *Uint) {
	C.callMultiTexCoordP3uiv((C.GLenum)(texture), (C.GLenum)(type_), (*C.GLuint)(coords))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindVertexArray
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindVertexArray.xml
func BindVertexArray(array Uint) {
	C.callBindVertexArray((C.GLuint)(array))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glColorP3ui.xml
func ColorP3ui(type_ Enum, color Uint) {
	C.callColorP3ui((C.GLenum)(type_), (C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform2i.xml
func Uniform2i(location Int, v0 Int, v1 Int) {
	C.callUniform2i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveAttrib
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveAttrib.xml
func GetActiveAttrib(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char) {
	C.callGetActiveAttrib((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginConditionalRender
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBeginConditionalRender.xml
func BeginConditionalRender(id Uint, mode Enum) {
	C.callBeginConditionalRender((C.GLuint)(id), (C.GLenum)(mode))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3d.xml
func ProgramUniform3d(program Uint, location Int, v0 Double, v1 Double, v2 Double) {
	C.callProgramUniform3d((C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glReadPixels
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glReadPixels.xml
func ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Ptr) {
	C.callReadPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(type_), (unsafe.Pointer)(pixels))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL1d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL1d.xml
func VertexAttribL1d(index Uint, x Double) {
	C.callVertexAttribL1d((C.GLuint)(index), (C.GLdouble)(x))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribPointer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribPointer.xml
func VertexAttribPointer(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, pointer Ptr) {
	C.callVertexAttribPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(type_), (C.GLboolean)(normalized), (C.GLsizei)(stride), (unsafe.Pointer)(pointer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribIuiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribIuiv.xml
func GetVertexAttribIuiv(index Uint, pname Enum, params *Uint) {
	C.callGetVertexAttribIuiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3ui.xml
func ProgramUniform3ui(program Uint, location Int, v0 Uint, v1 Uint, v2 Uint) {
	C.callProgramUniform3ui((C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2iv.xml
func ProgramUniform2iv(program Uint, location Int, count Sizei, value *Int) {
	C.callProgramUniform2iv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameteriv.xml
func SamplerParameteriv(sampler Uint, pname Enum, param *Int) {
	C.callSamplerParameteriv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(param))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4sv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4sv.xml
func VertexAttribI4sv(index Uint, v *Short) {
	C.callVertexAttribI4sv((C.GLuint)(index), (*C.GLshort)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x2fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x2fv.xml
func ProgramUniformMatrix4x2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float) {
	C.callProgramUniformMatrix4x2fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindFramebuffer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindFramebuffer.xml
func BindFramebuffer(target Enum, framebuffer Uint) {
	C.callBindFramebuffer((C.GLenum)(target), (C.GLuint)(framebuffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x3dv.xml
func ProgramUniformMatrix2x3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix2x3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x4fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x4fv.xml
func UniformMatrix2x4fv(location Int, count Sizei, transpose Boolean, value *Float) {
	C.callUniformMatrix2x4fv((C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexImage2D.xml
func CopyTexImage2D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, height Sizei, border Int) {
	C.callCopyTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexP3ui.xml
func VertexP3ui(type_ Enum, value Uint) {
	C.callVertexP3ui((C.GLenum)(type_), (C.GLuint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveSubroutineUniformiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveSubroutineUniformiv.xml
func GetActiveSubroutineUniformiv(program Uint, shadertype Enum, index Uint, pname Enum, values *Int) {
	C.callGetActiveSubroutineUniformiv((C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(values))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x4dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x4dv.xml
func ProgramUniformMatrix3x4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix3x4dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferfi
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferfi.xml
func ClearBufferfi(buffer Enum, drawbuffer Int, depth Float, stencil Int) {
	C.callClearBufferfi((C.GLenum)(buffer), (C.GLint)(drawbuffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1fv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1fv.xml
func ProgramUniform1fv(program Uint, location Int, count Sizei, value *Float) {
	C.callProgramUniform1fv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindTransformFeedback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindTransformFeedback.xml
func BindTransformFeedback(target Enum, id Uint) {
	C.callBindTransformFeedback((C.GLenum)(target), (C.GLuint)(id))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenerateMipmap
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenerateMipmap.xml
func GenerateMipmap(target Enum) {
	C.callGenerateMipmap((C.GLenum)(target))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexSubImage3D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage3D.xml
func CompressedTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, imageSize Sizei, data Ptr) {
	C.callCompressedTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameteriv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameteriv.xml
func GetSamplerParameteriv(sampler Uint, pname Enum, params *Int) {
	C.callGetSamplerParameteriv((C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenProgramPipelines
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGenProgramPipelines.xml
func GenProgramPipelines(n Sizei, pipelines *Uint) {
	C.callGenProgramPipelines((C.GLsizei)(n), (*C.GLuint)(pipelines))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMapBufferRange
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMapBufferRange.xml
func MapBufferRange(target Enum, offset Intptr, length Sizeiptr, access Bitfield) *Ptr {
	return (*Ptr)(C.callMapBufferRange((C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length), (C.GLbitfield)(access)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawArrays
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawArrays.xml
func MultiDrawArrays(mode Enum, first *Int, count *Sizei, drawcount Sizei) {
	C.callMultiDrawArrays((C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(drawcount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4f
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4f.xml
func ProgramUniform4f(program Uint, location Int, v0 Float, v1 Float, v2 Float, v3 Float) {
	C.callProgramUniform4f((C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteTransformFeedbacks
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteTransformFeedbacks.xml
func DeleteTransformFeedbacks(n Sizei, ids *Uint) {
	C.callDeleteTransformFeedbacks((C.GLsizei)(n), (*C.GLuint)(ids))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTextureView
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glTextureView.xml
func TextureView(texture Uint, target Enum, origtexture Uint, internalformat Enum, minlevel Uint, numlevels Uint, minlayer Uint, numlayers Uint) {
	C.callTextureView((C.GLuint)(texture), (C.GLenum)(target), (C.GLuint)(origtexture), (C.GLenum)(internalformat), (C.GLuint)(minlevel), (C.GLuint)(numlevels), (C.GLuint)(minlayer), (C.GLuint)(numlayers))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1d
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1d.xml
func Uniform1d(location Int, x Double) {
	C.callUniform1d((C.GLint)(location), (C.GLdouble)(x))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribIFormat
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribIFormat.xml
func VertexAttribIFormat(attribindex Uint, size Int, type_ Enum, relativeoffset Uint) {
	C.callVertexAttribIFormat((C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(type_), (C.GLuint)(relativeoffset))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4ui.xml
func Uniform4ui(location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint) {
	C.callUniform4ui((C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform3iv.xml
func Uniform3iv(location Int, count Sizei, value *Int) {
	C.callUniform3iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsQuery
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glIsQuery.xml
func IsQuery(id Uint) Boolean {
	return (Boolean)(C.callIsQuery((C.GLuint)(id)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetDoublei_v
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetDoublei_v.xml
func GetDoublei_v(target Enum, index Uint, data *Double) {
	C.callGetDoublei_v((C.GLenum)(target), (C.GLuint)(index), (*C.GLdouble)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewport
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glViewport.xml
func Viewport(x Int, y Int, width Sizei, height Sizei) {
	C.callViewport((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFunc
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBlendFunc.xml
func BlendFunc(sfactor Enum, dfactor Enum) {
	C.callBlendFunc((C.GLenum)(sfactor), (C.GLenum)(dfactor))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4i
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform4i.xml
func Uniform4i(location Int, v0 Int, v1 Int, v2 Int, v3 Int) {
	C.callUniform4i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteTextures
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteTextures.xml
func DeleteTextures(n Sizei, textures *Uint) {
	C.callDeleteTextures((C.GLsizei)(n), (*C.GLuint)(textures))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSecondaryColorP3ui
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColorP3ui.xml
func SecondaryColorP3ui(type_ Enum, color Uint) {
	C.callSecondaryColorP3ui((C.GLenum)(type_), (C.GLuint)(color))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTextureLayer
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTextureLayer.xml
func FramebufferTextureLayer(target Enum, attachment Enum, texture Uint, level Int, layer Int) {
	C.callFramebufferTextureLayer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTransformFeedbackVarying
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTransformFeedbackVarying.xml
func GetTransformFeedbackVarying(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Sizei, type_ *Enum, name *Char) {
	C.callGetTransformFeedbackVarying((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLsizei)(size), (*C.GLenum)(type_), (*C.GLchar)(name))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDebugMessageCallback
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDebugMessageCallback.xml
func DebugMessageCallback(callback Ptr, userParam Ptr) {
	C.callDebugMessageCallback((*[0]byte)(callback), (unsafe.Pointer)(userParam))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexSubImage2D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage2D.xml
func CompressedTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, imageSize Sizei, data Ptr) {
	C.callCompressedTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), (unsafe.Pointer)(data))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjectiv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectiv.xml
func GetQueryObjectiv(id Uint, pname Enum, params *Int) {
	C.callGetQueryObjectiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteQueries
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDeleteQueries.xml
func DeleteQueries(n Sizei, ids *Uint) {
	C.callDeleteQueries((C.GLsizei)(n), (*C.GLuint)(ids))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedbackStreamInstanced
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedbackStreamInstanced.xml
func DrawTransformFeedbackStreamInstanced(mode Enum, id Uint, stream Uint, instancecount Sizei) {
	C.callDrawTransformFeedbackStreamInstanced((C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream), (C.GLsizei)(instancecount))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2iv.xml
func VertexAttribI2iv(index Uint, v *Int) {
	C.callVertexAttribI2iv((C.GLuint)(index), (*C.GLint)(v))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexImage1D
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexImage1D.xml
func CopyTexImage1D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, border Int) {
	C.callCopyTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1iv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glUniform1iv.xml
func Uniform1iv(location Int, count Sizei, value *Int) {
	C.callUniform1iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilMaskSeparate
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glStencilMaskSeparate.xml
func StencilMaskSeparate(face Enum, mask Uint) {
	C.callStencilMaskSeparate((C.GLenum)(face), (C.GLuint)(mask))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFenceSync
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFenceSync.xml
func FenceSync(condition Enum, flags Bitfield) Sync {
	return (Sync)(C.callFenceSync((C.GLenum)(condition), (C.GLbitfield)(flags)))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glReleaseShaderCompiler
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glReleaseShaderCompiler.xml
func ReleaseShaderCompiler() {
	C.callReleaseShaderCompiler()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstancedBaseInstance
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstancedBaseInstance.xml
func DrawElementsInstancedBaseInstance(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, baseinstance Uint) {
	C.callDrawElementsInstancedBaseInstance((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(type_), (unsafe.Pointer)(indices), (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3dv
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3dv.xml
func ProgramUniformMatrix3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double) {
	C.callProgramUniformMatrix3dv((C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(transpose), (*C.GLdouble)(value))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFlush
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glFlush.xml
func Flush() {
	C.callFlush()
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindBufferBase
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindBufferBase.xml
func BindBufferBase(target Enum, index Uint, buffer Uint) {
	C.callBindBufferBase((C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEnableVertexAttribArray
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glEnableVertexAttribArray.xml
func EnableVertexAttribArray(index Uint) {
	C.callEnableVertexAttribArray((C.GLuint)(index))
}

//	GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceLocation
//	SDK doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceLocation.xml
func GetProgramResourceLocation(program Uint, programInterface Enum, name *Char) Int {
	return (Int)(C.callGetProgramResourceLocation((C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(name)))
}
