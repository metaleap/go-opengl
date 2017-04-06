# gl
--
    import "github.com/metaleap/go-opengl/core"

A cross-platform Go binding for OpenGL core profile 3.3 and/or higher with
extensions EXT_texture_filter_anisotropic, EXT_texture_compression_s3tc,
EXT_texture_sRGB. Auto-generated using
github.com/metaleap/go-opengl/cmd/gen-opengl-bindings

To initialize this OpenGL binding, after having created your GL context (with
SDL, GLFW or in some other way) call glcore.Util.Init().

## Usage

```go
const (
	FRAGMENT_SHADER                                            = 0x8B30
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	BUFFER_MAP_LENGTH                                          = 0x9120
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	PACK_ROW_LENGTH                                            = 0x0D02
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	NONE                                                       = 0
	CLAMP_TO_EDGE                                              = 0x812F
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	TEXTURE8                                                   = 0x84C8
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	POLYGON_MODE                                               = 0x0B40
	SRGB8_ALPHA8_EXT                                           = 0x8C43
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	READ_PIXELS_TYPE                                           = 0x828E
	NUM_SHADING_LANGUAGE_VERSIONS                              = 0x82E9
	ACTIVE_UNIFORMS                                            = 0x8B86
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	SCISSOR_BOX                                                = 0x0C10
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	DRAW_BUFFER6                                               = 0x882B
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	POINTS                                                     = 0x0000
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	RGBA32F                                                    = 0x8814
	RENDERBUFFER_WIDTH                                         = 0x8D42
	MAX_IMAGE_SAMPLES                                          = 0x906D
	RG8                                                        = 0x822B
	COLOR_ATTACHMENT13                                         = 0x8CED
	SAMPLER_2D                                                 = 0x8B5E
	RGBA8UI                                                    = 0x8D7C
	OUT_OF_MEMORY                                              = 0x0505
	BUFFER                                                     = 0x82E0
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	RG8I                                                       = 0x8237
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	STREAM_DRAW                                                = 0x88E0
	BLEND_EQUATION_RGB                                         = 0x8009
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	TEXTURE_MAG_FILTER                                         = 0x2800
	RGBA12                                                     = 0x805A
	HALF_FLOAT                                                 = 0x140B
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	DECR_WRAP                                                  = 0x8508
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	TEXTURE_1D_ARRAY                                           = 0x8C18
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	MAX_VARYING_VECTORS                                        = 0x8DFC
	MINOR_VERSION                                              = 0x821C
	SAMPLE_SHADING                                             = 0x8C36
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	DECR                                                       = 0x1E03
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	COLOR_ATTACHMENT10                                         = 0x8CEA
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	DOUBLE_VEC3                                                = 0x8FFD
	TEXTURE23                                                  = 0x84D7
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	SCISSOR_TEST                                               = 0x0C11
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	SRGB_ALPHA_EXT                                             = 0x8C42
	DRAW_BUFFER10                                              = 0x882F
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	NEVER                                                      = 0x0200
	STENCIL_VALUE_MASK                                         = 0x0B93
	IMAGE_BINDING_FORMAT                                       = 0x906E
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	CULL_FACE_MODE                                             = 0x0B45
	DOUBLE                                                     = 0x140A
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	UNSIGNED_SHORT                                             = 0x1403
	PIXEL_PACK_BUFFER                                          = 0x88EB
	FRAMEBUFFER_BLEND                                          = 0x828B
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	COLOR_ATTACHMENT0                                          = 0x8CE0
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = TRANSFORM_FEEDBACK_ACTIVE
	DEBUG_TYPE_MARKER                                          = 0x8268
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	SRGB_EXT                                                   = 0x8C40
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	UNPACK_SKIP_IMAGES                                         = 0x806D
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	NOR                                                        = 0x1508
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	COLOR_WRITEMASK                                            = 0x0C23
	RG32UI                                                     = 0x823C
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	IMAGE_BINDING_NAME                                         = 0x8F3A
	ZERO                                                       = 0
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	STENCIL_BUFFER_BIT                                         = 0x00000400
	DEPTH32F_STENCIL8                                          = 0x8CAD
	PACK_ALIGNMENT                                             = 0x0D05
	DEPTH_BUFFER_BIT                                           = 0x00000100
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	COMPUTE_TEXTURE                                            = 0x82A0
	XOR                                                        = 0x1506
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	SRGB_ALPHA                                                 = 0x8C42
	DRAW_BUFFER0                                               = 0x8825
	RGB32I                                                     = 0x8D83
	DRAW_FRAMEBUFFER_BINDING                                   = FRAMEBUFFER_BINDING
	MEDIUM_INT                                                 = 0x8DF4
	POINT                                                      = 0x1B00
	R3_G3_B2                                                   = 0x2A10
	IMAGE_1D                                                   = 0x904C
	STATIC_COPY                                                = 0x88E6
	TEXTURE_BORDER_COLOR                                       = 0x1004
	VIEWPORT                                                   = 0x0BA2
	ARRAY_SIZE                                                 = 0x92FB
	COMPUTE_SUBROUTINE                                         = 0x92ED
	TESS_GEN_POINT_MODE                                        = 0x8E79
	VERTEX_SHADER                                              = 0x8B31
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	R8I                                                        = 0x8231
	EXTENSIONS                                                 = 0x1F03
	WAIT_FAILED                                                = 0x911D
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	R11F_G11F_B10F                                             = 0x8C3A
	ACTIVE_PROGRAM                                             = 0x8259
	PROGRAM_OUTPUT                                             = 0x92E4
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	LINE_SMOOTH                                                = 0x0B20
	DEPTH_TEST                                                 = 0x0B71
	CLIP_DISTANCE3                                             = 0x3003
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	RGBA8I                                                     = 0x8D8E
	COLOR_ENCODING                                             = 0x8296
	DEPTH24_STENCIL8                                           = 0x88F0
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	TEXTURE_MAX_LEVEL                                          = 0x813D
	R16UI                                                      = 0x8234
	FRONT                                                      = 0x0404
	BUFFER_BINDING                                             = 0x9302
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	UNSIGNALED                                                 = 0x9118
	COLOR_RENDERABLE                                           = 0x8286
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	INVERT                                                     = 0x150A
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	R16I                                                       = 0x8233
	LINE_WIDTH                                                 = 0x0B21
	NAME_LENGTH                                                = 0x92F9
	INT_SAMPLER_3D                                             = 0x8DCB
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	TRANSFORM_FEEDBACK                                         = 0x8E22
	VIEW_CLASS_16_BITS                                         = 0x82CA
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	FALSE                                                      = 0
	PROXY_TEXTURE_1D                                           = 0x8063
	DEPTH_RENDERABLE                                           = 0x8287
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	CURRENT_PROGRAM                                            = 0x8B8D
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	FILL                                                       = 0x1B02
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	VERTEX_SUBROUTINE                                          = 0x92E8
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	IMAGE_2D                                                   = 0x904D
	TEXTURE_MIN_FILTER                                         = 0x2801
	READ_ONLY                                                  = 0x88B8
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	VIEW_CLASS_64_BITS                                         = 0x82C6
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	POLYGON_OFFSET_POINT                                       = 0x2A01
	SAMPLER_2D_SHADOW                                          = 0x8B62
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	DISPLAY_LIST                                               = 0x82E7
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	POLYGON_OFFSET_LINE                                        = 0x2A02
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	FRONT_RIGHT                                                = 0x0401
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	TEXTURE_BUFFER                                             = 0x8C2A
	COLOR_ATTACHMENT15                                         = 0x8CEF
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	NEAREST                                                    = 0x2600
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	SLUMINANCE_EXT                                             = 0x8C46
	SYNC_CONDITION                                             = 0x9113
	RED_SNORM                                                  = 0x8F90
	TIMEOUT_EXPIRED                                            = 0x911B
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	UNSIGNED_NORMALIZED                                        = 0x8C17
	FLOAT_MAT4                                                 = 0x8B5C
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	BUFFER_MAPPED                                              = 0x88BC
	READ_BUFFER                                                = 0x0C02
	CLEAR_BUFFER                                               = 0x82B4
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	GREATER                                                    = 0x0204
	ONE                                                        = 1
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	IMAGE_2D_RECT                                              = 0x904F
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	QUERY_NO_WAIT                                              = 0x8E14
	MAX_PATCH_VERTICES                                         = 0x8E7D
	LEQUAL                                                     = 0x0203
	SHADER                                                     = 0x82E1
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	MAX_LAYERS                                                 = 0x8281
	ACTIVE_TEXTURE                                             = 0x84E0
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	RG16_SNORM                                                 = 0x8F99
	COMPUTE_LOCAL_WORK_SIZE                                    = 0x8267
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	LOCATION                                                   = 0x930E
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	GREEN                                                      = 0x1904
	TEXTURE9                                                   = 0x84C9
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	TEXTURE_WIDTH                                              = 0x1000
	TEXTURE_MIN_LOD                                            = 0x813A
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	TIMESTAMP                                                  = 0x8E28
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	LINES_ADJACENCY                                            = 0x000A
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	INVALID_INDEX                                              = 0xFFFFFFFF
	DRAW_BUFFER8                                               = 0x882D
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	VIEW_CLASS_8_BITS                                          = 0x82CB
	DEPTH_WRITEMASK                                            = 0x0B72
	RGBA_SNORM                                                 = 0x8F93
	TEXTURE2                                                   = 0x84C2
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	OR_REVERSE                                                 = 0x150B
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	DEBUG_TYPE_OTHER                                           = 0x8251
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	PROVOKING_VERTEX                                           = 0x8E4F
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	NUM_SAMPLE_COUNTS                                          = 0x9380
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	STATIC_READ                                                = 0x88E5
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	PROGRAM_PIPELINE                                           = 0x82E4
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	COMPRESSED_RG11_EAC                                        = 0x9272
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	INT                                                        = 0x1404
	INT_VEC3                                                   = 0x8B54
	BYTE                                                       = 0x1400
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	STENCIL_COMPONENTS                                         = 0x8285
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	UNIFORM                                                    = 0x92E1
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	BLUE                                                       = 0x1905
	VERTEX_ATTRIB_ARRAY_LONG                                   = 0x874E
	RGB                                                        = 0x1907
	QUERY_WAIT                                                 = 0x8E13
	FLOAT_VEC3                                                 = 0x8B51
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	TESS_CONTROL_SHADER                                        = 0x8E88
	DEPTH_COMPONENT16                                          = 0x81A5
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	VALIDATE_STATUS                                            = 0x8B83
	TYPE                                                       = 0x92FA
	COMPRESSED_SRGB_ALPHA_EXT                                  = 0x8C49
	LOGIC_OP_MODE                                              = 0x0BF0
	MAX_COMPUTE_LOCAL_INVOCATIONS                              = 0x90EB
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	SAMPLE_MASK_VALUE                                          = 0x8E52
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	TEXTURE6                                                   = 0x84C6
	MULTISAMPLE                                                = 0x809D
	RG8_SNORM                                                  = 0x8F95
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	FILTER                                                     = 0x829A
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	INT_SAMPLER_1D                                             = 0x8DC9
	VIEW_CLASS_24_BITS                                         = 0x82C9
	DEPTH_COMPONENT32F                                         = 0x8CAC
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	COMPRESSED_SLUMINANCE_ALPHA_EXT                            = 0x8C4B
	FRAGMENT_TEXTURE                                           = 0x829F
	RGB10                                                      = 0x8052
	FLOAT_VEC2                                                 = 0x8B50
	COMPRESSED_SRGB_EXT                                        = 0x8C48
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	PROXY_TEXTURE_3D                                           = 0x8070
	SIGNALED                                                   = 0x9119
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	QUERY                                                      = 0x82E3
	TEXTURE_MAX_LOD                                            = 0x813B
	COPY_READ_BUFFER                                           = COPY_READ_BUFFER_BINDING
	PROGRAM                                                    = 0x82E2
	TEXTURE18                                                  = 0x84D2
	DOUBLE_MAT3                                                = 0x8F47
	DITHER                                                     = 0x0BD0
	COLOR_ATTACHMENT3                                          = 0x8CE3
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	TEXTURE3                                                   = 0x84C3
	TEXTURE_BLUE_SIZE                                          = 0x805E
	TEXTURE25                                                  = 0x84D9
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	TEXTURE_BINDING_1D                                         = 0x8068
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	BUFFER_MAP_OFFSET                                          = 0x9121
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	DOUBLE_VEC4                                                = 0x8FFE
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	READ_PIXELS_FORMAT                                         = 0x828D
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	SYNC_STATUS                                                = 0x9114
	OBJECT_TYPE                                                = 0x9112
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	STENCIL_FAIL                                               = 0x0B94
	SAMPLER_BINDING                                            = 0x8919
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	BACK_LEFT                                                  = 0x0402
	TRIANGLE_STRIP                                             = 0x0005
	VENDOR                                                     = 0x1F00
	TEXTURE19                                                  = 0x84D3
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	DYNAMIC_COPY                                               = 0x88EA
	SAMPLE_MASK                                                = 0x8E51
	SRC1_COLOR                                                 = 0x88F9
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	RGBA16I                                                    = 0x8D88
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	COLOR_LOGIC_OP                                             = 0x0BF2
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	DYNAMIC_READ                                               = 0x88E9
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	UNIFORM_TYPE                                               = 0x8A37
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	SHADER_IMAGE_LOAD                                          = 0x82A4
	FLOAT_MAT2                                                 = 0x8B5A
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	PROGRAM_POINT_SIZE                                         = 0x8642
	COMPRESSED_RED                                             = 0x8225
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	BOOL_VEC2                                                  = 0x8B57
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	DEBUG_SEVERITY_LOW                                         = 0x9148
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	RG16F                                                      = 0x822F
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	AND_INVERTED                                               = 0x1504
	INT_IMAGE_2D                                               = 0x9058
	RED_INTEGER                                                = 0x8D94
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	SLUMINANCE8_EXT                                            = 0x8C47
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	RGB16UI                                                    = 0x8D77
	TEXTURE_WRAP_T                                             = 0x2803
	PRIMITIVES_GENERATED                                       = 0x8C87
	R32I                                                       = 0x8235
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	SHADER_STORAGE_BARRIER_BIT                                 = 0x2000
	KEEP                                                       = 0x1E00
	LINES                                                      = 0x0001
	BLEND                                                      = 0x0BE2
	FASTEST                                                    = 0x1101
	BOOL_VEC3                                                  = 0x8B58
	FRONT_AND_BACK                                             = 0x0408
	INT_2_10_10_10_REV                                         = 0x8D9F
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	STENCIL_INDEX8                                             = 0x8D48
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	ARRAY_STRIDE                                               = 0x92FE
	DRAW_BUFFER14                                              = 0x8833
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	DRAW_BUFFER                                                = 0x0C01
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	TEXTURE_BINDING_2D                                         = 0x8069
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	TEXTURE_3D                                                 = 0x806F
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	RGB8UI                                                     = 0x8D7D
	MIN_SAMPLE_SHADING_VALUE                                   = 0x8C37
	BLEND_DST                                                  = 0x0BE0
	TRUE                                                       = 1
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	ARRAY_BUFFER_BINDING                                       = 0x8894
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	ACTIVE_VARIABLES                                           = 0x9305
	RENDERBUFFER                                               = 0x8D41
	REPLACE                                                    = 0x1E01
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	BOOL_VEC4                                                  = 0x8B59
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	BLOCK_INDEX                                                = 0x92FD
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT                        = 0x8C4D
	PATCH_VERTICES                                             = 0x8E72
	RG_INTEGER                                                 = 0x8228
	AND_REVERSE                                                = 0x1502
	MAX_VARYING_FLOATS                                         = 0x8B4B
	TEXTURE10                                                  = 0x84CA
	TIME_ELAPSED                                               = 0x88BF
	DEPTH_COMPONENT32                                          = 0x81A7
	FRACTIONAL_EVEN                                            = 0x8E7C
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	RGB_INTEGER                                                = 0x8D98
	IMAGE_CUBE                                                 = 0x9050
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	READ_WRITE                                                 = 0x88BA
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	FRACTIONAL_ODD                                             = 0x8E7B
	SAMPLER_1D                                                 = 0x8B5D
	INT_IMAGE_3D                                               = 0x9059
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	BACK_RIGHT                                                 = 0x0403
	CONTEXT_FLAGS                                              = 0x821E
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	UNIFORM_BUFFER_START                                       = 0x8A29
	DEBUG_OUTPUT                                               = 0x92E0
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	STENCIL_BACK_REF                                           = 0x8CA3
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	FLOAT                                                      = 0x1406
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	ALREADY_SIGNALED                                           = 0x911A
	TEXTURE14                                                  = 0x84CE
	UNIFORM_BLOCK                                              = 0x92E2
	VIEW_CLASS_128_BITS                                        = 0x82C4
	CULL_FACE                                                  = 0x0B44
	DYNAMIC_DRAW                                               = 0x88E8
	STENCIL_INDEX1                                             = 0x8D46
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	CLIP_DISTANCE4                                             = 0x3004
	BGRA                                                       = 0x80E1
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x0001
	COMPRESSED_RG                                              = 0x8226
	UPPER_LEFT                                                 = 0x8CA2
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	DRAW_BUFFER7                                               = 0x882C
	INT_IMAGE_2D_RECT                                          = 0x905A
	COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT                        = 0x8C4F
	STREAM_COPY                                                = 0x88E2
	TESS_GEN_MODE                                              = 0x8E76
	INT_IMAGE_1D                                               = 0x9057
	RGBA16F                                                    = 0x881A
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	DRAW_BUFFER11                                              = 0x8830
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	REPEAT                                                     = 0x2901
	COMMAND_BARRIER_BIT                                        = 0x00000040
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	ISOLINES                                                   = 0x8E7A
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	CLAMP_TO_BORDER                                            = 0x812D
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS                      = 0x8F9F
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	RED                                                        = 0x1903
	RGB16I                                                     = 0x8D89
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	BACK                                                       = 0x0405
	NUM_EXTENSIONS                                             = 0x821D
	CONTEXT_PROFILE_MASK                                       = 0x9126
	SRC1_ALPHA                                                 = 0x8589
	LINE_LOOP                                                  = 0x0002
	RG_SNORM                                                   = 0x8F91
	RGB12                                                      = 0x8053
	TEXTURE31                                                  = 0x84DF
	R32F                                                       = 0x822E
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	UNIFORM_BUFFER                                             = 0x8A11
	BUFFER_MAP_POINTER                                         = 0x88BD
	FULL_SUPPORT                                               = 0x82B7
	COLOR_BUFFER_BIT                                           = 0x00004000
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	SYNC_FLAGS                                                 = 0x9115
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	DEPTH                                                      = 0x1801
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	NICEST                                                     = 0x1102
	MEDIUM_FLOAT                                               = 0x8DF1
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	MAX_WIDTH                                                  = 0x827E
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	VERTEX_TEXTURE                                             = 0x829B
	INT_SAMPLER_2D                                             = 0x8DCA
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	FIXED                                                      = 0x140C
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	COLOR_ATTACHMENT1                                          = 0x8CE1
	TEXTURE_COMPARE_MODE                                       = 0x884C
	PACK_IMAGE_HEIGHT                                          = 0x806C
	COPY                                                       = 0x1503
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	RASTERIZER_DISCARD                                         = 0x8C89
	VERSION                                                    = 0x1F02
	STENCIL                                                    = 0x1802
	BLEND_DST_RGB                                              = 0x80C8
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	MATRIX_STRIDE                                              = 0x92FF
	PACK_SKIP_IMAGES                                           = 0x806B
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	TEXTURE1                                                   = 0x84C1
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	INVALID_ENUM                                               = 0x0500
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	FLOAT_MAT4x2                                               = 0x8B69
	DRAW_BUFFER3                                               = 0x8828
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	R16                                                        = 0x822A
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	DEBUG_SOURCE_OTHER                                         = 0x824B
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	INVALID_VALUE                                              = 0x0501
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	SAMPLER_2D_RECT                                            = 0x8B63
	READ_PIXELS                                                = 0x828C
	PRIMITIVE_RESTART                                          = 0x8F9D
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	MAX_VERTEX_STREAMS                                         = 0x8E71
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	STEREO                                                     = 0x0C33
	SAMPLER_1D_SHADOW                                          = 0x8B61
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	MAJOR_VERSION                                              = 0x821B
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	DOUBLEBUFFER                                               = 0x0C32
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	RG16UI                                                     = 0x823A
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	RGB5_A1                                                    = 0x8057
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	SAMPLER_3D                                                 = 0x8B5F
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	TEXTURE_CUBE_MAP                                           = 0x8513
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	DOUBLE_MAT4x2                                              = 0x8F4D
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	IMAGE_3D                                                   = 0x904E
	SAMPLE_BUFFERS                                             = 0x80A8
	SEPARATE_ATTRIBS                                           = 0x8C8D
	BGR_INTEGER                                                = 0x8D9A
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	BUFFER_VARIABLE                                            = 0x92E5
	STACK_UNDERFLOW                                            = 0x0504
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	DEBUG_TYPE_ERROR                                           = 0x824C
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	TESS_EVALUATION_SHADER                                     = 0x8E87
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	BGRA_INTEGER                                               = 0x8D9B
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	COMPILE_STATUS                                             = 0x8B81
	INT_VEC4                                                   = 0x8B55
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	LOW_INT                                                    = 0x8DF3
	DRAW_BUFFER15                                              = 0x8834
	R16_SNORM                                                  = 0x8F98
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	TEXTURE_WRAP_S                                             = 0x2802
	MAX_SUBROUTINES                                            = 0x8DE7
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	RGB5                                                       = 0x8050
	TESS_GEN_SPACING                                           = 0x8E77
	IMAGE_BUFFER                                               = 0x9051
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	RG16                                                       = 0x822C
	RGBA32I                                                    = 0x8D82
	TEXTURE29                                                  = 0x84DD
	COMPRESSED_RGB                                             = 0x84ED
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	RGBA4                                                      = 0x8056
	COMPRESSED_R11_EAC                                         = 0x9270
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	RGB16_SNORM                                                = 0x8F9A
	DEPTH_COMPONENTS                                           = 0x8284
	SUBPIXEL_BITS                                              = 0x0D50
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	RGB8I                                                      = 0x8D8F
	QUERY_COUNTER_BITS                                         = 0x8864
	LEFT                                                       = 0x0406
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	TEXTURE_2D                                                 = 0x0DE1
	COPY_WRITE_BUFFER                                          = COPY_WRITE_BUFFER_BINDING
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	UNIFORM_OFFSET                                             = 0x8A3B
	NAND                                                       = 0x150E
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	IMAGE_2D_ARRAY                                             = 0x9053
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	GEQUAL                                                     = 0x0206
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	FRONT_FACE                                                 = 0x0B46
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	SHADER_IMAGE_STORE                                         = 0x82A5
	BGR                                                        = 0x80E0
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	TEXTURE_BASE_LEVEL                                         = 0x813C
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	LINE_SMOOTH_HINT                                           = 0x0C52
	CW                                                         = 0x0900
	RG8UI                                                      = 0x8238
	DELETE_STATUS                                              = 0x8B80
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	SRGB_WRITE                                                 = 0x8298
	DEPTH_ATTACHMENT                                           = 0x8D00
	HIGH_FLOAT                                                 = 0x8DF2
	ONE_MINUS_DST_COLOR                                        = 0x0307
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	TEXTURE5                                                   = 0x84C5
	BLEND_SRC_ALPHA                                            = 0x80CB
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	STENCIL_BACK_FUNC                                          = 0x8800
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	STATIC_DRAW                                                = 0x88E4
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	LINK_STATUS                                                = 0x8B82
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	AND                                                        = 0x1501
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	FRONT_LEFT                                                 = 0x0400
	STACK_OVERFLOW                                             = 0x0503
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	MAX_INTEGER_SAMPLES                                        = 0x9110
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	STENCIL_ATTACHMENT                                         = 0x8D20
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	DRAW_BUFFER13                                              = 0x8832
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	RGB4                                                       = 0x804F
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	LINE_STRIP                                                 = 0x0003
	TEXTURE_LOD_BIAS                                           = 0x8501
	POLYGON_OFFSET_FILL                                        = 0x8037
	RGB8                                                       = 0x8051
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	TEXTURE4                                                   = 0x84C4
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	SYNC_FENCE                                                 = 0x9116
	CCW                                                        = 0x0901
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	CLIP_DISTANCE7                                             = 0x3007
	INT_IMAGE_BUFFER                                           = 0x905C
	R16F                                                       = 0x822D
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	DST_COLOR                                                  = 0x0306
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	STENCIL_INDEX16                                            = 0x8D49
	TEXTURE_VIEW                                               = 0x82B5
	DST_ALPHA                                                  = 0x0304
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	STENCIL_FUNC                                               = 0x0B92
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	CLIP_DISTANCE1                                             = 0x3001
	BLEND_SRC_RGB                                              = 0x80C9
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	INT_SAMPLER_CUBE                                           = 0x8DCC
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	INCR                                                       = 0x1E02
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	MAX_CLIP_DISTANCES                                         = 0x0D32
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	SRC_ALPHA                                                  = 0x0302
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	STENCIL_BACK_FAIL                                          = 0x8801
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	TEXTURE_1D                                                 = 0x0DE0
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	COLOR_COMPONENTS                                           = 0x8283
	CONDITION_SATISFIED                                        = 0x911C
	UNSIGNED_BYTE                                              = 0x1401
	FLOAT_MAT3                                                 = 0x8B5B
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	SAMPLER_CUBE                                               = 0x8B60
	RGBA                                                       = 0x1908
	COLOR_ATTACHMENT2                                          = 0x8CE2
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	SRGB8                                                      = 0x8C41
	TEXTURE_GREEN_SIZE                                         = 0x805D
	COPY_INVERTED                                              = 0x150C
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	SAMPLES                                                    = 0x80A9
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	COLOR_ATTACHMENT9                                          = 0x8CE9
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	DRAW_BUFFER9                                               = 0x882E
	BUFFER_USAGE                                               = 0x8765
	R8UI                                                       = 0x8232
	UNPACK_ALIGNMENT                                           = 0x0CF5
	TEXTURE                                                    = 0x1702
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	PROGRAM_SEPARABLE                                          = 0x8258
	MAX_IMAGE_UNITS                                            = 0x8F38
	DEBUG_SOURCE_API                                           = 0x8246
	SET                                                        = 0x150F
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	DOUBLE_MAT4                                                = 0x8F48
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	INFO_LOG_LENGTH                                            = 0x8B84
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	BLEND_SRC                                                  = 0x0BE1
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	RGBA8_SNORM                                                = 0x8F97
	DEPTH_COMPONENT                                            = 0x1902
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	MAP_READ_BIT                                               = 0x0001
	STENCIL_TEST                                               = 0x0B90
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	UNIFORM_SIZE                                               = 0x8A38
	UNDEFINED_VERTEX                                           = 0x8260
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	DOUBLE_MAT2x4                                              = 0x8F4A
	FIXED_ONLY                                                 = 0x891D
	RGB565                                                     = 0x8D62
	RGB16F                                                     = 0x881B
	VERTEX_SHADER_BIT                                          = 0x00000001
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	RGBA2                                                      = 0x8055
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	TEXTURE_GATHER                                             = 0x82A2
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	SLUMINANCE_ALPHA_EXT                                       = 0x8C44
	DOUBLE_MAT2x3                                              = 0x8F49
	FLOAT_MAT3x2                                               = 0x8B67
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	COLOR_ATTACHMENT11                                         = 0x8CEB
	TEXTURE_RED_TYPE                                           = 0x8C10
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	COLOR                                                      = 0x1800
	COLOR_ATTACHMENT8                                          = 0x8CE8
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	TEXTURE_COMPRESSED                                         = 0x86A1
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	RGB32UI                                                    = 0x8D71
	DOUBLE_MAT2                                                = 0x8F46
	LINEAR                                                     = 0x2601
	GEOMETRY_TEXTURE                                           = 0x829E
	RENDERER                                                   = 0x1F01
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	R8                                                         = 0x8229
	TEXTURE_DEPTH                                              = 0x8071
	TEXTURE_SHADOW                                             = 0x82A1
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	TEXTURE28                                                  = 0x84DC
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	TIMEOUT_IGNORED                                            = 0xFFFFFFFFFFFFFFFF
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	QUADS                                                      = 0x0007
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	WRITE_ONLY                                                 = 0x88B9
	TEXTURE_BINDING_3D                                         = 0x806A
	TEXTURE_HEIGHT                                             = 0x1001
	PROGRAM_INPUT                                              = 0x92E3
	DRAW_BUFFER4                                               = 0x8829
	SHADER_COMPILER                                            = 0x8DFA
	MAX_SAMPLES                                                = 0x8D57
	COMPRESSED_SRGB                                            = 0x8C48
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	DRAW_BUFFER5                                               = 0x882A
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	COLOR_CLEAR_VALUE                                          = 0x0C22
	VIEW_CLASS_48_BITS                                         = 0x82C7
	DRAW_BUFFER12                                              = 0x8831
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	BOOL                                                       = 0x8B56
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	IS_ROW_MAJOR                                               = 0x9300
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	ARRAY_BUFFER                                               = 0x8892
	DONT_CARE                                                  = 0x1100
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	ALPHA                                                      = 0x1906
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	MAX_LABEL_LENGTH                                           = 0x82E8
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	LINE                                                       = 0x1B01
	PATCHES                                                    = 0x000E
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	RGB_SNORM                                                  = 0x8F92
	TEXTURE21                                                  = 0x84D5
	BUFFER_SIZE                                                = 0x8764
	DEPTH_FUNC                                                 = 0x0B74
	TESS_CONTROL_TEXTURE                                       = 0x829C
	STENCIL_WRITEMASK                                          = 0x0B98
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	MIRRORED_REPEAT                                            = 0x8370
	BLEND_DST_ALPHA                                            = 0x80CA
	FLOAT_MAT4x3                                               = 0x8B6A
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	TEXTURE7                                                   = 0x84C7
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	TEXTURE15                                                  = 0x84CF
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	SAMPLES_PASSED                                             = 0x8914
	UNSIGNED_INT_24_8                                          = 0x84FA
	SIGNED_NORMALIZED                                          = 0x8F9C
	RGB32F                                                     = 0x8815
	TEXTURE_WRAP_R                                             = 0x8072
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	RGBA16UI                                                   = 0x8D76
	TEXTURE_SAMPLES                                            = 0x9106
	TEXTURE0                                                   = 0x84C0
	VIEW_CLASS_96_BITS                                         = 0x82C5
	BUFFER_DATA_SIZE                                           = 0x9303
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	R32UI                                                      = 0x8236
	DEPTH_STENCIL                                              = 0x84F9
	IMAGE_1D_ARRAY                                             = 0x9052
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	TEXTURE12                                                  = 0x84CC
	COLOR_ATTACHMENT6                                          = 0x8CE6
	COLOR_ATTACHMENT14                                         = 0x8CEE
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	PACK_SKIP_PIXELS                                           = 0x0D04
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	CLIP_DISTANCE5                                             = 0x3005
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	RIGHT                                                      = 0x0407
	RG32I                                                      = 0x823B
	DOUBLE_MAT3x4                                              = 0x8F4C
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	RGBA16                                                     = 0x805B
	COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT                        = 0x8C4E
	INCR_WRAP                                                  = 0x8507
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	COMPRESSED_RGBA                                            = 0x84EE
	SAMPLER                                                    = 0x82E6
	INT_IMAGE_CUBE                                             = 0x905B
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	DOUBLE_VEC2                                                = 0x8FFC
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	COLOR_ATTACHMENT4                                          = 0x8CE4
	NO_ERROR                                                   = 0
	TEXTURE11                                                  = 0x84CB
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	MAX_HEIGHT                                                 = 0x827F
	PROXY_TEXTURE_2D                                           = 0x8064
	RGB10_A2UI                                                 = 0x906F
	HIGH_INT                                                   = 0x8DF5
	COLOR_ATTACHMENT5                                          = 0x8CE5
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	OR                                                         = 0x1507
	FLOAT_MAT2x3                                               = 0x8B65
	BLEND_EQUATION_ALPHA                                       = 0x883D
	TEXTURE27                                                  = 0x84DB
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	RGB9_E5                                                    = 0x8C3D
	RG32F                                                      = 0x8230
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	LOCATION_INDEX                                             = 0x930F
	SAMPLE_COVERAGE                                            = 0x80A0
	RGBA32UI                                                   = 0x8D70
	SRGB8_EXT                                                  = 0x8C41
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	MAX_DEPTH                                                  = 0x8280
	ALWAYS                                                     = 0x0207
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	MAX_VIEWPORTS                                              = 0x825B
	CAVEAT_SUPPORT                                             = 0x82B8
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	RGB10_A2                                                   = 0x8059
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	CLAMP_READ_COLOR                                           = 0x891C
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	NOOP                                                       = 0x1505
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	COMPRESSED_SRGB_S3TC_DXT1_EXT                              = 0x8C4C
	SAMPLE_POSITION                                            = 0x8E50
	MAP_WRITE_BIT                                              = 0x0002
	RGB8_SNORM                                                 = 0x8F96
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	FLOAT_MAT3x4                                               = 0x8B68
	SHORT                                                      = 0x1402
	READ_FRAMEBUFFER                                           = 0x8CA8
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	ACTIVE_RESOURCES                                           = 0x92F5
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	TEXTURE_RECTANGLE                                          = 0x84F5
	COLOR_ATTACHMENT12                                         = 0x8CEC
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	BLUE_INTEGER                                               = 0x8D96
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	SHADER_TYPE                                                = 0x8B4F
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	TRIANGLES_ADJACENCY                                        = 0x000C
	STENCIL_INDEX4                                             = 0x8D47
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	RG16I                                                      = 0x8239
	CLIP_DISTANCE2                                             = 0x3002
	QUERY_RESULT                                               = 0x8866
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	MIPMAP                                                     = 0x8293
	RENDERBUFFER_BINDING                                       = 0x8CA7
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	DOUBLE_MAT4x3                                              = 0x8F4E
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	LOWER_LEFT                                                 = 0x8CA1
	CLEAR                                                      = 0x1500
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	MAX_DRAW_BUFFERS                                           = 0x8824
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	INT_VEC2                                                   = 0x8B53
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	SRC_ALPHA_SATURATE                                         = 0x0308
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	RGB16                                                      = 0x8054
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	LINE_WIDTH_RANGE                                           = 0x0B22
	COLOR_ATTACHMENT7                                          = 0x8CE7
	LOW_FLOAT                                                  = 0x8DF0
	FLOAT_VEC4                                                 = 0x8B52
	EQUIV                                                      = 0x1509
	LESS                                                       = 0x0201
	MAX_TEXTURE_SIZE                                           = 0x0D33
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	STENCIL_REF                                                = 0x0B97
	INVALID_OPERATION                                          = 0x0502
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	PACK_SWAP_BYTES                                            = 0x0D00
	DRAW_BUFFER1                                               = 0x8826
	DEPTH_CLAMP                                                = 0x864F
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	SRGB8_ALPHA8                                               = 0x8C43
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	SLUMINANCE8_ALPHA8_EXT                                     = 0x8C45
	PACK_SKIP_ROWS                                             = 0x0D03
	CURRENT_QUERY                                              = 0x8865
	POINT_SIZE_RANGE                                           = 0x0B12
	NOTEQUAL                                                   = 0x0205
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	SRGB_READ                                                  = 0x8297
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	TEXTURE13                                                  = 0x84CD
	TEXTURE20                                                  = 0x84D4
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	TEXTURE17                                                  = 0x84D1
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	DEPTH_RANGE                                                = 0x0B70
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	IS_PER_PATCH                                               = 0x92E7
	GREEN_INTEGER                                              = 0x8D95
	STREAM_READ                                                = 0x88E1
	GEOMETRY_SHADER                                            = 0x8DD9
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	STENCIL_INDEX                                              = 0x1901
	RGBA16_SNORM                                               = 0x8F9B
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	UNPACK_LSB_FIRST                                           = 0x0CF1
	OR_INVERTED                                                = 0x150D
	DRAW_BUFFER2                                               = 0x8827
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = TRANSFORM_FEEDBACK_PAUSED
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	STENCIL_RENDERABLE                                         = 0x8288
	TEXTURE_RED_SIZE                                           = 0x805C
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	SRGB                                                       = 0x8C40
	VIEW_CLASS_32_BITS                                         = 0x82C8
	SRC_COLOR                                                  = 0x0300
	BUFFER_ACCESS                                              = 0x88BB
	LINE_STRIP_ADJACENCY                                       = 0x000B
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	MAX_NAME_LENGTH                                            = 0x92F6
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	CLIP_DISTANCE0                                             = 0x3000
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	RG                                                         = 0x8227
	EQUAL                                                      = 0x0202
	POINT_SIZE                                                 = 0x0B11
	TRIANGLES                                                  = 0x0004
	SAMPLER_BUFFER                                             = 0x8DC2
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	RGBA_INTEGER                                               = 0x8D99
	TEXTURE16                                                  = 0x84D0
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	FRAMEBUFFER                                                = 0x8D40
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	COMPRESSED_SLUMINANCE_EXT                                  = 0x8C4A
	DOUBLE_MAT3x2                                              = 0x8F4B
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	UNSIGNED_INT                                               = 0x1405
	POLYGON_SMOOTH                                             = 0x0B41
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	RGBA8                                                      = 0x8058
	OFFSET                                                     = 0x92FC
	CLIP_DISTANCE6                                             = 0x3006
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	TEXTURE26                                                  = 0x84DA
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	DEPTH_COMPONENT24                                          = 0x81A6
	COMPUTE_SHADER                                             = 0x91B9
	TRIANGLE_FAN                                               = 0x0006
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	TEXTURE22                                                  = 0x84D6
	PACK_LSB_FIRST                                             = 0x0D01
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	R8_SNORM                                                   = 0x8F94
	TEXTURE30                                                  = 0x84DE
	ATTACHED_SHADERS                                           = 0x8B85
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	TEXTURE24                                                  = 0x84D8
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	FLOAT_MAT2x4                                               = 0x8B66
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
)
```

#### func  ActiveShaderProgram

```go
func ActiveShaderProgram(pipeline Uint, program Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glActiveShaderProgram SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glActiveShaderProgram.xml

#### func  ActiveTexture

```go
func ActiveTexture(texture Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glActiveTexture SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glActiveTexture.xml

#### func  AttachShader

```go
func AttachShader(program Uint, shader Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glAttachShader SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glAttachShader.xml

#### func  BeginConditionalRender

```go
func BeginConditionalRender(id Uint, mode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginConditionalRender SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBeginConditionalRender.xml

#### func  BeginQuery

```go
func BeginQuery(target Enum, id Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginQuery SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBeginQuery.xml

#### func  BeginQueryIndexed

```go
func BeginQueryIndexed(target Enum, index Uint, id Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginQueryIndexed SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBeginQueryIndexed.xml

#### func  BeginTransformFeedback

```go
func BeginTransformFeedback(primitiveMode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBeginTransformFeedback SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBeginTransformFeedback.xml

#### func  BindAttribLocation

```go
func BindAttribLocation(program Uint, index Uint, name *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindAttribLocation SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindAttribLocation.xml

#### func  BindBuffer

```go
func BindBuffer(target Enum, buffer Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindBuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindBuffer.xml

#### func  BindBufferBase

```go
func BindBufferBase(target Enum, index Uint, buffer Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindBufferBase SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindBufferBase.xml

#### func  BindBufferRange

```go
func BindBufferRange(target Enum, index Uint, buffer Uint, offset Intptr, size Sizeiptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindBufferRange SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindBufferRange.xml

#### func  BindFragDataLocation

```go
func BindFragDataLocation(program Uint, color Uint, name *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindFragDataLocation SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindFragDataLocation.xml

#### func  BindFragDataLocationIndexed

```go
func BindFragDataLocationIndexed(program Uint, colorNumber Uint, index Uint, name *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindFragDataLocationIndexed SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glBindFragDataLocationIndexed.xml

#### func  BindFramebuffer

```go
func BindFramebuffer(target Enum, framebuffer Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindFramebuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindFramebuffer.xml

#### func  BindImageTexture

```go
func BindImageTexture(unit Uint, texture Uint, level Int, layered Boolean, layer Int, access Enum, format Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindImageTexture SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindImageTexture.xml

#### func  BindProgramPipeline

```go
func BindProgramPipeline(pipeline Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindProgramPipeline SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindProgramPipeline.xml

#### func  BindRenderbuffer

```go
func BindRenderbuffer(target Enum, renderbuffer Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindRenderbuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindRenderbuffer.xml

#### func  BindSampler

```go
func BindSampler(unit Uint, sampler Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindSampler SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindSampler.xml

#### func  BindTexture

```go
func BindTexture(target Enum, texture Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindTexture SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindTexture.xml

#### func  BindTransformFeedback

```go
func BindTransformFeedback(target Enum, id Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindTransformFeedback SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindTransformFeedback.xml

#### func  BindVertexArray

```go
func BindVertexArray(array Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindVertexArray SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindVertexArray.xml

#### func  BindVertexBuffer

```go
func BindVertexBuffer(bindingindex Uint, buffer Uint, offset Intptr, stride Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBindVertexBuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBindVertexBuffer.xml

#### func  BlendColor

```go
func BlendColor(red Float, green Float, blue Float, alpha Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendColor SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBlendColor.xml

#### func  BlendEquation

```go
func BlendEquation(mode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquation SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquation.xml

#### func  BlendEquationSeparate

```go
func BlendEquationSeparate(modeRGB Enum, modeAlpha Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquationSeparate SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquationSeparate.xml

#### func  BlendEquationSeparatei

```go
func BlendEquationSeparatei(buf Uint, modeRGB Enum, modeAlpha Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquationSeparatei SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquationSeparatei.xml

#### func  BlendEquationi

```go
func BlendEquationi(buf Uint, mode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendEquationi SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBlendEquationi.xml

#### func  BlendFunc

```go
func BlendFunc(sfactor Enum, dfactor Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFunc SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBlendFunc.xml

#### func  BlendFuncSeparate

```go
func BlendFuncSeparate(sfactorRGB Enum, dfactorRGB Enum, sfactorAlpha Enum, dfactorAlpha Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFuncSeparate SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBlendFuncSeparate.xml

#### func  BlendFuncSeparatei

```go
func BlendFuncSeparatei(buf Uint, srcRGB Enum, dstRGB Enum, srcAlpha Enum, dstAlpha Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFuncSeparatei SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBlendFuncSeparatei.xml

#### func  BlendFunci

```go
func BlendFunci(buf Uint, src Enum, dst Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlendFunci SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBlendFunci.xml

#### func  BlitFramebuffer

```go
func BlitFramebuffer(srcX0 Int, srcY0 Int, srcX1 Int, srcY1 Int, dstX0 Int, dstY0 Int, dstX1 Int, dstY1 Int, mask Bitfield, filter Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBlitFramebuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBlitFramebuffer.xml

#### func  BufferData

```go
func BufferData(target Enum, size Sizeiptr, data Ptr, usage Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBufferData SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBufferData.xml

#### func  BufferSubData

```go
func BufferSubData(target Enum, offset Intptr, size Sizeiptr, data Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glBufferSubData SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glBufferSubData.xml

#### func  ClampColor

```go
func ClampColor(target Enum, clamp Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClampColor SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClampColor.xml

#### func  Clear

```go
func Clear(mask Bitfield)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClear SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClear.xml

#### func  ClearBufferData

```go
func ClearBufferData(target Enum, internalformat Enum, format Enum, type_ Enum, data Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferData SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferData.xml

#### func  ClearBufferSubData

```go
func ClearBufferSubData(target Enum, internalformat Enum, offset Intptr, size Sizeiptr, format Enum, type_ Enum, data Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferSubData SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferSubData.xml

#### func  ClearBufferfi

```go
func ClearBufferfi(buffer Enum, drawbuffer Int, depth Float, stencil Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferfi SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferfi.xml

#### func  ClearBufferfv

```go
func ClearBufferfv(buffer Enum, drawbuffer Int, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferfv.xml

#### func  ClearBufferiv

```go
func ClearBufferiv(buffer Enum, drawbuffer Int, value *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferiv.xml

#### func  ClearBufferuiv

```go
func ClearBufferuiv(buffer Enum, drawbuffer Int, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearBufferuiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClearBufferuiv.xml

#### func  ClearColor

```go
func ClearColor(red Float, green Float, blue Float, alpha Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearColor SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClearColor.xml

#### func  ClearDepth

```go
func ClearDepth(depth Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearDepth SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClearDepth.xml

#### func  ClearDepthf

```go
func ClearDepthf(d Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearDepthf SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClearDepthf.xml

#### func  ClearStencil

```go
func ClearStencil(s Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClearStencil SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClearStencil.xml

#### func  ColorMask

```go
func ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorMask SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glColorMask.xml

#### func  ColorMaski

```go
func ColorMaski(index Uint, r Boolean, g Boolean, b Boolean, a Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorMaski SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glColorMaski.xml

#### func  ColorP3ui

```go
func ColorP3ui(type_ Enum, color Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP3ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glColorP3ui.xml

#### func  ColorP3uiv

```go
func ColorP3uiv(type_ Enum, color *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP3uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glColorP3uiv.xml

#### func  ColorP4ui

```go
func ColorP4ui(type_ Enum, color Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP4ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glColorP4ui.xml

#### func  ColorP4uiv

```go
func ColorP4uiv(type_ Enum, color *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glColorP4uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glColorP4uiv.xml

#### func  CompileShader

```go
func CompileShader(shader Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompileShader SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCompileShader.xml

#### func  CompressedTexImage1D

```go
func CompressedTexImage1D(target Enum, level Int, internalformat Enum, width Sizei, border Int, imageSize Sizei, data Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexImage1D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage1D.xml

#### func  CompressedTexImage2D

```go
func CompressedTexImage2D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, border Int, imageSize Sizei, data Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexImage2D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage2D.xml

#### func  CompressedTexImage3D

```go
func CompressedTexImage3D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, depth Sizei, border Int, imageSize Sizei, data Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexImage3D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexImage3D.xml

#### func  CompressedTexSubImage1D

```go
func CompressedTexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, imageSize Sizei, data Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexSubImage1D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage1D.xml

#### func  CompressedTexSubImage2D

```go
func CompressedTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, imageSize Sizei, data Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexSubImage2D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage2D.xml

#### func  CompressedTexSubImage3D

```go
func CompressedTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, imageSize Sizei, data Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCompressedTexSubImage3D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCompressedTexSubImage3D.xml

#### func  CopyBufferSubData

```go
func CopyBufferSubData(readTarget Enum, writeTarget Enum, readOffset Intptr, writeOffset Intptr, size Sizeiptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyBufferSubData SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCopyBufferSubData.xml

#### func  CopyImageSubData

```go
func CopyImageSubData(srcName Uint, srcTarget Enum, srcLevel Int, srcX Int, srcY Int, srcZ Int, dstName Uint, dstTarget Enum, dstLevel Int, dstX Int, dstY Int, dstZ Int, srcWidth Sizei, srcHeight Sizei, srcDepth Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyImageSubData SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCopyImageSubData.xml

#### func  CopyTexImage1D

```go
func CopyTexImage1D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, border Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexImage1D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexImage1D.xml

#### func  CopyTexImage2D

```go
func CopyTexImage2D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, height Sizei, border Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexImage2D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexImage2D.xml

#### func  CopyTexSubImage1D

```go
func CopyTexSubImage1D(target Enum, level Int, xoffset Int, x Int, y Int, width Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexSubImage1D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage1D.xml

#### func  CopyTexSubImage2D

```go
func CopyTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, x Int, y Int, width Sizei, height Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexSubImage2D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage2D.xml

#### func  CopyTexSubImage3D

```go
func CopyTexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, x Int, y Int, width Sizei, height Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCopyTexSubImage3D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCopyTexSubImage3D.xml

#### func  CullFace

```go
func CullFace(mode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCullFace SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCullFace.xml

#### func  DebugMessageCallback

```go
func DebugMessageCallback(callback Ptr, userParam Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDebugMessageCallback SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDebugMessageCallback.xml

#### func  DebugMessageControl

```go
func DebugMessageControl(source Enum, type_ Enum, severity Enum, count Sizei, ids *Uint, enabled Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDebugMessageControl SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDebugMessageControl.xml

#### func  DebugMessageInsert

```go
func DebugMessageInsert(source Enum, type_ Enum, id Uint, severity Enum, length Sizei, buf *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDebugMessageInsert SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDebugMessageInsert.xml

#### func  DeleteBuffers

```go
func DeleteBuffers(n Sizei, buffers *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteBuffers SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteBuffers.xml

#### func  DeleteFramebuffers

```go
func DeleteFramebuffers(n Sizei, framebuffers *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteFramebuffers SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteFramebuffers.xml

#### func  DeleteProgram

```go
func DeleteProgram(program Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteProgram SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteProgram.xml

#### func  DeleteProgramPipelines

```go
func DeleteProgramPipelines(n Sizei, pipelines *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteProgramPipelines SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteProgramPipelines.xml

#### func  DeleteQueries

```go
func DeleteQueries(n Sizei, ids *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteQueries SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteQueries.xml

#### func  DeleteRenderbuffers

```go
func DeleteRenderbuffers(n Sizei, renderbuffers *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteRenderbuffers SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteRenderbuffers.xml

#### func  DeleteSamplers

```go
func DeleteSamplers(count Sizei, samplers *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteSamplers SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteSamplers.xml

#### func  DeleteShader

```go
func DeleteShader(shader Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteShader SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteShader.xml

#### func  DeleteSync

```go
func DeleteSync(sync Sync)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteSync SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteSync.xml

#### func  DeleteTextures

```go
func DeleteTextures(n Sizei, textures *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteTextures SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteTextures.xml

#### func  DeleteTransformFeedbacks

```go
func DeleteTransformFeedbacks(n Sizei, ids *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteTransformFeedbacks SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteTransformFeedbacks.xml

#### func  DeleteVertexArrays

```go
func DeleteVertexArrays(n Sizei, arrays *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDeleteVertexArrays SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDeleteVertexArrays.xml

#### func  DepthFunc

```go
func DepthFunc(func_ Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthFunc SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDepthFunc.xml

#### func  DepthMask

```go
func DepthMask(flag Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthMask SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDepthMask.xml

#### func  DepthRange

```go
func DepthRange(near_ Double, far_ Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRange SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDepthRange.xml

#### func  DepthRangeArrayv

```go
func DepthRangeArrayv(first Uint, count Sizei, v *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRangeArrayv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDepthRangeArrayv.xml

#### func  DepthRangeIndexed

```go
func DepthRangeIndexed(index Uint, n Double, f Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRangeIndexed SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDepthRangeIndexed.xml

#### func  DepthRangef

```go
func DepthRangef(n Float, f Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDepthRangef SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDepthRangef.xml

#### func  DetachShader

```go
func DetachShader(program Uint, shader Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDetachShader SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDetachShader.xml

#### func  Disable

```go
func Disable(cap Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDisable SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDisable.xml

#### func  DisableVertexAttribArray

```go
func DisableVertexAttribArray(index Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDisableVertexAttribArray SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDisableVertexAttribArray.xml

#### func  Disablei

```go
func Disablei(target Enum, index Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDisablei SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDisablei.xml

#### func  DispatchCompute

```go
func DispatchCompute(num_groups_x Uint, num_groups_y Uint, num_groups_z Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDispatchCompute SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDispatchCompute.xml

#### func  DispatchComputeIndirect

```go
func DispatchComputeIndirect(indirect Intptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDispatchComputeIndirect SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDispatchComputeIndirect.xml

#### func  DrawArrays

```go
func DrawArrays(mode Enum, first Int, count Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArrays SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawArrays.xml

#### func  DrawArraysIndirect

```go
func DrawArraysIndirect(mode Enum, indirect Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArraysIndirect SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawArraysIndirect.xml

#### func  DrawArraysInstanced

```go
func DrawArraysInstanced(mode Enum, first Int, count Sizei, instancecount Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArraysInstanced SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawArraysInstanced.xml

#### func  DrawArraysInstancedBaseInstance

```go
func DrawArraysInstancedBaseInstance(mode Enum, first Int, count Sizei, instancecount Sizei, baseinstance Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawArraysInstancedBaseInstance
SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawArraysInstancedBaseInstance.xml

#### func  DrawBuffer

```go
func DrawBuffer(mode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawBuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawBuffer.xml

#### func  DrawBuffers

```go
func DrawBuffers(n Sizei, bufs *Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawBuffers SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawBuffers.xml

#### func  DrawElements

```go
func DrawElements(mode Enum, count Sizei, type_ Enum, indices Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElements SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawElements.xml

#### func  DrawElementsBaseVertex

```go
func DrawElementsBaseVertex(mode Enum, count Sizei, type_ Enum, indices Ptr, basevertex Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsBaseVertex SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsBaseVertex.xml

#### func  DrawElementsIndirect

```go
func DrawElementsIndirect(mode Enum, type_ Enum, indirect Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsIndirect SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsIndirect.xml

#### func  DrawElementsInstanced

```go
func DrawElementsInstanced(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstanced SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstanced.xml

#### func  DrawElementsInstancedBaseInstance

```go
func DrawElementsInstancedBaseInstance(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, baseinstance Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstancedBaseInstance
SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstancedBaseInstance.xml

#### func  DrawElementsInstancedBaseVertex

```go
func DrawElementsInstancedBaseVertex(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, basevertex Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawElementsInstancedBaseVertex
SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstancedBaseVertex.xml

#### func  DrawElementsInstancedBaseVertexBaseInstance

```go
func DrawElementsInstancedBaseVertexBaseInstance(mode Enum, count Sizei, type_ Enum, indices Ptr, instancecount Sizei, basevertex Int, baseinstance Uint)
```
GLAPI Wiki:
http://www.opengl.org/wiki/GLAPI/glDrawElementsInstancedBaseVertexBaseInstance
SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawElementsInstancedBaseVertexBaseInstance.xml

#### func  DrawRangeElements

```go
func DrawRangeElements(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawRangeElements SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawRangeElements.xml

#### func  DrawRangeElementsBaseVertex

```go
func DrawRangeElementsBaseVertex(mode Enum, start Uint, end Uint, count Sizei, type_ Enum, indices Ptr, basevertex Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawRangeElementsBaseVertex SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawRangeElementsBaseVertex.xml

#### func  DrawTransformFeedback

```go
func DrawTransformFeedback(mode Enum, id Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedback SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedback.xml

#### func  DrawTransformFeedbackInstanced

```go
func DrawTransformFeedbackInstanced(mode Enum, id Uint, instancecount Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedbackInstanced
SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedbackInstanced.xml

#### func  DrawTransformFeedbackStream

```go
func DrawTransformFeedbackStream(mode Enum, id Uint, stream Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedbackStream SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedbackStream.xml

#### func  DrawTransformFeedbackStreamInstanced

```go
func DrawTransformFeedbackStreamInstanced(mode Enum, id Uint, stream Uint, instancecount Sizei)
```
GLAPI Wiki:
http://www.opengl.org/wiki/GLAPI/glDrawTransformFeedbackStreamInstanced SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glDrawTransformFeedbackStreamInstanced.xml

#### func  Enable

```go
func Enable(cap Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEnable SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glEnable.xml

#### func  EnableVertexAttribArray

```go
func EnableVertexAttribArray(index Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEnableVertexAttribArray SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glEnableVertexAttribArray.xml

#### func  Enablei

```go
func Enablei(target Enum, index Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEnablei SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glEnablei.xml

#### func  EndConditionalRender

```go
func EndConditionalRender()
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndConditionalRender SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glEndConditionalRender.xml

#### func  EndQuery

```go
func EndQuery(target Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndQuery SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glEndQuery.xml

#### func  EndQueryIndexed

```go
func EndQueryIndexed(target Enum, index Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndQueryIndexed SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glEndQueryIndexed.xml

#### func  EndTransformFeedback

```go
func EndTransformFeedback()
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glEndTransformFeedback SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glEndTransformFeedback.xml

#### func  Finish

```go
func Finish()
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFinish SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFinish.xml

#### func  Flush

```go
func Flush()
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFlush SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFlush.xml

#### func  FlushMappedBufferRange

```go
func FlushMappedBufferRange(target Enum, offset Intptr, length Sizeiptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFlushMappedBufferRange SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFlushMappedBufferRange.xml

#### func  FramebufferParameteri

```go
func FramebufferParameteri(target Enum, pname Enum, param Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferParameteri SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferParameteri.xml

#### func  FramebufferRenderbuffer

```go
func FramebufferRenderbuffer(target Enum, attachment Enum, renderbuffertarget Enum, renderbuffer Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferRenderbuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferRenderbuffer.xml

#### func  FramebufferTexture

```go
func FramebufferTexture(target Enum, attachment Enum, texture Uint, level Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture.xml

#### func  FramebufferTexture1D

```go
func FramebufferTexture1D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture1D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture1D.xml

#### func  FramebufferTexture2D

```go
func FramebufferTexture2D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture2D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture2D.xml

#### func  FramebufferTexture3D

```go
func FramebufferTexture3D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int, zoffset Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTexture3D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTexture3D.xml

#### func  FramebufferTextureLayer

```go
func FramebufferTextureLayer(target Enum, attachment Enum, texture Uint, level Int, layer Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFramebufferTextureLayer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFramebufferTextureLayer.xml

#### func  FrontFace

```go
func FrontFace(mode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFrontFace SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFrontFace.xml

#### func  GenBuffers

```go
func GenBuffers(n Sizei, buffers *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenBuffers SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGenBuffers.xml

#### func  GenFramebuffers

```go
func GenFramebuffers(n Sizei, framebuffers *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenFramebuffers SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGenFramebuffers.xml

#### func  GenProgramPipelines

```go
func GenProgramPipelines(n Sizei, pipelines *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenProgramPipelines SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGenProgramPipelines.xml

#### func  GenQueries

```go
func GenQueries(n Sizei, ids *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenQueries SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGenQueries.xml

#### func  GenRenderbuffers

```go
func GenRenderbuffers(n Sizei, renderbuffers *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenRenderbuffers SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGenRenderbuffers.xml

#### func  GenSamplers

```go
func GenSamplers(count Sizei, samplers *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenSamplers SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGenSamplers.xml

#### func  GenTextures

```go
func GenTextures(n Sizei, textures *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenTextures SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGenTextures.xml

#### func  GenTransformFeedbacks

```go
func GenTransformFeedbacks(n Sizei, ids *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenTransformFeedbacks SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGenTransformFeedbacks.xml

#### func  GenVertexArrays

```go
func GenVertexArrays(n Sizei, arrays *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenVertexArrays SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGenVertexArrays.xml

#### func  GenerateMipmap

```go
func GenerateMipmap(target Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGenerateMipmap SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGenerateMipmap.xml

#### func  GetActiveAtomicCounterBufferiv

```go
func GetActiveAtomicCounterBufferiv(program Uint, bufferIndex Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveAtomicCounterBufferiv
SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveAtomicCounterBufferiv.xml

#### func  GetActiveAttrib

```go
func GetActiveAttrib(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveAttrib SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveAttrib.xml

#### func  GetActiveSubroutineName

```go
func GetActiveSubroutineName(program Uint, shadertype Enum, index Uint, bufsize Sizei, length *Sizei, name *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveSubroutineName SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveSubroutineName.xml

#### func  GetActiveSubroutineUniformName

```go
func GetActiveSubroutineUniformName(program Uint, shadertype Enum, index Uint, bufsize Sizei, length *Sizei, name *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveSubroutineUniformName
SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveSubroutineUniformName.xml

#### func  GetActiveSubroutineUniformiv

```go
func GetActiveSubroutineUniformiv(program Uint, shadertype Enum, index Uint, pname Enum, values *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveSubroutineUniformiv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveSubroutineUniformiv.xml

#### func  GetActiveUniform

```go
func GetActiveUniform(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Int, type_ *Enum, name *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniform SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniform.xml

#### func  GetActiveUniformBlockName

```go
func GetActiveUniformBlockName(program Uint, uniformBlockIndex Uint, bufSize Sizei, length *Sizei, uniformBlockName *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformBlockName SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformBlockName.xml

#### func  GetActiveUniformBlockiv

```go
func GetActiveUniformBlockiv(program Uint, uniformBlockIndex Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformBlockiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformBlockiv.xml

#### func  GetActiveUniformName

```go
func GetActiveUniformName(program Uint, uniformIndex Uint, bufSize Sizei, length *Sizei, uniformName *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformName SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformName.xml

#### func  GetActiveUniformsiv

```go
func GetActiveUniformsiv(program Uint, uniformCount Sizei, uniformIndices *Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetActiveUniformsiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetActiveUniformsiv.xml

#### func  GetAttachedShaders

```go
func GetAttachedShaders(program Uint, maxCount Sizei, count *Sizei, obj *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetAttachedShaders SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetAttachedShaders.xml

#### func  GetBooleani_v

```go
func GetBooleani_v(target Enum, index Uint, data *Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBooleani_v SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetBooleani_v.xml

#### func  GetBooleanv

```go
func GetBooleanv(pname Enum, params *Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBooleanv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetBooleanv.xml

#### func  GetBufferParameteri64v

```go
func GetBufferParameteri64v(target Enum, pname Enum, params *Int64)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferParameteri64v SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferParameteri64v.xml

#### func  GetBufferParameteriv

```go
func GetBufferParameteriv(target Enum, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferParameteriv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferParameteriv.xml

#### func  GetBufferPointerv

```go
func GetBufferPointerv(target Enum, pname Enum, params *Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferPointerv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferPointerv.xml

#### func  GetBufferSubData

```go
func GetBufferSubData(target Enum, offset Intptr, size Sizeiptr, data Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetBufferSubData SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetBufferSubData.xml

#### func  GetCompressedTexImage

```go
func GetCompressedTexImage(target Enum, level Int, img Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetCompressedTexImage SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetCompressedTexImage.xml

#### func  GetDoublei_v

```go
func GetDoublei_v(target Enum, index Uint, data *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetDoublei_v SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetDoublei_v.xml

#### func  GetDoublev

```go
func GetDoublev(pname Enum, params *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetDoublev SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetDoublev.xml

#### func  GetFloati_v

```go
func GetFloati_v(target Enum, index Uint, data *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFloati_v SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetFloati_v.xml

#### func  GetFloatv

```go
func GetFloatv(pname Enum, params *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFloatv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetFloatv.xml

#### func  GetFramebufferAttachmentParameteriv

```go
func GetFramebufferAttachmentParameteriv(target Enum, attachment Enum, pname Enum, params *Int)
```
GLAPI Wiki:
http://www.opengl.org/wiki/GLAPI/glGetFramebufferAttachmentParameteriv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetFramebufferAttachmentParameteriv.xml

#### func  GetFramebufferParameteriv

```go
func GetFramebufferParameteriv(target Enum, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFramebufferParameteriv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetFramebufferParameteriv.xml

#### func  GetInteger64i_v

```go
func GetInteger64i_v(target Enum, index Uint, data *Int64)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInteger64i_v SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetInteger64i_v.xml

#### func  GetInteger64v

```go
func GetInteger64v(pname Enum, params *Int64)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInteger64v SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetInteger64v.xml

#### func  GetIntegeri_v

```go
func GetIntegeri_v(target Enum, index Uint, data *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetIntegeri_v SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetIntegeri_v.xml

#### func  GetIntegerv

```go
func GetIntegerv(pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetIntegerv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetIntegerv.xml

#### func  GetInternalformati64v

```go
func GetInternalformati64v(target Enum, internalformat Enum, pname Enum, bufSize Sizei, params *Int64)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInternalformati64v SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetInternalformati64v.xml

#### func  GetInternalformativ

```go
func GetInternalformativ(target Enum, internalformat Enum, pname Enum, bufSize Sizei, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetInternalformativ SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetInternalformativ.xml

#### func  GetMultisamplefv

```go
func GetMultisamplefv(pname Enum, index Uint, val *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetMultisamplefv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetMultisamplefv.xml

#### func  GetObjectLabel

```go
func GetObjectLabel(identifier Enum, name Uint, bufSize Sizei, length *Sizei, label *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetObjectLabel SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetObjectLabel.xml

#### func  GetObjectPtrLabel

```go
func GetObjectPtrLabel(ptr Ptr, bufSize Sizei, length *Sizei, label *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetObjectPtrLabel SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetObjectPtrLabel.xml

#### func  GetPointerv

```go
func GetPointerv(pname Enum, params *Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetPointerv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetPointerv.xml

#### func  GetProgramBinary

```go
func GetProgramBinary(program Uint, bufSize Sizei, length *Sizei, binaryFormat *Enum, binary Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramBinary SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramBinary.xml

#### func  GetProgramInfoLog

```go
func GetProgramInfoLog(program Uint, bufSize Sizei, length *Sizei, infoLog *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramInfoLog SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramInfoLog.xml

#### func  GetProgramInterfaceiv

```go
func GetProgramInterfaceiv(program Uint, programInterface Enum, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramInterfaceiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramInterfaceiv.xml

#### func  GetProgramPipelineInfoLog

```go
func GetProgramPipelineInfoLog(pipeline Uint, bufSize Sizei, length *Sizei, infoLog *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramPipelineInfoLog SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramPipelineInfoLog.xml

#### func  GetProgramPipelineiv

```go
func GetProgramPipelineiv(pipeline Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramPipelineiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramPipelineiv.xml

#### func  GetProgramResourceName

```go
func GetProgramResourceName(program Uint, programInterface Enum, index Uint, bufSize Sizei, length *Sizei, name *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceName SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceName.xml

#### func  GetProgramResourceiv

```go
func GetProgramResourceiv(program Uint, programInterface Enum, index Uint, propCount Sizei, props *Enum, bufSize Sizei, length *Sizei, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceiv.xml

#### func  GetProgramStageiv

```go
func GetProgramStageiv(program Uint, shadertype Enum, pname Enum, values *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramStageiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramStageiv.xml

#### func  GetProgramiv

```go
func GetProgramiv(program Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramiv.xml

#### func  GetQueryIndexediv

```go
func GetQueryIndexediv(target Enum, index Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryIndexediv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryIndexediv.xml

#### func  GetQueryObjecti64v

```go
func GetQueryObjecti64v(id Uint, pname Enum, params *Int64)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjecti64v SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjecti64v.xml

#### func  GetQueryObjectiv

```go
func GetQueryObjectiv(id Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjectiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectiv.xml

#### func  GetQueryObjectui64v

```go
func GetQueryObjectui64v(id Uint, pname Enum, params *Uint64)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjectui64v SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectui64v.xml

#### func  GetQueryObjectuiv

```go
func GetQueryObjectuiv(id Uint, pname Enum, params *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryObjectuiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryObjectuiv.xml

#### func  GetQueryiv

```go
func GetQueryiv(target Enum, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetQueryiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetQueryiv.xml

#### func  GetRenderbufferParameteriv

```go
func GetRenderbufferParameteriv(target Enum, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetRenderbufferParameteriv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetRenderbufferParameteriv.xml

#### func  GetSamplerParameterIiv

```go
func GetSamplerParameterIiv(sampler Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameterIiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameterIiv.xml

#### func  GetSamplerParameterIuiv

```go
func GetSamplerParameterIuiv(sampler Uint, pname Enum, params *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameterIuiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameterIuiv.xml

#### func  GetSamplerParameterfv

```go
func GetSamplerParameterfv(sampler Uint, pname Enum, params *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameterfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameterfv.xml

#### func  GetSamplerParameteriv

```go
func GetSamplerParameteriv(sampler Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSamplerParameteriv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetSamplerParameteriv.xml

#### func  GetShaderInfoLog

```go
func GetShaderInfoLog(shader Uint, bufSize Sizei, length *Sizei, infoLog *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderInfoLog SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderInfoLog.xml

#### func  GetShaderPrecisionFormat

```go
func GetShaderPrecisionFormat(shadertype Enum, precisiontype Enum, range_ *Int, precision *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderPrecisionFormat SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderPrecisionFormat.xml

#### func  GetShaderSource

```go
func GetShaderSource(shader Uint, bufSize Sizei, length *Sizei, source *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderSource SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderSource.xml

#### func  GetShaderiv

```go
func GetShaderiv(shader Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetShaderiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetShaderiv.xml

#### func  GetSynciv

```go
func GetSynciv(sync Sync, pname Enum, bufSize Sizei, length *Sizei, values *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSynciv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetSynciv.xml

#### func  GetTexImage

```go
func GetTexImage(target Enum, level Int, format Enum, type_ Enum, pixels Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexImage SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetTexImage.xml

#### func  GetTexLevelParameterfv

```go
func GetTexLevelParameterfv(target Enum, level Int, pname Enum, params *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexLevelParameterfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetTexLevelParameterfv.xml

#### func  GetTexLevelParameteriv

```go
func GetTexLevelParameteriv(target Enum, level Int, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexLevelParameteriv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetTexLevelParameteriv.xml

#### func  GetTexParameterIiv

```go
func GetTexParameterIiv(target Enum, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameterIiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterIiv.xml

#### func  GetTexParameterIuiv

```go
func GetTexParameterIuiv(target Enum, pname Enum, params *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameterIuiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterIuiv.xml

#### func  GetTexParameterfv

```go
func GetTexParameterfv(target Enum, pname Enum, params *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameterfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameterfv.xml

#### func  GetTexParameteriv

```go
func GetTexParameteriv(target Enum, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTexParameteriv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetTexParameteriv.xml

#### func  GetTransformFeedbackVarying

```go
func GetTransformFeedbackVarying(program Uint, index Uint, bufSize Sizei, length *Sizei, size *Sizei, type_ *Enum, name *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetTransformFeedbackVarying SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetTransformFeedbackVarying.xml

#### func  GetUniformIndices

```go
func GetUniformIndices(program Uint, uniformCount Sizei, uniformNames **Char, uniformIndices *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformIndices SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformIndices.xml

#### func  GetUniformSubroutineuiv

```go
func GetUniformSubroutineuiv(shadertype Enum, location Int, params *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformSubroutineuiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformSubroutineuiv.xml

#### func  GetUniformdv

```go
func GetUniformdv(program Uint, location Int, params *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformdv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformdv.xml

#### func  GetUniformfv

```go
func GetUniformfv(program Uint, location Int, params *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformfv.xml

#### func  GetUniformiv

```go
func GetUniformiv(program Uint, location Int, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformiv.xml

#### func  GetUniformuiv

```go
func GetUniformuiv(program Uint, location Int, params *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformuiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformuiv.xml

#### func  GetVertexAttribIiv

```go
func GetVertexAttribIiv(index Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribIiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribIiv.xml

#### func  GetVertexAttribIuiv

```go
func GetVertexAttribIuiv(index Uint, pname Enum, params *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribIuiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribIuiv.xml

#### func  GetVertexAttribLdv

```go
func GetVertexAttribLdv(index Uint, pname Enum, params *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribLdv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribLdv.xml

#### func  GetVertexAttribPointerv

```go
func GetVertexAttribPointerv(index Uint, pname Enum, pointer *Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribPointerv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribPointerv.xml

#### func  GetVertexAttribdv

```go
func GetVertexAttribdv(index Uint, pname Enum, params *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribdv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribdv.xml

#### func  GetVertexAttribfv

```go
func GetVertexAttribfv(index Uint, pname Enum, params *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribfv.xml

#### func  GetVertexAttribiv

```go
func GetVertexAttribiv(index Uint, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetVertexAttribiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetVertexAttribiv.xml

#### func  Hint

```go
func Hint(target Enum, mode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glHint SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glHint.xml

#### func  InvalidateBufferData

```go
func InvalidateBufferData(buffer Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateBufferData SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateBufferData.xml

#### func  InvalidateBufferSubData

```go
func InvalidateBufferSubData(buffer Uint, offset Intptr, length Sizeiptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateBufferSubData SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateBufferSubData.xml

#### func  InvalidateFramebuffer

```go
func InvalidateFramebuffer(target Enum, numAttachments Sizei, attachments *Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateFramebuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateFramebuffer.xml

#### func  InvalidateSubFramebuffer

```go
func InvalidateSubFramebuffer(target Enum, numAttachments Sizei, attachments *Enum, x Int, y Int, width Sizei, height Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateSubFramebuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateSubFramebuffer.xml

#### func  InvalidateTexImage

```go
func InvalidateTexImage(texture Uint, level Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateTexImage SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateTexImage.xml

#### func  InvalidateTexSubImage

```go
func InvalidateTexSubImage(texture Uint, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glInvalidateTexSubImage SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glInvalidateTexSubImage.xml

#### func  LineWidth

```go
func LineWidth(width Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glLineWidth SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glLineWidth.xml

#### func  LinkProgram

```go
func LinkProgram(program Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glLinkProgram SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glLinkProgram.xml

#### func  LogicOp

```go
func LogicOp(opcode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glLogicOp SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glLogicOp.xml

#### func  MemoryBarrier

```go
func MemoryBarrier(barriers Bitfield)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMemoryBarrier SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMemoryBarrier.xml

#### func  MinSampleShading

```go
func MinSampleShading(value Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMinSampleShading SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMinSampleShading.xml

#### func  MultiDrawArrays

```go
func MultiDrawArrays(mode Enum, first *Int, count *Sizei, drawcount Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawArrays SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawArrays.xml

#### func  MultiDrawArraysIndirect

```go
func MultiDrawArraysIndirect(mode Enum, indirect Ptr, drawcount Sizei, stride Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawArraysIndirect SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawArraysIndirect.xml

#### func  MultiDrawElements

```go
func MultiDrawElements(mode Enum, count *Sizei, type_ Enum, indices *Ptr, drawcount Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawElements SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawElements.xml

#### func  MultiDrawElementsBaseVertex

```go
func MultiDrawElementsBaseVertex(mode Enum, count *Sizei, type_ Enum, indices *Ptr, drawcount Sizei, basevertex *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawElementsBaseVertex SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawElementsBaseVertex.xml

#### func  MultiDrawElementsIndirect

```go
func MultiDrawElementsIndirect(mode Enum, type_ Enum, indirect Ptr, drawcount Sizei, stride Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiDrawElementsIndirect SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glMultiDrawElementsIndirect.xml

#### func  MultiTexCoordP1ui

```go
func MultiTexCoordP1ui(texture Enum, type_ Enum, coords Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP1ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP1ui.xml

#### func  MultiTexCoordP1uiv

```go
func MultiTexCoordP1uiv(texture Enum, type_ Enum, coords *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP1uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP1uiv.xml

#### func  MultiTexCoordP2ui

```go
func MultiTexCoordP2ui(texture Enum, type_ Enum, coords Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP2ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP2ui.xml

#### func  MultiTexCoordP2uiv

```go
func MultiTexCoordP2uiv(texture Enum, type_ Enum, coords *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP2uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP2uiv.xml

#### func  MultiTexCoordP3ui

```go
func MultiTexCoordP3ui(texture Enum, type_ Enum, coords Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP3ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP3ui.xml

#### func  MultiTexCoordP3uiv

```go
func MultiTexCoordP3uiv(texture Enum, type_ Enum, coords *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP3uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP3uiv.xml

#### func  MultiTexCoordP4ui

```go
func MultiTexCoordP4ui(texture Enum, type_ Enum, coords Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP4ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP4ui.xml

#### func  MultiTexCoordP4uiv

```go
func MultiTexCoordP4uiv(texture Enum, type_ Enum, coords *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMultiTexCoordP4uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMultiTexCoordP4uiv.xml

#### func  NormalP3ui

```go
func NormalP3ui(type_ Enum, coords Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glNormalP3ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glNormalP3ui.xml

#### func  NormalP3uiv

```go
func NormalP3uiv(type_ Enum, coords *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glNormalP3uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glNormalP3uiv.xml

#### func  ObjectLabel

```go
func ObjectLabel(identifier Enum, name Uint, length Sizei, label *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glObjectLabel SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glObjectLabel.xml

#### func  ObjectPtrLabel

```go
func ObjectPtrLabel(ptr Ptr, length Sizei, label *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glObjectPtrLabel SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glObjectPtrLabel.xml

#### func  PatchParameterfv

```go
func PatchParameterfv(pname Enum, values *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPatchParameterfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPatchParameterfv.xml

#### func  PatchParameteri

```go
func PatchParameteri(pname Enum, value Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPatchParameteri SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPatchParameteri.xml

#### func  PauseTransformFeedback

```go
func PauseTransformFeedback()
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPauseTransformFeedback SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPauseTransformFeedback.xml

#### func  PixelStoref

```go
func PixelStoref(pname Enum, param Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPixelStoref SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPixelStoref.xml

#### func  PixelStorei

```go
func PixelStorei(pname Enum, param Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPixelStorei SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPixelStorei.xml

#### func  PointParameterf

```go
func PointParameterf(pname Enum, param Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameterf SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPointParameterf.xml

#### func  PointParameterfv

```go
func PointParameterfv(pname Enum, params *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameterfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPointParameterfv.xml

#### func  PointParameteri

```go
func PointParameteri(pname Enum, param Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameteri SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPointParameteri.xml

#### func  PointParameteriv

```go
func PointParameteriv(pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointParameteriv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPointParameteriv.xml

#### func  PointSize

```go
func PointSize(size Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPointSize SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPointSize.xml

#### func  PolygonMode

```go
func PolygonMode(face Enum, mode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPolygonMode SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPolygonMode.xml

#### func  PolygonOffset

```go
func PolygonOffset(factor Float, units Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPolygonOffset SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPolygonOffset.xml

#### func  PopDebugGroup

```go
func PopDebugGroup()
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPopDebugGroup SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPopDebugGroup.xml

#### func  PrimitiveRestartIndex

```go
func PrimitiveRestartIndex(index Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPrimitiveRestartIndex SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPrimitiveRestartIndex.xml

#### func  ProgramBinary

```go
func ProgramBinary(program Uint, binaryFormat Enum, binary Ptr, length Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramBinary SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramBinary.xml

#### func  ProgramParameteri

```go
func ProgramParameteri(program Uint, pname Enum, value Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramParameteri SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramParameteri.xml

#### func  ProgramUniform1d

```go
func ProgramUniform1d(program Uint, location Int, v0 Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1d.xml

#### func  ProgramUniform1dv

```go
func ProgramUniform1dv(program Uint, location Int, count Sizei, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1dv.xml

#### func  ProgramUniform1f

```go
func ProgramUniform1f(program Uint, location Int, v0 Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1f SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1f.xml

#### func  ProgramUniform1fv

```go
func ProgramUniform1fv(program Uint, location Int, count Sizei, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1fv.xml

#### func  ProgramUniform1i

```go
func ProgramUniform1i(program Uint, location Int, v0 Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1i.xml

#### func  ProgramUniform1iv

```go
func ProgramUniform1iv(program Uint, location Int, count Sizei, value *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1iv.xml

#### func  ProgramUniform1ui

```go
func ProgramUniform1ui(program Uint, location Int, v0 Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1ui.xml

#### func  ProgramUniform1uiv

```go
func ProgramUniform1uiv(program Uint, location Int, count Sizei, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform1uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform1uiv.xml

#### func  ProgramUniform2d

```go
func ProgramUniform2d(program Uint, location Int, v0 Double, v1 Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2d.xml

#### func  ProgramUniform2dv

```go
func ProgramUniform2dv(program Uint, location Int, count Sizei, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2dv.xml

#### func  ProgramUniform2f

```go
func ProgramUniform2f(program Uint, location Int, v0 Float, v1 Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2f SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2f.xml

#### func  ProgramUniform2fv

```go
func ProgramUniform2fv(program Uint, location Int, count Sizei, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2fv.xml

#### func  ProgramUniform2i

```go
func ProgramUniform2i(program Uint, location Int, v0 Int, v1 Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2i.xml

#### func  ProgramUniform2iv

```go
func ProgramUniform2iv(program Uint, location Int, count Sizei, value *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2iv.xml

#### func  ProgramUniform2ui

```go
func ProgramUniform2ui(program Uint, location Int, v0 Uint, v1 Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2ui.xml

#### func  ProgramUniform2uiv

```go
func ProgramUniform2uiv(program Uint, location Int, count Sizei, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform2uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform2uiv.xml

#### func  ProgramUniform3d

```go
func ProgramUniform3d(program Uint, location Int, v0 Double, v1 Double, v2 Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3d.xml

#### func  ProgramUniform3dv

```go
func ProgramUniform3dv(program Uint, location Int, count Sizei, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3dv.xml

#### func  ProgramUniform3f

```go
func ProgramUniform3f(program Uint, location Int, v0 Float, v1 Float, v2 Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3f SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3f.xml

#### func  ProgramUniform3fv

```go
func ProgramUniform3fv(program Uint, location Int, count Sizei, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3fv.xml

#### func  ProgramUniform3i

```go
func ProgramUniform3i(program Uint, location Int, v0 Int, v1 Int, v2 Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3i.xml

#### func  ProgramUniform3iv

```go
func ProgramUniform3iv(program Uint, location Int, count Sizei, value *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3iv.xml

#### func  ProgramUniform3ui

```go
func ProgramUniform3ui(program Uint, location Int, v0 Uint, v1 Uint, v2 Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3ui.xml

#### func  ProgramUniform3uiv

```go
func ProgramUniform3uiv(program Uint, location Int, count Sizei, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform3uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform3uiv.xml

#### func  ProgramUniform4d

```go
func ProgramUniform4d(program Uint, location Int, v0 Double, v1 Double, v2 Double, v3 Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4d.xml

#### func  ProgramUniform4dv

```go
func ProgramUniform4dv(program Uint, location Int, count Sizei, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4dv.xml

#### func  ProgramUniform4f

```go
func ProgramUniform4f(program Uint, location Int, v0 Float, v1 Float, v2 Float, v3 Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4f SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4f.xml

#### func  ProgramUniform4fv

```go
func ProgramUniform4fv(program Uint, location Int, count Sizei, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4fv.xml

#### func  ProgramUniform4i

```go
func ProgramUniform4i(program Uint, location Int, v0 Int, v1 Int, v2 Int, v3 Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4i.xml

#### func  ProgramUniform4iv

```go
func ProgramUniform4iv(program Uint, location Int, count Sizei, value *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4iv.xml

#### func  ProgramUniform4ui

```go
func ProgramUniform4ui(program Uint, location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4ui.xml

#### func  ProgramUniform4uiv

```go
func ProgramUniform4uiv(program Uint, location Int, count Sizei, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniform4uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniform4uiv.xml

#### func  ProgramUniformMatrix2dv

```go
func ProgramUniformMatrix2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2dv.xml

#### func  ProgramUniformMatrix2fv

```go
func ProgramUniformMatrix2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2fv.xml

#### func  ProgramUniformMatrix2x3dv

```go
func ProgramUniformMatrix2x3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x3dv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x3dv.xml

#### func  ProgramUniformMatrix2x3fv

```go
func ProgramUniformMatrix2x3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x3fv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x3fv.xml

#### func  ProgramUniformMatrix2x4dv

```go
func ProgramUniformMatrix2x4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x4dv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x4dv.xml

#### func  ProgramUniformMatrix2x4fv

```go
func ProgramUniformMatrix2x4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix2x4fv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix2x4fv.xml

#### func  ProgramUniformMatrix3dv

```go
func ProgramUniformMatrix3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3dv.xml

#### func  ProgramUniformMatrix3fv

```go
func ProgramUniformMatrix3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3fv.xml

#### func  ProgramUniformMatrix3x2dv

```go
func ProgramUniformMatrix3x2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x2dv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x2dv.xml

#### func  ProgramUniformMatrix3x2fv

```go
func ProgramUniformMatrix3x2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x2fv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x2fv.xml

#### func  ProgramUniformMatrix3x4dv

```go
func ProgramUniformMatrix3x4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x4dv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x4dv.xml

#### func  ProgramUniformMatrix3x4fv

```go
func ProgramUniformMatrix3x4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix3x4fv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix3x4fv.xml

#### func  ProgramUniformMatrix4dv

```go
func ProgramUniformMatrix4dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4dv.xml

#### func  ProgramUniformMatrix4fv

```go
func ProgramUniformMatrix4fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4fv.xml

#### func  ProgramUniformMatrix4x2dv

```go
func ProgramUniformMatrix4x2dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x2dv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x2dv.xml

#### func  ProgramUniformMatrix4x2fv

```go
func ProgramUniformMatrix4x2fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x2fv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x2fv.xml

#### func  ProgramUniformMatrix4x3dv

```go
func ProgramUniformMatrix4x3dv(program Uint, location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x3dv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x3dv.xml

#### func  ProgramUniformMatrix4x3fv

```go
func ProgramUniformMatrix4x3fv(program Uint, location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProgramUniformMatrix4x3fv SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glProgramUniformMatrix4x3fv.xml

#### func  ProvokingVertex

```go
func ProvokingVertex(mode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glProvokingVertex SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glProvokingVertex.xml

#### func  PushDebugGroup

```go
func PushDebugGroup(source Enum, id Uint, length Sizei, message *Char)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glPushDebugGroup SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glPushDebugGroup.xml

#### func  QueryCounter

```go
func QueryCounter(id Uint, target Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glQueryCounter SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glQueryCounter.xml

#### func  ReadBuffer

```go
func ReadBuffer(mode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glReadBuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glReadBuffer.xml

#### func  ReadPixels

```go
func ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glReadPixels SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glReadPixels.xml

#### func  ReleaseShaderCompiler

```go
func ReleaseShaderCompiler()
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glReleaseShaderCompiler SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glReleaseShaderCompiler.xml

#### func  RenderbufferStorage

```go
func RenderbufferStorage(target Enum, internalformat Enum, width Sizei, height Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glRenderbufferStorage SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glRenderbufferStorage.xml

#### func  RenderbufferStorageMultisample

```go
func RenderbufferStorageMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glRenderbufferStorageMultisample
SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glRenderbufferStorageMultisample.xml

#### func  ResumeTransformFeedback

```go
func ResumeTransformFeedback()
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glResumeTransformFeedback SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glResumeTransformFeedback.xml

#### func  SampleCoverage

```go
func SampleCoverage(value Float, invert Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSampleCoverage SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glSampleCoverage.xml

#### func  SampleMaski

```go
func SampleMaski(index Uint, mask Bitfield)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSampleMaski SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glSampleMaski.xml

#### func  SamplerParameterIiv

```go
func SamplerParameterIiv(sampler Uint, pname Enum, param *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterIiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterIiv.xml

#### func  SamplerParameterIuiv

```go
func SamplerParameterIuiv(sampler Uint, pname Enum, param *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterIuiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterIuiv.xml

#### func  SamplerParameterf

```go
func SamplerParameterf(sampler Uint, pname Enum, param Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterf SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterf.xml

#### func  SamplerParameterfv

```go
func SamplerParameterfv(sampler Uint, pname Enum, param *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameterfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameterfv.xml

#### func  SamplerParameteri

```go
func SamplerParameteri(sampler Uint, pname Enum, param Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameteri SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameteri.xml

#### func  SamplerParameteriv

```go
func SamplerParameteriv(sampler Uint, pname Enum, param *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSamplerParameteriv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glSamplerParameteriv.xml

#### func  Scissor

```go
func Scissor(x Int, y Int, width Sizei, height Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissor SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glScissor.xml

#### func  ScissorArrayv

```go
func ScissorArrayv(first Uint, count Sizei, v *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissorArrayv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glScissorArrayv.xml

#### func  ScissorIndexed

```go
func ScissorIndexed(index Uint, left Int, bottom Int, width Sizei, height Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissorIndexed SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glScissorIndexed.xml

#### func  ScissorIndexedv

```go
func ScissorIndexedv(index Uint, v *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glScissorIndexedv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glScissorIndexedv.xml

#### func  SecondaryColorP3ui

```go
func SecondaryColorP3ui(type_ Enum, color Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSecondaryColorP3ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColorP3ui.xml

#### func  SecondaryColorP3uiv

```go
func SecondaryColorP3uiv(type_ Enum, color *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glSecondaryColorP3uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glSecondaryColorP3uiv.xml

#### func  ShaderBinary

```go
func ShaderBinary(count Sizei, shaders *Uint, binaryformat Enum, binary Ptr, length Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glShaderBinary SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glShaderBinary.xml

#### func  ShaderSource

```go
func ShaderSource(shader Uint, count Sizei, string_ **Char, length *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glShaderSource SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glShaderSource.xml

#### func  ShaderStorageBlockBinding

```go
func ShaderStorageBlockBinding(program Uint, storageBlockIndex Uint, storageBlockBinding Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glShaderStorageBlockBinding SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glShaderStorageBlockBinding.xml

#### func  StencilFunc

```go
func StencilFunc(func_ Enum, ref Int, mask Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilFunc SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glStencilFunc.xml

#### func  StencilFuncSeparate

```go
func StencilFuncSeparate(face Enum, func_ Enum, ref Int, mask Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilFuncSeparate SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glStencilFuncSeparate.xml

#### func  StencilMask

```go
func StencilMask(mask Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilMask SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glStencilMask.xml

#### func  StencilMaskSeparate

```go
func StencilMaskSeparate(face Enum, mask Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilMaskSeparate SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glStencilMaskSeparate.xml

#### func  StencilOp

```go
func StencilOp(fail Enum, zfail Enum, zpass Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilOp SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glStencilOp.xml

#### func  StencilOpSeparate

```go
func StencilOpSeparate(face Enum, sfail Enum, dpfail Enum, dppass Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glStencilOpSeparate SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glStencilOpSeparate.xml

#### func  TexBuffer

```go
func TexBuffer(target Enum, internalformat Enum, buffer Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexBuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexBuffer.xml

#### func  TexBufferRange

```go
func TexBufferRange(target Enum, internalformat Enum, buffer Uint, offset Intptr, size Sizeiptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexBufferRange SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexBufferRange.xml

#### func  TexCoordP1ui

```go
func TexCoordP1ui(type_ Enum, coords Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP1ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP1ui.xml

#### func  TexCoordP1uiv

```go
func TexCoordP1uiv(type_ Enum, coords *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP1uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP1uiv.xml

#### func  TexCoordP2ui

```go
func TexCoordP2ui(type_ Enum, coords Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP2ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP2ui.xml

#### func  TexCoordP2uiv

```go
func TexCoordP2uiv(type_ Enum, coords *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP2uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP2uiv.xml

#### func  TexCoordP3ui

```go
func TexCoordP3ui(type_ Enum, coords Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP3ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP3ui.xml

#### func  TexCoordP3uiv

```go
func TexCoordP3uiv(type_ Enum, coords *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP3uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP3uiv.xml

#### func  TexCoordP4ui

```go
func TexCoordP4ui(type_ Enum, coords Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP4ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP4ui.xml

#### func  TexCoordP4uiv

```go
func TexCoordP4uiv(type_ Enum, coords *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexCoordP4uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexCoordP4uiv.xml

#### func  TexImage1D

```go
func TexImage1D(target Enum, level Int, internalformat Int, width Sizei, border Int, format Enum, type_ Enum, pixels Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage1D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexImage1D.xml

#### func  TexImage2D

```go
func TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, type_ Enum, pixels Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage2D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexImage2D.xml

#### func  TexImage2DMultisample

```go
func TexImage2DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, fixedsamplelocations Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage2DMultisample SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexImage2DMultisample.xml

#### func  TexImage3D

```go
func TexImage3D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, depth Sizei, border Int, format Enum, type_ Enum, pixels Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage3D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexImage3D.xml

#### func  TexImage3DMultisample

```go
func TexImage3DMultisample(target Enum, samples Sizei, internalformat Int, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexImage3DMultisample SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexImage3DMultisample.xml

#### func  TexParameterIiv

```go
func TexParameterIiv(target Enum, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterIiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterIiv.xml

#### func  TexParameterIuiv

```go
func TexParameterIuiv(target Enum, pname Enum, params *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterIuiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterIuiv.xml

#### func  TexParameterf

```go
func TexParameterf(target Enum, pname Enum, param Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterf SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterf.xml

#### func  TexParameterfv

```go
func TexParameterfv(target Enum, pname Enum, params *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameterfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexParameterfv.xml

#### func  TexParameteri

```go
func TexParameteri(target Enum, pname Enum, param Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameteri SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexParameteri.xml

#### func  TexParameteriv

```go
func TexParameteriv(target Enum, pname Enum, params *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexParameteriv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexParameteriv.xml

#### func  TexStorage1D

```go
func TexStorage1D(target Enum, levels Sizei, internalformat Enum, width Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage1D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage1D.xml

#### func  TexStorage2D

```go
func TexStorage2D(target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage2D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage2D.xml

#### func  TexStorage2DMultisample

```go
func TexStorage2DMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei, fixedsamplelocations Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage2DMultisample SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage2DMultisample.xml

#### func  TexStorage3D

```go
func TexStorage3D(target Enum, levels Sizei, internalformat Enum, width Sizei, height Sizei, depth Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage3D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage3D.xml

#### func  TexStorage3DMultisample

```go
func TexStorage3DMultisample(target Enum, samples Sizei, internalformat Enum, width Sizei, height Sizei, depth Sizei, fixedsamplelocations Boolean)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexStorage3DMultisample SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexStorage3DMultisample.xml

#### func  TexSubImage1D

```go
func TexSubImage1D(target Enum, level Int, xoffset Int, width Sizei, format Enum, type_ Enum, pixels Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexSubImage1D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage1D.xml

#### func  TexSubImage2D

```go
func TexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, type_ Enum, pixels Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexSubImage2D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage2D.xml

#### func  TexSubImage3D

```go
func TexSubImage3D(target Enum, level Int, xoffset Int, yoffset Int, zoffset Int, width Sizei, height Sizei, depth Sizei, format Enum, type_ Enum, pixels Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTexSubImage3D SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTexSubImage3D.xml

#### func  TextureView

```go
func TextureView(texture Uint, target Enum, origtexture Uint, internalformat Enum, minlevel Uint, numlevels Uint, minlayer Uint, numlayers Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTextureView SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glTextureView.xml

#### func  TransformFeedbackVaryings

```go
func TransformFeedbackVaryings(program Uint, count Sizei, varyings **Char, bufferMode Enum)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glTransformFeedbackVaryings SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glTransformFeedbackVaryings.xml

#### func  Uniform1d

```go
func Uniform1d(location Int, x Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform1d.xml

#### func  Uniform1dv

```go
func Uniform1dv(location Int, count Sizei, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform1dv.xml

#### func  Uniform1f

```go
func Uniform1f(location Int, v0 Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1f SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform1f.xml

#### func  Uniform1fv

```go
func Uniform1fv(location Int, count Sizei, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform1fv.xml

#### func  Uniform1i

```go
func Uniform1i(location Int, v0 Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform1i.xml

#### func  Uniform1iv

```go
func Uniform1iv(location Int, count Sizei, value *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform1iv.xml

#### func  Uniform1ui

```go
func Uniform1ui(location Int, v0 Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform1ui.xml

#### func  Uniform1uiv

```go
func Uniform1uiv(location Int, count Sizei, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform1uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform1uiv.xml

#### func  Uniform2d

```go
func Uniform2d(location Int, x Double, y Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform2d.xml

#### func  Uniform2dv

```go
func Uniform2dv(location Int, count Sizei, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform2dv.xml

#### func  Uniform2f

```go
func Uniform2f(location Int, v0 Float, v1 Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2f SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform2f.xml

#### func  Uniform2fv

```go
func Uniform2fv(location Int, count Sizei, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform2fv.xml

#### func  Uniform2i

```go
func Uniform2i(location Int, v0 Int, v1 Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform2i.xml

#### func  Uniform2iv

```go
func Uniform2iv(location Int, count Sizei, value *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform2iv.xml

#### func  Uniform2ui

```go
func Uniform2ui(location Int, v0 Uint, v1 Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform2ui.xml

#### func  Uniform2uiv

```go
func Uniform2uiv(location Int, count Sizei, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform2uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform2uiv.xml

#### func  Uniform3d

```go
func Uniform3d(location Int, x Double, y Double, z Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform3d.xml

#### func  Uniform3dv

```go
func Uniform3dv(location Int, count Sizei, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform3dv.xml

#### func  Uniform3f

```go
func Uniform3f(location Int, v0 Float, v1 Float, v2 Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3f SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform3f.xml

#### func  Uniform3fv

```go
func Uniform3fv(location Int, count Sizei, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform3fv.xml

#### func  Uniform3i

```go
func Uniform3i(location Int, v0 Int, v1 Int, v2 Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform3i.xml

#### func  Uniform3iv

```go
func Uniform3iv(location Int, count Sizei, value *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform3iv.xml

#### func  Uniform3ui

```go
func Uniform3ui(location Int, v0 Uint, v1 Uint, v2 Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform3ui.xml

#### func  Uniform3uiv

```go
func Uniform3uiv(location Int, count Sizei, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform3uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform3uiv.xml

#### func  Uniform4d

```go
func Uniform4d(location Int, x Double, y Double, z Double, w Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform4d.xml

#### func  Uniform4dv

```go
func Uniform4dv(location Int, count Sizei, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform4dv.xml

#### func  Uniform4f

```go
func Uniform4f(location Int, v0 Float, v1 Float, v2 Float, v3 Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4f SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform4f.xml

#### func  Uniform4fv

```go
func Uniform4fv(location Int, count Sizei, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform4fv.xml

#### func  Uniform4i

```go
func Uniform4i(location Int, v0 Int, v1 Int, v2 Int, v3 Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform4i.xml

#### func  Uniform4iv

```go
func Uniform4iv(location Int, count Sizei, value *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform4iv.xml

#### func  Uniform4ui

```go
func Uniform4ui(location Int, v0 Uint, v1 Uint, v2 Uint, v3 Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform4ui.xml

#### func  Uniform4uiv

```go
func Uniform4uiv(location Int, count Sizei, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniform4uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniform4uiv.xml

#### func  UniformBlockBinding

```go
func UniformBlockBinding(program Uint, uniformBlockIndex Uint, uniformBlockBinding Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformBlockBinding SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformBlockBinding.xml

#### func  UniformMatrix2dv

```go
func UniformMatrix2dv(location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2dv.xml

#### func  UniformMatrix2fv

```go
func UniformMatrix2fv(location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2fv.xml

#### func  UniformMatrix2x3dv

```go
func UniformMatrix2x3dv(location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x3dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x3dv.xml

#### func  UniformMatrix2x3fv

```go
func UniformMatrix2x3fv(location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x3fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x3fv.xml

#### func  UniformMatrix2x4dv

```go
func UniformMatrix2x4dv(location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x4dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x4dv.xml

#### func  UniformMatrix2x4fv

```go
func UniformMatrix2x4fv(location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix2x4fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix2x4fv.xml

#### func  UniformMatrix3dv

```go
func UniformMatrix3dv(location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3dv.xml

#### func  UniformMatrix3fv

```go
func UniformMatrix3fv(location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3fv.xml

#### func  UniformMatrix3x2dv

```go
func UniformMatrix3x2dv(location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x2dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x2dv.xml

#### func  UniformMatrix3x2fv

```go
func UniformMatrix3x2fv(location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x2fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x2fv.xml

#### func  UniformMatrix3x4dv

```go
func UniformMatrix3x4dv(location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x4dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x4dv.xml

#### func  UniformMatrix3x4fv

```go
func UniformMatrix3x4fv(location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix3x4fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix3x4fv.xml

#### func  UniformMatrix4dv

```go
func UniformMatrix4dv(location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4dv.xml

#### func  UniformMatrix4fv

```go
func UniformMatrix4fv(location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4fv.xml

#### func  UniformMatrix4x2dv

```go
func UniformMatrix4x2dv(location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x2dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x2dv.xml

#### func  UniformMatrix4x2fv

```go
func UniformMatrix4x2fv(location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x2fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x2fv.xml

#### func  UniformMatrix4x3dv

```go
func UniformMatrix4x3dv(location Int, count Sizei, transpose Boolean, value *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x3dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x3dv.xml

#### func  UniformMatrix4x3fv

```go
func UniformMatrix4x3fv(location Int, count Sizei, transpose Boolean, value *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformMatrix4x3fv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformMatrix4x3fv.xml

#### func  UniformSubroutinesuiv

```go
func UniformSubroutinesuiv(shadertype Enum, count Sizei, indices *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUniformSubroutinesuiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUniformSubroutinesuiv.xml

#### func  UseProgram

```go
func UseProgram(program Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUseProgram SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUseProgram.xml

#### func  UseProgramStages

```go
func UseProgramStages(pipeline Uint, stages Bitfield, program Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUseProgramStages SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUseProgramStages.xml

#### func  ValidateProgram

```go
func ValidateProgram(program Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glValidateProgram SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glValidateProgram.xml

#### func  ValidateProgramPipeline

```go
func ValidateProgramPipeline(pipeline Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glValidateProgramPipeline SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glValidateProgramPipeline.xml

#### func  VertexAttribBinding

```go
func VertexAttribBinding(attribindex Uint, bindingindex Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribBinding SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribBinding.xml

#### func  VertexAttribDivisor

```go
func VertexAttribDivisor(index Uint, divisor Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribDivisor SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribDivisor.xml

#### func  VertexAttribFormat

```go
func VertexAttribFormat(attribindex Uint, size Int, type_ Enum, normalized Boolean, relativeoffset Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribFormat SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribFormat.xml

#### func  VertexAttribI1i

```go
func VertexAttribI1i(index Uint, x Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1i.xml

#### func  VertexAttribI1iv

```go
func VertexAttribI1iv(index Uint, v *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1iv.xml

#### func  VertexAttribI1ui

```go
func VertexAttribI1ui(index Uint, x Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1ui.xml

#### func  VertexAttribI1uiv

```go
func VertexAttribI1uiv(index Uint, v *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI1uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI1uiv.xml

#### func  VertexAttribI2i

```go
func VertexAttribI2i(index Uint, x Int, y Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2i.xml

#### func  VertexAttribI2iv

```go
func VertexAttribI2iv(index Uint, v *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2iv.xml

#### func  VertexAttribI2ui

```go
func VertexAttribI2ui(index Uint, x Uint, y Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2ui.xml

#### func  VertexAttribI2uiv

```go
func VertexAttribI2uiv(index Uint, v *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI2uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI2uiv.xml

#### func  VertexAttribI3i

```go
func VertexAttribI3i(index Uint, x Int, y Int, z Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3i.xml

#### func  VertexAttribI3iv

```go
func VertexAttribI3iv(index Uint, v *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3iv.xml

#### func  VertexAttribI3ui

```go
func VertexAttribI3ui(index Uint, x Uint, y Uint, z Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3ui.xml

#### func  VertexAttribI3uiv

```go
func VertexAttribI3uiv(index Uint, v *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI3uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI3uiv.xml

#### func  VertexAttribI4bv

```go
func VertexAttribI4bv(index Uint, v *Byte)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4bv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4bv.xml

#### func  VertexAttribI4i

```go
func VertexAttribI4i(index Uint, x Int, y Int, z Int, w Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4i SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4i.xml

#### func  VertexAttribI4iv

```go
func VertexAttribI4iv(index Uint, v *Int)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4iv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4iv.xml

#### func  VertexAttribI4sv

```go
func VertexAttribI4sv(index Uint, v *Short)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4sv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4sv.xml

#### func  VertexAttribI4ubv

```go
func VertexAttribI4ubv(index Uint, v *Ubyte)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4ubv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4ubv.xml

#### func  VertexAttribI4ui

```go
func VertexAttribI4ui(index Uint, x Uint, y Uint, z Uint, w Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4ui.xml

#### func  VertexAttribI4uiv

```go
func VertexAttribI4uiv(index Uint, v *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4uiv.xml

#### func  VertexAttribI4usv

```go
func VertexAttribI4usv(index Uint, v *Ushort)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribI4usv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribI4usv.xml

#### func  VertexAttribIFormat

```go
func VertexAttribIFormat(attribindex Uint, size Int, type_ Enum, relativeoffset Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribIFormat SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribIFormat.xml

#### func  VertexAttribIPointer

```go
func VertexAttribIPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribIPointer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribIPointer.xml

#### func  VertexAttribL1d

```go
func VertexAttribL1d(index Uint, x Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL1d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL1d.xml

#### func  VertexAttribL1dv

```go
func VertexAttribL1dv(index Uint, v *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL1dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL1dv.xml

#### func  VertexAttribL2d

```go
func VertexAttribL2d(index Uint, x Double, y Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL2d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL2d.xml

#### func  VertexAttribL2dv

```go
func VertexAttribL2dv(index Uint, v *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL2dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL2dv.xml

#### func  VertexAttribL3d

```go
func VertexAttribL3d(index Uint, x Double, y Double, z Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL3d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL3d.xml

#### func  VertexAttribL3dv

```go
func VertexAttribL3dv(index Uint, v *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL3dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL3dv.xml

#### func  VertexAttribL4d

```go
func VertexAttribL4d(index Uint, x Double, y Double, z Double, w Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL4d SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL4d.xml

#### func  VertexAttribL4dv

```go
func VertexAttribL4dv(index Uint, v *Double)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribL4dv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribL4dv.xml

#### func  VertexAttribLFormat

```go
func VertexAttribLFormat(attribindex Uint, size Int, type_ Enum, relativeoffset Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribLFormat SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribLFormat.xml

#### func  VertexAttribLPointer

```go
func VertexAttribLPointer(index Uint, size Int, type_ Enum, stride Sizei, pointer Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribLPointer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribLPointer.xml

#### func  VertexAttribP1ui

```go
func VertexAttribP1ui(index Uint, type_ Enum, normalized Boolean, value Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP1ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP1ui.xml

#### func  VertexAttribP1uiv

```go
func VertexAttribP1uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP1uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP1uiv.xml

#### func  VertexAttribP2ui

```go
func VertexAttribP2ui(index Uint, type_ Enum, normalized Boolean, value Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP2ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP2ui.xml

#### func  VertexAttribP2uiv

```go
func VertexAttribP2uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP2uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP2uiv.xml

#### func  VertexAttribP3ui

```go
func VertexAttribP3ui(index Uint, type_ Enum, normalized Boolean, value Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP3ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP3ui.xml

#### func  VertexAttribP3uiv

```go
func VertexAttribP3uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP3uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP3uiv.xml

#### func  VertexAttribP4ui

```go
func VertexAttribP4ui(index Uint, type_ Enum, normalized Boolean, value Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP4ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP4ui.xml

#### func  VertexAttribP4uiv

```go
func VertexAttribP4uiv(index Uint, type_ Enum, normalized Boolean, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribP4uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribP4uiv.xml

#### func  VertexAttribPointer

```go
func VertexAttribPointer(index Uint, size Int, type_ Enum, normalized Boolean, stride Sizei, pointer Ptr)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexAttribPointer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexAttribPointer.xml

#### func  VertexBindingDivisor

```go
func VertexBindingDivisor(bindingindex Uint, divisor Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexBindingDivisor SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexBindingDivisor.xml

#### func  VertexP2ui

```go
func VertexP2ui(type_ Enum, value Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP2ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexP2ui.xml

#### func  VertexP2uiv

```go
func VertexP2uiv(type_ Enum, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP2uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexP2uiv.xml

#### func  VertexP3ui

```go
func VertexP3ui(type_ Enum, value Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP3ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexP3ui.xml

#### func  VertexP3uiv

```go
func VertexP3uiv(type_ Enum, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP3uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexP3uiv.xml

#### func  VertexP4ui

```go
func VertexP4ui(type_ Enum, value Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP4ui SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexP4ui.xml

#### func  VertexP4uiv

```go
func VertexP4uiv(type_ Enum, value *Uint)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glVertexP4uiv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glVertexP4uiv.xml

#### func  Viewport

```go
func Viewport(x Int, y Int, width Sizei, height Sizei)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewport SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glViewport.xml

#### func  ViewportArrayv

```go
func ViewportArrayv(first Uint, count Sizei, v *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewportArrayv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glViewportArrayv.xml

#### func  ViewportIndexedf

```go
func ViewportIndexedf(index Uint, x Float, y Float, w Float, h Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewportIndexedf SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glViewportIndexedf.xml

#### func  ViewportIndexedfv

```go
func ViewportIndexedfv(index Uint, v *Float)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glViewportIndexedfv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glViewportIndexedfv.xml

#### func  WaitSync

```go
func WaitSync(sync Sync, flags Bitfield, timeout Uint64)
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glWaitSync SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glWaitSync.xml

#### type Bitfield

```go
type Bitfield C.GLbitfield
```


#### type Boolean

```go
type Boolean C.GLboolean
```


#### func  IsBuffer

```go
func IsBuffer(buffer Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsBuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsBuffer.xml

#### func  IsEnabled

```go
func IsEnabled(cap Enum) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsEnabled SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsEnabled.xml

#### func  IsEnabledi

```go
func IsEnabledi(target Enum, index Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsEnabledi SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsEnabledi.xml

#### func  IsFramebuffer

```go
func IsFramebuffer(framebuffer Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsFramebuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsFramebuffer.xml

#### func  IsProgram

```go
func IsProgram(program Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsProgram SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsProgram.xml

#### func  IsProgramPipeline

```go
func IsProgramPipeline(pipeline Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsProgramPipeline SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsProgramPipeline.xml

#### func  IsQuery

```go
func IsQuery(id Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsQuery SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsQuery.xml

#### func  IsRenderbuffer

```go
func IsRenderbuffer(renderbuffer Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsRenderbuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsRenderbuffer.xml

#### func  IsSampler

```go
func IsSampler(sampler Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsSampler SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsSampler.xml

#### func  IsShader

```go
func IsShader(shader Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsShader SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsShader.xml

#### func  IsSync

```go
func IsSync(sync Sync) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsSync SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsSync.xml

#### func  IsTexture

```go
func IsTexture(texture Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsTexture SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsTexture.xml

#### func  IsTransformFeedback

```go
func IsTransformFeedback(id Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsTransformFeedback SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsTransformFeedback.xml

#### func  IsVertexArray

```go
func IsVertexArray(array Uint) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glIsVertexArray SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glIsVertexArray.xml

#### func  UnmapBuffer

```go
func UnmapBuffer(target Enum) Boolean
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glUnmapBuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glUnmapBuffer.xml

#### type Byte

```go
type Byte C.GLbyte
```


#### type Char

```go
type Char C.GLchar
```


#### type Clampd

```go
type Clampd C.GLclampd
```


#### type Clampf

```go
type Clampf C.GLclampf
```


#### type Double

```go
type Double C.GLdouble
```


#### type Enum

```go
type Enum C.GLenum
```


#### func  CheckFramebufferStatus

```go
func CheckFramebufferStatus(target Enum) Enum
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCheckFramebufferStatus SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCheckFramebufferStatus.xml

#### func  ClientWaitSync

```go
func ClientWaitSync(sync Sync, flags Bitfield, timeout Uint64) Enum
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glClientWaitSync SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glClientWaitSync.xml

#### func  GetError

```go
func GetError() Enum
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetError SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetError.xml

#### type Float

```go
type Float C.GLfloat
```


#### type GlUtil

```go
type GlUtil struct{}
```

Provides utility methods for working with this OpenGL binding package.

```go
var Util GlUtil
```
Provides utility methods for working with this OpenGL binding package.

#### func (GlUtil) CString

```go
func (_ GlUtil) CString(str string) *Char
```
Go string to GL string.

#### func (GlUtil) CStringAlloc

```go
func (_ GlUtil) CStringAlloc(length Sizei) *Char
```
Allocates a GL string.

#### func (GlUtil) CStringArray

```go
func (_ GlUtil) CStringArray(strs ...string) []*Char
```
Converts a list of Go strings to a slice of GL strings.

#### func (GlUtil) CStringArrayFree

```go
func (_ GlUtil) CStringArrayFree(strs []*Char)
```
Free GL string slice allocated by GLStringArray().

#### func (GlUtil) CStringFree

```go
func (_ GlUtil) CStringFree(str *Char)
```
Frees GL string.

#### func (GlUtil) ErrorFlags

```go
func (_ GlUtil) ErrorFlags() (flags []Enum)
```
Returns all error flags collected from GetError() until the latter yields
NO_ERROR.

#### func (GlUtil) Init

```go
func (_ GlUtil) Init() bool
```
Initializes this OpenGL binding after having created your GL context (with SDL,
GLFW or in some other way).

#### func (GlUtil) PtrOffset

```go
func (_ GlUtil) PtrOffset(p Ptr, o uintptr) Ptr
```
Add offset to a pointer. Useful for VertexAttribPointer and friends.

#### func (GlUtil) StringFromChar

```go
func (_ GlUtil) StringFromChar(str *Char) string
```
GL string (GLchar*) to Go string.

#### func (GlUtil) StringFromUbyte

```go
func (_ GlUtil) StringFromUbyte(str *Ubyte) string
```
GL string (GLubyte*) to Go string.

#### func (GlUtil) StringN

```go
func (_ GlUtil) StringN(str *Char, length Sizei) string
```
GL string (GLchar*) with length to Go string.

#### type Half

```go
type Half C.GLhalf
```


#### type Int

```go
type Int C.GLint
```


#### func  GetAttribLocation

```go
func GetAttribLocation(program Uint, name *Char) Int
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetAttribLocation SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetAttribLocation.xml

#### func  GetFragDataIndex

```go
func GetFragDataIndex(program Uint, name *Char) Int
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFragDataIndex SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetFragDataIndex.xml

#### func  GetFragDataLocation

```go
func GetFragDataLocation(program Uint, name *Char) Int
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetFragDataLocation SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetFragDataLocation.xml

#### func  GetProgramResourceLocation

```go
func GetProgramResourceLocation(program Uint, programInterface Enum, name *Char) Int
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceLocation SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceLocation.xml

#### func  GetProgramResourceLocationIndex

```go
func GetProgramResourceLocationIndex(program Uint, programInterface Enum, name *Char) Int
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceLocationIndex
SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceLocationIndex.xml

#### func  GetSubroutineUniformLocation

```go
func GetSubroutineUniformLocation(program Uint, shadertype Enum, name *Char) Int
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSubroutineUniformLocation SDK
doc: http://www.opengl.org/sdk/docs/man/xhtml/glGetSubroutineUniformLocation.xml

#### func  GetUniformLocation

```go
func GetUniformLocation(program Uint, name *Char) Int
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformLocation SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformLocation.xml

#### type Int64

```go
type Int64 C.GLint64
```


#### type Intptr

```go
type Intptr C.GLintptr
```


#### type Ptr

```go
type Ptr unsafe.Pointer
```


#### func  MapBuffer

```go
func MapBuffer(target Enum, access Enum) *Ptr
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMapBuffer SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMapBuffer.xml

#### func  MapBufferRange

```go
func MapBufferRange(target Enum, offset Intptr, length Sizeiptr, access Bitfield) *Ptr
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glMapBufferRange SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glMapBufferRange.xml

#### type Short

```go
type Short C.GLshort
```


#### type Sizei

```go
type Sizei C.GLsizei
```


#### type Sizeiptr

```go
type Sizeiptr C.GLsizeiptr
```


#### type Sync

```go
type Sync C.GLsync
```


#### func  FenceSync

```go
func FenceSync(condition Enum, flags Bitfield) Sync
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glFenceSync SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glFenceSync.xml

#### type Ubyte

```go
type Ubyte C.GLubyte
```


#### func  GetString

```go
func GetString(name Enum) *Ubyte
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetString SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetString.xml

#### func  GetStringi

```go
func GetStringi(name Enum, index Uint) *Ubyte
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetStringi SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetStringi.xml

#### type Uint

```go
type Uint C.GLuint
```


#### func  CreateProgram

```go
func CreateProgram() Uint
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCreateProgram SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCreateProgram.xml

#### func  CreateShader

```go
func CreateShader(type_ Enum) Uint
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCreateShader SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCreateShader.xml

#### func  CreateShaderProgramv

```go
func CreateShaderProgramv(type_ Enum, count Sizei, strings **Char) Uint
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glCreateShaderProgramv SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glCreateShaderProgramv.xml

#### func  GetDebugMessageLog

```go
func GetDebugMessageLog(count Uint, bufsize Sizei, sources *Enum, types *Enum, ids *Uint, severities *Enum, lengths *Sizei, messageLog *Char) Uint
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetDebugMessageLog SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetDebugMessageLog.xml

#### func  GetProgramResourceIndex

```go
func GetProgramResourceIndex(program Uint, programInterface Enum, name *Char) Uint
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetProgramResourceIndex SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetProgramResourceIndex.xml

#### func  GetSubroutineIndex

```go
func GetSubroutineIndex(program Uint, shadertype Enum, name *Char) Uint
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetSubroutineIndex SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetSubroutineIndex.xml

#### func  GetUniformBlockIndex

```go
func GetUniformBlockIndex(program Uint, uniformBlockName *Char) Uint
```
GLAPI Wiki: http://www.opengl.org/wiki/GLAPI/glGetUniformBlockIndex SDK doc:
http://www.opengl.org/sdk/docs/man/xhtml/glGetUniformBlockIndex.xml

#### type Uint64

```go
type Uint64 C.GLuint64
```


#### type Ushort

```go
type Ushort C.GLushort
```

--
**godocdown** http://github.com/robertkrimen/godocdown
